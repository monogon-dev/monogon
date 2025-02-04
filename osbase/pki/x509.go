// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package pki

import (
	"crypto"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"time"
)

var (
	// From RFC 5280 Section 4.1.2.5
	UnknownNotAfter = time.Unix(253402300799, 0)
)

// Workaround for https://github.com/golang/go/issues/26676 in Go's
// crypto/x509. Specifically Go violates Section 4.2.1.2 of RFC 5280 without
// this. Fixed for 1.15 in https://go-review.googlesource.com/c/go/+/227098/.
//
// Taken from https://github.com/FiloSottile/mkcert/blob/master/cert.go#L295
// Written by one of Go's crypto engineers
//
// TODO(lorenz): remove this once we migrate to Go 1.15.
func calculateSKID(pubKey crypto.PublicKey) ([]byte, error) {
	spkiASN1, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	var spki struct {
		Algorithm        pkix.AlgorithmIdentifier
		SubjectPublicKey asn1.BitString
	}
	_, err = asn1.Unmarshal(spkiASN1, &spki)
	if err != nil {
		return nil, err
	}
	skid := sha1.Sum(spki.SubjectPublicKey.Bytes)
	return skid[:], nil
}
