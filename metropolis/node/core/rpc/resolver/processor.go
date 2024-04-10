package resolver

import (
	"context"
	"fmt"
	"net"
	"sort"

	"google.golang.org/grpc"

	cpb "source.monogon.dev/metropolis/proto/common"
)

// request contains all possible requests passed to the processor. Only one
// field can be set at a time. See the documentation of member structures for
// more information about the possible requests.
type request struct {
	cmg   *requestCuratorMapGet
	nu    *requestNodesUpdate
	sa    *requestSeedAdd
	oa    *requestOverrideAdd
	lu    *requestLeaderUpdate
	ds    *requestDialOptionsSet
	sub   *requestSubscribe
	unsub *requestUnsubscribe
	dbg   *requestDebug
}

// requestCuratorMapGet is received from any subsystem which wants to
// immediately receive the current curatorMap as seen by the processor.
type requestCuratorMapGet struct {
	// resC carries the current curatorMap. It must be drained by the caller,
	// otherwise the processor will get stuck.
	resC chan *curatorMap
}

// requestNodesUpdate is received from the curator updater, and carries
// information about the current curators as seen by the cluster control plane.
type requestNodesUpdate struct {
	// nodes is a map from node ID to received status
	nodes map[string]*cpb.NodeStatus
}

// requestSeedAdd is received from AddEndpoint calls. It updates the processor's
// curator map with the given seed.
type requestSeedAdd struct {
	endpoint *NodeEndpoint
}

// requestOverrideAdd is received from AddOverride calls. It updates the
// processor's curator map with the given override.
type requestOverrideAdd struct {
	nodeID   string
	endpoint *NodeEndpoint
}

// requestLeaderUpdate is received from the leader watcher whenever a new leader
// is found from any curator.
type requestLeaderUpdate struct {
	nodeID   string
	endpoint *NodeEndpoint
}

// requestDialOptionsSet is received from any subordinate watchers when a client
// connects with the given dial options. The processor will use the first
// options received this way to establish connectivity to curators.
type requestDialOptionsSet struct {
	options []grpc.DialOption
}

// requestSubscribe is received from subordinate watchers. The processor will
// then create a subscription channel that will be populated with updates about
// the current leader.
type requestSubscribe struct {
	resC chan *responseSubscribe
}

// requestDebug is received from the curator updater on failure, and is used to
// provide the user with up-to-date debug information about the processor's state
// at time of failure.
type requestDebug struct {
	resC chan *responseDebug
}

type responseDebug struct {
	// curmap is the copy of the curator map as seen by the processor at time of
	// request.
	curmap *curatorMap
	// leader is the current leader, if any, as seen by the processor at the time of
	// the request.
	leader *requestLeaderUpdate
}

type responseSubscribe struct {
	// id is the ID of the subscription, used to cancel the subscription by the
	// subscriber.
	id int64
	// subC carries updates about the current leader. The subscriber must drain the
	// updates as fast as possible, otherwise the processing loop will be stopped.
	subC chan *update
}

type update struct {
	// node ID of the current leader.
	nodeID string
	// endpoint of the current leader.
	endpoint *NodeEndpoint
}

// requestUnsubscribe is received from subordinate watcher to cancel a given
// subscription.
type requestUnsubscribe struct {
	id int64
}

