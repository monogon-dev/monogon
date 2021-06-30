package curator

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus/client"
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
//
// TODO(q3k): don't require privkey, but that needs some //metropolis/pkg/pki changes first.
func BootstrapNodeCredentials(ctx context.Context, etcd client.Namespaced, priv, pub []byte) (ca, node []byte, err error) {
	id := NodeID(pub)

	ca, _, err = pkiCA.Ensure(ctx, etcd)
	if err != nil {
		err = fmt.Errorf("when ensuring CA: %w", err)
		return
	}
	nodeCert := pkiNamespace.New(pkiCA, "", pki.Server([]string{id}, nil))
	nodeCert.UseExistingKey(priv)
	node, _, err = nodeCert.Ensure(ctx, etcd)
	if err != nil {
		err = fmt.Errorf("when ensuring node cert: %w", err)
		return
	}

	return
}

// BootstrapStore saves the Node into etcd, without regard for any other cluster
// state and directly using a given etcd client.
//
// This can only be used by the cluster bootstrap logic.
func (n *Node) BootstrapStore(ctx context.Context, etcd client.Namespaced) error {
	// Currently the only flow to store a node to etcd is a write-once flow:
	// once a node is created, it cannot be deleted or updated. In the future,
	// flows to change cluster node roles might be introduced (ie. to promote
	// nodes to consensus members, etc).
	key := n.etcdPath()
	msg := n.proto()
	nodeRaw, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal node: %w", err)
	}

	res, err := etcd.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(key), "=", 0),
	).Then(
		clientv3.OpPut(key, string(nodeRaw)),
	).Commit()
	if err != nil {
		return fmt.Errorf("failed to store node: %w", err)
	}

	if !res.Succeeded {
		return fmt.Errorf("attempted to re-register node (unsupported flow)")
	}
	return nil
}
