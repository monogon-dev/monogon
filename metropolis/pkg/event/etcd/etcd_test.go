package etcd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.etcd.io/etcd/integration"
	"google.golang.org/grpc/codes"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/pkg/event"
)

var (
	cluster   *integration.ClusterV3
	endpoints []string
)

// TestMain brings up a 3 node etcd cluster for tests to use.
func TestMain(m *testing.M) {
	cfg := integration.ClusterConfig{
		Size:                 3,
		GRPCKeepAliveMinTime: time.Millisecond,
	}
	cluster = integration.NewClusterV3(nil, &cfg)
	endpoints = make([]string, 3)
	for i := range endpoints {
		endpoints[i] = cluster.Client(i).Endpoints()[0]
	}

	v := m.Run()
	cluster.Terminate(nil)
	os.Exit(v)
}

// setRaceWg creates a new WaitGroup and sets the given watcher to wait on this
// WG after it performs the initial retrieval of a value from etcd, but before
// it starts the watcher. This is used to test potential race conditions
// present between these two steps.
func setRaceWg(w event.Watcher) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	w.(*watcher).testRaceWG = &wg
	return &wg
}

// setSetupWg creates a new WaitGroup and sets the given watcher to wait on
// thie WG after an etcd watch channel is created. This is used in tests to
// ensure that the watcher is fully created before it is tested.
func setSetupWg(w event.Watcher) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	w.(*watcher).testSetupWG = &wg
	return &wg
}

// testClient is an etcd connection to the test cluster.
type testClient struct {
	client     *clientv3.Client
	namespaced client.Namespaced
}

func newTestClient(t *testing.T) *testClient {
	t.Helper()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          1 * time.Second,
		DialKeepAliveTime:    1 * time.Second,
		DialKeepAliveTimeout: 1 * time.Second,
	})
	if err != nil {
		t.Fatalf("clientv3.New: %v", err)
	}

	namespaced := client.NewLocal(cli)
	return &testClient{
		client:     cli,
		namespaced: namespaced,
	}
}

func (d *testClient) close() {
	d.client.Close()
}

// setEndpoints configures which endpoints (from {0,1,2}) the testClient is
// connected to.
func (d *testClient) setEndpoints(nums ...uint) {
	var eps []string
	for _, num := range nums {
		eps = append(eps, endpoints[num])
	}
	d.client.SetEndpoints(eps...)
}

// put uses the testClient to store key with a given string value in etcd. It
// contains retry logic that will block until the put is successful.
func (d *testClient) put(t *testing.T, key, value string) {
	t.Helper()
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	for {
		ctxT, ctxC := context.WithTimeout(ctx, 100*time.Millisecond)
		_, err := d.namespaced.Put(ctxT, key, value)
		ctxC()
		if err == nil {
			return
		}
		if err == ctxT.Err() {
			log.Printf("Retrying after %v", err)
			continue
		}
		// Retry on etcd unavailability - this will happen in this code as the
		// etcd cluster repeatedly loses quorum.
		var eerr rpctypes.EtcdError
		if errors.As(err, &eerr) && eerr.Code() == codes.Unavailable {
			log.Printf("Retrying after %v", err)
			continue
		}
		t.Fatalf("Put: %v", err)
	}

}

// remove uses the testClient to remove the given key from etcd. It contains
// retry logic that will block until the removal is successful.
func (d *testClient) remove(t *testing.T, key string) {
	t.Helper()
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	_, err := d.namespaced.Delete(ctx, key)
	if err == nil {
		return
	}
	t.Fatalf("Delete: %v", err)
}

