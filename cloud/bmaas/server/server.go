// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/bmaas/bmdb/webug"
	apb "source.monogon.dev/cloud/bmaas/server/api"
	"source.monogon.dev/cloud/lib/component"
)

type Config struct {
	Component component.ComponentConfig
	BMDB      bmdb.BMDB
	Webug     webug.Config

	// PublicListenAddress is the address at which the 'public' (agent-facing) gRPC
	// server listener will run.
	PublicListenAddress string
}

// TODO(q3k): factor this out to BMDB library?
func runtimeInfo() string {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "UNKNOWN"
	}
	return fmt.Sprintf("host %s", hostname)
}

func (c *Config) RegisterFlags() {
	c.Component.RegisterFlags("srv")
	c.BMDB.ComponentName = "srv"
	c.BMDB.RuntimeInfo = runtimeInfo()
	c.BMDB.Database.RegisterFlags("bmdb")
	c.Webug.RegisterFlags()

	flag.StringVar(&c.PublicListenAddress, "srv_public_grpc_listen_address", ":8080", "Address to listen at for public/user gRPC connections for bmdbsrv")
}

type Server struct {
	Config Config

	// ListenGRPC will contain the address at which the internal gRPC server is
	// listening after .Start() has been called. This can differ from the configured
	// value if the configuration requests any port (via :0).
	ListenGRPC string
	// ListenPublic will contain the address at which the 'public' (agent-facing)
	// gRPC server is lsitening after .Start() has been called.
	ListenPublic string

	bmdb  *bmdb.Connection
	acsvc *agentCallbackService

	sessionC chan *bmdb.Session
}

// sessionWorker emits a valid BMDB session to sessionC as long as ctx is active.
func (s *Server) sessionWorker(ctx context.Context) {
	var session *bmdb.Session
	for {
		if session == nil || session.Expired() {
			klog.Infof("Starting new session...")
			bo := backoff.NewExponentialBackOff()
			err := backoff.Retry(func() error {
				var err error
				session, err = s.bmdb.StartSession(ctx, bmdb.SessionOption{Processor: metrics.ProcessorBMSRV})
				if err != nil {
					klog.Errorf("Failed to start session: %v", err)
					return err
				} else {
					return nil
				}
			}, backoff.WithContext(bo, ctx))
			if err != nil {
				// If something's really wrong just crash.
				klog.Exitf("Gave up on starting session: %v", err)
			}
			klog.Infof("New session: %s", session.UUID)
		}

		select {
		case <-ctx.Done():
			return
		case s.sessionC <- session:
		}
	}
}

func (s *Server) session(ctx context.Context) (*bmdb.Session, error) {
	select {
	case sess := <-s.sessionC:
		return sess, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *Server) startPublic(ctx context.Context) {
	g := grpc.NewServer(s.Config.Component.GRPCServerOptionsPublic()...)
	lis, err := net.Listen("tcp", s.Config.PublicListenAddress)
	if err != nil {
		klog.Exitf("Could not listen: %v", err)
	}
	s.ListenPublic = lis.Addr().String()
	apb.RegisterAgentCallbackServer(g, s.acsvc)
	reflection.Register(g)

	klog.Infof("Public API listening on %s", s.ListenPublic)
	go func() {
		err := g.Serve(lis)
		if !errors.Is(err, ctx.Err()) {
			klog.Exitf("Public gRPC serve failed: %v", err)
		}
	}()
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
		if !errors.Is(err, ctx.Err()) {
			klog.Exitf("Internal gRPC serve failed: %v", err)
		}
	}()
}

// Start the BMaaS Server in background goroutines. This should only be called
// once. The process will exit with debug logs if starting the server failed.
func (s *Server) Start(ctx context.Context) {
	reg := s.Config.Component.PrometheusRegistry()
	s.Config.BMDB.EnableMetrics(reg)
	s.Config.Component.StartPrometheus(ctx)

	conn, err := s.Config.BMDB.Open(true)
	if err != nil {
		klog.Exitf("Failed to connect to BMDB: %v", err)
	}
	s.acsvc = &agentCallbackService{
		s: s,
	}
	s.bmdb = conn
	s.sessionC = make(chan *bmdb.Session)
	go s.sessionWorker(ctx)
	s.startInternalGRPC(ctx)
	s.startPublic(ctx)
	go func() {
		if err := s.Config.Webug.Start(ctx, conn); err != nil && !errors.Is(err, ctx.Err()) {
			klog.Exitf("Failed to start webug: %v", err)
		}
	}()
}
