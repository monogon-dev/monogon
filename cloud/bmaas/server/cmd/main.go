package main

import (
	"context"
	"flag"

	"source.monogon.dev/cloud/bmaas/server"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

func main() {
	s := &server.Server{}
	s.Config.RegisterFlags()
	flag.Parse()

	ctx := clicontext.WithInterrupt(context.Background())
	s.Start(ctx)
	select {}
}
