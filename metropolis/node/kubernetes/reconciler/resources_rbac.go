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
	clusterRoleBindingOwnerAdmin             = builtinRBACName("owner-admin")
	clusterRoleCSIProvisioner                = builtinRBACName("csi-provisioner")
	clusterRoleBindingCSIProvisioners        = builtinRBACName("csi-provisioner")
	clusterRoleNetServices                   = builtinRBACName("netservices")
	clusterRoleBindingNetServices            = builtinRBACName("netservices")
)

type resourceClusterRoles struct {
	kubernetes.Interface
}

func (r resourceClusterRoles) List(ctx context.Context) ([]meta.Object, error) {
	res, err := r.RbacV1().ClusterRoles().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]meta.Object, len(res.Items))
	for i := range res.Items {
		objs[i] = &res.Items[i]
	}
	return objs, nil
}

func (r resourceClusterRoles) Create(ctx context.Context, el meta.Object) error {
	_, err := r.RbacV1().ClusterRoles().Create(ctx, el.(*rbac.ClusterRole), meta.CreateOptions{})
	return err
}

func (r resourceClusterRoles) Update(ctx context.Context, el meta.Object) error {
	_, err := r.RbacV1().ClusterRoles().Update(ctx, el.(*rbac.ClusterRole), meta.UpdateOptions{})
	return err
}

func (r resourceClusterRoles) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
	return r.RbacV1().ClusterRoles().Delete(ctx, name, opts)
}

func (r resourceClusterRoles) Expected() []meta.Object {
	return []meta.Object{
		&rbac.ClusterRole{
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
		&rbac.ClusterRole{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRoleCSIProvisioner,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This role grants access to PersistentVolumes, PersistentVolumeClaims and StorageClassses, as used by the CSI provisioner running on nodes.",
				},
			},
			Rules: []rbac.PolicyRule{
				{
					APIGroups: []string{""},
					Resources: []string{"events"},
					Verbs:     []string{"get", "list", "watch", "create", "update", "patch"},
				},
				{
					APIGroups: []string{"storage.k8s.io"},
					Resources: []string{"storageclasses"},
					Verbs:     []string{"get", "list", "watch"},
				},
				{
					APIGroups: []string{""},
					Resources: []string{"persistentvolumes", "persistentvolumeclaims"},
					Verbs:     []string{"*"},
				},
			},
		},
		&rbac.ClusterRole{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRoleNetServices,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This role grants access to the minimum set of resources that are needed to run networking services for a node.",
				},
			},
			Rules: []rbac.PolicyRule{
				{
					APIGroups: []string{"discovery.k8s.io"},
					Resources: []string{"endpointslices"},
					Verbs:     []string{"get", "list", "watch"},
				},
				{
					APIGroups: []string{""},
					Resources: []string{"services", "nodes", "namespaces"},
					Verbs:     []string{"get", "list", "watch"},
				},
			},
		},
	}
}

type resourceClusterRoleBindings struct {
	kubernetes.Interface
}

func (r resourceClusterRoleBindings) List(ctx context.Context) ([]meta.Object, error) {
	res, err := r.RbacV1().ClusterRoleBindings().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]meta.Object, len(res.Items))
	for i := range res.Items {
		objs[i] = &res.Items[i]
	}
	return objs, nil
}

func (r resourceClusterRoleBindings) Create(ctx context.Context, el meta.Object) error {
	_, err := r.RbacV1().ClusterRoleBindings().Create(ctx, el.(*rbac.ClusterRoleBinding), meta.CreateOptions{})
	return err
}

func (r resourceClusterRoleBindings) Update(ctx context.Context, el meta.Object) error {
	_, err := r.RbacV1().ClusterRoleBindings().Update(ctx, el.(*rbac.ClusterRoleBinding), meta.UpdateOptions{})
	return err
}

func (r resourceClusterRoleBindings) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
	return r.RbacV1().ClusterRoleBindings().Delete(ctx, name, opts)
}

func (r resourceClusterRoleBindings) Expected() []meta.Object {
	return []meta.Object{
		&rbac.ClusterRoleBinding{
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
		&rbac.ClusterRoleBinding{
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
					Name: "metropolis:apiserver-kubelet-client",
				},
			},
		},
		&rbac.ClusterRoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRoleBindingOwnerAdmin,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This binding grants the Metropolis Cluster owner access to the " +
						"cluster-admin role on Kubernetes.",
				},
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     "cluster-admin",
			},
			Subjects: []rbac.Subject{
				{
					APIGroup: rbac.GroupName,
					Kind:     "User",
					Name:     "owner",
				},
			},
		},
		&rbac.ClusterRoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRoleBindingCSIProvisioners,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This role binding grants CSI provisioners running on nodes access to the necessary resources.",
				},
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     clusterRoleCSIProvisioner,
			},
			Subjects: []rbac.Subject{
				{
					APIGroup: rbac.GroupName,
					Kind:     "Group",
					Name:     "metropolis:csi-provisioner",
				},
			},
		},
		&rbac.ClusterRoleBinding{
			ObjectMeta: meta.ObjectMeta{
				Name:   clusterRoleBindingNetServices,
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"kubernetes.io/description": "This role binding grants node network services access to necessary resources.",
				},
			},
			RoleRef: rbac.RoleRef{
				APIGroup: rbac.GroupName,
				Kind:     "ClusterRole",
				Name:     clusterRoleNetServices,
			},
			Subjects: []rbac.Subject{
				{
					APIGroup: rbac.GroupName,
					Kind:     "Group",
					Name:     "metropolis:netservices",
				},
			},
		},
	}
}
