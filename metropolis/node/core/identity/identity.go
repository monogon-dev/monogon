package identity

import (
	"crypto/ed25519"
	"crypto/subtle"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"source.monogon.dev/metropolis/node/core/localstorage"
)

// Node is the public part of the credentials of a node. They are
// emitted for a node by the cluster CA contained within the curator.
type Node struct {
	node *x509.Certificate
	ca   *x509.Certificate
}

// NewNode wraps a pair CA and node DER-encoded certificates into
// Node, ensuring the given certificate data is valid and compatible
// with Metropolis assumptions.
func NewNode(cert, ca []byte) (*Node, error) {
	certParsed, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, fmt.Errorf("could not parse node certificate: %w", err)
	}
	caCertParsed, err := x509.ParseCertificate(ca)
	if err != nil {
		return nil, fmt.Errorf("could not parse ca certificate: %w", err)
	}

	if _, err := VerifyNodeInCluster(certParsed, caCertParsed); err != nil {
		return nil, fmt.Errorf("could not node certificate within cluster CA: %w", err)
	}

	return &Node{
		node: certParsed,
		ca:   caCertParsed,
	}, nil
}

// PublicKey returns the Ed25519 public key corresponding to this node's
// certificate/credentials.
func (n *Node) PublicKey() ed25519.PublicKey {
	// Safe: we have ensured that the given certificate has an Ed25519 public key on
	// NewNode.
	return n.node.PublicKey.(ed25519.PublicKey)
}

// ClusterCA returns the CA certificate of the cluster for which this
// Node is emitted.
func (n *Node) ClusterCA() *x509.Certificate {
	return n.ca
}

// ID returns the canonical ID/name of the node for which this
// certificate/credentials were emitted.
func (n *Node) ID() string {
	return NodeID(n.PublicKey())
}

func (n *Node) Certificate() *x509.Certificate {
	return n.node
}

// NodeCredentials are the public and private part of the credentials of a node.
//
// It represents all the data necessary for a node to authenticate over mTLS to
// other nodes and the rest of the cluster.
//
// It must never be made available to any node other than the node it has been
// emitted for.
type NodeCredentials struct {
	Node
	private ed25519.PrivateKey
}

// NewNodeCredentials wraps a pair of CA and node DER-encoded certificates plus
// a private key into NodeCredentials, ensuring that the given data is valid and
// compatible with Metropolis assumptions.
func NewNodeCredentials(priv, cert, ca []byte) (*NodeCredentials, error) {
	nc, err := NewNode(cert, ca)
	if err != nil {
		return nil, err
	}

	// Ensure that the private key is a valid length.
	if want, got := ed25519.PrivateKeySize, len(priv); want != got {
		return nil, fmt.Errorf("private key is not the correct length, wanted %d, got %d", want, got)
	}

	// Ensure that the given private key matches the given public key.
	if want, got := ed25519.PrivateKey(priv).Public().(ed25519.PublicKey), nc.PublicKey(); subtle.ConstantTimeCompare(want, got) != 1 {
		return nil, fmt.Errorf("public key does not match private key")
	}

	return &NodeCredentials{
		Node:    *nc,
		private: priv,
	}, nil
}

func (n *NodeCredentials) TLSCredentials() tls.Certificate {
	return tls.Certificate{
		Leaf:        n.node,
		Certificate: [][]byte{n.node.Raw},
		PrivateKey:  n.private,
	}
}

// Save stores the given node credentials in local storage.
func (n *NodeCredentials) Save(d *localstorage.PKIDirectory) error {
	if err := d.CACertificate.Write(n.ca.Raw, 0400); err != nil {
		return fmt.Errorf("when writing CA certificate: %w", err)
	}
	if err := d.Certificate.Write(n.node.Raw, 0400); err != nil {
		return fmt.Errorf("when writing node certificate: %w", err)
	}
	if err := d.Key.Write(n.private, 0400); err != nil {
		return fmt.Errorf("when writing node private key: %w", err)
	}
	return nil
}

// NodeIDBare returns the `{pubkeyHash}` part of the node ID.
func NodeIDBare(pub []byte) string {
	return hex.EncodeToString(pub[:16])
}

// NodeID returns the name of this node, which is `metropolis-{pubkeyHash}`.
// This name should be the primary way to refer to Metropoils nodes within a
// cluster, and is guaranteed to be unique by relying on cryptographic
// randomness.
func NodeID(pub []byte) string {
	return fmt.Sprintf("metropolis-%s", NodeIDBare(pub))
}
