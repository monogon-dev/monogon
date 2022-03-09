package rpc

import (
	"context"
	"crypto/ed25519"
	"crypto/tls"
	"crypto/x509"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/pkg/logtree"
	apb "source.monogon.dev/metropolis/proto/api"
)

// authenticationStrategy is implemented by ExternalServerSecurity. Historically
// it has also been implemented by LocalServerSecurity (listening on a local
// domain socket), but this implementation has since been removed.
//
// TODO(q3k): simplify this code and remove this interface now that there's only
// ExternalServerSecurity.
type authenticationStrategy interface {
	// getPeerInfo will be called by the stream and unary gRPC server interceptors
	// to authenticate incoming gRPC calls. It's given the gRPC context of the call
	// (therefore allowing access to information about the underlying gRPC
	// transport), and should return a PeerInfo structure describing the
	// authenticated other end of the connection, or a gRPC status if the other
	// side could not be successfully authenticated.
	//
	// The returned PeerInfo will then be used to perform authorization checks based
	// on the configured authentication of a given gRPC method, as described by the
	// metropolis.proto.ext.authorization extension. The same PeerInfo will then be
	// available to the gRPC handler for this method by retrieving it from the
	// context (via GetPeerInfo).
	getPeerInfo(ctx context.Context) (*PeerInfo, error)

	// getPeerInfoUnauthenticated is an equivalent to getPeerInfo, but called by the
	// interceptors when a method is marked as 'unauthenticated'. The implementation
	// should return a PeerInfo containing Unauthenticated, potentially populating
	// it with UnauthenticatedPublicKey if such a public key could be retrieved.
	getPeerInfoUnauthenticated(ctx context.Context) (*PeerInfo, error)
}

// authenticationCheck is called by the unary and server interceptors to perform
// authentication and authorization checks for a given RPC, calling the
// serverInterceptors' authenticate function if needed.
func authenticationCheck(ctx context.Context, a authenticationStrategy, methodName string) (*PeerInfo, error) {
	mi, err := getMethodInfo(methodName)
	if err != nil {
		return nil, err
	}

	if mi.unauthenticated {
		return a.getPeerInfoUnauthenticated(ctx)
	}

	pi, err := a.getPeerInfo(ctx)
	if err != nil {
		return nil, err
	}
	if err := pi.CheckPermissions(mi.need); err != nil {
		return nil, err
	}
	return pi, nil
}

// ServerSecurity are the security options of a RPC server that will run
// ClusterServices on a Metropolis node. It contains all the data for the
// server implementation to authenticate itself to the clients and authenticate
// and authorize clients connecting to it.
//
// It implements authenticationStrategy.
type ExternalServerSecurity struct {
	// NodeCredentials which will be used to run the gRPC server, and whose CA
	// certificate will be used to authenticate incoming requests.
	NodeCredentials *identity.NodeCredentials

	// nodePermissions is used by tests to inject the permissions available to a
	// node. When not set, it defaults to the global nodePermissions map.
	nodePermissions Permissions
}

// SetupExternalGRPC returns a grpc.Server ready to listen and serve all external
// gRPC APIs that the cluster server implementation should run, with all calls
// being authenticated and authorized based on the data in ServerSecurity. The
// argument 'impls' is the object implementing the gRPC APIs.
//
// This effectively configures gRPC interceptors that verify
// metropolis.proto.ext.authorization options and authenticate/authorize
// incoming connections. It also runs the gRPC server with the correct TLS
// settings for authenticating itself to callers.
func (l *ExternalServerSecurity) SetupExternalGRPC(logger logtree.LeveledLogger, impls ClusterExternalServices) *grpc.Server {
	externalCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{l.NodeCredentials.TLSCredentials()},
		ClientAuth:   tls.RequestClientCert,
	})

	s := grpc.NewServer(
		grpc.Creds(externalCreds),
		grpc.UnaryInterceptor(unaryInterceptor(logger, l)),
		grpc.StreamInterceptor(streamInterceptor(logger, l)),
	)
	cpb.RegisterCuratorServer(s, impls)
	apb.RegisterAAAServer(s, impls)
	apb.RegisterManagementServer(s, impls)
	return s
}

func (l *ExternalServerSecurity) getPeerInfo(ctx context.Context) (*PeerInfo, error) {
	cert, err := getPeerCertificate(ctx)
	if err != nil {
		return nil, err
	}

	// Ensure that the certificate is signed by the cluster CA.
	if err := cert.CheckSignatureFrom(l.NodeCredentials.ClusterCA()); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "certificate not signed by cluster CA: %v", err)
	}

	nodepk, errNode := identity.VerifyNodeInCluster(cert, l.NodeCredentials.ClusterCA())
	if errNode == nil {
		// This is a Metropolis node.
		np := l.nodePermissions
		if np == nil {
			np = nodePermissions
		}
		return &PeerInfo{
			Node: &PeerInfoNode{
				PublicKey:   nodepk,
				Permissions: np,
			},
		}, nil
	}

	userid, errUser := identity.VerifyUserInCluster(cert, l.NodeCredentials.ClusterCA())
	if errUser == nil {
		// This is a Metropolis user/manager.
		return &PeerInfo{
			User: &PeerInfoUser{
				Identity: userid,
			},
		}, nil
	}

	// Could not parse as either node or user certificate.
	return nil, status.Errorf(codes.Unauthenticated, "presented certificate is neither user certificate (%v) nor node certificate (%v)", errUser, errNode)
}

func (l *ExternalServerSecurity) getPeerInfoUnauthenticated(ctx context.Context) (*PeerInfo, error) {
	res := PeerInfo{
		Unauthenticated: &PeerInfoUnauthenticated{},
	}

	// If peer presented a valid self-signed certificate, attach that to the
	// Unauthenticated struct.
	cert, err := getPeerCertificate(ctx)
	if err == nil {
		if err := cert.CheckSignature(cert.SignatureAlgorithm, cert.RawTBSCertificate, cert.Signature); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "presented certificate must be self-signed (check error: %v)", err)
		}
		res.Unauthenticated.SelfSignedPublicKey = cert.PublicKey.(ed25519.PublicKey)
	}

	return &res, nil
}

// getPeerCertificate returns the x509 certificate associated with the given
// gRPC connection's context and ensures that it is a certificate for an Ed25519
// keypair. The certificate is _not_ checked against the cluster CA.
//
// A gRPC status is returned if the certificate is invalid / unauthenticated for
// any reason.
func getPeerCertificate(ctx context.Context) (*x509.Certificate, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unavailable, "could not retrive peer info")
	}
	tlsInfo, ok := p.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "connection not secure")
	}
	count := len(tlsInfo.State.PeerCertificates)
	if count == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "no client certificate presented")
	}
	if count > 1 {
		return nil, status.Errorf(codes.Unauthenticated, "exactly one client certificate must be sent (got %d)", count)
	}
	cert := tlsInfo.State.PeerCertificates[0]
	if _, ok := cert.PublicKey.(ed25519.PublicKey); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "certificate must be issued for an ED25519 keypair")
	}

	return cert, nil
}