// run the resolver's processor, which is the main processing loop. It received
// updates from users, watchers, the curator updater and the leader updater.
func (r *Resolver) run(ctx context.Context) error {
	// Current curator map.
	curMap := newCuratorMap()

	// Current leader.
	var leader *requestLeaderUpdate

	// Subscribers.
	subscribers := make(map[int64]chan *update)
	subscriberIDNext := int64(0)

	// Whether the curator updater and leader updater have been started. This is
	// only done once we receive dial options from a watcher.
	running := false

	for {
		// Receive a request. Quit if our context gets canceled in the meantime.
		var req *request
		select {
		case <-ctx.Done():
			// Close all subscription channels, ensuring all the watchers get notified that
			// the resolver has closed.
			for _, subC := range subscribers {
				close(subC)
			}
			return ctx.Err()
		case req = <-r.reqC:
		}

		// Dispatch request.
		switch {
		case req.cmg != nil:
			// Curator Map Get
			req.cmg.resC <- curMap.copy()
		case req.nu != nil:
			// Nodes Update
			for nid, status := range req.nu.nodes {
				// Skip nodes which aren't running the curator right now.
				if status == nil || status.RunningCurator == nil {
					continue
				}
				addr := net.JoinHostPort(status.ExternalAddress, fmt.Sprintf("%d", status.RunningCurator.Port))
				if a, ok := curMap.overrides[nid]; ok {
					addr = a.endpoint
				}

				curMap.curators[nid] = &NodeEndpoint{
					endpoint: addr,
				}
			}
			toDelete := make(map[string]bool)
			for nid := range curMap.curators {
				if req.nu.nodes[nid] == nil {
					toDelete[nid] = true
				}
			}
			for nid := range toDelete {
				delete(curMap.curators, nid)
			}
		case req.sa != nil:
			// Seed Add
			curMap.seeds[req.sa.endpoint.endpoint] = true
		case req.oa != nil:
			// Override Add
			curMap.overrides[req.oa.nodeID] = req.oa.endpoint
		case req.lu != nil:
			// Leader Update
			leader = req.lu
			for _, s := range subscribers {
				s <- &update{
					nodeID:   leader.nodeID,
					endpoint: leader.endpoint,
				}
			}
		case req.ds != nil:
			// Dial options Set
			if !running {
				if !r.noCuratorUpdater {
					go r.runCuratorUpdater(ctx, req.ds.options)
				}
				go r.runLeaderUpdater(ctx, req.ds.options)
			}
			running = true
		case req.sub != nil:
			// Subscribe
			id := subscriberIDNext
			subC := make(chan *update)
			req.sub.resC <- &responseSubscribe{
				id:   id,
				subC: subC,
			}
			subscriberIDNext++
			subscribers[id] = subC

			// Provide current leader if missing.
			if leader != nil {
				subC <- &update{
					nodeID:   leader.nodeID,
					endpoint: leader.endpoint,
				}
			}
		case req.unsub != nil:
			// Unsubscribe
			if subscribers[req.unsub.id] != nil {
				close(subscribers[req.unsub.id])
				delete(subscribers, req.unsub.id)
			}
		case req.dbg != nil:
			// Debug
			var leader2 *requestLeaderUpdate
			if leader != nil {
				endpoint := *leader.endpoint
				leader2 = &requestLeaderUpdate{
					nodeID:   leader.nodeID,
					endpoint: &endpoint,
				}
			}
			req.dbg.resC <- &responseDebug{
				curmap: curMap.copy(),
				leader: leader2,
			}
		default:
			panic(fmt.Sprintf("unhandled request: %+v", req))
		}
	}
}

// curatorMap is the main state of the cluster as seen by the resolver's processor.
type curatorMap struct {
	// curators is a map from node ID to endpoint of nodes that are currently
	// running the curator. This is updated by the processor through the curator
	// updater.
	curators map[string]*NodeEndpoint
	// overrides are user-provided node ID to endpoint overrides. They are applied
	// to the curators in the curator map (above) and the leader information as
	// retrieved by the leader updater.
	overrides map[string]*NodeEndpoint
	// seeds are endpoints provided by the user that the leader updater will use to
	// bootstrap initial cluster connectivity.
	seeds map[string]bool
}

func newCuratorMap() *curatorMap {
	return &curatorMap{
		curators:  make(map[string]*NodeEndpoint),
		overrides: make(map[string]*NodeEndpoint),
		seeds:     make(map[string]bool),
	}
}

func (m *curatorMap) copy() *curatorMap {
	res := newCuratorMap()
	for k, v := range m.curators {
		res.curators[k] = v
	}
	for k, v := range m.overrides {
		res.overrides[k] = v
	}
	for k, v := range m.seeds {
		res.seeds[k] = v
	}
	return res
}

// candidates returns the curator endpoints that should be used to attempt to
// retrieve the current leader from. This is a combination of the curators
// received by the curator updater, and the seeds provided by the user.
func (m *curatorMap) candidates() []string {
	resMap := make(map[string]bool)
	for ep := range m.seeds {
		resMap[ep] = true
	}
	for nid, v := range m.curators {
		if o, ok := m.overrides[nid]; ok {
			resMap[o.endpoint] = true
		} else {
			resMap[v.endpoint] = true
		}
	}
	var res []string
	for ep := range resMap {
		res = append(res, ep)
	}
	sort.Strings(res)
	return res
}

// curatorMap returns the current curator map as seen by the resolver processor.
func (r *Resolver) curatorMap() *curatorMap {
	req := &request{
		cmg: &requestCuratorMapGet{
			resC: make(chan *curatorMap),
		},
	}
	r.reqC <- req
	return <-req.cmg.resC
}
