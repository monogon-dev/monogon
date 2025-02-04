// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/apigw/server"
)

func main() {
	s := &server.Server{}
	s.Config.RegisterFlags()
	flag.Parse()
	if flag.NArg() > 0 {
		klog.Exitf("unexpected positional arguments: %v", flag.Args())
	}

	ctx, ctxC := context.WithCancel(context.Background())
	// TODO: context cancel on interrupt.
	_ = ctxC

	s.Start(ctx)
	select {}
}
