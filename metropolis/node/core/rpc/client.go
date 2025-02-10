// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package rpc

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"math/big"
	"time"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
)

// UnknownNotAfter is a copy of //metroplis/pkg/pki.UnknownNotAfter.
//
// We copy it so that we can decouple the rpc package from the pki package, the
// former being used by metroctl (and thus needing to be portable), the latter
// having a dependency on fileargs (which isn't portable). The correct solution
// here is to clarify portability policy of each workspace path, and apply it.
// But this will do for now.
//
// TODO(issues/252): clean up and merge this back.
var UnknownNotAfter = time.Unix(253402300799, 0)

type verifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

func verifyClusterCertificateAndNodeID(ca *x509.Certificate, nodeID string) verifyPeerCertificate {
	return func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
		if len(rawCerts) != 1 {
			return fmt.Errorf("server presented %d certificates, wanted exactly one", len(rawCerts))
		}
		serverCert, err := x509.ParseCertificate(rawCerts[0])
		if err != nil {
			return fmt.Errorf("server presented unparseable certificate: %w", err)
		}
		id, err := identity.VerifyNodeInCluster(serverCert, ca)
		if err != nil {
			return fmt.Errorf("node certificate verification failed: %w", err)
		}
		if nodeID != "" && id != nodeID {
			return fmt.Errorf("wanted to reach node %q, got %q", nodeID, id)
		}

		return nil
	}
}

func verifyFail(err error) verifyPeerCertificate {
	return func(_ [][]byte, _ [][]*x509.Certificate) error {
		return err
	}
}

// NewEphemeralCredentials returns gRPC TransportCredentials that can be used to
// dial a cluster without authenticating with a certificate, but instead
// authenticating by proving the possession of a private key, via an ephemeral
// self-signed certificate.
//
// Currently these credentials are used in two flows:
//
//  1. Registration of nodes into a cluster, after which a node receives a proper
//     node certificate
//
//  2. Escrow of initial owner credentials into a proper manager
//     certificate
//
// The given opts can be used to lock down the remote side of the connection, eg.
// expecting a given cluster CA certificate or disabling remote side verification
// by using WantInsecure().
func NewEphemeralCredentials(private ed25519.PrivateKey, opts ...CredentialsOpt) (credentials.TransportCredentials, error) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     UnknownNotAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	certificateBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, private.Public(), private)
	if err != nil {
		return nil, fmt.Errorf("when generating self-signed certificate: %w", err)
	}
	certificate := tls.Certificate{
		Certificate: [][]byte{certificateBytes},
		PrivateKey:  private,
	}
	return NewAuthenticatedCredentials(certificate, opts...), nil
}

// CredentialsOpt are created using WantXXX functions and used in
// NewCredentials.
type CredentialsOpt struct {
	wantCA       *x509.Certificate
	wantNodeID   string
	insecureOkay bool
}

func (a *CredentialsOpt) merge(o *CredentialsOpt) {
	if a.wantNodeID == "" && o.wantNodeID != "" {
		a.wantNodeID = o.wantNodeID
	}
	if a.wantCA == nil && o.wantCA != nil {
		a.wantCA = o.wantCA
	}
	if !a.insecureOkay && o.insecureOkay {
		a.insecureOkay = o.insecureOkay
	}
}

// WantRemoteCluster enables the verification of the remote cluster identity when
// using NewAuthanticatedCredentials. If the connection is not terminated at a
// cluster with the given CA certificate, an error will be returned.
//
// This is the bare minimum option required to implement secure connections to
// clusters.
func WantRemoteCluster(ca *x509.Certificate) CredentialsOpt {
	return CredentialsOpt{
		wantCA: ca,
	}
}

// WantRemoteNode enables the verification of the remote node identity when using
// NewCredentials. If the connection is not terminated at the node
// ID 'id', an error will be returned. For this function to work,
// WantRemoteCluster must also be set.
func WantRemoteNode(id string) CredentialsOpt {
	return CredentialsOpt{
		wantNodeID: id,
	}
}

