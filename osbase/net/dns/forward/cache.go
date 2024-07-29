package forward

import (
	"hash/maphash"
	"math/rand/v2"
	"slices"
	"sync"
	"time"

	"github.com/miekg/dns"

	netDNS "source.monogon.dev/osbase/net/dns"
)

// The cache uses at most cacheMaxItemSize * cacheCapacity = 20 MB of memory.
// Actual memory usage may be slightly higher due to the overhead of in-memory
// data structures compared to the serialized, uncompressed length.
const (
	cacheMaxItemSize = 2048
	cacheCapacity    = 10000
	cacheMinSeconds  = 1
	cacheMaxSeconds  = 5
)

// cacheKey is the key used for cache lookups. Both the DNSSEC ok and the
// Checking Disabled flag influence the reply. While it would be possible to
// always make upstream queries with DNSSEC, and then strip the authenticating
// records if the client did not request it, this would mostly just waste
// bandwidth. In theory, it would be possible to cache NXDOMAINs independently
// of the QTYPE (RFC 2308, Section 5). However, the additional complexity and
// second lookup for each query does not seem worth it.
type cacheKey struct {
	Name             string
	Qtype            uint16
	DNSSEC           bool
	CheckingDisabled bool
}

type cacheItem struct {
	key cacheKey

	// lock protects all fields except key. It also doubles as a way to wait for
	// the reply. A write lock is held for as long as a query is pending.
	lock sync.RWMutex

	reply  proxyReply
	stored time.Time
	// ttl is the number of seconds during which the cached reply can be used.
	ttl uint32
	// seenTruncated is true if we ever saw a truncated response for this key.
	// We will then always use TCP when refetching after the item expires.
	seenTruncated bool
}

func (k cacheKey) hash(seed maphash.Seed) uint64 {
	var h maphash.Hash
	h.SetSeed(seed)
	h.WriteByte(byte(k.Qtype))
	h.WriteByte(byte(k.Qtype >> 8))
	var flags byte
	if k.DNSSEC {
		flags += 1
	}
	if k.CheckingDisabled {
		flags += 2
	}
	h.WriteByte(flags)
	h.WriteString(k.Name)
	return h.Sum64()
}

// valid returns true if the cache item can be used for this query.
func (i *cacheItem) valid(now time.Time, tcp bool) bool {
	expired := now.After(i.stored.Add(time.Duration(i.ttl) * time.Second))
	return !expired && (!tcp || !i.reply.Truncated)
}

func (f *Forward) HandleDNS(r *netDNS.Request) {
	if !r.Reply.RecursionDesired {
		// Only forward queries if the RD flag is set. If the question has been
		// redirected by CNAME, return the reply as is without following the CNAME,
		// else set a REFUSED rcode.
		if r.Qname == r.Reply.Question[0].Name {
			r.Reply.Rcode = dns.RcodeRefused
			rejectsCount.WithLabelValues("no_recursion_desired").Inc()
		}
	} else {
		f.lookupOrForward(r)
	}
	r.SendReply()
}

func (f *Forward) lookupOrForward(r *netDNS.Request) {
	key := cacheKey{
		Name:             r.QnameCanonical,
		Qtype:            r.Qtype,
		DNSSEC:           r.Ropt != nil && r.Ropt.Do(),
		CheckingDisabled: r.Reply.CheckingDisabled,
	}
	hash := key.hash(f.seed)
	tcp := r.Writer.RemoteAddr().Network() == "tcp"

	item, exists := f.cache.Get(hash)
	if !exists {
		// The lookup failed; allocate a new item and try to insert it.
		// Lock the new item before inserting it, such that concurrent queries
		// are blocked until we receive the reply and store it in the item.
		newItem := &cacheItem{key: key}
		newItem.lock.Lock()
		item, exists = f.cache.GetOrPut(hash, newItem)
		if !exists {
			cacheLookupsCount.WithLabelValues("miss").Inc()
			f.forward(r, newItem, hash, tcp)
			newItem.lock.Unlock()
			return
		}
	}
	if item.key != key {
		// We have a hash collision. Replace the existing item.
		cacheLookupsCount.WithLabelValues("miss").Inc()
		newItem := &cacheItem{key: key}
		newItem.lock.Lock()
		f.cache.Put(hash, newItem)
		f.forward(r, newItem, hash, tcp)
		newItem.lock.Unlock()
		return
	}

	// Take a read lock and check if the reply is valid for this query.
	// This blocks if a query for this item is pending.
	item.lock.RLock()
	now := f.now()
	if item.valid(now, tcp) {
		replyFromCache(r, item, now)
		item.lock.RUnlock()
		return
	}
	item.lock.RUnlock()

	item.lock.Lock()
	now = f.now()
	if item.valid(now, tcp) {
		replyFromCache(r, item, now)
		item.lock.Unlock()
		return
	}
	cacheLookupsCount.WithLabelValues("refresh").Inc()
	f.forward(r, item, hash, tcp || item.seenTruncated)
	item.lock.Unlock()
}

