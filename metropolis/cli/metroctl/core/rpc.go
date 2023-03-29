package core

import (
	"context"
	"crypto/ed25519"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net"

	"golang.org/x/net/proxy"
	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	"source.monogon.dev/metropolis/proto/api"
)

type ResolverLogger func(format string, args ...interface{})

// DialCluster dials the cluster control address. The owner certificate, and
// proxy address parameters are optional and can be left nil, and empty,
// respectively. At least one cluster endpoint must be provided. A missing
// owner certificate will result in a connection that is authenticated with
// ephemeral credentials, restricting the available API surface. proxyAddr
// must point at a SOCKS5 endpoint.
func DialCluster(ctx context.Context, opkey ed25519.PrivateKey, ocert *x509.Certificate, proxyAddr string, clusterEndpoints []string, rlf ResolverLogger) (*grpc.ClientConn, error) {
	var dialOpts []grpc.DialOption

	if opkey == nil {
		return nil, fmt.Errorf("an owner's private key must be provided")
	}
	if len(clusterEndpoints) == 0 {
		return nil, fmt.Errorf("at least one cluster endpoint must be provided")
	}

	if proxyAddr != "" {
		socksDialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("failed to build a SOCKS dialer: %v", err)
		}
		grpcd := func(_ context.Context, addr string) (net.Conn, error) {
			return socksDialer.Dial("tcp", addr)
		}
		dialOpts = append(dialOpts, grpc.WithContextDialer(grpcd))
	}

	if ocert == nil {
		creds, err := rpc.NewEphemeralCredentials(opkey, nil)
		if err != nil {
			return nil, fmt.Errorf("while building ephemeral credentials: %v", err)
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	} else {
		tlsc := tls.Certificate{
			Certificate: [][]byte{ocert.Raw},
			PrivateKey:  opkey,
		}
		creds := rpc.NewAuthenticatedCredentials(tlsc, rpc.WantInsecure())
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	}

	var resolverOpts []resolver.ResolverOption
	if rlf != nil {
		resolverOpts = append(resolverOpts, resolver.WithLogger(rlf))
	}
	r := resolver.New(ctx, resolverOpts...)

	for _, eps := range clusterEndpoints {
		ep := resolver.NodeByHostPort(eps, uint16(node.CuratorServicePort))
		r.AddEndpoint(ep)
	}
	dialOpts = append(dialOpts, grpc.WithResolvers(r))

	c, err := grpc.Dial(resolver.MetropolisControlAddress, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("could not dial: %v", err)
	}
	return c, nil
}

func DialNode(ctx context.Context, opkey ed25519.PrivateKey, ocert, ca *x509.Certificate, proxyAddr, nodeId, nodeAddr string) (*grpc.ClientConn, error) {
	var dialOpts []grpc.DialOption

	if opkey == nil {
		return nil, fmt.Errorf("an owner's private key must be provided")
	}
	if proxyAddr != "" {
		socksDialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("failed to build a SOCKS dialer: %v", err)
		}
		grpcd := func(_ context.Context, addr string) (net.Conn, error) {
			return socksDialer.Dial("tcp", addr)
		}
		dialOpts = append(dialOpts, grpc.WithContextDialer(grpcd))
	}
	tlsc := tls.Certificate{
		Certificate: [][]byte{ocert.Raw},
		PrivateKey:  opkey,
	}
	creds := rpc.NewAuthenticatedCredentials(tlsc, rpc.WantRemoteCluster(ca), rpc.WantRemoteNode(nodeId))
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))

	endpoint := net.JoinHostPort(nodeAddr, node.NodeManagement.PortString())
	return grpc.Dial(endpoint, dialOpts...)
}

// GetNodes retrieves node records, filtered by the supplied node filter
// expression fexp.
func GetNodes(ctx context.Context, mgmt api.ManagementClient, fexp string) ([]*api.Node, error) {
	resN, err := mgmt.GetNodes(ctx, &api.GetNodesRequest{
		Filter: fexp,
	})
	if err != nil {
		return nil, err
	}

	var nodes []*api.Node
	for {
		node, err := resN.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
