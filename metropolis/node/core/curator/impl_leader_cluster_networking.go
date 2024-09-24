package curator

import (
	"context"
	"errors"
	"net/netip"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/osbase/event"
	"source.monogon.dev/osbase/event/etcd"
)

// preapreClusternetCacheUnlocked makes sure the leader's clusternetCache exists,
// and loads it from etcd otherwise.
func (l *leaderCurator) prepareClusternetCacheUnlocked(ctx context.Context) error {
	if l.ls.clusternetCache != nil {
		return nil
	}

	cache := make(map[string]string)

	// Get all nodes.
	start, end := NodeEtcdPrefix.KeyRange()
	value := etcd.NewValue[*nodeAtID](l.etcd, start, nodeValueConverter, etcd.Range(end))
	w := value.Watch()
	defer w.Close()
	for {
		nodeKV, err := w.Get(ctx, event.BacklogOnly[*nodeAtID]())
		if errors.Is(err, event.ErrBacklogDone) {
			break
		}
		if err != nil {
			rpc.Trace(ctx).Printf("etcd watch failed (initial fetch): %v", err)
			return status.Error(codes.Unavailable, "internal error during clusternet cache load")
		}
		n := nodeKV.value
		if n == nil {
			continue
		}

		// Ignore nodes without cluster networking.
		if n.wireguardKey == "" {
			continue
		}

		// If we have an inconsistency in the database, just pretend it's not there.
		//
		// TODO(q3k): try to recover from this.
		if id, ok := cache[n.wireguardKey]; ok && id != n.ID() {
			continue
		}
		cache[n.wireguardKey] = n.ID()
	}

	l.ls.clusternetCache = cache
	return nil
}

func (l *leaderCurator) UpdateNodeClusterNetworking(ctx context.Context, req *ipb.UpdateNodeClusterNetworkingRequest) (*ipb.UpdateNodeClusterNetworkingResponse, error) {
	// Ensure that the given node_id matches the calling node. We currently
	// only allow for direct self-reporting of status by nodes.
	pi := rpc.GetPeerInfo(ctx)
	if pi == nil || pi.Node == nil {
		return nil, status.Error(codes.PermissionDenied, "only nodes can update node cluster networking")
	}
	id := pi.Node.ID

	if req.Clusternet == nil {
		return nil, status.Error(codes.InvalidArgument, "clusternet must be set")
	}
	cn := req.Clusternet
	if cn.WireguardPubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "clusternet.wireguard_pubkey must be set")
	}
	key, err := wgtypes.ParseKey(cn.WireguardPubkey)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "clusternet.wireguard_pubkey must be a valid wireguard public key")
	}

	// Lock everything, as we're doing a complex read/modify/store here.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	if err := l.prepareClusternetCacheUnlocked(ctx); err != nil {
		return nil, err
	}

	if nid, ok := l.ls.clusternetCache[key.String()]; ok && nid != id {
		return nil, status.Error(codes.InvalidArgument, "public key alread used by another node")
	}

	// TODO(q3k): unhardcode this and synchronize with Kubernetes code.
	clusterNet := netip.MustParsePrefix("10.192.0.0/11")

	// Retrieve node ...
	node, err := nodeLoad(ctx, l.leadership, id)
	if err != nil {
		return nil, err
	}

	if node.status == nil {
		return nil, status.Error(codes.FailedPrecondition, "node needs to submit at least one status update")
	}
	externalIP := node.status.ExternalAddress

	// Parse/validate given prefixes.
	var prefixes []netip.Prefix
	for i, prefix := range cn.Prefixes {
		// Parse them.
		p, err := netip.ParsePrefix(prefix.Cidr)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "clusternet.prefixes[%d].cidr invalid: %v", i, err)
		}

		// Make sure they're in canonical form.
		masked := p.Masked()
		if masked.String() != p.String() {
			return nil, status.Errorf(codes.InvalidArgument, "clusternet.prefixes[%d].cidr (%s) must be in canonical format (ie. all address bits within the subnet must be zero)", i, p.String())
		}

		// Make sure they're fully contained within clusterNet or are the /32 of a node's
		// externalIP.

		okay := false
		if clusterNet.Contains(p.Addr()) && p.Bits() >= clusterNet.Bits() {
			okay = true
		}
		if p.IsSingleIP() && p.Addr().String() == externalIP {
			okay = true
		}

		if !okay {
			return nil, status.Errorf(codes.InvalidArgument, "clusternet.prefixes[%d].cidr (%s) must be fully contained within cluster network (%s) or be the node's external IP (%s)", i, p.String(), clusterNet.String(), externalIP)
		}

		prefixes = append(prefixes, p)

	}

	// Modify and save node.
	node.wireguardKey = key.String()
	node.networkPrefixes = prefixes
	if err := nodeSave(ctx, l.leadership, node); err != nil {
		return nil, err
	}

	// Now that etcd is saved, also modify our cache.
	l.ls.clusternetCache[key.String()] = id

	return &ipb.UpdateNodeClusterNetworkingResponse{}, nil
}

func (l *curatorLeader) GetCACertificate(ctx context.Context, _ *ipb.GetCACertificateRequest) (*ipb.GetCACertificateResponse, error) {
	return &ipb.GetCACertificateResponse{
		IdentityCaCertificate: l.node.ClusterCA().Raw,
	}, nil
}
