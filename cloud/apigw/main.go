package main

import (
	"context"
	"flag"

	"source.monogon.dev/cloud/apigw/server"
)

func main() {
	s := &server.Server{}
	s.Config.RegisterFlags()
	flag.Parse()

	ctx, ctxC := context.WithCancel(context.Background())
	// TODO: context cancel on interrupt.
	_ = ctxC

	s.Start(ctx)
	select {}
}
