package main

import (
	"context"
	"crypto/ed25519"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"

	"golang.org/x/net/proxy"
	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
)

// dialCluster dials the cluster control address. The owner certificate, and
// proxy address parameters are optional and can be left nil, and empty,
// respectively. At least one cluster endpoint must be provided. A missing
// owner certificate will result in a connection that is authenticated with
// ephemeral credentials, restricting the available API surface. proxyAddr
// must point at a SOCKS5 endpoint.
func dialCluster(ctx context.Context, opkey ed25519.PrivateKey, ocert *x509.Certificate, proxyAddr string, clusterEndpoints []string) (*grpc.ClientConn, error) {
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
		creds := rpc.NewAuthenticatedCredentials(tlsc, nil)
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	}

	r := resolver.New(ctx)
	for _, ep := range clusterEndpoints {
		r.AddEndpoint(resolver.NodeByHostPort(ep, uint16(node.CuratorServicePort)))
	}
	dialOpts = append(dialOpts, grpc.WithResolvers(r))

	c, err := grpc.Dial(resolver.MetropolisControlAddress, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("could not dial: %v", err)
	}
	return c, nil
}