// WantInsecure disables the verification of the remote side of the connection
// via NewCredentials. This is unsafe.
func WantInsecure() CredentialsOpt {
	return CredentialsOpt{
		insecureOkay: true,
	}
}

// NewAuthenticatedTLSConfig returns a tls.Config that can be used to dial a
// cluster with a given TLS certificate (from node or manager credentials).
//
// The provided CredentialsOpt specify the verification of the remote side of the
// connection. When connecting to a cluster (any node), use WantRemoteCluster. If
// you also want to verify the connection to a particular node, specify
// WantRemoteNode alongside it. If no verification should be performed use
// WantInsecure.
//
// The given options are parsed on a first-wins basis.
func NewAuthenticatedTLSConfig(cert tls.Certificate, opts ...CredentialsOpt) *tls.Config {
	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	var merged CredentialsOpt
	for _, o := range opts {
		merged.merge(&o)
	}

	if merged.insecureOkay {
		if merged.wantNodeID != "" {
			config.VerifyPeerCertificate = verifyFail(fmt.Errorf("WantInsecure specified alongside WantRemoteNode"))
		} else if merged.wantCA != nil {
			config.VerifyPeerCertificate = verifyFail(fmt.Errorf("WantInsecure specified alongside WantRemoteCluster"))
		}
	} else {
		switch {
		case merged.wantNodeID == "" && merged.wantCA == nil:
			config.VerifyPeerCertificate = verifyFail(fmt.Errorf("WantRemoteNode/WantRemoteCluster/WantInsecure not specified"))
		case merged.wantNodeID != "" && merged.wantCA == nil:
			config.VerifyPeerCertificate = verifyFail(fmt.Errorf("WantRemoteNode also requires WantRemoteCluster"))
		case merged.wantCA == nil:
			config.VerifyPeerCertificate = verifyFail(fmt.Errorf("no AuthenticaedCreentialsOpts specified"))
		default:
			config.VerifyPeerCertificate = verifyClusterCertificateAndNodeID(merged.wantCA, merged.wantNodeID)
		}
	}

	return config
}

// NewAuthenticatedCredentials returns gRPC TransportCredentials that can be used
// to dial a cluster with a given TLS certificate (from node or manager
// credentials).
//
// The provided CredentialsOpt specify the verification of the remote side of the
// connection. When connecting to a cluster (any node), use WantRemoteCluster. If
// you also want to verify the connection to a particular node, specify
// WantRemoteNode alongside it. If no verification should be performed use
// WantInsecure.
//
// The given options are parsed on a first-wins basis.
func NewAuthenticatedCredentials(cert tls.Certificate, opts ...CredentialsOpt) credentials.TransportCredentials {
	return credentials.NewTLS(NewAuthenticatedTLSConfig(cert, opts...))
}

// RetrieveOwnerCertificate uses AAA.Escrow to retrieve a cluster manager
// certificate for the initial owner of the cluster, authenticated by the
// public/private key set in the clusters NodeParameters.ClusterBoostrap.
//
// The retrieved certificate can be used to dial further cluster RPCs.
func RetrieveOwnerCertificate(ctx context.Context, aaa apb.AAAClient, private ed25519.PrivateKey) (*tls.Certificate, error) {
	srv, err := aaa.Escrow(ctx)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			return nil, status.Errorf(st.Code(), "Escrow call failed: %s", st.Message())
		}
		return nil, err
	}
	if err := srv.Send(&apb.EscrowRequest{
		Parameters: &apb.EscrowRequest_Parameters{
			RequestedIdentityName: "owner",
			PublicKey:             private.Public().(ed25519.PublicKey),
		},
	}); err != nil {
		return nil, fmt.Errorf("when sending client parameters: %w", err)
	}
	resp, err := srv.Recv()
	if err != nil {
		return nil, fmt.Errorf("when receiving server message: %w", err)
	}
	if len(resp.EmittedCertificate) == 0 {
		return nil, fmt.Errorf("expected certificate, instead got needed proofs: %+v", resp.Needed)
	}

	return &tls.Certificate{
		Certificate: [][]byte{resp.EmittedCertificate},
		PrivateKey:  private,
	}, nil
}
