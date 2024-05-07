package metrics

import (
	"fmt"
	"io"
	"net/http"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/osbase/supervisor"
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
	// Executable to run to start the exporter.
	Executable string
	// Arguments to start the exporter. The exporter should listen at 127.0.0.1 and
	// the port specified by Port, and serve its metrics on /metrics.
	Arguments []string
	// Path to scrape metrics at. Defaults to /metrics.
	Path string
}

// DefaultExporters are the exporters which we run by default in Metropolis.
var DefaultExporters = []*Exporter{
	{
		Name:       "node",
		Port:       node.MetricsNodeListenerPort,
		Executable: "/metrics/bin/node_exporter",
		Arguments: []string{
			"--web.listen-address=127.0.0.1:" + node.MetricsNodeListenerPort.PortString(),
			"--collector.buddyinfo",
			"--collector.zoneinfo",
			"--collector.tcpstat",
			"--collector.cpu.info",
			"--collector.ethtool",
			"--collector.cpu_vulnerabilities",
			"--collector.ethtool.device-exclude=^(veth.*|sit.*|lo|clusternet)$",
			"--collector.netclass.ignored-devices=^(veth.*)$",
			"--collector.netdev.device-exclude=^(veth.*)$",
			"--collector.filesystem.mount-points-exclude=^/(dev|proc|sys|data/kubernetes/kubelet/pods/.+|tmp/.+|ephemeral/containerd/.+)($|/)",
		},
	},
	{
		Name: "etcd",
		Port: node.MetricsEtcdListenerPort,
	},
	{
		Name: "kubernetes-scheduler",
		Port: node.MetricsKubeSchedulerListenerPort,
	},
	{
		Name: "kubernetes-controller-manager",
		Port: node.MetricsKubeControllerManagerListenerPort,
	},
	{
		Name: "kubernetes-apiserver",
		Port: node.MetricsKubeAPIServerListenerPort,
	},
	{
		Name: "containerd",
		Port: node.MetricsContainerdListenerPort,
		Path: "/v1/metrics",
	},
}

func (e *Exporter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("method %q not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	// We are supplying the http.Server with a BaseContext that contains the
	// context from our runnable which contains the logger.
	logger := supervisor.Logger(ctx)

	path := e.Path
	if e.Path == "" {
		path = "/metrics"
	}

	url := "http://127.0.0.1:" + e.Port.PortString() + path
	outReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		logger.Errorf("%s: forwarding to %q failed: %v", r.RemoteAddr, e.Name, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultTransport.RoundTrip(outReq)
	if err != nil {
		logger.Errorf("%s: forwarding to %q failed: %v", r.RemoteAddr, e.Name, err)
		http.Error(w, "could not reach exporter", http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	copyHeader(w.Header(), res.Header)
	w.WriteHeader(res.StatusCode)

	if _, err := io.Copy(w, res.Body); err != nil {
		logger.Errorf("%s: copying response from %q failed: %v", r.RemoteAddr, e.Name, err)
		return
	}
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
