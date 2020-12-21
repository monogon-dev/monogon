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
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"git.monogon.dev/source/nexantic.git/metropolis/node/core/localstorage/declarative"
)

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

type CertificateTemplateNamer func(pubkey []byte) x509.Certificate

func CertificateForNode(pubkey []byte) x509.Certificate {
	// TODO(q3k): this should be unified with metroopolis/node/cluster:node.ID()
	name := "metropolis-" + hex.EncodeToString([]byte(pubkey[:16]))

	// This has no SANs because it authenticates by public key, not by name
	return x509.Certificate{
		Subject: pkix.Name{
			// We identify nodes by their ID public keys (not hashed since a strong hash is longer and serves no benefit)
			CommonName: name,
		},
		IsCA:                  false,
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              unknownNotAfter,
		// Certificate is used both as server & client
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}
}

func (p *PKIDirectory) EnsureSelfSigned(namer CertificateTemplateNamer) (*tls.Certificate, error) {
	create := false
	for _, f := range []*declarative.File{&p.Certificate, &p.Key} {
		exists, err := f.Exists()
		if err != nil {
			return nil, fmt.Errorf("could not check existence of file %q: %w", f.FullPath(), err)
		}
		if !exists {
			create = true
			break
		}
	}

	if !create {
		certRaw, err := p.Certificate.Read()
		if err != nil {
			return nil, fmt.Errorf("could not read certificate: %w", err)
		}
		privKeyRaw, err := p.Key.Read()
		if err != nil {
			return nil, fmt.Errorf("could not read key: %w", err)
		}
		cert, err := x509.ParseCertificate(certRaw)
		if err != nil {
			return nil, fmt.Errorf("could not parse certificate: %w", err)
		}
		privKey, err := x509.ParsePKCS8PrivateKey(privKeyRaw)
		if err != nil {
			return nil, fmt.Errorf("could not parse key: %w", err)
		}
		return &tls.Certificate{
			Certificate: [][]byte{certRaw},
			PrivateKey:  privKey,
			Leaf:        cert,
		}, nil
	}

	pubKey, privKeyRaw, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}

	privKey, err := x509.MarshalPKCS8PrivateKey(privKeyRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal key: %w", err)
	}

	if err := p.Key.Write(privKey, 0600); err != nil {
		return nil, fmt.Errorf("failed to write new private key: %w", err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %w", err)
	}

	template := namer(pubKey)
	template.SerialNumber = serialNumber

	certRaw, err := x509.CreateCertificate(rand.Reader, &template, &template, pubKey, privKeyRaw)
	if err != nil {
		return nil, fmt.Errorf("could not sign certificate: %w", err)
	}

	cert, err := x509.ParseCertificate(certRaw)
	if err != nil {
		return nil, fmt.Errorf("could not parse newly created certificate: %w", err)
	}

	if err := p.Certificate.Write(certRaw, 0600); err != nil {
		return nil, fmt.Errorf("failed to write new certificate: %w", err)
	}

	return &tls.Certificate{
		Certificate: [][]byte{certRaw},
		PrivateKey:  privKey,
		Leaf:        cert,
	}, nil
}

// AllExist returns true if all PKI files (cert, key, CA cert) are present on the backing
// store.
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

// AllAbsent returns true if all PKI files (cert, key, CA cert) are missing from the backing
// store.
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
