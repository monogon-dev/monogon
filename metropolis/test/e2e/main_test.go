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
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	_ "net/http"
	_ "net/http/pprof"
	"net/url"
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
	"source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/metropolis/test/launch/cluster"
	"source.monogon.dev/metropolis/test/util"
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

// TestE2ECore exercisees the core functionality of Metropolis: maintaining a
// control plane, changing node roles, ...
//
// The tests are performed against an in-memory cluster.
func TestE2ECore(t *testing.T) {
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

	launch.Log("E2E: Cluster running, starting tests...")

	// Dial first node's curator.
	creds := rpc.NewAuthenticatedCredentials(cluster.Owner, rpc.WantInsecure())
	remote := net.JoinHostPort(cluster.NodeIDs[0], common.CuratorServicePort.PortString())
	cl, err := grpc.Dial(remote, grpc.WithContextDialer(cluster.DialNode), grpc.WithTransportCredentials(creds))
	if err != nil {
		t.Fatalf("failed to dial first node's curator: %v", err)
	}
	defer cl.Close()
	mgmt := apb.NewManagementClient(cl)

	util.TestEventual(t, "Retrieving cluster directory sucessful", ctx, 60*time.Second, func(ctx context.Context) error {
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
	util.TestEventual(t, "Heartbeat test successful", ctx, 20*time.Second, cluster.AllNodesHealthy)
	util.TestEventual(t, "Node rejoin successful", ctx, 60*time.Second, func(ctx context.Context) error {
		// Ensure nodes rejoin the cluster after a reboot by reboting the 1st node.
		if err := cluster.RebootNode(ctx, 1); err != nil {
			return fmt.Errorf("while rebooting a node: %w", err)
		}
		return nil
	})
	util.TestEventual(t, "Heartbeat test successful", ctx, 20*time.Second, cluster.AllNodesHealthy)
	util.TestEventual(t, "Prometheus node metrics retrieved", ctx, smallTestTimeout, func(ctx context.Context) error {
		pool := x509.NewCertPool()
		pool.AddCert(cluster.CACertificate)
		cl := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates: []tls.Certificate{cluster.Owner},
					RootCAs:      pool,
				},
				DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
					return cluster.DialNode(ctx, addr)
				},
			},
		}
		u := url.URL{
			Scheme: "https",
			Host:   net.JoinHostPort(cluster.NodeIDs[0], common.MetricsPort.PortString()),
			Path:   "/metrics/node",
		}
		res, err := cl.Get(u.String())
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return fmt.Errorf("status code %d", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		needle := "node_uname_info"
		if !strings.Contains(string(body), needle) {
			return util.Permanent(fmt.Errorf("could not find %q in returned response", needle))
		}
		return nil
	})
}

// TestE2ECore exercisees the Kubernetes functionality of Metropolis.
//
// The tests are performed against an in-memory cluster.
func TestE2EKubernetes(t *testing.T) {
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

	clientSet, err := cluster.GetKubeClientSet()
	if err != nil {
		t.Fatal(err)
	}
	util.TestEventual(t, "Add KubernetesWorker roles", ctx, smallTestTimeout, func(ctx context.Context) error {
		// Make everything but the first node into KubernetesWorkers.
		for i := 1; i < clusterOptions.NumNodes; i++ {
			err := cluster.MakeKubernetesWorker(ctx, cluster.NodeIDs[i])
			if err != nil {
				return util.Permanent(fmt.Errorf("MakeKubernetesWorker: %w", err))
			}
		}
		return nil
	})
	util.TestEventual(t, "Node is registered and ready", ctx, largeTestTimeout, func(ctx context.Context) error {
		nodes, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		if err != nil {
			return err
		}
		if len(nodes.Items) < 1 {
			return errors.New("node not yet registered")
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
	util.TestEventual(t, "Simple deployment", ctx, largeTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.AppsV1().Deployments("default").Create(ctx, makeTestDeploymentSpec("test-deploy-1"), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple deployment is running", ctx, largeTestTimeout, func(ctx context.Context) error {
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
	util.TestEventual(t, "Simple deployment with gvisor", ctx, largeTestTimeout, func(ctx context.Context) error {
		deployment := makeTestDeploymentSpec("test-deploy-2")
		gvisorStr := "gvisor"
		deployment.Spec.Template.Spec.RuntimeClassName = &gvisorStr
		_, err := clientSet.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple deployment is running on gvisor", ctx, largeTestTimeout, func(ctx context.Context) error {
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
	util.TestEventual(t, "Simple StatefulSet with PVC", ctx, largeTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.AppsV1().StatefulSets("default").Create(ctx, makeTestStatefulSet("test-statefulset-1", corev1.PersistentVolumeFilesystem), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple StatefulSet with PVC is running", ctx, largeTestTimeout, func(ctx context.Context) error {
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
	util.TestEventual(t, "Simple StatefulSet with Block PVC", ctx, largeTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.AppsV1().StatefulSets("default").Create(ctx, makeTestStatefulSet("test-statefulset-2", corev1.PersistentVolumeBlock), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple StatefulSet with Block PVC is running", ctx, largeTestTimeout, func(ctx context.Context) error {
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
	util.TestEventual(t, "In-cluster self-test job", ctx, smallTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.BatchV1().Jobs("default").Create(ctx, makeSelftestSpec("selftest"), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "In-cluster self-test job passed", ctx, smallTestTimeout, func(ctx context.Context) error {
		res, err := clientSet.BatchV1().Jobs("default").Get(ctx, "selftest", metav1.GetOptions{})
		if err != nil {
			return err
		}
		if res.Status.Failed > 0 {
			pods, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
				LabelSelector: "job-name=selftest",
			})
			if err != nil {
				return util.Permanent(fmt.Errorf("job failed but failed to find pod: %w", err))
			}
			if len(pods.Items) < 1 {
				return fmt.Errorf("job failed but pod does not exist")
			}
			lines, err := getPodLogLines(ctx, clientSet, pods.Items[0].Name, 1)
			if err != nil {
				return fmt.Errorf("job failed but could not get logs: %w", err)
			}
			if len(lines) > 0 {
				return util.Permanent(fmt.Errorf("job failed, last log line: %s", lines[0]))
			}
			return util.Permanent(fmt.Errorf("job failed, empty log"))
		}
		if res.Status.Succeeded > 0 {
			return nil
		}
		return fmt.Errorf("job still running")
	})
	if os.Getenv("HAVE_NESTED_KVM") != "" {
		util.TestEventual(t, "Pod for KVM/QEMU smoke test", ctx, smallTestTimeout, func(ctx context.Context) error {
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
		util.TestEventual(t, "KVM/QEMU smoke test completion", ctx, smallTestTimeout, func(ctx context.Context) error {
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
}
