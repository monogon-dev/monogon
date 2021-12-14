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

package localstorage

import (
	"errors"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/pkg/tpm"
	apb "source.monogon.dev/metropolis/proto/api"
	ppb "source.monogon.dev/metropolis/proto/private"
)

// ESPDirectory is the EFI System Partition. It is a cleartext partition
// available to the system at early boot, and must contain all data required
// for the system to bootstrap, register into, or join a cluster.
type ESPDirectory struct {
	declarative.Directory
	Metropolis ESPMetropolisDirectory `dir:"metropolis"`
}

// ESPMetropolisDirectory is the directory inside the EFI System Partition where
// Metropolis-related data is stored that's not read by EFI itself like
// bootstrap-related data.
type ESPMetropolisDirectory struct {
	declarative.Directory
	SealedConfiguration ESPSealedConfiguration `file:"sealed_configuration.pb"`
	NodeParameters      ESPNodeParameters      `file:"parameters.pb"`
}

// ESPSealedConfiguration is a TPM sealed serialized
// private.SealedConfiguration protobuf. It contains all data required for a
// node to be able to join a cluster after startup.
type ESPSealedConfiguration struct {
	declarative.File
}

// ESPNodeParameters is the configuration for this node when first
// bootstrapping a cluster or registering into an existing one. It's a
// api.NodeParameters protobuf message.
type ESPNodeParameters struct {
	declarative.File
}

var (
	ErrNoSealed            = errors.New("no sealed configuration exists")
	ErrSealedUnavailable   = errors.New("sealed configuration temporary unavailable")
	ErrSealedCorrupted     = errors.New("sealed configuration corrupted")
	ErrNoParameters        = errors.New("no parameters found")
	ErrParametersCorrupted = errors.New("parameters corrupted")
)

func (e *ESPNodeParameters) Unmarshal() (*apb.NodeParameters, error) {
	bytes, err := e.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoParameters
		}
		return nil, fmt.Errorf("%w: when reading sealed data: %v", ErrNoParameters, err)
	}

	config := apb.NodeParameters{}
	err = proto.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("%w: when unmarshaling: %v", ErrParametersCorrupted, err)
	}

	return &config, nil
}

func (e *ESPSealedConfiguration) Unseal() (*ppb.SealedConfiguration, error) {
	bytes, err := e.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoSealed
		}
		return nil, fmt.Errorf("%w: when reading sealed data: %v", ErrSealedUnavailable, err)
	}

	bytes, err = tpm.Unseal(bytes)
	if err != nil {
		return nil, fmt.Errorf("%w: when unsealing: %v", ErrSealedCorrupted, err)
	}

	config := ppb.SealedConfiguration{}
	err = proto.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("%w: when unmarshaling: %v", ErrSealedCorrupted, err)
	}

	return &config, nil
}
