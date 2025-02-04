// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package proxy

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"github.com/prometheus/client_golang/prometheus"

	"source.monogon.dev/osbase/net/dns"
)

// Variables declared for monitoring.
var (
	healthcheckFailureCount = dns.MetricsFactory.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dnsserver",
		Subsystem: "forward",
		Name:      "healthcheck_failures_total",
		Help:      "Counter of the number of failed healthchecks.",
	}, []string{"to"})
)
