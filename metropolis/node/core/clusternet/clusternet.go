// Package clusternet implements a Cluster Networking mesh service running on all
// Metropolis nodes.
//
// The mesh is based on wireguard and a centralized configuration store in the
// cluster Curator (in etcd).
//
// While the implementation is nearly generic, it currently makes an assumption
// that it is used only for Kubernetes pod networking. That has a few
// implications:
//
// First, we only have a single real route on the host into the wireguard
// networking mesh / interface, and that is configured ahead of time in the
// Service as ClusterNet. All destination addresses that should be carried by the
// mesh must thus be part of this single route. Otherwise, traffic will be able
// to flow into the node from other nodes, but will exit through another
// interface. This is used in practice to allow other host nodes (whose external
// addresses are outside the cluster network) to access the cluster network.
//
// Second, we have two hardcoded/purpose-specific sources of prefixes:
//  1. Pod networking node prefixes from the kubelet
//  2. The host's external IP address (as a /32) from the network service.
package clusternet

import (
	"context"
	"fmt"
	"net"
	"net/netip"

	"github.com/cenkalti/backoff/v4"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/supervisor"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// Service implements the Cluster Networking Mesh. See package-level docs for
// more details.
type Service struct {
	// Curator is the gRPC client that the service will use to reach the cluster's
	// Curator, for pushing locally announced prefixes and pulling information about
	// other nodes.
	Curator apb.CuratorClient
	// ClusterNet is the prefix that will be programmed to exit through the wireguard
	// mesh.
	ClusterNet net.IPNet
	// DataDirectory is where the WireGuard key of this node will be stored.
	DataDirectory *localstorage.DataKubernetesClusterNetworkingDirectory
	// LocalKubernetesPodNetwork is an event.Value watched for prefixes that should
	// be announced into the mesh. This is to be Set by the Kubernetes service once
	// it knows about the local node's IPAM address assignment.
	LocalKubernetesPodNetwork event.Value[*Prefixes]
	// Network service used to get the local node's IP address to submit it as a /32.
	Network event.Value[*network.Status]

	// wg is the interface to all the low-level interactions with WireGuard (and
	// kernel routing). If not set, this defaults to a production implementation.
	// This can be overridden by test to a test implementation instead.
	wg wireguard
}

// Run the Service. This must be used in a supervisor Runnable.
func (s *Service) Run(ctx context.Context) error {
	if s.wg == nil {
		s.wg = &localWireguard{}
	}
	if err := s.wg.ensureOnDiskKey(s.DataDirectory); err != nil {
		return fmt.Errorf("could not ensure wireguard key: %w", err)
	}
	if err := s.wg.setup(&s.ClusterNet); err != nil {
		return fmt.Errorf("could not setup wireguard: %w", err)
	}

	supervisor.Logger(ctx).Infof("Wireguard setup complete, starting updaters...")

	kubeC := make(chan *Prefixes)
	netC := make(chan *network.Status)
	if err := supervisor.RunGroup(ctx, map[string]supervisor.Runnable{
		"source-kubernetes": event.Pipe(s.LocalKubernetesPodNetwork, kubeC),
		"source-network":    event.Pipe(s.Network, netC),
		"push": func(ctx context.Context) error {
			return s.push(ctx, kubeC, netC)
		},
	}); err != nil {
		return err
	}

	if err := supervisor.Run(ctx, "pull", s.pull); err != nil {
		return err
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return ctx.Err()
}

// push is the sub-runnable responsible for letting the Curator know about what
// prefixes that are originated by this node.
func (s *Service) push(ctx context.Context, kubeC chan *Prefixes, netC chan *network.Status) error {
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	var kubePrefixes *Prefixes
	var prevKubePrefixes *Prefixes

	var localAddr net.IP
	var prevLocalAddr net.IP

	for {
		kubeChanged := false
		localChanged := false

		select {
		case <-ctx.Done():
			return ctx.Err()
		case kubePrefixes = <-kubeC:
			if !kubePrefixes.Equal(prevKubePrefixes) {
				kubeChanged = true
			}
		case n := <-netC:
			localAddr = n.ExternalAddress
			if !localAddr.Equal(prevLocalAddr) {
				localChanged = true
			}
		}

		// Ignore spurious updates.
		if !localChanged && !kubeChanged {
			continue
		}

		// Prepare prefixes to submit to cluster.
		var prefixes Prefixes

		// Do we have a local node address? Add it to the prefixes.
		if len(localAddr) > 0 {
			addr, ok := netip.AddrFromSlice(localAddr)
			if ok {
				prefixes = append(prefixes, netip.PrefixFrom(addr, 32))
			}
		}
		// Do we have any kubelet prefixes? Add them, too.
		if kubePrefixes != nil {
			prefixes.Update(kubePrefixes)
		}

		supervisor.Logger(ctx).Infof("Submitting prefixes: %s (kube update: %v, local update: %v)", prefixes, kubeChanged, localChanged)

		err := backoff.Retry(func() error {
			_, err := s.Curator.UpdateNodeClusterNetworking(ctx, &apb.UpdateNodeClusterNetworkingRequest{
				Clusternet: &cpb.NodeClusterNetworking{
					WireguardPubkey: s.wg.key().PublicKey().String(),
					Prefixes:        prefixes.proto(),
				},
			})
			if err != nil {
				supervisor.Logger(ctx).Warningf("Could not submit cluster networking update: %v", err)
			}
			return err
		}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
		if err != nil {
			return fmt.Errorf("couldn't update curator: %w", err)
		}

		prevKubePrefixes = kubePrefixes
		prevLocalAddr = localAddr

	}
}

// pull is the sub-runnable responsible for fetching information about the
// cluster networking setup/status of other nodes, and programming it as
// WireGuard peers.
func (s *Service) pull(ctx context.Context) error {
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

	nodes := newNodemap()
	for {
		ev, err := srv.Recv()
		if err != nil {
			return fmt.Errorf("curator watch recv failed: %w", err)
		}

		updated, removed := nodes.update(ctx, ev)

		for _, n := range removed {
			supervisor.Logger(ctx).Infof("Node %s removed, unconfiguring", n.id)
			if err := s.wg.unconfigurePeer(n.copy()); err != nil {
				// Do nothing and hope whatever caused this will go away at some point.
				supervisor.Logger(ctx).Errorf("Node %s couldn't be unconfigured: %v", n.id, err)
			}
		}
		var newNodes []*node
		for _, n := range updated {
			newNodes = append(newNodes, n.copy())
			supervisor.Logger(ctx).Infof("Node %s updated: pk %s, address %s, prefixes %v", n.id, n.pubkey, n.address, n.prefixes)
		}
		succeeded := 0
		if err := s.wg.configurePeers(newNodes); err != nil {
			// If configuring all nodes at once failed, go node-by-node to make sure we've
			// done as much as possible.
			supervisor.Logger(ctx).Warningf("Bulk node update call failed, trying node-by-node..: %v", err)
			for _, n := range newNodes {
				if err := s.wg.configurePeers([]*node{n}); err != nil {
					supervisor.Logger(ctx).Errorf("Node %s failed: %v", n.id, err)
				} else {
					succeeded += 1
				}
			}
		} else {
			succeeded = len(newNodes)
		}

		if len(newNodes) != 0 {
			supervisor.Logger(ctx).Infof("Successfully updated %d out of %d nodes", succeeded, len(newNodes))

			numNodes, numPrefixes := nodes.stats()
			supervisor.Logger(ctx).Infof("Total: %d nodes, %d prefixes.", numNodes, numPrefixes)
		}
	}
}
