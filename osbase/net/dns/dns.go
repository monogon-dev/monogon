// Package dns provides a DNS server for resolving services against.
package dns

import (
	"context"
	"fmt"
	"net/netip"
	"runtime/debug"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/miekg/dns"

	"source.monogon.dev/osbase/supervisor"
)

// Service is a DNS server service with configurable handlers.
//
// The number and names of handlers is fixed when New is called. For each name
// in handlerNames there is a corresponding pointer to a handler in the handlers
// slice at the same index, which can be atomically updated at runtime through
// its atomic.Pointer via the SetHandler function.
type Service struct {
	handlerNames []string
	handlers     []atomic.Pointer[Handler]
}

type serviceCtx struct {
	service *Service
	ctx     context.Context
}

// New creates a Service instance. DNS handlers with the names given in
// handlerNames must be set with SetHandler. When serving DNS queries, they will
// be tried in the order they appear here. Doing it this way instead of directly
// passing a []Handler avoids circular dependencies.
func New(handlerNames []string) *Service {
	return &Service{
		handlerNames: handlerNames,
		handlers:     make([]atomic.Pointer[Handler], len(handlerNames)),
	}
}

// Run runs the DNS service.
func (s *Service) Run(ctx context.Context) error {
	addr4 := "127.0.0.1:53"
	addr6 := "[::1]:53"
	supervisor.Run(ctx, "udp4", func(ctx context.Context) error {
		return s.runListener(ctx, addr4, "udp")
	})
	supervisor.Run(ctx, "tcp4", func(ctx context.Context) error {
		return s.runListener(ctx, addr4, "tcp")
	})
	supervisor.Run(ctx, "udp6", func(ctx context.Context) error {
		return s.runListener(ctx, addr6, "udp")
	})
	supervisor.Run(ctx, "tcp6", func(ctx context.Context) error {
		return s.runListener(ctx, addr6, "tcp")
	})
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}

// RunListenerAddr runs a DNS listener on a specific address.
func (s *Service) RunListenerAddr(ctx context.Context, addr string) error {
	supervisor.Run(ctx, "udp", func(ctx context.Context) error {
		return s.runListener(ctx, addr, "udp")
	})
	supervisor.Run(ctx, "tcp", func(ctx context.Context) error {
		return s.runListener(ctx, addr, "tcp")
	})
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}

func (s *Service) runListener(ctx context.Context, addr string, network string) error {
	handler := &serviceCtx{service: s, ctx: ctx}
	server := &dns.Server{Addr: addr, Net: network, ReusePort: true, Handler: handler}
	server.NotifyStartedFunc = func() {
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		supervisor.Logger(ctx).Infof("DNS server listening on %s %s", addr, network)
		go func() {
			<-ctx.Done()
			server.Shutdown()
		}()
	}
	return server.ListenAndServe()
}

// Requests

// Request represents an incoming DNS query that is being handled.
type Request struct {
	// Reply is the reply that will be sent, and should be filled in by the
	// handler. It is guaranteed to contain exactly one question.
	Reply *dns.Msg
	// Writer will be used to send the reply, and contains network information.
	Writer dns.ResponseWriter

	// Qopt is the OPT record from the query, or nil if not present.
	Qopt *dns.OPT
	// Ropt, if non-nil, is the OPT record that will be added to the reply. The
	// handler can modify this as needed. Ropt is nil when Qopt is nil.
	Ropt *dns.OPT

	// Qname contains the current question name. This is different from the
	// original question in Reply.Question[0].Name if a CNAME has been followed
	// already.
	Qname string

	// QnameCanonical contains the canonicalized name of the question. This means
	// that ASCII letters are lowercased.
	QnameCanonical string

	// Qtype contains the question type for convenient access.
	Qtype uint16

	// Handled is set to true when the current question name has been handled and
	// no other handlers should be attempted.
	Handled bool

	// done is set to true when a reply has been sent. When a CNAME is
	// encountered, Handled is set to true, but done is false.
	done bool
}

