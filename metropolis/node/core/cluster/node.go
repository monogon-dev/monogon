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

package cluster

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"strings"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/pkg/supervisor"
	ppb "source.monogon.dev/metropolis/proto/private"
)

// Node is a Metropolis cluster member. A node is a virtual or physical machine
// running Metropolis. This object represents a node only as part of a cluster
// - ie., this object will never be available outside of
// //metropolis/node/core/cluster if the Node is not part of a Cluster.  Nodes
// are inherently tied to their long term storage, which is etcd. As such,
// methods on this object relate heavily to the Node's expected lifecycle on
// etcd.
type Node struct {
	// clusterUnlockKey is half of the unlock key required to mount the node's
	// data partition. It's stored in etcd, and will only be provided to the
	// Node if it can prove its identity via an integrity mechanism (ie. via
	// TPM), or when the Node was just created (as the key is generated locally
	// by localstorage on first format/mount).  The other part of the unlock
	// key is the LocalUnlockKey that's present on the node's ESP partition.
	clusterUnlockKey []byte

	pubkey []byte

	state ppb.Node_FSMState

	// A Node can have multiple Roles. Each Role is represented by the presence
	// of NodeRole* structures in this structure, with a nil pointer
	// representing the lack of a role.
	consensusMember  *NodeRoleConsensusMember
	kubernetesWorker *NodeRoleKubernetesWorker

	// At runtime, this represents an etcd client to the consensus cluster. This
	// is used by applications (like Kubernetes).
	KV clientv3.KV
}

// NodeRoleConsensusMember defines that the Node is a consensus (etcd) cluster
// member.
type NodeRoleConsensusMember struct {
}

// NodeRoleKubernetesWorker defines that the Node should be running the
// Kubernetes control and data plane.
type NodeRoleKubernetesWorker struct {
}

// ID returns the name of this node, which is `metropolis-{pubkeyHash}`. This
// name should be the primary way to refer to Metropoils nodes within a
// cluster, and is guaranteed to be unique by relying on cryptographic
// randomness.
func (n *Node) ID() string {
	return fmt.Sprintf("metropolis-%s", n.IDBare())
}

// IDBare returns the `{pubkeyHash}` part of the node ID.
func (n Node) IDBare() string {
	return hex.EncodeToString(n.pubkey[:16])
}

func (n *Node) String() string {
	return n.ID()
}

// ConsensusMember returns a copy of the NodeRoleConsensusMember struct if the
// Node is a consensus member, otherwise nil.
func (n *Node) ConsensusMember() *NodeRoleConsensusMember {
	if n.consensusMember == nil {
		return nil
	}
	cm := *n.consensusMember
	return &cm
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

// etcdPath builds the etcd path in which this node's protobuf-serialized state
// is stored in etcd.
func (n *Node) etcdPath() string {
	return fmt.Sprintf("/nodes/%s", n.ID())
}

// proto serializes the Node object into protobuf, to be used for saving to
// etcd.
func (n *Node) proto() *ppb.Node {
	msg := &ppb.Node{
		ClusterUnlockKey: n.clusterUnlockKey,
		PublicKey:        n.pubkey,
		FsmState:         n.state,
		Roles:            &ppb.Node_Roles{},
	}
	if n.consensusMember != nil {
		msg.Roles.ConsensusMember = &ppb.Node_Roles_ConsensusMember{}
	}
	if n.kubernetesWorker != nil {
		msg.Roles.KubernetesWorker = &ppb.Node_Roles_KubernetesWorker{}
	}
	return msg
}

// Store saves the Node into etcd. This should be called only once per Node
// (ie. when the Node has been created).
func (n *Node) Store(ctx context.Context, kv clientv3.KV) error {
	// Currently the only flow to store a node to etcd is a write-once flow:
	// once a node is created, it cannot be deleted or updated. In the future,
	// flows to change cluster node roles might be introduced (ie. to promote
	// nodes to consensus members, etc).
	key := n.etcdPath()
	msg := n.proto()
	nodeRaw, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal node: %w", err)
	}

	res, err := kv.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(key), "=", 0),
	).Then(
		clientv3.OpPut(key, string(nodeRaw)),
	).Commit()
	if err != nil {
		return fmt.Errorf("failed to store node: %w", err)
	}

	if !res.Succeeded {
		return fmt.Errorf("attempted to re-register node (unsupported flow)")
	}
	return nil
}

// MakeConsensusMember turns the node into a consensus member. This only
// configures internal fields, and does not actually start any services.
func (n *Node) MakeConsensusMember() error {
	if n.consensusMember != nil {
		return fmt.Errorf("node already is consensus member")
	}
	n.consensusMember = &NodeRoleConsensusMember{}
	return nil
}

// MakeKubernetesWorker turns the node into a kubernetes worker. This only
// configures internal fields, and does not actually start any services.
func (n *Node) MakeKubernetesWorker() error {
	if n.kubernetesWorker != nil {
		return fmt.Errorf("node is already kubernetes worker")
	}
	n.kubernetesWorker = &NodeRoleKubernetesWorker{}
	return nil
}

// ConfigureLocalHostname uses the node's ID as a hostname, and sets the
// current hostname, and local files like hosts and machine-id accordingly.
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
	if err := ephemeral.MachineID.Write([]byte(n.IDBare()), 0644); err != nil {
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
