// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/grpc"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	cpb "source.monogon.dev/metropolis/proto/common"
	ppb "source.monogon.dev/metropolis/proto/private"
	"source.monogon.dev/osbase/supervisor"
)

// join implements Join Flow of an already registered node.
func (m *Manager) join(ctx context.Context, sc *ppb.SealedConfiguration, cd *cpb.ClusterDirectory, sealed bool) error {
	// Generate a complete ED25519 Join Key based on the seed included in Sealed
	// Configuration.
	var jpriv ed25519.PrivateKey = sc.JoinKey

	// Get Cluster CA from Sealed Configuration.
	ca, err := x509.ParseCertificate(sc.ClusterCa)
	if err != nil {
		return fmt.Errorf("cluster CA certificate present in Sealed Configuration could not be parsed: %w", err)
	}

	// Tell the user what we're doing.
	hpkey := hex.EncodeToString(jpriv.Public().(ed25519.PublicKey))
	supervisor.Logger(ctx).Infof("Joining an existing cluster.")
	supervisor.Logger(ctx).Infof("  Using TPM-secured configuration: %v", sealed)
	supervisor.Logger(ctx).Infof("  Node Join public key: %s", hpkey)

	// Build resolver used by the join process, authenticating with join
	// credentials. Once the join is complete, the rolesever will start its own
	// long-term resolver.
	rctx, rctxC := context.WithCancel(ctx)
	defer rctxC()
	r := resolver.New(rctx, resolver.WithoutCuratorUpdater(), resolver.WithLogger(supervisor.Logger(ctx)))
	addedNodes := 0
	for _, node := range cd.Nodes {
		if len(node.Addresses) == 0 {
			continue
		}
		// MVP: handle curators at non-default ports
		r.AddEndpoint(resolver.NodeAtAddressWithDefaultPort(node.Addresses[0].Host))
		addedNodes += 1
	}
	if addedNodes == 0 {
		return fmt.Errorf("no remote node available, cannot join cluster")
	}

	ephCreds, err := rpc.NewEphemeralCredentials(jpriv, rpc.WantRemoteCluster(ca))
	if err != nil {
		return fmt.Errorf("could not create ephemeral credentials: %w", err)
	}
	eph, err := grpc.Dial(resolver.MetropolisControlAddress, grpc.WithTransportCredentials(ephCreds), grpc.WithResolvers(r))
	if err != nil {
		return fmt.Errorf("could not dial cluster with join credentials: %w", err)
	}
	cur := ipb.NewCuratorClient(eph)

	// Retrieve CUK from cluster and reconstruct encryption key if we're not in
	// insecure mode.
	var cuk []byte
	if sc.StorageSecurity != cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_INSECURE {
		if want, got := 32, len(sc.NodeUnlockKey); want != got {
			return fmt.Errorf("sealed configuration has invalid node unlock key (wanted %d bytes, got %d)", want, got)
		}

		// Join the cluster and use the newly obtained CUK to mount the data
		// partition.
		var jr *ipb.JoinNodeResponse
		bo := backoff.NewExponentialBackOff()
		bo.MaxElapsedTime = 0
		backoff.Retry(func() error {
			jr, err = cur.JoinNode(ctx, &ipb.JoinNodeRequest{
				UsingSealedConfiguration: sealed,
			})
			if err != nil {
				supervisor.Logger(ctx).Warningf("Join failed: %v", err)
				// This is never used.
				return fmt.Errorf("join call failed")
			}
			return nil
		}, bo)
		cuk = jr.ClusterUnlockKey

		if want, got := 32, len(cuk); want != got {
			return fmt.Errorf("cluster returned invalid cluster unlock key (wanted %d bytes, got %d)", want, got)
		}
	}

	if err := m.storageRoot.Data.MountExisting(sc, cuk); err != nil {
		return fmt.Errorf("while mounting Data: %w", err)
	}

	// Use the node credentials found in the data partition.
	var creds identity.NodeCredentials
	if err := creds.Read(&m.storageRoot.Data.Node.Credentials); err != nil {
		return fmt.Errorf("while reading node credentials: %w", err)
	}
	m.roleServer.ProvideJoinData(creds, cd)

	// After successfully joining cluster, mark boot as successful.
	// This allows the update service to mark the currently-booted slot as good
	// if an update has been performed.
	if err := m.updateService.MarkBootSuccessful(); err != nil {
		supervisor.Logger(ctx).Errorf("Failed to mark boot as successful: %v", err)
	}

	supervisor.Logger(ctx).Infof("Joined the cluster.")
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
