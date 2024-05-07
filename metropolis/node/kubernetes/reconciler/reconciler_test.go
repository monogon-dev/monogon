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

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	installnode "k8s.io/kubernetes/pkg/apis/node/install"
	installpolicy "k8s.io/kubernetes/pkg/apis/policy/install"
	installrbac "k8s.io/kubernetes/pkg/apis/rbac/install"
	installstorage "k8s.io/kubernetes/pkg/apis/storage/install"

	"source.monogon.dev/osbase/supervisor"
)

// TestExpectedUniqueNames ensures that all the Expected objects of any
// given resource type have a unique name.
func TestExpectedUniqueNames(t *testing.T) {
	for reconciler, r := range allResources(nil) {
		names := make(map[string]bool)
		for _, v := range r.Expected() {
			if names[v.GetName()] {
				t.Errorf("reconciler %q: duplicate name %q", reconciler, v.GetName())
				continue
			}
			names[v.GetName()] = true
		}
	}
}

// TestExpectedLabeledCorrectly ensures that all the Expected objects of all
// resource types have a Kubernetes metadata label that signifies it's a
// builtin object, to be retrieved afterwards. This contract must be met in
// order for the reconciler to not keep overwriting objects (and possibly
// failing), when a newly created object is not then retrievable using a
// selector corresponding to this label.
func TestExpectedLabeledCorrectly(t *testing.T) {
	for reconciler, r := range allResources(nil) {
		for _, v := range r.Expected() {
			if data := v.GetLabels()[BuiltinLabelKey]; data != BuiltinLabelValue {
				t.Errorf("reconciler %q, object %q: %q=%q, wanted =%q", reconciler, v.GetName(), BuiltinLabelKey, data, BuiltinLabelValue)
				continue
			}
		}
	}
}

// TestExpectedDefaulted ensures that all the Expected objects of all
// resource types have defaults already applied. If this were not the case,
// the reconciler would think that the object has changed and try to update it
// in each iteration. If this test fails, the most likely fix is to add the
// missing default values to the expected objects.
func TestExpectedDefaulted(t *testing.T) {
	scheme := runtime.NewScheme()
	installnode.Install(scheme)
	installpolicy.Install(scheme)
	installrbac.Install(scheme)
	installstorage.Install(scheme)

	for reconciler, r := range allResources(nil) {
		for _, v := range r.Expected() {
			v_defaulted := v.(runtime.Object).DeepCopyObject()
			if _, ok := scheme.IsUnversioned(v_defaulted); !ok {
				t.Errorf("reconciler %q: type not installed in scheme", reconciler)
			}
			scheme.Default(v_defaulted)
			if !apiequality.Semantic.DeepEqual(v, v_defaulted) {
				t.Errorf("reconciler %q, object %q changed after defaulting\ngot: %+v\nwanted: %+v", reconciler, v.GetName(), v, v_defaulted)
			}
		}
	}
}

// testObject is the object type managed by testResource.
type testObject struct {
	meta.ObjectMeta
	Val int
}

func makeTestObject(name string, val int) *testObject {
	return &testObject{
		ObjectMeta: meta.ObjectMeta{
			Name:   name,
			Labels: builtinLabels(nil),
		},
		Val: val,
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

func (r *testResource) Update(ctx context.Context, el meta.Object) error {
	r.current[el.GetName()] = el.(*testObject)
	return nil
}

func (r *testResource) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
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
	// This needs to run in a TestHarness to make logging work.
	supervisor.TestHarness(t, func(ctx context.Context) error {
		r := newTestResource(makeTestObject("foo", 0), makeTestObject("bar", 0), makeTestObject("baz", 0))
		rname := "testresource"

		// nothing should have happened yet (testing the test)
		if diff := r.currentDiff(); diff != "" {
			return fmt.Errorf("wrong state after creation: %s", diff)
		}

		if err := reconcile(ctx, r, rname); err != nil {
			return fmt.Errorf("reconcile: %v", err)
		}
		// everything requested should have been created
		if diff := r.currentDiff(makeTestObject("foo", 0), makeTestObject("bar", 0), makeTestObject("baz", 0)); diff != "" {
			return fmt.Errorf("wrong state after reconciliation: %s", diff)
		}

		delete(r.expected, "foo")
		if err := reconcile(ctx, r, rname); err != nil {
			return fmt.Errorf("reconcile: %v", err)
		}
		// foo should now be missing
		if diff := r.currentDiff(makeTestObject("bar", 0), makeTestObject("baz", 0)); diff != "" {
			return fmt.Errorf("wrong state after deleting foo: %s", diff)
		}

		r.expected["bar"] = makeTestObject("bar", 1)
		if err := reconcile(ctx, r, rname); err != nil {
			return fmt.Errorf("reconcile: %v", err)
		}
		// bar should be updated
		if diff := r.currentDiff(makeTestObject("bar", 1), makeTestObject("baz", 0)); diff != "" {
			return fmt.Errorf("wrong state after deleting foo: %s", diff)
		}

		return nil
	})
}

func TestIsImmutableError(t *testing.T) {
	gk := schema.GroupKind{Group: "someGroup", Kind: "someKind"}
	cases := []struct {
		err         error
		isImmutable bool
	}{
		{fmt.Errorf("something wrong"), false},
		{apierrors.NewApplyConflict(nil, "conflict"), false},
		{apierrors.NewInvalid(gk, "name", field.ErrorList{}), false},
		{apierrors.NewInvalid(gk, "name", field.ErrorList{
			field.Invalid(field.NewPath("field1"), true, apivalidation.FieldImmutableErrorMsg),
			field.Invalid(field.NewPath("field2"), true, "some other error"),
		}), false},
		{apierrors.NewInvalid(gk, "name", field.ErrorList{
			field.Invalid(field.NewPath("field1"), true, apivalidation.FieldImmutableErrorMsg),
		}), true},
		{apierrors.NewInvalid(gk, "name", field.ErrorList{
			field.Invalid(field.NewPath("field1"), true, apivalidation.FieldImmutableErrorMsg),
			field.Invalid(field.NewPath("field2"), true, apivalidation.FieldImmutableErrorMsg),
		}), true},
	}
	for _, c := range cases {
		actual := isImmutableError(c.err)
		if actual != c.isImmutable {
			t.Errorf("Expected %v, got %v for error: %v", c.isImmutable, actual, c.err)
		}
	}
}
