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

	core "k8s.io/api/core/v1"
	storage "k8s.io/api/storage/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var reclaimPolicyDelete = core.PersistentVolumeReclaimDelete
var waitForConsumerBinding = storage.VolumeBindingWaitForFirstConsumer

type resourceStorageClasses struct {
	kubernetes.Interface
}

func (r resourceStorageClasses) List(ctx context.Context) ([]string, error) {
	res, err := r.StorageV1().StorageClasses().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]string, len(res.Items))
	for i, el := range res.Items {
		objs[i] = el.ObjectMeta.Name
	}
	return objs, nil
}

func (r resourceStorageClasses) Create(ctx context.Context, el interface{}) error {
	_, err := r.StorageV1().StorageClasses().Create(ctx, el.(*storage.StorageClass), meta.CreateOptions{})
	return err
}

func (r resourceStorageClasses) Delete(ctx context.Context, name string) error {
	return r.StorageV1().StorageClasses().Delete(ctx, name, meta.DeleteOptions{})
}

func (r resourceStorageClasses) Expected() map[string]interface{} {
	return map[string]interface{}{
		"local": &storage.StorageClass{
			ObjectMeta: meta.ObjectMeta{
				Name:   "local",
				Labels: builtinLabels(nil),
				Annotations: map[string]string{
					"storageclass.kubernetes.io/is-default-class": "true",
				},
			},
			AllowVolumeExpansion: True(),
			Provisioner:          csiProvisionerName,
			ReclaimPolicy:        &reclaimPolicyDelete,
			VolumeBindingMode:    &waitForConsumerBinding,
		},
	}
}
