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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"source.monogon.dev/metropolis/pkg/pki"
	apb "source.monogon.dev/metropolis/proto/api"
)

// NewEphemeralClient dials a cluster's services using just a self-signed
// certificate and can be used to then escrow real cluster credentials for the
// owner.
//
// These self-signed certificates are used by clients connecting to the cluster
// which want to prove ownership of an ED25519 keypair but don't have any
// 'real' client certificate (yet). Current users include users of AAA.Escrow
// and new nodes Registering into the Cluster.
//
// If ca is given, the other side of the connection is verified to be served by
// a node presenting a certificate signed by that CA. Otherwise, no
// verification of the other side is performed (however, any attacker
// impersonating the cluster cannot use the escrowed credentials as the private
// key is never passed to the server).
func NewEphemeralClient(remote string, private ed25519.PrivateKey, ca *x509.Certificate) (*grpc.ClientConn, error) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     pki.UnknownNotAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
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
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{
			certificate,
		},
		InsecureSkipVerify: true,
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			if len(rawCerts) < 1 {
				return fmt.Errorf("server presented no certificate")
			}
			certs := make([]*x509.Certificate, len(rawCerts))
			for i, rawCert := range rawCerts {
				cert, err := x509.ParseCertificate(rawCert)
				if err != nil {
					return fmt.Errorf("could not parse server certificate %d: %v", i, err)
				}
				certs[i] = cert
			}

			if ca != nil {
				// CA given, perform full chain verification.
				roots := x509.NewCertPool()
				roots.AddCert(ca)
				opts := x509.VerifyOptions{
					Roots:         roots,
					Intermediates: x509.NewCertPool(),
				}
				for _, cert := range certs[1:] {
					opts.Intermediates.AddCert(cert)
				}
				_, err := certs[0].Verify(opts)
				if err != nil {
					return err
				}
			}

			// Regardless of CA given, ensure that the leaf certificate has the
			// right ExtKeyUsage.
			for _, ku := range certs[0].ExtKeyUsage {
				if ku == x509.ExtKeyUsageServerAuth {
					return nil
				}
			}
			return fmt.Errorf("server presented a certificate without server auth ext key usage")
		},
	})

	return grpc.Dial(remote, grpc.WithTransportCredentials(creds))
}

// RetrieveOwnerCertificates uses AAA.Escrow to retrieve a cluster manager
// certificate for the initial owner of the cluster, authenticated by the
// public/private key set in the clusters NodeParameters.ClusterBoostrap.
//
// The retrieved certificate can be used to dial further cluster RPCs.
func RetrieveOwnerCertificate(ctx context.Context, aaa apb.AAAClient, private ed25519.PrivateKey) (*tls.Certificate, error) {
	srv, err := aaa.Escrow(ctx)
	if err != nil {
		return nil, fmt.Errorf("when opening Escrow RPC: %w", err)
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
