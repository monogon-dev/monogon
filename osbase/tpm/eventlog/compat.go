// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
