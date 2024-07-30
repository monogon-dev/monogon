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
	"errors"
	"fmt"
	"runtime/debug"
	"sort"
	"time"
)

// The processor maintains runnable goroutines - ie., when requested will start
// one, and then once it exists it will record the result and act accordingly.
// It is also responsible for detecting and acting upon supervision subtrees
// that need to be restarted after death (via a 'GC' process)

// processorRequest is a request for the processor. Only one of the fields can
// be set.
type processorRequest struct {
	schedule    *processorRequestSchedule
	died        *processorRequestDied
	waitSettled *processorRequestWaitSettled
}

// processorRequestSchedule requests that a given node's runnable be started.
type processorRequestSchedule struct {
	dn string
}

// processorRequestDied is a signal from a runnable goroutine that the runnable
// has died.
type processorRequestDied struct {
	dn  string
	err error
}

type processorRequestWaitSettled struct {
	waiter chan struct{}
}

// processor is the main processing loop.
func (s *supervisor) processor(ctx context.Context) {
	s.ilogger.Info("supervisor processor started")

	// Waiters waiting for the GC to be settled.
	var waiters []chan struct{}

	// The GC will run every millisecond if needed. Any time the processor
	// requests a change in the supervision tree (ie a death or a new runnable)
	// it will mark the state as dirty and run the GC on the next millisecond
	// cycle.
	gc := time.NewTicker(1 * time.Millisecond)
	defer gc.Stop()
	clean := true

	// How long has the GC been clean. This is used to notify 'settled' waiters.
	cleanCycles := 0

	markDirty := func() {
		clean = false
		cleanCycles = 0
	}

	for {
		select {
		case <-ctx.Done():
			s.ilogger.Infof("supervisor processor exiting: %v", ctx.Err())
			s.processKill()
			s.ilogger.Info("supervisor exited, starting liquidator to clean up remaining runnables...")
			go s.liquidator()
			return
		case <-gc.C:
			if !clean {
				s.processGC()
			}
			clean = true
			cleanCycles += 1

			// This threshold is somewhat arbitrary. It's a balance between
			// test speed and test reliability.
			if cleanCycles > 50 {
				for _, w := range waiters {
					close(w)
				}
				waiters = nil
			}
		case r := <-s.pReq:
			switch {
			case r.schedule != nil:
				s.processSchedule(r.schedule)
				markDirty()
			case r.died != nil:
				s.processDied(r.died)
				markDirty()
			case r.waitSettled != nil:
				waiters = append(waiters, r.waitSettled.waiter)
			default:
				panic(fmt.Errorf("unhandled request %+v", r))
			}
		}
	}
}

// The liquidator is a context-free goroutine which the supervisor starts after
// its context has been canceled. Its job is to take over listening on the
// processing channels that the supervisor processor would usually listen on,
// and implement the minimum amount of logic required to mark existing runnables
// as DEAD.
//
// It exits when all runnables have exited one way or another, and the
// supervision tree is well and truly dead. This will also be reflected by
// liveRunnables returning an empty list.
func (s *supervisor) liquidator() {
	for {
		r := <-s.pReq
		switch {
		case r.schedule != nil:
			s.ilogger.Infof("liquidator: refusing to schedule %s", r.schedule.dn)
			s.mu.Lock()
			n := s.nodeByDN(r.schedule.dn)
			n.state = NodeStateDead
			s.mu.Unlock()
		case r.died != nil:
			s.ilogger.Infof("liquidator: %s exited", r.died.dn)
			s.mu.Lock()
			n := s.nodeByDN(r.died.dn)
			n.state = NodeStateDead
			s.mu.Unlock()
		}
		live := s.liveRunnables()
		if len(live) == 0 {
			s.ilogger.Infof("liquidator: complete, all runnables dead or done")
			return
		}
	}
}

// liveRunnables returns a list of runnable DNs that aren't DONE/DEAD. This is
// used by the liquidator to figure out when its job is done, and by the
// TestHarness to know when to unblock the test cleanup function.
func (s *supervisor) liveRunnables() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// DFS through supervision tree, making not of live (non-DONE/DEAD runnables).
	var live []string
	seen := make(map[string]bool)
	q := []*node{s.root}
	for {
		if len(q) == 0 {
			break
		}

		// Pop from DFS queue.
		el := q[0]
		q = q[1:]

		// Skip already visited runnables (this shouldn't happen because the supervision
		// tree is, well, a tree - but better stay safe than get stuck in a loop).
		eldn := el.dn()
		if seen[eldn] {
			continue
		}
		seen[eldn] = true

		if el.state != NodeStateDead && el.state != NodeStateDone {
			live = append(live, eldn)
		}

		// Recurse.
		for _, child := range el.children {
			q = append(q, child)
		}
	}

	sort.Strings(live)
	return live
}

