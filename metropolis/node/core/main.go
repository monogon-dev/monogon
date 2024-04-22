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
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/cluster"
	"source.monogon.dev/metropolis/node/core/devmgr"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/roleserve"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	timesvc "source.monogon.dev/metropolis/node/core/time"
	"source.monogon.dev/metropolis/node/core/update"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"
	"source.monogon.dev/metropolis/pkg/tpm"
	mversion "source.monogon.dev/metropolis/version"
	"source.monogon.dev/version"
)

func main() {
	// Set up basic mounts (like /dev, /sys...).
	if err := setupMounts(); err != nil {
		panic(fmt.Errorf("could not set up basic mounts: %w", err))
	}

	// Root system logtree.
	lt := logtree.New()

	// Set up logger for Metropolis. Currently logs everything to /dev/tty0 and
	// /dev/ttyS{0,1}.
	consoles := []*console{
		{
			path:     "/dev/tty0",
			maxWidth: 80,
		},
		{
			path:     "/dev/ttyS0",
			maxWidth: 120,
		},
		{
			path:     "/dev/ttyS1",
			maxWidth: 120,
		},
	}
	// Alternative channel that crash handling writes to, and which gets distributed
	// to the consoles.
	crash := make(chan string)

	// Open up consoles and set up logging from logtree and crash channel.
	for _, c := range consoles {
		f, err := os.OpenFile(c.path, os.O_WRONLY, 0)
		if err != nil {
			continue
		}
		reader, err := lt.Read("", logtree.WithChildren(), logtree.WithStream())
		if err != nil {
			panic(fmt.Sprintf("could not set up root log reader: %v", err))
		}
		c.reader = reader
		go func() {
			fmt.Fprintf(f, "\nMetropolis: this is %s. Verbose node logs follow.\n\n", c.path)
			for {
				select {
				case p := <-reader.Stream:
					if consoleFilter(p) {
						fmt.Fprintf(f, "%s\n", p.ConciseString(logtree.MetropolisShortenDict, c.maxWidth))
					}
				case s := <-crash:
					fmt.Fprintf(f, "%s\n", s)
				}
			}
		}()
	}

	// Initialize persistent panic handler early
	initPanicHandler(lt, consoles)

	// Initial logger. Used until we get to a supervisor.
	logger := lt.MustLeveledFor("init")

	// Linux kernel default is 4096 which is far too low. Raise it to 1M which
	// is what gVisor suggests.
	if err := unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{Cur: 1048576, Max: 1048576}); err != nil {
		logger.Fatalf("Failed to raise rlimits: %v", err)
	}

	logger.Info("Starting Metropolis node init")
	logger.Infof("Version: %s", version.Semver(mversion.Version))

	haveTPM := true
	if err := tpm.Initialize(logger); err != nil {
		logger.Warningf("Failed to initialize TPM 2.0: %v", err)
		haveTPM = false
	}

	networkSvc := network.New(nil)
	networkSvc.DHCPVendorClassID = "dev.monogon.metropolis.node.v1"
	networkSvc.ExtraDNSListenerIPs = []net.IP{node.ContainerDNSIP}
	timeSvc := timesvc.New()
	devmgrSvc := devmgr.New()

	// This function initializes a headless Delve if this is a debug build or
	// does nothing if it's not
	initializeDebugger(networkSvc)

	// Prepare local storage.
	root := &localstorage.Root{}
	if err := declarative.PlaceFS(root, "/"); err != nil {
		panic(fmt.Errorf("when placing root FS: %w", err))
	}

	updateSvc := &update.Service{
		Logger: lt.MustLeveledFor("update"),
	}

	// Make context for supervisor. We cancel it when we reach the trapdoor.
	ctxS, ctxC := context.WithCancel(context.Background())

	// Make node-wide cluster resolver.
	res := resolver.New(ctxS, resolver.WithLogger(func(f string, args ...interface{}) {
		lt.MustLeveledFor("resolver").WithAddedStackDepth(1).Infof(f, args...)
	}))

	// Function which performs core, one-way initialization of the node. This means
	// waiting for the network, starting the cluster manager, and then starting all
	// services related to the node's roles.
	init := func(ctx context.Context) error {
		// Start storage and network - we need this to get anything else done.
		if err := root.Start(ctx, updateSvc); err != nil {
			return fmt.Errorf("cannot start root FS: %w", err)
		}
		nodeParams, err := getNodeParams(ctx, root)
		if err != nil {
			return fmt.Errorf("cannot get node parameters: %w", err)
		}
		if nodeParams.NetworkConfig != nil {
			networkSvc.StaticConfig = nodeParams.NetworkConfig
			if err := root.ESP.Metropolis.NetworkConfiguration.Marshal(nodeParams.NetworkConfig); err != nil {
				logger.Errorf("Error writing back network_config from NodeParameters: %v", err)
			}
		}
		if networkSvc.StaticConfig == nil {
			staticConfig, err := root.ESP.Metropolis.NetworkConfiguration.Unmarshal()
			if err == nil {
				networkSvc.StaticConfig = staticConfig
			} else {
				logger.Errorf("Unable to load static config, proceeding without it: %v", err)
			}
		}
		if err := supervisor.Run(ctx, "devmgr", devmgrSvc.Run); err != nil {
			return fmt.Errorf("when starting devmgr: %w", err)
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
		if err := supervisor.Run(ctx, "sysctl", nodeSysctls); err != nil {
			return fmt.Errorf("when applying sysctls: %w", err)
		}

		// The kernel does of course not run in this runnable, only the log pipe
		// runs in it.
		if err := supervisor.Run(ctx, "kernel", func(ctx context.Context) error {
			return logtree.KmsgPipe(ctx, supervisor.Logger(ctx))
		}); err != nil {
			return fmt.Errorf("when starting kernel log pipe: %w", err)
		}

		// Start the role service. The role service connects to the curator and runs
		// all node-specific role code (eg. Kubernetes services).
		//   supervisor.Logger(ctx).Infof("Starting role service...")
		rs := roleserve.New(roleserve.Config{
			StorageRoot: root,
			Network:     networkSvc,
			Resolver:    res,
			LogTree:     lt,
			Update:      updateSvc,
		})
		if err := supervisor.Run(ctx, "role", rs.Run); err != nil {
			return fmt.Errorf("failed to start role service: %w", err)
		}

		if err := runDebugService(ctx, rs, lt, root); err != nil {
			return fmt.Errorf("when starting debug service: %w", err)
		}

		// Start cluster manager. This kicks off cluster membership machinery,
		// which will either start a new cluster, enroll into one or join one.
		m := cluster.NewManager(root, networkSvc, rs, updateSvc, nodeParams, haveTPM)
		return m.Run(ctx)
	}

	// Start the init function in a one-shot runnable. Smuggle out any errors from
	// the init function and stuff them into the fatal channel. This is where the
	// system supervisor takes over as the main process management system.
	fatal := make(chan error)
	supervisor.New(ctxS, func(ctx context.Context) error {
		err := init(ctx)
		if err != nil {
			fatal <- err
			select {}
		}
		return nil
	}, supervisor.WithExistingLogtree(lt))

	// Meanwhile, wait for any fatal error from the init process, and handle it
	// accordingly.
	err := <-fatal
	// Log error with primary logging mechanism still active.
	logger.Infof("Node startup failed: %v", err)
	// Start shutting down the supervision tree...
	ctxC()
	time.Sleep(time.Second)
	// After a bit, kill all console log readers.
	for _, c := range consoles {
		if c.reader == nil {
			continue
		}
		c.reader.Close()
		c.reader.Stream = nil
	}
	// Wait for final logs to flush to console...
	time.Sleep(time.Second)
	// Present final message to the console.
	crash <- ""
	crash <- ""
	crash <- fmt.Sprintf(" Fatal error: %v", err)
	crash <- fmt.Sprintf(" This node could not be started. Rebooting...")
	time.Sleep(time.Second)
	// Return to minit, which will reboot this node.
	os.Exit(0)
}

