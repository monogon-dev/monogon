package curator

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/pkg/combinectx"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// listener is the curator runnable responsible for listening for gRPC
// connections and forwarding them over to one of two possible implementations:
// - a local leader implementation which is backed by etcd
// - a follower implementation that forwards the RPCs over to a remote leader.
//
// Its goal is to make any switches over between leader and follower painless to
// the gRPC callers.
// Any pending calls will be canceled with UNAVAILABLE and an error message
// describing the fact that the implementation has been switched over.
// The gRPC sockets will always be listening for connections, and block until
// able to serve a request (either locally or by forwarding).
// No retries will be attempted on switchover, as some calls might not be
// idempotent and the caller is better equipped to know when to retry.
type listener struct {
	// etcd is a client to the locally running consensus (etcd) server which is used
	// both for storing lock/leader election status and actual Curator data.
	etcd client.Namespaced
	// directory is the ephemeral directory in which the local gRPC socket will
	// be available for node-local consumers.
	directory *localstorage.EphemeralCuratorDirectory
	// electionWatch is a function that returns an active electionWatcher for the
	// listener to use when determining local leadership. As the listener may
	// restart on error, this factory-function is used instead of an electionWatcher
	// directly.
	electionWatch func() electionWatcher

	dispatchC chan dispatchRequest
}

// dispatcher is the 'listener dispatcher', the listener's runnable responsible
// for keeping track of the currently selected curator implementation and
// switching over when necessary.
//
// It listens for 'dispatch' requests from the listener's RPC handlers and
// returns a Curator implementation that should be used to handle this request,
// alongside a context expressing the lifespan of this implementation.
func (l *listener) dispatcher(ctx context.Context) error {
	supervisor.Logger(ctx).Info("Dispatcher starting...")

	// Start with an empty 'active target'. This will be populated before the
	// first dispatch request is served.
	t := activeTarget{}
	w := l.electionWatch()

	supervisor.Signal(ctx, supervisor.SignalHealthy)

	// Channel containing electionStatus updates from value.
	c := make(chan *electionStatus)
	defer close(c)

	go func() {
		// Wait for initial status.
		s, ok := <-c
		if !ok {
			return
		}
		t.switchTo(ctx, l, s)

		// Respond to requests and status updates.
		for {
			select {
			case r := <-l.dispatchC:
				// Handle request.
				r.resC <- listenerTarget{
					ctx:  *t.ctx,
					impl: t.impl,
				}
			case s, ok := <-c:
				// Handle status update, or quit on  status update error.
				if !ok {
					return
				}
				t.switchTo(ctx, l, s)
			}
		}
	}()

	// Convert event electionStatus updates to channel sends. If we cannot retrieve
	// the newest electionStatus, we kill the dispatcher runner.
	for {
		s, err := w.get(ctx)
		if err != nil {
			return fmt.Errorf("could not get newest electionStatus: %w", err)
		}
		c <- s
	}
}

// activeTarget is the active implementation used by the listener dispatcher, or
// nil if none is active yet.
type activeTarget struct {
	// context describing the lifecycle of the active implementation, or nil if the
	// impl is nil.
	ctx *context.Context
	// context cancel function for ctx, or nil if ctx is nil.
	ctxC *context.CancelFunc
	// active Curator implementation, or nil if not yet set up.
	impl apb.CuratorServer
}

// switchTo switches the activeTarget over to a Curator implementation as per
// the electionStatus and leader configuration. If the activeTarget already had
// an implementation set, its associated context is canceled.
func (t *activeTarget) switchTo(ctx context.Context, l *listener, status *electionStatus) {
	if t.ctxC != nil {
		(*t.ctxC)()
	}
	implCtx, implCtxC := context.WithCancel(ctx)
	t.ctx = &implCtx
	t.ctxC = &implCtxC
	if leader := status.leader; leader != nil {
		supervisor.Logger(ctx).Info("Dispatcher switching over to local leader")
		t.impl = &curatorLeader{
			lockKey: leader.lockKey,
			lockRev: leader.lockRev,
			etcd:    l.etcd,
		}
	} else {
		supervisor.Logger(ctx).Info("Dispatcher switching over to follower")
		t.impl = &curatorFollower{}
	}
}

// dispatchRequest is a request sent to the dispatcher by the listener when it
// needs an up to date listenerTarget to run RPC calls against.
type dispatchRequest struct {
	resC chan listenerTarget
}

