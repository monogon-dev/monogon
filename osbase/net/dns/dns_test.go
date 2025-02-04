// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"net"
	"testing"

	"github.com/miekg/dns"

	"source.monogon.dev/osbase/net/dns/test"
)

func testQuery(t *testing.T, service *Service, query, wantReply *dns.Msg) {
	t.Helper()
	wantReply.RecursionAvailable = true
	testMsg(t, service, query, wantReply)
}

func testMsg(t *testing.T, service *Service, query, wantReply *dns.Msg) {
	sCtx := &serviceCtx{service: service}
	t.Helper()
	wantReply.Response = true
	writer := &testWriter{addr: &net.UDPAddr{}}
	sCtx.ServeDNS(writer, query)
	if got, want := writer.msg.String(), wantReply.String(); got != want {
		t.Errorf("Want reply:\n%s\nGot:\n%s", want, got)
	}
}

func TestBuiltinHandlers(t *testing.T) {
	service := New(nil)

	cases := []struct {
		name   string
		qtype  uint16
		rcode  int
		answer []dns.RR
	}{
		{
			name:   "localhost.",
			qtype:  dns.TypeA,
			answer: []dns.RR{test.RR("localhost.	300	IN	A	127.0.0.1")},
		},
		{
			name:   "foo.bar.localhost.",
			qtype:  dns.TypeA,
			answer: []dns.RR{test.RR("foo.bar.localhost.	300	IN	A	127.0.0.1")},
		},
		{
			name:   "localhost.",
			qtype:  dns.TypeAAAA,
			answer: []dns.RR{test.RR("localhost.	300	IN	AAAA	::1")},
		},
		{
			name:  "localhost.",
			qtype: dns.TypeANY,
			answer: []dns.RR{
				test.RR("localhost.	300	IN	A	127.0.0.1"),
				test.RR("localhost.	300	IN	AAAA	::1"),
			},
		},
		{
			name:  "localhost.",
			qtype: dns.TypeMX,
		},
		{
			name:   "1.0.0.127.in-addr.arpa.",
			qtype:  dns.TypePTR,
			answer: []dns.RR{test.RR("1.0.0.127.in-addr.arpa.	300	IN	PTR	localhost.")},
		},
		{
			name:  "1.0.0.127.in-addr.arpa.",
			qtype: dns.TypeNS,
		},
		{
			name:  "2.127.in-addr.arpa.",
			qtype: dns.TypePTR,
		},
		{
			name:  "foo.127.in-addr.arpa.",
			qtype: dns.TypePTR,
			rcode: dns.RcodeNameError,
		},
		{
			name:   "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.",
			qtype:  dns.TypePTR,
			answer: []dns.RR{test.RR("1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.	300	IN	PTR	localhost.")},
		},
		{
			name:  "foo.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.",
			qtype: dns.TypePTR,
			rcode: dns.RcodeNameError,
		},
		{
			name:  "invalid.",
			qtype: dns.TypeA,
			rcode: dns.RcodeNameError,
		},
		{
			name:  "foo.bar.invalid.",
			qtype: dns.TypeA,
			rcode: dns.RcodeNameError,
		},
	}

	for _, c := range cases {
		query := new(dns.Msg)
		query.SetQuestion(c.name, c.qtype)
		query.RecursionDesired = false
		wantReply := query.Copy()
		wantReply.Authoritative = true
		wantReply.Rcode = c.rcode
		wantReply.Answer = c.answer
		testQuery(t, service, query, wantReply)
	}
}

type handlerFunc func(*Request)

func (f handlerFunc) HandleDNS(r *Request) {
	f(r)
}

