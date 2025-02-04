// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package scruffy implements Scruffy, The Janitor.
//
// Scruffy is a BMaaS component which runs a bunch of important, housekeeping-ish
// processes that aren't tied to any particular provider and are mostly
// batch-oriented.
//
// Currently Scruffy just collects metrics from the BMDB.
package scruffy

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/bmaas/bmdb/webug"
	"source.monogon.dev/cloud/lib/component"
)

type Config struct {
	Component component.ComponentConfig
	BMDB      bmdb.BMDB
	Webug     webug.Config

	StatsRunnerRate time.Duration
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
	c.Component.RegisterFlags("scruffy")
	c.BMDB.ComponentName = "scruffy"
	c.BMDB.RuntimeInfo = runtimeInfo()
	c.BMDB.Database.RegisterFlags("bmdb")
	c.Webug.RegisterFlags()

	flag.DurationVar(&c.StatsRunnerRate, "scruffy_stats_collection_rate", time.Minute, "How often the stats collection loop will run against BMDB")
}

type Server struct {
	Config Config

	bmdb     *bmdb.Connection
	sessionC chan *bmdb.Session
}

func (s *Server) Start(ctx context.Context) {
	reg := s.Config.Component.PrometheusRegistry()
	s.Config.BMDB.EnableMetrics(reg)
	s.Config.Component.StartPrometheus(ctx)

	conn, err := s.Config.BMDB.Open(true)
	if err != nil {
		klog.Exitf("Failed to connect to BMDB: %v", err)
	}
	s.bmdb = conn
	s.sessionC = make(chan *bmdb.Session)
	go s.sessionWorker(ctx)

	bsr := newBMDBStatsRunner(s, reg)
	go bsr.run(ctx)

	hwr := newHWStatsRunner(s, reg)
	go hwr.run(ctx)

	go func() {
		if err := s.Config.Webug.Start(ctx, conn); err != nil && !errors.Is(err, ctx.Err()) {
			klog.Exitf("Failed to start webug: %v", err)
		}
	}()
}

// sessionWorker emits a valid BMDB session to sessionC as long as ctx is active.
//
// TODO(q3k): factor out into bmdb client lib
func (s *Server) sessionWorker(ctx context.Context) {
	var session *bmdb.Session
	for {
		if session == nil || session.Expired() {
			klog.Infof("Starting new session...")
			bo := backoff.NewExponentialBackOff()
			err := backoff.Retry(func() error {
				var err error
				session, err = s.bmdb.StartSession(ctx, bmdb.SessionOption{Processor: metrics.ProcessorScruffyStats})
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
