// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package watcher

import (
	"context"
	"fmt"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

// WatchNode runs a WatchRequest for NodeInCluster with the given Curator
// channel. The returned Watcher can then be queries in a for loop to detect
// changes to the targeted node.
func WatchNode(ctx context.Context, cur ipb.CuratorClient, nid string) *Watcher {
	wa, err := cur.Watch(ctx, &ipb.WatchRequest{
		Kind: &ipb.WatchRequest_NodeInCluster_{
			NodeInCluster: &ipb.WatchRequest_NodeInCluster{
				NodeId: nid,
			},
		},
	})
	if err != nil {
		return &Watcher{
			err: fmt.Errorf("could not watch node: %w", err),
		}
	}

	return &Watcher{
		wa: wa,
	}
}

// Watcher returned by WatchNode. Must be closed by calling Close().
type Watcher struct {
	err error
	wa  ipb.Curator_WatchClient
	ev  *ipb.WatchEvent
	ix  int
}

// Close RPC call associted with this Watcher. Must be called at least once after
// the Watcher is not used anymore.
func (w *Watcher) Close() {
	if w.wa != nil {
		w.wa.CloseSend()
		w.wa = nil
	}
}

// Next returns true if the next call to Node() is valid, false otherwise. Each
// call to Next blocks until an update to the node data is available.
//
// If false is returned, Error() should be called to get to the underlying error
// which caused this call to fail.
func (w *Watcher) Next() bool {
	if w.err != nil {
		w.Close()
		return false
	}
	if w.wa == nil {
		w.err = fmt.Errorf("watcher closed")
		return false
	}

	w.ix += 1
	if w.ev == nil || w.ix >= len(w.ev.Nodes) {
		ev, err := w.wa.Recv()
		if err != nil {
			w.err = err
			return false
		}
		w.ev = ev
		w.ix = 0
	}
	return true
}

// Error returns underlying error for this Watcher, nil if no error is present.
// After an error is returned, the Watcher cannot be used anymore.
func (w *Watcher) Error() error {
	return w.err
}

// Node returns the cached node state for this Watcher. The same node data is
// returned until Next() is called. The caller can hold on to the returned Node
// pointer, as the node data will not be modified in place with updates -
// instead, a new Node object will be returned.
func (w *Watcher) Node() *ipb.Node {
	return w.ev.Nodes[w.ix]
}
