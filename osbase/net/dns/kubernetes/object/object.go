// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package object holds functions that convert the objects from the k8s API in
// to a more memory efficient structures.
//
// Adding new fields to any of the structures defined in pod.go, endpoint.go
// and service.go should not be done lightly as this increases the memory use
// and will leads to OOMs in the k8s scale test.
package object

// Taken and modified from the Kubernetes plugin of CoreDNS, under Apache 2.0.

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
)

// ToFunc converts one v1.Object to another v1.Object.
type ToFunc func(v1.Object) (v1.Object, error)

// ProcessorBuilder returns function to process cache events.
type ProcessorBuilder func(cache.Indexer, cache.ResourceEventHandler) cache.ProcessFunc

// Empty is an empty struct.
type Empty struct{}

// GetObjectKind implements the ObjectKind interface as a noop.
func (e *Empty) GetObjectKind() schema.ObjectKind { return schema.EmptyObjectKind }

// GetGenerateName implements the metav1.Object interface.
func (e *Empty) GetGenerateName() string { return "" }

// SetGenerateName implements the metav1.Object interface.
func (e *Empty) SetGenerateName(name string) {}

// GetUID implements the metav1.Object interface.
func (e *Empty) GetUID() types.UID { return "" }

// SetUID implements the metav1.Object interface.
func (e *Empty) SetUID(uid types.UID) {}

// GetGeneration implements the metav1.Object interface.
func (e *Empty) GetGeneration() int64 { return 0 }

// SetGeneration implements the metav1.Object interface.
func (e *Empty) SetGeneration(generation int64) {}

// GetSelfLink implements the metav1.Object interface.
func (e *Empty) GetSelfLink() string { return "" }

// SetSelfLink implements the metav1.Object interface.
func (e *Empty) SetSelfLink(selfLink string) {}

// GetCreationTimestamp implements the metav1.Object interface.
func (e *Empty) GetCreationTimestamp() v1.Time { return v1.Time{} }

// SetCreationTimestamp implements the metav1.Object interface.
func (e *Empty) SetCreationTimestamp(timestamp v1.Time) {}

// GetDeletionTimestamp implements the metav1.Object interface.
func (e *Empty) GetDeletionTimestamp() *v1.Time { return &v1.Time{} }

// SetDeletionTimestamp implements the metav1.Object interface.
func (e *Empty) SetDeletionTimestamp(timestamp *v1.Time) {}

// GetDeletionGracePeriodSeconds implements the metav1.Object interface.
func (e *Empty) GetDeletionGracePeriodSeconds() *int64 { return nil }

// SetDeletionGracePeriodSeconds implements the metav1.Object interface.
func (e *Empty) SetDeletionGracePeriodSeconds(*int64) {}

// GetLabels implements the metav1.Object interface.
func (e *Empty) GetLabels() map[string]string { return nil }

// SetLabels implements the metav1.Object interface.
func (e *Empty) SetLabels(labels map[string]string) {}

// GetAnnotations implements the metav1.Object interface.
func (e *Empty) GetAnnotations() map[string]string { return nil }

// SetAnnotations implements the metav1.Object interface.
func (e *Empty) SetAnnotations(annotations map[string]string) {}

// GetFinalizers implements the metav1.Object interface.
func (e *Empty) GetFinalizers() []string { return nil }

// SetFinalizers implements the metav1.Object interface.
func (e *Empty) SetFinalizers(finalizers []string) {}

// GetOwnerReferences implements the metav1.Object interface.
func (e *Empty) GetOwnerReferences() []v1.OwnerReference { return nil }

// SetOwnerReferences implements the metav1.Object interface.
func (e *Empty) SetOwnerReferences([]v1.OwnerReference) {}

// GetZZZ_DeprecatedClusterName implements the metav1.Object interface.
func (e *Empty) GetZZZ_DeprecatedClusterName() string { return "" }

// SetZZZ_DeprecatedClusterName implements the metav1.Object interface.
func (e *Empty) SetZZZ_DeprecatedClusterName(clusterName string) {}

// GetManagedFields implements the metav1.Object interface.
func (e *Empty) GetManagedFields() []v1.ManagedFieldsEntry { return nil }

// SetManagedFields implements the metav1.Object interface.
func (e *Empty) SetManagedFields(managedFields []v1.ManagedFieldsEntry) {}
