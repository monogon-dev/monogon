// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package supervisor

// The service supervision library allows for writing of reliable,
// service-style software within a Metropolis node.  It builds upon the
// Erlang/OTP supervision tree system, adapted to be more Go-ish.  For detailed
// design see go/supervision.

import (
	"context"
	"fmt"
	"io"
	"sync"

	"source.monogon.dev/osbase/logtree"
)

// A Runnable is a function that will be run in a goroutine, and supervised
// throughout its lifetime. It can in turn start more runnables as its
// children, and those will form part of a supervision tree.
// The context passed to a runnable is very important and needs to be handled
// properly. It will be live (non-errored) as long as the runnable should be
// running, and canceled (ctx.Err() will be non-nil) when the supervisor wants
// it to exit. This means this context is also perfectly usable for performing
// any blocking operations.
type Runnable func(ctx context.Context) error

// RunGroup starts a set of runnables as a group. These runnables will run
// together, and if any one of them quits unexpectedly, the result will be
// canceled and restarted.
// The context here must be an existing Runnable context, and the spawned
// runnables will run under the node that this context represents.
func RunGroup(ctx context.Context, runnables map[string]Runnable) error {
	node, unlock := fromContext(ctx)
	defer unlock()
	return node.runGroup(runnables)
}

// Run starts a single runnable in its own group.
func Run(ctx context.Context, name string, runnable Runnable) error {
	return RunGroup(ctx, map[string]Runnable{
		name: runnable,
	})
}

// Signal tells the supervisor that the calling runnable has reached a certain
// state of its lifecycle. All runnables should SignalHealthy when they are
// ready with set up, running other child runnables and are now 'serving'.
func Signal(ctx context.Context, signal SignalType) {
	node, unlock := fromContext(ctx)
	defer unlock()
	node.signal(signal)
}

type SignalType int

const (
	// The runnable is healthy, done with setup, done with spawning more
	// Runnables, and ready to serve in a loop.  The runnable needs to check
	// the parent context and ensure that if that context is done, the runnable
	// exits.
	SignalHealthy SignalType = iota
	// The runnable is done - it does not need to run any loop. This is useful
	// for Runnables that only set up other child runnables. This runnable will
	// be restarted if a related failure happens somewhere in the supervision
	// tree.
	SignalDone
)

// supervisor represents and instance of the supervision system. It keeps track
// of a supervision tree and a request channel to its internal processor
// goroutine.
type supervisor struct {
	// mu guards the entire state of the supervisor.
	mu sync.RWMutex
	// root is the root node of the supervision tree, named 'root'. It
	// represents the Runnable started with the supervisor.New call.
	root *node
	// logtree is the main logtree exposed to runnables and used internally.
	logtree *logtree.LogTree
	// ilogger is the internal logger logging to "supervisor" in the logtree.
	ilogger logtree.LeveledLogger

	// pReq is an interface channel to the lifecycle processor of the
	// supervisor.
	pReq chan *processorRequest

	// propagate panics, ie. don't catch them.
	propagatePanic bool
}

// SupervisorOpt are runtime configurable options for the supervisor.
type SupervisorOpt func(s *supervisor)

// WithPropagatePanic prevents the Supervisor from catching panics in
// runnables and treating them as failures. This is useful to enable for
// testing and local debugging.
func WithPropagatePanic(s *supervisor) {
	s.propagatePanic = true
}

func WithExistingLogtree(lt *logtree.LogTree) SupervisorOpt {
	return func(s *supervisor) {
		s.logtree = lt
	}
}

// New creates a new supervisor with its root running the given root runnable.
// The given context can be used to cancel the entire supervision tree.
//
// For tests, we reccomend using TestHarness instead, which will also stream
// logs to stderr and take care of propagating root runnable errors to the test
// output.
func New(ctx context.Context, rootRunnable Runnable, opts ...SupervisorOpt) *supervisor {
	sup := &supervisor{
		logtree: logtree.New(),
		pReq:    make(chan *processorRequest),
	}

	for _, o := range opts {
		o(sup)
	}

	sup.ilogger = sup.logtree.MustLeveledFor("supervisor")
	sup.root = newNode("root", rootRunnable, sup, nil)

	go sup.processor(ctx)

	sup.pReq <- &processorRequest{
		schedule: &processorRequestSchedule{dn: "root"},
	}

	return sup
}

func Logger(ctx context.Context) logtree.LeveledLogger {
	node, unlock := fromContext(ctx)
	defer unlock()
	return node.sup.logtree.MustLeveledFor(logtree.DN(node.dn()))
}

func RawLogger(ctx context.Context) io.Writer {
	node, unlock := fromContext(ctx)
	defer unlock()
	return node.sup.logtree.MustRawFor(logtree.DN(node.dn()))
}

// SubLogger returns a LeveledLogger for a given name. The name is used to
// placed that logger within the logtree hierarchy. For example, if the
// runnable `root.foo` requests a SubLogger for name `bar`, the returned logger
// will log to `root.foo.bar` in the logging tree.
//
// An error is returned if the given name is invalid or conflicts with a child
// runnable of the current runnable. In addition, whenever a node uses a
// sub-logger with a given name, that name also becomes unavailable for use as
// a child runnable (no runnable and sub-logger may ever log into the same
// logtree DN).
func SubLogger(ctx context.Context, name string) (logtree.LeveledLogger, error) {
	node, unlock := fromContext(ctx)
	defer unlock()

	if _, ok := node.children[name]; ok {
		return nil, fmt.Errorf("name %q already in use by child runnable", name)
	}
	if !reNodeName.MatchString(name) {
		return nil, fmt.Errorf("sub-logger name %q is invalid", name)
	}
	node.reserved[name] = true

	dn := fmt.Sprintf("%s.%s", node.dn(), name)
	return node.sup.logtree.LeveledFor(logtree.DN(dn))
}

// MustSubLogger is a wrapper around SubLogger which panics on error. Errors
// should only happen due to invalid names, so as long as the given name is
// compile-time constant and valid, this function is safe to use.
func MustSubLogger(ctx context.Context, name string) logtree.LeveledLogger {
	l, err := SubLogger(ctx, name)
	if err != nil {
		panic(err)
	}
	return l
}
