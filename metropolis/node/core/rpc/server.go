package rpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	apb "source.monogon.dev/metropolis/proto/api"
	epb "source.monogon.dev/metropolis/proto/ext"
)

// ClusterServices is the interface containing all gRPC services that a
// Metropolis Cluster implements on its public interface. With the current
// implementaiton of Metropolis, this is all implemented by the Curator.
type ClusterServices interface {
	cpb.CuratorServer
	apb.AAAServer
	apb.ManagementServer
}

// ServerSecurity are the security options of a RPC server that will run
// ClusterServices on a Metropolis node. It contains all the data for the
// server implementation to authenticate itself to the clients and authenticate
// and authorize clients connecting to it.
type ServerSecurity struct {
	// NodeCredentials is the TLS certificate/key of the node that the server
	// implementation is running on. It should be signed by
	// ClusterCACertificate.
	NodeCredentials tls.Certificate
	// ClusterCACertificate is the cluster's CA certificate. It will be used to
	// authenticate the client certificates of incoming gRPC connections.
	ClusterCACertificate *x509.Certificate
}

// SetupPublicGRPC returns a grpc.Server ready to listen and serve all public
// gRPC APIs that the cluster server implementation should run, with all calls
// being authenticated and authorized based on the data in ServerSecurity. The
// argument 'impls' is the object implementing the gRPC APIs.
//
// This effectively configures gRPC interceptors that verify
// metropolis.proto.ext.authorizaton options and authenticate/authorize
// incoming connections. It also runs the gRPC server with the correct TLS
// settings for authenticating itself to callers.
func (l *ServerSecurity) SetupPublicGRPC(impls ClusterServices) *grpc.Server {
	publicCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{l.NodeCredentials},
		ClientAuth:   tls.RequestClientCert,
	})

	s := grpc.NewServer(
		grpc.Creds(publicCreds),
		grpc.UnaryInterceptor(l.unaryInterceptor),
		grpc.StreamInterceptor(l.streamInterceptor),
	)
	cpb.RegisterCuratorServer(s, impls)
	apb.RegisterAAAServer(s, impls)
	apb.RegisterManagementServer(s, impls)
	return s
}

// authorize performs an authorization check for the given gRPC context
// (containing peer information) and given RPC method name (as obtained from
// FullMethodName in {Unary,Stream}ServerInfo). The actual authorization
// requirements per method are retrieved from the Authorization protobuf
// option applied to the RPC method.
//
// If the peer (as retrieved from the context) is authorized to run this method,
// no error is returned. Otherwise, a gRPC status is returned outlining the
// reason the authorization being rejected.
func (l *ServerSecurity) authorize(ctx context.Context, methodName string) error {
	if !strings.HasPrefix(methodName, "/") {
		return status.Errorf(codes.InvalidArgument, "invalid method name %q", methodName)
	}
	methodName = strings.ReplaceAll(methodName[1:], "/", ".")
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(methodName))
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "could not retrieve descriptor for method: %v", err)
	}
	method, ok := desc.(protoreflect.MethodDescriptor)
	if !ok {
		return status.Error(codes.InvalidArgument, "querying method name did not yield a MethodDescriptor")
	}

	// Get authorization extension, defaults to no options set.
	authz, ok := proto.GetExtension(method.Options(), epb.E_Authorization).(*epb.Authorization)
	if !ok || authz == nil {
		authz = &epb.Authorization{}
	}

	// If unauthenticated connections are allowed, let them through immediately.
	if authz.AllowUnauthenticated && len(authz.Need) == 0 {
		return nil
	}

	// Otherwise, we check that the other side of the connection is authenticated
	// using a valid cluster CA client certificate.
	p, ok := peer.FromContext(ctx)
	if !ok {
		return status.Error(codes.Unavailable, "could not retrive peer info")
	}
	tlsInfo, ok := p.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return status.Error(codes.Unauthenticated, "connection not secure")
	}
	count := len(tlsInfo.State.PeerCertificates)
	if count == 0 {
		return status.Errorf(codes.Unauthenticated, "no client certificate presented")
	}
	if count > 1 {
		return status.Errorf(codes.Unauthenticated, "exactly one client certificate must be sent (got %d)", count)
	}
	pCert := tlsInfo.State.PeerCertificates[0]

	// Ensure that the certificate is signed by the cluster CA.
	if err := pCert.CheckSignatureFrom(l.ClusterCACertificate); err != nil {
		return status.Errorf(codes.Unauthenticated, "invalid client certificate: %v", err)
	}
	// Ensure that the certificate is a client certificate.
	// TODO(q3k): synchronize this with //metropolis/pkg/pki Client()/Server()/...
	isClient := false
	for _, ku := range pCert.ExtKeyUsage {
		if ku == x509.ExtKeyUsageClientAuth {
			isClient = true
			break
		}
	}
	if !isClient {
		return status.Error(codes.PermissionDenied, "presented certificate is not a client certificate")
	}

	// MVP: all permissions are granted to all users.
	// TODO(q3k): check authz.Need once we have a user/identity system implemented.
	return nil
}

// streamInterceptor is a gRPC server stream interceptor that performs
// authentication and authorization of incoming RPCs based on the Authorization
// option set on each method.
func (l *ServerSecurity) streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := l.authorize(ss.Context(), info.FullMethod); err != nil {
		return err
	}
	return handler(srv, ss)
}

// unaryInterceptor is a gRPC server unary interceptor that performs
// authentication and authorization of incoming RPCs based on the Authorization
// option set on each method.
func (l *ServerSecurity) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if err := l.authorize(ctx, info.FullMethod); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
