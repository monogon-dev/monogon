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

import (
	"context"
	"fmt"
	"testing"
	"time"

	"source.monogon.dev/osbase/logtree"
)

// waitSettle waits until the supervisor reaches a 'settled' state - ie., one
// where no actions have been performed for a number of GC cycles.
// This is used in tests only.
func (s *supervisor) waitSettle(ctx context.Context) error {
	waiter := make(chan struct{})
	s.pReq <- &processorRequest{
		waitSettled: &processorRequestWaitSettled{
			waiter: waiter,
		},
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-waiter:
		return nil
	}
}

// waitSettleError wraps waitSettle to fail a test if an error occurs, eg. the
// context is canceled.
func (s *supervisor) waitSettleError(ctx context.Context, t *testing.T) {
	err := s.waitSettle(ctx)
	if err != nil {
		t.Fatalf("waitSettle: %v", err)
	}
}

func runnableBecomesHealthy(healthy, done chan struct{}) Runnable {
	return func(ctx context.Context) error {
		Signal(ctx, SignalHealthy)

		go func() {
			if healthy != nil {
				healthy <- struct{}{}
			}
		}()

		<-ctx.Done()

		if done != nil {
			done <- struct{}{}
		}

		return ctx.Err()
	}
}

func runnableSpawnsMore(healthy, done chan struct{}, levels int) Runnable {
	return func(ctx context.Context) error {
		if levels > 0 {
			err := RunGroup(ctx, map[string]Runnable{
				"a": runnableSpawnsMore(nil, nil, levels-1),
				"b": runnableSpawnsMore(nil, nil, levels-1),
			})
			if err != nil {
				return err
			}
		}

		Signal(ctx, SignalHealthy)

		go func() {
			if healthy != nil {
				healthy <- struct{}{}
			}
		}()

		<-ctx.Done()

		if done != nil {
			done <- struct{}{}
		}
		return ctx.Err()
	}
}

// rc is a Remote Controlled runnable. It is a generic runnable used for
// testing the supervisor.
type rc struct {
	req chan rcRunnableRequest
}

type rcRunnableRequest struct {
	cmd    rcRunnableCommand
	stateC chan rcRunnableState
}

type rcRunnableCommand int

const (
	rcRunnableCommandBecomeHealthy rcRunnableCommand = iota
	rcRunnableCommandBecomeDone
	rcRunnableCommandDie
	rcRunnableCommandPanic
	rcRunnableCommandState
)

type rcRunnableState int

const (
	rcRunnableStateNew rcRunnableState = iota
	rcRunnableStateHealthy
	rcRunnableStateDone
)

func (r *rc) becomeHealthy() {
	r.req <- rcRunnableRequest{cmd: rcRunnableCommandBecomeHealthy}
}

func (r *rc) becomeDone() {
	r.req <- rcRunnableRequest{cmd: rcRunnableCommandBecomeDone}
}
func (r *rc) die() {
	r.req <- rcRunnableRequest{cmd: rcRunnableCommandDie}
}

func (r *rc) panic() {
	r.req <- rcRunnableRequest{cmd: rcRunnableCommandPanic}
}

func (r *rc) state() rcRunnableState {
	c := make(chan rcRunnableState)
	r.req <- rcRunnableRequest{
		cmd:    rcRunnableCommandState,
		stateC: c,
	}
	return <-c
}

