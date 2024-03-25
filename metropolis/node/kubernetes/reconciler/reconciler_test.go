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

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestExpectedLabeledCorrectly ensures that all the Expected objects of all
// resource types have a Kubernetes metadata label that signifies it's a
// builtin object, to be retrieved afterwards. This contract must be met in
// order for the reconciler to not keep overwriting objects (and possibly
// failing), when a newly created object is not then retrievable using a
// selector corresponding to this label.
func TestExpectedLabeledCorrectly(t *testing.T) {
	for reconciler, r := range allResources(nil) {
		for outer, v := range r.Expected() {
			if data := v.GetLabels()[BuiltinLabelKey]; data != BuiltinLabelValue {
				t.Errorf("reconciler %q, object %q: %q=%q, wanted =%q", reconciler, outer, BuiltinLabelKey, data, BuiltinLabelValue)
				continue
			}
		}
	}
}

// testObject is the object type managed by testResource.
type testObject struct {
	meta.ObjectMeta
}

func makeTestObject(name string) *testObject {
	return &testObject{
		ObjectMeta: meta.ObjectMeta{
			Name:   name,
			Labels: builtinLabels(nil),
		},
	}
}

// testResource is a resource type used for testing. It simulates a target
// (ie. k8s apiserver mock) that always acts nominally (all resources are
// created, deleted as requested, and the state is consistent with requests).
type testResource struct {
	// current is the simulated state of resources in the target.
	current map[string]*testObject
	// expected is what this type will report as the Expected() resources.
	expected map[string]*testObject
}

func (r *testResource) List(ctx context.Context) ([]meta.Object, error) {
	var cur []meta.Object
	for _, v := range r.current {
		v_copy := *v
		cur = append(cur, &v_copy)
	}
	return cur, nil
}

func (r *testResource) Create(ctx context.Context, el meta.Object) error {
	r.current[el.GetName()] = el.(*testObject)
	return nil
}

func (r *testResource) Delete(ctx context.Context, name string) error {
	delete(r.current, name)
	return nil
}

func (r *testResource) Expected() []meta.Object {
	var exp []meta.Object
	for _, v := range r.expected {
		v_copy := *v
		exp = append(exp, &v_copy)
	}
	return exp
}

// newTestResource creates a test resource with a list of expected objects.
func newTestResource(want ...*testObject) *testResource {
	expected := make(map[string]*testObject)
	for _, w := range want {
		expected[w.GetName()] = w
	}
	return &testResource{
		current:  make(map[string]*testObject),
		expected: expected,
	}
}

// currentDiff returns a human-readable string showing the difference between
// the current state and the given objects. If no difference is
// present, the returned string is empty.
func (r *testResource) currentDiff(want ...*testObject) string {
	expected := make(map[string]*testObject)
	for _, w := range want {
		if _, ok := r.current[w.GetName()]; !ok {
			return fmt.Sprintf("%q missing in current", w.GetName())
		}
		expected[w.GetName()] = w
	}
	for _, g := range r.current {
		if _, ok := expected[g.GetName()]; !ok {
			return fmt.Sprintf("%q spurious in current", g.GetName())
		}
	}
	return ""
}

// TestBasicReconciliation ensures that the reconcile function does manipulate
// a target state based on a set of expected resources.
func TestBasicReconciliation(t *testing.T) {
	ctx := context.Background()
	r := newTestResource(makeTestObject("foo"), makeTestObject("bar"), makeTestObject("baz"))

	// nothing should have happened yet (testing the test)
	if diff := r.currentDiff(); diff != "" {
		t.Fatalf("wrong state after creation: %s", diff)
	}

	if err := reconcile(ctx, r); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	// everything requested should have been created
	if diff := r.currentDiff(makeTestObject("foo"), makeTestObject("bar"), makeTestObject("baz")); diff != "" {
		t.Fatalf("wrong state after reconciliation: %s", diff)
	}

	delete(r.expected, "foo")
	if err := reconcile(ctx, r); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	// foo should now be missing
	if diff := r.currentDiff(makeTestObject("bar"), makeTestObject("baz")); diff != "" {
		t.Fatalf("wrong state after deleting foo: %s", diff)
	}
}
