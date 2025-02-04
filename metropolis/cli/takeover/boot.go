// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node/core/devmgr"
	"source.monogon.dev/osbase/supervisor"
)

// Main runnable for the installer.
func takeoverRunnable(ctx context.Context) error {
	l := supervisor.Logger(ctx)

	devmgrSvc := devmgr.New()
	supervisor.Run(ctx, "devmgr", devmgrSvc.Run)
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	for {
		devicePath := filepath.Join("/dev", os.Getenv(EnvInstallTarget))
		l.Infof("Waiting for device: %s", devicePath)
		_, err := os.Stat(devicePath)
		if os.IsNotExist(err) {
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			return err
		}
		break
	}

	if err := installMetropolis(l); err != nil {
		l.Errorf("Installation failed: %v", err)
	} else {
		l.Info("Installation succeeded")
	}

	time.Sleep(1 * time.Second)
	unix.Sync()
	unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)

	return nil
}
