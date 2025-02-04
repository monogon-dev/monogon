// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package kubernetes

import (
	"math/rand/v2"
	"net"
	"net/netip"

	"github.com/miekg/dns"

	netDNS "source.monogon.dev/osbase/net/dns"
	"source.monogon.dev/osbase/net/dns/kubernetes/object"
)

const (
	// DNSSchemaVersion is the schema version: https://github.com/kubernetes/dns/blob/master/docs/specification.md
	DNSSchemaVersion = "1.1.0"
	// defaultTTL to apply to all answers.
	defaultTTL = 5
)

func (k *Kubernetes) HandleDNS(r *netDNS.Request) {
	if netDNS.IsSubDomain(k.clusterDomain, r.QnameCanonical) {
		r.SetAuthoritative()

		subdomain := r.QnameCanonical[:len(r.QnameCanonical)-len(k.clusterDomain)]
		subdomain, last := netDNS.SplitLastLabel(subdomain)
		if last == "svc" {
			k.handleService(r, subdomain)
		} else if last == "" {
			if r.Qtype == dns.TypeSOA || r.Qtype == dns.TypeANY {
				r.Reply.Answer = append(r.Reply.Answer, k.makeSOA(r.Qname))
			}
			if r.Qtype == dns.TypeNS || r.Qtype == dns.TypeANY {
				r.Reply.Answer = append(r.Reply.Answer, k.makeNS(r.Qname))
			}
		} else if last == "dns-version" && subdomain == "" {
			if r.Qtype == dns.TypeTXT || r.Qtype == dns.TypeANY {
				rr := new(dns.TXT)
				rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: defaultTTL}
				rr.Txt = []string{DNSSchemaVersion}
				r.Reply.Answer = append(r.Reply.Answer, rr)
			}
		} else if last == "dns" && (subdomain == "" || subdomain == "ns.") {
			// Name exists but has no records.
		} else {
			r.Reply.Rcode = dns.RcodeNameError
		}

		if r.Handled {
			return
		}
		if len(r.Reply.Answer) == 0 {
			zone := r.Qname[len(r.Qname)-len(k.clusterDomain):]
			r.Reply.Ns = []dns.RR{k.makeSOA(zone)}
		}
		r.SendReply()
		return
	}

	reverseIP, reverseBits, extra := netDNS.ParseReverse(r.QnameCanonical)
	if reverseIP.IsValid() {
		for _, ipRange := range k.ipRanges {
			if !ipRange.Contains(reverseIP) || reverseBits < ipRange.Bits() {
				continue
			}

			r.SetAuthoritative()

			zoneBits := 0
			if reverseIP.BitLen() == 32 {
				zoneBits = (ipRange.Bits() + 7) & ^7
			} else {
				zoneBits = (ipRange.Bits() + 3) & ^3
			}

			if extra {
				// Name with extra labels does not exist.
				r.Reply.Rcode = dns.RcodeNameError
			} else {
				if reverseBits == reverseIP.BitLen() {
					k.handleReverse(r, reverseIP)
				}
				if reverseBits == zoneBits {
					if r.Qtype == dns.TypeSOA || r.Qtype == dns.TypeANY {
						r.Reply.Answer = append(r.Reply.Answer, k.makeSOA(r.Qname))
					}
					if r.Qtype == dns.TypeNS || r.Qtype == dns.TypeANY {
						r.Reply.Answer = append(r.Reply.Answer, k.makeNS(r.Qname))
					}
				}
			}

			if len(r.Reply.Answer) == 0 {
				zoneDots := 0
				if reverseIP.BitLen() == 32 {
					zoneDots = 3 + zoneBits/8
				} else {
					zoneDots = 3 + zoneBits/4
				}
				zoneStart := len(r.Qname)
				for zoneStart > 0 {
					if r.Qname[zoneStart-1] == '.' {
						zoneDots--
						if zoneDots == 0 {
							break
						}
					}
					zoneStart--
				}
				zone := r.Qname[zoneStart:]
				r.Reply.Ns = []dns.RR{k.makeSOA(zone)}
			}
			r.SendReply()
			return
		}
	}
}

