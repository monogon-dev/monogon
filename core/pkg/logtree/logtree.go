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

package logtree

import (
	"fmt"
	"strings"
	"sync"
)

// LogTree is a tree-shapped logging system. For more information, see the package-level documentation.
type LogTree struct {
	// journal is the tree's journal, storing all log data and managing subscribers.
	journal *journal
	// root is the root node of the actual tree of the log tree. The nodes contain per-DN configuration options, notably
	// the current verbosity level of that DN.
	root *node
}

func New() *LogTree {
	lt := &LogTree{
		journal: newJournal(),
	}
	lt.root = newNode(lt, "")
	return lt
}

// node represents a given DN as a discrete 'logger'. It implementes the LeveledLogger interface for log publishing,
// entries from which it passes over to the logtree's journal.
type node struct {
	// dn is the DN which this node represents (or "" if this is the root node).
	dn DN
	// tree is the LogTree to which this node belongs.
	tree *LogTree
	// verbosity is the current verbosity level of this DN/node, affecting .V(n) LeveledLogger calls
	verbosity VerbosityLevel

	// mu guards children.
	mu sync.Mutex
	// children is a map of DN-part to a children node in the logtree. A DN-part is a string representing a part of the
	// DN between the deliming dots, as returned by DN.Path.
	children map[string]*node
}

// newNode returns a node at a given DN in the LogTree - but doesn't set up the LogTree to insert it accordingly.
func newNode(tree *LogTree, dn DN) *node {
	n := &node{
		dn:       dn,
		tree:     tree,
		children: make(map[string]*node),
	}
	return n
}

// nodeByDN returns the LogTree node corresponding to a given DN. If either the node or some of its parents do not
// exist they will be created as needed.
func (l *LogTree) nodeByDN(dn DN) (*node, error) {
	traversal, err := newTraversal(dn)
	if err != nil {
		return nil, fmt.Errorf("traversal failed: %w", err)
	}
	return traversal.execute(l.root), nil
}

// nodeTraversal represents a request to traverse the LogTree in search of a given node by DN.
type nodeTraversal struct {
	// want is the DN of the node's that requested to be found.
	want DN
	// current is the path already taken to find the node, in the form of DN parts. It starts out as want.Parts() and
	// progresses to become empty as the traversal continues.
	current []string
	// left is the path that's still needed to be taken in order to find the node, in the form of DN parts. It starts
	// out empty and progresses to become wants.Parts() as the traversal continues.
	left []string
}

// next adjusts the traversal's current/left slices to the next element of the traversal, returns the part that's now
// being looked for (or "" if the traveral is done) and the full DN of the element that's being looked for.
//
// For example, a traversal of foo.bar.baz will cause .next() to return the following on each invocation:
//  - part: foo, full: foo
//  - part: bar, full: foo.bar
//  - part: baz, full: foo.bar.baz
//  - part: "",  full: foo.bar.baz
func (t *nodeTraversal) next() (part string, full DN) {
	if len(t.left) == 0 {
		return "", t.want
	}
	part = t.left[0]
	t.current = append(t.current, part)
	t.left = t.left[1:]
	full = DN(strings.Join(t.current, "."))
	return
}

// newTraversal returns a nodeTraversal fora a given wanted DN.
func newTraversal(dn DN) (*nodeTraversal, error) {
	parts, err := dn.Path()
	if err != nil {
		return nil, err
	}
	return &nodeTraversal{
		want: dn,
		left: parts,
	}, nil
}

// execute the traversal in order to find the node. This can only be called once per traversal.
// Nodes will be created within the tree until the target node is reached. Existing nodes will be reused.
// This is effectively an idempotent way of accessing a node in the tree based on a traversal.
func (t *nodeTraversal) execute(n *node) *node {
	cur := n
	for {
		part, full := t.next()
		if part == "" {
			return cur
		}

		mu := &cur.mu
		mu.Lock()
		if _, ok := cur.children[part]; !ok {
			cur.children[part] = newNode(n.tree, DN(full))
		}
		cur = cur.children[part]
		mu.Unlock()
	}
}
