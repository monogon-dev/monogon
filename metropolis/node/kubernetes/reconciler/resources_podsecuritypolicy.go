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

package reconciler

import (
	"context"

	core "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type resourcePodSecurityPolicies struct {
	kubernetes.Interface
}

func (r resourcePodSecurityPolicies) List(ctx context.Context) ([]meta.Object, error) {
	res, err := r.PolicyV1beta1().PodSecurityPolicies().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]meta.Object, len(res.Items))
	for i := range res.Items {
		objs[i] = &res.Items[i]
	}
	return objs, nil
}

func (r resourcePodSecurityPolicies) Create(ctx context.Context, el meta.Object) error {
	_, err := r.PolicyV1beta1().PodSecurityPolicies().Create(ctx, el.(*policy.PodSecurityPolicy), meta.CreateOptions{})
	return err
}

func (r resourcePodSecurityPolicies) Delete(ctx context.Context, name string) error {
	return r.PolicyV1beta1().PodSecurityPolicies().Delete(ctx, name, meta.DeleteOptions{})
}

func (r resourcePodSecurityPolicies) Expected() []meta.Object {
	return []meta.Object{
		&policy.PodSecurityPolicy{
			ObjectMeta: meta.ObjectMeta{
				Name:   "default",
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This default PSP allows the creation of pods using features that are" +
						" generally considered safe against any sort of escape.",
				},
			},
			Spec: policy.PodSecurityPolicySpec{
				AllowPrivilegeEscalation: True(),
				AllowedCapabilities: []core.Capability{ // runc's default list of allowed capabilities
					"SETPCAP",
					"MKNOD",
					"AUDIT_WRITE",
					"CHOWN",
					"NET_RAW",
					"DAC_OVERRIDE",
					"FOWNER",
					"FSETID",
					"KILL",
					"SETGID",
					"SETUID",
					"NET_BIND_SERVICE",
					"SYS_CHROOT",
					"SETFCAP",
				},
				HostNetwork: false,
				HostIPC:     false,
				HostPID:     false,
				FSGroup: policy.FSGroupStrategyOptions{
					Rule: policy.FSGroupStrategyRunAsAny,
				},
				RunAsUser: policy.RunAsUserStrategyOptions{
					Rule: policy.RunAsUserStrategyRunAsAny,
				},
				SELinux: policy.SELinuxStrategyOptions{
					Rule: policy.SELinuxStrategyRunAsAny,
				},
				SupplementalGroups: policy.SupplementalGroupsStrategyOptions{
					Rule: policy.SupplementalGroupsStrategyRunAsAny,
				},
				Volumes: []policy.FSType{ // Volumes considered safe to use
					policy.ConfigMap,
					policy.EmptyDir,
					policy.Projected,
					policy.Secret,
					policy.DownwardAPI,
					policy.PersistentVolumeClaim,
				},
			},
		},
	}
}
