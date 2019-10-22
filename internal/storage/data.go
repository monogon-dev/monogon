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

package storage

import (
	"fmt"
	"git.monogon.dev/source/smalltown.git/internal/common"
	"git.monogon.dev/source/smalltown.git/pkg/tpm"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/sys/unix"
)

const (
	dataMountPath         = "/data"
	espMountPath          = "/esp"
	espDataPath           = espMountPath + "/EFI/smalltown"
	etcdSealedKeyLocation = espDataPath + "/data-key.bin"
)

type Manager struct {
	logger              *zap.Logger
	dataReady           bool
	initializationError error
	mutex               sync.RWMutex
}

func Initialize(logger *zap.Logger) (*Manager, error) {
	if err := FindPartitions(); err != nil {
		return nil, err
	}

	if err := os.Mkdir("/esp", 0755); err != nil {
		return nil, err
	}

	// We're mounting ESP sync for reliability, this lowers our chances of getting half-written files
	if err := unix.Mount(ESPDevicePath, espMountPath, "vfat", unix.MS_NOEXEC|unix.MS_NODEV|unix.MS_SYNC, ""); err != nil {
		return nil, err
	}

	manager := &Manager{
		logger:    logger,
		dataReady: false,
	}

	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	sealedKeyFile, err := os.Open(etcdSealedKeyLocation)
	if os.IsNotExist(err) {
		logger.Info("Initializing encrypted storage, this might take a while...")
		go manager.initializeData()
	} else if err != nil {
		return nil, err
	} else {
		sealedKey, err := ioutil.ReadAll(sealedKeyFile)
		sealedKeyFile.Close()
		if err != nil {
			return nil, err
		}
		key, err := tpm.Unseal(sealedKey)
		if err != nil {
			return nil, err
		}
		if err := MapEncryptedBlockDevice("data", SmalltownDataCryptPath, key); err != nil {
			return nil, err
		}
		if err := manager.mountData(); err != nil {
			return nil, err
		}
		logger.Info("Mounted encrypted storage")
	}
	return manager, nil
}

func (s *Manager) initializeData() {
	key, err := tpm.GenerateSafeKey(256 / 8)
	if err != nil {
		s.logger.Error("Failed to generate master key", zap.Error(err))
		s.initializationError = fmt.Errorf("Failed to generate master key: %w", err)
		return
	}
	sealedKey, err := tpm.Seal(key, tpm.FullSystemPCRs)
	if err != nil {
		s.logger.Error("Failed to seal master key", zap.Error(err))
		s.initializationError = fmt.Errorf("Failed to seal master key: %w", err)
		return
	}
	if err := InitializeEncryptedBlockDevice("data", SmalltownDataCryptPath, key); err != nil {
		s.logger.Error("Failed to initialize encrypted block device", zap.Error(err))
		s.initializationError = fmt.Errorf("Failed to initialize encrypted block device: %w", err)
		return
	}
	mkfsCmd := exec.Command("/bin/mkfs.xfs", "-qf", "/dev/data")
	if _, err := mkfsCmd.Output(); err != nil {
		s.logger.Error("Failed to format encrypted block device", zap.Error(err))
		s.initializationError = fmt.Errorf("Failed to format encrypted block device: %w", err)
		return
	}
	// This file is the marker if the partition has
	if err := ioutil.WriteFile(etcdSealedKeyLocation, sealedKey, 0600); err != nil {
		panic(err)
	}

	if err := s.mountData(); err != nil {
		s.initializationError = err
		return
	}

	s.mutex.Lock()
	s.dataReady = true
	s.mutex.Unlock()

	s.logger.Info("Initialized encrypted storage")
}

func (s *Manager) mountData() error {
	if err := os.Mkdir("/data", 0755); err != nil {
		return err
	}

	if err := unix.Mount("/dev/data", "/data", "xfs", unix.MS_NOEXEC|unix.MS_NODEV, ""); err != nil {
		return err
	}
	return nil
}

// GetPathInPlace returns a path in the given place
// It may return ErrNotInitialized if the place you're trying to access
// is not initialized or ErrUnknownPlace if the place is not known
func (s *Manager) GetPathInPlace(place common.DataPlace, path string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	switch place {
	case common.PlaceESP:
		return filepath.Join(espDataPath, path), nil
	case common.PlaceData:
		if s.dataReady {
			return filepath.Join(dataMountPath, path), nil
		}
		return "", common.ErrNotInitialized
	default:
		return "", common.ErrUnknownPlace
	}
}
