package etcd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/cenkalti/backoff/v4"
	clientv3 "go.etcd.io/etcd/client/v3"

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
	keyEnd  string
	decoder BytesDecoder
}

type Option struct {
	rangeEnd string
}

// Range creates a Value that is backed a range of etcd key/value pairs from
// 'key' passed to NewValue to 'end' passed to Range.
//
// The key range semantics (ie. lexicographic ordering) are the same as in etcd
// ranges, so for example to retrieve all keys prefixed by `foo/` key should be
// `foo/` and end should be `foo0`.
//
// For any update in the given range, the decoder will be called and its result
// will trigger the return of a Get() call. The decoder should return a type
// that lets the user distinguish which of the multiple objects in the range got
// updated, as the Get() call returns no additional information about the
// location of the retrieved object by itself.
//
// The order of values retrieved by Get() is currently fully arbitrary and must
// not be relied on. It's possible that in the future the order of updates and
// the blocking behaviour of Get will be formalized, but this is not yet the
// case. Instead, the data returned should be treated as eventually consistent
// with the etcd state.
//
// For some uses, it might be necessary to first retrieve all the objects
// contained within the range before starting to block on updates - in this
// case, the BacklogOnly option should be used when calling Get.
func Range(end string) *Option {
	return &Option{
		rangeEnd: end,
	}
}

// NewValue creates a new Value for a given key(s) in an etcd client. The
// given decoder will be used to convert bytes retrieved from etcd into the
// interface{} value retrieved by Get by this value's watcher.
func NewValue(etcd client.Namespaced, key string, decoder BytesDecoder, options ...*Option) *Value {
	res := &Value{
		etcd:    etcd,
		key:     key,
		keyEnd:  key,
		decoder: decoder,
	}

	for _, opt := range options {
		if end := opt.rangeEnd; end != "" {
			res.keyEnd = end
		}
	}

	return res
}

// BytesDecoder is a function that converts bytes retrieved from etcd into an
// end-user facing value. Additionally, a key is available so that returned
// values can be augmented with the location they were retrieved from. This is
// especially useful when returning values resulting from an etcd range.
//
// If an error is returned, the Get call performed on a watcher configured with
// this decoder will fail, swallowing that particular update, but the watcher
// will continue to work. Any provided BytesDecoder implementations must be safe
// to copy.
type BytesDecoder = func(key []byte, data []byte) (interface{}, error)

// NoDecoder is a no-op decoder which passes through the retrieved bytes as a
// []byte type to the user.
func NoDecoder(key []byte, data []byte) (interface{}, error) {
	return data, nil
}