// SetAuthoritative marks the reply as authoritative.
func (r *Request) SetAuthoritative() {
	// Only set the AA bit if the question has not yet been redirected by CNAME.
	// See RFC 1034 6.2.7
	if r.Qname == r.Reply.Question[0].Name {
		r.Reply.Authoritative = true
	}
}

// SendReply sends the reply. It may only be called once.
func (r *Request) SendReply() {
	if r.Handled {
		panic("SendReply called twice for the same DNS request")
	}
	r.Handled = true
	r.done = true

	if r.Ropt != nil {
		r.Reply.Extra = append(r.Reply.Extra, r.Ropt)
	} else {
		// Cannot use extended RCODEs without an OPT, so replace with SERVFAIL.
		if r.Reply.Rcode > 0xF {
			r.Reply.Rcode = dns.RcodeServerFailure
		}
	}

	size := uint16(0)
	if r.Writer.RemoteAddr().Network() == "tcp" {
		size = dns.MaxMsgSize
	} else if r.Qopt != nil {
		size = r.Qopt.UDPSize()
	}
	if size < dns.MinMsgSize {
		size = dns.MinMsgSize
	}
	r.Reply.Truncate(int(size))
	if !r.Reply.Compress && r.Reply.Len() >= 1024 {
		r.Reply.Compress = true
	}

	r.Writer.WriteMsg(r.Reply)
}

// SendRcode sets the reply RCODE and sends the reply.
func (r *Request) SendRcode(rcode int) {
	r.Reply.Rcode = rcode
	r.SendReply()
}

// AddExtendedError adds an Extended DNS Error Option if the reply has an OPT.
// See RFC 8914.
func (r *Request) AddExtendedError(infoCode uint16, extraText string) {
	if r.Ropt != nil {
		r.Ropt.Option = append(r.Ropt.Option, &dns.EDNS0_EDE{InfoCode: infoCode, ExtraText: extraText})
	}
}

// AddCNAME adds a CNAME record to the answer section, and either sends the
// reply if the query is for the CNAME itself, or else marks the lookup to be
// restarted at the new name. target must be fully qualified.
func (r *Request) AddCNAME(target string, ttl uint32) {
	rr := new(dns.CNAME)
	rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: ttl}
	rr.Target = target
	r.Reply.Answer = append(r.Reply.Answer, rr)

	if r.Qtype == dns.TypeCNAME || r.Qtype == dns.TypeANY {
		r.SendReply()
	} else {
		r.Handled = true
		r.Qname = target
		r.QnameCanonical = dns.CanonicalName(r.Qname)
	}
}

// Handlers

// Handler can handle DNS requests. The handler should first inspect the query
// and decide if it wants to handle it. If not, it should return immediately.
// The next handler will then be tried. Otherwise, it should fill in the Reply,
// and then call SendReply. The Answer section may already contain CNAMEs that
// have been followed.
type Handler interface {
	HandleDNS(r *Request)
}

// SetHandler sets the handler of the given name. This name must have been
// registered when creating the Service. As long as SetHandler has not been
// called for a registered name, any queries that are not already handled by an
// earlier handler in the sequence return SERVFAIL. SetHandler may be called
// multiple times, each call replaces the previous handler of the same name.
func (s *Service) SetHandler(name string, h Handler) {
	for i, iname := range s.handlerNames {
		if iname == name {
			s.handlers[i].Store(&h)
			return
		}
	}
	panic(fmt.Sprintf("Attempted to set undeclared DNS handler: %q", name))
}

// EmptyDNSHandler is a handler that does not handle any queries. It can be used
// as a placeholder with SetHandler when a handler is inactive.
type EmptyDNSHandler struct{}

func (EmptyDNSHandler) HandleDNS(*Request) {}

// Serving requests

// advertiseUDPSize is the maximum message size that we advertise in the OPT RR
// of UDP messages. This is calculated as the minimum IPv6 MTU (1280) minus size
// of IPv6 (40) and UDP (8) headers.
const advertiseUDPSize = 1232

