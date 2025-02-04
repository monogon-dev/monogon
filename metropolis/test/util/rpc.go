// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"testing"

	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/osbase/pki"
)

// NewEphemeralClusterCredentials creates a set of TLS certificates for use in a
// test Metropolis cluster. These are a CA certificate, a Manager certificate
// and an arbitrary amount of Node certificates (per the nodes argument).
//
// All of these are ephemeral, ie. not stored anywhere - including the CA
// certificate. This function is for use by tests which want to bring up a
// minimum set of PKI credentials for a fake Metropolis cluster.
func NewEphemeralClusterCredentials(t *testing.T, nodes int) *EphemeralClusterCredentials {
	ctx := context.Background()
	t.Helper()

	ns := pki.Namespaced("unused")
	caCert := pki.Certificate{
		Namespace: &ns,
		Issuer:    pki.SelfSigned,
		Template:  identity.CACertificate("test cluster ca"),
		Mode:      pki.CertificateEphemeral,
	}
	caBytes, err := caCert.Ensure(ctx, nil)
	if err != nil {
		t.Fatalf("Could not ensure CA certificate: %v", err)
	}
	ca, err := x509.ParseCertificate(caBytes)
	if err != nil {
		t.Fatalf("Could not parse new CA certificate: %v", err)
	}

	managerCert := pki.Certificate{
		Namespace: &ns,
		Issuer:    &caCert,
		Template:  identity.UserCertificate("owner"),
		Mode:      pki.CertificateEphemeral,
	}
	managerBytes, err := managerCert.Ensure(ctx, nil)
	if err != nil {
		t.Fatalf("Could not ensure manager certificate: %v", err)
	}
	res := &EphemeralClusterCredentials{
		Nodes: make([]*identity.NodeCredentials, nodes),
		Manager: tls.Certificate{
			Certificate: [][]byte{managerBytes},
			PrivateKey:  managerCert.PrivateKey,
		},
		CA: ca,
	}

	for i := 0; i < nodes; i++ {
		npk, npr, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			t.Fatalf("Could not generate node keypair: %v", err)
		}
		nodeCert := pki.Certificate{
			Namespace: &ns,
			Issuer:    &caCert,
			Template:  identity.NodeCertificate(identity.NodeID(npk)),
			Mode:      pki.CertificateEphemeral,
			PublicKey: npk,
			Name:      "",
		}
		nodeBytes, err := nodeCert.Ensure(ctx, nil)
		if err != nil {
			t.Fatalf("Could not ensure node certificate: %v", err)
		}
		node, err := identity.NewNodeCredentials(npr, nodeBytes, caBytes)
		if err != nil {
			t.Fatalf("Could not build node credentials: %v", err)
		}
		res.Nodes[i] = node
	}

	return res
}

// EphemeralClusterCredentials are TLS/PKI credentials for use in a Metropolis
// test cluster.
type EphemeralClusterCredentials struct {
	// Nodes are the node credentials for the cluster. Each contains a private
	// key and x509 certificate authenticating the bearer as a Metropolis node.
	Nodes []*identity.NodeCredentials
	// Manager TLS certificate for the cluster. Contains a private key and x509
	// certificate authenticating the bearer as a Metropolis manager.
	Manager tls.Certificate
	// CA is the x509 certificate of the CA certificate for the cluster. Manager and
	// Node certificates are signed by this CA.
	CA *x509.Certificate
}
