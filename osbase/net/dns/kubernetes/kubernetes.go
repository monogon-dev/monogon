// Package kubernetes provides the kubernetes backend.
package kubernetes

// Taken and modified from the Kubernetes plugin of CoreDNS, under Apache 2.0.

import (
	"context"
	"net/netip"

	"github.com/miekg/dns"
	"k8s.io/client-go/kubernetes"

	"source.monogon.dev/osbase/supervisor"
)

// Kubernetes is a DNS handler that implements the Kubernetes
// DNS-Based Service Discovery specification.
// https://github.com/kubernetes/dns/blob/master/docs/specification.md
type Kubernetes struct {
	clusterDomain string
	nsDomain      string
	ipRanges      []netip.Prefix
	// A Kubernetes ClientSet with read access to endpoints and services
	ClientSet kubernetes.Interface
	apiConn   dnsController
}

// New returns an initialized Kubernetes. Kubernetes DNS records will be served
// under the clusterDomain. Additionally, reverse queries for services and pods
// are served under the given ipRanges.
func New(clusterDomain string, ipRanges []netip.Prefix) *Kubernetes {
	k := new(Kubernetes)
	k.clusterDomain = dns.CanonicalName(clusterDomain)
	k.nsDomain = "ns.dns." + k.clusterDomain
	k.ipRanges = ipRanges
	return k
}

// Run maintains the in-memory cache of Kubernetes services and endpoints.
func (k *Kubernetes) Run(ctx context.Context) error {
	k.apiConn = newdnsController(ctx, k.ClientSet)
	k.apiConn.Start(ctx.Done())

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return ctx.Err()
}
