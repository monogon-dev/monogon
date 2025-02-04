// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package forward

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"github.com/prometheus/client_golang/prometheus"

	"source.monogon.dev/osbase/net/dns"
)

// Variables declared for monitoring.
var (
	// Possible results:
	//   * hit: Item found and returned from cache.
	//   * miss: Item not found in cache.
	//   * refresh: Item found in cache, but is either expired, or
	//     truncated while the client used TCP.
	cacheLookupsCount = dns.MetricsFactory.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dnsserver",
		Subsystem: "forward",
		Name:      "cache_lookups_total",
		Help:      "Counter of the number of cache lookups.",
	}, []string{"result"})

	// protocol is one of:
	//   * udp
	//   * udp_truncated
	//   * tcp
	// rcode can be an uppercase rcode name, a numeric rcode if the rcode is not
	// known, or one of:
	//   * timeout
	//   * network_error
	//   * protocol_error
	upstreamDuration = dns.MetricsFactory.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "dnsserver",
		Subsystem: "forward",
		Name:      "upstream_duration_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.00025, 2, 16), // from 0.25ms to 8 seconds
		Help:      "Histogram of the time each upstream request took.",
	}, []string{"to", "protocol", "rcode"})

	// Possible reasons:
	//   * concurrency_limit: Too many concurrent upstream queries.
	//   * no_upstreams: There are no upstreams configured.
	//   * no_recursion_desired: Client did not set Recursion Desired flag.
	rejectsCount = dns.MetricsFactory.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dnsserver",
		Subsystem: "forward",
		Name:      "rejects_total",
		Help:      "Counter of the number of queries rejected and not forwarded to an upstream.",
	}, []string{"reason"})

	healthcheckBrokenCount = dns.MetricsFactory.NewCounter(prometheus.CounterOpts{
		Namespace: "dnsserver",
		Subsystem: "forward",
		Name:      "healthcheck_broken_total",
		Help:      "Counter of the number of complete failures of the healthchecks.",
	})
)
