// Package forward implements a forwarding proxy.
//
// A cache is used to reduce load on the upstream servers. Cached items are only
// used for a short time, because otherwise, we would need to provide a feature
// for flushing the cache. The cache is most useful for taking the load from
// applications making very frequent repeated queries. The cache also doubles as
// a way to merge concurrent identical queries, since items are inserted into
// the cache before sending the query upstream (see also RFC 5452, Section 5).
package forward

// Taken and modified from the Forward plugin of CoreDNS, under Apache 2.0.

import (
	"context"
	"errors"
	"hash/maphash"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/miekg/dns"

	"source.monogon.dev/osbase/event/memory"
	"source.monogon.dev/osbase/net/dns/forward/cache"
	"source.monogon.dev/osbase/net/dns/forward/proxy"
	"source.monogon.dev/osbase/supervisor"
)

const (
	connectionExpire    = 10 * time.Second
	healthcheckInterval = 500 * time.Millisecond
	forwardTimeout      = 5 * time.Second
	maxFails            = 2
	maxConcurrent       = 5000
	maxUpstreams        = 15
)

// Forward represents a plugin instance that can proxy requests to another (DNS)
// server. It has a list of proxies each representing one upstream proxy.
type Forward struct {
	DNSServers memory.Value[[]string]
	upstreams  atomic.Pointer[[]*proxy.Proxy]

	concurrent atomic.Int64

	seed  maphash.Seed
	cache *cache.Cache[*cacheItem]

	// now can be used to override time for testing.
	now func() time.Time
}

// New returns a new Forward.
func New() *Forward {
	return &Forward{
		seed:  maphash.MakeSeed(),
		cache: cache.New[*cacheItem](cacheCapacity),
		now:   time.Now,
	}
}

func (f *Forward) Run(ctx context.Context) error {
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	var lastAddrs []string
	var upstreams []*proxy.Proxy

	w := f.DNSServers.Watch()
	defer w.Close()
	for {
		addrs, err := w.Get(ctx)
		if err != nil {
			for _, p := range upstreams {
				p.Stop()
			}
			f.upstreams.Store(nil)
			return err
		}

		if len(addrs) > maxUpstreams {
			addrs = addrs[:maxUpstreams]
		}

		if slices.Equal(addrs, lastAddrs) {
			continue
		}
		lastAddrs = addrs
		supervisor.Logger(ctx).Infof("New upstream DNS servers: %s", addrs)

		newAddrs := make(map[string]bool)
		for _, addr := range addrs {
			newAddrs[addr] = true
		}
		var newUpstreams []*proxy.Proxy
		for _, p := range upstreams {
			if newAddrs[p.Addr()] {
				delete(newAddrs, p.Addr())
				newUpstreams = append(newUpstreams, p)
			} else {
				p.Stop()
			}
		}
		for newAddr := range newAddrs {
			p := proxy.NewProxy(newAddr)
			p.SetExpire(connectionExpire)
			p.GetHealthchecker().SetRecursionDesired(true)
			p.GetHealthchecker().SetDomain(".")
			p.Start(healthcheckInterval)
			newUpstreams = append(newUpstreams, p)
		}
		upstreams = newUpstreams
		f.upstreams.Store(&newUpstreams)
	}
}

type proxyReply struct {
	// NoStore indicates that the reply should not be stored in the cache.
	// This could be because it is cheap to obtain or expensive to store.
	NoStore bool

	Truncated bool
	Rcode     int
	Answer    []dns.RR
	Ns        []dns.RR
	Extra     []dns.RR
	Options   []dns.EDNS0
}

var (
	replyConcurrencyLimit = proxyReply{
		NoStore: true,
		Rcode:   dns.RcodeServerFailure,
		Options: []dns.EDNS0{&dns.EDNS0_EDE{
			InfoCode:  dns.ExtendedErrorCodeOther,
			ExtraText: "too many concurrent queries",
		}},
	}
	replyNoUpstreams = proxyReply{
		NoStore: true,
		Rcode:   dns.RcodeRefused,
		Options: []dns.EDNS0{&dns.EDNS0_EDE{
			InfoCode:  dns.ExtendedErrorCodeOther,
			ExtraText: "no upstream DNS servers configured",
		}},
	}
	replyProtocolError = proxyReply{
		Rcode: dns.RcodeServerFailure,
		Options: []dns.EDNS0{&dns.EDNS0_EDE{
			InfoCode:  dns.ExtendedErrorCodeNetworkError,
			ExtraText: "DNS protocol error when querying upstream DNS server",
		}},
	}
	replyTimeout = proxyReply{
		Rcode: dns.RcodeServerFailure,
		Options: []dns.EDNS0{&dns.EDNS0_EDE{
			InfoCode:  dns.ExtendedErrorCodeNetworkError,
			ExtraText: "timeout when querying upstream DNS server",
		}},
	}
	replyNetworkError = proxyReply{
		Rcode: dns.RcodeServerFailure,
		Options: []dns.EDNS0{&dns.EDNS0_EDE{
			InfoCode:  dns.ExtendedErrorCodeNetworkError,
			ExtraText: "network error when querying upstream DNS server",
		}},
	}
)

