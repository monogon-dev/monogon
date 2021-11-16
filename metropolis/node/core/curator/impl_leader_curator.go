package curator

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/event/etcd"
	"source.monogon.dev/metropolis/pkg/pki"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// leaderCurator implements the Curator gRPC API (ipb.Curator) as a curator
// leader.
type leaderCurator struct {
	*leadership
}

// Watch returns a stream of updates concerning some part of the cluster
// managed by the curator.
//
// See metropolis.node.core.curator.proto.api.Curator for more information about
// the RPC semantics.
//
// TODO(q3k): Currently the watch RPCs are individually backed by etcd cluster
// watches (via individual etcd event values), which might be problematic in
// case of a significant amount of parallel Watches being issued to the Curator.
// It might make sense to combine all pending Watch requests into a single watch
// issued to the cluster, with an intermediary caching stage within the curator
// instance. However, that is effectively implementing etcd learner/relay logic,
// which has has to be carefully considered, especially with regards to serving
// stale data.
func (l *leaderCurator) Watch(req *ipb.WatchRequest, srv ipb.Curator_WatchServer) error {
	switch x := req.Kind.(type) {
	case *ipb.WatchRequest_NodeInCluster_:
		return l.watchNodeInCluster(x.NodeInCluster, srv)
	case *ipb.WatchRequest_NodesInCluster_:
		return l.watchNodesInCluster(x.NodesInCluster, srv)
	default:
		return status.Error(codes.Unimplemented, "unsupported watch kind")
	}
}

// watchNodeInCluster implements the Watch API when dealing with a single
// node-in-cluster request. Effectively, it pipes an etcd value watcher into the
// Watch API.
func (l *leaderCurator) watchNodeInCluster(nic *ipb.WatchRequest_NodeInCluster, srv ipb.Curator_WatchServer) error {
	ctx := srv.Context()

	// Constructing arbitrary etcd path: this is okay, as we only have node objects
	// underneath the NodeEtcdPrefix. Worst case an attacker can do is request a node
	// that doesn't exist, and that will just hang . All access is privileged, so
	// there's also no need to filter anything.
	nodePath, err := nodeEtcdPrefix.Key(nic.NodeId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid node name: %v", err)
	}
	value := etcd.NewValue(l.etcd, nodePath, nodeValueConverter)

	w := value.Watch()
	defer w.Close()

	for {
		v, err := w.Get(ctx)
		if err != nil {
			if rpcErr, ok := rpcError(err); ok {
				return rpcErr
			}
			rpc.Trace(ctx).Printf("etcd watch failed: %v", err)
			return status.Error(codes.Unavailable, "internal error")
		}

		ev := &ipb.WatchEvent{}
		nodeKV := v.(nodeAtID)
		nodeKV.appendToEvent(ev)
		if err := srv.Send(ev); err != nil {
			return err
		}
	}
}

// watchNodesInCluster implements the Watch API when dealing with a
// all-nodes-in-cluster request. Effectively, it pipes a ranged etcd value
// watcher into the Watch API.
func (l *leaderCurator) watchNodesInCluster(_ *ipb.WatchRequest_NodesInCluster, srv ipb.Curator_WatchServer) error {
	ctx := srv.Context()

	start, end := nodeEtcdPrefix.KeyRange()
	value := etcd.NewValue(l.etcd, start, nodeValueConverter, etcd.Range(end))

	w := value.Watch()
	defer w.Close()

	// Perform initial fetch from etcd.
	nodes := make(map[string]*Node)
	for {
		v, err := w.Get(ctx, etcd.BacklogOnly)
		if err == etcd.BacklogDone {
			break
		}
		if err != nil {
			rpc.Trace(ctx).Printf("etcd watch failed (initial fetch): %v", err)
			return status.Error(codes.Unavailable, "internal error during initial fetch")
		}
		nodeKV := v.(nodeAtID)
		if nodeKV.value != nil {
			nodes[nodeKV.id] = nodeKV.value
		}
	}

	// Initial send, chunked to not go over 2MiB (half of the default gRPC message
	// size limit).
	//
	// TODO(q3k): formalize message limits, set const somewhere.
	we := &ipb.WatchEvent{}
	for _, n := range nodes {
		we.Nodes = append(we.Nodes, &ipb.Node{
			Id:     n.ID(),
			Roles:  n.proto().Roles,
			Status: n.status,
		})
		if proto.Size(we) > (2 << 20) {
			if err := srv.Send(we); err != nil {
				return err
			}
			we = &ipb.WatchEvent{}
		}
	}
	// Send last update message. This might be empty, but we need to send the
	// LAST_BACKLOGGED marker.
	we.Progress = ipb.WatchEvent_PROGRESS_LAST_BACKLOGGED
	if err := srv.Send(we); err != nil {
		return err
	}

	// Send updates as they arrive from etcd watcher.
	for {
		v, err := w.Get(ctx)
		if err != nil {
			rpc.Trace(ctx).Printf("etcd watch failed (update): %v", err)
			return status.Errorf(codes.Unavailable, "internal error during update")
		}
		we := &ipb.WatchEvent{}
		nodeKV := v.(nodeAtID)
		nodeKV.appendToEvent(we)
		if err := srv.Send(we); err != nil {
			return err
		}
	}
}

