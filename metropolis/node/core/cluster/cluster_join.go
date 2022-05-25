package cluster

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"google.golang.org/grpc"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
	ppb "source.monogon.dev/metropolis/proto/private"
)

// join implements Join Flow of an already registered node.
func (m *Manager) join(ctx context.Context, sc *ppb.SealedConfiguration, cd *cpb.ClusterDirectory) error {
	// Generate a complete ED25519 Join Key based on the seed included in Sealed
	// Configuration.
	var jpriv ed25519.PrivateKey = sc.JoinKey

	// Get Cluster CA from Sealed Configuration.
	ca, err := x509.ParseCertificate(sc.ClusterCa)
	if err != nil {
		return fmt.Errorf("Cluster CA certificate present in Sealed Configuration could not be parsed: %w", err)
	}

	// Tell the user what we're doing.
	hpkey := hex.EncodeToString(jpriv.Public().(ed25519.PublicKey))
	supervisor.Logger(ctx).Infof("Joining an existing cluster.")
	supervisor.Logger(ctx).Infof("  Node Join public key: %s", hpkey)
	supervisor.Logger(ctx).Infof("  Directory:")
	logClusterDirectory(ctx, cd)

	// Attempt to connect to the first node in the cluster directory.
	r, err := curatorRemote(cd)
	if err != nil {
		return fmt.Errorf("while picking a Curator endpoint: %w", err)
	}
	ephCreds, err := rpc.NewEphemeralCredentials(jpriv, ca)
	if err != nil {
		return fmt.Errorf("could not create ephemeral credentials: %w", err)
	}
	eph, err := grpc.Dial(r, grpc.WithTransportCredentials(ephCreds))
	if err != nil {
		return fmt.Errorf("could not create ephemeral client to %q: %w", r, err)
	}
	cur := ipb.NewCuratorClient(eph)

	// Join the cluster and use the newly obtained CUK to mount the data
	// partition.
	jr, err := cur.JoinNode(ctx, &ipb.JoinNodeRequest{})
	if err != nil {
		return fmt.Errorf("join call failed: %w", err)
	}
	if err := m.storageRoot.Data.MountExisting(sc, jr.ClusterUnlockKey); err != nil {
		return fmt.Errorf("while mounting Data: %w", err)
	}

	// Use the node credentials found in the data partition.
	var creds identity.NodeCredentials
	if err := creds.Read(&m.storageRoot.Data.Node.Credentials); err != nil {
		return fmt.Errorf("while reading node credentials: %w", err)
	}
	m.roleServer.ProvideJoinData(creds, cd)

	supervisor.Logger(ctx).Infof("Joined the cluster.")
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}