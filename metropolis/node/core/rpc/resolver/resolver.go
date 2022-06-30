package resolver

import (
	"context"
	"fmt"
	"io"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/grpc"

	common "source.monogon.dev/metropolis/node"
	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

const (
	// MetropolisControlAddress is the address of the current Metropolis leader as
	// accepted by the Resolver. Dialing a gRPC channel to this address while the
	// Resolver is used will open the channel to the current leader of the
	// Metropolis control plane.
	MetropolisControlAddress = "metropolis:///control"
)

// Resolver is a gRPC resolver Builder that can be passed to
// grpc.WithResolvers() when dialing a gRPC endpoint.
//
// It's responsible for resolving the magic MetropolisControlAddress
// (metropolis:///control) into an address of the node that is currently the
// leader of the cluster's control plane.
//
// To function, the ClusterResolver needs to be provided with at least one
// control plane node address. It will use these addresses to retrieve the
// address of the node which is the current leader of the control plane.
//
// Then, having established communication with the leader, it will continuously
// update an internal set of control plane node endpoints (the curator map) that
// will be contacted in the future about the state of the leadership when the
// current leader fails over.
//
// The resolver will wait for a first gRPC connection established through it to
// extract the transport credentials used, then use these credentials to call
// the Curator and CuratorLocal services on control plane nodes to perform its
// logic.
//
// This resolver is designed to be used as a long-running object which multiple
// gRPC client connections can use. Usually one ClusterResolver instance should
// be used per application.
//
// .------------------------.        .--------------------------------------.
// | Metropolis Cluster     |        | Resolver                             |
// :------------------------:        :--------------------------------------:
// :                        :        :                                      :
// : .--------------------. :        :   .----------------.                 :
// : | curator (follower) |<---.---------| Leader Updater |------------.    :
// : '--------------------' :  |     :   '----------------'            |    :
// : .--------------------. :  |     :   .------------------------.    |    :
// : | curator (follower) |<---:     :   | Processor (CuratorMap) |<-.-'-.  :
// : '--------------------' :  |     :   '------------------------'  |   |  :
// : .--------------------.<---'     :   .-----------------.         |   |  :
// : | curator (leader)   |<-------------| Curator Updater |---------'   |  :
// : '--------------------' :        :   '-----------------'             |  :
// :                        :        :                                   |  :
// '------------------------'        :   .----------.                    |  :
//                                   :   | Watchers |-.                  |  :
//                                   :   '----------' |------------------'  :
//                                   :     '-^--------'                     :
//                                   :       |  ^                           :
//                                   :       |  |                           :
//                                        .---------------.
//                                        | gRPC channels |
//                                        '---------------'
type Resolver struct {
	reqC chan *request
	ctx  context.Context

	// logger, if set, will be called with fmt.Sprintf-like arguments containing
	// debug logs from the running ClusterResolver, subordinate watchers and
	// updaters.
	logger func(f string, args ...interface{})

	// noCuratorUpdater makes the resolver not run a curator updater. This is used
	// in one-shot resolvers which are given an ahead-of-time list of curators to
	// attempt to contact, eg. joining and registering nodes.
	noCuratorUpdater bool
}

// New starts a new Resolver, ready to be used as a gRPC via WithResolvers.
// However, it needs to be populated with at least one endpoint first (via
// AddEndpoint).
func New(ctx context.Context, opts ...ResolverOption) *Resolver {
	r := &Resolver{
		reqC:   make(chan *request),
		ctx:    ctx,
		logger: func(string, ...interface{}) {},
	}
	for _, opt := range opts {
		opt(r)
	}
	go r.run(ctx)
	return r
}

// ResolverOptions are passed to a Resolver being created.
type ResolverOption func(r *Resolver)

// WithLogger configures a given function as the logger of the resolver. The
// function should take a printf-style format string and arguments.
func WithLogger(logger func(f string, args ...interface{})) ResolverOption {
	return func(r *Resolver) {
		r.logger = logger
	}
}

// WithoutCuratorUpdater configures the Resolver to not attmept to update
// curators from the cluster. This is useful in one-shot resolvers, eg.
// unauthenticated ones.
func WithoutCuratorUpdater() ResolverOption {
	return func(r *Resolver) {
		r.noCuratorUpdater = true
	}
}

// NodeEndpoint is the gRPC endpoint (host+port) of a Metropolis control plane
// node.
type NodeEndpoint struct {
	endpoint string
}

// NodeWithDefaultPort returns a NodeEndpoint referencing the default control
// plane port (the Curator port) of a node resolved by its ID over DNS. This is
// the easiest way to construct a NodeEndpoint provided DNS is fully set up.
func NodeWithDefaultPort(id string) (*NodeEndpoint, error) {
	if m, _ := regexp.MatchString(`metropolis-[a-f0-9]+`, id); !m {
		return nil, fmt.Errorf("invalid node ID")
	}
	return NodeByHostPort(id, uint16(common.CuratorServicePort)), nil
}

// NodeAtAddressWithDefaultPort returns a NodeEndpoint referencing the default
// control plane port (the Curator port) of a node at a given address.
func NodeAtAddressWithDefaultPort(host string) *NodeEndpoint {
	return NodeByHostPort(host, uint16(common.CuratorServicePort))
}

// NodeByHostPort returns a NodeEndpoint for a fully specified host + port pair.
// The host can either be a hostname or an IP address.
func NodeByHostPort(host string, port uint16) *NodeEndpoint {
	return &NodeEndpoint{
		endpoint: net.JoinHostPort(host, fmt.Sprintf("%d", port)),
	}
}

// nodeAtListener is used in tests to connect to the address of a given listener.
func nodeAtListener(lis net.Listener) *NodeEndpoint {
	return &NodeEndpoint{
		endpoint: lis.Addr().String(),
	}
}

// AddEndpoint tells the resolver that it should attempt to reach the cluster
// through a node available at the given NodeEndpoint.
//
// The resolver will make use of this during the next leadership find routine,
// but this node might then get overridden when the resolver retrieves the
// newest set of Curators from the acquired leader.
func (r *Resolver) AddEndpoint(endpoint *NodeEndpoint) {
	select {
	case <-r.ctx.Done():
		return
	case r.reqC <- &request{
		sa: &requestSeedAdd{
			endpoint: endpoint,
		},
	}:
	}
}

// AddOverride adds a long-lived override which forces the resolver to assume
// that a given node (by ID) is available at the given endpoint, instead of at
// whatever endpoint is reported by the cluster. This should be used sparingly
// outside the cluster, and is mostly designed so that nodes which connect to
// themselves can do so over the loopback address instead of their (possibly
// changing) external address.
func (r *Resolver) AddOverride(id string, ep *NodeEndpoint) {
	select {
	case <-r.ctx.Done():
		return
	case r.reqC <- &request{
		oa: &requestOverrideAdd{
			nodeID:   id,
			endpoint: ep,
		},
	}:
	}
}

// runCuratorUpdater runs the curator updater, noted in logs as CURUPDATE. It
// uses the resolver itself to contact the current leader, retrieve all nodes
// which are running a curator, and populate the processor's curator list in the
// curatorMap. That curatorMap will then be used by the leader updater to find
// the current leader.
func (r *Resolver) runCuratorUpdater(ctx context.Context, opts []grpc.DialOption) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 0
	bo.MaxInterval = 10 * time.Second

	return backoff.RetryNotify(func() error {
		opts = append(opts, grpc.WithResolvers(r))
		cl, err := grpc.Dial(MetropolisControlAddress, opts...)
		if err != nil {
			// This generally shouldn't happen.
			return fmt.Errorf("could not dial gRPC: %v", err)
		}
		defer cl.Close()

		cur := apb.NewCuratorClient(cl)
		w, err := cur.Watch(ctx, &apb.WatchRequest{
			Kind: &apb.WatchRequest_NodesInCluster_{
				NodesInCluster: &apb.WatchRequest_NodesInCluster{},
			},
		})
		if err != nil {
			return fmt.Errorf("could not watch nodes: %v", err)
		}

		// Map from node ID to status.
		nodes := make(map[string]*cpb.NodeStatus)

		// Keep updating map from watcher.
		for {
			ev, err := w.Recv()
			if err != nil {
				return fmt.Errorf("when receiving node: %w", err)
			}
			bo.Reset()

			// Update internal map.
			for _, n := range ev.Nodes {
				nodes[n.Id] = n.Status
			}
			for _, n := range ev.NodeTombstones {
				delete(nodes, n.NodeId)
			}

			// Make a copy, this time only curators.
			curators := make(map[string]*cpb.NodeStatus)
			var curatorNames []string
			for k, v := range nodes {
				if v == nil || v.RunningCurator == nil {
					continue
				}
				curators[k] = v
				curatorNames = append(curatorNames, k)
			}
			r.logger("CURUPDATE: got new curators: %s", strings.Join(curatorNames, ", "))

			select {
			case r.reqC <- &request{nu: &requestNodesUpdate{nodes: curators}}:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}, backoff.WithContext(bo, ctx), func(err error, t time.Duration) {
		r.logger("CURUPDATE: error in loop: %v", err)
		r.logger("CURUPDATE: retrying in %s...", t.String())
	})
}

// runLeaderUpdater runs the leader updater, noted in logs as FINDLEADER and
// WATCHLEADER. It uses the curator map from the resolver processor to find the
// current leader.
func (r *Resolver) runLeaderUpdater(ctx context.Context, opts []grpc.DialOption) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 0
	bo.MaxInterval = 10 * time.Second

	return backoff.RetryNotify(func() error {
		curMap := r.curatorMap()
		for _, endpoint := range curMap.candidates() {
			ok := r.watchLeaderVia(ctx, endpoint, opts)
			if ok {
				bo.Reset()
			}
		}
		return fmt.Errorf("out of endpoints")
	}, backoff.WithContext(bo, ctx), func(err error, t time.Duration) {
		r.logger("FINDLEADER: error in loop: %v, retrying in %s...", err, t.String())
	})
}

// watchLeaderVia connects to the endpoint defined by 'via' and attempts to
// continuously update the current leader (b.leader) based on data returned from
// it. Whenever new information about a leader is available, b.condLeader is
// updated.
//
// A boolean value is returned indicating whether the update was at all
// successful. This is used by retry logic to figure out whether to wait before
// retrying or not.
func (r *Resolver) watchLeaderVia(ctx context.Context, via string, opts []grpc.DialOption) bool {
	cl, err := grpc.Dial(via, opts...)
	if err != nil {
		r.logger("WATCHLEADER: dialing %s failed: %v", via, err)
		return false
	}
	defer cl.Close()
	cpl := apb.NewCuratorLocalClient(cl)

	cur, err := cpl.GetCurrentLeader(ctx, &apb.GetCurrentLeaderRequest{})
	if err != nil {
		r.logger("WATCHLEADER: failed to retrieve current leader from %s: %v", via, err)
		return false
	}
	ok := false
	for {
		leaderInfo, err := cur.Recv()
		if err == io.EOF {
			r.logger("WATCHLEADER: connection with %s closed", via)
			return ok
		}
		if err != nil {
			r.logger("WATCHLEADER: connection with %s failed: %v", via, err)
			return ok
		}

		curMap := r.curatorMap()

		viaID := leaderInfo.ThisNodeId
		if viaID == "" {
			// This shouldn't happen, but let's handle this just in case
			viaID = fmt.Sprintf("UNKNOWN NODE ID (%s)", via)
		}

		if leaderInfo.LeaderNodeId == "" {
			r.logger("WATCHLEADER: %s does not know the leader, trying next", viaID)
			return false
		}
		endpoint := ""
		if leaderInfo.LeaderHost == "" {
			// This node knows the leader, but doesn't know its host. Perhaps we have an
			// override for this?
			if ep, ok := curMap.overrides[leaderInfo.LeaderNodeId]; ok {
				endpoint = ep.endpoint
			}
		} else {
			if leaderInfo.LeaderPort == 0 {
				r.logger("WATCHLEADER: %s knows the leader's host (%s), but not its' port", viaID, leaderInfo.LeaderHost)
				return false
			}
			endpoint = net.JoinHostPort(leaderInfo.LeaderHost, fmt.Sprintf("%d", leaderInfo.LeaderPort))
		}

		r.logger("WATCHLEADER: got new leader: %s (%s) via %s", leaderInfo.LeaderNodeId, endpoint, viaID)

		select {
		case <-ctx.Done():
			return ok
		case r.reqC <- &request{lu: &requestLeaderUpdate{
			nodeID:   leaderInfo.LeaderNodeId,
			endpoint: &NodeEndpoint{endpoint: endpoint},
		}}:
		}

		ok = true
	}
}
