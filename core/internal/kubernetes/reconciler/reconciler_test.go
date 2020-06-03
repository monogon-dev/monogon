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
	"fmt"
	"testing"

	policy "k8s.io/api/policy/v1beta1"
	rbac "k8s.io/api/rbac/v1"
	storage "k8s.io/api/storage/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// kubernetesMeta unwraps an interface{} that might contain a Kubernetes resource of type that is managed by the
// reconciler. Any time a new Kubernetes type is managed by the reconciler, the following switch should be extended
// to cover that type.
func kubernetesMeta(v interface{}) *meta.ObjectMeta {
	switch v2 := v.(type) {
	case *rbac.ClusterRole:
		return &v2.ObjectMeta
	case *rbac.ClusterRoleBinding:
		return &v2.ObjectMeta
	case *storage.CSIDriver:
		return &v2.ObjectMeta
	case *storage.StorageClass:
		return &v2.ObjectMeta
	case *policy.PodSecurityPolicy:
		return &v2.ObjectMeta
	}
	return nil
}

// TestExpectedNamedCorrectly ensures that all the Expected objects of all resource types have a correspondence between
// their returned key and inner name. This contract must be met in order for the reconciler to not create runaway
// resources. This assumes all managed resources are Kubernetes resources.
func TestExpectedNamedCorrectly(t *testing.T) {
	for reconciler, r := range allResources(nil) {
		for outer, v := range r.Expected() {
			meta := kubernetesMeta(v)
			if meta == nil {
				t.Errorf("reconciler %q, object %q: could not decode kubernetes metadata", reconciler, outer)
				continue
			}
			if inner := meta.Name; outer != inner {
				t.Errorf("reconciler %q, object %q: inner name mismatch (%q)", reconciler, outer, inner)
				continue
			}
		}
	}
}

// TestExpectedLabeledCorrectly ensures that all the Expected objects of all resource types have a Kubernetes metadata
// label that signifies it's a builtin object, to be retrieved afterwards. This contract must be met in order for the
// reconciler to not keep overwriting objects (and possibly failing), when a newly created object is not then
// retrievable using a selector corresponding to this label. This assumes all managed resources are Kubernetes objects.
func TestExpectedLabeledCorrectly(t *testing.T) {
	for reconciler, r := range allResources(nil) {
		for outer, v := range r.Expected() {
			meta := kubernetesMeta(v)
			if meta == nil {
				t.Errorf("reconciler %q, object %q: could not decode kubernetes metadata", reconciler, outer)
				continue
			}
			if data := meta.Labels[BuiltinLabelKey]; data != BuiltinLabelValue {
				t.Errorf("reconciler %q, object %q: %q=%q, wanted =%q", reconciler, outer, BuiltinLabelKey, data, BuiltinLabelValue)
				continue
			}
		}
	}
}

// testResource is a resource type used for testing. The inner type is a string that is equal to its name (key).
// It simulates a target (ie. k8s apiserver mock) that always acts nominally (all resources are created, deleted as
// requested, and the state is consistent with requests).
type testResource struct {
	// current is the simulated state of resources in the target.
	current map[string]string
	// expected is what this type will report as the Expected() resources.
	expected map[string]string
}

func (r *testResource) List(ctx context.Context) ([]string, error) {
	var keys []string
	for k, _ := range r.current {
		keys = append(keys, k)
	}
	return keys, nil
}

func (r *testResource) Create(ctx context.Context, el interface{}) error {
	r.current[el.(string)] = el.(string)
	return nil
}

func (r *testResource) Delete(ctx context.Context, name string) error {
	delete(r.current, name)
	return nil
}

func (r *testResource) Expected() map[string]interface{} {
	exp := make(map[string]interface{})
	for k, v := range r.expected {
		exp[k] = v
	}
	return exp
}

// newTestResource creates a test resource with a list of expected resource strings.
func newTestResource(want ...string) *testResource {
	expected := make(map[string]string)
	for _, w := range want {
		expected[w] = w
	}
	return &testResource{
		current:  make(map[string]string),
		expected: expected,
	}
}

// currentDiff returns a human-readable string showing the different between the current state and the given resource
// strings. If no difference is present, the returned string is empty.
func (r *testResource) currentDiff(want ...string) string {
	expected := make(map[string]string)
	for _, w := range want {
		if _, ok := r.current[w]; !ok {
			return fmt.Sprintf("%q missing in current", w)
		}
		expected[w] = w
	}
	for _, g := range r.current {
		if _, ok := expected[g]; !ok {
			return fmt.Sprintf("%q spurious in current", g)
		}
	}
	return ""
}

// TestBasicReconciliation ensures that the reconcile function does manipulate a target state based on a set of
// expected resources.
func TestBasicReconciliation(t *testing.T) {
	ctx := context.Background()
	r := newTestResource("foo", "bar", "baz")

	// nothing should have happened yet (testing the test)
	if diff := r.currentDiff(); diff != "" {
		t.Fatalf("wrong state after creation: %s", diff)
	}

	if err := reconcile(ctx, r); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	// everything requested should have been created
	if diff := r.currentDiff("foo", "bar", "baz"); diff != "" {
		t.Fatalf("wrong state after reconciliation: %s", diff)
	}

	delete(r.expected, "foo")
	if err := reconcile(ctx, r); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	// foo should not be missing
	if diff := r.currentDiff("bar", "baz"); diff != "" {
		t.Fatalf("wrong state after deleting foo: %s", diff)
	}
}
