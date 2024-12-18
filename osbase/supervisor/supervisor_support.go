// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package supervisor

// Supporting infrastructure to allow running some non-Go payloads under
// supervision.

import (
	"context"
	"errors"
	"net"
	"os"
	"os/exec"

	"google.golang.org/grpc"

	"source.monogon.dev/osbase/logtree"
)

// GRPCServer creates a Runnable that serves gRPC requests as longs as it's not
// canceled.
// If graceful is set to true, the server will be gracefully stopped instead of
// plain stopped. This means all pending RPCs will finish, but also requires
// streaming gRPC handlers to check their context liveliness and exit
// accordingly.  If the server code does not support this, `graceful` should be
// false and the server will be killed violently instead.
func GRPCServer(srv *grpc.Server, lis net.Listener, graceful bool) Runnable {
	return func(ctx context.Context) error {
		Signal(ctx, SignalHealthy)
		defer func() {
			if graceful {
				srv.GracefulStop()
			} else {
				srv.Stop()
			}
		}()
		errC := make(chan error)
		go func() {
			errC <- srv.Serve(lis)
		}()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errC:
			return err
		}
	}
}

// RunCommand will create a Runnable that starts a long-running command, whose
// exit is determined to be a failure.
// cmd should be created with [exec.CommandContext] so that it will be killed
// when the context is canceled.
func RunCommand(ctx context.Context, cmd *exec.Cmd, opts ...RunCommandOption) error {
	Signal(ctx, SignalHealthy)

	var parseKLog bool
	var signal <-chan os.Signal
	for _, opt := range opts {
		if opt.parseKlog {
			parseKLog = true
		}
		if opt.signal != nil {
			signal = opt.signal
		}
	}

	if parseKLog {
		// We make two klogs, one for each of stdout/stderr. This is to prevent
		// accidental interleaving of both.
		klogStdout := logtree.KLogParser(Logger(ctx))
		defer klogStdout.Close()
		klogStderr := logtree.KLogParser(Logger(ctx))
		defer klogStderr.Close()

		cmd.Stdout = klogStdout
		cmd.Stderr = klogStderr
	} else {
		cmd.Stdout = RawLogger(ctx)
		cmd.Stderr = RawLogger(ctx)
	}
	err := cmd.Start()
	if err != nil {
		return err
	}

	exited := make(chan struct{})
	if signal != nil {
		go func() {
			for {
				var err error
				select {
				case s := <-signal:
					err = cmd.Process.Signal(s)
				case <-exited:
					return
				}
				if err != nil && !errors.Is(err, os.ErrProcessDone) {
					Logger(ctx).Warningf("Failed sending signal to process: %v", err)
				}
			}
		}()
	}

	err = cmd.Wait()
	if signal != nil {
		exited <- struct{}{}
	}
	Logger(ctx).Infof("Command returned: %v", err)
	return err
}

type RunCommandOption struct {
	parseKlog bool
	signal    <-chan os.Signal
}

// ParseKLog signals that the command being run will return klog-compatible
// logs to stdout and/or stderr, and these will be re-interpreted as structured
// logging and emitted to the supervisor's logger.
func ParseKLog() RunCommandOption {
	return RunCommandOption{
		parseKlog: true,
	}
}

// SignalChan takes a channel which can be used to send signals to the
// supervised process.
//
// The given channel will be read from as long as the underlying process is
// running. If the process doesn't start successfully the channel will not be
// read. When the process exits, the channel will stop being read.
//
// With the above in mind, and also taking into account the inherent lack of
// reliability in delivering any process-handled signals in POSIX/Linux, it is
// recommended to use unbuffered channels, always write to them in a non-blocking
// fashion (eg. in a select { ... default: } block), and to not rely only on the
// signal delivery mechanism for the intended behaviour.
//
// For example, if the signals are used to trigger some configuration reload,
// these configuration reloads should either be verified and signal delivery should
// be retried until confirmed successful, or there should be a backup periodic
// reload performed by the target process independently of signal-based reload
// triggers.
//
// Another example: if the signal delivered is a SIGTERM used to gracefully
// terminate some process, it should be attempted to be delivered a number of
// times before finally SIGKILLing the process.
func SignalChan(s <-chan os.Signal) RunCommandOption {
	return RunCommandOption{
		signal: s,
	}
}