// ServeDNS implements dns.Handler.
func (s *serviceCtx) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	defer func() {
		if rec := recover(); rec != nil {
			supervisor.Logger(s.ctx).Errorf("panic in DNS handler: %v, stacktrace: %s", rec, string(debug.Stack()))
		}
	}()

	// Only QUERY opcode is implemented.
	if r.Opcode != dns.OpcodeQuery {
		m := new(dns.Msg)
		m.SetRcode(r, dns.RcodeNotImplemented)
		w.WriteMsg(m)
		return
	}

	if r.Truncated {
		m := new(dns.Msg)
		m.RecursionAvailable = true
		m.SetRcode(r, dns.RcodeFormatError)
		w.WriteMsg(m)
		return
	}

	// Look for an OPT RR.
	var opt *dns.OPT
	for _, rr := range r.Extra {
		if rr, ok := rr.(*dns.OPT); ok {
			if opt != nil {
				// RFC 6891 6.1.1
				// If a query has more than one OPT RR, FORMERR MUST be returned.
				m := new(dns.Msg)
				m.RecursionAvailable = true
				m.SetRcode(r, dns.RcodeFormatError)
				w.WriteMsg(m)
				return
			}
			opt = rr
		}
	}

	// RFC 6891 6.1.3
	// If the VERSION of the query is not implemented, BADVERS MUST be returned.
	if opt != nil && opt.Version() != 0 {
		m := new(dns.Msg)
		m.RecursionAvailable = true
		m.SetRcode(r, dns.RcodeBadVers)
		m.SetEdns0(advertiseUDPSize, false)
		w.WriteMsg(m)
		return
	}

	// If the OPT name is not the root, the OPT is invalid.
	if opt != nil && opt.Hdr.Name != "." {
		m := new(dns.Msg)
		m.RecursionAvailable = true
		m.SetRcode(r, dns.RcodeFormatError)
		w.WriteMsg(m)
		return
	}

	// Reuse the query message as the reply message.
	r.Response = true
	r.Authoritative = false
	r.RecursionAvailable = true
	r.Zero = false
	r.AuthenticatedData = false
	r.Rcode = dns.RcodeSuccess
	r.Extra = nil

	req := &Request{
		Reply:  r,
		Writer: w,
		Qopt:   opt,
	}
	if opt != nil {
		req.Ropt = new(dns.OPT)
		req.Ropt.Hdr.Name = "."
		req.Ropt.Hdr.Rrtype = dns.TypeOPT
		req.Ropt.SetUDPSize(advertiseUDPSize)
		if opt.Do() {
			req.Ropt.SetDo()
		}
	}

	// Refuse queries that don't have exactly one question of class INET, or that
	// have non-empty answer or authority sections.
	if len(r.Question) != 1 || r.Question[0].Qclass != dns.ClassINET || len(r.Answer) != 0 || len(r.Ns) != 0 {
		r.Answer = nil
		r.Ns = nil
		req.SendRcode(dns.RcodeRefused)
		return
	}
	req.Qtype = r.Question[0].Qtype
	req.Qname = r.Question[0].Name
	req.QnameCanonical = dns.CanonicalName(req.Qname)

	switch req.Qtype {
	case dns.TypeOPT:
		// OPT is a pseudo-RR and may only appear in the additional section.
		req.SendRcode(dns.RcodeFormatError)
		return
	case dns.TypeAXFR, dns.TypeIXFR:
		// Zone transfer is not supported.
		req.AddExtendedError(dns.ExtendedErrorCodeNotSupported, "")
		req.SendRcode(dns.RcodeRefused)
		return
	}

	// If we encounter a CNAME, DNS resolution must be restarted with the new
	// name. That's what this loop is for.
	i := 0
	seen := make(map[string]bool)
	for {
		prevName := req.QnameCanonical
		s.service.HandleDNS(req)
		if req.done {
			break
		}
		req.Handled = false
		i++
		seen[prevName] = true
		if seen[req.QnameCanonical] || i > 7 {
			if seen[req.QnameCanonical] {
				req.AddExtendedError(dns.ExtendedErrorCodeOther, "CNAME loop")
			} else {
				req.AddExtendedError(dns.ExtendedErrorCodeOther, "too many CNAME redirects")
			}
			req.SendRcode(dns.RcodeServerFailure)
			break
		}
	}
}

