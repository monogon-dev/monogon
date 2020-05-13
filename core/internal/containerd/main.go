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

package containerd

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"

	"git.monogon.dev/source/nexantic.git/core/pkg/logbuffer"

	"golang.org/x/sys/unix"
)

const runscLogsFIFOPath = "/containerd/run/runsc-logs.fifo"

type Service struct {
	Log      *logbuffer.LogBuffer
	RunscLog *logbuffer.LogBuffer
}

func New() (*Service, error) {
	return &Service{Log: logbuffer.New(5000, 16384), RunscLog: logbuffer.New(5000, 16384)}, nil
}

func (s *Service) Run() supervisor.Runnable {
	return func(ctx context.Context) error {
		cmd := exec.CommandContext(ctx, "/containerd/bin/containerd", "--config", "/containerd/conf/config.toml")
		cmd.Stdout = s.Log
		cmd.Stderr = s.Log
		cmd.Env = []string{"PATH=/containerd/bin", "TMPDIR=/containerd/run/tmp"}

		if err := unix.Mount("tmpfs", "/containerd/run", "tmpfs", 0, ""); err != nil {
			panic(err)
		}
		if err := os.MkdirAll("/containerd/run/tmp", 0755); err != nil {
			panic(err)
		}

		runscFifo, err := os.OpenFile(runscLogsFIFOPath, os.O_CREATE|os.O_RDONLY, os.ModeNamedPipe|0777)
		if err != nil {
			return err
		}
		go func() {
			for {
				n, err := io.Copy(s.RunscLog, runscFifo)
				if n == 0 && err == nil {
					// Hack because pipes/FIFOs can return zero reads when nobody is writing. To avoid busy-looping,
					// sleep a bit before retrying. This does not loose data since the FIFO internal buffer will
					// stall writes when it becomes full. 10ms maximum stall in a non-latency critical process (reading
					// debug logs) is not an issue for us.
					time.Sleep(10 * time.Millisecond)
				} else if err != nil {
					// TODO: Use supervisor.Logger() and Error() before exiting. Should never happen.
					fmt.Println(err)
					return // It's likely that this will busy-loop printing errors if it encounters one, so bail
				}
			}
		}()

		// TODO(lorenz): Healthcheck against CRI RuntimeService.Status() and SignalHealthy

		err = cmd.Run()
		fmt.Fprintf(s.Log, "containerd stopped: %v\n", err)
		return err
	}
}
