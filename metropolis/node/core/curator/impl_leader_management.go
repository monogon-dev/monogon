package curator

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"sort"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

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

// GetClusterInfo implements Management.GetClusterInfo, which returns summary
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
			// TODO(issues/85): log this
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

// GetNodes implements Management.GetNodes, which returns a list of nodes from
// the point of view of the cluster.
func (l *leaderManagement) GetNodes(_ *apb.GetNodesRequest, srv apb.Management_GetNodesServer) error {
	ctx := srv.Context()

	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Retrieve all nodes from etcd in a single Get call.
	res, err := l.txnAsLeader(ctx, nodeEtcdPrefix.Range())
	if err != nil {
		return status.Errorf(codes.Unavailable, "could not retrieve list of nodes: %v", err)
	}

	// Convert etcd data into proto nodes, send one streaming response for each
	// node.
	kvs := res.Responses[0].GetResponseRange().Kvs
	for _, kv := range kvs {
		node, err := nodeUnmarshal(kv.Value)
		if err != nil {
			// TODO(issues/85): log this
			continue
		}

		// Convert node roles.
		roles := &cpb.NodeRoles{}
		if node.kubernetesWorker != nil {
			roles.KubernetesWorker = &cpb.NodeRoles_KubernetesWorker{}
		}

		if err := srv.Send(&apb.Node{
			Pubkey: node.pubkey,
			State:  node.state,
			Status: node.status,
			Roles:  roles,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (l *leaderManagement) ApproveNode(ctx context.Context, req *apb.ApproveNodeRequest) (*apb.ApproveNodeResponse, error) {
	// MVP: check if policy allows for this node to be approved for this cluster.
	// This should happen automatically, if possible, via hardware attestation
	// against policy, not manually.

	if len(req.Pubkey) != ed25519.PublicKeySize {
		return nil, status.Errorf(codes.InvalidArgument, "pubkey must be %d bytes long", ed25519.PublicKeySize)
	}

	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Find node by pubkey/ID.
	id := identity.NodeID(req.Pubkey)
	key, err := nodeEtcdPrefix.Key(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "pubkey invalid: %v", err)
	}
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(key))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not retrieve node: %v", err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	if len(kvs) != 1 {
		return nil, status.Errorf(codes.NotFound, "node not found")
	}
	node, err := nodeUnmarshal(kvs[0].Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not deserialize node: %v", err)
	}

	// Ensure node is either UP/STANDBY (no-op) or NEW (set to STANDBY).
	switch node.state {
	case cpb.NodeState_NODE_STATE_UP, cpb.NodeState_NODE_STATE_STANDBY:
		// No-op for idempotency.
		return &apb.ApproveNodeResponse{}, nil
	case cpb.NodeState_NODE_STATE_NEW:
		// What we can act on.
	default:
		return nil, status.Errorf(codes.FailedPrecondition, "node in state %s cannot be approved", node.state)
	}

	node.state = cpb.NodeState_NODE_STATE_STANDBY
	nodeBytes, err := proto.Marshal(node.proto())
	if err != nil {
		// TODO(issues/85): log this
		return nil, status.Errorf(codes.Unavailable, "could not marshal updated node")
	}
	_, err = l.txnAsLeader(ctx, clientv3.OpPut(key, string(nodeBytes)))
	if err != nil {
		// TODO(issues/85): log this
		return nil, status.Error(codes.Unavailable, "could not save updated node")
	}

	return &apb.ApproveNodeResponse{}, nil
}