// processKill cancels all nodes in the supervision tree. This is only called
// right before exiting the processor, so they do not get automatically
// restarted.
func (s *supervisor) processKill() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Gather all context cancel functions.
	var cancels []func()
	queue := []*node{s.root}
	for {
		if len(queue) == 0 {
			break
		}

		cur := queue[0]
		queue = queue[1:]

		cancels = append(cancels, cur.ctxC)
		for _, c := range cur.children {
			queue = append(queue, c)
		}
	}

	// Call all context cancels.
	for _, c := range cancels {
		c()
	}
}

// processSchedule starts a node's runnable in a goroutine and records its
// output once it's done.
func (s *supervisor) processSchedule(r *processorRequestSchedule) {
	s.mu.Lock()
	defer s.mu.Unlock()

	n := s.nodeByDN(r.dn)
	go func() {
		if !s.propagatePanic {
			defer func() {
				if rec := recover(); rec != nil {
					s.pReq <- &processorRequest{
						died: &processorRequestDied{
							dn:  r.dn,
							err: fmt.Errorf("panic: %v, stacktrace: %s", rec, string(debug.Stack())),
						},
					}
				}
			}()
		}

		res := n.runnable(n.ctx)

		s.pReq <- &processorRequest{
			died: &processorRequestDied{
				dn:  r.dn,
				err: res,
			},
		}
	}()
}

// processDied records the result from a runnable goroutine, and updates its
// node state accordingly. If the result is a death and not an expected exit,
// related nodes (ie. children and group siblings) are canceled accordingly.
func (s *supervisor) processDied(r *processorRequestDied) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Okay, so a Runnable has quit. What now?
	n := s.nodeByDN(r.dn)
	ctx := n.ctx

	// Simple case: it was marked as Done and quit with no error.
	if n.state == NodeStateDone && r.err == nil {
		// Do nothing. This was supposed to happen. Keep the process as DONE.
		return
	}

	// Simple case: the context was canceled and the returned error is the
	// context error.
	if r.err != nil && ctx.Err() != nil && errors.Is(r.err, ctx.Err()) {
		// Mark the node as canceled successfully.
		n.state = NodeStateCanceled
		return
	}

	// Otherwise, the Runnable should not have died or quit. Handle
	// accordingly.
	err := r.err
	// A lack of returned error is also an error.
	if err == nil {
		err = fmt.Errorf("returned nil when %s", n.state)
	}

	s.ilogger.Errorf("%s: %v", n.dn(), err)
	// Mark as dead.
	n.state = NodeStateDead

	// Cancel that node's context, just in case something still depends on it.
	n.ctxC()

	// Cancel all siblings.
	if n.parent != nil {
		for name := range n.parent.groupSiblings(n.name) {
			if name == n.name {
				continue
			}
			sibling := n.parent.children[name]
			// TODO(q3k): does this need to run in a goroutine, ie. can a
			// context cancel block?
			sibling.ctxC()
		}
	}
}

