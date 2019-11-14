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

import "git.monogon.dev/source/nexantic.git/core/generated/api"

// TODO(leo): merge api and node packages and get rid of this extra layer of indirection?

type (
	SetupService interface {
		CurrentState() SmalltownState
		GetJoinClusterToken() string
		SetupNewCluster() error
		EnterJoinClusterMode() error
		JoinCluster(initialCluster string, certs *api.ConsensusCertificates) error
	}

	SmalltownState string
)

const (
	// Node is unprovisioned and waits for Setup to be called.
	StateSetupMode SmalltownState = "setup"
	// Setup() has been called, node waits for a JoinCluster or BootstrapCluster call.
	StateClusterJoinMode SmalltownState = "join"
	// Node is fully provisioned.
	StateConfigured SmalltownState = "configured"
)
