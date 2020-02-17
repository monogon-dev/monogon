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

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"

	"git.monogon.dev/source/nexantic.git/core/internal/network"
	"git.monogon.dev/source/nexantic.git/core/internal/node"
	"git.monogon.dev/source/nexantic.git/core/internal/storage"
	"git.monogon.dev/source/nexantic.git/core/pkg/tpm"

	"go.uber.org/zap"
	"golang.org/x/sys/unix"
)

const (
	apiPort       = 7833
	consensusPort = 7834
)

func main() {
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Init panicked:", r)
			debug.PrintStack()
		}
		unix.Sync()
		// TODO(lorenz): Switch this to Reboot when init panics are less likely
		// Best effort, nothing we can do if this fails except printing the error to the console.
		if err := unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF); err != nil {
			panic(fmt.Sprintf("failed to halt node: %v\n", err))
		}
	}()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger.Info("Starting Smalltown Init")

	// Set up bare minimum mounts
	if err := os.Mkdir("/sys", 0755); err != nil {
		panic(err)
	}
	if err := unix.Mount("sysfs", "/sys", "sysfs", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, ""); err != nil {
		panic(err)
	}

	if err := os.Mkdir("/proc", 0755); err != nil {
		panic(err)
	}
	if err := unix.Mount("procfs", "/proc", "proc", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, ""); err != nil {
		panic(err)
	}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel)

	if err := tpm.Initialize(logger.With(zap.String("component", "tpm"))); err != nil {
		logger.Panic("Failed to initialize TPM 2.0", zap.Error(err))
	}

	storageManager, err := storage.Initialize(logger.With(zap.String("component", "storage")))
	if err != nil {
		panic(err)
	}

	networkSvc, err := network.NewNetworkService(network.Config{}, logger.With(zap.String("component", "network")))
	if err != nil {
		panic(err)
	}

	if err := networkSvc.Start(); err != nil {
		logger.Panic("Failed to start network service", zap.Error(err))
	}

	nodeInstance, err := node.NewSmalltownNode(logger, networkSvc, storageManager)
	if err != nil {
		panic(err)
	}

	err = nodeInstance.Start(ctx)
	if err != nil {
		panic(err)
	}

	// We're PID1, so orphaned processes get reparented to us to clean up
	for {
		sig := <-signalChannel
		switch sig {
		case unix.SIGCHLD:
			var status unix.WaitStatus
			var rusage unix.Rusage
			for {
				res, err := unix.Wait4(-1, &status, unix.WNOHANG, &rusage)
				if err != nil && err != unix.ECHILD {
					logger.Error("Failed to wait on orphaned child", zap.Error(err))
					break
				}
				if res <= 0 {
					break
				}
			}
		// TODO(lorenz): We can probably get more than just SIGCHLD as init, but I can't think
		// of any others right now, just log them in case we hit any of them.
		default:
			logger.Warn("Got unexpected signal", zap.String("signal", sig.String()))
		}
	}
}
