// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"errors"
	"fmt"
	"net"

	"github.com/miekg/dns"
)

// CreateTestRequest creates a Request for use in tests.
func CreateTestRequest(qname string, qtype uint16, proto string) *Request {
	var addr net.Addr
	switch proto {
	case "udp":
		addr = &net.UDPAddr{}
	case "tcp":
		addr = &net.TCPAddr{}
	default:
		panic(fmt.Sprintf("Unknown protocol %q", proto))
	}
	req := &Request{
		Reply:          new(dns.Msg),
		Writer:         &testWriter{addr: addr},
		Qopt:           &dns.OPT{Hdr: dns.RR_Header{Name: ".", Rrtype: dns.TypeOPT}},
		Ropt:           &dns.OPT{Hdr: dns.RR_Header{Name: ".", Rrtype: dns.TypeOPT}},
		Qname:          qname,
		QnameCanonical: dns.CanonicalName(qname),
		Qtype:          qtype,
	}
	req.Reply.Response = true
	req.Reply.Question = []dns.Question{{
		Name:   qname,
		Qtype:  qtype,
		Qclass: dns.ClassINET,
	}}
	req.Reply.RecursionAvailable = true
	req.Reply.RecursionDesired = true
	req.Qopt.SetUDPSize(advertiseUDPSize)
	req.Ropt.SetUDPSize(advertiseUDPSize)
	return req
}

type testWriter struct {
	addr net.Addr
	msg  *dns.Msg
}

func (t *testWriter) LocalAddr() net.Addr  { return t.addr }
func (t *testWriter) RemoteAddr() net.Addr { return t.addr }
func (*testWriter) Write([]byte) (int, error) {
	return 0, errors.New("testWriter only supports WriteMsg")
}
func (*testWriter) Close() error        { return nil }
func (*testWriter) TsigStatus() error   { return nil }
func (*testWriter) TsigTimersOnly(bool) {}
func (*testWriter) Hijack()             {}

func (t *testWriter) WriteMsg(msg *dns.Msg) error {
	if msg == nil {
		panic("WriteMsg(nil)")
	}
	if t.msg != nil {
		panic("duplicate WriteMsg()")
	}
	t.msg = msg
	return nil
}
