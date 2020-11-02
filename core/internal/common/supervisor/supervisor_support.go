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

// Supporting infrastructure to allow running some non-Go payloads under supervision.

import (
	"context"
	"net"
	"os/exec"

	"google.golang.org/grpc"
)

// GRPCServer creates a Runnable that serves gRPC requests as longs as it's not canceled.
// If graceful is set to true, the server will be gracefully stopped instead of plain stopped. This means all pending
// RPCs will finish, but also requires streaming gRPC handlers to check their context liveliness and exit accordingly.
// If the server code does not support this, `graceful` should be false and the server will be killed violently instead.
func GRPCServer(srv *grpc.Server, lis net.Listener, graceful bool) Runnable {
	return func(ctx context.Context) error {
		Signal(ctx, SignalHealthy)
		errC := make(chan error)
		go func() {
			errC <- srv.Serve(lis)
		}()
		select {
		case <-ctx.Done():
			if graceful {
				srv.GracefulStop()
			} else {
				srv.Stop()
			}
			return ctx.Err()
		case err := <-errC:
			return err
		}
	}
}

// RunCommand will create a Runnable that starts a long-running command, whose exit is determined to be a failure.
func RunCommand(ctx context.Context, cmd *exec.Cmd) error {
	Signal(ctx, SignalHealthy)
	cmd.Stdout = RawLogger(ctx)
	cmd.Stderr = RawLogger(ctx)
	err := cmd.Run()
	Logger(ctx).Infof("Command returned: %v", err)
	return err
}
