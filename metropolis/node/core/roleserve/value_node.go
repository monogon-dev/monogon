package roleserve

import (
	"context"

	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type localRolesValue struct {
	value memory.Value
}

func (c *localRolesValue) Watch() *localRolesWatcher {
	return &localRolesWatcher{
		Watcher: c.value.Watch(),
	}
}

func (c *localRolesValue) set(v *cpb.NodeRoles) {
	c.value.Set(v)
}

type localRolesWatcher struct {
	event.Watcher
}

// Get retrieves the roles assigned to the local node by the cluster.
func (c *localRolesWatcher) Get(ctx context.Context) (*cpb.NodeRoles, error) {
	v, err := c.Watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	return v.(*cpb.NodeRoles), nil
}
