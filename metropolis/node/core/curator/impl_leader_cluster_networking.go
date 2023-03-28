package curator

import (
	"context"
	"net/netip"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

func (l *leaderCurator) UpdateNodeClusterNetworking(ctx context.Context, req *ipb.UpdateNodeClusterNetworkingRequest) (*ipb.UpdateNodeClusterNetworkingResponse, error) {
	// Ensure that the given node_id matches the calling node. We currently
	// only allow for direct self-reporting of status by nodes.
	pi := rpc.GetPeerInfo(ctx)
	if pi == nil || pi.Node == nil {
		return nil, status.Error(codes.PermissionDenied, "only nodes can update node cluster networking")
	}
	id := identity.NodeID(pi.Node.PublicKey)

	if req.Clusternet == nil {
		return nil, status.Error(codes.InvalidArgument, "clusternet must be set")
	}
	cn := req.Clusternet
	if cn.WireguardPubkey == "" {
		return nil, status.Error(codes.InvalidArgument, "clusternet.wireguard_pubkey must be set")
	}
	_, err := wgtypes.ParseKey(cn.WireguardPubkey)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "clusternet.wireguard_pubkey must be a valid wireguard public key")
	}

	// TODO(q3k): unhardcode this and synchronize with Kubernetes code.
	clusterNet := netip.MustParsePrefix("10.0.0.0/16")

	// Update node with new clusternetworking data. We're doing a load/modify/store,
	// so lock here.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

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

	// ... update its' clusternetworking bits ...
	node.wireguardKey = cn.WireguardPubkey
	node.networkPrefixes = prefixes
	// ... and save it to etcd.
	if err := nodeSave(ctx, l.leadership, node); err != nil {
		return nil, err
	}

	return &ipb.UpdateNodeClusterNetworkingResponse{}, nil
}
