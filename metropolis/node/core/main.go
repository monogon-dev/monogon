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
	"io"
	"os"
	"runtime/debug"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node/core/cluster"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/network/hostsfile"
	"source.monogon.dev/metropolis/node/core/roleserve"
	timesvc "source.monogon.dev/metropolis/node/core/time"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"
	"source.monogon.dev/metropolis/pkg/tpm"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "\n\n")
			fmt.Fprintf(os.Stderr, "  Metropolis encountered an uncorrectable error and this node must be restarted.\n")
			fmt.Fprintf(os.Stderr, "  Core panicked: %v\n\n", r)
			debug.PrintStack()
		}
		unix.Sync()
		// TODO(lorenz): Switch this to Reboot when init panics are less likely.
		if err := unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF); err != nil {
			// Best effort, nothing we can do if this fails except printing the error to the
			// console.
			panic(fmt.Sprintf("failed to halt node: %v\n", err))
		}
	}()

	// Set up basic mounts (like /dev, /sys...).
	if err := setupMounts(); err != nil {
		panic(fmt.Errorf("could not set up basic mounts: %w", err))
	}

	// Set up logger for Metropolis. Currently logs everything to /dev/tty0 and
	// /dev/ttyS0.
	lt := logtree.New()
	for _, p := range []string{
		"/dev/tty0", "/dev/ttyS0",
	} {
		f, err := os.OpenFile(p, os.O_WRONLY, 0)
		if err != nil {
			continue
		}
		reader, err := lt.Read("", logtree.WithChildren(), logtree.WithStream())
		if err != nil {
			panic(fmt.Errorf("could not set up root log reader: %v", err))
		}
		go func(path string, f io.Writer) {
			fmt.Fprintf(f, "\nMetropolis: this is %s. Verbose node logs follow.\n\n", path)
			for {
				p := <-reader.Stream
				fmt.Fprintf(f, "%s\n", p.String())
			}
		}(p, f)
	}

	// Initial logger. Used until we get to a supervisor.
	logger := lt.MustLeveledFor("init")

	// Linux kernel default is 4096 which is far too low. Raise it to 1M which
	// is what gVisor suggests.
	if err := unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{Cur: 1048576, Max: 1048576}); err != nil {
		logger.Fatalf("Failed to raise rlimits: %v", err)
	}

	logger.Info("Starting Metropolis node init")

	if err := tpm.Initialize(logger); err != nil {
		logger.Warningf("Failed to initialize TPM 2.0, attempting fallback to untrusted: %v", err)
	}

	networkSvc := network.New()
	timeSvc := timesvc.New()

	// This function initializes a headless Delve if this is a debug build or
	// does nothing if it's not
	initializeDebugger(networkSvc)

	// Prepare local storage.
	root := &localstorage.Root{}
	if err := declarative.PlaceFS(root, "/"); err != nil {
		panic(fmt.Errorf("when placing root FS: %w", err))
	}

	// trapdoor is a channel used to signal to the init service that a very
	// low-level, unrecoverable failure occured.
	trapdoor := make(chan struct{})

	// Make context for supervisor. We cancel it when we reach the trapdoor.
	ctxS, ctxC := context.WithCancel(context.Background())

	// Start root initialization code as a supervisor one-shot runnable. This
	// means waiting for the network, starting the cluster manager, and then
	// starting all services related to the node's roles.
	// TODO(q3k): move this to a separate 'init' service.
	supervisor.New(ctxS, func(ctx context.Context) error {
		// Start storage and network - we need this to get anything else done.
		if err := root.Start(ctx); err != nil {
			return fmt.Errorf("cannot start root FS: %w", err)
		}
		if err := supervisor.Run(ctx, "network", networkSvc.Run); err != nil {
			return fmt.Errorf("when starting network: %w", err)
		}
		if err := supervisor.Run(ctx, "time", timeSvc.Run); err != nil {
			return fmt.Errorf("when starting time: %w", err)
		}
		if err := supervisor.Run(ctx, "pstore", dumpAndCleanPstore); err != nil {
			return fmt.Errorf("when starting pstore: %w", err)
		}

		// Start the role service. The role service connects to the curator and runs
		// all node-specific role code (eg. Kubernetes services).
		//   supervisor.Logger(ctx).Infof("Starting role service...")
		rs := roleserve.New(roleserve.Config{
			StorageRoot: root,
			Network:     networkSvc,
		})
		if err := supervisor.Run(ctx, "role", rs.Run); err != nil {
			close(trapdoor)
			return fmt.Errorf("failed to start role service: %w", err)
		}

		// Start the hostsfile service.
		hostsfileSvc := hostsfile.Service{
			Config: hostsfile.Config{
				Roleserver: rs,
				Network:    networkSvc,
				Ephemeral:  &root.Ephemeral,
				ESP:        &root.ESP,
			},
		}
		if err := supervisor.Run(ctx, "hostsfile", hostsfileSvc.Run); err != nil {
			close(trapdoor)
			return fmt.Errorf("failed to start hostsfile service: %w", err)
		}

		// Start cluster manager. This kicks off cluster membership machinery,
		// which will either start a new cluster, enroll into one or join one.
		m := cluster.NewManager(root, networkSvc, rs)
		if err := supervisor.Run(ctx, "enrolment", m.Run); err != nil {
			close(trapdoor)
			return fmt.Errorf("when starting enrolment: %w", err)
		}

		if err := runDebugService(ctx, rs, lt, root); err != nil {
			return fmt.Errorf("when starting debug service: %w", err)
		}

		supervisor.Signal(ctx, supervisor.SignalHealthy)
		supervisor.Signal(ctx, supervisor.SignalDone)
		return nil
	}, supervisor.WithExistingLogtree(lt))

	<-trapdoor
	logger.Infof("Trapdoor closed, exiting core.")
	ctxC()
}
