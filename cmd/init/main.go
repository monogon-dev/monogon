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
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	node2 "smalltown/internal/node"
	"smalltown/internal/storage"
	"smalltown/pkg/tpm"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sys/unix"
)

func main() {
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

	if err := storage.FindPartitions(); err != nil {
		logger.Panic("Failed to search for partitions", zap.Error(err))
	}

	if err := os.Mkdir("/esp", 0755); err != nil {
		panic(err)
	}

	if err := unix.Mount(storage.ESPDevicePath, "/esp", "vfat", unix.MS_NOEXEC|unix.MS_NODEV|unix.MS_SYNC, ""); err != nil {
		logger.Panic("Failed to mount ESP partition", zap.Error(err))
	}

	if err := tpm.Initialize(logger.With(zap.String("component", "tpm"))); err != nil {
		logger.Panic("Failed to initialize TPM 2.0", zap.Error(err))
	}

	// TODO(lorenz): This really doesn't belong here and needs to be asynchronous as well
	var keyLocation = "/esp/EFI/smalltown/data-key.bin"
	sealedKeyFile, err := os.Open(keyLocation)
	if os.IsNotExist(err) {
		logger.Info("Initializing encrypted storage, this might take a while...")
		key, err := tpm.GenerateSafeKey(256 / 8)
		if err != nil {
			panic(err)
		}
		sealedKey, err := tpm.Seal(key, tpm.SecureBootPCRs)
		if err != nil {
			panic(err)
		}
		if err := storage.InitializeEncryptedBlockDevice("data", storage.SmalltownDataCryptPath, key); err != nil {
			panic(err)
		}
		mkfsCmd := exec.Command("/bin/mkfs.xfs", "-qf", "/dev/data")
		if _, err := mkfsCmd.Output(); err != nil {
			panic(err)
		}
		// Existence of this file indicates that the encrypted storage has been successfully initialized
		if err := ioutil.WriteFile(keyLocation, sealedKey, 0600); err != nil {
			panic(err)
		}
		logger.Info("Initialized encrypted storage")
	} else if err != nil {
		panic(err)
	} else {
		sealedKey, err := ioutil.ReadAll(sealedKeyFile)
		if err != nil {
			panic(err)
		}
		key, err := tpm.Unseal(sealedKey)
		if err != nil {
			panic(err)
		}
		if err := storage.MapEncryptedBlockDevice("data", storage.SmalltownDataCryptPath, key); err != nil {
			panic(err)
		}
		logger.Info("Opened encrypted storage")
	}
	sealedKeyFile.Close()

	if err := os.Mkdir("/data", 0755); err != nil {
		panic(err)
	}

	if err := unix.Mount("/dev/data", "/data", "xfs", unix.MS_NOEXEC|unix.MS_NODEV, ""); err != nil {
		panic(err)
	}

	node, err := node2.NewSmalltownNode(logger, "/esp/EFI/smalltown", "/data", 7833, 7834)
	if err != nil {
		panic(err)
	}

	err = node.Start()
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
				res, err := unix.Wait4(-1, &status, syscall.WNOHANG, &rusage)
				if err != nil {
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
