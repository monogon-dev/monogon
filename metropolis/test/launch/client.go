package launch

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	apb "source.monogon.dev/metropolis/proto/api"
)

// InitialClient implements a gRPC wrapper for dialing a Metropolis cluster
// while not (yet) authenticated, i.e. using only a self-signed public
// certificate to prove ownership of an ed25519 public key.
//
// This is used to dial a cluster's AAA.Escrow service after cluster bootstrap
// using the owner key configured in NodeParams.
type InitialClient struct {
	// conn is the underlying dialed gRPC connection to the cluster.
	conn *grpc.ClientConn
	// aaa is a stub to the AAA service running on conn.
	aaa apb.AAAClient
	// options are the options this client has been opened with.
	options *InitialClientOptions
}

type InitialClientOptions struct {
	// Remote is an address:port to connect to. This should be a cluster node's
	// curator port.
	Remote string
	// Private is the cluster owner private key, which should correspond to the
	// owner public key defined in NodeParametrs.ClusterBootstrap when a cluster
	// is bootstrapped.
	Private ed25519.PrivateKey
}

// NewInitialClient dials a cluster's curator service using just a self-signed
// certificate and can be used to then escrow real cluster credentials for the
// owner.
//
// MVP SECURITY: this does not verify the identity of the cluster/node. However,
// because any intercepting party cannot forward the presented client
// certificate to any real cluster, no danger of intercepting administrative
// access to the expected cluster is possible. Instead, the interceptor can just
// pretend to be the cluster which was expected.
func NewInitialClient(o *InitialClientOptions) (*InitialClient, error) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}
	certificateBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, o.Private.Public(), o.Private)
	if err != nil {
		return nil, fmt.Errorf("when generating self-signed certificate: %w", err)
	}
	keyBytes, err := x509.MarshalPKCS8PrivateKey(o.Private)
	if err != nil {
		return nil, fmt.Errorf("when marshaling private key: %w", err)
	}
	key := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})
	certificate := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificateBytes})
	clientCert, err := tls.X509KeyPair(certificate, key)
	if err != nil {
		return nil, fmt.Errorf("when building self-signed TLS client certificate: %w", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{
			clientCert,
		},
		InsecureSkipVerify:    true,
		VerifyPeerCertificate: o.verify,
	})

	conn, err := grpc.Dial(o.Remote, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("when dialing: %w", err)
	}

	return &InitialClient{
		conn:    conn,
		aaa:     apb.NewAAAClient(conn),
		options: o,
	}, nil
}

// Close must be called when the InitialClient is not used anymore. This closes
// the underlying gRPC connection(s).
func (i *InitialClient) Close() error {
	return i.conn.Close()
}

func (o *InitialClientOptions) verify(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	// SECURITY: Always permit all server certificates. See NewInitialClient godoc
	// for more information.
	return nil
}

// RetrieveOwnerCertificates uses AAA.Escrow to retrieve a cluster manager
// certificate for the initial owner of the cluster, authenticated by the
// public/private key set in the clusters NodeParameters.ClusterBoostrap.
//
// The retrieved certificate can be used to dial further cluster RPCs.
func (i *InitialClient) RetrieveOwnerCertificate(ctx context.Context) (*tls.Certificate, error) {
	srv, err := i.aaa.Escrow(ctx)
	if err != nil {
		return nil, fmt.Errorf("when opening Escrow RPC: %w", err)
	}
	if err := srv.Send(&apb.EscrowFromClient{
		Parameters: &apb.EscrowFromClient_Parameters{
			RequestedIdentityName: "owner",
			PublicKey:             i.options.Private.Public().(ed25519.PublicKey),
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

	certificateBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: resp.EmittedCertificate})
	key, err := x509.MarshalPKCS8PrivateKey(i.options.Private)
	if err != nil {
		return nil, fmt.Errorf("while marshalling private key: %w", err)
	}
	keyBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})
	ownerCert, err := tls.X509KeyPair(certificateBytes, keyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not build certificate from data received from cluster: %w", err)
	}

	return &ownerCert, nil
}
