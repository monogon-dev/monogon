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
	"crypto/x509"
	"encoding/pem"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"source.monogon.dev/metropolis/test/launch/cluster"
)

// GetKubeClientSet gets a Kubernetes client set accessing the Metropolis
// Kubernetes authenticating proxy using the cluster owner identity.
// It currently has access to everything (i.e. the cluster-admin role)
// via the owner-admin binding.
func GetKubeClientSet(cluster *cluster.Cluster, port uint16) (kubernetes.Interface, error) {
	pkcs8Key, err := x509.MarshalPKCS8PrivateKey(cluster.Owner.PrivateKey)
	if err != nil {
		// We explicitly pass an Ed25519 private key in, so this can't happen
		panic(err)
	}
	var clientConfig = rest.Config{
		Host: fmt.Sprintf("localhost:%v", port),
		TLSClientConfig: rest.TLSClientConfig{
			ServerName: "kubernetes.default.svc.cluster.local",
			Insecure:   true,
			CertData:   pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cluster.Owner.Certificate[0]}),
			KeyData:    pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: pkcs8Key}),
		},
	}
	return kubernetes.NewForConfig(&clientConfig)
}

// makeTestDeploymentSpec generates a Deployment spec for a single pod running
// NGINX with a readiness probe. This allows verifying that the control plane
// is capable of scheduling simple pods and that kubelet works, its runtime is
// set up well enough to run a simple container and the network to the pod can
// pass readiness probe traffic.
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
								ProbeHandler: corev1.ProbeHandler{
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
func makeTestStatefulSet(name string, volumeMode corev1.PersistentVolumeMode) *appsv1.StatefulSet {
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
						VolumeMode: &volumeMode,
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
								ProbeHandler: corev1.ProbeHandler{
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
