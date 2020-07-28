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
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	apb "git.monogon.dev/source/nexantic.git/core/proto/api"
)

// GetKubeClientSet gets a Kubeconfig from the debug API and creates a K8s ClientSet using it. The identity used has
// the system:masters group and thus has RBAC access to everything.
func GetKubeClientSet(ctx context.Context, client apb.NodeDebugServiceClient, port uint16) (kubernetes.Interface, error) {
	var lastErr = errors.New("context canceled before any operation completed")
	for {
		reqT, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		res, err := client.GetDebugKubeconfig(reqT, &apb.GetDebugKubeconfigRequest{Id: "debug-user", Groups: []string{"system:masters"}})
		if err == nil {
			rawClientConfig, err := clientcmd.NewClientConfigFromBytes([]byte(res.DebugKubeconfig))
			if err != nil {
				return nil, err // Invalid Kubeconfigs are immediately fatal
			}

			clientConfig, err := rawClientConfig.ClientConfig()
			clientConfig.Host = fmt.Sprintf("localhost:%v", port)
			clientSet, err := kubernetes.NewForConfig(clientConfig)
			if err != nil {
				return nil, err
			}
			return clientSet, nil
		}
		if err != nil && err == ctx.Err() {
			return nil, lastErr
		}
		lastErr = err
		select {
		case <-ctx.Done():
			return nil, lastErr
		case <-time.After(1 * time.Second):
		}
	}
}

// makeTestDeploymentSpec generates a Deployment spec for a single pod running NGINX with a readiness probe. This allows
// verifying that the control plane is capable of scheduling simple pods and that kubelet works, its runtime is set up
// well enough to run a simple container and the network to the pod can pass readiness probe traffic.
func makeTestDeploymentSpec(name string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{MatchLabels: map[string]string{
				"name": name,
			}},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: "test",
							// TODO(phab/T793): Build and preseed our own container images
							Image: "nginx:alpine",
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromInt(80)},
								},
							},
						},
					},
				},
			},
		},
	}
}

// makeTestStatefulSet generates a StatefulSet spec
func makeTestStatefulSet(name string) *appsv1.StatefulSet {
	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: appsv1.StatefulSetSpec{
			Selector: &metav1.LabelSelector{MatchLabels: map[string]string{
				"name": name,
			}},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "www"},
					Spec: corev1.PersistentVolumeClaimSpec{
						AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
						Resources: corev1.ResourceRequirements{
							Requests: map[corev1.ResourceName]resource.Quantity{corev1.ResourceStorage: resource.MustParse("50Mi")},
						},
					},
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test",
							Image: "nginx:alpine",
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromInt(80)},
								},
							},
						},
					},
				},
			},
		},
	}
}
