package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

func dialAuthenticated() *grpc.ClientConn {
	if len(flags.clusterEndpoints) == 0 {
		log.Fatal("Please provide at least one cluster endpoint using the --endpoint parameter.")
	}

	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := getCredentials()
	if err == noCredentialsError {
		log.Fatalf("You have to take ownership of the cluster first: %v", err)
	}

	ctx := clicontext.WithInterrupt(context.Background())
	cc, err := core.DialCluster(ctx, opkey, ocert, flags.proxyAddr, flags.clusterEndpoints, rpcLogger)
	if err != nil {
		log.Fatalf("While dialing the cluster: %v", err)
	}
	return cc
}
