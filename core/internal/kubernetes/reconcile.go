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

// The reconciler ensures that a base set of K8s resources is always available in the cluster. These are necessary to
// ensure correct out-of-the-box functionality. All resources containing the smalltown.com/builtin=true label are assumed
// to be managed by the reconciler.
// It currently does not revert modifications made by admins, it is  planned to create an admission plugin prohibiting
// such modifications to resources with the smalltown.com/builtin label to deal with that problem. This would also solve a
// potential issue where you could delete resources just by adding the smalltown.com/builtin=true label.
package kubernetes

import (
	"context"
	"time"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const builtinRBACPrefix = "smalltown:"

// Sad workaround for all the pointer booleans in K8s specs
func True() *bool {
	val := true
	return &val
}
func False() *bool {
	val := false
	return &val
}

func rbac(name string) string {
	return builtinRBACPrefix + name
}

// Extended from https://github.com/kubernetes/kubernetes/blob/master/cluster/gce/addons/podsecuritypolicies/unprivileged-addon.yaml
var builtinPSPs = []*v1beta1.PodSecurityPolicy{
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
			Labels: map[string]string{
				"smalltown.com/builtin": "true",
			},
			Annotations: map[string]string{
				"kubernetes.io/description": "This default PSP allows the creation of pods using features that are" +
					" generally considered safe against any sort of escape.",
			},
		},
		Spec: v1beta1.PodSecurityPolicySpec{
			AllowPrivilegeEscalation: True(),
			AllowedCapabilities: []corev1.Capability{ // runc's default list of allowed capabilities
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
			FSGroup: v1beta1.FSGroupStrategyOptions{
				Rule: v1beta1.FSGroupStrategyRunAsAny,
			},
			RunAsUser: v1beta1.RunAsUserStrategyOptions{
				Rule: v1beta1.RunAsUserStrategyRunAsAny,
			},
			SELinux: v1beta1.SELinuxStrategyOptions{
				Rule: v1beta1.SELinuxStrategyRunAsAny,
			},
			SupplementalGroups: v1beta1.SupplementalGroupsStrategyOptions{
				Rule: v1beta1.SupplementalGroupsStrategyRunAsAny,
			},
			Volumes: []v1beta1.FSType{ // Volumes considered safe to use
				v1beta1.ConfigMap,
				v1beta1.EmptyDir,
				v1beta1.Projected,
				v1beta1.Secret,
				v1beta1.DownwardAPI,
				v1beta1.PersistentVolumeClaim,
			},
		},
	},
}

var builtinClusterRoles = []*rbacv1.ClusterRole{
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: rbac("psp-default"),
			Annotations: map[string]string{
				"kubernetes.io/description": "This role grants access to the \"default\" PSP.",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups:     []string{"policy"},
				Resources:     []string{"podsecuritypolicies"},
				ResourceNames: []string{"default"},
				Verbs:         []string{"use"},
			},
		},
	},
}

var builtinClusterRoleBindings = []*rbacv1.ClusterRoleBinding{
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: rbac("default-psp-for-sa"),
			Annotations: map[string]string{
				"kubernetes.io/description": "This binding grants every service account access to the \"default\" PSP. " +
					"Creation of Pods is still restricted by other RBAC roles. Otherwise no pods (unprivileged or not) " +
					"can be created.",
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "ClusterRole",
			Name:     rbac("psp-default"),
		},
		Subjects: []rbacv1.Subject{
			{
				APIGroup: rbacv1.GroupName,
				Kind:     "Group",
				Name:     "system:serviceaccounts",
			},
		},
	},
	{
		ObjectMeta: metav1.ObjectMeta{
			Name: rbac("apiserver-kubelet-client"),
			Annotations: map[string]string{
				"kubernetes.io/description": "This binding grants the apiserver access to the kubelets. This enables " +
					"lots of built-in functionality like reading logs or forwarding ports via the API.",
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "ClusterRole",
			Name:     "system:kubelet-api-admin",
		},
		Subjects: []rbacv1.Subject{
			{
				APIGroup: rbacv1.GroupName,
				Kind:     "User",
				Name:     "smalltown:apiserver-kubelet-client",
			},
		},
	},
}

