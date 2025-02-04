// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package consensus

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/test/util"
	"source.monogon.dev/osbase/supervisor"
)

type boilerplate struct {
	ctx     context.Context
	ctxC    context.CancelFunc
	root    *localstorage.Root
	privkey ed25519.PrivateKey
	tmpdir  string
}

func prep(t *testing.T) *boilerplate {
	t.Helper()
	ctx, ctxC := context.WithCancel(context.Background())
	root := &localstorage.Root{}
	// Force usage of /tmp as temp directory root, otherwsie TMPDIR from Bazel
	// returns a path long enough that socket binds in the localstorage fail
	// (as socket names are limited to 108 characters).
	tmp, err := os.MkdirTemp("/tmp", "metropolis-consensus-test")
	if err != nil {
		t.Fatal(err)
	}
	err = declarative.PlaceFS(root, tmp)
	if err != nil {
		t.Fatal(err)
	}
	os.MkdirAll(root.Data.Etcd.FullPath(), 0700)
	os.MkdirAll(root.Ephemeral.Consensus.FullPath(), 0700)

	_, privkey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	return &boilerplate{
		ctx:     ctx,
		ctxC:    ctxC,
		root:    root,
		privkey: privkey,
		tmpdir:  tmp,
	}
}

func (b *boilerplate) close() {
	b.ctxC()
	os.RemoveAll(b.tmpdir)
}

func TestEtcdMetrics(t *testing.T) {
	b := prep(t)
	defer b.close()
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		NodeID:         "node1",
		NodePrivateKey: b.privkey,
		testOverrides: testOverrides{
			externalPort:    2345,
			etcdMetricsPort: 4100,
		},
	})

	ctxC, _ := supervisor.TestHarness(t, etcd.Run)
	defer ctxC()

	w := etcd.Watch()
	// Wait until etcd state is populated as the metrics endpoint comes up
	// before the bootstrap process is complete. The test context then gets
	// cancelled as the test has succeeded, but the consensus runnable is not
	// done bootstrapping, causing it to return an error and TestHarness to fail
	// the test.
	_, err := w.Get(b.ctx)
	if err != nil {
		t.Fatalf("status get failed: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	util.TestEventual(t, "metrics-reachable", ctx, 10*time.Second, func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:4100/metrics", nil)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("Get: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("StatusCode: wanted 200, got %d", resp.StatusCode)
		}
		return nil
	})
}

func TestBootstrap(t *testing.T) {
	b := prep(t)
	defer b.close()
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		NodeID:         "node1",
		NodePrivateKey: b.privkey,
		testOverrides: testOverrides{
			externalPort: 1234,
		},
	})

	ctxC, _ := supervisor.TestHarness(t, etcd.Run)
	defer ctxC()

	w := etcd.Watch()
	st, err := w.Get(b.ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	cl, err := st.CuratorClient()
	if err != nil {
		t.Fatalf("CuratorClient: %v", err)
	}
	defer cl.Close()

	if _, err := cl.Put(b.ctx, "/foo", "bar"); err != nil {
		t.Fatalf("test key creation failed: %v", err)
	}
	if _, err := cl.Get(b.ctx, "/foo"); err != nil {
		t.Fatalf("test key retrieval failed: %v", err)
	}
}

