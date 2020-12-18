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

	rbac "k8s.io/api/rbac/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	clusterRolePSPDefault                    = builtinRBACName("psp-default")
	clusterRoleBindingDefaultPSP             = builtinRBACName("default-psp-for-sa")
	clusterRoleBindingAPIServerKubeletClient = builtinRBACName("apiserver-kubelet-client")
)

type resourceClusterRoles struct {
	kubernetes.Interface
}

func (r resourceClusterRoles) List(ctx context.Context) ([]string, error) {
	res, err := r.RbacV1().ClusterRoles().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]string, len(res.Items))
	for i, el := range res.Items {
		objs[i] = el.ObjectMeta.Name
	}
	return objs, nil
}

func (r resourceClusterRoles) Create(ctx context.Context, el interface{}) error {
	_, err := r.RbacV1().ClusterRoles().Create(ctx, el.(*rbac.ClusterRole), meta.CreateOptions{})
	return err
}

func (r resourceClusterRoles) Delete(ctx context.Context, name string) error {
	return r.RbacV1().ClusterRoles().Delete(ctx, name, meta.DeleteOptions{})
}

func (r resourceClusterRoles) Expected() map[string]interface{} {
	return map[string]interface{}{
		clusterRolePSPDefault: &rbac.ClusterRole{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRolePSPDefault,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This role grants access to the \"default\" PSP.",
				},
			},
			Rules: []rbac.PolicyRule{
				{
					APIGroups:     []string{"policy"},
					Resources:     []string{"podsecuritypolicies"},
					ResourceNames: []string{"default"},
					Verbs:         []string{"use"},
				},
			},
		},
	}
}

type resourceClusterRoleBindings struct {
	kubernetes.Interface
}

func (r resourceClusterRoleBindings) List(ctx context.Context) ([]string, error) {
	res, err := r.RbacV1().ClusterRoleBindings().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]string, len(res.Items))
	for i, el := range res.Items {
		objs[i] = el.ObjectMeta.Name
	}
	return objs, nil
}

func (r resourceClusterRoleBindings) Create(ctx context.Context, el interface{}) error {
	_, err := r.RbacV1().ClusterRoleBindings().Create(ctx, el.(*rbac.ClusterRoleBinding), meta.CreateOptions{})
	return err
}

func (r resourceClusterRoleBindings) Delete(ctx context.Context, name string) error {
	return r.RbacV1().ClusterRoleBindings().Delete(ctx, name, meta.DeleteOptions{})
}

func (r resourceClusterRoleBindings) Expected() map[string]interface{} {
	return map[string]interface{}{
		clusterRoleBindingDefaultPSP: &rbac.ClusterRoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRoleBindingDefaultPSP,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This binding grants every service account access to the \"default\" PSP. " +
						"Creation of Pods is still restricted by other RBAC roles. Otherwise no pods (unprivileged or not) " +
						"can be created.",
				},
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     clusterRolePSPDefault,
			},
			Subjects: []rbac.Subject{
				{
					APIGroup: rbac.GroupName,
					Kind:     "Group",
					Name:     "system:serviceaccounts",
				},
			},
		},
		clusterRoleBindingAPIServerKubeletClient: &rbac.ClusterRoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRoleBindingAPIServerKubeletClient,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This binding grants the apiserver access to the kubelets. This enables " +
						"lots of built-in functionality like reading logs or forwarding ports via the API.",
				},
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     "system:kubelet-api-admin",
			},
			Subjects: []rbac.Subject{
				{
					APIGroup: rbac.GroupName,
					Kind:     "User",
					// TODO(q3k): describe this name's contract, or unify with whatever creates this.
					Name: "smalltown:apiserver-kubelet-client",
				},
			},
		},
	}
}
