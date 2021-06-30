package cluster

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"

	"source.monogon.dev/metropolis/node/core/curator"
)

type alterCert func(t *x509.Certificate)

func createPKI(t *testing.T, fca, fnode alterCert) (caCertBytes, nodeCertBytes, nodePriv []byte) {
	t.Helper()

	caPub, caPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	var nodePub ed25519.PublicKey
	nodePub, nodePriv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}

	caTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "CA"},
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageOCSPSigning},
		NotBefore:             time.Now(),
		NotAfter:              time.Unix(253402300799, 0),
		BasicConstraintsValid: true,
	}
	fca(caTemplate)

	caCertBytes, err = x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, caPub, caPriv)
	if err != nil {
		t.Fatalf("CreateCertificate (CA): %v", err)
	}
	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		t.Fatalf("ParseCertificate (CA): %v", err)
	}

	nodeTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		NotBefore:    time.Now(),
		NotAfter:     time.Unix(253402300799, 0),
		DNSNames:     []string{curator.NodeID(nodePub)},
	}
	fnode(nodeTemplate)

	nodeCertBytes, err = x509.CreateCertificate(rand.Reader, nodeTemplate, caCert, nodePub, caPriv)
	if err != nil {
		t.Fatalf("CreateCertificate (node): %v", err)
	}

	return
}

// TestNodeCertificateX509 exercises X509 validity checks performed by
// NewNodeCertificate.
func TestNodeCertificateX509(t *testing.T) {
	for i, te := range []struct {
		fca     alterCert
		fnode   alterCert
		success bool
	}{
		// Case 0: everything should work.
		{
			func(ca *x509.Certificate) {},
			func(n *x509.Certificate) {},
			true,
		},
		// Case 1: CA must be IsCA
		{
			func(ca *x509.Certificate) { ca.IsCA = false },
			func(n *x509.Certificate) {},
			false,
		},
		// Case 2: node must have its ID as a DNS name.
		{
			func(ca *x509.Certificate) {},
			func(n *x509.Certificate) { n.DNSNames = []string{"node"} },
			false,
		},
	} {
		caCert, nodeCert, nodePriv := createPKI(t, te.fca, te.fnode)
		_, err := NewNodeCredentials(nodePriv, nodeCert, caCert)
		if te.success && err != nil {
			t.Fatalf("Case %d: NewNodeCredentials failed: %v", i, err)
		}
		if !te.success && err == nil {
			t.Fatalf("Case %d: NewNodeCredentials succeeded, wanted failure", i)
		}
	}
}
