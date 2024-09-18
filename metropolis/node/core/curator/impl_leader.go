package curator

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
)

// leaderState is the transient state of the Curator leader. All the
// information kept inside is lost whenever another leader is elected.
type leaderState struct {
	// heartbeatTimestamps maps node IDs to monotonic clock timestamps matching
	// the last corresponding node heartbeats received by the current Curator
	// leader.
	heartbeatTimestamps sync.Map

	// startTs is a local monotonic clock timestamp associated with this node's
	// assumption of Curator leadership.
	startTs time.Time

	// clusternetCache maps wireguard public keys (as strings) into node IDs. It is
	// used to detect possibly re-used WireGuard public keys without having to get
	// all nodes from etcd.
	clusternetCache map[string]string
}

// leadership represents the curator leader's ability to perform actions as a
// leader. It is available to all services implemented by the leader.
type leadership struct {
	// lockKey is the etcd key which backs this leader-elected instance.
	lockKey string
	// lockRev is the revision at which lockKey was created. The leader will use it
	// in combination with lockKey to ensure all mutations/reads performed to etcd
	// succeed only if this leader election is still current.
	lockRev int64
	// leaderID is the node ID of this curator's node, ie. the one acting as a
	// curator leader.
	leaderID string
	// etcd is the etcd client in which curator data and leader election state is
	// stored.
	etcd client.Namespaced

	// muNodes guards any changes to nodes, and prevents race conditions where the
	// curator performs a read-modify-write operation to node data. The curator's
	// leadership ensure no two curators run simultaneously, and this lock ensures
	// no two parallel curator operations race eachother.
	//
	// This lock has to be taken any time such RMW operation takes place when not
	// additionally guarded using etcd transactions.
	muNodes sync.Mutex

	consensusStatus *consensus.Status
	consensus       consensus.ServiceHandle

	// muRegisterTicket guards changes to the register ticket. Its usage semantics
	// are the same as for muNodes, as described above.
	muRegisterTicket sync.Mutex

	// ls contains the current leader's non-persistent local state.
	ls leaderState
}

var (
	// errLostLeadership is returned by txnAsLeader if the transaction got canceled
	// because leadership was lost.
	errLostLeadership = errors.New("lost leadership")
)

// txnAsLeader performs an etcd transaction guarded by continued leadership.
// errLostLeadership will be returned as an error in case the leadership is lost.
func (l *leadership) txnAsLeader(ctx context.Context, ops ...clientv3.Op) (*clientv3.TxnResponse, error) {
	var opsStr []string
	for _, op := range ops {
		opstr := "unk"
		switch {
		case op.IsGet():
			opstr = "get"
		case op.IsDelete():
			opstr = "delete"
		case op.IsPut():
			opstr = "put"
		}
		opsStr = append(opsStr, fmt.Sprintf("%s: %s", opstr, op.KeyBytes()))
	}
	rpc.Trace(ctx).Printf("txnAsLeader(%s)...", strings.Join(opsStr, ","))
	resp, err := l.etcd.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(l.lockKey), "=", l.lockRev),
	).Then(ops...).Commit()
	if err != nil {
		rpc.Trace(ctx).Printf("txnAsLeader(...): failed: %v", err)
		return nil, fmt.Errorf("when running leader transaction: %w", err)
	}
	if !resp.Succeeded {
		// Transaction failed because leadership was lost. Log error with
		// detailed information about lock key, expected revision and found
		// revision to aid debugging.
		checkRes, err := l.etcd.Get(ctx, l.lockKey)
		var lockRev string
		if err != nil {
			lockRev = fmt.Sprintf("couldn't check: %v", err)
		} else {
			if len(checkRes.Kvs) > 0 {
				lockRev = fmt.Sprintf("%d", checkRes.Kvs[0].CreateRevision)
			} else {
				lockRev = "no revision?"
			}
		}
		rpc.Trace(ctx).Printf("txnAsLeader(...): rejected (lost leadership (key %s should've been at rev %d, is at rev %s)", l.lockKey, l.lockRev, lockRev)
		return nil, errLostLeadership
	}
	rpc.Trace(ctx).Printf("txnAsLeader(...): ok")
	return resp, nil
}

// rpcError attempts to convert a given error to a high-level error that can be
// directly exposed to RPC clients. If false is returned, the error was not
// converted and is returned verbatim.
func rpcError(err error) (error, bool) {
	if errors.Is(err, errLostLeadership) {
		return status.Error(codes.Unavailable, "lost leadership"), true
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Error(codes.DeadlineExceeded, err.Error()), true
	}
	if errors.Is(err, context.Canceled) {
		return status.Error(codes.Canceled, err.Error()), true
	}
	return err, false
}

// curatorLeader implements the curator acting as the elected leader of a
// cluster. It performs direct reads/writes from/to etcd as long as it remains
// leader.
//
// Its made up of different subcomponents implementing gRPC services, each of
// which has access to the leadership structure.
type curatorLeader struct {
	leaderCurator
	leaderAAA
	leaderManagement
	leaderBackground
}

func newCuratorLeader(l *leadership, node *identity.Node) *curatorLeader {
	// Mark the start of this leader's tenure.
	l.ls.startTs = time.Now()

	return &curatorLeader{
		leaderCurator{leadership: l},
		leaderAAA{leadership: l},
		leaderManagement{leadership: l, node: node},
		leaderBackground{leadership: l},
	}
}
