package metrics

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"sync"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// Service is the Metropolis Metrics Service.
//
// Currently, metrics means Prometheus metrics.
//
// It runs a forwarding proxy from a public HTTPS listener to a number of
// locally-running exporters, themselves listening over HTTP. The listener uses
// the main cluster CA and the node's main certificate, authenticating incoming
// connections with the same CA.
//
// Each exporter is exposed on a separate path, /metrics/<name>, where <name> is
// the name of the exporter.
//
// The HTTPS listener is bound to node.MetricsPort.
type Service struct {
	// Credentials used to run the TLS/HTTPS listener and verify incoming
	// connections.
	Credentials *identity.NodeCredentials
	// Curator is the gRPC client that the service will use to reach the cluster's
	// Curator, for pulling a list of all nodes.
	Curator ipb.CuratorClient
	// LocalRoles contains the local node roles which gets listened on and
	// is required to decide whether or not to start the discovery routine
	LocalRoles *memory.Value[*cpb.NodeRoles]
	// List of Exporters to run and to forward HTTP requests to. If not set, defaults
	// to DefaultExporters.
	Exporters []Exporter
	// enableDynamicAddr enables listening on a dynamically chosen TCP port. This is
	// used by tests to make sure we don't fail due to the default port being already
	// in use.
	enableDynamicAddr bool

	// dynamicAddr will contain the picked dynamic listen address after the service
	// starts, if enableDynamicAddr is set.
	dynamicAddr chan string
	// sdResp will contain the cached sdResponse
	sdResp sdResponse
	// sdRespMtx is the mutex for sdResp to allow usage inside the http handler.
	sdRespMtx sync.RWMutex
}

// listen starts the public TLS listener for the service.
func (s *Service) listen() (net.Listener, error) {
	cert := s.Credentials.TLSCredentials()

	pool := x509.NewCertPool()
	pool.AddCert(s.Credentials.ClusterCA())

	tlsc := tls.Config{
		Certificates: []tls.Certificate{
			cert,
		},
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  pool,
		// TODO(q3k): use VerifyPeerCertificate/VerifyConnection to check that the
		// incoming client is allowed to access metrics. Currently we allow
		// anyone/anything with a valid cluster certificate to access them.
	}

	addr := net.JoinHostPort("", node.MetricsPort.PortString())
	if s.enableDynamicAddr {
		addr = ""
	}
	return tls.Listen("tcp", addr, &tlsc)
}