func TestCustomHandlers(t *testing.T) {
	service := New([]string{"handler1", "handler2"})
	service.SetHandler("handler2", handlerFunc(func(r *Request) {
		if IsSubDomain("example.com.", r.QnameCanonical) {
			r.SetAuthoritative()
			if r.Qtype == dns.TypeA || r.Qtype == dns.TypeANY {
				rr := new(dns.A)
				rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300}
				rr.A = net.IP{1, 2, 3, 4}
				r.Reply.Answer = append(r.Reply.Answer, rr)
			}
			r.SendReply()
		}
	}))

	// Because handler1 is not yet set, this query should fail.
	query := new(dns.Msg)
	query.SetQuestion("example.com.", dns.TypeA)
	wantReply := query.Copy()
	wantReply.Rcode = dns.RcodeServerFailure
	testQuery(t, service, query, wantReply)

	service.SetHandler("handler1", EmptyDNSHandler{})

	// Now, we should get the result from handler2.
	query = new(dns.Msg)
	query.SetQuestion("example.com.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Authoritative = true
	wantReply.Answer = []dns.RR{test.RR("example.com.	300	IN	A	1.2.3.4")}
	testQuery(t, service, query, wantReply)

	service.SetHandler("handler1", handlerFunc(func(r *Request) {
		if IsSubDomain("example.com.", r.QnameCanonical) {
			r.SetAuthoritative()
			if r.Qtype == dns.TypeA || r.Qtype == dns.TypeANY {
				rr := new(dns.A)
				rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300}
				rr.A = net.IP{5, 6, 7, 8}
				r.Reply.Answer = append(r.Reply.Answer, rr)
			}
			r.SendReply()
		}
	}))

	// Handlers can be updated, and are tried in the order in which they were
	// declared when creating the Service.
	query = new(dns.Msg)
	query.SetQuestion("example.com.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Authoritative = true
	wantReply.Answer = []dns.RR{test.RR("example.com.	300	IN	A	5.6.7.8")}
	testQuery(t, service, query, wantReply)

	// Names which are not handled by any handler get refused.
	query = new(dns.Msg)
	query.SetQuestion("example.net.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeRefused
	testQuery(t, service, query, wantReply)
}

func TestRedirect(t *testing.T) {
	service := New([]string{"handler1", "handler2"})
	service.SetHandler("handler1", handlerFunc(func(r *Request) {
		if IsSubDomain("example.net.", r.QnameCanonical) {
			r.SetAuthoritative()
			if r.Qtype == dns.TypeA || r.Qtype == dns.TypeANY {
				rr := new(dns.A)
				rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 300}
				rr.A = net.IP{1, 2, 3, 4}
				r.Reply.Answer = append(r.Reply.Answer, rr)
			}
			r.SendReply()
		}
	}))
	service.SetHandler("handler2", handlerFunc(func(r *Request) {
		if IsSubDomain("example.com.", r.QnameCanonical) {
			switch r.QnameCanonical {
			case "1.example.com.":
				r.AddCNAME("2.example.com.", 30)
			case "2.example.com.":
				r.AddCNAME("example.net.", 30)

			case "loop.example.com.":
				r.AddCNAME("loop.example.com.", 30)

			case "loop1.example.com.":
				r.AddCNAME("loop2.example.com.", 30)
			case "loop2.example.com.":
				r.AddCNAME("loop3.example.com.", 30)
			case "loop3.example.com.":
				r.AddCNAME("loop1.example.com.", 30)

			case "chain1.example.com.":
				r.AddCNAME("chain2.example.com.", 30)
			case "chain2.example.com.":
				r.AddCNAME("chain3.example.com.", 30)
			case "chain3.example.com.":
				r.AddCNAME("chain4.example.com.", 30)
			case "chain4.example.com.":
				r.AddCNAME("chain5.example.com.", 30)
			case "chain5.example.com.":
				r.AddCNAME("chain6.example.com.", 30)
			case "chain6.example.com.":
				r.AddCNAME("chain7.example.com.", 30)
			case "chain7.example.com.":
				r.AddCNAME("chain8.example.com.", 30)
			case "chain8.example.com.":
				r.AddCNAME("chain9.example.com.", 30)
			case "chain9.example.com.":
				r.AddCNAME("chain10.example.com.", 30)

			default:
				r.SendRcode(dns.RcodeNameError)
			}
		}
	}))

	// CNAME redirects are followed.
	query := new(dns.Msg)
	query.SetQuestion("1.example.com.", dns.TypeA)
	wantReply := query.Copy()
	wantReply.Answer = []dns.RR{
		test.RR("1.example.com.	30	IN	CNAME	2.example.com."),
		test.RR("2.example.com.	30	IN	CNAME	example.net."),
		test.RR("example.net.	300	IN	A	1.2.3.4"),
	}
	testQuery(t, service, query, wantReply)

	// Queries of type CNAME or ANY do not follow the redirect.
	query = new(dns.Msg)
	query.SetQuestion("1.example.com.", dns.TypeCNAME)
	wantReply = query.Copy()
	wantReply.Answer = []dns.RR{test.RR("1.example.com.	30	IN	CNAME	2.example.com.")}
	testQuery(t, service, query, wantReply)

	query = new(dns.Msg)
	query.SetQuestion("1.example.com.", dns.TypeANY)
	wantReply = query.Copy()
	wantReply.Answer = []dns.RR{test.RR("1.example.com.	30	IN	CNAME	2.example.com.")}
	testQuery(t, service, query, wantReply)

	// Loops are detected.
	query = new(dns.Msg)
	query.SetQuestion("loop.example.com.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Answer = []dns.RR{test.RR("loop.example.com.	30	IN	CNAME	loop.example.com.")}
	wantReply.Rcode = dns.RcodeServerFailure
	testQuery(t, service, query, wantReply)

	// Loops are detected.
	query = new(dns.Msg)
	query.SetQuestion("loop1.example.com.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Answer = []dns.RR{
		test.RR("loop1.example.com.	30	IN	CNAME	loop2.example.com."),
		test.RR("loop2.example.com.	30	IN	CNAME	loop3.example.com."),
		test.RR("loop3.example.com.	30	IN	CNAME	loop1.example.com."),
	}
	wantReply.Rcode = dns.RcodeServerFailure
	testQuery(t, service, query, wantReply)

	// Number of redirects is limited.
	query = new(dns.Msg)
	query.SetQuestion("chain1.example.com.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Answer = []dns.RR{
		test.RR("chain1.example.com.	30	IN	CNAME	chain2.example.com."),
		test.RR("chain2.example.com.	30	IN	CNAME	chain3.example.com."),
		test.RR("chain3.example.com.	30	IN	CNAME	chain4.example.com."),
		test.RR("chain4.example.com.	30	IN	CNAME	chain5.example.com."),
		test.RR("chain5.example.com.	30	IN	CNAME	chain6.example.com."),
		test.RR("chain6.example.com.	30	IN	CNAME	chain7.example.com."),
		test.RR("chain7.example.com.	30	IN	CNAME	chain8.example.com."),
		test.RR("chain8.example.com.	30	IN	CNAME	chain9.example.com."),
	}
	wantReply.Rcode = dns.RcodeServerFailure
	testQuery(t, service, query, wantReply)
}

func TestFlags(t *testing.T) {
	service := New(nil)

	query := new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)

	// Set flags which should be copied to the reply.
	query.RecursionDesired = true
	query.CheckingDisabled = true

	wantReply := query.Copy()
	wantReply.Authoritative = true
	wantReply.Answer = []dns.RR{test.RR("localhost.	300	IN	A	127.0.0.1")}

	// Set flags which should be ignored.
	query.Authoritative = true
	query.RecursionAvailable = true
	query.Zero = true
	query.AuthenticatedData = true
	query.Rcode = dns.RcodeRefused

	testQuery(t, service, query, wantReply)
}

