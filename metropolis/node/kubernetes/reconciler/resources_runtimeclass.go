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

	node "k8s.io/api/node/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type resourceRuntimeClasses struct {
	kubernetes.Interface
}

func (r resourceRuntimeClasses) List(ctx context.Context) ([]meta.Object, error) {
	res, err := r.NodeV1beta1().RuntimeClasses().List(ctx, listBuiltins)
	if err != nil {
		return nil, err
	}
	objs := make([]meta.Object, len(res.Items))
	for i := range res.Items {
		objs[i] = &res.Items[i]
	}
	return objs, nil
}

func (r resourceRuntimeClasses) Create(ctx context.Context, el meta.Object) error {
	_, err := r.NodeV1beta1().RuntimeClasses().Create(ctx, el.(*node.RuntimeClass), meta.CreateOptions{})
	return err
}

func (r resourceRuntimeClasses) Delete(ctx context.Context, name string) error {
	return r.NodeV1beta1().RuntimeClasses().Delete(ctx, name, meta.DeleteOptions{})
}

func (r resourceRuntimeClasses) Expected() []meta.Object {
	return []meta.Object{
		&node.RuntimeClass{
			ObjectMeta: meta.ObjectMeta{
				Name:   "gvisor",
				Labels: builtinLabels(nil),
			},
			Handler: "runsc",
		},
		&node.RuntimeClass{
			ObjectMeta: meta.ObjectMeta{
				Name:   "runc",
				Labels: builtinLabels(nil),
			},
			Handler: "runc",
		},
	}
}
