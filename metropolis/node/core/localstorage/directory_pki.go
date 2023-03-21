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
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
)

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

type CertificateTemplateNamer func(pubkey []byte) x509.Certificate

// AllExist returns true if all PKI files (cert, key, CA cert) are present on
// the backing store.
func (p *PKIDirectory) AllExist() (bool, error) {
	for _, d := range []*declarative.File{&p.CACertificate, &p.Certificate, &p.Key} {
		exists, err := d.Exists()
		if err != nil {
			return false, fmt.Errorf("failed to check %q: %v", d.FullPath(), err)
		}
		if !exists {
			return false, nil
		}
	}
	return true, nil
}

// AllAbsent returns true if all PKI files (cert, key, CA cert) are missing
// from the backing store.
func (p *PKIDirectory) AllAbsent() (bool, error) {
	for _, d := range []*declarative.File{&p.CACertificate, &p.Certificate, &p.Key} {
		exists, err := d.Exists()
		if err != nil {
			return false, fmt.Errorf("failed to check %q: %v", d.FullPath(), err)
		}
		if exists {
			return false, nil
		}
	}
	return true, nil
}

// WriteAll (over)writes the PKI data in this directory with the given private
// key, certificate and CA certificate.
//
// For ease of use, the accepted certificates are expected to already be in
// DER-encoded form (eg. from the Raw field in a x509.Certificate).
func (p *PKIDirectory) WriteAll(cert []byte, key ed25519.PrivateKey, ca []byte) error {
	if err := p.MkdirAll(0700); err != nil {
		return fmt.Errorf("failed to make PKI directory: %w", err)
	}
	keyBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return fmt.Errorf("failed to marshal key: %w", err)
	}
	if err := p.Key.Write(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes}), 0600); err != nil {
		return fmt.Errorf("failed to write key: %w", err)
	}
	if err := p.Certificate.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert}), 0600); err != nil {
		return fmt.Errorf("failed to write certificate: %w", err)
	}
	if err := p.CACertificate.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ca}), 0600); err != nil {
		return fmt.Errorf("failed to write CA certificate: %w", err)
	}
	return nil
}
