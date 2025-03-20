// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package structfs_test

import (
	"errors"
	"io"
	"io/fs"
	"slices"
	"syscall"
	"testing"
	"time"

	. "source.monogon.dev/osbase/structfs"
)

func TestOptions(t *testing.T) {
	testTimestamp := time.Date(2022, 03, 04, 5, 6, 8, 0, time.UTC)
	var tree Tree
	tree.PlaceDir("dir", Tree{},
		WithModTime(testTimestamp),
		WithPerm(0o700|fs.ModeSetuid|fs.ModeDevice),
		WithSys("fakesys"),
	)
	node := tree[0]
	if node.ModTime != testTimestamp {
		t.Errorf("Got ModTime %v, expected %v", node.ModTime, testTimestamp)
	}
	expectMode := 0o700 | fs.ModeSetuid | fs.ModeDir
	if node.Mode != expectMode {
		t.Errorf("Got Mode %s, expected %s", node.Mode, expectMode)
	}
	if node.Sys != "fakesys" {
		t.Errorf("Got Sys %v, expected %v", node.Sys, "fakesys")
	}
}

func treeToStrings(t *testing.T, tree Tree) []string {
	var out []string
	for path, node := range tree.Walk() {
		s := path + " " + node.Mode.String()[:1]
		if node.Mode.IsRegular() {
			content, err := node.Content.Open()
			if err != nil {
				t.Errorf("Failed to open %q: %v", path, err)
				continue
			}
			b, err := io.ReadAll(content)
			if err != nil {
				t.Errorf("Failed to read %q: %v", path, err)
				continue
			}
			s += " " + string(b)
			content.Close()
		}
		out = append(out, s)
	}
	return out
}

func TestWalk(t *testing.T) {
	testCases := []struct {
		desc     string
		tree     Tree
		expected []string
	}{
		{
			desc: "example",
			tree: Tree{
				File("file1a", Bytes("content1a")),
				Dir("dir1", Tree{
					File("file2", Bytes("content2")),
					Dir("dir2", nil),
				}),
				File("file1b", Bytes("content1b")),
			},
			expected: []string{
				"file1a - content1a",
				"dir1 d",
				"dir1/file2 - content2",
				"dir1/dir2 d",
				"file1b - content1b",
			},
		},
		{
			desc:     "empty",
			tree:     nil,
			expected: nil,
		},
		{
			desc: "ignore file children",
			// Non-directories should not have children and Walk should ignore them.
			tree: Tree{{
				Name:    "file1",
				Content: Bytes("content1"),
				Children: Tree{
					File("file2", Bytes("content2")),
					Dir("dir2", nil),
				},
			}},
			expected: []string{
				"file1 - content1",
			},
		},
		{
			desc: "skip invalid name",
			tree: Tree{
				File("", Bytes("invalid")),
				File(".", Bytes("invalid")),
				File("a/b", Bytes("invalid")),
				File("file1a", Bytes("content1a")),
				Dir("..", Tree{
					File("file2", Bytes("content2")),
					Dir("dir2", nil),
				}),
				File("file1b", Bytes("content1b")),
			},
			expected: []string{
				"file1a - content1a",
				"file1b - content1b",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := treeToStrings(t, tC.tree)
			if !slices.Equal(actual, tC.expected) {
				t.Errorf("Walk result %v differs from expected %v", actual, tC.expected)
			}
		})
	}
}

