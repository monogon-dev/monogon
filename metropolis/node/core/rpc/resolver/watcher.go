// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package resolver

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

// clientWatcher is a subordinate structure to a given ClusterResolver,
// updating a gRPC ClientConn with information about current endpoints.
type clientWatcher struct {
	resolver     *Resolver
	clientConn   resolver.ClientConn
	subscription *responseSubscribe
}

var (
	// ErrResolverClosed will be returned by the resolver to gRPC machinery whenever a
	// resolver cannot be used anymore because it was Closed.
	ErrResolverClosed = errors.New("cluster resolver closed")
)

// Build is called by gRPC on each Dial call. It spawns a new clientWatcher,
// whose goroutine receives information about currently available nodes from the
// parent ClusterResolver and actually updates a given gRPC client connection
// with information about the current set of nodes.
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// We can only connect to "metropolis://control".
	if target.URL.Scheme != "metropolis" || target.URL.Host != "" || target.URL.Path != "/control" {
		return nil, fmt.Errorf("invalid target: must be %s, is: %s", MetropolisControlAddress, target.URL.String())
	}

	if opts.DialCreds == nil {
		return nil, fmt.Errorf("can only be used with clients containing TransportCredentials")
	}

	// Submit the dial options to the resolver's processor, quitting if the resolver
	// gets canceled in the meantime.
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(opts.DialCreds),
		grpc.WithContextDialer(opts.Dialer),
	}

	select {
	case <-r.ctx.Done():
		return nil, ErrResolverClosed
	case r.reqC <- &request{
		ds: &requestDialOptionsSet{
			options: options,
		},
	}:
	}

	// Submit a subscription request to the resolver's processor, quitting if the
	// resolver gets canceled in the meantime.

	req := &request{
		sub: &requestSubscribe{resC: make(chan *responseSubscribe)},
	}
	select {
	case <-r.ctx.Done():
		return nil, ErrResolverClosed
	case r.reqC <- req:
	}
	// This receive is uninterruptible by contract - as it's also uninterruptible on
	// the processor side.
	subscription := <-req.sub.resC

	watcher := &clientWatcher{
		resolver:     r,
		clientConn:   cc,
		subscription: subscription,
	}
	go watcher.watch()

	return watcher, nil
}

func (r *Resolver) Scheme() string {
	return "metropolis"
}

func (w *clientWatcher) watch() {
	// Craft a trivial gRPC service config which forces round-robin behaviour. This
	// doesn't really matter for us, as we only ever submit the single leader as a
	// connection endpoint.
	svcConfig := w.clientConn.ParseServiceConfig(`{ "loadBalancingConfig": [{"round_robin": {}}]}`)

	// Watch for leader to be updated.
	for {
		update := <-w.subscription.subC
		if update == nil {
			// A nil result means the channel is closed, which means this watcher has either
			// closed or the resolver has been canceled. Abort loop.
			w.clientConn.ReportError(ErrResolverClosed)
			break
		}
		w.clientConn.UpdateState(resolver.State{
			Addresses: []resolver.Address{
				{
					Addr:       update.endpoint.endpoint,
					ServerName: update.nodeID,
				},
			},
			ServiceConfig: svcConfig,
		})
	}
}

func (w *clientWatcher) ResolveNow(_ resolver.ResolveNowOptions) {
	// No-op. The clientWatcher's watcher runs as fast as possible.
}

func (w *clientWatcher) Close() {
	select {
	case <-w.resolver.ctx.Done():
	case w.resolver.reqC <- &request{
		unsub: &requestUnsubscribe{
			id: w.subscription.id,
		},
	}:
	}
}
