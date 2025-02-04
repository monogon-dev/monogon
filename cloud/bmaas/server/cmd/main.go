// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/server"
)

func main() {
	s := &server.Server{}
	s.Config.RegisterFlags()
	flag.Parse()
	if flag.NArg() > 0 {
		klog.Exitf("unexpected positional arguments: %v", flag.Args())
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	s.Start(ctx)
	select {}
}
