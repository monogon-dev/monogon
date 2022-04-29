package cluster

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
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
	//// Do a validation pass on the provided NodeParameters.Register data, fail early
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
				return fmt.Errorf("NodeParameters.ClusterDirectory.Nodes[%d].Addresses[%d] (%q) invalid: not a parseable address", i, j, add.Host)
			}
		}
		if len(node.PublicKey) != ed25519.PublicKeySize {
			return fmt.Errorf("NodeParameters.ClusterDirectory.Nodes[%d] invalid: PublicKey invalid length or not set", i)
		}
	}

	// Strip the initial ClusterDirectory of any node public keys that might have
	// been included, as it can't be relied on beyond providing cluster endpoint
	// addresses, considering its untrusted origin (ESP). This explicitly enforces
	// suggested usage described in ClusterDirectory's protofile.
	for i, _ := range register.ClusterDirectory.Nodes {
		register.ClusterDirectory.Nodes[i].PublicKey = nil
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
	logClusterDirectory(ctx, register.ClusterDirectory)

	// Mount new storage with generated CUK, MountNew will save NUK into sc, to be
	// saved into the ESP after successful registration.
	var sc ppb.SealedConfiguration
	state.configuration = &sc
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
	r, err := curatorRemote(register.ClusterDirectory)
	if err != nil {
		return fmt.Errorf("while picking a Curator endpoint: %w", err)
	}
	ephCreds, err := rpc.NewEphemeralCredentials(priv, ca)
	if err != nil {
		return fmt.Errorf("could not create ephemeral credentials: %w", err)
	}
	eph, err := grpc.Dial(r, grpc.WithTransportCredentials(ephCreds))
	if err != nil {
		return fmt.Errorf("could not create ephemeral client to %q: %w", r, err)
	}
	cur := ipb.NewCuratorClient(eph)

	// Generate Join Credentials. The private key will be stored in
	// SealedConfiguration only if RegisterNode succeeds.
	jpub, jpriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate join keypair: %w", err)
	}
	sc.JoinKey = jpriv
	supervisor.Logger(ctx).Infof("Registering: join public key: %s", hex.EncodeToString([]byte(jpub)))

	// Register this node.
	//
	// MVP: From this point on forward, we have very little resiliency to failure,
	// we barely scrape by with usage of retry loops. We should break down all the
	// logic into some sort of state machine where we can atomically make progress
	// on each of the stages and get rid of the retry loops. The cluster enrolment
	// code should let us do this quite easily.
	_, err = cur.RegisterNode(ctx, &ipb.RegisterNodeRequest{
		RegisterTicket: register.RegisterTicket,
		JoinKey:        jpub,
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
	m.roleServer.ProvideRegisterData(*creds, register.ClusterDirectory)

	// Save Node Credentials
	if err = creds.Save(&m.storageRoot.Data.Node.Credentials); err != nil {
		return fmt.Errorf("while saving node credentials: %w", err)
	}
	// Save the Cluster Directory into the ESP.
	cdirRaw, err := proto.Marshal(register.ClusterDirectory)
	if err != nil {
		return fmt.Errorf("couldn't marshal ClusterDirectory: %w", err)
	}
	if err = m.storageRoot.ESP.Metropolis.ClusterDirectory.Write(cdirRaw, 0644); err != nil {
		return err
	}
	// Include the Cluster CA in Sealed Configuration.
	sc.ClusterCa = register.CaCertificate
	// Save Cluster CA, NUK and Join Credentials into Sealed Configuration.
	if err = m.storageRoot.ESP.Metropolis.SealedConfiguration.SealSecureBoot(&sc); err != nil {
		return err
	}
	unix.Sync()

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
