package main

import (
	"context"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/roleserve"
	"source.monogon.dev/osbase/logtree"
)

// runDebugService runs the debug service if this is a debug build. Otherwise
// it does nothing.
func runDebugService(_ context.Context, _ *roleserve.Service, _ *logtree.LogTree, _ *localstorage.Root) error {
	// This code is included in the production build, do nothing.
	return nil
}
