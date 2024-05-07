package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"source.monogon.dev/cloud/bmaas/scruffy"
)

func main() {
	s := &scruffy.Server{}
	s.Config.RegisterFlags()
	flag.Parse()

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	s.Start(ctx)
	<-ctx.Done()
}
