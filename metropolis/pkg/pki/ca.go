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
	// Issue will generate a certificate signed by the Issuer. The returned
	// certificate is x509 DER-encoded.
	Issue(ctx context.Context, req *Certificate, kv clientv3.KV) (cert []byte, err error)
}

// issueCertificate is a generic low level certificate-and-key issuance
// function. If ca is null, the certificate will be self-signed. The returned
// certificate is DER-encoded
func issueCertificate(req *Certificate, ca *x509.Certificate, caKey ed25519.PrivateKey) (cert []byte, err error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("failed to generate serial number: %w", err)
		return
	}

	skid, err := calculateSKID(req.PublicKey)
	if err != nil {
		return nil, err
	}

	req.Template.SerialNumber = serialNumber
	req.Template.NotBefore = time.Now()
	req.Template.NotAfter = UnknownNotAfter
	req.Template.BasicConstraintsValid = true
	req.Template.SubjectKeyId = skid

	// Set the AuthorityKeyID to the SKID of the signing certificate (or self,
	// if self-signing).
	if ca != nil {
		req.Template.AuthorityKeyId = ca.AuthorityKeyId
	} else {
		req.Template.AuthorityKeyId = req.Template.SubjectKeyId
		ca = &req.Template
	}

	return x509.CreateCertificate(rand.Reader, &req.Template, ca, req.PublicKey, caKey)
}

type selfSigned struct{}

var (
	// SelfSigned is an Issuer that generates self-signed certificates.
	SelfSigned = &selfSigned{}
)

// Issue will generate a key and certificate that is self-signed.
func (s *selfSigned) Issue(ctx context.Context, req *Certificate, kv clientv3.KV) (cert []byte, err error) {
	if err := req.ensureKey(ctx, kv); err != nil {
		return nil, err
	}
	if req.PrivateKey == nil {
		return nil, fmt.Errorf("cannot issue self-signed certificate without a private key")
	}
	return issueCertificate(req, nil, req.PrivateKey)
}

// CACertificate returns nil for self-signed issuers.
func (s *selfSigned) CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error) {
	return nil, nil
}

// Issue will generate a key and certificate that is signed by this
// Certificate, if the Certificate is a CA.
func (c *Certificate) Issue(ctx context.Context, req *Certificate, kv clientv3.KV) (cert []byte, err error) {
	if err := c.ensureKey(ctx, kv); err != nil {
		return nil, fmt.Errorf("could not ensure CA %q key exists: %w", c.Name, err)
	}
	if err := req.ensureKey(ctx, kv); err != nil {
		return nil, fmt.Errorf("could not subject %q key exists: %w", req.Name, err)
	}
	if c.PrivateKey == nil {
		return nil, fmt.Errorf("cannot use certificate without private key as CA")
	}

	caCert, err := c.ensure(ctx, kv)
	if err != nil {
		return nil, fmt.Errorf("could not ensure CA %q certificate exists: %w", c.Name, err)
	}

	ca, err := x509.ParseCertificate(caCert)
	if err != nil {
		return nil, fmt.Errorf("could not parse CA certificate: %w", err)
	}
	// Ensure only one level of CAs exist, and that they are created explicitly.
	req.Template.IsCA = false
	return issueCertificate(req, ca, c.PrivateKey)
}

// CACertificate returns the DER encoded x509 form of this Certificate that
// will be the used to issue child certificates.
func (c *Certificate) CACertificate(ctx context.Context, kv clientv3.KV) ([]byte, error) {
	return c.ensure(ctx, kv)
}
