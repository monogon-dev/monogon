package identity

import (
	"crypto/ed25519"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
)

// UserCertificate makes a Metropolis-compatible user certificate template.
func UserCertificate(identity string) x509.Certificate {
	return x509.Certificate{
		Subject: pkix.Name{
			CommonName: identity,
		},
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
		},
	}
}

// NodeCertificate makes a Metropolis-compatible node certificate template.
func NodeCertificate(nodeID string) x509.Certificate {
	return x509.Certificate{
		Subject: pkix.Name{
			CommonName: nodeID,
		},
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{
			// Note: node certificates are also effectively being used to perform client
			// authentication to other node certificates, but they don't have the ClientAuth
			// bit set. Instead, Metropolis uses the ClientAuth and ServerAuth bits
			// exclusively to distinguish Metropolis nodes from Metropolis users.
			x509.ExtKeyUsageServerAuth,
		},
		// We populate the Node's ID (metropolis-xxxx) as the DNS name for this
		// certificate for ease of use within Metropolis, where the local DNS setup
		// allows each node's IP address to be resolvable through the Node's ID.
		DNSNames: []string{
			nodeID,
		},
	}
}

// CACertificate makes a Metropolis-compatible CA certificate template.
//
// cn is a human-readable string that can be used to distinguish Metropolis
// clusters, if needed. It is not machine-parsed, instead only signature
// verification and CA pinning is performed.
func CACertificate(cn string) x509.Certificate {
	return x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: cn,
		},
		IsCA:        true,
		KeyUsage:    x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageOCSPSigning},
	}
}

// VerifyCAInsecure ensures that the given certificate is a valid certificate
// that is allowed to act as a CA and which is emitted for an Ed25519 keypair.
//
// It does _not_ ensure that the certificate is the local node's CA, and should
// not be used for security checks, just for data validation checks.
func VerifyCAInsecure(ca *x509.Certificate) error {
	if ca == nil {
		return fmt.Errorf("ca must be set")
	}
	// Ensure ca certificate uses ED25519 keypair.
	if _, ok := ca.PublicKey.(ed25519.PublicKey); !ok {
		return fmt.Errorf("not issued for ed25519 keypair")
	}
	// Ensure CA certificate has the X.509 basic constraints extension. Everything
	// else is legacy, we might as well weed that out early.
	if !ca.BasicConstraintsValid {
		return fmt.Errorf("does not have basic constraints")
	}
	// Ensure CA certificate can act as CA per BasicConstraints.
	if !ca.IsCA {
		return fmt.Errorf("not permitted to act as CA")
	}
	if ca.KeyUsage != 0 && ca.KeyUsage&x509.KeyUsageCertSign == 0 {
		return fmt.Errorf("not permitted to sign certificates")
	}
	return nil
}

// VerifyInCluster ensures that the given certificate has been signed by a CA
// certificate and are both certificates emitted for ed25519 keypairs.
//
// The subject certificate's public key is returned if verification is
// successful, and error is returned otherwise.
func VerifyInCluster(cert, ca *x509.Certificate) (ed25519.PublicKey, error) {
	if err := VerifyCAInsecure(ca); err != nil {
		return nil, fmt.Errorf("ca certificate invalid: %w", err)
	}

	// Ensure subject cert is signed by ca.
	if err := cert.CheckSignatureFrom(ca); err != nil {
		return nil, fmt.Errorf("signature verification failed: %w", err)
	}

	// Ensure subject certificate is _not_ CA. CAs (cluster or possibly
	// intermediaries) are not supposed to either directly serve traffic or perform
	// client actions on the cluster.
	if cert.IsCA {
		return nil, fmt.Errorf("subject certificate is a CA")
	}

	// Extract subject ED25519 public key.
	pubkey, ok := cert.PublicKey.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("certificate not issued for ed25519 keypair")
	}

	return pubkey, nil
}

// VerifyNodeInCluster ensures that a given certificate is a Metropolis node
// certificate emitted by a given Metropolis CA.
//
// The node's ID is returned if verification is successful, and error is
// returned otherwise.
func VerifyNodeInCluster(node, ca *x509.Certificate) (string, error) {
	pk, err := VerifyInCluster(node, ca)
	if err != nil {
		return "", err
	}

	// Ensure certificate has ServerAuth bit, thereby marking it as a node certificate.
	found := false
	for _, ku := range node.ExtKeyUsage {
		if ku == x509.ExtKeyUsageServerAuth {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("not a node certificate (missing ServerAuth key usage)")
	}

	id := NodeID(pk)

	// Ensure node ID is present in Subject.CommonName and at least one DNS name.
	if node.Subject.CommonName != id {
		return "", fmt.Errorf("node ID not found in CommonName")
	}

	found = false
	for _, n := range node.DNSNames {
		if n == id {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("node ID not found in DNSNames")
	}

	return id, nil
}

// VerifyUserInCluster ensures that a given certificate is a Metropolis user
// certificate emitted by a given Metropolis CA.
//
// The user certificate's identity is returned if verification is successful,
// and error is returned otherwise.
func VerifyUserInCluster(user, ca *x509.Certificate) (string, error) {
	_, err := VerifyInCluster(user, ca)
	if err != nil {
		return "", err
	}

	// Ensure certificate has ClientAuth bit, thereby marking it as a user certificate.
	found := false
	for _, ku := range user.ExtKeyUsage {
		if ku == x509.ExtKeyUsageClientAuth {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("not a user certificate (missing ClientAuth key usage)")
	}

	// Extract identity from CommonName, ensure set.
	identity := user.Subject.CommonName
	if identity == "" {
		return "", fmt.Errorf("CommonName not set")
	}
	return identity, nil
}