func (e *Value) Watch() event.Watcher {
	ctx, ctxC := context.WithCancel(context.Background())
	return &watcher{
		Value: *e,

		ctx:  ctx,
		ctxC: ctxC,

		current: make(map[string][]byte),

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

	// backlogged is a list of keys retrieved from etcd but not yet returned via
	// Get. These items are not a replay of all the updates from etcd, but are
	// already compacted to deduplicate updates to the same object (ie., if the
	// update stream from etcd is for keys A, B, and A, the backlogged list will
	// only contain one update for A and B each, with the first update for A being
	// discarded upon arrival of the second update).
	//
	// The keys are an index into the current map, which contains the values
	// retrieved, including ones that have already been returned via Get. This
	// persistence allows us to deduplicate spurious updates to the user, in which
	// etcd returned a new revision of a key, but the data stayed the same.
	backlogged [][]byte
	// current map, keyed from etcd key into etcd value at said key. This map
	// persists alongside an etcd connection, permitting deduplication of spurious
	// etcd updates even across multiple Get calls.
	current map[string][]byte

	// prev is the etcd store revision of a previously completed etcd Get/Watch
	// call, used to resume a Watch call in case of failures.
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

// setup initiates wc (the watch channel from etcd) after retrieving the initial
// value(s) with a get operation.
func (w *watcher) setup(ctx context.Context) error {
	if w.wc != nil {
		return nil
	}
	ranged := w.key != w.keyEnd

	// First, check if some data under this key/range already exists.

	// We use an exponential backoff and retry here as the initial Get can fail
	// if the cluster is unstable (eg. failing over). We only fail the retry if
	// the context expires.
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 0

	err := backoff.Retry(func() error {

		var getOpts []clientv3.OpOption
		if ranged {
			getOpts = append(getOpts, clientv3.WithRange(w.keyEnd))
		}
		get, err := w.etcd.Get(ctx, w.key, getOpts...)
		if err != nil {
			return fmt.Errorf("when retrieving initial value: %w", err)
		}

		// Assert that the etcd API is behaving as expected.
		if !ranged && len(get.Kvs) > 1 {
			panic("More than one key returned in unary GET response")
		}

		// After a successful Get, save the revision to watch from and re-build the
		// backlog from scratch based on what was available in the etcd store at that
		// time.
		w.prev = &get.Header.Revision

		w.backlogged = nil
		w.current = make(map[string][]byte)
		for _, kv := range get.Kvs {
			w.backlogged = append(w.backlogged, kv.Key)
			w.current[string(kv.Key)] = kv.Value
		}
		return nil

	}, backoff.WithContext(bo, ctx))

	if w.testRaceWG != nil {
		w.testRaceWG.Wait()
	}
	if err != nil {
		return err
	}

	watchOpts := []clientv3.OpOption{
		clientv3.WithRev(*w.prev + 1),
	}
	if ranged {
		watchOpts = append(watchOpts, clientv3.WithRange(w.keyEnd))
	}
	w.wc = w.etcd.Watch(w.ctx, w.key, watchOpts...)

	if w.testSetupWG != nil {
		w.testSetupWG.Wait()
	}
	return nil
}

// backfill blocks until a backlog of items is available. An error is returned
// if the context is canceled.
func (w *watcher) backfill(ctx context.Context) error {
	// Keep watching for watch events.
	for {
		var resp *clientv3.WatchResponse
		select {
		case r := <-w.wc:
			resp = &r
		case <-ctx.Done():
			return ctx.Err()
		}

		if resp.Canceled {
			// Only allow for watches to be canceled due to context
			// cancellations. Any other error is something we need to handle,
			// eg. a client close or compaction error.
			if errors.Is(resp.Err(), ctx.Err()) {
				return fmt.Errorf("watch canceled: %w", resp.Err())
			}

			// Attempt to reconnect.
			if w.wc != nil {
				// If a wc already exists, close it. This forces a reconnection
				// by the next setup call.
				w.ctxC()
				w.ctx, w.ctxC = context.WithCancel(context.Background())
				w.wc = nil
			}
			if err := w.setup(ctx); err != nil {
				return fmt.Errorf("failed to setup watcher: %w", err)
			}
			continue
		}

		w.prev = &resp.Header.Revision
		// Spurious watch event with no update? Keep trying.
		if len(resp.Events) == 0 {
			continue
		}

		// Process updates into compacted list, transforming deletions into value: nil
		// keyValues. This maps an etcd key into a pointer in the already existing
		// backlog list. It will then be used to compact all updates into the smallest
		// backlog possible (by overriding previously backlogged items for a key if this
		// key is encountered again).
		//
		// TODO(q3k): this could be stored in the watcher state to not waste time on
		// each update, but it's good enough for now.
		for _, ev := range resp.Events {
			var value []byte
			switch ev.Type {
			case clientv3.EventTypeDelete:
			case clientv3.EventTypePut:
				value = ev.Kv.Value
			default:
				return fmt.Errorf("invalid event type %v", ev.Type)
			}

			keyS := string(ev.Kv.Key)
			prev := w.current[keyS]
			if !bytes.Equal(prev, value) {
				w.backlogged = append(w.backlogged, ev.Kv.Key)
				w.current[keyS] = value
			}
		}

		// Still nothing in backlog? Keep trying.
		if len(w.backlogged) == 0 {
			continue
		}

		return nil
	}
}

type GetOption struct {
	backlogOnly bool
}

var (
	// BacklogOnly will prevent Get from blocking on waiting for more updates from
	// etcd, by instead returning BacklogDone whenever no more data is currently
	// locally available. This is different however, from establishing that there
	// are no more pending updates from the etcd cluster - the only way to ensure
	// the local client is up to date is by performing Get calls without this option
	// set.
	//
	// This mode of retrieval should only be used for the retrieval of the existing
	// data in the etcd cluster on the initial creation of the Watcher (by
	// repeatedly calling Get until BacklogDone isreturned), and shouldn't be set
	// for any subsequent call. Any use of this option after that initial fetch is
	// undefined behaviour that exposes the internals of the Get implementation, and
	// must not be relied on. However, in the future, this behaviour might be
	// formalized.
	//
	// This mode is particularly useful for ranged watchers. Non-ranged watchers can
	// still use this option to distinguish between blocking because of the
	// nonexistence of an object vs. blocking because of networking issues. However,
	// non-ranged retrieval semantics generally will rarely need to make this
	// distinction.
	BacklogOnly = GetOption{backlogOnly: true}

	// BacklogDone is returned by Get when BacklogOnly is set and there is no more
	// event data stored in the Watcher client, ie. when the initial cluster state
	// of the requested key has been retrieved.
	BacklogDone = errors.New("no more backlogged data")
)

// Get implements the Get method of the Watcher interface.
// It can return an error in three cases:
//  - the given context is canceled (in which case, the given error will wrap
//    the context error)
//  - the watcher's BytesDecoder returned an error (in which case the error
//    returned by the BytesDecoder will be returned verbatim)
//  - it has been called with BacklogOnly and the Watcher has no more local
//    event data to return (see BacklogOnly for more information on the
//    semantics of this mode of operation)
// Note that transient and permanent etcd errors are never returned, and the
// Get call will attempt to recover from these errors as much as possible. This
// also means that the user of the Watcher will not be notified if the
// underlying etcd client disconnects from the cluster, or if the cluster loses
// quorum.
//
// TODO(q3k): implement leases to allow clients to be notified when there are
// transient cluster/quorum/partition errors, if needed.
//
// TODO(q3k): implement internal, limited buffering for backlogged data not yet
// consumed by client, as etcd client library seems to use an unbound buffer in
// case this happens ( see: watcherStream.buf in clientv3).
func (w *watcher) Get(ctx context.Context, opts ...event.GetOption) (interface{}, error) {
	select {
	case w.getSem <- struct{}{}:
	default:
		return nil, fmt.Errorf("cannot Get() concurrently on a single waiter")
	}
	defer func() {
		<-w.getSem
	}()

	backlogOnly := false
	for _, optI := range opts {
		opt, ok := optI.(GetOption)
		if !ok {
			return nil, fmt.Errorf("get options must be of type etcd.GetOption")
		}
		if opt.backlogOnly {
			backlogOnly = true
		}
	}

	// Early check for context cancelations, preventing spurious contact with etcd
	// if there's no need to.
	if w.ctx.Err() != nil {
		return nil, w.ctx.Err()
	}

	if err := w.setup(ctx); err != nil {
		return nil, fmt.Errorf("when setting up watcher: %w", err)
	}

	if backlogOnly && len(w.backlogged) == 0 {
		return nil, BacklogDone
	}

	// Update backlog from etcd if needed.
	if len(w.backlogged) < 1 {
		err := w.backfill(ctx)
		if err != nil {
			return nil, fmt.Errorf("when watching for new value: %w", err)
		}
	}
	// Backlog is now guaranteed to contain at least one element.

	ranged := w.key != w.keyEnd
	if !ranged {
		// For non-ranged queries, drain backlog fully.
		if len(w.backlogged) != 1 {
			panic("multiple keys in nonranged value")
		}
		k := w.backlogged[0]
		v := w.current[string(k)]
		w.backlogged = nil
		return w.decoder(k, v)
	} else {
		// For ranged queries, pop one ranged query off the backlog.
		k := w.backlogged[0]
		v := w.current[string(k)]
		w.backlogged = w.backlogged[1:]
		return w.decoder(k, v)
	}
}

func (w *watcher) Close() error {
	w.ctxC()
	return nil
}