func TestRestartFromDisk(t *testing.T) {
	b := prep(t)
	defer b.close()

	// Start once.
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		NodeID:         "node1",
		NodePrivateKey: b.privkey,
		testOverrides: testOverrides{
			externalPort: 1235,
		},
	})
	ctxC, _ := supervisor.TestHarness(t, etcd.Run)
	defer ctxC()

	w := etcd.Watch()
	st, err := w.Get(b.ctx)
	if err != nil {
		t.Fatalf("status get failed: %v", err)
	}
	cl, err := st.CuratorClient()
	if err != nil {
		t.Fatalf("CuratorClient: %v", err)
	}
	defer cl.Close()

	if _, err := cl.Put(b.ctx, "/foo", "bar"); err != nil {
		t.Fatalf("test key creation failed: %v", err)
	}
	firstCA, err := etcd.config.Data.PeerPKI.CACertificate.Read()
	if err != nil {
		t.Fatalf("could not read CA file: %v", err)
	}

	// Stop and wait until reported stopped.
	ctxC()
	ctxWait, ctxWaitC := context.WithCancel(context.Background())
	for {
		if st.stopped {
			break
		}
		st, err = w.Get(ctxWait)
		if err != nil {
			t.Fatalf("status get failed: %v", err)
		}
		if st.stopped {
			break
		}
	}
	ctxWaitC()

	// Restart.
	etcd = New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		NodeID:         "node1",
		NodePrivateKey: b.privkey,
		testOverrides: testOverrides{
			externalPort: 1235,
		},
	})
	ctxC, _ = supervisor.TestHarness(t, etcd.Run)
	defer ctxC()

	w = etcd.Watch()
	st, err = w.Get(b.ctx)
	if err != nil {
		t.Fatalf("status get failed: %v", err)
	}
	cl, err = st.CuratorClient()
	if err != nil {
		t.Fatalf("CuratorClient: %v", err)
	}
	defer cl.Close()

	res, err := cl.Get(b.ctx, "/foo")
	if err != nil {
		t.Fatalf("test key retrieval failed: %v", err)
	}
	if len(res.Kvs) != 1 || string(res.Kvs[0].Value) != "bar" {
		t.Fatalf("test key value missing: %v", res.Kvs)
	}

	secondCA, err := etcd.config.Data.PeerPKI.CACertificate.Read()
	if err != nil {
		t.Fatalf("could not read CA file: %v", err)
	}
	ctxC()

	if !bytes.Equal(firstCA, secondCA) {
		t.Fatalf("wanted same, got different CAs accross runs")
	}
}

func TestJoin(t *testing.T) {
	b := prep(t)
	defer b.close()

	// Start first node and perform write.
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		NodeID:         "node1",
		NodePrivateKey: b.privkey,
		testOverrides: testOverrides{
			externalPort:    3000,
			externalAddress: "localhost",
			etcdMetricsPort: 3100,
		},
	})
	ctxC, _ := supervisor.TestHarness(t, etcd.Run)
	defer ctxC()

	w := etcd.Watch()
	st, err := w.Get(b.ctx)
	if err != nil {
		t.Fatalf("could not get status: %v", err)
	}
	cl, err := st.CuratorClient()
	if err != nil {
		t.Fatalf("CuratorClient: %v", err)
	}
	defer cl.Close()
	if _, err := cl.Put(b.ctx, "/foo", "bar"); err != nil {
		t.Fatalf("test key creation failed: %v", err)
	}

	// Start second node and ensure data is present.
	b2 := prep(t)
	defer b2.close()

	join, err := st.AddNode(b.ctx, "node2", b2.privkey.Public().(ed25519.PublicKey), &AddNodeOption{
		externalAddress: "localhost",
		externalPort:    3001,
	})
	if err != nil {
		t.Fatalf("could not add node: %v", err)
	}

	etcd2 := New(Config{
		Data:           &b2.root.Data.Etcd,
		Ephemeral:      &b2.root.Ephemeral.Consensus,
		NodeID:         "node2",
		NodePrivateKey: b2.privkey,
		JoinCluster:    join,
		testOverrides: testOverrides{
			externalPort:    3001,
			externalAddress: "localhost",
			etcdMetricsPort: 3101,
		},
	})
	ctxC, _ = supervisor.TestHarness(t, etcd2.Run)
	defer ctxC()

	w2 := etcd2.Watch()
	st2, err := w2.Get(b.ctx)
	if err != nil {
		t.Fatalf("could not get status: %v", err)
	}
	cl2, err := st2.CuratorClient()
	if err != nil {
		t.Fatalf("CuratorClient: %v", err)
	}
	defer cl2.Close()

	res, err := cl2.Get(b.ctx, "/foo")
	if err != nil {
		t.Fatalf("test key retrieval failed: %v", err)
	}
	if len(res.Kvs) != 1 || string(res.Kvs[0].Value) != "bar" {
		t.Fatalf("test key value missing: %v", res.Kvs)
	}
}
