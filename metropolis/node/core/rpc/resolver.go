package rpc

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

const (
	MetropolisControlAddress = "metropolis:///control"
)

// ClusterResolver is a gRPC resolver Builder that can be passed to
// grpc.WithResolvers() when dialing a gRPC endpoint.
//
// It's responsible for resolving the magic MetropolisControlAddress
// (metropolis:///control) into all Metropolis nodes running control plane
// services, ie. the Curator.
//
// To function, the ClusterResolver needs to be provided with at least one node
// address. Afterwards, it will continuously update an internal list of nodes
// which can be contacted for access to control planes services, and gRPC
// clients using this resolver will automatically try the available addresses
// for each RPC call in a round-robin fashion.
//
// The ClusterResolver is designed to be used as a long-running objects which
// multiple gRPC client connections can use. Usually one ClusterResolver
// instance should be used per application.
type ClusterResolver struct {
	ctx  context.Context
	ctxC context.CancelFunc

	// logger, if set, will be called with fmt.Sprintf-like arguments containing
	// debug logs from the running ClusterResolver, subordinate watchers and
	// updaters.
	logger func(f string, args ...interface{})

	condCurators  *sync.Cond
	curators      map[string]string
	condTLSConfig *sync.Cond
	tlsConfig     credentials.TransportCredentials
}

// AddNode provides a given node ID at a given address as an initial (or
// additional) node for the ClusterResolver to update cluster information
// from.
func (b *ClusterResolver) AddNode(name, remote string) {
	b.condCurators.L.Lock()
	defer b.condCurators.L.Unlock()

	b.curators[name] = remote
	b.condCurators.Broadcast()
}

// NewClusterResolver creates an empty ClusterResolver. It must be populated
// with initial node information for any gRPC call that uses it to succeed.
func NewClusterResolver() *ClusterResolver {
	ctx, ctxC := context.WithCancel(context.Background())
	b := &ClusterResolver{
		ctx:           ctx,
		ctxC:          ctxC,
		logger:        func(f string, args ...interface{}) {},
		condCurators:  sync.NewCond(&sync.Mutex{}),
		curators:      make(map[string]string),
		condTLSConfig: sync.NewCond(&sync.Mutex{}),
	}

	go b.run(b.ctx)

	return b
}

var (
	ResolverClosed = errors.New("cluster resolver closed")
)

// Close the ClusterResolver to clean up background goroutines. The node address
// resolution process stops and all future connections done via this
// ClusterResolver will continue to use whatever node addresses were last known.
// However, new attempts to dial using this ClusterResolver will fail.
func (b *ClusterResolver) Close() {
	b.ctxC()
}

// run is the main loop of the ClusterResolver. Its job is to wait for a TLS
// config from a gRPC client, and iterate through available node addresses to
// start an updater on. The updater will then communicate back to this goroutine
// with up to date node information. In case an updater cannot run anymore (eg.
// a node stopped working), the main loop of run restarts and another endpoint
// will be picked.
func (b *ClusterResolver) run(ctx context.Context) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 0

	// Helper to update internal node list and notify all gRPC clients of it.
	updateCurators := func(nodes map[string]string) {
		b.condCurators.L.Lock()
		b.curators = nodes
		b.condCurators.L.Unlock()
		b.condCurators.Broadcast()
	}

	// Helper to sleep for a given time, but with possible interruption by the
	// resolver being stopped.
	waitTimeout := func(t time.Duration) bool {
		select {
		case <-time.After(t):
			return true
		case <-ctx.Done():
			return false
		}
	}

	for {
		b.logger("RESOLVER: waiting for TLS config...")
		// Wait for a TLS config to be set.
		b.condTLSConfig.L.Lock()
		for b.tlsConfig == nil {
			b.condTLSConfig.Wait()
		}
		creds := b.tlsConfig
		b.condTLSConfig.L.Unlock()
		b.logger("RESOLVER: have TLS config...")

		// Iterate over endpoints to find a working one, and retrieve cluster-provided
		// node info from there.
		endpoints := b.addresses()
		if len(endpoints) == 0 {
			w := bo.NextBackOff()
			b.logger("RESOLVER: no endpoints, waiting %s...", w)
			if waitTimeout(w) {
				b.logger("RESOLVER: canceled")
				return
			}
			continue
		}

		b.logger("RESOLVER: starting endpoint loop with %v...", endpoints)
		for name, endpoint := range endpoints {
			upC := make(chan map[string]string)
			b.logger("RESOLVER: starting updater pointed at %s/%s", name, endpoint)

			// Start updater, which actually connects to the endpoint and provides back the
			// newest set of nodes via upC.
			go b.runUpdater(ctx, endpoint, creds, upC)

			// Keep using this updater as long as possible. If it fails, restart the main
			// loop.
			failed := false
			for {
				var newNodes map[string]string
				failed := false
				select {
				case newNodes = <-upC:
					if newNodes == nil {
						// Updater quit.
						failed = true
					}
				case <-ctx.Done():
					b.logger("RESOLVER: canceled")
					updateCurators(nil)
					return
				}

				if failed {
					w := bo.NextBackOff()
					b.logger("RESOLVER: updater failed, waiting %s...", w)
					if waitTimeout(w) {
						b.logger("RESOLVER: canceled")
						return
					}
					b.logger("RESOLVER: done waiting")
					break
				} else {
					bo.Reset()
					updateCurators(newNodes)
				}

			}
			// Restart entire ClusterResolver loop on failure.
			if failed {
				break
			}
		}
	}
}

