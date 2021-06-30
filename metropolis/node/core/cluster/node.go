package cluster

import (
	"crypto/ed25519"
	"crypto/subtle"
	"crypto/x509"
	"fmt"

	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/localstorage"
)

// NodeCertificate is the public part of the credentials of a node. They are
// emitted for a node by the cluster CA contained within the curator.
type NodeCertificate struct {
	node *x509.Certificate
	ca   *x509.Certificate
}

// NodeCredentials are the public and private part of the credentials of a node.
//
// It represents all the data necessary for a node to authenticate over mTLS to
// other nodes and the rest of the cluster.
//
// It must never be made available to any node other than the node it has been
// emitted for.
type NodeCredentials struct {
	NodeCertificate
	private ed25519.PrivateKey
}

// NewNodeCertificate wraps a pair CA and node DER-encoded certificates into
// NodeCertificate, ensuring the given certificate data is valid and compatible
// Metropolis assumptions.
//
// It does _not_ verify that the given CA is a known/trusted Metropolis CA for a
// running cluster.
func NewNodeCertificate(cert, ca []byte) (*NodeCertificate, error) {
	certParsed, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, fmt.Errorf("could not parse node certificate: %w", err)
	}
	caCertParsed, err := x509.ParseCertificate(ca)
	if err != nil {
		return nil, fmt.Errorf("could not parse ca certificate: %w", err)
	}

	// Ensure both CA and node certs use ED25519.
	if certParsed.PublicKeyAlgorithm != x509.Ed25519 {
		return nil, fmt.Errorf("node certificate must use ED25519, is %s", certParsed.PublicKeyAlgorithm.String())
	}
	if pub, ok := certParsed.PublicKey.(ed25519.PublicKey); !ok || len(pub) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("node certificate ED25519 key invalid")
	}
	if caCertParsed.PublicKeyAlgorithm != x509.Ed25519 {
		return nil, fmt.Errorf("CA certificate must use ED25519, is %s", caCertParsed.PublicKeyAlgorithm.String())
	}
	if pub, ok := caCertParsed.PublicKey.(ed25519.PublicKey); !ok || len(pub) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("CA certificate ED25519 key invalid")
	}

	// Ensure that the certificate is signed by the CA certificate.
	if err := certParsed.CheckSignatureFrom(caCertParsed); err != nil {
		return nil, fmt.Errorf("certificate not signed by given CA: %w", err)
	}

	// Ensure that the certificate has the node's calculated ID in its DNS names.
	found := false
	nid := curator.NodeID(certParsed.PublicKey.(ed25519.PublicKey))
	for _, n := range certParsed.DNSNames {
		if n == nid {
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("calculated node ID %q not found in node certificate's DNS names (%v)", nid, certParsed.DNSNames)
	}

	return &NodeCertificate{
		node: certParsed,
		ca:   caCertParsed,
	}, nil
}

// NewNodeCredentials wraps a pair of CA and node DER-encoded certificates plus
// a private key into NodeCredentials, ensuring that the given data is valid and
// compatible with Metropolis assumptions.
//
// It does _not_ verify that the given CA is a known/trusted Metropolis CA for a
// running cluster.
func NewNodeCredentials(priv, cert, ca []byte) (*NodeCredentials, error) {
	nc, err := NewNodeCertificate(cert, ca)
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
		NodeCertificate: *nc,
		private:         ed25519.PrivateKey(priv),
	}, nil
}

// Save stores the given node credentials in local storage.
func (c *NodeCredentials) Save(d *localstorage.PKIDirectory) error {
	if err := d.CACertificate.Write(c.ca.Raw, 0400); err != nil {
		return fmt.Errorf("when writing CA certificate: %w", err)
	}
	if err := d.Certificate.Write(c.node.Raw, 0400); err != nil {
		return fmt.Errorf("when writing node certificate: %w", err)
	}
	if err := d.Key.Write(c.private, 0400); err != nil {
		return fmt.Errorf("when writing node private key: %w", err)
	}
	return nil
}

// PublicKey returns the ED25519 public key corresponding to this node's
// certificate/credentials.
func (nc *NodeCertificate) PublicKey() ed25519.PublicKey {
	// Safe: we have ensured that the given certificate has an ed25519 public key on
	// NewNodeCertificate.
	return nc.node.PublicKey.(ed25519.PublicKey)
}

// ID returns the canonical ID/name of the node for which this
// certificate/credentials were emitted.
func (nc *NodeCertificate) ID() string {
	return curator.NodeID(nc.PublicKey())
}
