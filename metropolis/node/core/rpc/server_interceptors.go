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
		pi, err := authenticationCheck(ss.Context(), a, info.FullMethod)
		if err != nil {
			return err
		}
		if logger != nil {
			s := newLogtreeSpan(logger)
			s.Printf("RPC received: streaming request: %s", info.FullMethod)
			ss = &spanServerStream{
				ServerStream: ss,
				span:         s,
			}

		}
		return handler(srv, pi.serverStream(ss))
	}
}

// unaryInterceptor implements the gRPC UnaryInterceptor interface for use with
// grpc.NewServer, based on an authenticationStrategy. It's applied to gRPC
// servers started within Metropolis, notably to the Curator.
func unaryInterceptor(logger logtree.LeveledLogger, a authenticationStrategy) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		pi, err := authenticationCheck(ctx, a, info.FullMethod)
		if err != nil {
			return nil, err
		}
		ctx = pi.apply(ctx)
		var s *logtreeSpan
		if logger != nil {
			s = newLogtreeSpan(logger)
			s.Printf("RPC received: unary request: %s", info.FullMethod)
			ctx = contextWithSpan(ctx, s)
		}
		resp, err = handler(pi.apply(ctx), req)
		if s != nil {
			if err != nil {
				s.Printf("RPC send: error: %v", err)
			} else {
				s.Printf("RPC send: ok, %s", protoMessagePretty(resp))
			}
		}
		return
	}
}
