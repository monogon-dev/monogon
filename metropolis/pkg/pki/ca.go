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
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"math/big"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// Issuer is an entity that can issue certificates. This interface is
// implemented by SelfSigned, which is an issuer that emits self-signed
// certificates, and any other Certificate that has been created with CA(),
// which makes this Certificate act as a CA and issue (sign) ceritficates.
type Issuer interface {
	// CACertificate returns the DER-encoded x509 certificate of the CA that
	// will sign certificates when Issue is called, or nil if this is
	// self-signing issuer.
	CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error)
	// Issue will generate a key and certificate signed by the Issuer. The
	// returned certificate is x509 DER-encoded, while the key is a bare
	// ed25519 key.
	Issue(ctx context.Context, req *Certificate, kv clientv3.KV) (cert, key []byte, err error)
}

// issueCertificate is a generic low level certificate-and-key issuance
// function. If ca or cakey is null, the certificate will be self-signed. The
// returned certificate is DER-encoded, while the returned key is internal.
func issueCertificate(req *Certificate, ca *x509.Certificate, caKey interface{}) (cert, key []byte, err error) {
	var privKey ed25519.PrivateKey
	var pubKey ed25519.PublicKey
	if req.key != nil {
		privKey = req.key
		pubKey = privKey.Public().(ed25519.PublicKey)
	} else {
		var err error
		pubKey, privKey, err = ed25519.GenerateKey(rand.Reader)
		if err != nil {
			panic(err)
		}
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

	req.template.SerialNumber = serialNumber
	req.template.NotBefore = time.Now()
	req.template.NotAfter = unknownNotAfter
	req.template.BasicConstraintsValid = true
	req.template.SubjectKeyId = skid

	// Set the AuthorityKeyID to the SKID of the signing certificate (or self,
	// if self-signing).
	if ca != nil && caKey != nil {
		req.template.AuthorityKeyId = ca.AuthorityKeyId
	} else {
		req.template.AuthorityKeyId = req.template.SubjectKeyId
	}

	if ca == nil || caKey == nil {
		ca = &req.template
		caKey = privKey
	}

	caCertRaw, err := x509.CreateCertificate(rand.Reader, &req.template, ca, pubKey, caKey)
	return caCertRaw, privKey, err
}

type selfSigned struct{}

var (
	// SelfSigned is an Issuer that generates self-signed certificates.
	SelfSigned = &selfSigned{}
)

// Issue will generate a key and certificate that is self-signed.
func (s *selfSigned) Issue(ctx context.Context, req *Certificate, kv clientv3.KV) (cert, key []byte, err error) {
	return issueCertificate(req, nil, nil)
}

// CACertificate returns nil for self-signed issuers.
func (s *selfSigned) CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error) {
	return nil, nil
}

// Issue will generate a key and certificate that is signed by this
// Certificate, if the Certificate is a CA.
func (c *Certificate) Issue(ctx context.Context, req *Certificate, kv clientv3.KV) (cert, key []byte, err error) {
	caCert, caKey, err := c.ensure(ctx, kv)
	if err != nil {
		return nil, nil, fmt.Errorf("could not ensure CA certificate %q exists: %w", c.name, err)
	}

	ca, err := x509.ParseCertificate(caCert)
	if err != nil {
		return nil, nil, fmt.Errorf("could not parse CA certificate: %w", err)
	}
	// Ensure only one level of CAs exist, and that they are created explicitly.
	req.template.IsCA = false
	return issueCertificate(req, ca, ed25519.PrivateKey(caKey))
}

// CACertificate returns the DER encoded x509 form of this Certificate that
// will be the used to issue child certificates.
func (c *Certificate) CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error) {
	cert, _, err := c.ensure(ctx, kv)
	return cert, err
}