func runReconciler(ctx context.Context, masterKubeconfig []byte, log *zap.Logger) error {
	rawClientConfig, err := clientcmd.NewClientConfigFromBytes(masterKubeconfig)
	if err != nil {
		return err
	}

	clientConfig, err := rawClientConfig.ClientConfig()
	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return err
	}
	t := time.NewTicker(10 * time.Second)
	for {
		err = reconcile(ctx, clientset)
		select {
		case <-t.C:
			err = reconcile(ctx, clientset)
			if err != nil {
				log.Warn("Failed to reconcile built-in resources", zap.Error(err))
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func reconcile(ctx context.Context, clientset *kubernetes.Clientset) error {
	if err := reconcilePSPs(ctx, clientset); err != nil {
		return err
	}
	if err := reconcileClusterRoles(ctx, clientset); err != nil {
		return err
	}
	if err := reconcileClusterRoleBindings(ctx, clientset); err != nil {
		return err
	}
	return nil
}

func reconcilePSPs(ctx context.Context, clientset *kubernetes.Clientset) error {
	pspClient := clientset.PolicyV1beta1().PodSecurityPolicies()
	availablePSPs, err := pspClient.List(ctx, metav1.ListOptions{
		LabelSelector: "smalltown.com/builtin=true",
	})
	if err != nil {
		return err
	}
	availablePSPMap := make(map[string]struct{})
	for _, psp := range availablePSPs.Items {
		availablePSPMap[psp.Name] = struct{}{}
	}
	expectedPSPMap := make(map[string]*v1beta1.PodSecurityPolicy)
	for _, psp := range builtinPSPs {
		expectedPSPMap[psp.Name] = psp
	}
	for pspName, psp := range expectedPSPMap {
		if _, ok := availablePSPMap[pspName]; !ok {
			if _, err := pspClient.Create(ctx, psp, metav1.CreateOptions{}); err != nil {
				return err
			}
		}
	}
	for pspName, _ := range availablePSPMap {
		if _, ok := expectedPSPMap[pspName]; !ok {
			if err := pspClient.Delete(ctx, pspName, metav1.DeleteOptions{}); err != nil {
				return err
			}
		}
	}
	return nil
}

func reconcileClusterRoles(ctx context.Context, clientset *kubernetes.Clientset) error {
	crClient := clientset.RbacV1().ClusterRoles()
	availableCRs, err := crClient.List(ctx, metav1.ListOptions{
		LabelSelector: "smalltown.com/builtin=true",
	})
	if err != nil {
		return err
	}
	availableCRMap := make(map[string]struct{})
	for _, cr := range availableCRs.Items {
		availableCRMap[cr.Name] = struct{}{}
	}
	expectedCRMap := make(map[string]*rbacv1.ClusterRole)
	for _, cr := range builtinClusterRoles {
		expectedCRMap[cr.Name] = cr
	}
	for crName, psp := range expectedCRMap {
		if _, ok := availableCRMap[crName]; !ok {
			if _, err := crClient.Create(ctx, psp, metav1.CreateOptions{}); err != nil {
				return err
			}
		}
	}
	for crName, _ := range availableCRMap {
		if _, ok := expectedCRMap[crName]; !ok {
			if err := crClient.Delete(ctx, crName, metav1.DeleteOptions{}); err != nil {
				return err
			}
		}
	}
	return nil
}

func reconcileClusterRoleBindings(ctx context.Context, clientset *kubernetes.Clientset) error {
	crbClient := clientset.RbacV1().ClusterRoleBindings()
	availableCRBs, err := crbClient.List(ctx, metav1.ListOptions{
		LabelSelector: "smalltown.com/builtin=true",
	})
	if err != nil {
		return err
	}
	availableCRBMap := make(map[string]struct{})
	for _, crb := range availableCRBs.Items {
		availableCRBMap[crb.Name] = struct{}{}
	}
	expectedCRBMap := make(map[string]*rbacv1.ClusterRoleBinding)
	for _, crb := range builtinClusterRoleBindings {
		expectedCRBMap[crb.Name] = crb
	}
	for crbName, psp := range expectedCRBMap {
		if _, ok := availableCRBMap[crbName]; !ok {
			if _, err := crbClient.Create(ctx, psp, metav1.CreateOptions{}); err != nil {
				return err
			}
		}
	}
	for crbName, _ := range availableCRBMap {
		if _, ok := expectedCRBMap[crbName]; !ok {
			if err := crbClient.Delete(ctx, crbName, metav1.DeleteOptions{}); err != nil {
				return err
			}
		}
	}
	return nil
}
