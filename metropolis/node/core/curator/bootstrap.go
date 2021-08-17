package curator

import (
	"context"
	"crypto/ed25519"
	"fmt"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/pkg/pki"
)

// bootstrap.go contains functions specific for integration between the curator
// and cluster bootstrap code (//metropolis/node/core/cluster).
//
// These functions must only be called by the bootstrap code, and are
// effectively well-controlled abstraction leaks. An alternative would be to
// rework the curator API to explicitly support a well-contained and
// well-defined bootstrap procedure, formalized within bootstrap-specific types.
// However, that seems to not be worth the effort for a tightly coupled single
// consumer like the bootstrap code.

// BootstrapNodeCredentials creates node credentials for the first node in a
// cluster. It can only be called by cluster bootstrap code. It returns the
// generated x509 CA and node certificates.
func BootstrapNodeCredentials(ctx context.Context, etcd client.Namespaced, pubkey ed25519.PublicKey) (ca, node []byte, err error) {
	id := NodeID(pubkey)

	ca, err = pkiCA.Ensure(ctx, etcd)
	if err != nil {
		err = fmt.Errorf("when ensuring CA: %w", err)
		return
	}
	nodeCert := &pki.Certificate{
		Namespace: &pkiNamespace,
		Issuer:    pkiCA,
		Template:  pki.Server([]string{id}, nil),
		Mode:      pki.CertificateExternal,
		PublicKey: pubkey,
		Name:      fmt.Sprintf("node-%s", id),
	}
	node, err = nodeCert.Ensure(ctx, etcd)
	if err != nil {
		err = fmt.Errorf("when ensuring node cert: %w", err)
		return
	}

	return
}

// BootstrapFinish saves the given Node and initial cluster owner pubkey into
// etcd, without regard for any other cluster state and directly using a given
// etcd client.
//
// This is ran by the cluster bootstrap workflow to finish bootstrapping a
// cluster - afterwards, this cluster will be ready to serve.
//
// This can only be used by the cluster bootstrap logic, and may only be called
// once. It's guaranteed to either succeed fully or fail fully, without leaving
// the cluster in an inconsistent state.
func BootstrapFinish(ctx context.Context, etcd client.Namespaced, initialNode *Node, pubkey []byte) error {
	nodeKey := initialNode.etcdPath()
	nodeRaw, err := proto.Marshal(initialNode.proto())
	if err != nil {
		return fmt.Errorf("failed to marshal node: %w", err)
	}

	owner := &ppb.InitialOwner{
		PublicKey: pubkey,
	}
	ownerKey := initialOwnerEtcdPath
	ownerRaw, err := proto.Marshal(owner)
	if err != nil {
		return fmt.Errorf("failed to marshal iniail owner: %w", err)
	}

	res, err := etcd.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(nodeKey), "=", 0),
		clientv3.Compare(clientv3.CreateRevision(ownerKey), "=", 0),
	).Then(
		clientv3.OpPut(nodeKey, string(nodeRaw)),
		clientv3.OpPut(ownerKey, string(ownerRaw)),
	).Commit()
	if err != nil {
		return fmt.Errorf("failed to store initial cluster state: %w", err)
	}

	if !res.Succeeded {
		return fmt.Errorf("cluster already bootstrapped")
	}
	return nil
}
