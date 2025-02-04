// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package reconciler

import (
	"context"

	node "k8s.io/api/node/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type resourceRuntimeClasses struct {
	kubernetes.Interface
}

func (r resourceRuntimeClasses) List(ctx context.Context) ([]meta.Object, error) {
	res, err := r.NodeV1().RuntimeClasses().List(ctx, listBuiltins)
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
	_, err := r.NodeV1().RuntimeClasses().Create(ctx, el.(*node.RuntimeClass), meta.CreateOptions{})
	return err
}

func (r resourceRuntimeClasses) Update(ctx context.Context, el meta.Object) error {
	_, err := r.NodeV1().RuntimeClasses().Update(ctx, el.(*node.RuntimeClass), meta.UpdateOptions{})
	return err
}

func (r resourceRuntimeClasses) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
	return r.NodeV1().RuntimeClasses().Delete(ctx, name, opts)
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
