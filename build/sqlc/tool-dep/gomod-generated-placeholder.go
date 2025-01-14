//go:build tools
// +build tools

// This is a synthetic file which depends on the sqlc binary. That in turn
// makes `go mod tidy` pick it up, which in turns makes it available to Bazel.
package main

import (
	_ "github.com/sqlc-dev/sqlc/pkg/cli"
)

func main() {
}
