// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package curator

import (
	"context"
	"fmt"
	"net"
	"strings"

	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/pkg/supervisor"
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

	// pubkey is the ED25519 public key corresponding to the node's private key
	// which it stores on its local data partition. The private part of the key
	// never leaves the node.
	//
	// The public key is used to generate the Node's canonical ID.
	pubkey []byte

	// state is the state of this node as seen from the point of view of the
	// cluster. See //metropolis/proto:common.proto for more information.
	state cpb.NodeState

	// A Node can have multiple Roles. Each Role is represented by the presence
	// of NodeRole* structures in this structure, with a nil pointer
	// representing the lack of a role.

	// kubernetesWorker is set if this node is a Kubernetes worker, ie. running the
	// Kubernetes control plan and workload elements.
	// In the future, this will be split into a separate worker and control plane
	// role.
	kubernetesWorker *NodeRoleKubernetesWorker
}

// NewNodeForBootstrap creates a brand new node without regard for any other
// cluster state.
//
// This can only be used by the cluster bootstrap logic.
func NewNodeForBootstrap(cuk, pubkey []byte) Node {
	return Node{
		clusterUnlockKey: cuk,
		pubkey:           pubkey,
		state:            cpb.NodeState_NODE_STATE_UP,
		// TODO(q3k): make this configurable.
		kubernetesWorker: &NodeRoleKubernetesWorker{},
	}
}

// NodeRoleKubernetesWorker defines that the Node should be running the
// Kubernetes control and data plane.
type NodeRoleKubernetesWorker struct {
}

// ID returns the name of this node. See NodeID for more information.
func (n *Node) ID() string {
	return identity.NodeID(n.pubkey)
}

func (n *Node) String() string {
	return n.ID()
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

var (
	nodeEtcdPrefix = mustNewEtcdPrefix("/nodes/")
)

// etcdPath builds the etcd path in which this node's protobuf-serialized state
// is stored in etcd.
func (n *Node) etcdPath() (string, error) {
	return nodeEtcdPrefix.Key(n.ID())
}

// proto serializes the Node object into protobuf, to be used for saving to
// etcd.
func (n *Node) proto() *ppb.Node {
	msg := &ppb.Node{
		ClusterUnlockKey: n.clusterUnlockKey,
		PublicKey:        n.pubkey,
		FsmState:         n.state,
		Roles:            &cpb.NodeRoles{},
	}
	if n.kubernetesWorker != nil {
		msg.Roles.KubernetesWorker = &cpb.NodeRoles_KubernetesWorker{}
	}
	return msg
}

func nodeUnmarshal(data []byte) (*Node, error) {
	msg := ppb.Node{}
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("could not unmarshal proto: %w", err)
	}
	n := &Node{
		clusterUnlockKey: msg.ClusterUnlockKey,
		pubkey:           msg.PublicKey,
		state:            msg.FsmState,
	}
	if msg.Roles.KubernetesWorker != nil {
		n.kubernetesWorker = &NodeRoleKubernetesWorker{}
	}
	return n, nil
}

// ConfigureLocalHostname uses the node's ID as a hostname, and sets the
// current hostname, and local files like hosts and machine-id accordingly.
//
// TODO(q3k): move this to roleserver?
func (n *Node) ConfigureLocalHostname(ctx context.Context, ephemeral *localstorage.EphemeralDirectory, address net.IP) error {
	if err := unix.Sethostname([]byte(n.ID())); err != nil {
		return fmt.Errorf("failed to set runtime hostname: %w", err)
	}
	hosts := []string{
		"127.0.0.1 localhost",
		"::1 localhost",
		fmt.Sprintf("%s %s", address.String(), n.ID()),
	}
	if err := ephemeral.Hosts.Write([]byte(strings.Join(hosts, "\n")), 0644); err != nil {
		return fmt.Errorf("failed to write /ephemeral/hosts: %w", err)
	}
	if err := ephemeral.MachineID.Write([]byte(identity.NodeIDBare(n.pubkey)), 0644); err != nil {
		return fmt.Errorf("failed to write /ephemeral/machine-id: %w", err)
	}

	// Check that we are self-resolvable.
	ip, err := net.ResolveIPAddr("ip", n.ID())
	if err != nil {
		return fmt.Errorf("failed to self-resolve: %w", err)
	}
	supervisor.Logger(ctx).Infof("This is node %s at %v", n.ID(), ip)
	return nil
}
