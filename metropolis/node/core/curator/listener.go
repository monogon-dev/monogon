package curator

import (
	"context"
	"fmt"
	"net"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// listener is the curator runnable responsible for listening for gRPC
// connections and forwarding them over to one of two possible implementations:
// - a local leader implementation which is backed by etcd
// - a follower implementation that forwards the RPCs over to a remote leader.
//
// Its goal is to make any switches over between leader and follower painless to
// the gRPC callers. Each incoming RPC first goes into a shim defined directly
// on the listener, then goes on to be passed into either implementation with a
// context that is valid as long as that implementation is current.
//
// Any calls which are pending during a switchover will have their context
// canceled with UNAVAILABLE and an error message describing the fact that the
// implementation has been switched over. The gRPC sockets will always be
// listening for connections, and block until able to serve a request (either
// locally or by forwarding). No retries will be attempted on switchover, as
// some calls might not be idempotent and the caller is better equipped to know
// when to retry.
type listener struct {
	node *identity.NodeCredentials
	// etcd is a client to the locally running consensus (etcd) server which is used
	// both for storing lock/leader election status and actual Curator data.
	etcd client.Namespaced
	// electionWatch is a function that returns an active electionWatcher for the
	// listener to use when determining local leadership. As the listener may
	// restart on error, this factory-function is used instead of an electionWatcher
	// directly.
	electionWatch func() electionWatcher

	consensus consensus.ServiceHandle
}

// listenerTarget is where the listener should forward a given curator RPC. This
// is provided by the listener dispatcher on request (on 'dispatch').
type listenerTarget struct {
	// ctx is the context representing the lifetime of the given impl. It will be
	// canceled when that implementation switches over to a different one.
	ctx context.Context
	// impl is the CuratorServer implementation to which RPCs should be directed
	// according to the dispatcher.
	impl rpc.ClusterServices
}

// run is the listener runnable. It listens on gRPC sockets and serves RPCs.
func (l *listener) run(ctx context.Context) error {
	w := l.electionWatch()
	supervisor.Logger(ctx).Infof("Waiting for election status...")
	st, err := w.get(ctx)
	if err != nil {
		return fmt.Errorf("could not get election status: %w", err)
	}

	sec := rpc.ServerSecurity{
		NodeCredentials: l.node,
	}

	switch {
	case st.leader != nil:
		supervisor.Logger(ctx).Infof("This curator is a leader, starting listener.")
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", node.CuratorServicePort))
		if err != nil {
			return fmt.Errorf("failed to listen on curator socket: %w", err)
		}
		defer lis.Close()

		leader := newCuratorLeader(&leadership{
			lockKey:   st.leader.lockKey,
			lockRev:   st.leader.lockRev,
			etcd:      l.etcd,
			consensus: l.consensus,
		}, &l.node.Node)
		runnable := supervisor.GRPCServer(sec.SetupExternalGRPC(supervisor.MustSubLogger(ctx, "rpc"), leader), lis, true)
		if err := supervisor.Run(ctx, "server", runnable); err != nil {
			return fmt.Errorf("could not run server: %w", err)
		}
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		for {
			st, err := w.get(ctx)
			if err != nil {
				return fmt.Errorf("getting election status after starting listener failed, bailing just in case: %w", err)
			}
			if st.leader == nil {
				return fmt.Errorf("this curator stopped being a leader, quitting")
			}
		}
	case st.follower != nil && st.follower.lock != nil:
		supervisor.Logger(ctx).Infof("This curator is a follower (leader is %q), starting proxy.", st.follower.lock.NodeId)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", node.CuratorServicePort))
		if err != nil {
			return fmt.Errorf("failed to listen on curator socket: %w", err)
		}
		defer lis.Close()

		follower := &curatorFollower{}
		runnable := supervisor.GRPCServer(sec.SetupExternalGRPC(supervisor.MustSubLogger(ctx, "rpc"), follower), lis, true)
		if err := supervisor.Run(ctx, "server", runnable); err != nil {
			return fmt.Errorf("could not run server: %w", err)
		}
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		for {
			st, err := w.get(ctx)
			if err != nil {
				return fmt.Errorf("getting election status after starting listener failed, bailing just in case: %w", err)
			}
			if st.follower == nil {
				return fmt.Errorf("this curator stopped being a follower, quitting")
			}
		}
	default:
		return fmt.Errorf("curator is neither leader nor follower - this is likely transient, restarting listener now")
	}
}
