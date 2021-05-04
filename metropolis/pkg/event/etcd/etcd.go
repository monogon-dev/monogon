package etcd

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"go.etcd.io/etcd/clientv3"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/pkg/event"
)

var (
	// Type assert that *Value implements event.ValueWatcher. We do this
	// artificially, as there currently is no code path that needs this to be
	// strictly true.  However, users of this library might want to rely on the
	// Value type instead of particular Value implementations.
	_ event.ValueWatch = &Value{}
)

// Value is an 'Event Value' backed in an etcd cluster, accessed over an
// etcd client. This is a stateless handle and can be copied and shared across
// goroutines.
type Value struct {
	etcd    client.Namespaced
	key     string
	decoder BytesDecoder
}

// NewValue creates a new Value for a given key in an etcd client. The
// given decoder will be used to convert bytes retrieved from etcd into the
// interface{} value retrieved by Get by this value's watcher.
func NewValue(etcd client.Namespaced, key string, decoder BytesDecoder) *Value {
	return &Value{
		etcd:    etcd,
		key:     key,
		decoder: decoder,
	}
}

// BytesDecoder is a function that converts bytes retrieved from etcd into an
// end-user facing value. If an error is returned, the Get call performed on a
// watcher configured with this decoder will fail, swallowing that particular
// update, but the watcher will continue to work.
// Any provided BytesDecoder implementations must be safe to copy.
type BytesDecoder = func(data []byte) (interface{}, error)

// NoDecoder is a no-op decoder which passes through the retrieved bytes as a
// []byte type to the user.
func NoDecoder(data []byte) (interface{}, error) {
	return data, nil
}

func (e *Value) Watch() event.Watcher {
	ctx, ctxC := context.WithCancel(context.Background())
	return &watcher{
		Value: *e,

		ctx:  ctx,
		ctxC: ctxC,

		getSem: make(chan struct{}, 1),
	}
}

type watcher struct {
	// Value copy, used to configure the behaviour of this watcher.
	Value

	// ctx is the context that expresses the liveness of this watcher. It is
	// canceled when the watcher is closed, and the etcd Watch hangs off of it.
	ctx  context.Context
	ctxC context.CancelFunc

	// getSem is a semaphore used to limit concurrent Get calls and throw an
	// error if concurrent access is attempted.
	getSem chan struct{}

	// backlogged is the value retrieved by an initial KV Get from etcd that
	// should be returned at the next opportunity, or nil if there isn't any.
	backlogged *[]byte
	// prev is the revision of a previously retrieved value within this
	// watcher, or nil if there hasn't been any.
	prev *int64
	// wc is the etcd watch channel, or nil if no channel is yet open.
	wc clientv3.WatchChan

	// testRaceWG is an optional WaitGroup that, if set, will be waited upon
	// after the initial KV value retrieval, but before the watch is created.
	// This is only used for testing.
	testRaceWG *sync.WaitGroup
	// testSetupWG is an optional WaitGroup that, if set, will be waited upon
	// after the etcd watch is created.
	// This is only used for testing.
	testSetupWG *sync.WaitGroup
}

func (w *watcher) setup(ctx context.Context) error {
	if w.wc != nil {
		return nil
	}

	// First, check if some data under this key already exists.
	// We use an exponential backoff and retry here as the initial Get can fail
	// if the cluster is unstable (eg. failing over). We only fail the retry if
	// the context expires.
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 0
	err := backoff.Retry(func() error {
		get, err := w.etcd.Get(ctx, w.key)
		if err != nil {
			return fmt.Errorf("when retrieving initial value: %w", err)
		}
		w.prev = &get.Header.Revision
		if len(get.Kvs) != 0 {
			// Assert that the etcd API is behaving as expected.
			if len(get.Kvs) > 1 {
				panic("More than one key returned in unary GET response")
			}
			// If an existing value is present, backlog it and set the prev value
			// accordingly.
			kv := get.Kvs[0]
			w.backlogged = &kv.Value
		} else {
			w.backlogged = nil
		}
		return nil

	}, backoff.WithContext(bo, ctx))

	if w.testRaceWG != nil {
		w.testRaceWG.Wait()
	}
	if err != nil {
		return err
	}

	var watchOpts []clientv3.OpOption
	if w.prev != nil {
		watchOpts = append(watchOpts, clientv3.WithRev(*w.prev+1))
	} else {
	}
	w.wc = w.etcd.Watch(w.ctx, w.key, watchOpts...)

	if w.testSetupWG != nil {
		w.testSetupWG.Wait()
	}
	return nil
}

func (w *watcher) get(ctx context.Context) ([]byte, error) {
	// Return backlogged value, if present.
	if w.backlogged != nil {
		value := *w.backlogged
		w.backlogged = nil
		return value, nil
	}

	// Keep watching for a watch event.
	var event *clientv3.Event
	for {
		var resp *clientv3.WatchResponse
		select {
		case r := <-w.wc:
			resp = &r
		case <-ctx.Done():
			return nil, ctx.Err()
		}

		if resp.Canceled {
			// Only allow for watches to be canceled due to context
			// cancellations. Any other error is something we need to handle,
			// eg. a client close or compaction error.
			if errors.Is(resp.Err(), ctx.Err()) {
				return nil, fmt.Errorf("watch canceled: %w", resp.Err())
			}

			// Attempt to reconnect.
			w.wc = nil
			w.setup(ctx)
			continue
		}

		if len(resp.Events) < 1 {
			continue
		}

		event = resp.Events[len(resp.Events)-1]
		break
	}

	w.prev = &event.Kv.ModRevision
	// Return deletions as nil, and new values as their content.
	switch event.Type {
	case clientv3.EventTypeDelete:
		return nil, nil
	case clientv3.EventTypePut:
		return event.Kv.Value, nil
	default:
		return nil, fmt.Errorf("invalid event type %v", event.Type)
	}
}

// Get implements the Get method of the Watcher interface.
// It can return an error in two cases:
//  - the given context is canceled (in which case, the given error will wrap
//    the context error)
//  - the watcher's BytesDecoder returned an error (in which case the error
//    returned by the BytesDecoder will be returned verbatim)
// Note that transient and permanent etcd errors are never returned, and the
// Get call will attempt to recover from these errors as much as possible. This
// also means that the user of the Watcher will not be notified if the
// underlying etcd client disconnects from the cluster, or if the cluster loses
// quorum.
// TODO(q3k): implement leases to allow clients to be notified when there are
// transient cluster/quorum/partition errors, if needed.
// TODO(q3k): implement internal, limited buffering for backlogged data not yet
// consumed by client, as etcd client library seems to use an unbound buffer in
// case this happens ( see: watcherStream.buf in clientv3).
func (w *watcher) Get(ctx context.Context) (interface{}, error) {
	select {
	case w.getSem <- struct{}{}:
	default:
		return nil, fmt.Errorf("cannot Get() concurrently on a single waiter")
	}
	defer func() {
		<-w.getSem
	}()

	// Early check for context cancelations, preventing spurious contact with etcd
	// if there's no need to.
	if w.ctx.Err() != nil {
		return nil, w.ctx.Err()
	}

	if err := w.setup(ctx); err != nil {
		return nil, fmt.Errorf("when setting up watcher: %w", err)
	}

	value, err := w.get(ctx)
	if err != nil {
		return nil, fmt.Errorf("when watching for new value: %w", err)
	}
	return w.decoder(value)
}

func (w *watcher) Close() error {
	w.ctxC()
	return nil
}
