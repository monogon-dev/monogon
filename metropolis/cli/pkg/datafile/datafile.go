// datafile provides an abstraction for accessing files passed through the data
// attribute in a Bazel build rule.
//
// It thinly wraps around the Bazel/Go runfile library (to allow running from
// outside `bazel run`).
package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

// parseManifest takes a bazel runfile MANIFEST and parses it into a map from
// workspace-relative path to absolute path, flattening all workspaces into a
// single tree.
func parseManifest(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open MANIFEST: %v", err)
	}
	defer f.Close()

	manifest := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 2 {
			continue
		}
		fpathParts := strings.Split(parts[0], string(os.PathSeparator))
		fpath := strings.Join(fpathParts[1:], string(os.PathSeparator))
		manifest[fpath] = parts[1]
	}
	return manifest, nil
}

// ResolveRunfile tries to resolve a workspace-relative file path into an
// absolute path with the use of bazel runfiles, through either the original
// Bazel/Go runfile integration or a wrapper that also supports running from
// outside `bazel run`.
func ResolveRunfile(path string) (string, error) {
	var errEx error
	ep, err := os.Executable()
	if err == nil {
		rfdir := ep + ".runfiles"
		mfpath := filepath.Join(rfdir, "MANIFEST")
		if stat, err := os.Stat(rfdir); err == nil && stat.IsDir() {
			// We have a runfiles directory, parse MANIFEST and resolve files this way.
			manifest, err := parseManifest(mfpath)
			if err == nil {
				tpath := manifest[path]
				if tpath == "" {
					errEx = fmt.Errorf("not in MANIFEST")
				} else {
					return tpath, err
				}
			} else {
				errEx = err
			}
		} else {
			errEx = err
		}
	}

	// Try runfiles just in case.
	rf, errRF := bazel.Runfile(path)
	if errRF == nil {
		return rf, nil
	}
	return "", fmt.Errorf("could not resolve via executable location (%v) and runfile resolution failed: %v", errEx, errRF)
}

// Get tries to read a workspace-relative file path through the use of Bazel
// runfiles, including for cases when executables are running outside `bazel
// run`.
func Get(path string) ([]byte, error) {
	rfpath, err := ResolveRunfile(path)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(rfpath)
}

// MustGet either successfully resolves a file through Get() or logs an error
// (through the stdlib log library) and stops execution. This should thus only
// be used in binaries which use the log library.
func MustGet(path string) []byte {
	res, err := Get(path)
	if err != nil {
		log.Fatalf("Could not get datafile %s: %v", path, err)
	}
	return res
}
