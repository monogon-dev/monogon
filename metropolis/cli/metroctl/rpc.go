package main

import (
	"context"
	"crypto/x509"
	"log"

	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/cli/metroctl/core"
)

func dialAuthenticated(ctx context.Context) *grpc.ClientConn {
	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if err == core.NoCredentialsError {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
	}
	if len(flags.clusterEndpoints) == 0 {
		log.Fatal("Please provide at least one cluster endpoint using the --endpoint parameter.")
	}
	cc, err := core.DialCluster(ctx, opkey, ocert, flags.proxyAddr, flags.clusterEndpoints, rpcLogger)
	if err != nil {
		log.Fatalf("While dialing the cluster: %v", err)
	}
	return cc
}

func dialAuthenticatedNode(ctx context.Context, id, address string, cacert *x509.Certificate) *grpc.ClientConn {
	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if err == core.NoCredentialsError {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
	}
	cc, err := core.DialNode(ctx, opkey, ocert, cacert, flags.proxyAddr, id, address)
	if err != nil {
		log.Fatalf("While dialing node: %v", err)
	}
	return cc
}
