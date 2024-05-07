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

package eventlog

// This file contains compatibility functions for our TPM library

import (
	"crypto"
)

// ConvertRawPCRs converts from raw PCRs to eventlog PCR structures
func ConvertRawPCRs(pcrs [][]byte) []PCR {
	var evPCRs []PCR
	for i, digest := range pcrs {
		evPCRs = append(evPCRs, PCR{DigestAlg: crypto.SHA256, Index: i, Digest: digest})
	}
	return evPCRs
}
