package server

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"

	apb "source.monogon.dev/cloud/api"
	"source.monogon.dev/cloud/apigw/model"
	"source.monogon.dev/cloud/lib/component"
)

// Config is the main configuration of the apigw server. It's usually populated
// from flags via RegisterFlags, but can also be set manually (eg. in tests).
type Config struct {
	Component component.ComponentConfig
	Database  component.CockroachConfig

	PublicListenAddress string
}

// RegisterFlags registers the component configuration to be provided by flags.
// This must be called exactly once before then calling flags.Parse().
func (c *Config) RegisterFlags() {
	c.Component.RegisterFlags("apigw")
	c.Database.RegisterFlags("apigw_db")
	flag.StringVar(&c.PublicListenAddress, "apigw_public_grpc_listen_address", ":8080", "Address to listen at for public/user gRPC connections for apigw")
}

// Server runs the apigw server. It listens on two interfaces:
//  - Internal gRPC, which is authenticated using TLS and authorized by CA. This
//    is to be used for internal RPCs, eg. management/debug.
//  - Public gRPC-Web, which is currently unauthenticated.
type Server struct {
	Config Config

	// ListenGRPC will contain the address at which the internal gRPC server is
	// listening after .Start() has been called. This can differ from the configured
	// value if the configuration requests any port (via :0).
	ListenGRPC string
	// ListenPublic will contain the address at which the public API server is
	// listening after .Start() has been called. This can differ from the configured
	// value if the configuration requests any port (via :0).
	ListenPublic string
}

func (s *Server) startInternalGRPC(ctx context.Context) {
	g := grpc.NewServer(s.Config.Component.GRPCServerOptions()...)
	lis, err := net.Listen("tcp", s.Config.Component.GRPCListenAddress)
	if err != nil {
		klog.Exitf("Could not listen: %v", err)
	}
	s.ListenGRPC = lis.Addr().String()

	reflection.Register(g)

	klog.Infof("Internal gRPC listening on %s", s.ListenGRPC)
	go func() {
		err := g.Serve(lis)
		if err != ctx.Err() {
			klog.Exitf("Internal gRPC serve failed: %v", err)
		}
	}()
}

func (s *Server) startPublic(ctx context.Context) {
	g := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	lis, err := net.Listen("tcp", s.Config.PublicListenAddress)
	if err != nil {
		klog.Exitf("Could not listen: %v", err)
	}
	s.ListenPublic = lis.Addr().String()

	reflection.Register(g)
	apb.RegisterIAMServer(g, s)

	wrapped := grpcweb.WrapServer(g)
	server := http.Server{
		Addr:    s.Config.PublicListenAddress,
		Handler: http.HandlerFunc(wrapped.ServeHTTP),
	}
	klog.Infof("Public API listening on %s", s.ListenPublic)
	go func() {
		err := server.Serve(lis)
		if err != ctx.Err() {
			klog.Exitf("Public API serve failed: %v", err)
		}
	}()
}

// Start runs the two listeners of the server. The process will fail (via
// klog.Exit) if any of the listeners/servers fail to start.
func (s *Server) Start(ctx context.Context) {
	if s.Config.Database.Migrations == nil {
		klog.Infof("Using default migrations source.")
		m, err := model.MigrationsSource()
		if err != nil {
			klog.Exitf("failed to prepare migrations source: %v", err)
		}
		s.Config.Database.Migrations = m
	}

	klog.Infof("Running migrations...")
	if err := s.Config.Database.MigrateUp(); err != nil {
		klog.Exitf("Migrations failed: %v", err)
	}
	klog.Infof("Migrations done.")
	s.startInternalGRPC(ctx)
	s.startPublic(ctx)
}

func (s *Server) WhoAmI(ctx context.Context, req *apb.WhoAmIRequest) (*apb.WhoAmIResponse, error) {
	klog.Infof("req: %+v", req)
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