// processGC runs the GC process. It's not really Garbage Collection, as in, it
// doesn't remove unnecessary tree nodes - but it does find nodes that need to
// be restarted, find the subset that can and then schedules them for running.
// As such, it's less of a Garbage Collector and more of a Necromancer.
// However, GC is a friendlier name.
func (s *supervisor) processGC() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// The 'GC' serves is the main business logic of the supervision tree. It
	// traverses a locked tree and tries to find subtrees that must be
	// restarted (because of a DEAD/CANCELED runnable). It then finds which of
	// these subtrees that should be restarted can be restarted, ie. which ones
	// are fully recursively DEAD/CANCELED. It also finds the smallest set of
	// largest subtrees that can be restarted, ie. if there's multiple DEAD
	// runnables that can be restarted at once, it will do so.

	// Phase one: Find all leaves.
	// This is a simple DFS that finds all the leaves of the tree, ie all nodes
	// that do not have children nodes.
	leaves := make(map[string]bool)
	queue := []*node{s.root}
	for {
		if len(queue) == 0 {
			break
		}
		cur := queue[0]
		queue = queue[1:]

		for _, c := range cur.children {
			queue = append([]*node{c}, queue...)
		}

		if len(cur.children) == 0 {
			leaves[cur.dn()] = true
		}
	}

	// Phase two: traverse tree from node to root and make note of all subtrees
	// that can be restarted.
	// A subtree is restartable/ready iff every node in that subtree is either
	// CANCELED, DEAD or DONE.  Such a 'ready' subtree can be restarted by the
	// supervisor if needed.

	// DNs that we already visited.
	visited := make(map[string]bool)
	// DNs whose subtrees are ready to be restarted.
	// These are all subtrees recursively - ie., root.a.a and root.a will both
	// be marked here.
	ready := make(map[string]bool)

	// We build a queue of nodes to visit, starting from the leaves.
	queue = []*node{}
	for l := range leaves {
		queue = append(queue, s.nodeByDN(l))
	}

	for {
		if len(queue) == 0 {
			break
		}

		cur := queue[0]
		curDn := cur.dn()

		queue = queue[1:]

		// Do we have a decision about our children?
		allVisited := true
		for _, c := range cur.children {
			if !visited[c.dn()] {
				allVisited = false
				break
			}
		}

		// If no decision about children is available, it means we ended up in
		// this subtree through some shorter path of a shorter/lower-order
		// leaf. There is a path to a leaf that's longer than the one that
		// caused this node to be enqueued. Easy solution: just push back the
		// current element and retry later.
		if !allVisited {
			// Push back to queue and wait for a decision later.
			queue = append(queue, cur)
			continue
		}

		// All children have been visited and we have an idea about whether
		// they're ready/restartable. All of the node's children must be
		// restartable in order for this node to be restartable.
		childrenReady := true
		var childrenNotReady []string
		for _, c := range cur.children {
			if !ready[c.dn()] {
				childrenNotReady = append(childrenNotReady, c.dn())
				childrenReady = false
				break
			}
		}

		// In addition to children, the node itself must be restartable (ie.
		// DONE, DEAD or CANCELED).
		curReady := false
		switch cur.state {
		case NodeStateDone:
			curReady = true
		case NodeStateCanceled:
			curReady = true
		case NodeStateDead:
			curReady = true
		default:
		}

		if cur.state == NodeStateDead && !childrenReady {
			s.ilogger.Warningf("Not restarting %s: children not ready to be restarted: %v", curDn, childrenNotReady)
		}

		// Note down that we have an opinion on this node, and note that
		// opinion down.
		visited[curDn] = true
		ready[curDn] = childrenReady && curReady

		// Now we can also enqueue the parent of this node for processing.
		if cur.parent != nil && !visited[cur.parent.dn()] {
			queue = append(queue, cur.parent)
		}
	}

	// Phase 3: traverse tree from root to find largest subtrees that need to
	// be restarted and are ready to be restarted.

	// All DNs that need to be restarted by the GC process.
	want := make(map[string]bool)
	// All DNs that need to be restarted and can be restarted by the GC process
	// - a subset of 'want' DNs.
	can := make(map[string]bool)
	// The set difference between 'want' and 'can' are all nodes that should be
	// restarted but can't yet (ie. because a child is still in the process of
	// being canceled).

	// DFS from root.
	queue = []*node{s.root}
	for {
		if len(queue) == 0 {
			break
		}

		cur := queue[0]
		queue = queue[1:]

		// If this node is DEAD or CANCELED it should be restarted.
		if cur.state == NodeStateDead || cur.state == NodeStateCanceled {
			want[cur.dn()] = true
		}

		// If it should be restarted and is ready to be restarted...
		if want[cur.dn()] && ready[cur.dn()] {
			// And its parent context is valid (ie hasn't been canceled), mark
			// it as restartable.
			if cur.parent == nil || cur.parent.ctx.Err() == nil {
				can[cur.dn()] = true
				continue
			}
		}

		// Otherwise, traverse further down the tree to see if something else
		// needs to be done.
		for _, c := range cur.children {
			queue = append(queue, c)
		}
	}

	// Reinitialize and reschedule all subtrees
	for dn := range can {
		n := s.nodeByDN(dn)

		// Only back off when the node unexpectedly died - not when it got
		// canceled.
		bo := time.Duration(0)
		if n.state == NodeStateDead {
			bo = n.bo.NextBackOff()
		}

		// Prepare node for rescheduling - remove its children, reset its state
		// to new.
		n.reset()
		s.ilogger.Infof("rescheduling supervised node %s with backoff %s", dn, bo.String())

		// Reschedule node runnable to run after backoff.
		go func(n *node, bo time.Duration) {
			time.Sleep(bo)
			s.pReq <- &processorRequest{
				schedule: &processorRequestSchedule{dn: n.dn()},
			}
		}(n, bo)
	}
}
