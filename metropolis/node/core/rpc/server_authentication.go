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

// ServerSecurity are the security options of a RPC server that will run
// ClusterServices on a Metropolis node. It contains all the data for the
// server implementation to authenticate itself to the clients and authenticate
// and authorize clients connecting to it.
type ServerSecurity struct {
	// NodeCredentials which will be used to run the gRPC server, and whose CA
	// certificate will be used to authenticate incoming requests.
	NodeCredentials *identity.NodeCredentials

	// nodePermissions is used by tests to inject the permissions available to a
	// node. When not set, it defaults to the global nodePermissions map.
	nodePermissions Permissions
}

// SetupExternalGRPC returns a grpc.Server ready to listen and serve all gRPC
// services that the cluster server implementation should run, with all calls
// authenticated and authorized based on the data in ServerSecurity. The
// argument 'impls' is the object implementing the gRPC APIs.
//
// Under the hood, this configures gRPC interceptors that verify
// metropolis.proto.ext.authorization options and authenticate/authorize
// incoming connections. It also runs the gRPC server with the correct TLS
// settings for authenticating itself to callers.
func (s *ServerSecurity) SetupExternalGRPC(logger logtree.LeveledLogger, impls ClusterServices) *grpc.Server {
	externalCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{s.NodeCredentials.TLSCredentials()},
		ClientAuth:   tls.RequestClientCert,
	})

	srv := grpc.NewServer(
		grpc.Creds(externalCreds),
		grpc.UnaryInterceptor(s.unaryInterceptor(logger)),
		grpc.StreamInterceptor(s.streamInterceptor(logger)),
	)
	cpb.RegisterCuratorServer(srv, impls)
	apb.RegisterAAAServer(srv, impls)
	apb.RegisterManagementServer(srv, impls)
	return srv
}

// streamInterceptor returns a gRPC StreamInterceptor interface for use with
// grpc.NewServer. It's applied to gRPC servers started within Metropolis,
// notably to the Curator.
func (s *ServerSecurity) streamInterceptor(logger logtree.LeveledLogger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var span *logtreeSpan
		if logger != nil {
			span = newLogtreeSpan(logger)
			span.Printf("RPC invoked: streaming request: %s", info.FullMethod)
			ss = &spanServerStream{
				ServerStream: ss,
				span:         span,
			}
		}

		pi, err := s.authenticationCheck(ss.Context(), info.FullMethod)
		if err != nil {
			if s != nil {
				span.Printf("RPC send: authentication failed: %v", err)
			}
			return err
		}
		if span != nil {
			span.Printf("RPC peerInfo: %s", pi.String())
		}

		return handler(srv, pi.serverStream(ss))
	}
}

// unaryInterceptor returns a gRPC UnaryInterceptor interface for use with
// grpc.NewServer. It's applied to gRPC servers started within Metropolis,
// notably to the Curator.
func (s *ServerSecurity) unaryInterceptor(logger logtree.LeveledLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// Inject span if we have a logger.
		if logger != nil {
			ctx = contextWithSpan(ctx, newLogtreeSpan(logger))
		}

		Trace(ctx).Printf("RPC invoked: unary request: %s", info.FullMethod)

		// Perform authentication check and inject PeerInfo.
		pi, err := s.authenticationCheck(ctx, info.FullMethod)
		if err != nil {
			Trace(ctx).Printf("RPC send: authentication failed: %v", err)
			return nil, err
		}
		ctx = pi.apply(ctx)

		// Log authentication information.
		Trace(ctx).Printf("RPC peerInfo: %s", pi.String())

		// Call underlying handler.
		resp, err = handler(ctx, req)

		// Log result into span.
		if err != nil {
			Trace(ctx).Printf("RPC send: error: %v", err)
		} else {
			Trace(ctx).Printf("RPC send: ok, %s", protoMessagePretty(resp))
		}
		return
	}
}

// authenticationCheck is called by the unary and server interceptors to perform
// authentication and authorization checks for a given RPC.
func (s *ServerSecurity) authenticationCheck(ctx context.Context, methodName string) (*PeerInfo, error) {
	mi, err := getMethodInfo(methodName)
	if err != nil {
		return nil, err
	}

	if mi.unauthenticated {
		return s.getPeerInfoUnauthenticated(ctx)
	}

	pi, err := s.getPeerInfo(ctx)
	if err != nil {
		return nil, err
	}
	if err := pi.CheckPermissions(mi.need); err != nil {
		return nil, err
	}
	return pi, nil
}

// getPeerInfo is be called by authenticationCheck to authenticate incoming gRPC
// calls. It returns PeerInfo structure describing the authenticated other end
// of the connection, or a gRPC status if the other side could not be
// successfully authenticated.
//
// The returned PeerInfo can then be used to perform authorization checks based
// on the configured authentication of a given gRPC method, as described by the
// metropolis.proto.ext.authorization extension.
func (s *ServerSecurity) getPeerInfo(ctx context.Context) (*PeerInfo, error) {
	cert, err := getPeerCertificate(ctx)
	if err != nil {
		return nil, err
	}

	// Ensure that the certificate is signed by the cluster CA.
	if err := cert.CheckSignatureFrom(s.NodeCredentials.ClusterCA()); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "certificate not signed by cluster CA: %v", err)
	}

	nodepk, errNode := identity.VerifyNodeInCluster(cert, s.NodeCredentials.ClusterCA())
	if errNode == nil {
		// This is a Metropolis node.
		np := s.nodePermissions
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

	userid, errUser := identity.VerifyUserInCluster(cert, s.NodeCredentials.ClusterCA())
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

// getPeerInfoUnauthenticated is an equivalent to getPeerInfo, but called when a
// method is marked as 'unauthenticated'. The implementation should return a
// PeerInfo containing Unauthenticated, potentially populating it with
// UnauthenticatedPublicKey if such a public key could be retrieved.
func (s *ServerSecurity) getPeerInfoUnauthenticated(ctx context.Context) (*PeerInfo, error) {
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
