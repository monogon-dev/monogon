package metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"

	"source.monogon.dev/metropolis/pkg/supervisor"
)

type Discovery struct {
	Curator ipb.CuratorClient

	// sdResp will contain the cached sdResponse
	sdResp sdResponse
	// sdRespMtx is the mutex for sdResp to allow usage inside the http handler.
	sdRespMtx sync.RWMutex
}

type sdResponse []sdTarget

type sdTarget struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

// Run is the sub-runnable responsible for fetching and serving node updates.
func (s *Discovery) Run(ctx context.Context) error {
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

	defer func() {
		s.sdRespMtx.Lock()
		// disable the metrics endpoint until the new routine takes over
		s.sdResp = nil
		s.sdRespMtx.Unlock()
	}()

	nodes := make(map[string]*apb.Node)
	for {
		ev, err := srv.Recv()
		if err != nil {
			// The watcher wont return a properly wrapped error which confuses
			// our testing harness. Lets just return the context error directly
			// if it exists.
			if ctx.Err() != nil {
				return ctx.Err()
			}

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

		s.sdResp = nil
		for _, n := range nodes {
			// Only care about nodes that have all required configuration set.
			if n.Status == nil || n.Status.ExternalAddress == "" || n.Roles == nil {
				continue
			}

			s.sdResp = append(s.sdResp, sdTarget{
				Targets: []string{n.Status.ExternalAddress},
				Labels: map[string]string{
					"__meta_metropolis_role_kubernetes_worker":     fmt.Sprintf("%t", n.Roles.KubernetesWorker != nil),
					"__meta_metropolis_role_kubernetes_controller": fmt.Sprintf("%t", n.Roles.KubernetesController != nil),
					"__meta_metropolis_role_consensus_member":      fmt.Sprintf("%t", n.Roles.ConsensusMember != nil),
				},
			})
		}

		s.sdRespMtx.Unlock()
	}
}

func (s *Discovery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("method %q not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}

	s.sdRespMtx.RLock()
	defer s.sdRespMtx.RUnlock()

	// If sdResp is nil, which only happens if we are not a master node
	// or we are still booting, we respond with NotImplemented.
	if s.sdResp == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
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
