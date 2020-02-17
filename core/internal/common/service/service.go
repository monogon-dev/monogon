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

package service

import (
	"context"
	"errors"
	"sync"

	"go.uber.org/zap"
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

		// A context that represents the lifecycle of a service.
		// It is created right before impl.OnStart, and canceled
		// right after impl.OnStop is.
		// This is a transition mechanism from moving from OnStart/OnStop
		// based lifecycle management of services to a context-based supervision
		// tree.
		// Service implementations should access this via .Context()
		ctx  *context.Context
		ctxC *context.CancelFunc
	}
)

func NewBaseService(name string, logger *zap.Logger, impl Service) *BaseService {
	return &BaseService{
		Logger: logger,
		name:   name,
		impl:   impl,
		ctx:    nil,
		ctxC:   nil,
	}
}

// Start starts the service. This is an atomic operation and should not be called on an already running service.
func (b *BaseService) Start() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.running {
		return ErrAlreadyRunning
	}

	ctx, ctxC := context.WithCancel(context.Background())
	b.ctx = &ctx
	b.ctxC = &ctxC

	err := b.impl.OnStart()
	if err != nil {
		b.Logger.Error("Failed to start service", zap.String("service", b.name), zap.Error(err))
		return err
	}

	b.running = true
	b.Logger.Info("Started service", zap.String("service", b.name))
	return nil
}

// Stop stops the service. This is an atomic operation and should only be called on a running service.
func (b *BaseService) Stop() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if !b.running {
		return ErrNotRunning
	}

	err := b.impl.OnStop()
	if err != nil {
		b.Logger.Error("Failed to stop service", zap.String("service", b.name), zap.Error(err))

		return err
	}

	b.running = false

	// Kill context
	(*b.ctxC)()
	b.ctx = nil
	b.ctxC = nil

	b.Logger.Info("Stopped service", zap.String("service", b.name))
	return nil
}

// IsRunning returns whether the service is currently running.
func (b *BaseService) IsRunning() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.running
}

// Context returns a context that can be used within OnStart() to create new
// lightweight subservices that use a context for lifecycle management.
// This is a transition measure before the Service library is rewritten to use
// a more advanced context-and-returned-error supervision tree.
// This context can also be used for blocking operations like IO, etc.
func (b *BaseService) Context() context.Context {
	return *b.ctx
}
