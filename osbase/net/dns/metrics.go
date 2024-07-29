package dns

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// MetricsRegistry is the metrics registry in which all DNS metrics are
// registered.
var MetricsRegistry = prometheus.NewRegistry()
var MetricsFactory = promauto.With(MetricsRegistry)

var (
	// rcode can be an uppercase rcode name, a numeric rcode if the rcode is not
	// known, or one of:
	//   * redirected: The query was redirected by CNAME, so the final rcode
	//     is not yet known.
	//   * not_ready: The handler is not yet ready, SERVFAIL is replied.
	handlerDuration = MetricsFactory.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "dnsserver",
		Subsystem: "server",
		Name:      "handler_duration_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.00025, 2, 16), // from 0.25ms to 8 seconds
		Help:      "Histogram of the time each handler took.",
	}, []string{"handler", "rcode"})
)
