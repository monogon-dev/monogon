package curator

import (
	"context"
	"errors"
	"io/ioutil"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// TestListenerSwitch exercises the curator listener's
// switch-to-different-implementation functionality, notably ensuring that the
// correct implementation is called and that the context is canceled accordingly
// on implementation switch.
//
// It does not test the gRPC listener socket itself and the actual
// implementations - that is deferred to curator functionality tests.
func TestListenerSwitch(t *testing.T) {
	// Create ephemeral directory for curator and place it into /tmp.
	dir := localstorage.EphemeralCuratorDirectory{}
	// Force usage of /tmp as temp directory root, otherwsie TMPDIR from Bazel
	// returns a path long enough that socket binds in the localstorage fail
	// (as socket names are limited to 108 characters).
	tmp, err := ioutil.TempDir("/tmp", "curator-test-*")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	err = declarative.PlaceFS(&dir, tmp)
	if err != nil {
		t.Fatalf("PlaceFS: %v", err)
	}

	// Create test event value.
	var val memory.Value

	// Create DUT listener.
	l := &listener{
		etcd:      nil,
		directory: &dir,
		electionWatch: func() electionWatcher {
			return electionWatcher{
				Watcher: val.Watch(),
			}
		},
		dispatchC: make(chan dispatchRequest),
	}

	// Start listener under supervisor.
	supervisor.TestHarness(t, l.run)

	// Begin with a follower.
	val.Set(electionStatus{
		follower: &electionStatusFollower{},
	})

	// Context for this test.
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Simulate a request context.
	ctxR, ctxRC := context.WithCancel(ctx)

	// Check that canceling the request unblocks a pending dispatched call.
	errC := make(chan error)
	go func() {
		errC <- l.callImpl(ctxR, func(ctx context.Context, impl services) error {
			<-ctx.Done()
			return ctx.Err()
		})
	}()
	ctxRC()
	err = <-errC
	if err == nil || !errors.Is(err, context.Canceled) {
		t.Fatalf("callImpl context should have returned context error, got %v", err)
	}

	// Check that switching implementations unblocks a pending dispatched call.
	scheduledC := make(chan struct{})
	go func() {
		errC <- l.callImpl(ctx, func(ctx context.Context, impl services) error {
			close(scheduledC)
			<-ctx.Done()
			return ctx.Err()
		})
	}()
	// Block until we actually start executing on the follower listener.
	<-scheduledC
	// Switch over to leader listener.
	val.Set(electionStatus{
		leader: &electionStatusLeader{},
	})
	// Check returned error.
	err = <-errC
	if err == nil {
		t.Fatalf("callImpl context should have returned error, got nil")
	}
	if serr, ok := status.FromError(err); !ok || serr.Code() != codes.Unavailable {
		t.Fatalf("callImpl context should have returned unavailable, got %v", err)
	}
}
