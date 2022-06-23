package curator

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"sort"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	dpb "google.golang.org/protobuf/types/known/durationpb"

	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
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

	// HeartbeatPeriod is the duration between consecutive heartbeat update
	// messages sent by the node.
	HeartbeatInterval = time.Second * 5

	// HeartbeatTimeout is the duration after which a node is considered to be
	// timing out, given no recent heartbeat updates were received by the leader.
	HeartbeatTimeout = HeartbeatInterval * 2
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
			rpc.Trace(ctx).Printf("Unmarshalling node %q failed: %v", kv.Value, err)
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

// nodeHeartbeatTimestamp returns the node nid's last heartbeat timestamp, as
// seen from the Curator leader's perspective. If no heartbeats were received
// from the node, a zero time.Time value is returned.
func (l *leaderManagement) nodeHeartbeatTimestamp(nid string) time.Time {
	smv, ok := l.ls.heartbeatTimestamps.Load(nid)
	if ok {
		return smv.(time.Time)
	}
	return time.Time{}
}

// nodeHealth returns the node's health, along with the duration since last
// heartbeat was received, given a current timestamp.
func (l *leaderManagement) nodeHealth(node *Node, now time.Time) (apb.Node_Health, time.Duration) {
	// Get the last received node heartbeat's timestamp.
	nid := identity.NodeID(node.pubkey)
	nts := l.nodeHeartbeatTimestamp(nid)
	// lhb is the duration since the last heartbeat was received.
	lhb := now.Sub(nts)
	// Determine the node's health based on the heartbeat timestamp.
	var nh apb.Node_Health
	if node.state == cpb.NodeState_NODE_STATE_UP {
		// Only UP nodes can send heartbeats.
		switch {
		// If no heartbeats were received, but the leadership has only just
		// started, the node's health is unknown.
		case nts.IsZero() && (now.Sub(l.ls.startTs) < HeartbeatTimeout):
			nh = apb.Node_UNKNOWN
		// If the leader had received heartbeats from the node, but the last
		// heartbeat is stale, the node is timing out.
		case lhb > HeartbeatTimeout:
			nh = apb.Node_HEARTBEAT_TIMEOUT
		// Otherwise, the node can be declared healthy.
		default:
			nh = apb.Node_HEALTHY
		}
	} else {
		// Since node isn't UP, its health is unknown. Non-UP nodes can't access
		// the heartbeat RPC.
		nh = apb.Node_UNKNOWN
	}
	return nh, lhb
}

// GetNodes implements Management.GetNodes, which returns a list of nodes from
// the point of view of the cluster.
func (l *leaderManagement) GetNodes(req *apb.GetNodesRequest, srv apb.Management_GetNodesServer) error {
	ctx := srv.Context()

	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Retrieve all nodes from etcd in a single Get call.
	res, err := l.txnAsLeader(ctx, nodeEtcdPrefix.Range())
	if err != nil {
		return status.Errorf(codes.Unavailable, "could not retrieve list of nodes: %v", err)
	}

	// Create a CEL filter program, to be used in the reply loop below.
	filter, err := buildNodeFilter(ctx, req.Filter)
	if err != nil {
		return err
	}

	// Get a singular monotonic timestamp to reference node heartbeat timestamps
	// against.
	now := time.Now()

	// Convert etcd data into proto nodes, send one streaming response for each
	// node.
	kvs := res.Responses[0].GetResponseRange().Kvs
	for _, kv := range kvs {
		node, err := nodeUnmarshal(kv.Value)
		if err != nil {
			rpc.Trace(ctx).Printf("Unmarshalling node %q failed: %v", kv.Value, err)
			continue
		}

		// Convert node roles.
		roles := &cpb.NodeRoles{}
		if node.kubernetesWorker != nil {
			roles.KubernetesWorker = &cpb.NodeRoles_KubernetesWorker{}
		}
		if node.consensusMember != nil {
			roles.ConsensusMember = &cpb.NodeRoles_ConsensusMember{}
		}

		// Assess the node's health.
		health, lhb := l.nodeHealth(node, now)

		entry := apb.Node{
			Pubkey:             node.pubkey,
			State:              node.state,
			Status:             node.status,
			Roles:              roles,
			TimeSinceHeartbeat: dpb.New(lhb),
			Health:             health,
		}

		// Evaluate the filter expression for this node. Send the node, if it's
		// kept by the filter.
		keep, err := filter(ctx, &entry)
		if err != nil {
			return err
		}
		if !keep {
			continue
		}
		if err := srv.Send(&entry); err != nil {
			return err
		}
	}
	return nil
}

