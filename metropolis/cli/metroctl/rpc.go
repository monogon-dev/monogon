package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"

	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
)

func dialAuthenticated(ctx context.Context) *grpc.ClientConn {
	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.NoCredentialsError) {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
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
	if errors.Is(err, core.NoCredentialsError) {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
	}
	cc, err := core.DialNode(ctx, opkey, ocert, cacert, flags.proxyAddr, id, address)
	if err != nil {
		log.Fatalf("While dialing node: %v", err)
	}
	return cc
}