// listenerTarget is where the listener should forward a given curator RPC. This
// is provided by the listener dispatcher on request (on 'dispatch').
type listenerTarget struct {
	// ctx is the context representing the lifetime of the given impl. It will be
	// canceled when that implementation switches over to a different one.
	ctx context.Context
	// impl is the CuratorServer implementation to which RPCs should be directed
	// according to the dispatcher.
	impl apb.CuratorServer
}

// dispatch contacts the dispatcher to retrieve an up-to-date listenerTarget.
// This target is then used to serve RPCs. The given context is only used to
// time out the dispatch action and does not influence the listenerTarget
// returned.
func (l *listener) dispatch(ctx context.Context) (*listenerTarget, error) {
	req := dispatchRequest{
		// resC is non-blocking to ensure slow dispatch requests do not further cascade
		// into blocking the dispatcher.
		resC: make(chan listenerTarget, 1),
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case l.dispatchC <- req:
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-req.resC:
		return &res, nil
	}
}

// run is the listener runnable. It listens on gRPC sockets and serves RPCs.
func (l *listener) run(ctx context.Context) error {
	supervisor.Logger(ctx).Info("Listener starting...")
	if err := supervisor.Run(ctx, "dispatcher", l.dispatcher); err != nil {
		return fmt.Errorf("when starting dispatcher: %w", err)
	}

	// TODO(q3k): recreate socket if already exists? Is this needed?
	lis, err := net.ListenUnix("unix", &net.UnixAddr{Name: l.directory.ClientSocket.FullPath(), Net: "unix"})
	if err != nil {
		return fmt.Errorf("failed to listen on curator listener socket: %w", err)
	}

	// TODO(q3k): run remote/public gRPC listener.

	srv := grpc.NewServer()
	apb.RegisterCuratorServer(srv, l)

	if err := supervisor.Run(ctx, "local", supervisor.GRPCServer(srv, lis, true)); err != nil {
		return fmt.Errorf("while starting local gRPC listener: %w", err)
	}
	supervisor.Logger(ctx).Info("Listener running.")
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	// Keep the listener running, as its a parent to the gRPC listener.
	<-ctx.Done()
	return ctx.Err()
}

// callImpl gets the newest listenerTarget from the dispatcher, combines the
// given context with the context of the listenerTarget implementation and calls
// the given function with the combined context and implementation.
//
// It is effectively a helper wrapper used by the Curator implementation of the
// listener to run the RPC against the active listenerTarget.
func (l *listener) callImpl(ctx context.Context, f func(ctx context.Context, impl apb.CuratorServer) error) error {
	lt, err := l.dispatch(ctx)
	// dispatch will only return errors on context cancellations.
	if err != nil {
		return err
	}

	ctxCombined := combinectx.Combine(ctx, lt.ctx)
	err = f(ctxCombined, lt.impl)

	// No error occurred? Nothing else to do.
	if err == nil {
		return nil
	}
	cerr := &combinectx.Error{}
	// An error occurred. Was it a context error?
	if errors.As(err, &cerr) {
		if cerr.First() {
			// Request context got canceled. Return inner context error.
			return cerr.Unwrap()
		} else {
			// Leadership changed. Return an UNAVAILABLE so that the request gets retried by
			// the caller if needed.
			return status.Error(codes.Unavailable, "curator backend switched, request can be retried")
		}
	} else {
		// Not a context error, return verbatim.
		return err
	}
}

// curatorWatchServer implements Curator_WatchServer but overrides the context
// of the streaming RPC call with some other context (in this case, the combined
// context from callImpl).
type curatorWatchServer struct {
	grpc.ServerStream
	ctx context.Context
}

func (c *curatorWatchServer) Context() context.Context {
	return c.ctx
}

func (c *curatorWatchServer) Send(m *apb.WatchEvent) error {
	return c.ServerStream.SendMsg(m)
}

// Watch implements the Watch RPC from Curator by dispatching it against the
// correct implementation for this curator instance.
func (l *listener) Watch(req *apb.WatchRequest, srv apb.Curator_WatchServer) error {
	return l.callImpl(srv.Context(), func(ctx context.Context, impl apb.CuratorServer) error {
		return impl.Watch(req, &curatorWatchServer{
			ServerStream: srv,
			ctx:          ctx,
		})
	})
}
