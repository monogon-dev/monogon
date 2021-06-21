package curator

import (
	"fmt"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/pkg/pki"
)

var (
	// pkiNamespace is the etcd/pki namespace in which the Metropolis cluster CA
	// data will live.
	pkiNamespace = pki.Namespaced("/cluster-pki/")
	// pkiCA is the main cluster CA, stored in etcd. It is used to emit cluster,
	// node and user certificates.
	pkiCA = pkiNamespace.New(pki.SelfSigned, "cluster-ca", pki.CA("Metropolis Cluster CA"))
)

// NodeCredentials are the public and private part of the credentials of a node.
//
// It represents all the data necessary for a node to authenticate over mTLS to
// other nodes and the rest of the cluster.
//
// It must never be made available to any node other than the node it has been
// emitted for.
type NodeCredentials struct {
	NodeCertificate
	// PrivateKey is the ED25519 private key of the node, corresponding to
	// NodeCertificate.PublicKey.
	PrivateKey []byte
}

// NodeCertificate is the public part of the credential of a node.
type NodeCertificate struct {
	// PublicKey is the ED25519 public key of the node.
	PublicKey []byte
	// Certificate is the DER-encoded TLS certificate emitted for the node (ie.
	// PublicKey) by the cluster CA.
	Certificate []byte
	// CACertificate is the DER-encoded TLS certificate of the cluster CA at time of
	// emitting the Certificate.
	CACertificate []byte
}

// ID returns the Node ID of the node for which this NodeCertificate was
// emitted.
func (c *NodeCertificate) ID() string {
	return NodeID(c.PublicKey)
}

// Save stores the given node credentials in local storage.
func (c *NodeCredentials) Save(d *localstorage.PKIDirectory) error {
	if err := d.CACertificate.Write(c.CACertificate, 0400); err != nil {
		return fmt.Errorf("when writing CA certificate: %w", err)
	}
	if err := d.Certificate.Write(c.Certificate, 0400); err != nil {
		return fmt.Errorf("when writing node certificate: %w", err)
	}
	if err := d.Key.Write(c.PrivateKey, 0400); err != nil {
		return fmt.Errorf("when writing node private key: %w", err)
	}
	return nil
}
