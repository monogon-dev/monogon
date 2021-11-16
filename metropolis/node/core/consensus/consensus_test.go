// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consensus

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"os"
	"testing"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/pkg/supervisor"
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

func TestBootstrap(t *testing.T) {
	b := prep(t)
	defer b.close()
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
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

	if bytes.Compare(firstCA, secondCA) != 0 {
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
		NodePrivateKey: b.privkey,
		testOverrides: testOverrides{
			externalPort:    3000,
			externalAddress: "localhost",
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

	join, err := st.AddNode(b.ctx, b2.privkey.Public().(ed25519.PublicKey), &AddNodeOption{
		externalAddress: "localhost",
		externalPort:    3001,
	})
	if err != nil {
		t.Fatalf("could not add node: %v", err)
	}

	etcd2 := New(Config{
		Data:           &b2.root.Data.Etcd,
		Ephemeral:      &b2.root.Ephemeral.Consensus,
		NodePrivateKey: b2.privkey,
		JoinCluster:    join,
		testOverrides: testOverrides{
			externalPort:    3001,
			externalAddress: "localhost",
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
