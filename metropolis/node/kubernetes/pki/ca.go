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

package pki

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"math/big"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// Issuer is a CA that can issue certificates. Two issuers are currently implemented:
//  - SelfSigned, which will generated a certificate signed by its corresponding private key.
//  - Certificate, which will use another existing Certificate as a CA.
type Issuer interface {
	// CACertificate returns the DER-encoded x509 certificate of the CA that will sign certificates when Issue is
	// called, or nil if this is self-signing issuer.
	CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error)
	// Issue will generate a key and certificate signed by the Issuer. The returned certificate is x509 DER-encoded,
	// while the key is a bare ed25519 key.
	Issue(ctx context.Context, template x509.Certificate, kv clientv3.KV) (cert, key []byte, err error)
}

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

// Workaround for https://github.com/golang/go/issues/26676 in Go's crypto/x509. Specifically Go
// violates Section 4.2.1.2 of RFC 5280 without this.
// Fixed for 1.15 in https://go-review.googlesource.com/c/go/+/227098/.
//
// Taken from https://github.com/FiloSottile/mkcert/blob/master/cert.go#L295 written by one of Go's
// crypto engineers
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

// issueCertificate is a generic low level certificate-and-key issuance function. If ca or cakey is null, the
// certificate will be self-signed. The returned certificate is DER-encoded, while the returned key is internal.
func issueCertificate(template x509.Certificate, ca *x509.Certificate, caKey interface{}) (cert, key []byte, err error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("failed to generate serial number: %w", err)
		return
	}

	skid, err := calculateSKID(pubKey)
	if err != nil {
		return []byte{}, privKey, err
	}

	template.SerialNumber = serialNumber
	template.NotBefore = time.Now()
	template.NotAfter = unknownNotAfter
	template.BasicConstraintsValid = true
	template.SubjectKeyId = skid

	// Set the AuthorityKeyID to the SKID of the signing certificate (or self, if self-signing).
	if ca != nil && caKey != nil {
		template.AuthorityKeyId = ca.AuthorityKeyId
	} else {
		template.AuthorityKeyId = template.SubjectKeyId
	}

	if ca == nil || caKey == nil {
		ca = &template
		caKey = privKey
	}

	caCertRaw, err := x509.CreateCertificate(rand.Reader, &template, ca, pubKey, caKey)
	return caCertRaw, privKey, err
}

type selfSigned struct{}

func (s *selfSigned) Issue(ctx context.Context, template x509.Certificate, kv clientv3.KV) (cert, key []byte, err error) {
	return issueCertificate(template, nil, nil)
}

func (s *selfSigned) CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error) {
	return nil, nil
}

var (
	// SelfSigned is an Issuer that generates self-signed certificates.
	SelfSigned = &selfSigned{}
)

func (c *Certificate) Issue(ctx context.Context, template x509.Certificate, kv clientv3.KV) (cert, key []byte, err error) {
	caCert, caKey, err := c.ensure(ctx, kv)
	if err != nil {
		return nil, nil, fmt.Errorf("could not ensure CA certificate %q exists: %w", c.name, err)
	}

	ca, err := x509.ParseCertificate(caCert)
	if err != nil {
		return nil, nil, fmt.Errorf("could not parse CA certificate: %w", err)
	}
	// Ensure only one level of CAs exist, and that they are created explicitly.
	template.IsCA = false
	return issueCertificate(template, ca, ed25519.PrivateKey(caKey))
}

func (c *Certificate) CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error) {
	cert, _, err := c.ensure(ctx, kv)
	return cert, err
}
