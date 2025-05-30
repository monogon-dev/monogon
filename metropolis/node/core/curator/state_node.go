// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package curator

import (
	"context"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net/netip"
	"sort"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/osbase/pki"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// Node is a Metropolis cluster member. A node is a virtual or physical machine
// running Metropolis. This object represents a node only as part of a cluster.
// A machine running Metropolis that is not yet (attempting to be) part of a
// cluster is not considered a Node.
//
// This object is used internally within the curator code. Curator clients do
// not have access to this object and instead rely on protobuf representations
// of objects from the Curator gRPC API. An exception is the cluster bootstrap
// code which needs to bring up a new curator from scratch alongside the rest of
// the cluster.
type Node struct {
	// clusterUnlockKey is half of the unlock key required to mount the node's
	// data partition. It's stored in etcd, and will only be provided to the
	// Node if it can prove its identity via an integrity mechanism (ie. via
	// TPM), or when the Node was just created (as the key is generated locally
	// by localstorage on first format/mount).
	//
	// The other part of the unlock key is the LocalUnlockKey that's present on the
	// node's ESP partition.
	clusterUnlockKey []byte

	// id is the Node's ID.
	id string

	// pubkey is the ED25519 public key corresponding to the node's private key
	// which it stores on its local data partition. The private part of the key
	// never leaves the node.
	//
	// The public key is used to generate the Node's canonical ID.
	pubkey []byte

	// jkey is the node's ED25519 public Join Key. The private part of the key
	// never leaves the node. The key is generated by the node and passed to
	// Curator during the registration process.
	jkey []byte

	// state is the state of this node as seen from the point of view of the
	// cluster. See //metropolis/proto:common.proto for more information.
	state cpb.NodeState

	status *cpb.NodeStatus

	tpmUsage cpb.NodeTPMUsage

	// A Node can have multiple Roles. Each Role is represented by the presence
	// of NodeRole* structures in this structure, with a nil pointer
	// representing the lack of a role.

	consensusMember *NodeRoleConsensusMember

	// kubernetesController is set if this node is a Kubernetes controller, ie.
	// running the Kubernetes control plane.
	kubernetesController *NodeRoleKubernetesController

	// kubernetesWorker is set if this node is a Kubernetes worker, ie. running the
	// Kubernetes dataplane which runs user workloads.
	kubernetesWorker *NodeRoleKubernetesWorker

	// wireguardKey, if set, is the Wireguard key of the node's cluster networking
	// setup.
	wireguardKey string
	// networkingPrefixes are all the network routes exported by the node into the
	// cluster networking mesh. All of them will be programmed as allowedIPs into a
	// wireguard peer, but only the pod network will have a single large route
	// installed into the host routing table.
	networkPrefixes []netip.Prefix

	labels map[string]string
}

type NewNodeData struct {
	CUK      []byte
	ID       string
	Pubkey   []byte
	JPub     []byte
	TPMUsage cpb.NodeTPMUsage
	Labels   map[string]string
}

// NewNodeForBootstrap creates a brand new node without regard for any other
// cluster state.
//
// This can only be used by the cluster bootstrap logic.
func NewNodeForBootstrap(n *NewNodeData) Node {
	return Node{
		clusterUnlockKey: n.CUK,
		id:               n.ID,
		pubkey:           n.Pubkey,
		jkey:             n.JPub,
		state:            cpb.NodeState_NODE_STATE_UP,
		tpmUsage:         n.TPMUsage,
		labels:           n.Labels,
	}
}

// NodeRoleKubernetesController defines that the Node should be running the
// Kubernetes control plane.
type NodeRoleKubernetesController struct {
}

// NodeRoleKubernetesWorker defines that the Node should be running the
// Kubernetes data plane.
type NodeRoleKubernetesWorker struct {
}

// NodeRoleConsensusMember defines that the Node should be running a
// consensus/etcd instance.
type NodeRoleConsensusMember struct {
	// CACertificate, PeerCertificate are the X509 certificates to be used by the
	// node's etcd member to serve peer traffic.
	CACertificate, PeerCertificate *x509.Certificate
	// CRL is an initial certificate revocation list that the etcd member should
	// start with.
	//
	// TODO(q3k): don't store this in etcd like that, instead have the node retrieve
	// an initial CRL using gRPC/Curator.Watch.
	CRL *pki.CRL

	// Peers are a list of etcd members that the node's etcd member should attempt
	// to connect to.
	//
	// TODO(q3k): don't store this in etcd like that, instead have this be
	// dynamically generated at time of retrieval.
	Peers []NodeRoleConsensusMemberPeer
}

// NodeRoleConsensusMemberPeer is a name/URL pair pointing to an etcd member's
// peer listener.
type NodeRoleConsensusMemberPeer struct {
	// Name is the name of the etcd member, equal to the Metropolis node's ID that
	// the etcd member is running on.
	Name string
	// URL is a https://host:port string that can be passed to etcd on startup.
	URL string
}

// ID returns the name of this node.
func (n *Node) ID() string {
	return n.id
}

func (n *Node) String() string {
	return n.id
}

// KubernetesWorker returns a copy of the NodeRoleKubernetesWorker struct if
// the Node is a kubernetes worker, otherwise nil.
func (n *Node) KubernetesWorker() *NodeRoleKubernetesWorker {
	if n.kubernetesWorker == nil {
		return nil
	}
	kw := *n.kubernetesWorker
	return &kw
}

func (n *Node) EnableKubernetesWorker() {
	n.kubernetesWorker = &NodeRoleKubernetesWorker{}
}

func (n *Node) DisableKubernetesWorker() {
	n.kubernetesWorker = nil
}

func (n *Node) KubernetesController() *NodeRoleKubernetesController {
	if n.kubernetesController == nil {
		return nil
	}
	kcp := *n.kubernetesController
	return &kcp
}

func (n *Node) EnableKubernetesController() {
	n.kubernetesController = &NodeRoleKubernetesController{}
}

func (n *Node) DisableKubernetesController() {
	n.kubernetesController = nil
}

func (n *Node) EnableConsensusMember(jc *consensus.JoinCluster) {
	peers := make([]NodeRoleConsensusMemberPeer, len(jc.ExistingNodes))
	for i, n := range jc.ExistingNodes {
		peers[i].Name = n.Name
		peers[i].URL = n.URL
	}
	n.consensusMember = &NodeRoleConsensusMember{
		CACertificate:   jc.CACertificate,
		PeerCertificate: jc.NodeCertificate,
		Peers:           peers,
		CRL:             jc.InitialCRL,
	}
}

func (n *Node) DisableConsensusMember() {
	n.consensusMember = nil
}

var (
	// NodeEtcdPrefix is an etcd key prefix preceding cluster member node IDs,
	// mapping to ppb.Node values.
	NodeEtcdPrefix = mustNewEtcdPrefix("/nodes/")
	// joinCredPrefix is an etcd key prefix preceding hex-encoded cluster member
	// node join keys, mapping to node IDs.
	joinCredPrefix = mustNewEtcdPrefix("/join_keys/")
)

// etcdNodePath builds the etcd path in which this node's protobuf-serialized
// state is stored in etcd.
func (n *Node) etcdNodePath() (string, error) {
	return NodeEtcdPrefix.Key(n.ID())
}

func (n *Node) etcdJoinKeyPath() (string, error) {
	return joinCredPrefix.Key(hex.EncodeToString(n.jkey))
}

// proto serializes the Node object into protobuf, to be used for saving to
// etcd.
func (n *Node) proto() *ppb.Node {
	msg := &ppb.Node{
		Id:               n.id,
		ClusterUnlockKey: n.clusterUnlockKey,
		PublicKey:        n.pubkey,
		JoinKey:          n.jkey,
		FsmState:         n.state,
		Roles:            &cpb.NodeRoles{},
		Status:           n.status,
		TpmUsage:         n.tpmUsage,
		Labels:           &cpb.NodeLabels{},
	}
	if n.kubernetesWorker != nil {
		msg.Roles.KubernetesWorker = &cpb.NodeRoles_KubernetesWorker{}
	}
	if n.kubernetesController != nil {
		msg.Roles.KubernetesController = &cpb.NodeRoles_KubernetesController{}
	}
	if n.consensusMember != nil {
		peers := make([]*cpb.NodeRoles_ConsensusMember_Peer, len(n.consensusMember.Peers))
		for i, p := range n.consensusMember.Peers {
			peers[i] = &cpb.NodeRoles_ConsensusMember_Peer{
				Name: p.Name,
				Url:  p.URL,
			}
		}
		msg.Roles.ConsensusMember = &cpb.NodeRoles_ConsensusMember{
			CaCertificate:   n.consensusMember.CACertificate.Raw,
			PeerCertificate: n.consensusMember.PeerCertificate.Raw,
			InitialCrl:      n.consensusMember.CRL.Raw,
			Peers:           peers,
		}
	}
	if n.wireguardKey != "" {
		var prefixes []*cpb.NodeClusterNetworking_Prefix
		for _, prefix := range n.networkPrefixes {
			prefixes = append(prefixes, &cpb.NodeClusterNetworking_Prefix{
				Cidr: prefix.String(),
			})
		}
		msg.Clusternet = &cpb.NodeClusterNetworking{
			WireguardPubkey: n.wireguardKey,
			Prefixes:        prefixes,
		}
	}
	for k, v := range n.labels {
		msg.Labels.Pairs = append(msg.Labels.Pairs, &cpb.NodeLabels_Pair{
			Key:   k,
			Value: v,
		})
	}
	sort.Slice(msg.Labels.Pairs, func(i, j int) bool {
		return msg.Labels.Pairs[i].Key < msg.Labels.Pairs[j].Key
	})
	return msg
}

func nodeUnmarshal(kv *mvccpb.KeyValue) (*Node, error) {
	id := NodeEtcdPrefix.ExtractID(string(kv.Key))
	if id == "" {
		return nil, fmt.Errorf("invalid node key %q", kv.Key)
	}
	var msg ppb.Node
	if err := proto.Unmarshal(kv.Value, &msg); err != nil {
		return nil, fmt.Errorf("could not unmarshal proto of node %s: %w", id, err)
	}
	valueID := msg.Id
	if valueID == "" {
		// Backward compatibility
		valueID = identity.NodeID(msg.PublicKey)
	}
	if id != valueID {
		return nil, fmt.Errorf("node ID mismatch (etcd key: %q, value: %q)", id, valueID)
	}
	n := &Node{
		clusterUnlockKey: msg.ClusterUnlockKey,
		id:               id,
		pubkey:           msg.PublicKey,
		jkey:             msg.JoinKey,
		state:            msg.FsmState,
		status:           msg.Status,
		tpmUsage:         msg.TpmUsage,
		labels:           make(map[string]string),
	}
	if msg.Roles.KubernetesWorker != nil {
		n.kubernetesWorker = &NodeRoleKubernetesWorker{}
	}
	if msg.Roles.KubernetesController != nil {
		n.kubernetesController = &NodeRoleKubernetesController{}
	}
	if cm := msg.Roles.ConsensusMember; cm != nil {
		caCert, err := x509.ParseCertificate(cm.CaCertificate)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal consensus ca certificate: %w", err)
		}
		peerCert, err := x509.ParseCertificate(cm.PeerCertificate)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal consensus peer certificate: %w", err)
		}
		crl, err := x509.ParseCRL(cm.InitialCrl)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal consensus crl: %w", err)
		}
		var peers []NodeRoleConsensusMemberPeer
		for _, p := range cm.Peers {
			peers = append(peers, NodeRoleConsensusMemberPeer{
				Name: p.Name,
				URL:  p.Url,
			})
		}
		n.consensusMember = &NodeRoleConsensusMember{
			CACertificate:   caCert,
			PeerCertificate: peerCert,
			CRL: &pki.CRL{
				Raw:  cm.InitialCrl,
				List: crl,
			},
			Peers: peers,
		}
	}
	if cn := msg.Clusternet; cn != nil {
		n.wireguardKey = cn.WireguardPubkey
		for _, prefix := range cn.Prefixes {
			if prefix.Cidr == "" {
				continue
			}
			nip, err := netip.ParsePrefix(prefix.Cidr)
			if err != nil {
				// Eat error. When we serialize this back into a node, the invalid record will
				// just be removed.
				continue
			}
			n.networkPrefixes = append(n.networkPrefixes, nip)
		}
	}
	if l := msg.Labels; l != nil {
		for _, pair := range l.Pairs {
			// Skip invalid keys/values that were somehow persisted into etcd. They will be
			// removed on next marshal/save.
			if err := common.ValidateLabelKey(pair.Key); err != nil {
				continue
			}
			if err := common.ValidateLabelValue(pair.Value); err != nil {
				continue
			}
			n.labels[pair.Key] = pair.Value
		}
	}
	return n, nil
}