func TestPlace(t *testing.T) {
	testCases := []struct {
		desc     string
		tree     Tree
		place    func(*Tree) error
		expected []string
		err      error
		errPath  string
	}{
		{
			desc: "file",
			tree: Tree{
				File("file1a", Bytes("content1a")),
				Dir("dir1", Tree{
					File("file2", Bytes("content2")),
					Dir("dir2", nil),
				}),
				File("file1b", Bytes("content1b")),
			},
			place: func(tree *Tree) error {
				return tree.PlaceFile("dir1/dir3/file4", Bytes("content4"))
			},
			expected: []string{
				"file1a - content1a",
				"dir1 d",
				"dir1/file2 - content2",
				"dir1/dir2 d",
				"dir1/dir3 d",
				"dir1/dir3/file4 - content4",
				"file1b - content1b",
			},
		},
		{
			desc: "dir",
			tree: Tree{
				File("file1a", Bytes("content1a")),
				Dir("dir1", Tree{
					File("file2", Bytes("content2")),
				}),
			},
			place: func(tree *Tree) error {
				return tree.PlaceDir("dir1/dir3", Tree{
					File("file4", Bytes("content4")),
					Dir("dir4", nil),
				})
			},
			expected: []string{
				"file1a - content1a",
				"dir1 d",
				"dir1/file2 - content2",
				"dir1/dir3 d",
				"dir1/dir3/file4 - content4",
				"dir1/dir3/dir4 d",
			},
		},
		{
			desc: "empty",
			tree: nil,
			place: func(tree *Tree) error {
				return tree.PlaceFile("dir1/dir2/file3", Bytes("content"))
			},
			expected: []string{
				"dir1 d",
				"dir1/dir2 d",
				"dir1/dir2/file3 - content",
			},
		},
		{
			desc: "root",
			tree: Tree{
				File("file1", Bytes("content1")),
			},
			place: func(tree *Tree) error {
				return tree.PlaceFile("file2", Bytes("content2"))
			},
			expected: []string{
				"file1 - content1",
				"file2 - content2",
			},
		},
		{
			desc: "invalid path",
			place: func(tree *Tree) error {
				return tree.PlaceFile(".", Bytes("content"))
			},
			err:     fs.ErrInvalid,
			errPath: ".",
		},
		{
			desc: "not a directory",
			tree: Tree{
				Dir("dir1", Tree{
					File("file2", Bytes("content")),
				}),
			},
			place: func(tree *Tree) error {
				return tree.PlaceFile("dir1/file2/dir3/file4", Bytes("content"))
			},
			err:     syscall.ENOTDIR,
			errPath: "dir1/file2",
		},
		{
			desc: "already exists",
			tree: Tree{
				Dir("dir1", Tree{
					Dir("dir2", nil),
				}),
			},
			place: func(tree *Tree) error {
				return tree.PlaceDir("dir1/dir2", nil)
			},
			err:     fs.ErrExist,
			errPath: "dir1/dir2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.place(&tC.tree)
			if err != nil {
				if tC.err == nil {
					t.Fatalf("Place failed unexpectedly: %v", err)
				}
				if !errors.Is(err, tC.err) {
					t.Errorf("Place failed with error %v, expected %v", err, tC.err)
				}
				var pe *fs.PathError
				if !errors.As(err, &pe) {
					t.Fatalf("Place(): error is %T, want *fs.PathError", err)
				}
				if pe.Path != tC.errPath {
					t.Errorf("Place(): err.Path = %q, want %q", pe.Path, tC.errPath)
				}
			} else if tC.err != nil {
				t.Error("Expected place to fail but it did not")
			} else {
				actual := treeToStrings(t, tC.tree)
				if !slices.Equal(actual, tC.expected) {
					t.Errorf("Result %v differs from expected %v", actual, tC.expected)
				}
			}
		})
	}
}

func TestValidName(t *testing.T) {
	isValidNameTests := []struct {
		name string
		ok   bool
	}{
		{"x", true},
		{"", false},
		{"..", false},
		{".", false},
		{"x/y", false},
		{"/", false},
		{"x/", false},
		{"/x", false},
		{`x\y`, true},
	}
	for _, tt := range isValidNameTests {
		ok := ValidName(tt.name)
		if ok != tt.ok {
			t.Errorf("ValidName(%q) = %v, want %v", tt.name, ok, tt.ok)
		}
	}
}
