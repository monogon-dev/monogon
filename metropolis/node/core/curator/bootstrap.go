package curator

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/osbase/pki"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
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

// BootstrapNodeFinish saves the given Node and initial cluster owner pubkey
// into etcd, without regard for any other cluster state and directly using
// a given etcd client.
//
// This is ran by the cluster bootstrap workflow to finish bootstrapping a
// cluster - afterwards, this cluster will be ready to serve.
//
// This must only be used by the cluster bootstrap logic. It is idempotent, thus
// can be called repeatedly in case of intermittent failures in the bootstrap
// logic.
func BootstrapNodeFinish(ctx context.Context, etcd client.Namespaced, node *Node, ownerKey []byte, cluster *Cluster) (caCertBytes, nodeCertBytes []byte, err error) {
	// Workaround for pkiCA being a global, but BootstrapNodeFinish being called for
	// multiple, different etcd instances in tests. Doing this ensures that we
	// always resynchronize with etcd, ie. not keep certificates loaded in memory
	// even though the underlying certificate in etcd changed.
	//
	// TODO(q3k): pass pkiCA explicitly, eg. within a curator object?
	pkiCA.PrivateKey = nil
	pkiCA.PublicKey = nil

	// Issue CA and node certificates.
	caCertBytes, err = pkiCA.Ensure(ctx, etcd)
	if err != nil {
		return nil, nil, fmt.Errorf("when ensuring CA: %w", err)
	}
	nodeCert := &pki.Certificate{
		Namespace: &pkiNamespace,
		Issuer:    pkiCA,
		Template:  identity.NodeCertificate(node.ID()),
		Mode:      pki.CertificateExternal,
		PublicKey: node.pubkey,
		Name:      fmt.Sprintf("node-%s", node.ID()),
	}
	nodeCertBytes, err = nodeCert.Ensure(ctx, etcd)
	if err != nil {
		err = fmt.Errorf("when ensuring node cert: %w", err)
		return
	}

	nodePath, err := node.etcdNodePath()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get node key: %w", err)
	}
	nodeRaw, err := proto.Marshal(node.proto())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal node: %w", err)
	}
	ownerRaw, err := proto.Marshal(&ppb.InitialOwner{
		PublicKey: ownerKey,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal initial owner: %w", err)
	}
	joinKeyPath, err := node.etcdJoinKeyPath()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get join key: %w", err)
	}

	clusterProto, err := cluster.proto()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to serialize initial cluster configuration: %w", err)
	}
	clusterRaw, err := proto.Marshal(clusterProto)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal initial cluster configuration: %w", err)
	}

	// We don't care about the result's success - this is idempotent.
	_, err = etcd.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(nodePath), "=", 0),
		clientv3.Compare(clientv3.CreateRevision(initialOwnerEtcdPath), "=", 0),
		clientv3.Compare(clientv3.CreateRevision(joinKeyPath), "=", 0),
		clientv3.Compare(clientv3.CreateRevision(clusterConfigurationKey), "=", 0),
	).Then(
		clientv3.OpPut(nodePath, string(nodeRaw)),
		clientv3.OpPut(initialOwnerEtcdPath, string(ownerRaw)),
		clientv3.OpPut(joinKeyPath, node.ID()),
		clientv3.OpPut(clusterConfigurationKey, string(clusterRaw)),
	).Commit()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to store initial cluster state: %w", err)
	}

	return
}
