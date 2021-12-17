package importsort

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"source.monogon.dev/build/toolbase/gotoolchain"
)

//go:embed testdata/*
var testdata embed.FS

func init() {
	// analysistest uses x/go/packages which in turns uses runtime.GOROOT().
	// runtime.GOROOT itself is neutered by rules_go for hermeticity. We provide our
	// own GOROOT that we get from gotoolchain (which is still hermetic, but depends
	// on runfiles). However, for runtime.GOROOT to pick it up, this env var must be
	// set in init().
	os.Setenv("GOROOT", gotoolchain.Root)
}

func TestImportsort(t *testing.T) {
	// Add `go` to PATH for x/go/packages.
	os.Setenv("PATH", filepath.Dir(gotoolchain.Go))

	// Make an empty GOCACHE for x/go/packages.
	gocache, err := os.MkdirTemp("/tmp", "gocache")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(gocache)
	os.Setenv("GOCACHE", gocache)

	// Make an empty GOPATH for x/go/packages.
	gopath, err := os.MkdirTemp("/tmp", "gopath")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(gopath)
	os.Setenv("GOPATH", gopath)

	// Convert testdata from an fs.FS to a path->contents map as expected by
	// analysistest.WriteFiles, rewriting paths to build a correct GOPATH-like
	// layout.
	filemap := make(map[string]string)
	fs.WalkDir(testdata, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			t.Fatalf("WalkDir: %v", err)
		}
		if d.IsDir() {
			return nil
		}
		bytes, _ := testdata.ReadFile(path)
		path = strings.TrimPrefix(path, "testdata/")
		path = strings.ReplaceAll(path, ".notgo", ".go")
		filemap[path] = string(bytes)
		return nil
	})

	// Run the actual tests, which are all declared within testdata/**.
	dir, cleanup, err := analysistest.WriteFiles(filemap)
	if err != nil {
		t.Fatalf("WriteFiles: %v", err)
	}
	defer cleanup()
	analysistest.Run(t, dir, Analyzer, "source.monogon.dev/...")
}
