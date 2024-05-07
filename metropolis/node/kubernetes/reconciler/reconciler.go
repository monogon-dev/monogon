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

// The reconciler ensures that a base set of K8s resources is always available
// in the cluster. These are necessary to ensure correct out-of-the-box
// functionality. All resources containing the
// metropolis.monogon.dev/builtin=true label are assumed to be managed by the
// reconciler.
// It currently does not revert modifications made by admins, it is  planned to
// create an admission plugin prohibiting such modifications to resources with
// the metropolis.monogon.dev/builtin label to deal with that problem. This
// would also solve a potential issue where you could delete resources just by
// adding the metropolis.monogon.dev/builtin=true label.
package reconciler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"source.monogon.dev/osbase/supervisor"
)

// True is a sad workaround for all the pointer booleans in K8s specs
func True() *bool {
	val := true
	return &val
}
func False() *bool {
	val := false
	return &val
}

const (
	// BuiltinLabelKey is used as a k8s label to mark built-in objects (ie.,
	// managed by the reconciler)
	BuiltinLabelKey = "metropolis.monogon.dev/builtin"
	// BuiltinLabelValue is used as a k8s label value, under the
	// BuiltinLabelKey key.
	BuiltinLabelValue = "true"
	// BuiltinRBACPrefix is used to prefix all built-in objects that are part
	// of the rbac/v1 API (eg.  {Cluster,}Role{Binding,} objects). This
	// corresponds to the colon-separated 'namespaces' notation used by
	// Kubernetes system (system:) objects.
	BuiltinRBACPrefix = "metropolis:"
)

// builtinLabels makes a kubernetes-compatible label dictionary (key->value)
// that is used to mark objects that are built-in into Metropolis (ie., managed
// by the reconciler). These are then subsequently retrieved by listBuiltins.
// The extra argument specifies what other labels are to be merged into the the
// labels dictionary, for convenience. If nil or empty, no extra labels will be
// applied.
func builtinLabels(extra map[string]string) map[string]string {
	l := map[string]string{
		BuiltinLabelKey: BuiltinLabelValue,
	}
	for k, v := range extra {
		l[k] = v
	}
	return l
}

// listBuiltins returns a k8s client ListOptions structure that allows to
// retrieve all objects that are built-in into Metropolis currently present in
// the API server (ie., ones that are to be managed by the reconciler). These
// are created by applying builtinLabels to their metadata labels.
var listBuiltins = meta.ListOptions{
	LabelSelector: fmt.Sprintf("%s=%s", BuiltinLabelKey, BuiltinLabelValue),
}

// builtinRBACName returns a name that is compatible with colon-delimited
// 'namespaced' objects, a la system:*.
// These names are to be used by all builtins created as part of the rbac/v1
// Kubernetes API.
func builtinRBACName(name string) string {
	return BuiltinRBACPrefix + name
}

// resource is a type of resource to be managed by the reconciler. All
// built-ins/reconciled objects must implement this interface to be managed
// correctly by the reconciler.
type resource interface {
	// List returns a list of objects currently present on the target
	// (ie. k8s API server).
	List(ctx context.Context) ([]meta.Object, error)
	// Create creates an object on the target. The el argument is
	// an object returned by the Expected() call.
	Create(ctx context.Context, el meta.Object) error
	// Update updates an existing object, by name, on the target.
	// The el argument is an object returned by the Expected() call.
	Update(ctx context.Context, el meta.Object) error
	// Delete deletes an object, by name, from the target.
	Delete(ctx context.Context, name string, opts meta.DeleteOptions) error
	// Expected returns a list of all objects expected to be present on the
	// target. Objects are identified by their name, as returned by GetName.
	Expected() []meta.Object
}

func allResources(clientSet kubernetes.Interface) map[string]resource {
	return map[string]resource{
		"clusterroles":        resourceClusterRoles{clientSet},
		"clusterrolebindings": resourceClusterRoleBindings{clientSet},
		"storageclasses":      resourceStorageClasses{clientSet},
		"csidrivers":          resourceCSIDrivers{clientSet},
		"runtimeclasses":      resourceRuntimeClasses{clientSet},
	}
}

func reconcileAll(ctx context.Context, clientSet kubernetes.Interface) error {
	resources := allResources(clientSet)
	for name, resource := range resources {
		err := reconcile(ctx, resource, name)
		if err != nil {
			return fmt.Errorf("resource %s: %w", name, err)
		}
	}
	return nil
}

func reconcile(ctx context.Context, r resource, rname string) error {
	log := supervisor.Logger(ctx)
	present, err := r.List(ctx)
	if err != nil {
		return err
	}
	presentMap := make(map[string]meta.Object)
	for _, el := range present {
		presentMap[el.GetName()] = el
	}
	expected := r.Expected()
	expectedMap := make(map[string]meta.Object)
	for _, el := range expected {
		expectedMap[el.GetName()] = el
	}
	for name, expectedEl := range expectedMap {
		if presentEl, ok := presentMap[name]; ok {
			// The object already exists. Update it if it is different than expected.

			// The server rejects updates which don't have an up to date ResourceVersion.
			expectedEl.SetResourceVersion(presentEl.GetResourceVersion())

			// Clear out fields set by the server, such that comparison succeeds if
			// there are no other changes.
			presentEl.SetUID("")
			presentEl.SetGeneration(0)
			presentEl.SetCreationTimestamp(meta.Time{})
			presentEl.SetManagedFields(nil)

			if !apiequality.Semantic.DeepEqual(presentEl, expectedEl) {
				log.Infof("Updating %s object %q", rname, name)
				if err := r.Update(ctx, expectedEl); err != nil {
					if !isImmutableError(err) {
						return err
					}
					log.Infof("Failed to update object due to immutable fields; deleting and recreating: %v", err)

					// Only delete if the ResourceVersion has not changed. If it has
					// changed, that means another reconciler was faster than us and
					// has already recreated the object.
					resourceVersion := presentEl.GetResourceVersion()
					deleteOpts := meta.DeleteOptions{
						Preconditions: &meta.Preconditions{
							ResourceVersion: &resourceVersion,
						},
					}
					// ResourceVersion must be cleared when creating.
					expectedEl.SetResourceVersion("")

					if err := r.Delete(ctx, name, deleteOpts); err != nil {
						return err
					}
					if err := r.Create(ctx, expectedEl); err != nil {
						return err
					}
				}
			}
		} else {
			log.Infof("Creating %s object %q", rname, name)
			if err := r.Create(ctx, expectedEl); err != nil {
				return err
			}
		}
	}
	for name := range presentMap {
		if _, ok := expectedMap[name]; !ok {
			log.Infof("Deleting %s object %q", rname, name)
			if err := r.Delete(ctx, name, meta.DeleteOptions{}); err != nil {
				return err
			}
		}
	}
	return nil
}

// isImmutableError returns true if err indicates that an update failed because
// of an attempt to update one or more immutable fields.
func isImmutableError(err error) bool {
	if !apierrors.IsInvalid(err) {
		return false
	}
	var status apierrors.APIStatus
	if !errors.As(err, &status) {
		return false
	}
	details := status.Status().Details
	if details == nil || len(details.Causes) == 0 {
		return false
	}
	for _, cause := range details.Causes {
		if !strings.Contains(cause.Message, apivalidation.FieldImmutableErrorMsg) {
			return false
		}
	}
	return true
}
