// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package object

// Taken and modified from the Kubernetes plugin of CoreDNS, under Apache 2.0.

import (
	"testing"

	api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func TestDefaultProcessor(t *testing.T) {
	pbuild := DefaultProcessor(ToService)
	reh := cache.ResourceEventHandlerFuncs{}
	idx := cache.NewIndexer(cache.DeletionHandlingMetaNamespaceKeyFunc, cache.Indexers{})
	processor := pbuild(idx, reh)
	testProcessor(t, processor, idx)
}

func testProcessor(t *testing.T, processor cache.ProcessFunc, idx cache.Indexer) {
	obj := &api.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "service1", Namespace: "test1"},
		Spec: api.ServiceSpec{
			Type:         api.ServiceTypeExternalName,
			ExternalName: "example.com.",
			Ports:        []api.ServicePort{{Port: 80}},
		},
	}
	obj2 := &api.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "service2", Namespace: "test1"},
		Spec: api.ServiceSpec{
			ClusterIP:  "5.6.7.8",
			ClusterIPs: []string{"5.6.7.8"},
			Ports:      []api.ServicePort{{Port: 80}},
		},
	}

	// Add the objects
	err := processor(cache.Deltas{
		{Type: cache.Added, Object: obj.DeepCopy()},
		{Type: cache.Added, Object: obj2.DeepCopy()},
	}, false)
	if err != nil {
		t.Fatalf("add failed: %v", err)
	}
	got, exists, err := idx.Get(obj)
	if err != nil {
		t.Fatalf("get added object failed: %v", err)
	}
	if !exists {
		t.Fatal("added object not found in index")
	}
	svc, ok := got.(*Service)
	if !ok {
		t.Fatal("object in index was incorrect type")
	}
	if svc.ExternalName != obj.Spec.ExternalName {
		t.Fatalf("expected '%v', got '%v'", obj.Spec.ExternalName, svc.ExternalName)
	}

	// Update an object
	obj.Spec.ExternalName = "2.example.com."
	err = processor(cache.Deltas{{
		Type:   cache.Updated,
		Object: obj.DeepCopy(),
	}}, false)
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}
	got, exists, err = idx.Get(obj)
	if err != nil {
		t.Fatalf("get updated object failed: %v", err)
	}
	if !exists {
		t.Fatal("updated object not found in index")
	}
	svc, ok = got.(*Service)
	if !ok {
		t.Fatal("object in index was incorrect type")
	}
	if svc.ExternalName != obj.Spec.ExternalName {
		t.Fatalf("expected '%v', got '%v'", obj.Spec.ExternalName, svc.ExternalName)
	}

	// Delete an object
	err = processor(cache.Deltas{{
		Type:   cache.Deleted,
		Object: obj2.DeepCopy(),
	}}, false)
	if err != nil {
		t.Fatalf("delete test failed: %v", err)
	}
	_, exists, err = idx.Get(obj2)
	if err != nil {
		t.Fatalf("get deleted object failed: %v", err)
	}
	if exists {
		t.Fatal("deleted object found in index")
	}

	// Delete an object via tombstone
	key, _ := cache.MetaNamespaceKeyFunc(obj)
	tombstone := cache.DeletedFinalStateUnknown{Key: key, Obj: svc}
	err = processor(cache.Deltas{{
		Type:   cache.Deleted,
		Object: tombstone,
	}}, false)
	if err != nil {
		t.Fatalf("tombstone delete test failed: %v", err)
	}
	_, exists, err = idx.Get(svc)
	if err != nil {
		t.Fatalf("get tombstone deleted object failed: %v", err)
	}
	if exists {
		t.Fatal("tombstone deleted object found in index")
	}
}