func (f *Forward) queryProxies(
	question dns.Question,
	dnssec bool,
	checkingDisabled bool,
	queryOptions []dns.EDNS0,
	useTCP bool,
) proxyReply {
	count := f.concurrent.Add(1)
	defer f.concurrent.Add(-1)
	if count > maxConcurrent {
		rejectsCount.WithLabelValues("concurrency_limit").Inc()
		return replyConcurrencyLimit
	}

	// Construct outgoing query.
	qopt := new(dns.OPT)
	qopt.Hdr.Name = "."
	qopt.Hdr.Rrtype = dns.TypeOPT
	qopt.SetUDPSize(proxy.AdvertiseUDPSize)
	if dnssec {
		qopt.SetDo()
	}
	qopt.Option = queryOptions
	m := &dns.Msg{
		MsgHdr: dns.MsgHdr{
			Opcode:           dns.OpcodeQuery,
			RecursionDesired: true,
			CheckingDisabled: checkingDisabled,
		},
		Question: []dns.Question{question},
		Extra:    []dns.RR{qopt},
	}

	var list []*proxy.Proxy
	if upstreams := f.upstreams.Load(); upstreams != nil {
		list = randomList(*upstreams)
	}

	if len(list) == 0 {
		rejectsCount.WithLabelValues("no_upstreams").Inc()
		return replyNoUpstreams
	}

	proto := "udp"
	if useTCP {
		proto = "tcp"
	}

	var (
		curUpstream *proxy.Proxy
		curStart    time.Time
		ret         *dns.Msg
		err         error
	)
	recordDuration := func(rcode string) {
		upstreamDuration.WithLabelValues(curUpstream.Addr(), proto, rcode).Observe(time.Since(curStart).Seconds())
	}

	fails := 0
	i := 0
	listStart := time.Now()
	deadline := listStart.Add(forwardTimeout)
	for {
		if i >= len(list) {
			// reached the end of list, reset to begin
			i = 0
			fails = 0

			// Sleep for a bit if the last time we started going through the list was
			// very recent.
			time.Sleep(time.Until(listStart.Add(time.Second)))
			listStart = time.Now()
		}

		curUpstream = list[i]
		i++
		if curUpstream.Down(maxFails) {
			fails++
			if fails < len(list) {
				continue
			}
			// All upstream proxies are dead, assume healthcheck is completely broken
			// and connect to a random upstream.
			healthcheckBrokenCount.Inc()
		}

		curStart = time.Now()

		for {
			ret, err = curUpstream.Connect(m, useTCP)

			if errors.Is(err, proxy.ErrCachedClosed) {
				// Remote side closed conn, can only happen with TCP.
				continue
			}
			break
		}

		if err != nil {
			// Kick off health check to see if *our* upstream is broken.
			curUpstream.Healthcheck()

			retry := fails < len(list) && time.Now().Before(deadline)
			var dnsError *dns.Error
			switch {
			case errors.Is(err, os.ErrDeadlineExceeded):
				recordDuration("timeout")
				if !retry {
					return replyTimeout
				}
			case errors.As(err, &dnsError):
				recordDuration("protocol_error")
				if !retry {
					return replyProtocolError
				}
			default:
				recordDuration("network_error")
				if !retry {
					return replyNetworkError
				}
			}
			continue
		}

		break
	}

	if !ret.Response || ret.Opcode != dns.OpcodeQuery || len(ret.Question) != 1 {
		recordDuration("protocol_error")
		return replyProtocolError
	}

	if ret.Truncated && useTCP {
		recordDuration("protocol_error")
		return replyProtocolError
	}
	if ret.Truncated {
		proto = "udp_truncated"
	}

	// Check that the reply matches the question.
	retq := ret.Question[0]
	if retq.Qtype != question.Qtype || retq.Qclass != question.Qclass ||
		(retq.Name != question.Name && dns.CanonicalName(retq.Name) != question.Name) {
		recordDuration("protocol_error")
		return replyProtocolError
	}

	// Extract OPT from reply.
	var ropt *dns.OPT
	var options []dns.EDNS0
	for i := len(ret.Extra) - 1; i >= 0; i-- {
		if rr, ok := ret.Extra[i].(*dns.OPT); ok {
			if ropt != nil {
				// Found more than one OPT.
				recordDuration("protocol_error")
				return replyProtocolError
			}
			ropt = rr
			ret.Extra = append(ret.Extra[:i], ret.Extra[i+1:]...)
		}
	}
	if ropt != nil {
		if ropt.Version() != 0 || ropt.Hdr.Name != "." {
			recordDuration("protocol_error")
			return replyProtocolError
		}
		// Forward Extended DNS Error options.
		for _, option := range ropt.Option {
			switch option.(type) {
			case *dns.EDNS0_EDE:
				options = append(options, option)
			}
		}
	}

	rcode, ok := dns.RcodeToString[ret.Rcode]
	if !ok {
		// There are 4096 possible Rcodes, so it's probably still fine
		// to put it in a metric label.
		rcode = strconv.Itoa(ret.Rcode)
	}
	recordDuration(rcode)

	// AuthenticatedData is intentionally not copied from the proxy reply because
	// we don't have a secure channel to the proxy.
	return proxyReply{
		// Don't store large messages in the cache. Such large messages are very
		// rare, and this protects against the cache using huge amounts of memory.
		// DNS messages over TCP can be up to 64 KB in size, and after decompression
		// this could go over 1 MB of memory usage.
		NoStore: ret.Len() > cacheMaxItemSize,

		Truncated: ret.Truncated,
		Rcode:     ret.Rcode,
		Answer:    ret.Answer,
		Ns:        ret.Ns,
		Extra:     ret.Extra,
		Options:   options,
	}
}

func randomList(p []*proxy.Proxy) []*proxy.Proxy {
	switch len(p) {
	case 1:
		return p
	case 2:
		if rand.Int()%2 == 0 {
			return []*proxy.Proxy{p[1], p[0]} // swap
		}
		return p
	}

	shuffled := slices.Clone(p)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
