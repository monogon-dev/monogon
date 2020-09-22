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
package reconciler

import (
	"context"
	"fmt"
	"time"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

// Sad workaround for all the pointer booleans in K8s specs
func True() *bool {
	val := true
	return &val
}
func False() *bool {
	val := false
	return &val
}

const (
	// BuiltinLabelKey is used as a k8s label to mark built-in objects (ie., managed by the reconciler)
	BuiltinLabelKey = "smalltown.com/builtin"
	// BuiltinLabelValue is used as a k8s label value, under the BuiltinLabelKey key.
	BuiltinLabelValue = "true"
	// BuiltinRBACPrefix is used to prefix all built-in objects that are part of the rbac/v1 API (eg.
	// {Cluster,}Role{Binding,} objects). This corresponds to the colon-separated 'namespaces' notation used by
	// Kubernetes system (system:) objects.
	BuiltinRBACPrefix = "smalltown:"
)

// builtinLabels makes a kubernetes-compatible label dictionary (key->value) that is used to mark objects that are
// built-in into Smalltown (ie., managed by the reconciler). These are then subsequently retrieved by listBuiltins.
// The extra argument specifies what other labels are to be merged into the the labels dictionary, for convenience. If
// nil or empty, no extra labels will be applied.
func builtinLabels(extra map[string]string) map[string]string {
	l := map[string]string{
		BuiltinLabelKey: BuiltinLabelValue,
	}
	if extra != nil {
		for k, v := range extra {
			l[k] = v
		}
	}
	return l
}

// listBuiltins returns a k8s client ListOptions structure that allows to retrieve all objects that are built-in into
// Smalltown currently present in the API server (ie., ones that are to be managed by the reconciler). These are created
// by applying builtinLabels to their metadata labels.
var listBuiltins = meta.ListOptions{
	LabelSelector: fmt.Sprintf("%s=%s", BuiltinLabelKey, BuiltinLabelValue),
}

// builtinRBACName returns a name that is compatible with colon-delimited 'namespaced' objects, a la system:*.
// These names are to be used by all builtins created as part of the rbac/v1 Kubernetes API.
func builtinRBACName(name string) string {
	return BuiltinRBACPrefix + name
}

// resource is a type of resource to be managed by the reconciler. All builti-ins/reconciled objects must implement
// this interface to be managed correctly by the reconciler.
type resource interface {
	// List returns a list of names of objects current present on the target (ie. k8s API server).
	List(ctx context.Context) ([]string, error)
	// Create creates an object on the target. The el interface{} argument is the black box object returned by the
	// Expected() call.
	Create(ctx context.Context, el interface{}) error
	// Delete delete an object, by name, from the target.
	Delete(ctx context.Context, name string) error
	// Expected returns a map of all objects expected to be present on the target. The keys are names (which must
	// correspond to the names returned by List() and used by Delete(), and the values are blackboxes that will then
	// be passed to the Create() call if their corresponding key (name) does not exist on the target.
	Expected() map[string]interface{}
}

func allResources(clientSet kubernetes.Interface) map[string]resource {
	return map[string]resource{
		"psps":                resourcePodSecurityPolicies{clientSet},
		"clusterroles":        resourceClusterRoles{clientSet},
		"clusterrolebindings": resourceClusterRoleBindings{clientSet},
		"storageclasses":      resourceStorageClasses{clientSet},
		"csidrivers":          resourceCSIDrivers{clientSet},
		"runtimeclasses":      resourceRuntimeClasses{clientSet},
	}
}

func Run(clientSet kubernetes.Interface) supervisor.Runnable {
	return func(ctx context.Context) error {
		log := supervisor.Logger(ctx)
		resources := allResources(clientSet)
		t := time.NewTicker(10 * time.Second)
		reconcileAll := func() {
			for name, resource := range resources {
				if err := reconcile(ctx, resource); err != nil {
					log.Warn("Failed to reconcile built-in resources", zap.String("kind", name), zap.Error(err))
				}
			}
		}
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		reconcileAll()
		for {
			select {
			case <-t.C:
				reconcileAll()
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func reconcile(ctx context.Context, r resource) error {
	present, err := r.List(ctx)
	if err != nil {
		return err
	}
	presentSet := make(map[string]bool)
	for _, el := range present {
		presentSet[el] = true
	}
	expectedMap := r.Expected()
	for name, el := range expectedMap {
		if !presentSet[name] {
			if err := r.Create(ctx, el); err != nil {
				return err
			}
		}
	}
	for name, _ := range presentSet {
		if _, ok := expectedMap[name]; !ok {
			if err := r.Delete(ctx, name); err != nil {
				return err
			}
		}
	}
	return nil
}