func TestOPT(t *testing.T) {
	service := New(nil)

	query := new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	wantReply := query.Copy()
	wantReply.Authoritative = true
	wantReply.Answer = []dns.RR{test.RR("localhost.	300	IN	A	127.0.0.1")}
	wantReply.SetEdns0(advertiseUDPSize, false)
	query.SetEdns0(512, false)
	testQuery(t, service, query, wantReply)

	// DNSSEC ok flag.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Authoritative = true
	wantReply.Answer = []dns.RR{test.RR("localhost.	300	IN	A	127.0.0.1")}
	wantReply.SetEdns0(advertiseUDPSize, true)
	query.SetEdns0(512, true)
	testQuery(t, service, query, wantReply)
}

func TestInvalidQuery(t *testing.T) {
	service := New(nil)

	// Valid query.
	query := new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	wantReply := query.Copy()
	wantReply.Authoritative = true
	wantReply.Answer = []dns.RR{test.RR("localhost.	300	IN	A	127.0.0.1")}
	testQuery(t, service, query, wantReply)

	// Not query opcode.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	query.Opcode = dns.OpcodeNotify
	wantReply = query.Copy()
	wantReply.RecursionDesired = false
	wantReply.Rcode = dns.RcodeNotImplemented
	testMsg(t, service, query, wantReply)

	// Truncated.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeFormatError
	query.Truncated = true
	testQuery(t, service, query, wantReply)

	// Multiple OPTs.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeFormatError
	query.SetEdns0(512, false)
	query.SetEdns0(512, false)
	testQuery(t, service, query, wantReply)

	// Unknown OPT version.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeBadVers
	wantReply.SetEdns0(advertiseUDPSize, false)
	query.SetEdns0(512, false)
	query.Extra[0].(*dns.OPT).SetVersion(1)
	testQuery(t, service, query, wantReply)

	// Invalid OPT name.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeA)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeFormatError
	query.SetEdns0(512, false)
	query.Extra[0].(*dns.OPT).Hdr.Name = "localhost."
	testQuery(t, service, query, wantReply)

	// No question.
	query = new(dns.Msg)
	query.Id = dns.Id()
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeRefused
	testQuery(t, service, query, wantReply)

	// Multiple questions.
	query = new(dns.Msg)
	query.Id = dns.Id()
	query.Question = []dns.Question{
		{Name: "localhost.", Qtype: dns.TypeA, Qclass: dns.ClassINET},
		{Name: "localhost.", Qtype: dns.TypeAAAA, Qclass: dns.ClassINET},
	}
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeRefused
	testQuery(t, service, query, wantReply)

	// OPT qtype.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeOPT)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeFormatError
	testQuery(t, service, query, wantReply)

	// Zone transfer.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeAXFR)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeRefused
	testQuery(t, service, query, wantReply)

	// Zone transfer.
	query = new(dns.Msg)
	query.SetQuestion("localhost.", dns.TypeIXFR)
	wantReply = query.Copy()
	wantReply.Rcode = dns.RcodeRefused
	testQuery(t, service, query, wantReply)
}
