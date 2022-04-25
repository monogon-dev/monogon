package roleserve

import (
	"context"
	"crypto/ed25519"

	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
)

// bootstrapData is an internal EventValue structure which is populated by the
// Cluster Enrolment logic via ProvideBootstrapData. It contains data needed by
// the control plane logic to go into bootstrap mode and bring up a control
// plane from scratch.
type bootstrapData struct {
	nodePrivateKey     ed25519.PrivateKey
	clusterUnlockKey   []byte
	nodeUnlockKey      []byte
	initialOwnerKey    []byte
	nodePrivateJoinKey ed25519.PrivateKey
}

type bootstrapDataValue struct {
	value memory.Value
}

func (c *bootstrapDataValue) Watch() *bootstrapDataWatcher {
	return &bootstrapDataWatcher{
		Watcher: c.value.Watch(),
	}
}

func (c *bootstrapDataValue) set(v *bootstrapData) {
	c.value.Set(v)
}

type bootstrapDataWatcher struct {
	event.Watcher
}

func (c *bootstrapDataWatcher) Get(ctx context.Context) (*bootstrapData, error) {
	v, err := c.Watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	return v.(*bootstrapData), nil
}
