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
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/sys/unix"

	"git.monogon.dev/source/nexantic.git/metropolis/node/core/localstorage"
	ipb "git.monogon.dev/source/nexantic.git/metropolis/proto/internal"
)

// Node is a Smalltown cluster member. A node is a virtual or physical machine running Smalltown. This object represents a
// node only as part of a Cluster - ie., this object will never be available outside of //metropolis/node/core/cluster
// if the Node is not part of a Cluster.
// Nodes are inherently tied to their long term storage, which is etcd. As such, methods on this object relate heavily
// to the Node's expected lifecycle on etcd.
type Node struct {
	// clusterUnlockKey is half of the unlock key required to mount the node's data partition. It's stored in etcd, and
	// will only be provided to the Node if it can prove its identity via an integrity mechanism (ie. via TPM), or when
	// the Node was just created (as the key is generated locally by localstorage on first format/mount).
	// The other part of the unlock key is the LocalUnlockKey that's present on the node's ESP partition.
	clusterUnlockKey []byte
	// certificate is the node's TLS certificate, used to authenticate Smalltown gRPC calls/services (but not
	// consensus/etcd). The certificate for a node is permanent (and never expires). It's self-signed by the node on
	// startup, and contains the node's IP address in its SAN. Callers/services should check directly against the
	// expected certificate, and not against a CA.
	certificate x509.Certificate
	// address is the management IP address of the node. The management IP address of a node is permanent.
	address net.IP

	// A Node can have multiple Roles. Each Role is represented by the presence of NodeRole* structures in this
	// structure, with a nil pointer representing the lack of a role.

	consensusMember  *NodeRoleConsensusMember
	kubernetesWorker *NodeRoleKubernetesWorker
}

// NewNode creates a new Node. This is only called when a New node is supposed to be created as part of a cluster,
// otherwise it should be loaded from Etcd.
func NewNode(cuk []byte, address net.IP, certificate x509.Certificate) *Node {
	if certificate.Raw == nil {
		panic("new node must contain raw certificate")
	}
	return &Node{
		clusterUnlockKey: cuk,
		certificate:      certificate,
		address:          address,
	}
}

// NodeRoleConsensusMember defines that the Node is a consensus (etcd) cluster member.
type NodeRoleConsensusMember struct {
	// etcdMember is the name of the node in Kubernetes. This is for now usually the same as the ID() of the Node.
	etcdMemberName string
}

// NodeRoleKubernetesWorker defines that the Node should be running the Kubernetes control and data plane.
type NodeRoleKubernetesWorker struct {
	// nodeName is the name of the node in Kubernetes. This is for now usually the same as the ID() of the Node.
	nodeName string
}

// ID returns the name of this node, which is `smalltown-{pubkeyHash}`. This name should be the primary way to refer to
// Smalltown nodes within a cluster, and is guaranteed to be unique by relying on cryptographic randomness.
func (n *Node) ID() string {
	return fmt.Sprintf("smalltown-%s", n.IDBare())
}

// IDBare returns the `{pubkeyHash}` part of the node ID.
func (n Node) IDBare() string {
	pubKey, ok := n.certificate.PublicKey.(ed25519.PublicKey)
	if !ok {
		panic("node has non-ed25519 public key")
	}
	return hex.EncodeToString(pubKey[:16])
}

func (n *Node) String() string {
	return n.ID()
}

// ConsensusMember returns a copy of the NodeRoleConsensusMember struct if the Node is a consensus member, otherwise
// nil.
func (n *Node) ConsensusMember() *NodeRoleConsensusMember {
	if n.consensusMember == nil {
		return nil
	}
	cm := *n.consensusMember
	return &cm
}

// KubernetesWorker returns a copy of the NodeRoleKubernetesWorker struct if the Node is a kubernetes worker, otherwise
// nil.
func (n *Node) KubernetesWorker() *NodeRoleKubernetesWorker {
	if n.kubernetesWorker == nil {
		return nil
	}
	kw := *n.kubernetesWorker
	return &kw
}

// etcdPath builds the etcd path in which this node's protobuf-serialized state is stored in etcd.
func (n *Node) etcdPath() string {
	return fmt.Sprintf("/nodes/%s", n.ID())
}

// proto serializes the Node object into protobuf, to be used for saving to etcd.
func (n *Node) proto() *ipb.Node {
	msg := &ipb.Node{
		Certificate:      n.certificate.Raw,
		ClusterUnlockKey: n.clusterUnlockKey,
		Address:          n.address.String(),
		Roles:            &ipb.Node_Roles{},
	}
	if n.consensusMember != nil {
		msg.Roles.ConsensusMember = &ipb.Node_Roles_ConsensusMember{
			EtcdMemberName: n.consensusMember.etcdMemberName,
		}
	}
	if n.kubernetesWorker != nil {
		msg.Roles.KubernetesWorker = &ipb.Node_Roles_KubernetesWorker{
			NodeName: n.kubernetesWorker.nodeName,
		}
	}
	return msg
}

// Store saves the Node into etcd. This should be called only once per Node (ie. when the Node has been created).
func (n *Node) Store(ctx context.Context, kv clientv3.KV) error {
	// Currently the only flow to store a node to etcd is a write-once flow: once a node is created, it cannot be
	// deleted or updated. In the future, flows to change cluster node roles might be introduced (ie. to promote nodes
	// to consensus members, etc).
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

// MakeConsensusMember turns the node into a consensus member with a given name. This only configures internal fields,
// and does not actually start any services.
func (n *Node) MakeConsensusMember(etcdMemberName string) error {
	if n.consensusMember != nil {
		return fmt.Errorf("node already is consensus member")
	}
	n.consensusMember = &NodeRoleConsensusMember{
		etcdMemberName: etcdMemberName,
	}
	return nil
}

// MakeKubernetesWorker turns the node into a kubernetes worker with a given name. This only configures internal fields,
// and does not actually start any services.
func (n *Node) MakeKubernetesWorker(name string) error {
	if n.kubernetesWorker != nil {
		return fmt.Errorf("node is already kubernetes worker")
	}
	n.kubernetesWorker = &NodeRoleKubernetesWorker{
		nodeName: name,
	}
	return nil
}

func (n *Node) Address() net.IP {
	return n.address
}

// ConfigureLocalHostname uses the node's ID as a hostname, and sets the current hostname, and local files like hosts
// and machine-id accordingly.
func (n *Node) ConfigureLocalHostname(etc *localstorage.EtcDirectory) error {
	if err := unix.Sethostname([]byte(n.ID())); err != nil {
		return fmt.Errorf("failed to set runtime hostname: %w", err)
	}
	if err := etc.Hosts.Write([]byte(fmt.Sprintf("%s %s", "127.0.0.1", n.ID())), 0644); err != nil {
		return fmt.Errorf("failed to write /etc/hosts: %w", err)
	}
	if err := etc.MachineID.Write([]byte(n.IDBare()), 0644); err != nil {
		return fmt.Errorf("failed to write /etc/machine-id: %w", err)
	}
	return nil
}