var (
	errNodeNotFound = status.Error(codes.NotFound, "node not found")
)

// nodeLoad attempts to load a node by ID from etcd, within a given active
// leadership. All returned errors are gRPC statuses that are safe to return to
// untrusted callers. If the given node is not found, errNodeNotFound will be
// returned.
func nodeLoad(ctx context.Context, l *leadership, id string) (*Node, error) {
	rpc.Trace(ctx).Printf("loadNode(%s)...", id)
	key, err := NodeEtcdPrefix.Key(id)
	if err != nil {
		rpc.Trace(ctx).Printf("invalid node id: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid node id")
	}
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(key))
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return nil, rpcErr
		}
		rpc.Trace(ctx).Printf("could not retrieve node %s: %v", id, err)
		return nil, status.Errorf(codes.Unavailable, "could not retrieve node %s: %v", id, err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	rpc.Trace(ctx).Printf("loadNode(%s): %d KVs", id, len(kvs))
	if len(kvs) != 1 {
		return nil, errNodeNotFound
	}
	node, err := nodeUnmarshal(kvs[0])
	if err != nil {
		rpc.Trace(ctx).Printf("could not unmarshal node: %v", err)
		return nil, status.Errorf(codes.Unavailable, "could not unmarshal node")
	}
	rpc.Trace(ctx).Printf("loadNode(%s): unmarshal ok", id)
	return node, nil
}

