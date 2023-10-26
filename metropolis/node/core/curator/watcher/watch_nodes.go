package watcher

import (
	"context"
	"fmt"
	"slices"
	"sort"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

// nodeSet is a collection of Node data. It accumulates updates from Events
// returned from a Curator Watch RPC.
//
// Node data stored within a nodeSet are immutable - any modifications to Node
// data are performed by replacing the entirety of the structure. This means that
// it is safe to hold on to Node pointers acquired from the nodeSet and access
// their fields concurrently from other goroutines, but that these pointers
// represent a snapshot in time of a node's state, and will not get updated as
// the watched node gets updated.
//
// This structure is safe to use in its zero form.
type nodeSet struct {
	// nodes is a map from node ID to node data.
	nodes map[string]*ipb.Node
	// nodeNames is an ordered list of node IDs. In combination with the nodes map,
	// it provides an 'ordered set' semantic to the nodeSet.
	nodeNames []string
}

// clone performs a shallow copy of a nodeSet, not cloning the underlying Node
// structures. This is fine, as Node structures stored in a nodeSet are
// immutable.
func (n *nodeSet) clone() *nodeSet {
	nodes := make(map[string]*ipb.Node)
	nodeNames := make([]string, 0, len(n.nodeNames))
	for _, name := range n.nodeNames {
		nodeNames = append(nodeNames, name)
		nodes[name] = n.nodes[name]
	}
	return &nodeSet{
		nodes:     nodes,
		nodeNames: nodeNames,
	}
}

// updateFromEvent a node set based on data received from a Curator's WatchEvent.
//
// The stored Nodes are not mutated, but instead are replaced with new Node
// structures reflecting the new state of the cluster's nodes.
func (n *nodeSet) updateFromEvent(ev *ipb.WatchEvent) {
	if n.nodes == nil {
		n.nodes = make(map[string]*ipb.Node)
	}

	nodesAdded := false
	nodesDeleted := false

	for _, node := range ev.Nodes {
		// Add to nodeNames if this is the first time we see this node.
		if _, ok := n.nodes[node.Id]; !ok {
			n.nodeNames = append(n.nodeNames, node.Id)
			nodesAdded = true
		}
		// Replace node in map.
		n.nodes[node.Id] = node
	}

	// Delete nodes which have been tombstoned.
	deleted := make(map[string]bool)
	for _, node := range ev.NodeTombstones {
		deleted[node.NodeId] = true
		delete(n.nodes, node.NodeId)
		nodesDeleted = true
	}

	if nodesDeleted {
		n.nodeNames = slices.DeleteFunc(n.nodeNames, func(id string) bool { return deleted[id] })
	}
	if nodesAdded || nodesDeleted {
		sort.Strings(n.nodeNames)
	}
}

// follow updates a nodeSet from another nodeSet, but uses a Follower interface
// to filter out nodes and call back into external systems with information about
// node lifecycle events.
func (n *nodeSet) follow(origin *nodeSet, f Follower) error {
	if n.nodes == nil {
		n.nodes = make(map[string]*ipb.Node)
	}
	seen := make(map[string]bool)
	for _, name := range origin.nodeNames {
		if !f.Filter(origin.nodes[name]) {
			continue
		}
		seen[name] = true
		if _, ok := n.nodes[name]; !ok {
			// New node.
			n.nodes[name] = origin.nodes[name]
			if err := f.New(n.nodes[name]); err != nil {
				return fmt.Errorf("new node %s: %w", name, err)
			}
			n.nodeNames = append(n.nodeNames, name)
			continue
		} else {
			// Updated node.
			if !f.Equals(n.nodes[name], origin.nodes[name]) {
				if err := f.Updated(n.nodes[name], origin.nodes[name]); err != nil {
					return fmt.Errorf("updated node %s: %w", name, err)
				}
			}
			n.nodes[name] = origin.nodes[name]
		}
	}

	for _, name := range n.nodeNames {
		if seen[name] {
			continue
		}
		if err := f.Deleted(n.nodes[name]); err != nil {
			return fmt.Errorf("deleted node %s: %w", name, err)
		}
	}

	if err := f.BatchDone(); err != nil {
		return fmt.Errorf("batch done: %w", err)
	}

	n.nodeNames = slices.DeleteFunc(n.nodeNames, func(id string) bool { return !seen[id] })
	sort.Strings(n.nodeNames)
	return nil
}

// A Follower is some subsystem which wishes to be notified about changes to a
// cluster's node state.
//
// It provides function to filter out state and state transitions that are
// interesting to itself, and functions which will be called when the filtered
// state changes.
//
// The Filter and Equals functions make up a 'view' of the cluster state from the
// point of view of the Follower. That is, a Follower which only cares about some
// subset of nodes and expresses said subset with Filter will only see these
// nodes in its nodeSet and in its callbacks' calls. Similarly, updates to the
// nodes will also be filtered out accordingly to Equals.
//
// A simple callback-based implementation is available in SimpleFollower.
type Follower interface {
	// Filter should return true if a node is of interest to the follower - when it
	// has all required fields present and at a requested state.
	//
	// For example, a Follower which wishes to watch for nodes' external IP
	// addresses would filter out all nodes which don't have an address assigned.
	Filter(a *ipb.Node) bool

	// Equals should return true if a given node's state is identical, from the point
	// of view of the Follower, to some other state. Correctly implementing this
	// function allows the Follower to only receive calls to New/Updated/Deleted when
	// the node actually changed in a meaningful and actionable way.
	//
	// For example, a Follower which wishes to watch for nodes' external IP addresses
	// would return true only if the two nodes' external IP addresses actually
	// differed.
	Equals(a *ipb.Node, b *ipb.Node) bool

	// New will be called when a node has appeared from the point of view of the
	// Follower (i.e. started existing on the cluster and then also passed the Filter
	// function).
	//
	// Any returned error is considered fatal and will stop future use of the
	// Follower, e.g. WatchNodes will return.
	New(new *ipb.Node) error

	// Updated will be called when a node has been updated from the point of view of
	// the Follower (i.e. has not been filtered out, and Equals returned false).
	//
	// Any returned error is considered fatal and will stop future use of the
	// Follower, e.g. WatchNodes will return.
	Updated(prev *ipb.Node, new *ipb.Node) error

	// Deleted will be called when a node has been removed from the point of view of
	// the Follower (i.e. has been filtered out, or has been removed from the cluster
	// altogether).
	//
	// Any returned error is considered fatal and will stop future use of the
	// Follower, e.g. WatchNodes will return.
	Deleted(prev *ipb.Node) error

	// BatchDone will be called at the end of any batch of node updates (either New,
	// Updated or Deleted calls). This can be used by Followers to reduce the number
	// of mutations of an expensive resource, for example if the Nodes watch
	// mechanism is used to feed some other stateful system which also supports
	// batch-based updates.
	//
	// Just exactly how large batches are is an implementation detail of the
	// underlying Curator watch protocol and the way update events get created by the
	// Curator and sent over the wire.
	//
	// Note: BatchDone() will not be called if any of the New/Updated/Deleted
	// implementations returned an error - the follower will be terminated
	// immediately!
	BatchDone() error
}

// SimpleFollower is a callback struct based implementation of a Follower, with
// the additional collapse of New and Updated into a NewUpdated function.
//
// This is the simplest way to use the Follower / WatchNodes system from a
// function.
type SimpleFollower struct {
	// FilterFn corresponds to Follower.Filter - see its documentation for more
	// details.
	FilterFn func(a *ipb.Node) bool
	// EqualsFn corresponds to Follower.Equals - see its documentation for more
	// details.
	EqualsFn func(a *ipb.Node, b *ipb.Node) bool

	// OnNewUpdated will be called whenever a node is updated or appears for the
	// first time from the point of view of the Follower.
	OnNewUpdated func(new *ipb.Node) error
	// OnDeleted will be called whenever a node disappears from the point of view of
	// the Follower.
	OnDeleted func(prev *ipb.Node) error

	// OnBatchDone will be called at the end of a batch of NewUpdated/Deleted calls
	// from the underlying Curator watch mechanism.
	OnBatchDone func() error
}

func (f SimpleFollower) Filter(a *ipb.Node) bool {
	if f.FilterFn == nil {
		return true
	}
	return f.FilterFn(a)
}

func (f SimpleFollower) Equals(a *ipb.Node, b *ipb.Node) bool {
	return f.EqualsFn(a, b)
}

func (f SimpleFollower) New(new *ipb.Node) error {
	if f.OnNewUpdated == nil {
		return nil
	}
	return f.OnNewUpdated(new)
}

func (f SimpleFollower) Updated(_prev *ipb.Node, new *ipb.Node) error {
	if f.OnNewUpdated == nil {
		return nil
	}
	return f.OnNewUpdated(new)
}

func (f SimpleFollower) Deleted(prev *ipb.Node) error {
	if f.OnDeleted == nil {
		return nil
	}
	return f.OnDeleted(prev)
}

func (f SimpleFollower) BatchDone() error {
	if f.OnBatchDone == nil {
		return nil
	}
	return f.OnBatchDone()
}

// WatchNodes runs a WatchRequest for NodesInCluster with the given Curator
// channel. Any updates to the state of the nodes is processed through the given
// Follower.
//
// This is the main interface to follow a state of nodes in the cluster and act
// upon any changes. SimpleFollower is given as a type to implement the simplest
// kind of callback-driven interface to various events, but users are free to
// implement Follower on their own.
//
// This function will exit with a context error whenever the given context is
// canceled, or return with whatever error is returned by the Follower
// implementation.
func WatchNodes(ctx context.Context, cur ipb.CuratorClient, f Follower) error {
	wa, err := cur.Watch(ctx, &ipb.WatchRequest{
		Kind: &ipb.WatchRequest_NodesInCluster_{
			NodesInCluster: &ipb.WatchRequest_NodesInCluster{},
		},
	})
	if err != nil {
		return fmt.Errorf("watch request failed: %w", err)
	}
	defer wa.CloseSend()

	var ons nodeSet
	var ns nodeSet
	for {
		ev, err := wa.Recv()
		if err != nil {
			return fmt.Errorf("receive failed: %w", err)
		}
		ons.updateFromEvent(ev)

		if err := ns.follow(&ons, f); err != nil {
			return err
		}
	}
}
