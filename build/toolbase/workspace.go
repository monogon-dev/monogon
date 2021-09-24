package toolbase

import (
	"fmt"
	"os"
	"path"
)

// isWorkspace returns whether a given string is a valid path pointing to a
// Bazel workspace directory.
func isWorkspace(dir string) bool {
	w := path.Join(dir, "WORKSPACE")
	if _, err := os.Stat(w); err == nil {
		return true
	}
	return false
}

// WorkspaceDirectory returns the workspace directory from which a given
// command line tool is running. This handles the following cases:
//
// 1. The command line tool was invoked via `bazel run`.
// 2. The command line tool was started directly in a workspace directory (but
//    not a subdirectory).
//
// If the workspace directory path cannot be inferred based on the above
// assumptions, an error is returned.
func WorkspaceDirectory() (string, error) {
	if p := os.Getenv("BUILD_WORKSPACE_DIRECTORY"); p != "" && isWorkspace(p) {
		return p, nil
	}

	if p, err := os.Getwd(); err != nil && isWorkspace(p) {
		return p, nil
	}

	return "", fmt.Errorf("not invoked from `bazel run` and not running in workspace directory")
}