// nodeSave attempts to save a node into etcd, within a given active leadership.
// All returned errors are gRPC statuses that safe to return to untrusted callers.
func nodeSave(ctx context.Context, l *leadership, n *Node) error {
	// Build an etcd operation to save the node with a key based on its ID.
	id := n.ID()
	rpc.Trace(ctx).Printf("nodeSave(%s)...", id)
	nkey, err := NodeEtcdPrefix.Key(id)
	if err != nil {
		rpc.Trace(ctx).Printf("invalid node id: %v", err)
		return status.Errorf(codes.InvalidArgument, "invalid node id")
	}
	nodeBytes, err := proto.Marshal(n.proto())
	if err != nil {
		rpc.Trace(ctx).Printf("could not marshal updated node: %v", err)
		return status.Errorf(codes.Unavailable, "could not marshal updated node")
	}
	ons := clientv3.OpPut(nkey, string(nodeBytes))

	// Build an etcd operation to map the node's Join Key into its ID for use in
	// Join Flow.
	jkey, err := n.etcdJoinKeyPath()
	if err != nil {
		// This should never happen.
		rpc.Trace(ctx).Printf("invalid join key representation: %v", err)
		return status.Errorf(codes.InvalidArgument, "invalid join key representation")
	}
	// TODO(mateusz@monogon.tech): ensure that if the join key index already
	// exists, it points to the node we're saving. Refuse to save/update the
	// node if it doesn't.
	oks := clientv3.OpPut(jkey, id)

	// Execute both operations atomically.
	_, err = l.txnAsLeader(ctx, ons, oks)
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return rpcErr
		}
		rpc.Trace(ctx).Printf("could not save updated node: %v", err)
		return status.Error(codes.Unavailable, "could not save updated node")
	}
	rpc.Trace(ctx).Printf("nodeSave(%s): write ok", id)
	return nil
}