// consoleFilter is used to filter out some uselessly verbose logs from the
// console.
//
// This should be limited to external services, our internal services should
// instead just have good logging by default.
func consoleFilter(p *logtree.LogEntry) bool {
	if p.Raw != nil {
		return false
	}
	if p.Leveled == nil {
		return false
	}
	s := string(p.DN)
	if strings.HasPrefix(s, "root.role.controlplane.launcher.consensus.etcd") {
		return p.Leveled.Severity().AtLeast(logtree.WARNING)
	}
	// TODO(q3k): turn off RPC traces instead
	if strings.HasPrefix(s, "root.role.controlplane.launcher.curator.listener.rpc") {
		return false
	}
	if strings.HasPrefix(s, "root.role.kubernetes.run.kubernetes.networked.kubelet") {
		return p.Leveled.Severity().AtLeast(logtree.WARNING)
	}
	if strings.HasPrefix(s, "root.role.kubernetes.run.kubernetes.networked.apiserver") {
		return p.Leveled.Severity().AtLeast(logtree.WARNING)
	}
	if strings.HasPrefix(s, "root.role.kubernetes.run.kubernetes.controller-manager") {
		return p.Leveled.Severity().AtLeast(logtree.WARNING)
	}
	if strings.HasPrefix(s, "root.role.kubernetes.run.kubernetes.scheduler") {
		return p.Leveled.Severity().AtLeast(logtree.WARNING)
	}
	if strings.HasPrefix(s, "root.kernel") {
		// Linux writes high-severity logs directly to the console anyways and
		// its low-severity logs are too verbose.
		return false
	}
	if strings.HasPrefix(s, "supervisor") {
		return p.Leveled.Severity().AtLeast(logtree.WARNING)
	}
	return true
}

type console struct {
	path     string
	maxWidth int
	reader   *logtree.LogReader
}
