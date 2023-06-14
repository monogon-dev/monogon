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

	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	"source.monogon.dev/metropolis/pkg/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
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

	// Validation passed, let's start working on registering us into the cluster.
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate node keypair: %w", err)
	}

	// Build resolver used by the register process, authenticating with ephemeral
	// credentials. Once the join is complete, the rolesever will start its own
	// long-term resolver.
	rctx, rctxC := context.WithCancel(ctx)
	defer rctxC()
	r := resolver.New(rctx, resolver.WithoutCuratorUpdater(), resolver.WithLogger(func(f string, args ...interface{}) {
		supervisor.Logger(ctx).WithAddedStackDepth(1).Infof(f, args...)
	}))
	addedNodes := 0
	for _, node := range register.ClusterDirectory.Nodes {
		if len(node.Addresses) == 0 {
			continue
		}
		// MVP: handle curators at non-default ports
		r.AddEndpoint(resolver.NodeAtAddressWithDefaultPort(node.Addresses[0].Host))
		addedNodes += 1
	}
	if addedNodes == 0 {
		return fmt.Errorf("no remote node available, cannot register into cluster")
	}

	ephCreds, err := rpc.NewEphemeralCredentials(priv, ca)
	if err != nil {
		return fmt.Errorf("could not create ephemeral credentials: %w", err)
	}
	eph, err := grpc.Dial(resolver.MetropolisControlAddress, grpc.WithTransportCredentials(ephCreds), grpc.WithResolvers(r))
	if err != nil {
		return fmt.Errorf("could not dial cluster with ephemeral credentials: %w", err)
	}
	cur := ipb.NewCuratorClient(eph)

	// TODO(q3k): allow node to pick storage security per given policy

	// Tell the user what we're doing.
	supervisor.Logger(ctx).Infof("Registering into existing cluster.")
	supervisor.Logger(ctx).Infof("  Cluster CA public key: %s", hex.EncodeToString(ca.PublicKey.(ed25519.PublicKey)))
	supervisor.Logger(ctx).Infof("  Node public key: %s", hex.EncodeToString(pub))
	supervisor.Logger(ctx).Infof("  Register Ticket: %s", hex.EncodeToString(register.RegisterTicket))

	// Generate Join Credentials. The private key will be stored in
	// SealedConfiguration only if RegisterNode succeeds.
	jpub, jpriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate join keypair: %w", err)
	}
	supervisor.Logger(ctx).Infof("Registering: join public key: %s", hex.EncodeToString([]byte(jpub)))

	// Register this node.
	//
	// MVP: From this point on forward, we have very little resiliency to failure,
	// we barely scrape by with usage of retry loops. We should break down all the
	// logic into some sort of state machine where we can atomically make progress
	// on each of the stages and get rid of the retry loops. The cluster enrolment
	// code should let us do this quite easily.
	res, err := cur.RegisterNode(ctx, &ipb.RegisterNodeRequest{
		RegisterTicket: register.RegisterTicket,
		JoinKey:        jpub,
		HaveLocalTpm:   m.haveTPM,
	})
	if err != nil {
		return fmt.Errorf("register call failed: %w", err)
	}
	storageSecurity := res.RecommendedNodeStorageSecurity

	// Mount new storage with generated CUK, MountNew will save NUK into sc, to be
	// saved into the ESP after successful registration.
	var sc ppb.SealedConfiguration
	supervisor.Logger(ctx).Infof("Registering: mounting new storage...")
	cuk, err := m.storageRoot.Data.MountNew(&sc, storageSecurity)
	if err != nil {
		return fmt.Errorf("could not make and mount data partition: %w", err)
	}
	sc.JoinKey = jpriv

	supervisor.Logger(ctx).Infof("Storage Security: cluster policy: %s", res.ClusterConfiguration.StorageSecurityPolicy)
	supervisor.Logger(ctx).Infof("Storage Security: node: %s", storageSecurity)
	supervisor.Logger(ctx).Infof("TPM: cluster TPM mode: %s", res.ClusterConfiguration.TpmMode)
	supervisor.Logger(ctx).Infof("TPM: node TPM usage: %v", res.TpmUsage)

	// Attempt to commit in a loop. This will succeed once the node is approved.
	supervisor.Logger(ctx).Infof("Registering: success, attempting to commit...")
	var certBytes, caCertBytes []byte
	for {
		resC, err := cur.CommitNode(ctx, &ipb.CommitNodeRequest{
			ClusterUnlockKey: cuk,
			StorageSecurity:  storageSecurity,
		})
		if err == nil {
			supervisor.Logger(ctx).Infof("Registering: Commit successful, received certificate")
			certBytes = resC.NodeCertificate
			caCertBytes = resC.CaCertificate
			break
		}
		supervisor.Logger(ctx).Infof("Registering: Commit failed, retrying: %v", err)
		time.Sleep(time.Second)
	}

	// Node is now UP, build client/credentials and save them to ESP.
	creds, err := identity.NewNodeCredentials(priv, certBytes, caCertBytes)
	if err != nil {
		return fmt.Errorf("NewNodeCredentials failed after receiving certificate from cluster: %w", err)
	}

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
	if err = m.storageRoot.ESP.Metropolis.SealedConfiguration.SealSecureBoot(&sc, res.TpmUsage); err != nil {
		return err
	}
	unix.Sync()

	// All synced up, we can now let downstream know about the creds, which in turn
	// will start heartbeating the cluster and running role-specific jobs.
	m.roleServer.ProvideRegisterData(*creds, register.ClusterDirectory)

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
