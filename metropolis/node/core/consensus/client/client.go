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

// package client implements a higher-level client for consensus/etcd that is
// to be used within the Metropolis node code for unprivileged access (ie.
// access by local services that simply wish to access etcd KV without
// management access).
package client

import (
	"context"
	"fmt"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/namespace"
)

// Namespaced etcd/consensus client. Each Namespaced client allows access to a
// subtree of the etcd key/value space, and each can emit more clients that
// reside in their respective subtree - effectively permitting delegated,
// hierarchical access to the etcd store.
// Note: the namespaces should not be treated as a security boundary, as it's
// very likely possible that compromised services could navigate upwards in the
// k/v space if needed. Instead, this mechanism should only be seen as
// containerization for the purpose of simplifying code that needs to access
// etcd, and especially code that needs to pass this access around to its
// subordinate code.
// This client embeds the KV, Lease and Watcher etcd client interfaces to
// perform the actual etcd operations, and the Sub method to create subtree
// clients of this client.
type Namespaced interface {
	clientv3.KV
	clientv3.Lease
	clientv3.Watcher

	// Sub returns a child client from this client, at a sub-namespace 'space'.
	// The given 'space' path in a series of created clients (eg.
	// Namespace.Sub("a").Sub("b").Sub("c") are used to create an etcd k/v
	// prefix `a:b:c/` into which K/V access is remapped.
	Sub(space string) (Namespaced, error)

	// ThinClient returns a clientv3.Client which has the same namespacing as the
	// namespaced interface. It only implements the KV, Lease and Watcher interfaces
	// - all other interfaces are unimplemented and will panic when called. The
	// given context is returned by client.Ctx() and is used by some library code
	// (eg. etcd client-go's built-in concurrency library).
	ThinClient(ctx context.Context) *clientv3.Client
}

// ThinClient takes a set of KV, Lease and Watcher etcd clients and turns them
// into a full Client struct. The rest of the interfaces (Cluster, Auth,
// Maintenance) will all panic when called.
func ThinClient(ctx context.Context, kv clientv3.KV, lease clientv3.Lease, watcher clientv3.Watcher) *clientv3.Client {
	cli := clientv3.NewCtxClient(ctx)
	cli.Cluster = &unimplementedCluster{}
	cli.KV = kv
	cli.Lease = lease
	cli.Watcher = watcher
	cli.Auth = &unimplementedAuth{}
	cli.Maintenance = &unimplementedMaintenance{}
	return cli
}

// local implements the Namespaced client to access a locally running etc.
type local struct {
	root *clientv3.Client
	path []string

	clientv3.KV
	clientv3.Lease
	clientv3.Watcher
}

// NewLocal returns a local Namespaced client starting at the root of the given
// etcd client.
func NewLocal(cl *clientv3.Client) Namespaced {
	l := &local{
		root: cl,
		path: nil,
	}
	l.populate()
	return l
}

// populate prepares the namespaced KV/Watcher/Lease clients given the current
// root and path of the local client.
func (l *local) populate() {
	space := strings.Join(l.path, ":") + "/"
	l.KV = namespace.NewKV(l.root, space)
	l.Watcher = namespace.NewWatcher(l.root, space)
	l.Lease = namespace.NewLease(l.root, space)
}

func (l *local) Sub(space string) (Namespaced, error) {
	if strings.Contains(space, ":") {
		return nil, fmt.Errorf("sub-namespace name cannot contain ':' characters")
	}
	sub := &local{
		root: l.root,
		path: append(l.path, space),
	}
	sub.populate()
	return sub, nil
}

func (l *local) ThinClient(ctx context.Context) *clientv3.Client {
	return ThinClient(ctx, l.KV, l.Lease, l.Watcher)
}

func (l *local) Close() error {
	errW := l.Watcher.Close()
	errL := l.Lease.Close()
	if errW == nil && errL == nil {
		return nil
	}
	if errW != nil && errL == nil {
		return fmt.Errorf("closing watcher: %w", errW)
	}
	if errL != nil && errW == nil {
		return fmt.Errorf("closing lease: %w", errL)
	}
	return fmt.Errorf("closing watcher: %v, closing lease: %v", errW, errL)
}
