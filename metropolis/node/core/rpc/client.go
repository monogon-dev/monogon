package rpc

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"math/big"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/pkg/pki"
	apb "source.monogon.dev/metropolis/proto/api"
)

type verifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

func verifyClusterCertificate(ca *x509.Certificate) verifyPeerCertificate {
	return func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
		if len(rawCerts) != 1 {
			return fmt.Errorf("server presented %d certificates, wanted exactly one", len(rawCerts))
		}
		serverCert, err := x509.ParseCertificate(rawCerts[0])
		if err != nil {
			return fmt.Errorf("server presented unparseable certificate: %w", err)
		}
		if _, err := identity.VerifyNodeInCluster(serverCert, ca); err != nil {
			return fmt.Errorf("node certificate verification failed: %w", err)
		}

		return nil
	}
}

// NewEphemeralClient dials a cluster's services using just a self-signed
// certificate and can be used to then escrow real cluster credentials for the
// owner.
//
// These self-signed certificates are used by clients connecting to the cluster
// which want to prove ownership of an ED25519 keypair but don't have any
// 'real' client certificate (yet). Current users include users of AAA.Escrow
// and new nodes Registering into the Cluster.
//
// If 'ca' is given, the remote side will be cryptographically verified to be a
// node that's part of the cluster represented by the ca. Otherwise, no
// verification is performed and this function is unsafe.
func NewEphemeralClient(remote string, private ed25519.PrivateKey, ca *x509.Certificate, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     pki.UnknownNotAfter,

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
	return NewAuthenticatedClient(remote, certificate, ca, opts...)
}

func NewEphemeralClientTest(listener *bufconn.Listener, private ed25519.PrivateKey, ca *x509.Certificate) (*grpc.ClientConn, error) {
	return NewEphemeralClient("local", private, ca, grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return listener.Dial()
	}))
}

// RetrieveOwnerCertificates uses AAA.Escrow to retrieve a cluster manager
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
	if err := srv.Send(&apb.EscrowFromClient{
		Parameters: &apb.EscrowFromClient_Parameters{
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

// NewAuthenticatedClient dials a cluster's services using the given TLS
// credentials (either user or node credentials).
//
// If 'ca' is given, the remote side will be cryptographically verified to be a
// node that's part of the cluster represented by the ca. Otherwise, no
// verification is performed and this function is unsafe.
func NewAuthenticatedClient(remote string, cert tls.Certificate, ca *x509.Certificate, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	if ca != nil {
		config.VerifyPeerCertificate = verifyClusterCertificate(ca)
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	return grpc.Dial(remote, opts...)
}

func NewAuthenticatedClientTest(listener *bufconn.Listener, cert tls.Certificate, ca *x509.Certificate) (*grpc.ClientConn, error) {
	return NewAuthenticatedClient("local", cert, ca, grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return listener.Dial()
	}))
}
