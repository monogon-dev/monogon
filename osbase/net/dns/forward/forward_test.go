package forward

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/miekg/dns"

	netDNS "source.monogon.dev/osbase/net/dns"
	"source.monogon.dev/osbase/net/dns/test"
	"source.monogon.dev/osbase/supervisor"
)

func rrStrings(rrs []dns.RR) []string {
	s := make([]string, len(rrs))
	for i, rr := range rrs {
		s[i] = rr.String()
	}
	return s
}

func expectReply(t *testing.T, req *netDNS.Request, wantReply proxyReply) {
	t.Helper()
	if !req.Handled {
		t.Errorf("Request was not handled")
	}
	if got, want := req.Reply.Truncated, wantReply.Truncated; got != want {
		t.Errorf("Want truncated %v, got %v", want, got)
	}
	if got, want := req.Reply.Rcode, wantReply.Rcode; got != want {
		t.Errorf("Want rcode %v, got %v", want, got)
	}

	wantExtra := wantReply.Extra
	if req.Ropt != nil {
		wantOpt := &dns.OPT{Hdr: dns.RR_Header{Name: ".", Rrtype: dns.TypeOPT}}
		wantOpt.Option = wantReply.Options
		wantOpt.SetUDPSize(req.Ropt.UDPSize())
		wantOpt.SetDo(req.Qopt.Do())
		wantExtra = slices.Concat(wantExtra, []dns.RR{wantOpt})
	}
	checkReplySection(t, "answer", req.Reply.Answer, wantReply.Answer)
	checkReplySection(t, "ns", req.Reply.Ns, wantReply.Ns)
	checkReplySection(t, "extra", req.Reply.Extra, wantExtra)
}

func checkReplySection(t *testing.T, sectionName string, got []dns.RR, want []dns.RR) {
	t.Helper()
	gotStr := rrStrings(got)
	wantStr := rrStrings(want)
	if !slices.Equal(gotStr, wantStr) {
		t.Errorf("Want %s:\n%s\nGot:\n%v", sectionName,
			strings.Join(wantStr, "\n"), strings.Join(gotStr, "\n"))
	}
}

type fakeTime struct {
	time atomic.Pointer[time.Time]
}

func (f *fakeTime) now() time.Time {
	t := f.time.Load()
	if t != nil {
		return *t
	}
	return time.Time{}
}

func (f *fakeTime) set(t time.Time) {
	f.time.Store(&t)
}

func (f *fakeTime) add(t time.Duration) {
	f.set(f.now().Add(t))
}

func TestUpstreams(t *testing.T) {
	answerRecord1 := test.RR("example.com. IN A 127.0.0.1")
	s1 := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Answer = append(ret.Answer, answerRecord1)
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s1.Close()
	answerRecord2 := test.RR("2.example.com. IN A 127.0.0.1")
	s2 := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Answer = append(ret.Answer, answerRecord2)
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s2.Close()

	forward := New()
	supervisor.TestHarness(t, forward.Run)

	// If no upstreams are set, return an error.
	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, replyNoUpstreams)

	forward.DNSServers.Set([]string{s1.Addr})
	time.Sleep(10 * time.Millisecond)

	req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord1}})

	forward.DNSServers.Set([]string{s2.Addr})
	time.Sleep(10 * time.Millisecond)

	// New DNS server should be used.
	req = netDNS.CreateTestRequest("2.example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord2}})
}