// nodeAtID is a key/pair container for a node update received from an etcd
// watcher. The value will be nil if this update represents a node being
// deleted.
type nodeAtID struct {
	id    string
	value *Node
}

// nodeValueConverter is called by etcd node value watchers to convert updates
// from the cluster into nodeAtID, ensuring data integrity and checking
// invariants.
func nodeValueConverter(key, value []byte) (interface{}, error) {
	res := nodeAtID{
		id: nodeEtcdPrefix.ExtractID(string(key)),
	}
	if len(value) > 0 {
		node, err := nodeUnmarshal(value)
		if err != nil {
			return nil, err
		}
		res.value = node
		if res.id != res.value.ID() {
			return nil, fmt.Errorf("node ID mismatch (etcd key: %q, value: %q)", res.id, res.value.ID())
		}
	}
	if res.id == "" {
		// This shouldn't happen, to the point where this might be better handled by a
		// panic.
		return nil, fmt.Errorf("invalid node key %q", key)
	}
	return res, nil
}

// appendToId records a node update represented by nodeAtID into a Curator
// WatchEvent, either a Node or NodeTombstone.
func (kv nodeAtID) appendToEvent(ev *ipb.WatchEvent) {
	if node := kv.value; node != nil {
		ev.Nodes = append(ev.Nodes, &ipb.Node{
			Id:     node.ID(),
			Roles:  node.proto().Roles,
			Status: node.status,
		})
	} else {
		ev.NodeTombstones = append(ev.NodeTombstones, &ipb.WatchEvent_NodeTombstone{
			NodeId: kv.id,
		})
	}
}

// UpdateNodeStatus is called by nodes in the cluster to report their own
// status. This status is recorded by the curator and can be retrieed via
// Watch.
func (l *leaderCurator) UpdateNodeStatus(ctx context.Context, req *ipb.UpdateNodeStatusRequest) (*ipb.UpdateNodeStatusResponse, error) {
	// Ensure that the given node_id matches the calling node. We currently
	// only allow for direct self-reporting of status by nodes.
	pi := rpc.GetPeerInfo(ctx)
	if pi == nil || pi.Node == nil {
		return nil, status.Error(codes.PermissionDenied, "only nodes can update node status")
	}
	id := identity.NodeID(pi.Node.PublicKey)
	if id != req.NodeId {
		return nil, status.Errorf(codes.PermissionDenied, "node %q cannot update the status of node %q", id, req.NodeId)
	}

	// Verify sent status. Currently we assume the entire status must be set at
	// once, and cannot be unset.
	if req.Status == nil || req.Status.ExternalAddress == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Status and Status.ExternalAddress must be set")
	}

	if net.ParseIP(req.Status.ExternalAddress) == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Status.ExternalAddress must be a valid IP address")
	}

	// As we're performing a node update with two etcd transactions below (one
	// to retrieve, one to save and upate node), take a local lock to ensure
	// that we don't have a race between either two UpdateNodeStatus calls or
	// an UpdateNodeStatus call and some other mutation to the node store.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Retrieve node ...
	node, err := nodeLoad(ctx, l.leadership, id)
	if err != nil {
		return nil, err
	}
	// ... update its' status ...
	node.status = req.Status
	// ... and save it to etcd.
	if err := nodeSave(ctx, l.leadership, node); err != nil {
		return nil, err
	}

	return &ipb.UpdateNodeStatusResponse{}, nil
}

