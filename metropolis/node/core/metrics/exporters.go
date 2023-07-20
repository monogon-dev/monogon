package metrics

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	"k8s.io/kubernetes/cmd/kubeadm/app/constants"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/pkg/logtree"
)

// An Exporter is a Prometheus binary running under the Metrics service which
// collects some metrics and exposes them on a locally bound TCP port.
//
// The Metrics Service will forward requests from /metrics/<name> to the
// exporter.
type Exporter struct {
	// Name of the exporter, which becomes part of the metrics URL for this exporter.
	Name string
	// Port on which this exporter will be running.
	Port node.Port
	// ServerName used to verify the tls connection.
	ServerName string
	// TLSConfigFunc is used to configure tls authentication
	TLSConfigFunc func(*Service, *Exporter) *tls.Config
	// Executable to run to start the exporter.
	Executable string
	// Arguments to start the exporter. The exporter should listen at 127.0.0.1 and
	// the port specified by Port, and serve its metrics on /metrics.
	Arguments []string
}

// DefaultExporters are the exporters which we run by default in Metropolis.
var DefaultExporters = []Exporter{
	{
		Name:       "node",
		Port:       node.MetricsNodeListenerPort,
		Executable: "/metrics/bin/node_exporter",
		Arguments: []string{
			"--web.listen-address=127.0.0.1:" + node.MetricsNodeListenerPort.PortString(),
			"--collector.buddyinfo",
			"--collector.zoneinfo",
			"--collector.tcpstat",
			"--collector.filesystem.mount-points-exclude=^/(dev|proc|sys|data/kubernetes/kubelet/pods/.+|tmp/.+|ephermal/containerd/.+)($|/)",
		},
	},
	{
		Name: "etcd",
		Port: node.MetricsEtcdListenerPort,
	},
	{
		Name:          "kubernetes-scheduler",
		Port:          constants.KubeSchedulerPort,
		ServerName:    "kube-scheduler.local",
		TLSConfigFunc: (*Service).kubeTLSConfig,
	},
	{
		Name:          "kubernetes-controller-manager",
		Port:          constants.KubeControllerManagerPort,
		ServerName:    "kube-controller-manager.local",
		TLSConfigFunc: (*Service).kubeTLSConfig,
	},
}

func (s *Service) kubeTLSConfig(e *Exporter) *tls.Config {
	c := s.KubeTLSConfig.Clone()
	c.ServerName = e.ServerName
	return c
}

// forward a given HTTP request to this exporter.
func (e *Exporter) forward(s *Service, logger logtree.LeveledLogger, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outreq := r.Clone(ctx)

	outreq.URL = &url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort("127.0.0.1", e.Port.PortString()),
		Path:   "/metrics",
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	if e.TLSConfigFunc != nil {
		outreq.URL.Scheme = "https"
		transport.TLSClientConfig = e.TLSConfigFunc(s, e)
	}
	logger.V(1).Infof("%s: forwarding %s to %s", r.RemoteAddr, r.URL.String(), outreq.URL.String())

	if r.ContentLength == 0 {
		outreq.Body = nil
	}
	if outreq.Body != nil {
		defer outreq.Body.Close()
	}
	res, err := transport.RoundTrip(outreq)
	if err != nil {
		logger.Errorf("%s: forwarding to %q failed: %v", r.RemoteAddr, e.Name, err)
		w.WriteHeader(502)
		fmt.Fprintf(w, "could not reach exporter")
		return
	}

	copyHeader(w.Header(), res.Header)
	w.WriteHeader(res.StatusCode)

	if _, err := io.Copy(w, res.Body); err != nil {
		logger.Errorf("%s: copying response from %q failed: %v", r.RemoteAddr, e.Name, err)
		return
	}
	res.Body.Close()
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (e *Exporter) externalPath() string {
	return "/metrics/" + e.Name
}