// TestHealthcheck tests that if one of multiple upstreams is broken,
// this upstream receives health check queries, and that client queries
// succeed since they are retried on the good upstream.
func TestHealthcheck(t *testing.T) {
	var healthcheckCount atomic.Int64

	sGood := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Rcode = dns.RcodeNameError
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer sGood.Close()
	sBad := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		if r.Question[0] == (dns.Question{Name: ".", Qtype: dns.TypeNS, Qclass: dns.ClassINET}) {
			healthcheckCount.Add(1)
		}
		w.Write([]byte("this is not a dns message"))
	})
	defer sBad.Close()

	forward := New()
	forward.DNSServers.Set([]string{sGood.Addr, sBad.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	for i := range 100 {
		req := netDNS.CreateTestRequest(fmt.Sprintf("%v.example.com.", i), dns.TypeA, "udp")
		forward.HandleDNS(req)
		expectReply(t, req, proxyReply{Rcode: dns.RcodeNameError})
	}

	if healthcheckCount.Load() == 0 {
		t.Error("Expected to see at least one healthcheck query.")
	}
}

func TestRecursionDesired(t *testing.T) {
	forward := New()

	// If RecursionDesired is not set, refuse query.
	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	req.Reply.RecursionDesired = false
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Rcode: dns.RcodeRefused})

	// If RecursionDesired is not set and the query was redirected, send reply as is.
	req = netDNS.CreateTestRequest("external.default.scv.cluster.local.", dns.TypeA, "udp")
	req.Reply.RecursionDesired = false
	req.AddCNAME("example.com.", 10)
	req.Handled = false
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{
		Answer: []dns.RR{test.RR("external.default.scv.cluster.local. 10 IN CNAME example.com.")},
	})
}

// TestCache tests that both concurrent and sequential queries use the cache.
func TestCache(t *testing.T) {
	var queryCount atomic.Int64

	answerRecord := test.RR("example.com. IN A 127.0.0.1")
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		queryCount.Add(1)
		// Sleep a bit until all concurrent queries are blocked.
		time.Sleep(10 * time.Millisecond)
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Answer = append(ret.Answer, answerRecord)
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	wg := sync.WaitGroup{}
	for range 3 {
		wg.Add(1)
		go func() {
			req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
			forward.HandleDNS(req)
			expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord}})
			wg.Done()
		}()
	}
	wg.Wait()

	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord}})

	// tcp query
	req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "tcp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord}})

	// query without OPT
	req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	req.Qopt = nil
	req.Ropt = nil
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord}})

	if got, want := queryCount.Load(), int64(1); got != want {
		t.Errorf("Want %v queries, got %v", want, got)
	}
}

func TestTtl(t *testing.T) {
	var queryCount atomic.Int64
	answer := []dns.RR{
		test.RR("1.example.com. 3 CNAME 2.example.com."),
		test.RR("2.example.com. 3600 IN A 127.0.0.1"),
	}
	answerDecrement := []dns.RR{
		test.RR("1.example.com. 2 CNAME 2.example.com."),
		test.RR("2.example.com. 3599 IN A 127.0.0.1"),
	}
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		queryCount.Add(1)
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Answer = answer
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	ft := fakeTime{}
	ft.set(time.Now())
	forward.now = ft.now
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	req := netDNS.CreateTestRequest("1.example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: answer})

	ft.add(1500 * time.Millisecond)

	// TTL of cached reply should be decremented.
	req = netDNS.CreateTestRequest("1.example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: answerDecrement})

	ft.add(2000 * time.Millisecond)

	// Cache expired.
	req = netDNS.CreateTestRequest("1.example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: answer})

	if got, want := queryCount.Load(), int64(2); got != want {
		t.Errorf("Want %v queries, got %v", want, got)
	}
}

// TestShuffle tests that replies from cache have shuffled RRsets.
// In this example, only the A records should be shuffled,
// the CNAMEs and RRSIG should stay where they are.
func TestShuffle(t *testing.T) {
	testAnswer := []dns.RR{
		test.RR("1.example.com. CNAME 2.example.com."),
		test.RR("2.example.com. CNAME 3.example.com."),
	}
	// A random shuffle of 20 items is extremely unlikely (1/(20!))
	// to end up in the same order it was originally.
	for i := range 20 {
		testAnswer = append(testAnswer, test.RR(fmt.Sprintf("3.example.com. IN A 127.0.0.%v", i)))
	}
	testAnswer = append(testAnswer, test.RR("3.example.com. RRSIG A 8 2 3600 1 1 1 example.com AAAA AAAA AAAA"))

	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Answer = testAnswer
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "tcp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: testAnswer})

	req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "tcp")
	forward.HandleDNS(req)

	if slices.Equal(rrStrings(req.Reply.Answer), rrStrings(testAnswer)) {
		t.Error("Expected second reply to be shuffled.")
	}
	slices.SortFunc(req.Reply.Answer[2:len(testAnswer)-1], func(a, b dns.RR) int {
		return int(a.(*dns.A).A[3]) - int(b.(*dns.A).A[3])
	})
	expectReply(t, req, proxyReply{Answer: testAnswer})
}

func TestTruncated(t *testing.T) {
	var queryCount atomic.Int64
	answerRecord1 := test.RR("example.com. IN A 127.0.0.1")
	answerRecord2 := test.RR("example.com. IN A 127.0.0.2")
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		queryCount.Add(1)
		tcp := w.RemoteAddr().Network() == "tcp"
		ret := new(dns.Msg)
		ret.SetReply(r)
		if tcp {
			ret.Answer = append(ret.Answer, answerRecord2)
		} else {
			ret.Answer = append(ret.Answer, answerRecord1)
			ret.Truncated = true
		}
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	ft := fakeTime{}
	ft.set(time.Now())
	forward.now = ft.now
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	for range 2 {
		// Truncated replies are cached and returned for udp queries.
		req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
		forward.HandleDNS(req)
		expectReply(t, req, proxyReply{Truncated: true, Answer: []dns.RR{answerRecord1}})
	}

	// Cached truncated replies are not used for tcp queries.
	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "tcp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord2}})

	// Subsequent udp queries get the tcp reply.
	req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord2}})

	ft.add(10000 * time.Second)

	// After the cache expires, tcp is used.
	req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord2}})

	if got, want := queryCount.Load(), int64(3); got != want {
		t.Errorf("Want %v queries, got %v", want, got)
	}
}

