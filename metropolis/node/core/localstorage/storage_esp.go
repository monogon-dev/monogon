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
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	ppb "source.monogon.dev/metropolis/proto/private"
	npb "source.monogon.dev/net/proto"
	"source.monogon.dev/osbase/tpm"
)

// ESPDirectory is the EFI System Partition. It is a cleartext partition
// available to the system at early boot, and must contain all data required
// for the system to bootstrap, register into, or join a cluster.
type ESPDirectory struct {
	declarative.Directory
	Metropolis ESPMetropolisDirectory `dir:"metropolis"`
	EFI        ESPEFIDirectory        `dir:"ESP"`
}

type ESPEFIDirectory struct {
	declarative.Directory
	Boot       ESPBootDirectory          `dir:"BOOT"`
	Metropolis ESPEFIMetropolisDirectory `dir:"metropolis"`
}

type ESPEFIMetropolisDirectory struct {
	declarative.Directory
	BootA declarative.File `file:"boot-a.efi"`
	BootB declarative.File `file:"boot-b.efi"`
}

type ESPBootDirectory struct {
	declarative.Directory
}

// ESPMetropolisDirectory is the directory inside the EFI System Partition where
// Metropolis-related data is stored that's not read by EFI itself like
// bootstrap-related data.
type ESPMetropolisDirectory struct {
	declarative.Directory
	SealedConfiguration  ESPSealedConfiguration  `file:"sealed_configuration.pb"`
	NodeParameters       ESPNodeParameters       `file:"parameters.pb"`
	ClusterDirectory     ESPClusterDirectory     `file:"cluster_directory.pb"`
	NetworkConfiguration ESPNetworkConfiguration `file:"network_configuration.pb"`
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

// ESPClusterDirectory is a serialized common.ClusterDirectory protobuf. It
// contains a list of endpoints a registered node might connect to when joining
// a cluster.
type ESPClusterDirectory struct {
	declarative.File
}

// ESPNetworkConfiguration is a serialized net.Net protobuf. If present, it
// disables automatic network configuration and uses the given configuration
// to enable network connectivity.
type ESPNetworkConfiguration struct {
	declarative.File
}

var (
	ErrNoSealed               = errors.New("no sealed configuration exists")
	ErrSealedUnavailable      = errors.New("sealed configuration temporary unavailable")
	ErrSealedCorrupted        = errors.New("sealed configuration corrupted")
	ErrNoParameters           = errors.New("no parameters found")
	ErrParametersCorrupted    = errors.New("parameters corrupted")
	ErrNoDirectory            = errors.New("no cluster directory found")
	ErrDirectoryCorrupted     = errors.New("cluster directory corrupted")
	ErrNetworkConfigCorrupted = errors.New("network configuration corrupted")
)

func (e *ESPNodeParameters) Unmarshal() (*apb.NodeParameters, error) {
	bytes, err := e.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoParameters
		}
		return nil, fmt.Errorf("%w: when reading sealed data: %v", ErrNoParameters, err)
	}

	var config apb.NodeParameters
	err = proto.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("%w: when unmarshaling: %v", ErrParametersCorrupted, err)
	}

	return &config, nil
}

func (e *ESPClusterDirectory) Unmarshal() (*cpb.ClusterDirectory, error) {
	bytes, err := e.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoDirectory
		}
		return nil, fmt.Errorf("%w: when reading: %v", ErrNoDirectory, err)
	}

	var dir cpb.ClusterDirectory
	err = proto.Unmarshal(bytes, &dir)
	if err != nil {
		return nil, fmt.Errorf("%w: when unmarshaling: %v", ErrDirectoryCorrupted, err)
	}
	return &dir, nil
}

func (e *ESPNetworkConfiguration) Unmarshal() (*npb.Net, error) {
	bytes, err := e.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w: when reading: %v", ErrNetworkConfigCorrupted, err)
	}

	var netConf npb.Net
	err = proto.Unmarshal(bytes, &netConf)
	if err != nil {
		return nil, fmt.Errorf("%w: when unmarshaling: %v", ErrNetworkConfigCorrupted, err)
	}
	return &netConf, nil
}

func (e *ESPNetworkConfiguration) Marshal(n *npb.Net) error {
	netConfRaw, err := proto.Marshal(n)
	if err != nil {
		return fmt.Errorf("error marshaling Net: %w", err)
	}
	if err := e.Write(netConfRaw, 0666); err != nil {
		return fmt.Errorf("error writing static network config to ESP: %w", err)
	}
	return nil
}

func (e *ESPSealedConfiguration) SealSecureBoot(c *ppb.SealedConfiguration, tpmUsage cpb.NodeTPMUsage) error {
	bytes, err := proto.Marshal(c)
	if err != nil {
		return fmt.Errorf("while marshaling: %w", err)
	}

	switch tpmUsage {
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_AND_USED:
		// Use Secure Boot PCRs to seal the configuration.
		// See: TCG PC Client Platform Firmware Profile Specification v1.05,
		//      table 3.3.4.1
		// See: https://trustedcomputinggroup.org/wp-content/uploads/
		//      TCG_PCClient_PFP_r1p05_v22_02dec2020.pdf
		bytes, err = tpm.Seal(bytes, tpm.SecureBootPCRs)
		if err != nil {
			return fmt.Errorf("while using tpm: %w", err)
		}
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_BUT_UNUSED:
	case cpb.NodeTPMUsage_NODE_TPM_NOT_PRESENT:
	default:
		return fmt.Errorf("unknown tpmUsage %d", tpmUsage)
	}

	if err := e.Write(bytes, 0644); err != nil {
		return fmt.Errorf("while writing: %w", err)
	}
	return nil
}

func (e *ESPSealedConfiguration) Unseal(tpmUsage cpb.NodeTPMUsage) (*ppb.SealedConfiguration, error) {
	bytes, err := e.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoSealed
		}
		return nil, fmt.Errorf("%w: when reading sealed data: %v", ErrSealedUnavailable, err)
	}

	switch tpmUsage {
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_AND_USED:
		bytes, err = tpm.Unseal(bytes)
		if err != nil {
			return nil, fmt.Errorf("%w: when unsealing: %v", ErrSealedCorrupted, err)
		}
	case cpb.NodeTPMUsage_NODE_TPM_PRESENT_BUT_UNUSED:
	case cpb.NodeTPMUsage_NODE_TPM_NOT_PRESENT:
	default:
		return nil, fmt.Errorf("unknown tpmUsage %d", tpmUsage)
	}

	var config ppb.SealedConfiguration
	err = proto.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("%w: when unmarshaling: %v", ErrSealedCorrupted, err)
	}

	return &config, nil
}