func (k *Kubernetes) handleService(r *netDNS.Request, subdomain string) {
	if subdomain == "" {
		// Name exists but has no records.
		return
	}

	rest, namespace := netDNS.SplitLastLabel(subdomain)
	if rest == "" {
		// Name exists if the namespace exists, and has no records.
		if !k.apiConn.NamespaceExists(namespace) {
			k.notFound(r)
		}
		return
	}

	serviceSub, _ := netDNS.SplitLastLabel(rest)
	rest, hostnameOrProto := netDNS.SplitLastLabel(serviceSub)

	var proto string
	var portName string
	var hostname string
	switch hostnameOrProto {
	case "_tcp", "_udp", "_sctp":
		proto = hostnameOrProto[1:]
		rest, portName = netDNS.SplitLastLabel(rest)
		if len(portName) >= 2 && portName[0] == '_' {
			portName = portName[1:]
		} else if portName != "" {
			r.Reply.Rcode = dns.RcodeNameError
			return
		}
		// If portName is empty, the name exists if the parent exists,
		// but has no records.
	default:
		hostname = hostnameOrProto
	}

	if rest != "" {
		// The query name has too many labels.
		r.Reply.Rcode = dns.RcodeNameError
		return
	}

	// serviceKey is "<service>.<ns>"
	serviceKey := subdomain[len(serviceSub) : len(subdomain)-1]
	service := k.apiConn.GetSvc(serviceKey)
	if service == nil {
		k.notFound(r)
		return
	}

	// External service
	if service.ExternalName != "" {
		if serviceSub != "" {
			// External services don't have subdomains.
			r.Reply.Rcode = dns.RcodeNameError
			return
		}
		if service.ExternalName == object.ExternalNameInvalid {
			// The service has an invalid ExternalName, return an error.
			r.AddExtendedError(dns.ExtendedErrorCodeInvalidData, "Kubernetes service has invalid externalName")
			r.Reply.Rcode = dns.RcodeServerFailure
			return
		}
		// We already ensure that ExternalName is valid and fully qualified
		// when constructing the object.Service.
		r.AddCNAME(service.ExternalName, defaultTTL)
		return
	}

	// Headless service.
	if service.Headless {
		found := false
		haveIP := make(map[string]struct{})
		haveSRV := make(map[srvItem]struct{})
		existingAnswer := len(r.Reply.Answer)
		existingExtra := len(r.Reply.Extra)
		for _, ep := range k.apiConn.EpIndex(serviceKey) {
			if portName != "" {
				// _<port>._<proto>.<service>.<ns>.svc.
				var portNumber uint16
				for _, p := range ep.Ports {
					if p.Name == portName && p.Protocol == proto {
						portNumber = p.Port
						break
					}
				}
				if portNumber == 0 {
					continue
				}
				for _, addr := range ep.Addresses {
					found = true
					if r.Qtype == dns.TypeSRV || r.Qtype == dns.TypeANY {
						targetName := addr.Hostname + r.Qname[len(serviceSub)-1:]
						if !isDuplicateSRV(haveSRV, addr.Hostname, "", portNumber) {
							addSRV(r, portNumber, targetName)
						}
						if !isDuplicateSRV(haveSRV, addr.Hostname, addr.IP, 0) {
							addAddrExtra(r, targetName, net.IP(addr.IP))
						}
					}
				}
			} else {
				// <service>.<ns>.svc. or <hostname>.<service>.<ns>.svc.
				for _, addr := range ep.Addresses {
					if hostname != "" && hostname != addr.Hostname {
						continue
					}
					found = true
					if proto != "" {
						// _<proto>.<service>.<ns>.svc. has no records
						// and exists if its parent exists.
						break
					}
					if _, ok := haveIP[addr.IP]; !ok {
						haveIP[addr.IP] = struct{}{}
						addAddr(r, net.IP(addr.IP))
					}
				}
			}
		}
		shuffleRRs(r.Reply.Answer[existingAnswer:])
		shuffleRRs(r.Reply.Extra[existingExtra:])
		if !found {
			k.notFound(r)
		}
		return
	}

	if hostname != "" {
		// Non-headless services don't have hostname records.
		r.Reply.Rcode = dns.RcodeNameError
		return
	}

	// ClusterIP service
	if proto == "" {
		// <service>.<ns>.svc. for ClusterIP service.
		for _, ip := range service.ClusterIPs {
			addAddr(r, net.IP(ip))
		}
		// The specification does not define what to return if the service has
		// no (valid) clusterIP. We return an empty response with no error.
		return
	}

	if portName == "" {
		// _<proto>.<service>.<ns>.svc. exists but has no records.
		return
	}

	// _<port>._<proto>.<service>.<ns>.svc. for ClusterIP service.
	var portNumber uint16
	for _, p := range service.Ports {
		if p.Name == portName && p.Protocol == proto {
			portNumber = p.Port
			break
		}
	}
	if portNumber == 0 {
		r.Reply.Rcode = dns.RcodeNameError
		return
	}
	if r.Qtype == dns.TypeSRV || r.Qtype == dns.TypeANY {
		targetName := r.Qname[len(serviceSub):]
		addSRV(r, portNumber, targetName)
		for _, ip := range service.ClusterIPs {
			addAddrExtra(r, targetName, net.IP(ip))
		}
	}
}

