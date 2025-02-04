// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package reconciler

import (
	"context"

	storage "k8s.io/api/storage/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/ptr"
)

// TODO(q3k): this is duplicated with
// //metropolis/node/kubernetes:provisioner.go; integrate this once
// provisioner.go gets moved into a subpackage.
// ONCHANGE(//metropolis/node/kubernetes:provisioner.go): needs to match
// csiProvisionerName declared.
const csiProvisionerName = "dev.monogon.metropolis.vfs"

type resourceCSIDrivers struct {
	kubernetes.Interface
}

func (r resourceCSIDrivers) List(ctx context.Context) ([]meta.Object, error) {
	res, err := r.StorageV1().CSIDrivers().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]meta.Object, len(res.Items))
	for i := range res.Items {
		objs[i] = &res.Items[i]
	}
	return objs, nil
}

func (r resourceCSIDrivers) Create(ctx context.Context, el meta.Object) error {
	_, err := r.StorageV1().CSIDrivers().Create(ctx, el.(*storage.CSIDriver), meta.CreateOptions{})
	return err
}

func (r resourceCSIDrivers) Update(ctx context.Context, el meta.Object) error {
	_, err := r.StorageV1().CSIDrivers().Update(ctx, el.(*storage.CSIDriver), meta.UpdateOptions{})
	return err
}

func (r resourceCSIDrivers) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
	return r.StorageV1().CSIDrivers().Delete(ctx, name, opts)
}

func (r resourceCSIDrivers) Expected() []meta.Object {
	return []meta.Object{
		&storage.CSIDriver{
			ObjectMeta: meta.ObjectMeta{
				Name:   csiProvisionerName,
				Labels: builtinLabels(nil),
			},
			Spec: storage.CSIDriverSpec{
				AttachRequired:       ptr.To(false),
				PodInfoOnMount:       ptr.To(false),
				VolumeLifecycleModes: []storage.VolumeLifecycleMode{storage.VolumeLifecyclePersistent},
				StorageCapacity:      ptr.To(false),
				FSGroupPolicy:        ptr.To(storage.FileFSGroupPolicy),
				RequiresRepublish:    ptr.To(false),
				SELinuxMount:         ptr.To(false),
			},
		},
	}
}
