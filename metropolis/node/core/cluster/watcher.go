package cluster

import (
	"context"
	"fmt"

	"source.monogon.dev/metropolis/pkg/event"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type Watcher struct {
	event.Watcher
}

func (w *Watcher) Get(ctx context.Context) (*Status, error) {
	val, err := w.Watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	status := val.(Status)
	return &status, err
}

// GetHome waits until the cluster, from the point of view of this node, is in
// the ClusterHome state. This can be used to wait for the cluster manager to
// 'settle', before clients start more node services.
func (w *Watcher) GetHome(ctx context.Context) (*Status, error) {
	for {
		status, err := w.Get(ctx)
		if err != nil {
			return nil, err
		}
		switch status.State {
		case cpb.ClusterState_CLUSTER_STATE_HOME:
			return status, nil
		case cpb.ClusterState_CLUSTER_STATE_DISOWNING:
			return nil, fmt.Errorf("the cluster has disowned this node")
		}
	}
}

func (m *Manager) Watch() Watcher {
	return Watcher{
		Watcher: m.status.Watch(),
	}
}
