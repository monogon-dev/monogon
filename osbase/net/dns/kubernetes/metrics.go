// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package kubernetes

// Taken and modified from the Kubernetes plugin of CoreDNS, under Apache 2.0.

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"source.monogon.dev/osbase/net/dns"
)

var (
	// dnsProgrammingLatency is defined as the time it took to program a DNS
	// instance - from the time a service or pod has changed to the time the
	// change was propagated and was available to be served by a DNS server.
	// The definition of this SLI can be found at https://github.com/kubernetes/community/blob/master/sig-scalability/slos/dns_programming_latency.md
	// Note that the metrics is partially based on the time exported by the
	// endpoints controller on the master machine. The measurement may be
	// inaccurate if there is a clock drift between the node and master machine.
	// The service_kind label can be one of:
	//   * cluster_ip
	//   * headless_with_selector
	//   * headless_without_selector
	dnsProgrammingLatency = dns.MetricsFactory.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "dnsserver",
		Subsystem: "kubernetes",
		Name:      "dns_programming_duration_seconds",
		// From 1 millisecond to ~17 minutes.
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 20),
		Help:    "In Cluster DNS Programming Latency in seconds",
	}, []string{"service_kind"})
)

func recordDNSProgrammingLatency(lastChangeTriggerTime time.Time) {
	if !lastChangeTriggerTime.IsZero() {
		// If we're here it means that the Endpoints object is for a headless service
		// and that the Endpoints object was created by the endpoints-controller
		// (because the LastChangeTriggerTime annotation is set). It means that the
		// corresponding service is a "headless service with selector".
		dnsProgrammingLatency.WithLabelValues("headless_with_selector").
			Observe(time.Since(lastChangeTriggerTime).Seconds())
	}
}