type testQuery struct {
	qtype            uint16
	dnssec           bool
	checkingDisabled bool
}

// TestQueries tests that queries which differ in relevant fields
// result in separate upstream queries and are separately cached.
func TestQueries(t *testing.T) {
	var queryCount atomic.Int64

	answerRecord := test.RR("example.com. IN A 127.0.0.1")
	answerRecordAAAA := test.RR("example.com. IN AAAA ::1")
	answerRecordRRSIG := test.RR("example.com. IN RRSIG A 8 2 3600 1 1 1 example.com AAAA AAAA AAAA")
	answerRecordCD := test.RR("example.com. IN A 127.0.0.2")
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		queryCount.Add(1)
		ret := new(dns.Msg)
		ret.SetReply(r)
		if r.Question[0].Name != "example.com." || r.Question[0].Qclass != dns.ClassINET {
			t.Errorf("Unexpected Name or Qclass")
			return
		}
		switch (testQuery{r.Question[0].Qtype, r.IsEdns0().Do(), r.CheckingDisabled}) {
		case testQuery{dns.TypeA, false, false}:
			ret.Answer = append(ret.Answer, answerRecord)
		case testQuery{dns.TypeAAAA, false, false}:
			ret.Answer = append(ret.Answer, answerRecordAAAA)
		case testQuery{dns.TypeA, true, false}:
			ret.Answer = append(ret.Answer, answerRecord)
			ret.Answer = append(ret.Answer, answerRecordRRSIG)
		case testQuery{dns.TypeA, false, true}:
			ret.Answer = append(ret.Answer, answerRecordCD)
		}
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	for range 2 {
		req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
		forward.HandleDNS(req)
		expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord}})

		// different qtype
		req = netDNS.CreateTestRequest("example.com.", dns.TypeAAAA, "udp")
		forward.HandleDNS(req)
		expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecordAAAA}})

		// DNSSEC flag
		req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
		req.Qopt.SetDo()
		req.Ropt.SetDo()
		forward.HandleDNS(req)
		expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecord, answerRecordRRSIG}})

		// CheckingDisabled flag
		req = netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
		req.Reply.CheckingDisabled = true
		forward.HandleDNS(req)
		expectReply(t, req, proxyReply{Answer: []dns.RR{answerRecordCD}})
	}

	if got, want := queryCount.Load(), int64(4); got != want {
		t.Errorf("Want %v queries, got %v", want, got)
	}
}

