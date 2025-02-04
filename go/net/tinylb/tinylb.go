// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package tinylb implements a small and simple userland round-robin load
// balancer, mostly for TCP connections. However, it is entirely
// protocol-agnostic, and only expects net.Listener and net.Conn objects.
//
// Apart from the simple act of load-balancing across a set of backends, tinylb
// also automatically and immediately closes all open connections to backend
// targets that have been removed from the set. This is perhaps not the ideal
// behaviour for user-facing services, but it's the sort of behaviour that works
// very well for machine-to-machine communication.
package tinylb

import (
	"context"
	"io"
	"net"
	"sync"

	"source.monogon.dev/go/types/mapsets"
	"source.monogon.dev/osbase/event"
	"source.monogon.dev/osbase/supervisor"
)

// Backend is to be implemented by different kinds of loadbalancing backends, eg.
// one per network protocol.
type Backend interface {
	// TargetName returns the 'target name' of the backend, which is _not_ the same
	// as its key in the BackendSet. Instead, the TargetName should uniquely identify
	// some backend address, and will be used to figure out that while a backend
	// might still exist, its target address has changed - and thus, all existing
	// connections to the previous target address should be closed.
	//
	// For simple load balancing backends, this could be the connection string used
	// to connect to the backend.
	TargetName() string
	// Dial returns a new connection to this backend.
	Dial() (net.Conn, error)
}

// BackendSet is the main structure used to provide the current set of backends
// that should be targeted by tinylb. The key is some unique backend identifier.
type BackendSet = mapsets.OrderedMap[string, Backend]

// SimpleTCPBackend implements Backend for trivial TCP-based backends.
type SimpleTCPBackend struct {
	Remote string
}

func (t *SimpleTCPBackend) TargetName() string {
	return t.Remote
}

func (t *SimpleTCPBackend) Dial() (net.Conn, error) {
	return net.Dial("tcp", t.Remote)
}

// Server is a tiny round-robin loadbalancer for net.Listener/net.Conn compatible
// protocols.
//
// All fields must be set before the loadbalancer can be run.
type Server struct {
	// Provider is some event Value which provides the current BackendSet for the
	// loadbalancer to use. As the BackendSet is updated, the internal loadbalancing
	// algorithm will adjust to the updated set, and any connections to backend
	// TargetNames that are not present in the set anymore will be closed.
	Provider event.Value[BackendSet]
	// Listener is where the loadbalancer will listen on. After the loadbalancer
	// exits, this listener will be closed.
	Listener net.Listener
}

// Run the loadbalancer in a superervisor.Runnable and block until canceled.
// Because the Server's Listener will be closed after exit, it should be opened
// in the same runnable as this function is then started.
func (s *Server) Run(ctx context.Context) error {
	// Connection pool used to track connections/backends.
	pool := newConnectionPool()

	// Current backend set and its lock.
	var curSetMu sync.RWMutex
	var curSet BackendSet

	// Close listener on exit.
	go func() {
		<-ctx.Done()
		s.Listener.Close()
	}()

	// The acceptor is runs the main Accept() loop on the given Listener.
	err := supervisor.Run(ctx, "acceptor", func(ctx context.Context) error {
		// This doesn't need a lock, as it doesn't read any fields of curSet.
		it := curSet.Cycle()

		supervisor.Signal(ctx, supervisor.SignalHealthy)

		for {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			conn, err := s.Listener.Accept()
			if err != nil {
				return err
			}

			// Get next backend.
			curSetMu.RLock()
			id, backend, ok := it.Next()
			curSetMu.RUnlock()

			if !ok {
				supervisor.Logger(ctx).Warningf("Balancing %s: failed due to empty backend set", conn.RemoteAddr().String())
				conn.Close()
				continue
			}
			// Dial backend.
			r, err := backend.Dial()
			if err != nil {
				supervisor.Logger(ctx).Warningf("Balancing %s: failed due to backend %s error: %v", conn.RemoteAddr(), id, err)
				conn.Close()
				continue
			}
			// Insert connection/target name into connectionPool.
			target := backend.TargetName()
			cid := pool.Insert(target, r)

			// Pipe data. Close both connections on any side failing.
			go func() {
				defer conn.Close()
				defer pool.CloseConn(cid)
				io.Copy(r, conn)
			}()
			go func() {
				defer conn.Close()
				defer pool.CloseConn(cid)
				io.Copy(conn, r)
			}()
		}
	})
	if err != nil {
		return err
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)

	// Update curSet from Provider.
	w := s.Provider.Watch()
	defer w.Close()
	for {
		set, err := w.Get(ctx)
		if err != nil {
			return err
		}

		// Did we lose a backend? If so, kill all connections going through it.

		// First, gather a map from TargetName to backend ID for the current set.
		curTargets := make(map[string]string)
		curSetMu.Lock()
		for _, kv := range curSet.Values() {
			curTargets[kv.Value.TargetName()] = kv.Key
		}
		curSetMu.Unlock()

		// Then, gather it for the new set.
		targets := make(map[string]string)
		for _, kv := range set.Values() {
			targets[kv.Value.TargetName()] = kv.Key
		}

		// Then, if we have any target name in the connection pool that's not in the new
		// set, close all of its connections.
		for _, target := range pool.Targets() {
			if _, ok := targets[target]; ok {
				continue
			}
			// Use curTargets just for displaying the name of the backend that has changed.
			supervisor.Logger(ctx).Infof("Backend %s / target %s removed, killing connections", curTargets[target], target)
			pool.CloseTarget(target)
		}

		// Tell about new backend set and actually replace it.
		supervisor.Logger(ctx).Infof("New backend set (%d backends):", len(set.Keys()))
		for _, kv := range set.Values() {
			supervisor.Logger(ctx).Infof(" - %s, target %s", kv.Key, kv.Value.TargetName())
		}
		curSetMu.Lock()
		curSet.Replace(&set)
		curSetMu.Unlock()
	}
}