func (l *leaderCurator) RegisterNode(ctx context.Context, req *ipb.RegisterNodeRequest) (*ipb.RegisterNodeResponse, error) {
	// Call is unauthenticated - verify the other side has connected with an
	// ephemeral certificate. That certificate's pubkey will become the node's
	// pubkey.
	pi := rpc.GetPeerInfo(ctx)
	if pi == nil || pi.Unauthenticated == nil {
		return nil, status.Error(codes.Unauthenticated, "connection must be established with a self-signed ephemeral certificate")
	}
	pubkey := pi.Unauthenticated.SelfSignedPublicKey

	// Verify that call contains a RegisterTicket and that this RegisterTicket is
	// valid.
	wantTicket, err := l.ensureRegisterTicket(ctx)
	if err != nil {
		rpc.Trace(ctx).Printf("could not ensure register ticket: %v", err)
		return nil, status.Error(codes.Unavailable, "could not retrieve register ticket")
	}
	gotTicket := req.RegisterTicket
	if subtle.ConstantTimeCompare(wantTicket, gotTicket) != 1 {
		return nil, status.Error(codes.PermissionDenied, "registerticket invalid")
	}

	// Doing a read-then-write operation below, take lock.
	//
	// MVP: This can lock up the cluster if too many RegisterNode calls get issued,
	// we should either ratelimit these or remove the need to lock.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Check if there already is a node with this pubkey in the cluster.
	id := identity.NodeID(pubkey)
	node, err := nodeLoad(ctx, l.leadership, id)
	if err == nil {
		// If the existing node is in the NEW state already, there's nothing to do,
		// return no error. This can happen in case of spurious retries from the calling
		// node.
		if node.state == cpb.NodeState_NODE_STATE_NEW {
			return &ipb.RegisterNodeResponse{}, nil
		}
		// We can return a bit more information to the calling node here, as if it's in
		// possession of the private key corresponding to an existing node in the
		// cluster, it should have access to the status of the node without danger of
		// leaking data about other nodes.
		//
		rpc.Trace(ctx).Printf("node %s already exists in cluster, failing", id)
		return nil, status.Errorf(codes.FailedPrecondition, "node already exists in cluster, state %s", node.state.String())
	}
	if err != errNodeNotFound {
		return nil, err
	}

	// No node exists, create one.
	node = &Node{
		pubkey: pubkey,
		state:  cpb.NodeState_NODE_STATE_NEW,
	}
	if err := nodeSave(ctx, l.leadership, node); err != nil {
		return nil, err
	}
	return &ipb.RegisterNodeResponse{}, nil
}

func (l *leaderCurator) CommitNode(ctx context.Context, req *ipb.CommitNodeRequest) (*ipb.CommitNodeResponse, error) {
	// Call is unauthenticated - verify the other side has connected with an
	// ephemeral certificate. That certificate's pubkey will become the node's
	// pubkey.
	pi := rpc.GetPeerInfo(ctx)
	if pi == nil || pi.Unauthenticated == nil {
		return nil, status.Error(codes.Unauthenticated, "connection must be established with a self-signed ephemeral certificate")
	}
	pubkey := pi.Unauthenticated.SelfSignedPublicKey

	// Doing a read-then-write operation below, take lock.
	//
	// MVP: This can lock up the cluster if too many RegisterNode calls get issued,
	// we should either ratelimit these or remove the need to lock.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Retrieve the node and act on its current state, either returning early or
	// mutating it and continuing with the rest of the Commit logic.
	id := identity.NodeID(pubkey)
	node, err := nodeLoad(ctx, l.leadership, id)
	if err != nil {
		return nil, err
	}

	switch node.state {
	case cpb.NodeState_NODE_STATE_NEW:
		return nil, status.Error(codes.PermissionDenied, "node is NEW, wait for attestation/approval")
	case cpb.NodeState_NODE_STATE_DISOWNED:
		// This node has been since disowned by the cluster for some reason, the
		// register flow should be aborted.
		return nil, status.Error(codes.FailedPrecondition, "node is DISOWNED, abort register flow")
	case cpb.NodeState_NODE_STATE_UP:
		// This can happen due to a network failure when we already handled a
		// CommitNode, but we weren't able to respond to the user. CommitNode is
		// non-idempotent, so just abort, the node should retry from scratch and this
		// node should be manually disowned/deleted by system owners.
		return nil, status.Error(codes.FailedPrecondition, "node is already UP, abort register flow")
	case cpb.NodeState_NODE_STATE_STANDBY:
		// This is what we want.
	default:
		return nil, status.Errorf(codes.Internal, "node is in unknown state: %v", node.state)
	}

	// Check the given CUK is valid.
	// TODO(q3k): unify length with localstorage/crypt keySize.
	if want, got := 32, len(req.ClusterUnlockKey); want != got {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ClusterUnlockKey length, wanted %d bytes, got %d", want, got)
	}

	// Generate certificate for node, save new node state, return.

	// If this fails we are safe to let the client retry, as the PKI code is
	// idempotent.
	caCertBytes, err := pkiCA.Ensure(ctx, l.etcd)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not get CA certificate: %v", err)
	}
	nodeCert := &pki.Certificate{
		Namespace: &pkiNamespace,
		Issuer:    pkiCA,
		Template:  identity.NodeCertificate(node.pubkey),
		Mode:      pki.CertificateExternal,
		PublicKey: node.pubkey,
		Name:      fmt.Sprintf("node-%s", identity.NodeID(pubkey)),
	}
	nodeCertBytes, err := nodeCert.Ensure(ctx, l.etcd)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not emit node credentials: %v", err)
	}

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

	node.state = cpb.NodeState_NODE_STATE_UP
	node.clusterUnlockKey = req.ClusterUnlockKey
	node.EnableConsensusMember(join)
	node.EnableKubernetesWorker()
	if err := nodeSave(ctx, l.leadership, node); err != nil {
		return nil, err
	}

	// From this point on, any failure (in the server, or in the network, ...) dooms
	// the node from making progress in registering, as Commit is non-idempotent.

	return &ipb.CommitNodeResponse{
		CaCertificate:   caCertBytes,
		NodeCertificate: nodeCertBytes,
	}, nil
}
