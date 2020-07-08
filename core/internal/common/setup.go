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
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"strings"
)

type (
	SmalltownState string
)

const (
	// These are here to prevent depdendency loops
	NodeServicePort     = 7835
	ConsensusPort       = 7834
	MasterServicePort   = 7833
	ExternalServicePort = 7836
	DebugServicePort    = 7837
	KubernetesAPIPort   = 6443
	DebuggerPort        = 2345
)

const (
	// Node is provisioning a new cluster with itself as a master
	StateNewClusterMode SmalltownState = "setup"
	// Node is enrolling itself and waiting to be adopted
	StateEnrollMode SmalltownState = "join"
	// Node is fully provisioned.
	StateJoined SmalltownState = "enrolled"
)

func NameFromIDKey(pubKey ed25519.PublicKey) string {
	return "smalltown-" + hex.EncodeToString(pubKey[:16])
}

func IDKeyPrefixFromName(name string) ([]byte, error) {
	if !strings.HasPrefix(name, "smalltown-") {
		return []byte{}, errors.New("invalid name")
	}
	return hex.DecodeString(strings.TrimPrefix(name, "smalltown-"))
}
