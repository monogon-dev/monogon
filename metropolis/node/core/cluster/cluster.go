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

package cluster

import (
	"fmt"

	"source.monogon.dev/metropolis/pkg/pki"
)

type ClusterState int

const (
	ClusterUnknown ClusterState = iota
	ClusterForeign
	ClusterTrusted
	ClusterHome
	ClusterDisowning
)

type Cluster struct {
	State ClusterState
}

func (s ClusterState) String() string {
	switch s {
	case ClusterForeign:
		return "ClusterForeign"
	case ClusterTrusted:
		return "ClusterTrusted"
	case ClusterHome:
		return "ClusterHome"
	case ClusterDisowning:
		return "ClusterDisowning"
	}
	return fmt.Sprintf("Invalid(%d)", s)
}

var clusterStateTransitions = map[ClusterState][]ClusterState{
	ClusterUnknown: {ClusterForeign, ClusterHome, ClusterDisowning},
	ClusterForeign: {ClusterTrusted},
	ClusterTrusted: {ClusterHome},
	ClusterHome:    {ClusterHome, ClusterDisowning},
}

var (
	PKINamespace = pki.Namespaced("/cluster-pki/")
	PKICA        = PKINamespace.New(pki.SelfSigned, "cluster-ca", pki.CA("Metropolis Cluster CA"))
)