func (f *Forward) forward(r *netDNS.Request, item *cacheItem, hash uint64, tcp bool) {
	// Query proxies.
	var queryOptions []dns.EDNS0
	if r.Qopt != nil {
		// Forward DNSSEC algorithm understood options. These are only for
		// statistics and must not influence the reply, so we do not need to include
		// them in the cache key.
		for _, option := range r.Qopt.Option {
			switch option.(type) {
			case *dns.EDNS0_DAU, *dns.EDNS0_DHU, *dns.EDNS0_N3U:
				queryOptions = append(queryOptions, option)
			}
		}
	}

	question := dns.Question{
		Name:   item.key.Name,
		Qtype:  item.key.Qtype,
		Qclass: dns.ClassINET,
	}
	reply := f.queryProxies(question, item.key.DNSSEC, item.key.CheckingDisabled, queryOptions, tcp)

	r.Reply.Truncated = reply.Truncated
	r.Reply.Rcode = reply.Rcode
	r.Reply.Answer = appendOrClip(r.Reply.Answer, reply.Answer)
	r.Reply.Ns = appendOrClip(r.Reply.Ns, reply.Ns)
	r.Reply.Extra = appendOrClip(r.Reply.Extra, reply.Extra)
	if r.Ropt != nil {
		r.Ropt.Option = appendOrClip(r.Ropt.Option, reply.Options)
	}

	item.reply = reply
	if reply.Truncated {
		item.seenTruncated = true
	}
	item.stored = f.now()

	// Compute how long to cache the item.
	ttl := uint32(cacheMaxSeconds)
	// If the reply is an error, or contains no ttls, use the minimum cache time.
	if (reply.Rcode != dns.RcodeSuccess && reply.Rcode != dns.RcodeNameError) ||
		len(reply.Answer)+len(reply.Ns)+len(reply.Extra) == 0 {
		ttl = cacheMinSeconds
	}
	for _, rr := range reply.Answer {
		ttl = min(ttl, rr.Header().Ttl)
	}
	for _, rr := range reply.Ns {
		ttl = min(ttl, rr.Header().Ttl)
	}
	for _, rr := range reply.Extra {
		ttl = min(ttl, rr.Header().Ttl)
	}
	item.ttl = max(ttl, cacheMinSeconds)

	if reply.NoStore {
		f.cache.Remove(hash)
	}
}

func replyFromCache(r *netDNS.Request, item *cacheItem, now time.Time) {
	cacheLookupsCount.WithLabelValues("hit").Inc()
	decrementTtl := uint32(max(0, now.Sub(item.stored)/time.Second))

	r.Reply.Truncated = item.reply.Truncated
	r.Reply.Rcode = item.reply.Rcode

	existing_len := len(r.Reply.Answer)
	r.Reply.Answer = appendCached(r.Reply.Answer, item.reply.Answer, decrementTtl)
	shuffleAnswer(r.Reply.Answer[existing_len:])
	r.Reply.Ns = appendCached(r.Reply.Ns, item.reply.Ns, decrementTtl)
	r.Reply.Extra = appendCached(r.Reply.Extra, item.reply.Extra, decrementTtl)
	if r.Ropt != nil {
		r.Ropt.Option = appendOrClip(r.Ropt.Option, item.reply.Options)
	}
}

func appendCached(existing, add []dns.RR, decrementTtl uint32) []dns.RR {
	existing = slices.Grow(existing, len(add))
	for _, rr := range add {
		decRR := dns.Copy(rr)
		hdr := decRR.Header()
		if hdr.Ttl == 0 {
		} else if decrementTtl >= hdr.Ttl {
			// Don't decrement the TTL to 0, as that could cause problems.
			// https://00f.net/2011/11/17/how-long-does-a-dns-ttl-last/
			hdr.Ttl = 1
		} else {
			hdr.Ttl = hdr.Ttl - decrementTtl
		}
		existing = append(existing, decRR)
	}
	return existing
}

// shuffleAnswer randomizes the order of consecutive RRs which are part of the
// same RRset. This provides some load balancing.
func shuffleAnswer(rrs []dns.RR) {
	if len(rrs) < 2 {
		return
	}
	startIndex := 0
	startHdr := rrs[0].Header()
	for i := 1; i < len(rrs); i++ {
		hdr := rrs[i].Header()
		sameRRset := startHdr.Rrtype == hdr.Rrtype &&
			startHdr.Class == hdr.Class &&
			startHdr.Name == hdr.Name
		if sameRRset {
			swap := startIndex + rand.IntN(i+1-startIndex)
			rrs[i], rrs[swap] = rrs[swap], rrs[i]
		} else {
			startIndex = i
			startHdr = hdr
		}
	}
}

// appendOrClip is similar to append(a, b...) except that it avoids allocation
// if a is empty, in which case it returns b with any free capacity removed.
// The resulting slice can still be appended to without affecting b.
func appendOrClip[S ~[]E, E any](a, b S) S {
	if len(a) == 0 {
		return slices.Clip(b)
	}
	return append(a, b...)
}
