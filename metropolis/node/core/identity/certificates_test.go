package identity

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"math/big"
	"testing"
	"time"
)

// alterCert is used by test code to slightly alter certificates before they get
// signed.
type alterCert func(t *x509.Certificate)

// basic is the bare minimum for ceritifcates to be properly issued over what
// {CA,User,Node}Certificate return. The equivalent logic is present in the pki
// codebase, we replicate it here because we don't use pki.
func basic(t *x509.Certificate) {
	t.SerialNumber = big.NewInt(1)
	t.NotBefore = time.Now()
	t.NotAfter = time.Unix(253402300799, 0)
	t.BasicConstraintsValid = true
}

func noop(_ *x509.Certificate) {}

// createPKI builds a minimum viable cluster PKI. We do not use
// EphemeralClusterCredentials because we want to test the behaviour of the
// certificate verification code when the certificate templates are slightly
// altered, including in ways that the pki could would normally prevent us
// from doing.
func createPKI(t *testing.T, fca, fnode, fuser alterCert) (caCertBytes, nodeCertBytes, userCertBytes []byte) {
	t.Helper()

	caPub, caPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	nodePub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	userPub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}

	caTemplate := CACertificate("test metropolis CA")
	basic(&caTemplate)
	fca(&caTemplate)

	caCertBytes, err = x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, caPub, caPriv)
	if err != nil {
		t.Fatalf("CreateCertificate (CA): %v", err)
	}
	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		t.Fatalf("ParseCertificate (CA): %v", err)
	}

	nodeTemplate := NodeCertificate(NodeID(nodePub))
	basic(&nodeTemplate)
	fnode(&nodeTemplate)
	nodeCertBytes, err = x509.CreateCertificate(rand.Reader, &nodeTemplate, caCert, nodePub, caPriv)
	if err != nil {
		t.Fatalf("CreateCertificate (node): %v", err)
	}

	userTemplate := UserCertificate("test")
	basic(&userTemplate)
	fuser(&userTemplate)
	userCertBytes, err = x509.CreateCertificate(rand.Reader, &userTemplate, caCert, userPub, caPriv)
	if err != nil {
		t.Fatalf("CreateCertificate (node): %v", err)
	}

	return
}

func TestCertificates(t *testing.T) {
	for i, te := range []struct {
		fca         alterCert
		fnode       alterCert
		fuser       alterCert
		successNode bool
		successUser bool
	}{
		// Case 0: everything should work.
		{
			noop,
			noop,
			noop,
			true, true,
		},
		// Case 1: CA must be IsCA
		{
			func(ca *x509.Certificate) { ca.IsCA = false },
			noop,
			noop,
			false, false,
		},
		// Case 2: node must not have IsCA set
		{
			noop,
			func(n *x509.Certificate) { n.IsCA = true },
			noop,
			false, true,
		},
		// Case 3: user must not have IsCA set
		{
			noop,
			noop,
			func(u *x509.Certificate) { u.IsCA = true },
			true, false,
		},
		// Case 4: node must have its ID as a DNS name.
		{
			noop,
			func(n *x509.Certificate) { n.DNSNames = []string{"node"} },
			noop,
			false, true,
		},
		// Case 5: node must have its ID as CommoNName.
		{
			noop,
			func(n *x509.Certificate) { n.Subject.CommonName = "node" },
			noop,
			false, true,
		},
		// Case 6: user must have CommonName set.
		{
			noop,
			noop,
			func(u *x509.Certificate) { u.Subject.CommonName = "" },
			true, false,
		},
	} {
		caCert, nodeCert, userCert := createPKI(t, te.fca, te.fnode, te.fuser)
		caCertParsed, err := x509.ParseCertificate(caCert)
		if err != nil {
			t.Fatalf("Case %d: ParseCertificate(ca): %v", i, err)
		}
		nodeCertParsed, err := x509.ParseCertificate(nodeCert)
		if err != nil {
			t.Fatalf("Case %d: ParseCertificate(node): %v", i, err)
		}
		userCertParsed, err := x509.ParseCertificate(userCert)
		if err != nil {
			t.Fatalf("Case %d: ParseCertificate(node): %v", i, err)
		}

		// Check node certificate as node certificate. Should succeed iff successNode.
		_, err = VerifyNodeInCluster(nodeCertParsed, caCertParsed)
		if te.successNode && err != nil {
			t.Errorf("Case %d: VerifyNodeInCluster failed: %v", i, err)
		}
		if !te.successNode && err == nil {
			t.Errorf("Case %d: VerifyNodeInCluster succeeded, wanted failure", i)
		}

		// Check user certificate as user certificate. Should succeed iff successUser.
		_, err = VerifyUserInCluster(userCertParsed, caCertParsed)
		if te.successUser && err != nil {
			t.Errorf("Case %d: VerifyUserInCluster failed: %v", i, err)
		}
		if !te.successUser && err == nil {
			t.Errorf("Case %d: VerifyUserInCluster succeeded, wanted failure", i)
		}

		// Check user certificate as node certificate. Should always fail.
		if _, err := VerifyNodeInCluster(userCertParsed, caCertParsed); err == nil {
			t.Errorf("Case %d: User certificate erroneously verified as node ceritficate", i)
		}
		// Check node certificate as user certificate. Should always fail.
		if _, err := VerifyUserInCluster(nodeCertParsed, caCertParsed); err == nil {
			t.Errorf("Case %d: Node certificate erroneously verified as user ceritficate", i)
		}
	}
}