func (s *Service) Run(ctx context.Context) error {
	lis, err := s.listen()
	if err != nil {
		return fmt.Errorf("listen failed: %w", err)
	}
	if s.enableDynamicAddr {
		s.dynamicAddr <- lis.Addr().String()
	}

	if s.Exporters == nil {
		s.Exporters = DefaultExporters
	}

	// First, make sure we don't have duplicate exporters.
	seenNames := make(map[string]bool)
	for _, exporter := range s.Exporters {
		if seenNames[exporter.Name] {
			return fmt.Errorf("duplicate exporter name: %q", exporter.Name)
		}
		seenNames[exporter.Name] = true
	}

	// Start all exporters as sub-runnables.
	for _, exporter := range s.Exporters {
		cmd := exec.CommandContext(ctx, exporter.Executable, exporter.Arguments...)
		err := supervisor.Run(ctx, exporter.Name, func(ctx context.Context) error {
			return supervisor.RunCommand(ctx, cmd)
		})
		if err != nil {
			return fmt.Errorf("running %s failed: %w", exporter.Name, err)
		}

	}

	// And register all exporter forwarding functions on a mux.
	mux := http.NewServeMux()
	logger := supervisor.Logger(ctx)
	for _, exporter := range s.Exporters {
		exporter := exporter

		mux.HandleFunc(exporter.externalPath(), func(w http.ResponseWriter, r *http.Request) {
			exporter.forward(logger, w, r)
		})

		logger.Infof("Registered exporter %q", exporter.Name)
	}

	// And register a http_sd discovery endpoint.
	mux.HandleFunc("/discovery", s.handleDiscovery)

	if err := supervisor.Run(ctx, "watch-roles", s.watchRoles); err != nil {
		return err
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	// Start forwarding server.
	srv := http.Server{
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		<-ctx.Done()
		srv.Close()
	}()

	err = srv.Serve(lis)
	if err != nil && ctx.Err() != nil {
		return ctx.Err()
	}
	return fmt.Errorf("Serve: %w", err)
}

func shouldStartDiscovery(nr *cpb.NodeRoles) bool {
	return nr.ConsensusMember != nil
}

func (s *Service) watchRoles(ctx context.Context) error {
	w := s.LocalRoles.Watch()
	defer w.Close()

	r, err := w.Get(ctx)
	if err != nil {
		return err
	}

	if shouldStartDiscovery(r) {
		supervisor.Logger(ctx).Infof("Starting discovery endpoint")
		if err := supervisor.Run(ctx, "watch", s.watch); err != nil {
			return err
		}
	}

	for {
		nr, err := w.Get(ctx)
		if err != nil {
			return err
		}

		if shouldStartDiscovery(r) != shouldStartDiscovery(nr) {
			s.sdRespMtx.Lock()
			// disable the metrics endpoint until the new routine takes over
			s.sdResp = nil
			s.sdRespMtx.Unlock()

			supervisor.Logger(ctx).Infof("Discovery endpoint config changed, restarting")
			return fmt.Errorf("restarting")
		}
	}

}

// watch is the sub-runnable responsible for fetching node updates.
func (s *Service) watch(ctx context.Context) error {
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	srv, err := s.Curator.Watch(ctx, &apb.WatchRequest{
		Kind: &apb.WatchRequest_NodesInCluster_{
			NodesInCluster: &apb.WatchRequest_NodesInCluster{},
		},
	})
	if err != nil {
		return fmt.Errorf("curator watch failed: %w", err)
	}
	defer srv.CloseSend()

	nodes := make(map[string]*apb.Node)
	for {
		ev, err := srv.Recv()
		if err != nil {
			return fmt.Errorf("curator watch recv failed: %w", err)
		}

		for _, n := range ev.Nodes {
			nodes[n.Id] = n
		}

		for _, t := range ev.NodeTombstones {
			n, ok := nodes[t.NodeId]
			if !ok {
				// This is an indication of us losing data somehow. If this happens, it likely
				// means a Curator bug.
				supervisor.Logger(ctx).Warningf("Node %s: tombstone for unknown node", t.NodeId)
				continue
			}
			delete(nodes, n.Id)
		}

		s.sdRespMtx.Lock()

		// reset the existing response slice
		s.sdResp = s.sdResp[:0]
		for _, n := range nodes {
			// Only care about nodes that have all required configuration set.
			if n.Status == nil || n.Status.ExternalAddress == "" || n.Roles == nil {
				continue
			}

			s.sdResp = append(s.sdResp, sdTarget{
				Targets: []string{n.Status.ExternalAddress},
				Labels: map[string]string{
					"kubernetes_worker":     fmt.Sprintf("%t", n.Roles.KubernetesWorker != nil),
					"consensus_member":      fmt.Sprintf("%t", n.Roles.ConsensusMember != nil),
					"kubernetes_controller": fmt.Sprintf("%t", n.Roles.KubernetesController != nil),
				},
			})
		}

		s.sdRespMtx.Unlock()
	}
}

func (s *Service) handleDiscovery(w http.ResponseWriter, _ *http.Request) {
	s.sdRespMtx.RLock()
	defer s.sdRespMtx.RUnlock()

	// If sdResp is nil, which only happens if we are not a master node
	// or we are still booting, we respond with NotImplemented.
	if s.sdResp == nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(s.sdResp); err != nil {
		// If the encoder fails its mostly because of closed connections
		// so lets just ignore these errors.
		return
	}
}

type sdResponse []sdTarget

type sdTarget struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}
