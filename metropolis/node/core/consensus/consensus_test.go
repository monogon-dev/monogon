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
	"crypto/x509"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"git.monogon.dev/source/nexantic.git/metropolis/node/core/localstorage"
	"git.monogon.dev/source/nexantic.git/metropolis/node/core/localstorage/declarative"
	"git.monogon.dev/source/nexantic.git/metropolis/pkg/freeport"
	"git.monogon.dev/source/nexantic.git/metropolis/pkg/supervisor"
)

type boilerplate struct {
	ctx    context.Context
	ctxC   context.CancelFunc
	root   *localstorage.Root
	tmpdir string
}

func prep(t *testing.T) *boilerplate {
	ctx, ctxC := context.WithCancel(context.Background())
	root := &localstorage.Root{}
	tmp, err := ioutil.TempDir("", "metropolis-consensus-test")
	if err != nil {
		t.Fatal(err)
	}
	err = declarative.PlaceFS(root, tmp)
	if err != nil {
		t.Fatal(err)
	}
	os.MkdirAll(root.Data.Etcd.FullPath(), 0700)
	os.MkdirAll(root.Ephemeral.Consensus.FullPath(), 0700)

	return &boilerplate{
		ctx:    ctx,
		ctxC:   ctxC,
		root:   root,
		tmpdir: tmp,
	}
}

func (b *boilerplate) close() {
	b.ctxC()
	os.RemoveAll(b.tmpdir)
}

func waitEtcd(t *testing.T, s *Service) {
	deadline := time.Now().Add(5 * time.Second)
	for {
		if time.Now().After(deadline) {
			t.Fatalf("etcd did not start up on time")
		}
		if s.IsReady() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func TestBootstrap(t *testing.T) {
	b := prep(t)
	defer b.close()
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		Name:           "test",
		NewCluster:     true,
		InitialCluster: "127.0.0.1",
		ExternalHost:   "127.0.0.1",
		ListenHost:     "127.0.0.1",
		Port:           freeport.MustConsume(freeport.AllocateTCPPort()),
	})

	supervisor.New(b.ctx, etcd.Run)
	waitEtcd(t, etcd)

	kv := etcd.KV("foo", "bar")
	if _, err := kv.Put(b.ctx, "/foo", "bar"); err != nil {
		t.Fatalf("test key creation failed: %v", err)
	}
	if _, err := kv.Get(b.ctx, "/foo"); err != nil {
		t.Fatalf("test key retrieval failed: %v", err)
	}
}

func TestMemberInfo(t *testing.T) {
	b := prep(t)
	defer b.close()
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		Name:           "test",
		NewCluster:     true,
		InitialCluster: "127.0.0.1",
		ExternalHost:   "127.0.0.1",
		ListenHost:     "127.0.0.1",
		Port:           freeport.MustConsume(freeport.AllocateTCPPort()),
	})
	supervisor.New(b.ctx, etcd.Run)
	waitEtcd(t, etcd)

	id, name, err := etcd.MemberInfo(b.ctx)
	if err != nil {
		t.Fatalf("MemberInfo: %v", err)
	}

	// Compare name with configured name.
	if want, got := "test", name; want != got {
		t.Errorf("MemberInfo returned name %q, wanted %q (per config)", got, want)
	}

	// Compare name with cluster information.
	members, err := etcd.Cluster().MemberList(b.ctx)
	if err != nil {
		t.Errorf("MemberList: %v", err)
	}

	if want, got := 1, len(members.Members); want != got {
		t.Fatalf("expected one cluster member, got %d", got)
	}
	if want, got := id, members.Members[0].ID; want != got {
		t.Errorf("MemberInfo returned ID %d, Cluster endpoint says %d", want, got)
	}
	if want, got := name, members.Members[0].Name; want != got {
		t.Errorf("MemberInfo returned name %q, Cluster endpoint says %q", want, got)
	}
}

func TestRestartFromDisk(t *testing.T) {
	b := prep(t)
	defer b.close()

	startEtcd := func(new bool) (*Service, context.CancelFunc) {
		etcd := New(Config{
			Data:           &b.root.Data.Etcd,
			Ephemeral:      &b.root.Ephemeral.Consensus,
			Name:           "test",
			NewCluster:     new,
			InitialCluster: "127.0.0.1",
			ExternalHost:   "127.0.0.1",
			ListenHost:     "127.0.0.1",
			Port:           freeport.MustConsume(freeport.AllocateTCPPort()),
		})
		ctx, ctxC := context.WithCancel(b.ctx)
		supervisor.New(ctx, etcd.Run)
		waitEtcd(t, etcd)
		kv := etcd.KV("foo", "bar")
		if new {
			if _, err := kv.Put(b.ctx, "/foo", "bar"); err != nil {
				t.Fatalf("test key creation failed: %v", err)
			}
		}
		if _, err := kv.Get(b.ctx, "/foo"); err != nil {
			t.Fatalf("test key retrieval failed: %v", err)
		}

		return etcd, ctxC
	}

	etcd, ctxC := startEtcd(true)
	etcd.stateMu.Lock()
	firstCA := etcd.state.ca.CACertRaw
	etcd.stateMu.Unlock()
	ctxC()

	etcd, ctxC = startEtcd(false)
	etcd.stateMu.Lock()
	secondCA := etcd.state.ca.CACertRaw
	etcd.stateMu.Unlock()
	ctxC()

	if bytes.Compare(firstCA, secondCA) != 0 {
		t.Fatalf("wanted same, got different CAs accross runs")
	}
}

func TestCRL(t *testing.T) {
	b := prep(t)
	defer b.close()
	etcd := New(Config{
		Data:           &b.root.Data.Etcd,
		Ephemeral:      &b.root.Ephemeral.Consensus,
		Name:           "test",
		NewCluster:     true,
		InitialCluster: "127.0.0.1",
		ExternalHost:   "127.0.0.1",
		ListenHost:     "127.0.0.1",
		Port:           freeport.MustConsume(freeport.AllocateTCPPort()),
	})
	supervisor.New(b.ctx, etcd.Run)
	waitEtcd(t, etcd)

	etcd.stateMu.Lock()
	ca := etcd.state.ca
	kv := etcd.state.cl.KV
	etcd.stateMu.Unlock()

	certRaw, _, err := ca.Issue(b.ctx, kv, "revoketest", net.ParseIP("1.2.3.4"))
	if err != nil {
		t.Fatalf("cert issue failed: %v", err)
	}
	cert, err := x509.ParseCertificate(certRaw)
	if err != nil {
		t.Fatalf("cert parse failed: %v", err)
	}

	if err := ca.Revoke(b.ctx, kv, "revoketest"); err != nil {
		t.Fatalf("cert revoke failed: %v", err)
	}

	deadline := time.Now().Add(5 * time.Second)
	for {
		if time.Now().After(deadline) {
			t.Fatalf("CRL did not get updated in time")
		}
		time.Sleep(100 * time.Millisecond)

		crlRaw, err := b.root.Data.Etcd.PeerCRL.Read()
		if err != nil {
			// That's fine. Maybe it hasn't been written yet.
			continue
		}
		crl, err := x509.ParseCRL(crlRaw)
		if err != nil {
			// That's fine. Maybe it hasn't been written yet.
			continue
		}

		found := false
		for _, revoked := range crl.TBSCertList.RevokedCertificates {
			if revoked.SerialNumber.Cmp(cert.SerialNumber) == 0 {
				found = true
			}
		}
		if found {
			break
		}
	}
}