// TestOPT tests that only certains OPT options are forwarded
// in query and reply.
func TestOPT(t *testing.T) {
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		wantOpt := &dns.OPT{}
		wantOpt.Hdr.Name = "."
		wantOpt.Hdr.Rrtype = dns.TypeOPT
		wantOpt.SetUDPSize(r.IsEdns0().UDPSize())
		wantOpt.Option = []dns.EDNS0{
			&dns.EDNS0_DAU{AlgCode: []uint8{1, 4}},
			&dns.EDNS0_DHU{AlgCode: []uint8{5}},
			&dns.EDNS0_N3U{AlgCode: []uint8{6}},
		}
		if got, want := r.IsEdns0().String(), wantOpt.String(); got != want {
			t.Errorf("Wanted opt %q, got %q", want, got)
		}

		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Rcode = dns.RcodeBadAlg
		ret.SetEdns0(512, false)
		ropt := ret.Extra[0].(*dns.OPT)
		ropt.Option = []dns.EDNS0{
			&dns.EDNS0_NSID{Nsid: "c0ffee"},
			&dns.EDNS0_EDE{InfoCode: dns.ExtendedErrorCodeCensored, ExtraText: "****"},
			&dns.EDNS0_EDE{InfoCode: dns.ExtendedErrorCodeDNSKEYMissing, ExtraText: "second problem"},
			&dns.EDNS0_PADDING{Padding: []byte{0, 0, 0}},
		}
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	req.Qopt.Option = []dns.EDNS0{
		&dns.EDNS0_NSID{Nsid: ""},
		&dns.EDNS0_EDE{InfoCode: dns.ExtendedErrorCodeDNSBogus, ExtraText: "huh?"},
		&dns.EDNS0_DAU{AlgCode: []uint8{1, 4}},
		&dns.EDNS0_DHU{AlgCode: []uint8{5}},
		&dns.EDNS0_N3U{AlgCode: []uint8{6}},
		&dns.EDNS0_PADDING{Padding: []byte{0, 0, 0}},
	}
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{
		Rcode: dns.RcodeBadAlg,
		Options: []dns.EDNS0{
			&dns.EDNS0_EDE{InfoCode: dns.ExtendedErrorCodeCensored, ExtraText: "****"},
			&dns.EDNS0_EDE{InfoCode: dns.ExtendedErrorCodeDNSKEYMissing, ExtraText: "second problem"},
		},
	})
}

// TestBadReply tests that if the qname of the reply is not what was
// sent in the query, the reply is rejected.
func TestBadReply(t *testing.T) {
	answerRecord := test.RR("1.example.com. IN A 127.0.0.1")
	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Question[0].Name = "1.example.com."
		ret.Answer = append(ret.Answer, answerRecord)
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "udp")
	forward.HandleDNS(req)
	expectReply(t, req, replyProtocolError)
}

// TestLargeReply tests that large replies are not stored,
// but still shared with concurrent queries.
func TestLargeReply(t *testing.T) {
	var queryCount atomic.Int64

	var testAnswer []dns.RR
	for i := range 100 {
		testAnswer = append(testAnswer, test.RR(fmt.Sprintf("%v.example.com. IN A 127.0.0.1", i)))
	}

	s := test.NewServer(func(w dns.ResponseWriter, r *dns.Msg) {
		queryCount.Add(1)
		// Sleep a bit until all concurrent queries are blocked.
		time.Sleep(10 * time.Millisecond)
		ret := new(dns.Msg)
		ret.SetReply(r)
		ret.Answer = testAnswer
		err := w.WriteMsg(ret)
		if err != nil {
			t.Error(err)
		}
	})
	defer s.Close()

	forward := New()
	forward.DNSServers.Set([]string{s.Addr})
	supervisor.TestHarness(t, forward.Run)
	time.Sleep(10 * time.Millisecond)

	wg := sync.WaitGroup{}
	for range 2 {
		wg.Add(1)
		go func() {
			req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "tcp")
			forward.HandleDNS(req)
			expectReply(t, req, proxyReply{Answer: testAnswer})
			wg.Done()
		}()
	}
	wg.Wait()

	if got, want := queryCount.Load(), int64(1); got != want {
		t.Errorf("Want %v queries, got %v", want, got)
	}

	req := netDNS.CreateTestRequest("example.com.", dns.TypeA, "tcp")
	forward.HandleDNS(req)
	expectReply(t, req, proxyReply{Answer: testAnswer})

	if got, want := queryCount.Load(), int64(2); got != want {
		t.Errorf("Want %v queries, got %v", want, got)
	}
}
