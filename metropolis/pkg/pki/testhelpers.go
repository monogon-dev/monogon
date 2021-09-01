package pki

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"testing"
)

// EphemeralClusterCredentials returns a pair of node and manager
// tls.Certificates signed by a CA certificate.
//
// All of these are ephemeral, ie. not stored anywhere - including the CA
// certificate. This function is for use by tests which want to bring up a
// minimum set of PKI credentials for a fake Metropolis cluster.
func EphemeralClusterCredentials(t *testing.T) (node, manager tls.Certificate, ca *x509.Certificate) {
	ctx := context.Background()

	ns := Namespaced("unused")
	caCert := Certificate{
		Namespace: &ns,
		Issuer:    SelfSigned,
		Template:  CA("test cluster ca"),
		Mode:      CertificateEphemeral,
	}
	caBytes, err := caCert.Ensure(ctx, nil)
	if err != nil {
		t.Fatalf("Could not ensure CA certificate: %v", err)
	}
	ca, err = x509.ParseCertificate(caBytes)
	if err != nil {
		t.Fatalf("Could not parse new CA certificate: %v", err)
	}

	nodeCert := Certificate{
		Namespace: &ns,
		Issuer:    &caCert,
		Template:  Server([]string{"test-server"}, nil),
		Mode:      CertificateEphemeral,
	}
	nodeBytes, err := nodeCert.Ensure(ctx, nil)
	if err != nil {
		t.Fatalf("Could not ensure node certificate: %v", err)
	}
	node = tls.Certificate{
		Certificate: [][]byte{nodeBytes},
		PrivateKey:  nodeCert.PrivateKey,
	}

	managerCert := Certificate{
		Namespace: &ns,
		Issuer:    &caCert,
		Template:  Client("owner", nil),
		Mode:      CertificateEphemeral,
	}
	managerBytes, err := managerCert.Ensure(ctx, nil)
	if err != nil {
		t.Fatalf("Could not ensure manager certificate: %v", err)
	}
	manager = tls.Certificate{
		Certificate: [][]byte{managerBytes},
		PrivateKey:  managerCert.PrivateKey,
	}
	return
}
