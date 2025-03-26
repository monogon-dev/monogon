// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package proxy

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"errors"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/miekg/dns"

	"source.monogon.dev/osbase/net/dns/test"
)

func TestProxy(t *testing.T) {
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Answer = append(ret.Answer, test.RR("example.org. IN A 127.0.0.1"))
		w.WriteMsg(ret)
	})
	defer s.Close()

	p := NewProxy(s.Addr)
	p.Start(5 * time.Second)
	m := new(dns.Msg)

	m.SetQuestion("example.org.", dns.TypeA)

	resp, err := p.Connect(m, false)
	if err != nil {
		t.Errorf("Failed to connect to testdnsserver: %s", err)
	}

	if x := resp.Answer[0].Header().Name; x != "example.org." {
		t.Errorf("Expected %s, got %s", "example.org.", x)
	}
}

func TestProtocolSelection(t *testing.T) {
	p := NewProxy("bad_address")

	go func() {
		p.Connect(new(dns.Msg), false)
		p.Connect(new(dns.Msg), true)
	}()

	for i, exp := range []string{"udp", "tcp"} {
		proto := <-p.transport.dial
		p.transport.ret <- nil
		if proto != exp {
			t.Errorf("Unexpected protocol in case %d, expected %q, actual %q", i, exp, proto)
		}
	}
}

func TestProxyIncrementFails(t *testing.T) {
	var testCases = []struct {
		name        string
		fails       uint32
		expectFails uint32
	}{
		{
			name:        "increment fails counter overflows",
			fails:       math.MaxUint32,
			expectFails: math.MaxUint32,
		},
		{
			name:        "increment fails counter",
			fails:       0,
			expectFails: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewProxy("bad_address")
			p.fails = tc.fails
			p.incrementFails()
			if p.fails != tc.expectFails {
				t.Errorf("Expected fails to be %d, got %d", tc.expectFails, p.fails)
			}
		})
	}
}

func TestCoreDNSOverflow(t *testing.T) {
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)

		var answers []dns.RR
		for i := range 50 {
			answers = append(answers, test.RR(fmt.Sprintf("example.org. IN A 127.0.0.%v", i)))
		}
		ret.Answer = answers
		w.WriteMsg(ret)
	})
	defer s.Close()

	p := NewProxy(s.Addr)
	p.Start(5 * time.Second)
	defer p.Stop()

	// Test different connection modes
	testConnection := func(proto string, useTCP bool, expectTruncated bool) {
		t.Helper()

		queryMsg := new(dns.Msg)
		queryMsg.SetQuestion("example.org.", dns.TypeA)

		response, err := p.Connect(queryMsg, useTCP)
		if err != nil {
			t.Errorf("Failed to connect to testdnsserver: %s", err)
			return
		}

		if response.Truncated != expectTruncated {
			t.Errorf("Expected truncated response for %s, but got TC flag %v", proto, response.Truncated)
		}
	}

	// Test udp, expect truncated response
	testConnection("udp", false, true)

	// Test tcp, expect no truncated response
	testConnection("tcp", true, false)
}

func TestShouldTruncateResponse(t *testing.T) {
	testCases := []struct {
		testname string
		err      error
		expected bool
	}{
		{"BadAlgorithm", dns.ErrAlg, false},
		{"BufferSizeTooSmall", dns.ErrBuf, true},
		{"OverflowUnpackingA", errors.New("overflow unpacking a"), true},
		{"OverflowingHeaderSize", errors.New("overflowing header size"), true},
		{"OverflowpackingA", errors.New("overflow packing a"), true},
		{"ErrSig", dns.ErrSig, false},
	}

	for _, tc := range testCases {
		t.Run(tc.testname, func(t *testing.T) {
			result := shouldTruncateResponse(tc.err)
			if result != tc.expected {
				t.Errorf("For testname '%v', expected %v but got %v", tc.testname, tc.expected, result)
			}
		})
	}
}