// expect runs a Get on the given Watcher, ensuring the returned value is a
// given string.
func expect(t *testing.T, w event.Watcher, value string) {
	t.Helper()
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	got, err := w.Get(ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	if got, want := string(got.([]byte)), value; got != want {
		t.Errorf("Got value %q, wanted %q", want, got)
	}
}

// expectTimeout ensures that the given watcher blocks on a Get call for at
// least 100 milliseconds. This is used by tests to attempt to verify that the
// watcher Get is fully blocked, but can cause false positives (eg. when Get
// blocks for 101 milliseconds). Thus, this function should be used sparingly
// and in tests that perform other baseline behaviour checks alongside this
// test.
func expectTimeout(t *testing.T, w event.Watcher) {
	t.Helper()
	ctx, ctxC := context.WithTimeout(context.Background(), 100*time.Millisecond)
	got, err := w.Get(ctx)
	ctxC()

	if !errors.Is(err, ctx.Err()) {
		t.Fatalf("Expected timeout error, got %v, %v", got, err)
	}
}

// wait wraps a watcher into a channel of strings, ensuring that the watcher
// never errors on Get calls and always returns strings.
func wait(t *testing.T, w event.Watcher) (chan string, func()) {
	t.Helper()
	ctx, ctxC := context.WithCancel(context.Background())

	c := make(chan string)

	go func() {
		for {
			got, err := w.Get(ctx)
			if err != nil && errors.Is(err, ctx.Err()) {
				return
			}
			if err != nil {
				t.Fatalf("Get: %v", err)
			}
			c <- string(got.([]byte))
		}
	}()

	return c, ctxC
}

// TestSimple exercises the simplest possible interaction with a watched value.
func TestSimple(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()

	k := "test-simple"
	value := NewValue(tc.namespaced, k, NoDecoder)
	tc.put(t, k, "one")

	watcher := value.Watch()
	defer watcher.Close()
	expect(t, watcher, "one")

	tc.put(t, k, "two")
	expect(t, watcher, "two")

	tc.put(t, k, "three")
	tc.put(t, k, "four")
	tc.put(t, k, "five")
	tc.put(t, k, "six")

	q, cancel := wait(t, watcher)
	// Test will hang here if the above value does not receive the set "six".
	for el := range q {
		if el == "six" {
			break
		}
	}
	cancel()
}

// TestCancel ensures that watchers can resume after being canceled.
func TestCancel(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()

	k := "test-cancel"
	value := NewValue(tc.namespaced, k, NoDecoder)
	tc.put(t, k, "one")

	watcher := value.Watch()
	defer watcher.Close()
	expect(t, watcher, "one")

	ctx, ctxC := context.WithCancel(context.Background())
	errs := make(chan error, 1)
	go func() {
		_, err := watcher.Get(ctx)
		errs <- err
	}()
	ctxC()
	if want, got := ctx.Err(), <-errs; !errors.Is(got, want) {
		t.Fatalf("Wanted err %v, got %v", want, got)
	}

	// Successfully canceled watch, resuming should continue to work.
	q, cancel := wait(t, watcher)
	defer cancel()

	tc.put(t, k, "two")
	if want, got := "two", <-q; want != got {
		t.Fatalf("Wanted val %q, got %q", want, got)
	}
}

// TestCancelOnGet ensures that a context cancellation on an initial Get (which
// translates to an etcd Get in a backoff loop) doesn't block.
func TestCancelOnGet(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()

	k := "test-cancel-on-get"
	value := NewValue(tc.namespaced, k, NoDecoder)
	watcher := value.Watch()
	tc.put(t, k, "one")

	// Cause partition between client endpoint and rest of cluster. Any read/write
	// operations will now hang.
	tc.setEndpoints(0)
	cluster.Members[0].InjectPartition(t, cluster.Members[1], cluster.Members[2])

	// Perform the initial Get(), which should attempt to retrieve a KV entry from
	// the etcd service. This should hang. Unfortunately, there's no easy way to do
	// this without an arbitrary sleep hoping that the client actually gets to the
	// underlying etcd.Get call. This can cause false positives (eg. false 'pass'
	// results) in this test.
	ctx, ctxC := context.WithCancel(context.Background())
	errs := make(chan error, 1)
	go func() {
		_, err := watcher.Get(ctx)
		errs <- err
	}()
	time.Sleep(time.Second)

	// Now that the etcd.Get is hanging, cancel the context.
	ctxC()
	// And now unpartition the cluster, resuming reads.
	cluster.Members[0].RecoverPartition(t, cluster.Members[1], cluster.Members[2])

	// The etcd.Get() call should've returned with a context cancellation.
	err := <-errs
	switch {
	case err == nil:
		t.Errorf("watcher.Get() returned no error, wanted context error")
	case errors.Is(err, ctx.Err()):
		// Okay.
	default:
		t.Errorf("watcher.Get() returned %v, wanted context error", err)
	}
}

// TestClientReconnect forces a 'reconnection' of an active watcher from a
// running member to another member, by stopping the original member and
// explicitly reconnecting the client to other available members.
//
// This doe not reflect a situation expected during Metropolis runtime, as we
// do not expect splits between an etcd client and its connected member
// (instead, all etcd clients only connect to their local member). However, it
// is still an important safety test to perform, and it also exercies the
// equivalent behaviour of an etcd client re-connecting for any other reason.
func TestClientReconnect(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()
	tc.setEndpoints(0)

	k := "test-client-reconnect"
	value := NewValue(tc.namespaced, k, NoDecoder)
	tc.put(t, k, "one")

	watcher := value.Watch()
	defer watcher.Close()
	expect(t, watcher, "one")

	q, cancel := wait(t, watcher)
	defer cancel()

	cluster.Members[0].Stop(t)
	defer cluster.Members[0].Restart(t)
	cluster.WaitLeader(t)

	tc.setEndpoints(1, 2)
	tc.put(t, k, "two")

	if want, got := "two", <-q; want != got {
		t.Fatalf("Watcher received incorrect data after client restart, wanted %q, got %q", want, got)
	}
}

// TestClientPartition forces a temporary partition of the etcd member while a
// watcher is running, updates the value from across the partition, and undoes
// the partition.
// The partition is expected to be entirely transparent to the watcher.
func TestClientPartition(t *testing.T) {
	tcOne := newTestClient(t)
	defer tcOne.close()
	tcOne.setEndpoints(0)

	tcRest := newTestClient(t)
	defer tcRest.close()
	tcRest.setEndpoints(1, 2)

	k := "test-client-partition"
	valueOne := NewValue(tcOne.namespaced, k, NoDecoder)
	watcherOne := valueOne.Watch()
	defer watcherOne.Close()
	valueRest := NewValue(tcRest.namespaced, k, NoDecoder)
	watcherRest := valueRest.Watch()
	defer watcherRest.Close()

	tcRest.put(t, k, "a")
	expect(t, watcherOne, "a")
	expect(t, watcherRest, "a")

	cluster.Members[0].InjectPartition(t, cluster.Members[1], cluster.Members[2])

	tcRest.put(t, k, "b")
	expect(t, watcherRest, "b")
	expectTimeout(t, watcherOne)

	cluster.Members[0].RecoverPartition(t, cluster.Members[1], cluster.Members[2])

	expect(t, watcherOne, "b")
	tcRest.put(t, k, "c")
	expect(t, watcherOne, "c")
	expect(t, watcherRest, "c")

}

// TestEarlyUse exercises the correct behaviour of the value watcher on a value
// that is not yet set.
func TestEarlyUse(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()

	k := "test-early-use"

	value := NewValue(tc.namespaced, k, NoDecoder)
	watcher := value.Watch()
	defer watcher.Close()

	wg := setSetupWg(watcher)
	wg.Add(1)
	q, cancel := wait(t, watcher)
	defer cancel()

	wg.Done()

	tc.put(t, k, "one")

	if want, got := "one", <-q; want != got {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

// TestRemove exercises the basic functionality of handling deleted values.
func TestRemove(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()

	k := "test-remove"
	tc.put(t, k, "one")

	value := NewValue(tc.namespaced, k, NoDecoder)
	watcher := value.Watch()
	defer watcher.Close()

	expect(t, watcher, "one")
	tc.remove(t, k)
	expect(t, watcher, "")
}

// TestEmptyRace forces the watcher to retrieve an empty value from the K/V
// store at first, and establishing the watch channel after a new value has
// been stored in the same place.
func TestEmptyRace(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()

	k := "test-remove-race"
	tc.put(t, k, "one")
	tc.remove(t, k)

	value := NewValue(tc.namespaced, k, NoDecoder)
	watcher := value.Watch()
	defer watcher.Close()

	wg := setRaceWg(watcher)
	wg.Add(1)
	q, cancel := wait(t, watcher)
	defer cancel()

	tc.put(t, k, "two")
	wg.Done()

	if want, got := "two", <-q; want != got {
		t.Fatalf("Watcher received incorrect data after client restart, wanted %q, got %q", want, got)
	}
}

type errOrInt struct {
	val int64
	err error
}

// TestDecoder exercises the BytesDecoder functionality of the watcher, by
// creating a value with a decoder that only accepts string-encoded integers
// that are divisible by three. The test then proceeds to put a handful of
// values into etcd, ensuring that undecodable values correctly return an error
// on Get, but that the watcher continues to work after the error has been
// returned.
func TestDecoder(t *testing.T) {
	decodeStringifiedNumbersDivisibleBy3 := func(data []byte) (interface{}, error) {
		num, err := strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("not a valid number")
		}
		if (num % 3) != 0 {
			return nil, fmt.Errorf("not divisible by 3")
		}
		return num, nil
	}

	tc := newTestClient(t)
	defer tc.close()

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	k := "test-decoder"
	value := NewValue(tc.namespaced, k, decodeStringifiedNumbersDivisibleBy3)
	watcher := value.Watch()
	defer watcher.Close()
	tc.put(t, k, "3")
	_, err := watcher.Get(ctx)
	if err != nil {
		t.Fatalf("Initial Get: %v", err)
	}

	// Stream updates into arbitrarily-bounded test channel.
	queue := make(chan errOrInt, 100)
	go func() {
		for {
			res, err := watcher.Get(ctx)
			if err != nil && errors.Is(err, ctx.Err()) {
				return
			}
			if err != nil {
				queue <- errOrInt{
					err: err,
				}
			} else {
				queue <- errOrInt{
					val: res.(int64),
				}
			}
		}
	}()

	var wantList []*int64
	wantError := func(val string) {
		wantList = append(wantList, nil)
		tc.put(t, k, val)
	}
	wantValue := func(val string, decoded int64) {
		wantList = append(wantList, &decoded)
		tc.put(t, k, val)
	}

	wantError("")
	wantValue("9", 9)
	wantError("foo")
	wantValue("18", 18)
	wantError("10")
	wantError("10")
	wantValue("27", 27)
	wantValue("36", 36)

	for i, want := range wantList {
		q := <-queue
		if want == nil && q.err == nil {
			t.Errorf("Case %d: wanted error, got no error and value %d", i, q.val)
		}
		if want != nil && (*want) != q.val {
			t.Errorf("Case %d: wanted value %d, got error %v and value %d", i, *want, q.err, q.val)
		}
	}
}

// TestBacklog ensures that the watcher can handle a large backlog of changes
// in etcd that the client didnt' keep up with, and that whatever final state
// is available to the client when it actually gets to calling Get().
func TestBacklog(t *testing.T) {
	tc := newTestClient(t)
	defer tc.close()

	k := "test-backlog"
	value := NewValue(tc.namespaced, k, NoDecoder)
	watcher := value.Watch()
	defer watcher.Close()

	tc.put(t, k, "initial")
	expect(t, watcher, "initial")

	for i := 0; i < 1000; i++ {
		tc.put(t, k, fmt.Sprintf("val-%d", i))
	}

	ctx, ctxC := context.WithTimeout(context.Background(), time.Second)
	defer ctxC()
	for {
		valB, err := watcher.Get(ctx)
		if err != nil {
			t.Fatalf("Get() returned error before expected final value: %v", err)
		}
		val := string(valB.([]byte))
		if val == "val-999" {
			break
		}
	}
}
