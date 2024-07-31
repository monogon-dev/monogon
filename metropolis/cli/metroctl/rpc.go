package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/proxy"
	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
)

func dialAuthenticated(ctx context.Context) *grpc.ClientConn {
	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.ErrNoCredentials) {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to get owner credentials: %v", err)
	}
	if len(flags.clusterEndpoints) == 0 {
		log.Fatal("Please provide at least one cluster endpoint using the --endpoint parameter.")
	}

	ca, err := core.GetClusterCAWithTOFU(ctx, connectOptions())
	if err != nil {
		log.Fatalf("Failed to get cluster CA: %v", err)
	}

	tlsc := tls.Certificate{
		Certificate: [][]byte{ocert.Raw},
		PrivateKey:  opkey,
	}
	creds := rpc.NewAuthenticatedCredentials(tlsc, rpc.WantRemoteCluster(ca))
	opts, err := core.DialOpts(ctx, connectOptions())
	if err != nil {
		log.Fatalf("While configuring dial options: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	cc, err := grpc.Dial(resolver.MetropolisControlAddress, opts...)
	if err != nil {
		log.Fatalf("While dialing cluster: %v", err)
	}
	return cc
}

func dialAuthenticatedNode(ctx context.Context, id, address string, cacert *x509.Certificate) *grpc.ClientConn {
	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.ErrNoCredentials) {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
	}
	cc, err := core.DialNode(ctx, opkey, ocert, cacert, flags.proxyAddr, id, address)
	if err != nil {
		log.Fatalf("While dialing node: %v", err)
	}
	return cc
}

func newAuthenticatedNodeHTTPTransport(ctx context.Context, id string) *http.Transport {
	cacert, err := core.GetClusterCAWithTOFU(ctx, connectOptions())
	if err != nil {
		log.Fatalf("Could not get CA certificate: %v", err)
	}
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.ErrNoCredentials) {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
	}
	tlsc := tls.Certificate{
		Certificate: [][]byte{ocert.Raw},
		PrivateKey:  opkey,
	}
	tlsconf := rpc.NewAuthenticatedTLSConfig(tlsc, rpc.WantRemoteCluster(cacert), rpc.WantRemoteNode(id))
	transport := &http.Transport{
		TLSClientConfig: tlsconf,
	}
	if flags.proxyAddr != "" {
		dialer, err := proxy.SOCKS5("tcp", flags.proxyAddr, nil, proxy.Direct)
		if err != nil {
			log.Fatalf("Failed to create proxy dialer: %v", err)
		}
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			// TODO(q3k): handle context
			return dialer.Dial(network, addr)
		}
	}
	return transport
}
