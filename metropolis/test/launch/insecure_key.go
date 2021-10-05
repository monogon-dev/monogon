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

package launch

import (
	"crypto/ed25519"

	apb "source.monogon.dev/metropolis/proto/api"
)

var (
	// InsecurePrivateKey is a ED25519 key that can be used during development as
	// a fixed Owner Key when bootstrapping new clusters.
	InsecurePrivateKey = ed25519.NewKeyFromSeed([]byte(
		"\xb5\xcf\x35\x0f\xbf\xfb\xea\xfa\xa0\xf0\x29\x9d\xfa\xf7\xca\x6f" +
			"\xa2\xc2\xc7\x87\xd7\x03\x3e\xb2\x11\x4f\x36\xe0\x22\x73\x4f\x87"))
	// InsecurePublicKey is the ED25519 public key corresponding to
	// InsecurePrivateKey.
	InsecurePublicKey = InsecurePrivateKey.Public().(ed25519.PublicKey)
	// InsecureClusterBootstrap is a ClusterBootstrap message to be used within
	// NodeParameters when bootstrapping a development cluster that should be owned
	// by the InsecurePrivateKey.
	InsecureClusterBootstrap = &apb.NodeParameters_ClusterBootstrap{
		OwnerPublicKey: InsecurePublicKey,
	}
)
