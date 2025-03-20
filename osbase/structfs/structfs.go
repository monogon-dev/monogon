// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package structfs defines a data structure for a file system, similar to the
// [fs] package but based on structs instead of interfaces.
//
// The entire tree structure and directory entry metadata is stored in memory.
// File content is represented with [Blob], and may come from various sources.
package structfs

import (
	"io/fs"
	"iter"
	pathlib "path"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

// Tree represents a file system tree.
type Tree []*Node

// Node is a node in a file system tree, which is either a file or directory.
type Node struct {
	// Name of this node, which must be valid according to [ValidName].
	Name string
	// Mode contains the file type and permissions.
	Mode fs.FileMode
	// ModTime is the modification time.
	ModTime time.Time
	// Content is the file content, must be set for regular files.
	Content Blob
	// Children of a directory, must be empty if this is not a directory.
	Children Tree
	// Sys contains any system-specific directory entry fields.
	//
	// It should be accessed using interface type assertions, to allow combining
	// information for multiple target systems with struct embedding.
	Sys any
}

type Option func(*Node)

// WithModTime sets the ModTime of the Node.
func WithModTime(t time.Time) Option {
	return func(n *Node) {
		n.ModTime = t
	}
}

const permMask = fs.ModePerm | fs.ModeSetuid | fs.ModeSetgid | fs.ModeSticky

// WithPerm sets the permission bits of the Node.
func WithPerm(perm fs.FileMode) Option {
	return func(n *Node) {
		n.Mode = (n.Mode & ^permMask) | (perm & permMask)
	}
}

// WithSys sets the Sys field of the Node.
func WithSys(sys any) Option {
	return func(n *Node) {
		n.Sys = sys
	}
}

// File creates a regular file node with the given name and content.
//
// Permission defaults to 644.
func File(name string, content Blob, opts ...Option) *Node {
	n := &Node{
		Name:    name,
		Mode:    0o644,
		Content: content,
	}
	for _, f := range opts {
		f(n)
	}
	return n
}

// Dir creates a directory node with the given name and children.
//
// Permission defaults to 755.
func Dir(name string, children Tree, opts ...Option) *Node {
	n := &Node{
		Name:     name,
		Mode:     fs.ModeDir | 0o755,
		Children: children,
	}
	for _, f := range opts {
		f(n)
	}
	return n
}

// PlaceFile creates parent directories if necessary and places a file with the
// given content at the path. It fails if path already exists.
func (t *Tree) PlaceFile(path string, content Blob, opts ...Option) error {
	path, name, err := splitPlacePath(path)
	if err != nil {
		return err
	}
	return t.Place(path, File(name, content, opts...))
}

// PlaceDir creates parent directories if necessary and places a directory with
// the given children at the path. It fails if path already exists.
func (t *Tree) PlaceDir(path string, children Tree, opts ...Option) error {
	path, name, err := splitPlacePath(path)
	if err != nil {
		return err
	}
	return t.Place(path, Dir(name, children, opts...))
}

func splitPlacePath(path string) (dir string, name string, err error) {
	if !fs.ValidPath(path) || path == "." {
		return "", "", &fs.PathError{Op: "place", Path: path, Err: fs.ErrInvalid}
	}
	dir, name = pathlib.Split(path)
	if dir == "" {
		dir = "."
	} else {
		dir = dir[:len(dir)-1]
	}
	return
}

// Place creates directories if necessary and places the node in the directory
// at the path.
//
// The special path "." indicates the root.
func (t *Tree) Place(path string, node *Node) error {
	if !fs.ValidPath(path) {
		return &fs.PathError{Op: "place", Path: path, Err: fs.ErrInvalid}
	}
	treeRef := t
	if path != "." {
		pathlen := 0
	outer:
		for name := range strings.SplitSeq(path, "/") {
			pathlen += len(name) + 1
			for _, nodeRef := range *treeRef {
				if nodeRef.Name == name {
					if !nodeRef.Mode.IsDir() {
						return &fs.PathError{Op: "mkdir", Path: path[:pathlen-1], Err: syscall.ENOTDIR}
					}
					treeRef = &nodeRef.Children
					continue outer
				}
			}
			dir := Dir(name, nil)
			*treeRef = append(*treeRef, dir)
			treeRef = &dir.Children
		}
	}
	for _, nodeRef := range *treeRef {
		if nodeRef.Name == node.Name {
			return &fs.PathError{Op: "place", Path: path + "/" + nodeRef.Name, Err: fs.ErrExist}
		}
	}
	*treeRef = append(*treeRef, node)
	return nil
}

// Walk returns an iterator over all nodes in the tree in DFS pre-order.
// The key is the path of the node.
//
// Entries with invalid name are skipped.
func (t Tree) Walk() iter.Seq2[string, *Node] {
	return func(yield func(string, *Node) bool) {
		walk(t, ".", yield)
	}
}

func walk(t Tree, path string, yield func(string, *Node) bool) bool {
	for _, node := range t {
		if !ValidName(node.Name) {
			// Skip entries with invalid name.
			continue
		}
		nodePath := node.Name
		if path != "." {
			nodePath = path + "/" + nodePath
		}
		if !yield(nodePath, node) {
			return false
		}
		if node.Mode.IsDir() {
			if !walk(node.Children, nodePath, yield) {
				return false
			}
		}
	}
	return true
}

// ValidName reports whether the given name is a valid node name.
//
// The name must be UTF-8-encoded, must not be empty, "." or "..", and must not
// contain "/". These are the same rules as for a path element in [fs.ValidPath].
func ValidName(name string) bool {
	if !utf8.ValidString(name) {
		return false
	}
	if name == "" || name == "." || name == ".." {
		return false
	}
	if strings.ContainsRune(name, '/') {
		return false
	}
	return true
}