// runUpdaters runs the ClusterResolver's updater, which is a goroutine that
// connects to a Curator running on a given node and feeds back information
// about consensus members via updateC. If the endpoints fails (eg. because the
// node went down), updateC will be closed.
func (b *ClusterResolver) runUpdater(ctx context.Context, endpoint string, creds credentials.TransportCredentials, updateC chan map[string]string) {
	defer close(updateC)
	cl, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		b.logger("UPDATER: dial failed: %v", err)
		return
	}
	defer cl.Close()
	cur := cpb.NewCuratorClient(cl)
	w, err := cur.Watch(ctx, &cpb.WatchRequest{
		Kind: &cpb.WatchRequest_NodesInCluster_{
			NodesInCluster: &cpb.WatchRequest_NodesInCluster{},
		},
	})
	if err != nil {
		b.logger("UPDATER: watch failed: %v", err)
		return
	}

	// Maintain a long-term set of node ID to node external address, and populate it
	// from the Curator Watcher above.
	nodes := make(map[string]string)
	for {
		ev, err := w.Recv()
		if err != nil {
			b.logger("UPDATER: recv failed: %v", err)
			return
		}
		for _, node := range ev.Nodes {
			if node.Roles.ConsensusMember == nil {
				delete(nodes, node.Id)
				continue
			}
			st := node.Status
			if st == nil || st.ExternalAddress == "" {
				delete(nodes, node.Id)
				continue
			}
			nodes[node.Id] = st.ExternalAddress
		}
		for _, node := range ev.NodeTombstones {
			delete(nodes, node.NodeId)
		}
		b.logger("UPDATER: new nodes: %v", nodes)
		updateC <- nodes
	}
}

// addresses returns the current set of node addresses that the ClusterResolver
// considers as possible updater candidates.
func (b *ClusterResolver) addresses() map[string]string {
	b.condCurators.L.Lock()
	defer b.condCurators.L.Unlock()

	res := make(map[string]string)
	for k, v := range b.curators {
		res[k] = v
	}
	return res
}

// Build is called by gRPC on each Dial call. It spawns a new clientWatcher,
// whose goroutine receives information about currently available nodes from the
// parent ClusterResolver and actually updates a given gRPC client connection
// with information about the current set of nodes.
func (b *ClusterResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// We can only connect to "metropolis://control".
	if target.Scheme != "metropolis" || target.Authority != "" || target.Endpoint != "control" {
		return nil, fmt.Errorf("invalid target: must be %s, is: %s", MetropolisControlAddress, target.Endpoint)
	}

	if opts.DialCreds == nil {
		return nil, fmt.Errorf("can only be used with clients containing TransportCredentials")
	}

	if b.ctx.Err() != nil {
		return nil, ResolverClosed
	}

	b.condTLSConfig.L.Lock()
	// TODO(q3k): make sure we didn't receive different DialCreds for a different
	// cluster or something.
	b.tlsConfig = opts.DialCreds
	b.condTLSConfig.Broadcast()
	defer b.condTLSConfig.L.Unlock()

	ctx, ctxC := context.WithCancel(b.ctx)
	resolver := &clientWatcher{
		builder:    b,
		clientConn: cc,
		ctx:        ctx,
		ctxC:       ctxC,
	}
	go resolver.watch()
	return resolver, nil
}

func (b *ClusterResolver) Scheme() string {
	return "metropolis"
}

// clientWatcher is a subordinate structure to a given ClusterResolver,
// updating a gRPC ClientConn with information about current endpoints.
type clientWatcher struct {
	builder    *ClusterResolver
	clientConn resolver.ClientConn

	ctx  context.Context
	ctxC context.CancelFunc
}

func (r *clientWatcher) watch() {
	// Craft a trivial gRPC service config which forces round-robin behaviour for
	// RPCs. This makes the gRPC client contact all curators in a round-robin
	// fashion. Ideally, we would prioritize contacting the leader, but this will do
	// for now.
	svcConfig := r.clientConn.ParseServiceConfig(`{ "loadBalancingConfig": [{"round_robin": {}}]}`)

	// Watch for condCurators being updated.
	r.builder.condCurators.L.Lock()
	for {
		if r.ctx.Err() != nil {
			return
		}

		nodes := r.builder.curators
		var addresses []resolver.Address
		for n, addr := range nodes {
			addresses = append(addresses, resolver.Address{
				Addr:       addr,
				ServerName: n,
			})
		}
		r.builder.logger("WATCHER: new addresses: %v", addresses)
		r.clientConn.UpdateState(resolver.State{
			Addresses:     addresses,
			ServiceConfig: svcConfig,
		})
		r.builder.condCurators.Wait()
	}
}

func (r *clientWatcher) ResolveNow(_ resolver.ResolveNowOptions) {
	// No-op. The clientWatcher's watcher runs as fast as possible.
}

func (r *clientWatcher) Close() {
	r.ctxC()
	// Spuriously interrupt all clientWatchers on this ClusterResolver so that this
	// clientWatcher gets to notice it should quit. This isn't ideal.
	r.builder.condCurators.Broadcast()
}