func (s *Service) HandleDNS(r *Request) {
	start := time.Now()

	// Handle localhost.
	if IsSubDomain("localhost.", r.QnameCanonical) {
		handleLocalhost(r)
		return
	}
	if IsSubDomain("127.in-addr.arpa.", r.QnameCanonical) ||
		IsSubDomain("1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.", r.QnameCanonical) {
		handleLocalhostPtr(r)
		return
	}

	// Serve NXDOMAIN for the "invalid." domain, see RFC 6761.
	if IsSubDomain("invalid.", r.QnameCanonical) {
		r.SetAuthoritative()
		r.SendRcode(dns.RcodeNameError)
		return
	}

	for i := range s.handlers {
		handler := s.handlers[i].Load()
		if handler == nil {
			// The application is still starting up. Fail queries instead of leaking
			// local queries to the Internet and sending the wrong reply.
			r.AddExtendedError(dns.ExtendedErrorCodeNotReady, fmt.Sprintf("%s handler not ready", s.handlerNames[i]))
			r.SendRcode(dns.RcodeServerFailure)
			handlerDuration.WithLabelValues(s.handlerNames[i], "not_ready").Observe(time.Since(start).Seconds())
			return
		}
		(*handler).HandleDNS(r)
		if r.Handled {
			rcode, ok := dns.RcodeToString[r.Reply.Rcode]
			if !ok {
				// There are 4096 possible Rcodes, so it's probably still fine to put it
				// in a metric label.
				rcode = strconv.Itoa(r.Reply.Rcode)
			}
			if !r.done {
				rcode = "redirected"
			}
			handlerDuration.WithLabelValues(s.handlerNames[i], rcode).Observe(time.Since(start).Seconds())
			return
		}
	}

	// No handler can handle this request.
	r.SendRcode(dns.RcodeRefused)
}

var (
	localhostA    = netip.MustParseAddr("127.0.0.1").AsSlice()
	localhostAAAA = netip.MustParseAddr("::1").AsSlice()
)

const localhostTtl = 60 * 5

func handleLocalhost(r *Request) {
	r.SetAuthoritative()
	if r.Qtype == dns.TypeA || r.Qtype == dns.TypeANY {
		rr := new(dns.A)
		rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: localhostTtl}
		rr.A = localhostA
		r.Reply.Answer = append(r.Reply.Answer, rr)
	}
	if r.Qtype == dns.TypeAAAA || r.Qtype == dns.TypeANY {
		rr := new(dns.AAAA)
		rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: localhostTtl}
		rr.AAAA = localhostAAAA
		r.Reply.Answer = append(r.Reply.Answer, rr)
	}
	r.SendReply()
}

func handleLocalhostPtr(r *Request) {
	r.SetAuthoritative()
	ip, bits, extra := ParseReverse(r.QnameCanonical)
	if extra {
		// Name with extra labels does not exist (e.g. foo.1.0.0.127.in-addr.arpa.)
		r.Reply.Rcode = dns.RcodeNameError
	} else if bits != ip.BitLen() {
		// Partial reverse name (e.g. 127.in-addr.arpa.) exists but has no records.
	} else if r.Qtype == dns.TypePTR || r.Qtype == dns.TypeANY {
		rr := new(dns.PTR)
		rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypePTR, Class: dns.ClassINET, Ttl: localhostTtl}
		rr.Ptr = "localhost."
		r.Reply.Answer = append(r.Reply.Answer, rr)
	}
	r.SendReply()
}