func (k *Kubernetes) handleReverse(r *netDNS.Request, ip netip.Addr) {
	stringIP := string(ip.AsSlice())
	found := false
	for _, service := range k.apiConn.SvcIndexReverse(stringIP) {
		found = true
		if r.Qtype == dns.TypePTR || r.Qtype == dns.TypeANY {
			rr := new(dns.PTR)
			rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypePTR, Class: dns.ClassINET, Ttl: defaultTTL}
			rr.Ptr = service.Name + "." + service.Namespace + ".svc." + k.clusterDomain
			r.Reply.Answer = append(r.Reply.Answer, rr)
		}
	}
	haveName := make(map[string]struct{})
	for _, ep := range k.apiConn.EpIndexReverse(stringIP) {
		for _, addr := range ep.Addresses {
			if addr.IP == stringIP {
				found = true
				if r.Qtype == dns.TypePTR || r.Qtype == dns.TypeANY {
					ptr := addr.Hostname + "." + ep.Index + ".svc." + k.clusterDomain
					if _, ok := haveName[ptr]; ok {
						continue
					}
					haveName[ptr] = struct{}{}
					rr := new(dns.PTR)
					rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypePTR, Class: dns.ClassINET, Ttl: defaultTTL}
					rr.Ptr = ptr
					r.Reply.Answer = append(r.Reply.Answer, rr)
				}
			}
		}
	}
	if !found {
		k.notFound(r)
	}
}

func (k *Kubernetes) makeSOA(zone string) *dns.SOA {
	header := dns.RR_Header{Name: zone, Rrtype: dns.TypeSOA, Class: dns.ClassINET, Ttl: defaultTTL}
	return &dns.SOA{
		Hdr:     header,
		Mbox:    "nobody.invalid.",
		Ns:      k.nsDomain,
		Serial:  uint32(k.apiConn.Modified()),
		Refresh: 7200,
		Retry:   1800,
		Expire:  86400,
		Minttl:  defaultTTL,
	}
}

func (k *Kubernetes) makeNS(zone string) *dns.NS {
	rr := new(dns.NS)
	rr.Hdr = dns.RR_Header{Name: zone, Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: defaultTTL}
	rr.Ns = k.nsDomain
	return rr
}

func addAddr(r *netDNS.Request, ip net.IP) {
	if len(ip) == net.IPv4len && (r.Qtype == dns.TypeA || r.Qtype == dns.TypeANY) {
		rr := new(dns.A)
		rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: defaultTTL}
		rr.A = ip
		r.Reply.Answer = append(r.Reply.Answer, rr)
	}
	if len(ip) == net.IPv6len && (r.Qtype == dns.TypeAAAA || r.Qtype == dns.TypeANY) {
		rr := new(dns.AAAA)
		rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: defaultTTL}
		rr.AAAA = ip
		r.Reply.Answer = append(r.Reply.Answer, rr)
	}
}

func addAddrExtra(r *netDNS.Request, name string, ip net.IP) {
	if len(ip) == net.IPv4len {
		rr := new(dns.A)
		rr.Hdr = dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: defaultTTL}
		rr.A = ip
		r.Reply.Extra = append(r.Reply.Extra, rr)
	}
	if len(ip) == net.IPv6len {
		rr := new(dns.AAAA)
		rr.Hdr = dns.RR_Header{Name: name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: defaultTTL}
		rr.AAAA = ip
		r.Reply.Extra = append(r.Reply.Extra, rr)
	}
}

func addSRV(r *netDNS.Request, portNumber uint16, targetName string) {
	rr := new(dns.SRV)
	rr.Hdr = dns.RR_Header{Name: r.Qname, Rrtype: dns.TypeSRV, Class: dns.ClassINET, Ttl: defaultTTL}
	rr.Priority = 0
	rr.Weight = 0
	rr.Port = portNumber
	rr.Target = targetName
	r.Reply.Answer = append(r.Reply.Answer, rr)
}

// notFound should be called if a name was not found, but could exist
// if there are Kubernetes object that are not yet available locally.
func (k *Kubernetes) notFound(r *netDNS.Request) {
	if !k.apiConn.HasSynced() {
		// We don't know if the name exists or not, so return an error.
		r.AddExtendedError(dns.ExtendedErrorCodeNotReady, "Kubernetes objects not yet synced")
		r.Reply.Rcode = dns.RcodeServerFailure
	} else {
		r.Reply.Rcode = dns.RcodeNameError
	}
}

type srvItem struct {
	name string
	addr string
	port uint16
}

// isDuplicateSRV returns true if the (name, addr, port) combination already
// exists in m, and adds it to m if not.
func isDuplicateSRV(m map[srvItem]struct{}, name, addr string, port uint16) bool {
	_, ok := m[srvItem{name, addr, port}]
	if !ok {
		m[srvItem{name, addr, port}] = struct{}{}
	}
	return ok
}

// shuffleRRs shuffles a slice of RRs for some load balancing.
func shuffleRRs(rrs []dns.RR) {
	rand.Shuffle(len(rrs), func(i, j int) {
		rrs[i], rrs[j] = rrs[j], rrs[i]
	})
}
