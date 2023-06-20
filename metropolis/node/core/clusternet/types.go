package clusternet

import (
	"context"
	"net/netip"
	"sort"
	"strings"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// Prefixes are network prefixes that should be announced by a node to the
// Cluster Networking mesh.
type Prefixes []netip.Prefix

func (p *Prefixes) proto() (res []*cpb.NodeClusterNetworking_Prefix) {
	for _, prefix := range *p {
		res = append(res, &cpb.NodeClusterNetworking_Prefix{
			Cidr: prefix.String(),
		})
	}
	return
}

// Update by copying all prefixes from o into p, merging duplicates as necessary.
func (p *Prefixes) Update(o *Prefixes) {
	// Gather prefixes we already have.
	cur := make(map[netip.Prefix]bool)
	for _, pp := range *p {
		cur[pp] = true
	}

	// Copy over any prefix that we don't yet have.
	for _, pp := range *o {
		if cur[pp] {
			continue
		}
		cur[pp] = true
		*p = append(*p, pp)
	}
}

// String returns a stringified, comma-dalimited representation of the prefixes.
func (p *Prefixes) String() string {
	var strs []string
	for _, pp := range *p {
		strs = append(strs, pp.String())
	}
	return strings.Join(strs, ", ")
}

// node is used for internal statekeeping in the cluster networking service.
type node struct {
	id       string
	pubkey   string
	address  string
	prefixes []string
}

func (n *node) copy() *node {
	n2 := *n
	return &n2
}

// update mutates this node to whatever data is held in the given proto Node, and
// returns true if any data changed.
func (n *node) update(p *apb.Node) (changed bool) {
	if n.address != p.Status.ExternalAddress {
		n.address = p.Status.ExternalAddress
		changed = true
	}
	if n.pubkey != p.Clusternet.WireguardPubkey {
		n.pubkey = p.Clusternet.WireguardPubkey
		changed = true
	}

	var newPrefixes []string
	for _, prefix := range p.Clusternet.Prefixes {
		if prefix.Cidr == "" {
			continue
		}
		newPrefixes = append(newPrefixes, prefix.Cidr)
	}
	oldPrefixes := make([]string, len(n.prefixes))
	copy(oldPrefixes[:], n.prefixes)

	sort.Strings(newPrefixes)
	sort.Strings(oldPrefixes)
	if want, got := strings.Join(newPrefixes, ","), strings.Join(oldPrefixes, ","); want != got {
		n.prefixes = newPrefixes
		changed = true
	}

	return
}

// nodeMap is the main internal statekeeping structure of the pull sub-runnable.
type nodeMap struct {
	nodes map[string]*node
}

func newNodemap() *nodeMap {
	return &nodeMap{
		nodes: make(map[string]*node),
	}
}

// update updates the nodeMap from the given Curator WatchEvent, interpreting
// both node changes and deletions. Two nodeMaps are returned: the first one
// contains only nodes that have been added/changed by the given event, the other
// contains only nodes that have been deleted by the given event.
func (m *nodeMap) update(ctx context.Context, ev *apb.WatchEvent) (changed, removed map[string]*node) {
	changed = make(map[string]*node)
	removed = make(map[string]*node)

	// Make sure we're not getting multiple nodes with the same public key. This is
	// not expected to happen in practice as the Curator should prevent this from
	// happening, but we at least want to make sure we're not blowing up routing if
	// other defenses fail.

	// pkeys maps from public key to node ID.
	pkeys := make(map[string]string)
	for _, n := range m.nodes {
		// We don't have to check whether we have any collisions already in m.nodes, as
		// the check below prevents them from happening in the first place.
		pkeys[n.pubkey] = n.id
	}

	for _, n := range ev.Nodes {
		// Only care about nodes that have all required configuration set.
		if n.Status == nil || n.Status.ExternalAddress == "" || n.Clusternet == nil || n.Clusternet.WireguardPubkey == "" {
			// We could attempt to delete any matching node currently in this nodemap at this
			// point, but this is likely transient and we don't want to just kill routing for
			// no reason.
			continue
		}

		key := n.Clusternet.WireguardPubkey
		if id, ok := pkeys[key]; ok && id != n.Id {
			supervisor.Logger(ctx).Warningf("Nodes %q and %q share wireguard key %q. That should not have happened.", n.Id, id, key)
			continue
		}

		if _, ok := m.nodes[n.Id]; !ok {
			m.nodes[n.Id] = &node{
				id: n.Id,
			}
		}
		diff := m.nodes[n.Id].update(n)
		if diff {
			changed[n.Id] = m.nodes[n.Id]
		}
	}

	for _, t := range ev.NodeTombstones {
		n, ok := m.nodes[t.NodeId]
		if !ok {
			// This is an indication of us losing data somehow. If this happens, it likely
			// means a Curator bug.
			supervisor.Logger(ctx).Warningf("Node %s: tombstone for unknown node", t.NodeId)
			continue
		}
		removed[n.id] = n
		delete(m.nodes, n.id)
	}

	return
}
