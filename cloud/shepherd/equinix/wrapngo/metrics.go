package wrapngo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"
)

// metricsSet contains all the Prometheus metrics collected by wrapngo.
type metricsSet struct {
	requestLatencies *prometheus.HistogramVec
	waiting          prometheus.GaugeFunc
	inFlight         prometheus.GaugeFunc
}

func newMetricsSet(ser *serializer) *metricsSet {
	return &metricsSet{
		requestLatencies: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "equinix_api_latency",
				Help: "Equinix API request latency in seconds, partitioned by endpoint status code",
			},
			[]string{"endpoint", "status_code"},
		),
		waiting: prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: "equinix_api_waiting",
				Help: "Number of API requests pending to be sent to Equinix but waiting on semaphore",
			},
			func() float64 {
				_, waiting := ser.stats()
				return float64(waiting)
			},
		),
		inFlight: prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: "equinix_api_in_flight",
				Help: "Number of API requests currently being processed by Equinix",
			},
			func() float64 {
				inFlight, _ := ser.stats()
				return float64(inFlight)
			},
		),
	}
}

// getEndpointForPath converts from an Equinix API method and path (eg.
// /metal/v1/devices/deadbeef) into an 'endpoint' name, which is an imaginary,
// Monogon-specific name for the API endpoint accessed by this call.
//
// If the given path is unknown and thus cannot be converted to an endpoint name,
// 'Unknown' is return and a warning is logged.
//
// We use this function to partition request statistics per API 'endpoint'. An
// alternative to this would be to record high-level packngo function names, but
// one packngo function call might actually emit multiple HTTP API requests - so
// we're stuck recording the low-level requests and gathering statistics from
// there instead.
func getEndpointForPath(method, path string) string {
	path = strings.TrimPrefix(path, "/metal/v1")
	for name, match := range endpointNames {
		if match.matches(method, path) {
			return name
		}
	}
	klog.Warningf("Unknown Equinix API %s %s - cannot determine metric endpoint name", method, path)
	return "Unknown"
}

// requestMatch is used to match a HTTP request method/path.
type requestMatch struct {
	method string
	regexp *regexp.Regexp
}

func (r *requestMatch) matches(method, path string) bool {
	if r.method != method {
		return false
	}
	return r.regexp.MatchString(path)
}

var (
	endpointNames = map[string]requestMatch{
		"GetDevice":           {"GET", regexp.MustCompile(`^/devices/[^/]+$`)},
		"ListDevices":         {"GET", regexp.MustCompile(`^/(organizations|projects)/[^/]+/devices$`)},
		"CreateDevice":        {"POST", regexp.MustCompile(`^/projects/[^/]+/devices$`)},
		"ListReservations":    {"GET", regexp.MustCompile(`^/project/[^/]+/hardware-reservations$`)},
		"ListSSHKeys":         {"GET", regexp.MustCompile(`^/ssh-keys$`)},
		"CreateSSHKey":        {"POST", regexp.MustCompile(`^/project/[^/]+/ssh-keys$`)},
		"GetSSHKey":           {"GET", regexp.MustCompile(`^/ssh-keys/[^/]+$`)},
		"UpdateSSHKey":        {"PATCH", regexp.MustCompile(`^/ssh-keys/[^/]+$`)},
		"PerformDeviceAction": {"POST", regexp.MustCompile(`^/devices/[^/]+/actions$`)},
	}
)

// onAPIRequestDone is called by the wrapngo code on every API response from
// Equinix, and records the given parameters into metrics.
func (m *metricsSet) onAPIRequestDone(req *http.Request, res *http.Response, err error, latency time.Duration) {
	if m == nil {
		return
	}

	code := "unknown"
	if err == nil {
		code = fmt.Sprintf("%d", res.StatusCode)
	} else {
		switch {
		case errors.Is(err, context.Canceled):
			code = "ctx canceled"
		case errors.Is(err, context.DeadlineExceeded):
			code = "deadline exceeded"
		}
	}
	if code == "unknown" {
		klog.Warningf("Unexpected HTTP result: req %s %s, error: %v", req.Method, req.URL.Path, res)
	}

	endpoint := getEndpointForPath(req.Method, req.URL.Path)
	m.requestLatencies.With(prometheus.Labels{"endpoint": endpoint, "status_code": code}).Observe(latency.Seconds())
}