// nodeDestroy removes all traces of a node from etcd. It does not first check
// whether the node is safe to be removed.
func nodeDestroy(ctx context.Context, l *leadership, n *Node) error {
	// Build an etcd operation to save the node with a key based on its ID.
	id := n.ID()
	rpc.Trace(ctx).Printf("nodeDestroy(%s)...", id)

	// Get paths for node data and join key.
	nkey, err := NodeEtcdPrefix.Key(id)
	if err != nil {
		rpc.Trace(ctx).Printf("invalid node id: %v", err)
		return status.Errorf(codes.InvalidArgument, "invalid node id")
	}
	jkey, err := n.etcdJoinKeyPath()
	if err != nil {
		// This should never happen.
		rpc.Trace(ctx).Printf("invalid join key representation: %v", err)
		return status.Errorf(codes.InvalidArgument, "invalid join key representation")
	}
	// Delete both.
	_, err = l.txnAsLeader(ctx,
		clientv3.OpDelete(nkey),
		clientv3.OpDelete(jkey),
	)
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return rpcErr
		}
		rpc.Trace(ctx).Printf("could not destroy node: %v", err)
		return status.Error(codes.Unavailable, "could not destroy node")
	}

	// TODO(q3k): remove node's data from PKI.

	rpc.Trace(ctx).Printf("nodeDestroy(%s): destroy ok", id)
	return nil
}

// nodeIdByJoinKey attempts to fetch a Node ID corresponding to the given Join
// Key from etcd, within a given active leadership. All returned errors are
// gRPC statuses that are safe to return to untrusted callers. If the given
// Join Key is not found, errNodeNotFound will be returned along with an empty
// string.
func nodeIdByJoinKey(ctx context.Context, l *leadership, jkey []byte) (string, error) {
	if len(jkey) == 0 {
		return "", status.Errorf(codes.InvalidArgument, "join key is empty")
	}

	cred := hex.EncodeToString(jkey)
	key, err := joinCredPrefix.Key(cred)
	if err != nil {
		// This should never happen.
		rpc.Trace(ctx).Printf("invalid join key representation: %v", err)
		return "", status.Errorf(codes.InvalidArgument, "invalid join key representation")
	}
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(key))
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return "", rpcErr
		}
		return "", status.Errorf(codes.Unavailable, "could not retrieve node id matching join key %s: %v", cred, err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	if len(kvs) != 1 {
		return "", errNodeNotFound
	}
	id := string(kvs[0].Value[:])
	return id, nil
}
