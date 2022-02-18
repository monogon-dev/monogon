package rpc

import (
	"context"

	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/pkg/logtree"
)

// stream implements the gRPC StreamInterceptor interface for use with
// grpc.NewServer, based on an authenticationStrategy. It's applied to gRPC
// servers started within Metropolis, notably to the Curator.
func streamInterceptor(logger logtree.LeveledLogger, a authenticationStrategy) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var s *logtreeSpan
		if logger != nil {
			s = newLogtreeSpan(logger)
			s.Printf("RPC invoked: streaming request: %s", info.FullMethod)
			ss = &spanServerStream{
				ServerStream: ss,
				span:         s,
			}
		}

		pi, err := authenticationCheck(ss.Context(), a, info.FullMethod)
		if err != nil {
			if s != nil {
				s.Printf("RPC send: authentication failed: %v", err)
			}
			return err
		}
		if s != nil {
			s.Printf("RPC peerInfo: %s", pi.String())
		}

		return handler(srv, pi.serverStream(ss))
	}
}

// unaryInterceptor implements the gRPC UnaryInterceptor interface for use with
// grpc.NewServer, based on an authenticationStrategy. It's applied to gRPC
// servers started within Metropolis, notably to the Curator.
func unaryInterceptor(logger logtree.LeveledLogger, a authenticationStrategy) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// Inject span if we have a logger.
		if logger != nil {
			ctx = contextWithSpan(ctx, newLogtreeSpan(logger))
		}

		Trace(ctx).Printf("RPC invoked: unary request: %s", info.FullMethod)

		// Perform authentication check and inject PeerInfo.
		pi, err := authenticationCheck(ctx, a, info.FullMethod)
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
