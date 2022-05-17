// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	podv1 "k8s.io/kubernetes/pkg/api/v1/pod"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	apb "source.monogon.dev/metropolis/proto/api"
	"source.monogon.dev/metropolis/test/launch/cluster"
)

const (
	// Timeout for the global test context.
	//
	// Bazel would eventually time out the test after 900s ("large") if, for
	// some reason, the context cancellation fails to abort it.
	globalTestTimeout = 600 * time.Second

	// Timeouts for individual end-to-end tests of different sizes.
	smallTestTimeout = 60 * time.Second
	largeTestTimeout = 120 * time.Second
)

// TestE2E is the main E2E test entrypoint for single-node freshly-bootstrapped
// E2E tests. It starts a full Metropolis node in bootstrap mode and then runs
// tests against it. The actual tests it performs are located in the RunGroup
// subtest.
func TestE2E(t *testing.T) {
	// Run pprof server for debugging
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	pprofListen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on pprof port: %s", pprofListen.Addr())
	}

	log.Printf("E2E: pprof server listening on %s", pprofListen.Addr())
	go func() {
		log.Printf("E2E: pprof server returned an error: %v", http.Serve(pprofListen, nil))
		pprofListen.Close()
	}()

	// Set a global timeout to make sure this terminates
	ctx, cancel := context.WithTimeout(context.Background(), globalTestTimeout)
	defer cancel()

	// Launch cluster.
	clusterOptions := cluster.ClusterOptions{
		NumNodes: 2,
	}
	cluster, err := cluster.LaunchCluster(ctx, clusterOptions)
	if err != nil {
		t.Fatalf("LaunchCluster failed: %v", err)
	}
	defer func() {
		err := cluster.Close()
		if err != nil {
			t.Fatalf("cluster Close failed: %v", err)
		}
	}()

	log.Printf("E2E: Cluster running, starting tests...")

	// Dial first node's curator.
	creds := rpc.NewAuthenticatedCredentials(cluster.Owner, nil)
	remote := net.JoinHostPort(cluster.NodeIDs[0], common.CuratorServicePort.PortString())
	cl, err := grpc.Dial(remote, grpc.WithContextDialer(cluster.DialNode), grpc.WithTransportCredentials(creds))
	if err != nil {
		t.Fatalf("failed to dial first node's curator: %v", err)
	}
	defer cl.Close()
	mgmt := apb.NewManagementClient(cl)

	// This exists to keep the parent around while all the children race.
	// It currently tests both a set of OS-level conditions and Kubernetes
	// Deployments and StatefulSets
	t.Run("RunGroup", func(t *testing.T) {
		t.Run("Cluster", func(t *testing.T) {
			testEventual(t, "Retrieving cluster directory sucessful", ctx, 60*time.Second, func(ctx context.Context) error {
				res, err := mgmt.GetClusterInfo(ctx, &apb.GetClusterInfoRequest{})
				if err != nil {
					return fmt.Errorf("GetClusterInfo: %w", err)
				}

				// Ensure that the expected node count is present.
				nodes := res.ClusterDirectory.Nodes
				if want, got := clusterOptions.NumNodes, len(nodes); want != got {
					return fmt.Errorf("wanted %d nodes in cluster directory, got %d", want, got)
				}

				// Ensure the nodes have the expected addresses.
				addresses := make(map[string]bool)
				for _, n := range nodes {
					if len(n.Addresses) != 1 {
						return fmt.Errorf("node %s has no addresss", identity.NodeID(n.PublicKey))
					}
					address := n.Addresses[0].Host
					addresses[address] = true
				}

				for _, address := range []string{"10.1.0.2", "10.1.0.3"} {
					if !addresses[address] {
						return fmt.Errorf("address %q not found in directory", address)
					}
				}
				return nil
			})
			testEventual(t, "Node rejoin successful", ctx, 60*time.Second, func(ctx context.Context) error {
				// Ensure nodes rejoin the cluster after a reboot by reboting the 1st node.
				if err := cluster.RebootNode(ctx, 1); err != nil {
					return fmt.Errorf("while rebooting a node: %w", err)
				}
				return nil
			})
			testEventual(t, "Heartbeat test successful", ctx, 60*time.Second, func(ctx context.Context) error {
				// Ensure all cluster nodes are capable of sending heartbeat updates.
				// This test assumes the expected count of nodes is already present in
				// the cluster.
				for {
					srvN, err := mgmt.GetNodes(ctx, &apb.GetNodesRequest{})
					if err != nil {
						return fmt.Errorf("GetNodes: %w", err)
					}

					// Count the unhealthy nodes.
					var unhealthy int
					for {
						node, err := srvN.Recv()
						if err == io.EOF {
							break
						}
						if err != nil {
							return fmt.Errorf("GetNodes.Recv: %w", err)
						}

						if node.Health != apb.Node_HEALTHY {
							unhealthy++
						}
					}

					// If all nodes tested in this iteration are healthy, the test has
					// been passed.
					if unhealthy == 0 {
						break
					}
				}
				return nil
			})
		})
		t.Run("Kubernetes", func(t *testing.T) {
			t.Parallel()
			// TODO(q3k): use SOCKS proxy.
			clientSet, err := GetKubeClientSet(cluster, cluster.Ports[uint16(common.KubernetesAPIWrappedPort)])
			if err != nil {
				t.Fatal(err)
			}
			testEventual(t, "Nodes are registered and ready", ctx, largeTestTimeout, func(ctx context.Context) error {
				nodes, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
				if err != nil {
					return err
				}
				if len(nodes.Items) < 2 {
					return errors.New("nodes not yet registered")
				}
				node := nodes.Items[0]
				for _, cond := range node.Status.Conditions {
					if cond.Type != corev1.NodeReady {
						continue
					}
					if cond.Status != corev1.ConditionTrue {
						return fmt.Errorf("node not ready: %v", cond.Message)
					}
				}
				return nil
			})
			testEventual(t, "Simple deployment", ctx, largeTestTimeout, func(ctx context.Context) error {
				_, err := clientSet.AppsV1().Deployments("default").Create(ctx, makeTestDeploymentSpec("test-deploy-1"), metav1.CreateOptions{})
				return err
			})
			testEventual(t, "Simple deployment is running", ctx, largeTestTimeout, func(ctx context.Context) error {
				res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-deploy-1"})
				if err != nil {
					return err
				}
				if len(res.Items) == 0 {
					return errors.New("pod didn't get created")
				}
				pod := res.Items[0]
				if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
					return nil
				}
				events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
				if err != nil || len(events.Items) == 0 {
					return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
				} else {
					return fmt.Errorf("pod is not ready: %v", events.Items[0].Message)
				}
			})
			testEventual(t, "Simple deployment with runc", ctx, largeTestTimeout, func(ctx context.Context) error {
				deployment := makeTestDeploymentSpec("test-deploy-2")
				var runcStr = "runc"
				deployment.Spec.Template.Spec.RuntimeClassName = &runcStr
				_, err := clientSet.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
				return err
			})
			testEventual(t, "Simple deployment is running on runc", ctx, largeTestTimeout, func(ctx context.Context) error {
				res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-deploy-2"})
				if err != nil {
					return err
				}
				if len(res.Items) == 0 {
					return errors.New("pod didn't get created")
				}
				pod := res.Items[0]
				if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
					return nil
				}
				events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
				if err != nil || len(events.Items) == 0 {
					return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
				} else {
					var errorMsg strings.Builder
					for _, msg := range events.Items {
						errorMsg.WriteString(" | ")
						errorMsg.WriteString(msg.Message)
					}
					return fmt.Errorf("pod is not ready: %v", errorMsg.String())
				}
			})
			testEventual(t, "Simple StatefulSet with PVC", ctx, largeTestTimeout, func(ctx context.Context) error {
				_, err := clientSet.AppsV1().StatefulSets("default").Create(ctx, makeTestStatefulSet("test-statefulset-1", corev1.PersistentVolumeFilesystem), metav1.CreateOptions{})
				return err
			})
			testEventual(t, "Simple StatefulSet with PVC is running", ctx, largeTestTimeout, func(ctx context.Context) error {
				res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-statefulset-1"})
				if err != nil {
					return err
				}
				if len(res.Items) == 0 {
					return errors.New("pod didn't get created")
				}
				pod := res.Items[0]
				if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
					return nil
				}
				events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
				if err != nil || len(events.Items) == 0 {
					return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
				} else {
					return fmt.Errorf("pod is not ready: %v", events.Items[0].Message)
				}
			})
			testEventual(t, "Simple StatefulSet with Block PVC", ctx, largeTestTimeout, func(ctx context.Context) error {
				_, err := clientSet.AppsV1().StatefulSets("default").Create(ctx, makeTestStatefulSet("test-statefulset-2", corev1.PersistentVolumeBlock), metav1.CreateOptions{})
				return err
			})
			testEventual(t, "Simple StatefulSet with Block PVC is running", ctx, largeTestTimeout, func(ctx context.Context) error {
				res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-statefulset-2"})
				if err != nil {
					return err
				}
				if len(res.Items) == 0 {
					return errors.New("pod didn't get created")
				}
				pod := res.Items[0]
				if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
					return nil
				}
				events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
				if err != nil || len(events.Items) == 0 {
					return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
				} else {
					return fmt.Errorf("pod is not ready: %v", events.Items[0].Message)
				}
			})
			testEventual(t, "Pod with preseeded image", ctx, smallTestTimeout, func(ctx context.Context) error {
				_, err := clientSet.CoreV1().Pods("default").Create(ctx, &corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name: "preseed-test-1",
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{{
							Name:            "preseed-test-1",
							ImagePullPolicy: corev1.PullNever,
							Image:           "bazel/metropolis/test/e2e/preseedtest:preseedtest_image",
						}},
						RestartPolicy: corev1.RestartPolicyNever,
					},
				}, metav1.CreateOptions{})
				return err
			})
			testEventual(t, "Pod with preseeded image is completed", ctx, largeTestTimeout, func(ctx context.Context) error {
				pod, err := clientSet.CoreV1().Pods("default").Get(ctx, "preseed-test-1", metav1.GetOptions{})
				if err != nil {
					return fmt.Errorf("failed to get pod: %w", err)
				}
				if pod.Status.Phase == corev1.PodSucceeded {
					return nil
				}
				events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
				if err != nil || len(events.Items) == 0 {
					return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
				} else {
					return fmt.Errorf("pod is not ready: %v", events.Items[len(events.Items)-1].Message)
				}
			})
			if os.Getenv("HAVE_NESTED_KVM") != "" {
				testEventual(t, "Pod for KVM/QEMU smoke test", ctx, smallTestTimeout, func(ctx context.Context) error {
					runcRuntimeClass := "runc"
					_, err := clientSet.CoreV1().Pods("default").Create(ctx, &corev1.Pod{
						ObjectMeta: metav1.ObjectMeta{
							Name: "vm-smoketest",
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Name:            "vm-smoketest",
								ImagePullPolicy: corev1.PullNever,
								Image:           "bazel/metropolis/vm/smoketest:smoketest_container",
								Resources: corev1.ResourceRequirements{
									Limits: corev1.ResourceList{
										"devices.monogon.dev/kvm": *resource.NewQuantity(1, ""),
									},
								},
							}},
							RuntimeClassName: &runcRuntimeClass,
							RestartPolicy:    corev1.RestartPolicyNever,
						},
					}, metav1.CreateOptions{})
					return err
				})
				testEventual(t, "KVM/QEMU smoke test completion", ctx, smallTestTimeout, func(ctx context.Context) error {
					pod, err := clientSet.CoreV1().Pods("default").Get(ctx, "vm-smoketest", metav1.GetOptions{})
					if err != nil {
						return fmt.Errorf("failed to get pod: %w", err)
					}
					if pod.Status.Phase == corev1.PodSucceeded {
						return nil
					}
					events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
					if err != nil || len(events.Items) == 0 {
						return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
					} else {
						return fmt.Errorf("pod is not ready: %v", events.Items[len(events.Items)-1].Message)
					}
				})
			}
		})
	})
}