func (l *leaderManagement) ApproveNode(ctx context.Context, req *apb.ApproveNodeRequest) (*apb.ApproveNodeResponse, error) {
	// MVP: check if policy allows for this node to be approved for this cluster.
	// This should happen automatically, if possible, via hardware attestation
	// against policy, not manually.

	// Ensure the given key resembles a public key before using it to generate
	// a node iD. This key is then used to craft an arbitrary etcd path, so
	// let's do an early check in case the user set something that's obviously
	// not a public key.
	if len(req.Pubkey) != ed25519.PublicKeySize {
		return nil, status.Errorf(codes.InvalidArgument, "pubkey must be %d bytes long", ed25519.PublicKeySize)
	}

	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Find node for this pubkey.
	id := identity.NodeID(req.Pubkey)
	node, err := nodeLoad(ctx, l.leadership, id)
	if err != nil {
		return nil, err
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

	// Set node to be STANDBY.
	node.state = cpb.NodeState_NODE_STATE_STANDBY
	if err := nodeSave(ctx, l.leadership, node); err != nil {
		return nil, err
	}

	return &apb.ApproveNodeResponse{}, nil
}

// UpdateNodeRoles implements Management.UpdateNodeRoles, which in addition to
// adjusting the affected node's representation within the cluster, can also
// trigger the addition of a new etcd learner node.
func (l *leaderManagement) UpdateNodeRoles(ctx context.Context, req *apb.UpdateNodeRolesRequest) (*apb.UpdateNodeRolesResponse, error) {
	if len(req.Pubkey) != ed25519.PublicKeySize {
		return nil, status.Errorf(codes.InvalidArgument, "pubkey must be %d bytes long", ed25519.PublicKeySize)
	}

	// Take l.muNodes before modifying the node.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Find the node matching the requested public key.
	id := identity.NodeID(req.Pubkey)
	node, err := nodeLoad(ctx, l.leadership, id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "while loading node %s: %v", id, err)
	}

	// Adjust each role, if a corresponding value is set within the request. Do
	// nothing, if the role is already matches the requested value.

	if req.ConsensusMember != nil {
		if *req.ConsensusMember {
			// Add a new etcd learner node.
			w := l.consensus.Watch()
			defer w.Close()

			st, err := w.GetRunning(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Unavailable, "could not get running consensus: %v", err)
			}

			join, err := st.AddNode(ctx, node.pubkey)
			if err != nil {
				return nil, status.Errorf(codes.Unavailable, "could not add node: %v", err)
			}

			// Modify the node's state to reflect the change.
			node.EnableConsensusMember(join)
		} else {
			node.DisableConsensusMember()
		}
	}

	if req.KubernetesWorker != nil {
		if *req.KubernetesWorker {
			if node.consensusMember == nil {
				return nil, status.Errorf(codes.InvalidArgument, "could not set role: Kubernetes worker nodes must also be consensus members.")
			}
			node.EnableKubernetesWorker()
		} else {
			node.DisableKubernetesWorker()
		}
	}

	if err := nodeSave(ctx, l.leadership, node); err != nil {
		return nil, err
	}
	return &apb.UpdateNodeRolesResponse{}, nil
}
