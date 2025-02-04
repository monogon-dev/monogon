// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"

	"source.monogon.dev/osbase/bringup"
)

func main() {
	bringup.Runnable(func(ctx context.Context) error {
		// Disable the exception tracing of the kernel which captures unhandled
		// signals as it races with our logging which makes the test flaky.
		err := os.WriteFile("/proc/sys/debug/exception-trace", []byte("0"), 0755)
		if err != nil {
			return err
		}

		// Provoke a segfault, which produces a panic inside the root runnable
		//nolint:nilness
		_ = *(*string)(nil)
		return nil
	}).Run()
}
