package main

import (
	"context"
	"flag"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/server"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

func main() {
	s := &server.Server{}
	s.Config.RegisterFlags()
	flag.Parse()
	if flag.NArg() > 0 {
		klog.Exitf("unexpected positional arguments: %v", flag.Args())
	}

	ctx := clicontext.WithInterrupt(context.Background())
	s.Start(ctx)
	select {}
}
