package metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"source.monogon.dev/go/types/mapsets"
	"source.monogon.dev/metropolis/node/core/curator/watcher"
	"source.monogon.dev/metropolis/pkg/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

type Discovery struct {
	Curator ipb.CuratorClient

	// sdResp will contain the cached sdResponse
	sdResp mapsets.OrderedMap[string, sdTarget]
	// sdRespMtx is the mutex for sdResp to allow usage inside the http handler.
	sdRespMtx sync.RWMutex
}

type sdTarget struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

// Run is the sub-runnable responsible for fetching and serving node updates.
func (s *Discovery) Run(ctx context.Context) error {
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	defer func() {
		s.sdRespMtx.Lock()
		// disable the metrics endpoint until the new routine takes over
		s.sdResp.Clear()
		s.sdRespMtx.Unlock()
	}()

	return watcher.WatchNodes(ctx, s.Curator, watcher.SimpleFollower{
		FilterFn: func(a *ipb.Node) bool {
			if a.Status == nil {
				return false
			}
			if a.Status.ExternalAddress == "" {
				return false
			}
			if a.Roles == nil {
				return false
			}
			return true
		},
		EqualsFn: func(a *ipb.Node, b *ipb.Node) bool {
			if (a.Roles.ConsensusMember == nil) != (b.Roles.ConsensusMember == nil) {
				return false
			}
			if (a.Roles.KubernetesController == nil) != (b.Roles.KubernetesController == nil) {
				return false
			}
			if (a.Roles.ConsensusMember == nil) != (b.Roles.ConsensusMember == nil) {
				return false
			}
			if a.Status.ExternalAddress != b.Status.ExternalAddress {
				return false
			}
			return true
		},
		OnNewUpdated: func(new *ipb.Node) error {
			s.sdRespMtx.Lock()
			defer s.sdRespMtx.Unlock()

			s.sdResp.Insert(new.Id, sdTarget{
				Targets: []string{new.Status.ExternalAddress},
				Labels: map[string]string{
					"__meta_metropolis_role_kubernetes_worker":     fmt.Sprintf("%t", new.Roles.KubernetesWorker != nil),
					"__meta_metropolis_role_kubernetes_controller": fmt.Sprintf("%t", new.Roles.KubernetesController != nil),
					"__meta_metropolis_role_consensus_member":      fmt.Sprintf("%t", new.Roles.ConsensusMember != nil),
				},
			})
			return nil
		},
		OnDeleted: func(prev *ipb.Node) error {
			s.sdRespMtx.Lock()
			defer s.sdRespMtx.Unlock()

			s.sdResp.Delete(prev.Id)
			return nil
		},
	})
}

func (s *Discovery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("method %q not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}

	s.sdRespMtx.RLock()
	defer s.sdRespMtx.RUnlock()

	// If sdResp is empty, respond with Service Unavailable. This will only happen
	// early enough in the lifecycle of a control plane node that it doesn't know
	// about itself, or if this is not a control plane node:
	if s.sdResp.Count() == 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Turn into a plain array as expected by the service discovery API.
	var res []sdTarget
	for _, v := range s.sdResp.Values() {
		res = append(res, v.Value)
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		// If the encoder fails its mostly because of closed connections
		// so lets just ignore these errors.
		return
	}
}
