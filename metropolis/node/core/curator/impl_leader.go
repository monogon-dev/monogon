package curator

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
)

// leadership represents the curator leader's ability to perform actions as a
// leader. It is available to all services implemented by the leader.
type leadership struct {
	// lockKey is the etcd key which backs this leader-elected instance.
	lockKey string
	// lockRev is the revision at which lockKey was created. The leader will use it
	// in combination with lockKey to ensure all mutations/reads performed to etcd
	// succeed only if this leader election is still current.
	lockRev int64
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

	consensus consensus.ServiceHandle

	// muRegisterTicket guards changes to the register ticket. Its usage semantics
	// are the same as for muNodes, as described above.
	muRegisterTicket sync.Mutex
}

var (
	// lostLeadership is returned by txnAsLeader if the transaction got canceled
	// because leadership was lost.
	lostLeadership = errors.New("lost leadership")
)

// txnAsLeader performs an etcd transaction guarded by continued leadership.
// lostLeadership will be returned as an error in case the leadership is lost.
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
		rpc.Trace(ctx).Printf("txnAsLeader(...): rejected (lost leadership)")
		return nil, lostLeadership
	}
	rpc.Trace(ctx).Printf("txnAsLeader(...): ok")
	return resp, nil
}

// rpcError attempts to convert a given error to a high-level error that can be
// directly exposed to RPC clients. If false is returned, the error was not
// converted and is returned verbatim.
func rpcError(err error) (error, bool) {
	if errors.Is(err, lostLeadership) {
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
}

func newCuratorLeader(l *leadership, node *identity.Node) *curatorLeader {
	return &curatorLeader{
		leaderCurator{leadership: l},
		leaderAAA{leadership: l},
		leaderManagement{leadership: l, node: node},
	}
}
