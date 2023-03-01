package roleserve

import (
	"crypto/ed25519"
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
