package curator

import (
	"bytes"
	"context"
	"sort"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type leaderManagement struct {
	*leadership

	// node certificate on which leaderManagement runs. It's used by
	// GetClusterInformation which needs access to the CA pubkey.
	// Alternatively this could be stored in etcd, instead of being dependency
	// injected here.
	node *identity.Node
}

const (
	// registerTicketSize is the size, in bytes, of the RegisterTicket used to
	// perform early perimeter checks for nodes which wish to register into the
	// cluster.
	//
	// The size was picked to offer resistance against on-line bruteforcing attacks
	// in even the worst case scenario (no ratelimiting, no monitoring, zero latency
	// between attacker and cluster). 256 bits of entropy require 3.6e68 requests
	// per second to bruteforce the ticket within 10 years. The ticket doesn't need
	// to be manually copied by humans, so the relatively overkill size also doesn't
	// impact usability.
	registerTicketSize = 32
)

const (
	// registerTicketEtcdPath is the etcd key under which private.RegisterTicket is
	// stored.
	registerTicketEtcdPath = "/global/register_ticket"
)

func (l *leaderManagement) GetRegisterTicket(ctx context.Context, req *apb.GetRegisterTicketRequest) (*apb.GetRegisterTicketResponse, error) {
	ticket, err := l.ensureRegisterTicket(ctx)
	if err != nil {
		return nil, err
	}
	return &apb.GetRegisterTicketResponse{
		Ticket: ticket,
	}, nil
}

// GetClusterInfo implements Curator.GetClusterInfo, which returns summary
// information about the Metropolis cluster.
func (l *leaderManagement) GetClusterInfo(ctx context.Context, req *apb.GetClusterInfoRequest) (*apb.GetClusterInfoResponse, error) {
	res, err := l.txnAsLeader(ctx, nodeEtcdPrefix.Range())
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not retrieve list of nodes: %v", err)
	}

	// Sort nodes by public key, filter out Up, use top 15 in cluster directory
	// (limited to an arbitrary amount that doesn't overload callers with
	// unnecesssary information).
	//
	// MVP: this should be formalized and possibly re-designed/engineered.
	kvs := res.Responses[0].GetResponseRange().Kvs
	var nodes []*Node
	for _, kv := range kvs {
		node, err := nodeUnmarshal(kv.Value)
		if err != nil {
			// TODO(q3k): log this
			continue
		}
		if node.state != cpb.NodeState_NODE_STATE_UP {
			continue
		}
		nodes = append(nodes, node)
	}
	sort.Slice(nodes, func(i, j int) bool {
		return bytes.Compare(nodes[i].pubkey, nodes[j].pubkey) < 0
	})
	if len(nodes) > 15 {
		nodes = nodes[:15]
	}

	// Build cluster directory.
	directory := &cpb.ClusterDirectory{
		Nodes: make([]*cpb.ClusterDirectory_Node, len(nodes)),
	}
	for i, node := range nodes {
		var addresses []*cpb.ClusterDirectory_Node_Address
		if node.status != nil && node.status.ExternalAddress != "" {
			addresses = append(addresses, &cpb.ClusterDirectory_Node_Address{
				Host: node.status.ExternalAddress,
			})
		}
		directory.Nodes[i] = &cpb.ClusterDirectory_Node{
			PublicKey: node.pubkey,
			Addresses: addresses,
		}
	}

	return &apb.GetClusterInfoResponse{
		ClusterDirectory: directory,
		CaCertificate:    l.node.ClusterCA().Raw,
	}, nil
}
