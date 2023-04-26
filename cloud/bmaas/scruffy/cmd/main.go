package main

import (
	"context"
	"flag"

	"source.monogon.dev/cloud/bmaas/scruffy"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

func main() {
	s := &scruffy.Server{}
	s.Config.RegisterFlags()
	flag.Parse()

	ctx := clicontext.WithInterrupt(context.Background())
	s.Start(ctx)
	<-ctx.Done()
}
