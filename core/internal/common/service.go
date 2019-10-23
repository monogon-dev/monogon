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

package common

import (
	"errors"
	"go.uber.org/zap"
	"sync"
)

var (
	ErrAlreadyRunning = errors.New("service is already running")
	ErrNotRunning     = errors.New("service is not running")
)

type (
	// Service represents a subsystem of an application that can be used with a BaseService.
	Service interface {
		OnStart() error
		OnStop() error
	}

	// BaseService implements utility functionality around a service.
	BaseService struct {
		impl Service
		name string

		Logger *zap.Logger

		mutex   sync.Mutex
		running bool
	}
)

func NewBaseService(name string, logger *zap.Logger, impl Service) *BaseService {
	return &BaseService{
		Logger: logger,
		name:   name,
		impl:   impl,
	}
}

// Start starts the service. This is an atomic operation and should not be called on an already running service.
func (b *BaseService) Start() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.running {
		return ErrAlreadyRunning
	}

	err := b.impl.OnStart()
	if err != nil {
		b.Logger.Error("Failed to start service", zap.String("service", b.name), zap.Error(err))
		return err
	}

	b.running = true
	b.Logger.Info("Started service", zap.String("service", b.name))
	return nil
}

// Stop stops the service. THis is an atomic operation and should only be called on a running service.
func (b *BaseService) Stop() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if !b.running {
		return ErrNotRunning
	}

	err := b.impl.OnStart()
	if err != nil {
		b.Logger.Error("Failed to stop service", zap.String("service", b.name), zap.Error(err))

		return err
	}

	b.running = false
	b.Logger.Info("Stopped service", zap.String("service", b.name))
	return nil
}

// IsRunning returns whether the service is currently running.
func (b *BaseService) IsRunning() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.running
}
