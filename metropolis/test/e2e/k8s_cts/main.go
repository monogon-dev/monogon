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

// This package launches a Metropolis cluster with two nodes and spawns in the CTS container. Then it streams its output
// to the console. When the CTS has finished it exits with the appropriate error code.
package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/test/e2e"
	"source.monogon.dev/metropolis/test/launch"
)

// makeCTSPodSpec generates a spec for a standalone pod running the Kubernetes CTS. It also sets the test configuration
// for the Kubernetes E2E test suite to only run CTS tests and excludes known-broken ones.
func makeCTSPodSpec(name string, saName string) *corev1.Pod {
	skipRegexes := []string{
		// hostNetworking cannot be supported since we run different network stacks for the host and containers
		"should function for node-pod communication",
		// gVisor misreports statfs() syscalls: https://github.com/google/gvisor/issues/3339
		`should support \((non-)?root,`,
		"volume on tmpfs should have the correct mode",
		"volume on default medium should have the correct mode",
		// gVisor doesn't support the full Linux privilege machinery including SUID and NewPrivs
		// https://github.com/google/gvisor/issues/189#issuecomment-481064000
		"should run the container as unprivileged when false",
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"name": name,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "cts",
					Image: "bazel/metropolis/test/e2e/k8s_cts:k8s_cts_image",
					Args: []string{
						"-cluster-ip-range=10.0.0.0/17",
						"-dump-systemd-journal=false",
						"-ginkgo.focus=\\[Conformance\\]",
						"-ginkgo.skip=" + strings.Join(skipRegexes, "|"),
						"-test.parallel=8",
					},
					ImagePullPolicy: corev1.PullNever,
				},
			},
			Tolerations: []corev1.Toleration{{ // Tolerate all taints, otherwise the CTS likes to self-evict
				Operator: "Exists",
			}},
			PriorityClassName:  "system-cluster-critical", // Don't evict the CTS pod
			RestartPolicy:      corev1.RestartPolicyNever,
			ServiceAccountName: saName,
		},
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-sigs
		log.Printf("Got signal %s, aborting test", sig)
		cancel()
	}()

	debugClient, portMap, err := launch.LaunchCluster(ctx, launch.ClusterOptions{NumNodes: 2})
	if err != nil {
		log.Fatalf("Failed to launch cluster: %v", err)
	}
	log.Println("Cluster initialized")
	clientSet, err := e2e.GetKubeClientSet(ctx, debugClient, portMap[common.KubernetesAPIPort])
	if err != nil {
		log.Fatalf("Failed to get clientSet: %v", err)
	}
	log.Println("Credentials available")

	saName := "cts"
	ctsSA := &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: saName}}
	for {
		if _, err := clientSet.CoreV1().ServiceAccounts("default").Create(ctx, ctsSA, metav1.CreateOptions{}); err != nil {
			log.Printf("Failed to create ServiceAccount: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	ctsRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: saName,
		},
		Subjects: []rbacv1.Subject{
			{
				Namespace: "default",
				Name:      saName,
				Kind:      rbacv1.ServiceAccountKind,
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind: "ClusterRole",
			Name: "cluster-admin",
		},
	}
	podName := "cts"
	if _, err := clientSet.RbacV1().ClusterRoleBindings().Create(ctx, ctsRoleBinding, metav1.CreateOptions{}); err != nil {
		log.Fatalf("Failed to create ClusterRoleBinding: %v", err)
	}
	for {
		if _, err := clientSet.CoreV1().Pods("default").Create(ctx, makeCTSPodSpec(podName, saName), metav1.CreateOptions{}); err != nil {
			log.Printf("Failed to create Pod: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	var logs io.ReadCloser
	go func() {
		// This loops the whole .Stream()/io.Copy process because the API sometimes returns streams that immediately return EOF
		for {
			logs, err = clientSet.CoreV1().Pods("default").GetLogs(podName, &corev1.PodLogOptions{Follow: true}).Stream(ctx)
			if err == nil {
				if _, err := io.Copy(os.Stdout, logs); err != nil {
					log.Printf("Log pump error: %v", err)
				}
				logs.Close()
			} else if err == ctx.Err() {
				return // Exit if the context has been cancelled
			} else {
				log.Printf("Pod logs not ready yet: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		time.Sleep(1 * time.Second)
		pod, err := clientSet.CoreV1().Pods("default").Get(ctx, podName, metav1.GetOptions{})
		if err != nil && err == ctx.Err() {
			return // Exit if the context has been cancelled
		} else if err != nil {
			log.Printf("Failed to get CTS pod: %v", err)
			continue
		}
		if pod.Status.Phase == corev1.PodSucceeded {
			return
		}
		if pod.Status.Phase == corev1.PodFailed {
			log.Fatalf("CTS failed")
		}
	}
}
