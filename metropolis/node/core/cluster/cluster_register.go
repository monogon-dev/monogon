package cluster

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"source.monogon.dev/metropolis/node"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	ppb "source.monogon.dev/metropolis/proto/private"
)

// register performs the registration flow of the node in the cluster, ie. makes
// this node part of an existing cluster.
//
// This is a temporary implementation that's not well hardened against
// transitive failures, but is good enough to get us off the ground and able to
// test multiple nodes in a cluster. It does notably not run either
// etcd/consensus or the curator, which also prevents Kubernetes from running.
func (m *Manager) register(ctx context.Context, register *apb.NodeParameters_ClusterRegister) error {
	// Do a validation pass on the provided NodeParameters.Register data, fail early
	// if it looks invalid.
	ca, err := x509.ParseCertificate(register.CaCertificate)
	if err != nil {
		return fmt.Errorf("NodeParameters.Register invalid: CaCertificate could not parsed: %w", err)
	}
	if err := identity.VerifyCAInsecure(ca); err != nil {
		return fmt.Errorf("NodeParameters.Register invalid: CaCertificate invalid: %w", err)
	}
	if len(register.RegisterTicket) == 0 {
		return fmt.Errorf("NodeParameters.Register invalid: RegisterTicket not set")
	}
	if register.ClusterDirectory == nil || len(register.ClusterDirectory.Nodes) == 0 {
		return fmt.Errorf("NodeParameters.ClusterDirectory invalid: must contain at least one node")
	}
	for i, node := range register.ClusterDirectory.Nodes {
		if len(node.Addresses) == 0 {
			return fmt.Errorf("NodeParameters.ClusterDirectory.Nodes[%d] invalid: must have at least one address", i)
		}
		for j, add := range node.Addresses {
			if add.Host == "" || net.ParseIP(add.Host) == nil {
				return fmt.Errorf("NodeParameters.ClusterDirectory.Nodes[%d].Addresses[%d] invalid: not a parseable address", i, j)
			}
		}
		if len(node.PublicKey) != ed25519.PublicKeySize {
			return fmt.Errorf("NodeParameters.ClusterDirectory.Nodes[%d] invalid: PublicKey invalid length or not set", i)
		}
	}

	// Validation passed, let's take the state lock and start working on registering
	// us into the cluster.

	state, unlock := m.lock()
	defer unlock()

	// Tell the user what we're doing.
	supervisor.Logger(ctx).Infof("Registering into existing cluster.")
	supervisor.Logger(ctx).Infof("  Cluster CA public key: %s", hex.EncodeToString(ca.PublicKey.(ed25519.PublicKey)))
	supervisor.Logger(ctx).Infof("  Register Ticket: %s", hex.EncodeToString(register.RegisterTicket))
	supervisor.Logger(ctx).Infof("  Directory:")
	for _, node := range register.ClusterDirectory.Nodes {
		id := identity.NodeID(node.PublicKey)
		var addresses []string
		for _, add := range node.Addresses {
			addresses = append(addresses, add.Host)
		}
		supervisor.Logger(ctx).Infof("    Node ID: %s, Addresses: %s", id, strings.Join(addresses, ","))
	}

	// Mount new storage with generated CUK, and save LUK into sealed config proto.
	state.configuration = &ppb.SealedConfiguration{}
	supervisor.Logger(ctx).Infof("Registering: mounting new storage...")
	cuk, err := m.storageRoot.Data.MountNew(state.configuration)
	if err != nil {
		return fmt.Errorf("could not make and mount data partition: %w", err)
	}

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate node keypair: %w", err)
	}
	supervisor.Logger(ctx).Infof("Registering: node public key: %s", hex.EncodeToString([]byte(pub)))

	// Attempt to connect to first node in cluster directory and to call Register.
	//
	// MVP: this should be properly client-side loadbalanced.
	remote := register.ClusterDirectory.Nodes[0].Addresses[0].Host
	remote = net.JoinHostPort(remote, strconv.Itoa(int(node.CuratorServicePort)))
	eph, err := rpc.NewEphemeralClient(remote, priv, ca)
	if err != nil {
		return fmt.Errorf("could not create ephemeral client to %q: %w", remote, err)
	}
	cur := ipb.NewCuratorClient(eph)
	_, err = cur.RegisterNode(ctx, &ipb.RegisterNodeRequest{
		RegisterTicket: register.RegisterTicket,
	})
	if err != nil {
		return fmt.Errorf("register call failed: %w", err)
	}

	// Attempt to commit in a loop. This will succeed once the node is approved.
	supervisor.Logger(ctx).Infof("Registering: success, attempting to commit...")
	var certBytes, caCertBytes []byte
	for {
		resC, err := cur.CommitNode(ctx, &ipb.CommitNodeRequest{
			ClusterUnlockKey: cuk,
		})
		if err == nil {
			supervisor.Logger(ctx).Infof("Registering: Commit succesfull, received certificate")
			certBytes = resC.NodeCertificate
			caCertBytes = resC.CaCertificate
			break
		}
		supervisor.Logger(ctx).Infof("Registering: Commit failed, retrying: %v", err)
		time.Sleep(time.Second)
	}

	// Node is now UP, build client and report it to downstream code.
	creds, err := identity.NewNodeCredentials(priv, certBytes, caCertBytes)
	if err != nil {
		return fmt.Errorf("NewNodeCredentials failed after receiving certificate from cluster: %w", err)
	}
	status := Status{
		State:       cpb.ClusterState_CLUSTER_STATE_HOME,
		Credentials: creds,
	}
	m.status.Set(status)
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
