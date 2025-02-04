// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package reconciler

import (
	"context"

	core "k8s.io/api/core/v1"
	storage "k8s.io/api/storage/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/ptr"
)

type resourceStorageClasses struct {
	kubernetes.Interface
}

func (r resourceStorageClasses) List(ctx context.Context) ([]meta.Object, error) {
	res, err := r.StorageV1().StorageClasses().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]meta.Object, len(res.Items))
	for i := range res.Items {
		objs[i] = &res.Items[i]
	}
	return objs, nil
}

func (r resourceStorageClasses) Create(ctx context.Context, el meta.Object) error {
	_, err := r.StorageV1().StorageClasses().Create(ctx, el.(*storage.StorageClass), meta.CreateOptions{})
	return err
}

func (r resourceStorageClasses) Update(ctx context.Context, el meta.Object) error {
	_, err := r.StorageV1().StorageClasses().Update(ctx, el.(*storage.StorageClass), meta.UpdateOptions{})
	return err
}

func (r resourceStorageClasses) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
	return r.StorageV1().StorageClasses().Delete(ctx, name, opts)
}

func (r resourceStorageClasses) Expected() []meta.Object {
	return []meta.Object{
		&storage.StorageClass{
			ObjectMeta: meta.ObjectMeta{
				Name:   "local",
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"storageclass.kubernetes.io/is-default-class": "true",
					"kubernetes.io/description": "local is the default storage class on Metropolis. " +
						"It stores data on the node root disk and supports space limits, resizing and oversubscription but no snapshots. " +
						"It is backed by XFS.",
				},
			},
			AllowVolumeExpansion: ptr.To(true),
			Provisioner:          csiProvisionerName,
			ReclaimPolicy:        ptr.To(core.PersistentVolumeReclaimDelete),
			VolumeBindingMode:    ptr.To(storage.VolumeBindingWaitForFirstConsumer),
			MountOptions: []string{
				"exec",
				"dev",
				"suid",
			},
		},
	}
}