func (r *rc) waitState(s rcRunnableState) {
	// This is poll based. Making it non-poll based would make the RC runnable
	// logic a bit more complex for little gain.
	for {
		got := r.state()
		if got == s {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func newRC() *rc {
	return &rc{
		req: make(chan rcRunnableRequest),
	}
}

// Remote Controlled Runnable
func (r *rc) runnable() Runnable {
	return func(ctx context.Context) error {
		state := rcRunnableStateNew

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case r := <-r.req:
				switch r.cmd {
				case rcRunnableCommandBecomeHealthy:
					Signal(ctx, SignalHealthy)
					state = rcRunnableStateHealthy
				case rcRunnableCommandBecomeDone:
					Signal(ctx, SignalDone)
					state = rcRunnableStateDone
				case rcRunnableCommandDie:
					return fmt.Errorf("died on request")
				case rcRunnableCommandPanic:
					panic("at the disco")
				case rcRunnableCommandState:
					r.stateC <- state
				}
			}
		}
	}
}

func TestSimple(t *testing.T) {
	h1 := make(chan struct{})
	d1 := make(chan struct{})
	h2 := make(chan struct{})
	d2 := make(chan struct{})

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	s := New(ctx, func(ctx context.Context) error {
		err := RunGroup(ctx, map[string]Runnable{
			"one": runnableBecomesHealthy(h1, d1),
			"two": runnableBecomesHealthy(h2, d2),
		})
		if err != nil {
			return err
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	}, WithPropagatePanic)

	// Expect both to start running.
	s.waitSettleError(ctx, t)
	select {
	case <-h1:
	default:
		t.Fatalf("runnable 'one' didn't start")
	}
	select {
	case <-h2:
	default:
		t.Fatalf("runnable 'one' didn't start")
	}
}

func TestSimpleFailure(t *testing.T) {
	h1 := make(chan struct{})
	d1 := make(chan struct{})
	two := newRC()

	ctx, ctxC := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxC()
	s := New(ctx, func(ctx context.Context) error {
		err := RunGroup(ctx, map[string]Runnable{
			"one": runnableBecomesHealthy(h1, d1),
			"two": two.runnable(),
		})
		if err != nil {
			return err
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	}, WithPropagatePanic)
	s.waitSettleError(ctx, t)

	two.becomeHealthy()
	s.waitSettleError(ctx, t)
	// Expect one to start running.
	select {
	case <-h1:
	default:
		t.Fatalf("runnable 'one' didn't start")
	}

	// Kill off two, one should restart.
	two.die()
	<-d1

	// And one should start running again.
	s.waitSettleError(ctx, t)
	select {
	case <-h1:
	default:
		t.Fatalf("runnable 'one' didn't restart")
	}
}

func TestDeepFailure(t *testing.T) {
	h1 := make(chan struct{})
	d1 := make(chan struct{})
	two := newRC()

	ctx, ctxC := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxC()
	s := New(ctx, func(ctx context.Context) error {
		err := RunGroup(ctx, map[string]Runnable{
			"one": runnableSpawnsMore(h1, d1, 5),
			"two": two.runnable(),
		})
		if err != nil {
			return err
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	}, WithPropagatePanic)

	two.becomeHealthy()
	s.waitSettleError(ctx, t)
	// Expect one to start running.
	select {
	case <-h1:
	default:
		t.Fatalf("runnable 'one' didn't start")
	}

	// Kill off two, one should restart.
	two.die()
	<-d1

	// And one should start running again.
	s.waitSettleError(ctx, t)
	select {
	case <-h1:
	default:
		t.Fatalf("runnable 'one' didn't restart")
	}
}

func TestPanic(t *testing.T) {
	h1 := make(chan struct{})
	d1 := make(chan struct{})
	two := newRC()

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	s := New(ctx, func(ctx context.Context) error {
		err := RunGroup(ctx, map[string]Runnable{
			"one": runnableBecomesHealthy(h1, d1),
			"two": two.runnable(),
		})
		if err != nil {
			return err
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	})

	two.becomeHealthy()
	s.waitSettleError(ctx, t)
	// Expect one to start running.
	select {
	case <-h1:
	default:
		t.Fatalf("runnable 'one' didn't start")
	}

	// Kill off two, one should restart.
	two.panic()
	<-d1

	// And one should start running again.
	s.waitSettleError(ctx, t)
	select {
	case <-h1:
	default:
		t.Fatalf("runnable 'one' didn't restart")
	}
}

func TestMultipleLevelFailure(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	New(ctx, func(ctx context.Context) error {
		err := RunGroup(ctx, map[string]Runnable{
			"one": runnableSpawnsMore(nil, nil, 4),
			"two": runnableSpawnsMore(nil, nil, 4),
		})
		if err != nil {
			return err
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	}, WithPropagatePanic)
}

func TestBackoff(t *testing.T) {
	one := newRC()

	ctx, ctxC := context.WithTimeout(context.Background(), 20*time.Second)
	defer ctxC()

	s := New(ctx, func(ctx context.Context) error {
		if err := Run(ctx, "one", one.runnable()); err != nil {
			return err
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	}, WithPropagatePanic)

	one.becomeHealthy()
	// Die a bunch of times in a row, this brings up the next exponential
	// backoff to over a second.
	for i := 0; i < 4; i += 1 {
		one.die()
		one.waitState(rcRunnableStateNew)
	}
	// Measure how long it takes for the runnable to respawn after a number of
	// failures
	start := time.Now()
	one.die()
	one.becomeHealthy()
	one.waitState(rcRunnableStateHealthy)
	taken := time.Since(start)
	if taken < 1*time.Second {
		t.Errorf("Runnable took %v to restart, wanted at least a second from backoff", taken)
	}

	s.waitSettleError(ctx, t)
	// Now that we've become healthy, die again. Becoming healthy resets the backoff.
	start = time.Now()
	one.die()
	one.becomeHealthy()
	one.waitState(rcRunnableStateHealthy)
	taken = time.Since(start)
	if taken > 1*time.Second || taken < 100*time.Millisecond {
		t.Errorf("Runnable took %v to restart, wanted at least 100ms from backoff and at most 1s from backoff reset", taken)
	}
}

// TestCancelRestart fails a runnable, but before its restart timeout expires,
// also fails its parent. This should cause cancelation of the restart timeout.
func TestCancelRestart(t *testing.T) {
	startedOuter := make(chan struct{})
	failInner := make(chan struct{})
	failOuter := make(chan struct{})

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	New(ctx, func(ctx context.Context) error {
		<-startedOuter
		err := Run(ctx, "inner", func(ctx context.Context) error {
			<-failInner
			return fmt.Errorf("failed")
		})
		if err != nil {
			return err
		}
		<-failOuter
		return fmt.Errorf("failed")
	}, WithPropagatePanic)

	startedOuter <- struct{}{}
	failInner <- struct{}{}
	time.Sleep(10 * time.Millisecond)
	// Before the inner runnable has restarted, fail the outer runnable.
	failOuter <- struct{}{}

	start := time.Now()
	startedOuter <- struct{}{}
	taken := time.Since(start)
	// With the default backoff parameters, the initial backoff time is
	// 0.5s +- 0.25s because of randomization. If the inner restart timer is not
	// canceled, then it takes twice as long.
	if taken > 1*time.Second {
		t.Errorf("Runnable took %v to restart, wanted at most 1s", taken)
	}
}

// TestDoneDelay test that a node is only considered restartable once it has
// returned, not already when it has signaled Done. Otherwise, we can get into
// an inconsistent state and for example panic because the node no longer
// exists once the runnable returns.
func TestDoneDelay(t *testing.T) {
	startedInner := make(chan struct{})
	failOuter := make(chan struct{})

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	New(ctx, func(ctx context.Context) error {
		err := Run(ctx, "inner", func(ctx context.Context) error {
			Signal(ctx, SignalHealthy)
			Signal(ctx, SignalDone)
			<-startedInner
			time.Sleep(10 * time.Millisecond)
			return nil
		})
		if err != nil {
			return err
		}
		<-failOuter
		return fmt.Errorf("failed")
	}, WithPropagatePanic)

	startedInner <- struct{}{}
	failOuter <- struct{}{}
	time.Sleep(20 * time.Millisecond)
}

// TestCancelDoneSibling tests that a node in state DONE is restarted if it is
// canceled because a sibling has died.
func TestCancelDoneSibling(t *testing.T) {
	innerRunning := make(chan struct{})
	innerExit := make(chan struct{})
	failSibling := make(chan struct{})

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	New(ctx, func(ctx context.Context) error {
		err := RunGroup(ctx, map[string]Runnable{
			"done": func(ctx context.Context) error {
				err := Run(ctx, "inner", func(ctx context.Context) error {
					<-innerRunning
					<-ctx.Done()
					<-innerExit
					return ctx.Err()
				})
				if err != nil {
					return err
				}
				Signal(ctx, SignalHealthy)
				Signal(ctx, SignalDone)
				return nil
			},
			"sibling": func(ctx context.Context) error {
				<-failSibling
				return fmt.Errorf("failed")
			},
		})
		if err != nil {
			return err
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	}, WithPropagatePanic)

	innerRunning <- struct{}{}
	failSibling <- struct{}{}
	// The inner node should exit and start running again.
	innerExit <- struct{}{}
	innerRunning <- struct{}{}
}

// TestResilience throws some curveballs at the supervisor - either programming
// errors or high load. It then ensures that another runnable is running, and
// that it restarts on its sibling failure.
func TestResilience(t *testing.T) {
	// request/response channel for testing liveness of the 'one' runnable
	req := make(chan chan struct{})

	// A runnable that responds on the 'req' channel.
	one := func(ctx context.Context) error {
		Signal(ctx, SignalHealthy)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case r := <-req:
				r <- struct{}{}
			}
		}
	}
	oneSibling := newRC()

	oneTest := func() {
		timeout := time.NewTicker(1000 * time.Millisecond)
		ping := make(chan struct{})
		req <- ping
		select {
		case <-ping:
		case <-timeout.C:
			t.Fatalf("one ping response timeout")
		}
		timeout.Stop()
	}

	// A nasty runnable that calls Signal with the wrong context (this is a
	// programming error)
	two := func(ctx context.Context) error {
		Signal(context.TODO(), SignalHealthy)
		return nil
	}

	// A nasty runnable that calls Signal wrong (this is a programming error).
	three := func(ctx context.Context) error {
		Signal(ctx, SignalDone)
		return nil
	}

	// A nasty runnable that runs in a busy loop (this is a programming error).
	four := func(ctx context.Context) error {
		for {
			time.Sleep(0)
		}
	}

	// A nasty runnable that keeps creating more runnables.
	five := func(ctx context.Context) error {
		i := 1
		for {
			err := Run(ctx, fmt.Sprintf("r%d", i), runnableSpawnsMore(nil, nil, 2))
			if err != nil {
				return err
			}

			time.Sleep(100 * time.Millisecond)
			i += 1
		}
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	New(ctx, func(ctx context.Context) error {
		RunGroup(ctx, map[string]Runnable{
			"one":        one,
			"oneSibling": oneSibling.runnable(),
		})
		rs := map[string]Runnable{
			"two": two, "three": three, "four": four, "five": five,
		}
		for k, v := range rs {
			if err := Run(ctx, k, v); err != nil {
				return err
			}
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		return nil
	})

	// Five rounds of letting one run, then restarting it.
	for i := 0; i < 5; i += 1 {
		oneSibling.becomeHealthy()
		oneSibling.waitState(rcRunnableStateHealthy)

		// 'one' should work for at least a second.
		deadline := time.Now().Add(1 * time.Second)
		for {
			if time.Now().After(deadline) {
				break
			}

			oneTest()
		}

		// Killing 'oneSibling' should restart one.
		oneSibling.panic()
	}
	// Make sure 'one' is still okay.
	oneTest()
}

// TestSubLoggers exercises the reserved/sub-logger functionality of runnable
// nodes. It ensures a sub-logger and runnable cannot have colliding names, and
// that logging actually works.
func TestSubLoggers(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	errCA := make(chan error)
	errCB := make(chan error)
	lt := logtree.New()
	New(ctx, func(ctx context.Context) error {
		err := RunGroup(ctx, map[string]Runnable{
			// foo will first create a sublogger, then attempt to create a
			// colliding runnable.
			"foo": func(ctx context.Context) error {
				sl, err := SubLogger(ctx, "dut")
				if err != nil {
					errCA <- fmt.Errorf("creating sub-logger: %w", err)
					return nil
				}
				sl.Infof("hello from foo.dut")
				err = Run(ctx, "dut", runnableBecomesHealthy(nil, nil))
				if err == nil {
					errCA <- fmt.Errorf("creating colliding runnable should have failed")
					return nil
				}
				Signal(ctx, SignalHealthy)
				Signal(ctx, SignalDone)
				errCA <- nil
				return nil
			},
		})
		if err != nil {
			return err
		}
		_, err = SubLogger(ctx, "foo")
		if err == nil {
			errCB <- fmt.Errorf("creating collising sub-logger should have failed")
			return nil
		}
		Signal(ctx, SignalHealthy)
		Signal(ctx, SignalDone)
		errCB <- nil
		return nil
	}, WithPropagatePanic, WithExistingLogtree(lt))

	err := <-errCA
	if err != nil {
		t.Fatalf("from root.foo: %v", err)
	}
	err = <-errCB
	if err != nil {
		t.Fatalf("from root: %v", err)
	}

	// Now enure that the expected message appears in the logtree.
	dn := logtree.DN("root.foo.dut")
	r, err := lt.Read(dn, logtree.WithBacklog(logtree.BacklogAllAvailable))
	if err != nil {
		t.Fatalf("logtree read failed: %v", err)
	}
	defer r.Close()
	found := false
	for _, e := range r.Backlog {
		if e.DN != dn {
			continue
		}
		if e.Leveled == nil {
			continue
		}
		if e.Leveled.MessagesJoined() != "hello from foo.dut" {
			continue
		}
		found = true
		break
	}
	if !found {
		t.Fatalf("did not find expected logline in %s", dn)
	}
}

func TestMetrics(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Build a supervision tree with 'wait'/step channels per runnable:
	//
	// root: wait, start one, wait, healthy
	//   one: wait, start two, crash, wait, start two, healthy, wait, done
	//     two: wait, healthy, run forever
	//
	// This tree allows us to exercise a few flows, like two getting canceled when
	// one crashes, runnables returning done, runnables staying healthy, etc.

	stepRoot := make(chan struct{})
	stepOne := make(chan struct{})
	stepTwo := make(chan struct{})
	m := InMemoryMetrics{}

	New(ctx, func(ctx context.Context) error {
		<-stepRoot

		attempts := 0
		Run(ctx, "one", func(ctx context.Context) error {
			<-stepOne
			Run(ctx, "two", func(ctx context.Context) error {
				<-stepTwo
				Signal(ctx, SignalHealthy)
				<-ctx.Done()
				return ctx.Err()
			})
			if attempts == 0 {
				attempts += 1
				return fmt.Errorf("failed")
			}
			Signal(ctx, SignalHealthy)
			<-stepOne
			Signal(ctx, SignalDone)
			return nil
		})

		<-stepRoot
		Signal(ctx, SignalHealthy)
		return nil
	}, WithPropagatePanic, WithMetrics(&m))

	// expectDN waits a second until a given DN is at a given state and fails the
	// test otherwise.
	expectDN := func(dn string, state NodeState) {
		t.Helper()
		start := time.Now()
		for {
			snap := m.DNs()
			if _, ok := snap[dn]; !ok {
				if time.Since(start) > time.Second {
					t.Fatalf("No DN %q", dn)
				} else {
					time.Sleep(100 * time.Millisecond)
					continue
				}
			}
			if want, got := state, snap[dn].State; want != got {
				if time.Since(start) > time.Second {
					t.Fatalf("Expected %q to be %s, got %s", dn, want, got)
				} else {
					time.Sleep(100 * time.Millisecond)
					continue
				}
			}
			break
		}
	}

	// Make progress thorugh the runnable tree and check expected states.

	expectDN("root", NodeStateNew)

	stepRoot <- struct{}{}
	expectDN("root", NodeStateNew)
	expectDN("root.one", NodeStateNew)

	stepOne <- struct{}{}
	stepTwo <- struct{}{}
	expectDN("root", NodeStateNew)
	expectDN("root.one", NodeStateDead)
	expectDN("root.one.two", NodeStateCanceled)

	stepOne <- struct{}{}
	expectDN("root", NodeStateNew)
	expectDN("root.one", NodeStateHealthy)
	expectDN("root.one.two", NodeStateNew)

	stepOne <- struct{}{}
	expectDN("root", NodeStateNew)
	expectDN("root.one", NodeStateDone)
	expectDN("root.one.two", NodeStateNew)

	stepTwo <- struct{}{}
	expectDN("root", NodeStateNew)
	expectDN("root.one", NodeStateDone)
	expectDN("root.one.two", NodeStateHealthy)
}

func ExampleNew() {
	// Minimal runnable that is immediately done.
	childC := make(chan struct{})
	child := func(ctx context.Context) error {
		Signal(ctx, SignalHealthy)
		close(childC)
		Signal(ctx, SignalDone)
		return nil
	}

	// Start a supervision tree with a root runnable.
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	New(ctx, func(ctx context.Context) error {
		err := Run(ctx, "child", child)
		if err != nil {
			return fmt.Errorf("could not run 'child': %w", err)
		}
		Signal(ctx, SignalHealthy)

		t := time.NewTicker(time.Second)
		defer t.Stop()

		// Do something in the background, and exit on context cancel.
		for {
			select {
			case <-t.C:
				fmt.Printf("tick!")
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	// root.child will close this channel.
	<-childC
}
