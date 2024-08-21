package proxy

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"testing"
	"time"

	"github.com/miekg/dns"

	"source.monogon.dev/osbase/net/dns/test"
)

func TestCached(t *testing.T) {
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		w.WriteMsg(ret)
	})
	defer s.Close()

	tr := newTransport(s.Addr)
	tr.Start()
	defer tr.Stop()

	c1, cache1, _ := tr.Dial("udp")
	c2, cache2, _ := tr.Dial("udp")

	if cache1 || cache2 {
		t.Errorf("Expected non-cached connection")
	}

	tr.Yield(c1)
	tr.Yield(c2)
	c3, cached3, _ := tr.Dial("udp")
	if !cached3 {
		t.Error("Expected cached connection (c3)")
	}
	if c2 != c3 {
		t.Error("Expected c2 == c3")
	}

	tr.Yield(c3)

	// dial another protocol
	c4, cached4, _ := tr.Dial("tcp")
	if cached4 {
		t.Errorf("Expected non-cached connection (c4)")
	}
	tr.Yield(c4)
}

func TestCleanupByTimer(t *testing.T) {
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		w.WriteMsg(ret)
	})
	defer s.Close()

	tr := newTransport(s.Addr)
	tr.SetExpire(10 * time.Millisecond)
	tr.Start()
	defer tr.Stop()

	c1, _, _ := tr.Dial("udp")
	c2, _, _ := tr.Dial("udp")
	tr.Yield(c1)
	time.Sleep(2 * time.Millisecond)
	tr.Yield(c2)

	time.Sleep(15 * time.Millisecond)
	c3, cached, _ := tr.Dial("udp")
	if cached {
		t.Error("Expected non-cached connection (c3)")
	}
	tr.Yield(c3)

	time.Sleep(15 * time.Millisecond)
	c4, cached, _ := tr.Dial("udp")
	if cached {
		t.Error("Expected non-cached connection (c4)")
	}
	tr.Yield(c4)
}

func TestCleanupAll(t *testing.T) {
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		w.WriteMsg(ret)
	})
	defer s.Close()

	tr := newTransport(s.Addr)

	c1, _ := dns.DialTimeout("udp", tr.addr, maxDialTimeout)
	c2, _ := dns.DialTimeout("udp", tr.addr, maxDialTimeout)
	c3, _ := dns.DialTimeout("udp", tr.addr, maxDialTimeout)

	tr.conns[typeUDP] = []*persistConn{{c1, time.Now()}, {c2, time.Now()}, {c3, time.Now()}}

	if len(tr.conns[typeUDP]) != 3 {
		t.Error("Expected 3 connections")
	}
	tr.cleanup(true)

	if len(tr.conns[typeUDP]) > 0 {
		t.Error("Expected no cached connections")
	}
}
