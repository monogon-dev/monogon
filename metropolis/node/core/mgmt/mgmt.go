// Package mgmt implements the node-local management service, a.k.a.
// metropolis.proto.api.NodeManagement.
package mgmt

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"

	apb "source.monogon.dev/metropolis/proto/api"
)

// Service implements metropolis.proto.api.NodeManagement.
type Service struct {
	// NodeCredentials used to set up gRPC server.
	NodeCredentials *identity.NodeCredentials
	// LogTree from which NodeManagement.Logs will be served.
	LogTree *logtree.LogTree

	// Automatically populated on Run.
	LogService
}

// Run the Servie as a supervisor runnable.
func (s *Service) Run(ctx context.Context) error {
	if s.NodeCredentials == nil {
		return fmt.Errorf("NodeCredentials missing")
	}
	if s.LogTree == nil {
		return fmt.Errorf("LogTree missing")
	}

	s.LogService.LogTree = s.LogTree

	sec := rpc.ServerSecurity{
		NodeCredentials: s.NodeCredentials,
	}
	logger := supervisor.MustSubLogger(ctx, "rpc")
	opts := sec.GRPCOptions(logger)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", node.NodeManagement))
	if err != nil {
		return fmt.Errorf("failed to listen on node management socket socket: %w", err)
	}
	defer lis.Close()

	srv := grpc.NewServer(opts...)
	apb.RegisterNodeManagementServer(srv, s)

	runnable := supervisor.GRPCServer(srv, lis, false)
	if err := supervisor.Run(ctx, "server", runnable); err != nil {
		return fmt.Errorf("could not run server: %w", err)
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return ctx.Err()
}
