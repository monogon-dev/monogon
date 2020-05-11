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

// package ca implements a simple standards-compliant certificate authority.
// It only supports ed25519 keys, and does not maintain any persistent state.
//
// CA and certificates successfully pass https://github.com/zmap/zlint
// (minus the CA/B rules that a public CA would adhere to, which requires
// things like OCSP servers, Certificate Policies and ECDSA/RSA-only keys).
package ca

// TODO(leo): add zlint test

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
	"time"
)

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

type CA struct {
	// TODO: Potentially protect the key with memguard
	PrivateKey *ed25519.PrivateKey
	CACert     *x509.Certificate
	CACertRaw  []byte
	CRLRaw     []byte
	Revoked    []pkix.RevokedCertificate
}

// Workaround for https://github.com/golang/go/issues/26676 in Go's crypto/x509. Specifically Go
// violates Section 4.2.1.2 of RFC 5280 without this. Should eventually be redundant.
//
// Taken from https://github.com/FiloSottile/mkcert/blob/master/cert.go#L295 written by one of Go's
// crypto engineers (BSD 3-clause).
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

// New creates a new certificate authority with the given common name.
func New(name string) (*CA, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %w", err)
	}

	skid, err := calculateSKID(pubKey)
	if err != nil {
		return nil, err
	}

	caCert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: name,
		},
		IsCA:                  true,
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              unknownNotAfter,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageOCSPSigning},
		AuthorityKeyId:        skid,
		SubjectKeyId:          skid,
	}

	caCertRaw, err := x509.CreateCertificate(rand.Reader, caCert, caCert, pubKey, privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create root certificate: %w", err)
	}

	ca := &CA{
		PrivateKey: &privKey,
		CACertRaw:  caCertRaw,
		CACert:     caCert,
	}
	if ca.ReissueCRL() != nil {
		return nil, fmt.Errorf("failed to create initial CRL: %w", err)
	}

	return ca, nil
}

// FromCertificates restores CA state.
func FromCertificates(caCert []byte, caKey []byte, crl []byte) (*CA, error) {
	if len(caKey) != ed25519.PrivateKeySize {
		return nil, errors.New("invalid CA private key size")
	}
	privateKey := ed25519.PrivateKey(caKey)

	caCertVal, err := x509.ParseCertificate(caCert)
	if err != nil {
		return nil, err
	}
	crlVal, err := x509.ParseCRL(crl)
	if err != nil {
		return nil, err
	}
	return &CA{
		PrivateKey: &privateKey,
		CACertRaw:  caCert,
		CACert:     caCertVal,
		Revoked:    crlVal.TBSCertList.RevokedCertificates,
	}, nil
}

// IssueCertificate issues a certificate
func (ca *CA) IssueCertificate(commonName string) (cert []byte, privkey []byte, err error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("Failed to generate serial number: %w", err)
		return
	}

	pubKey, privKeyRaw, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	privkey, err = x509.MarshalPKCS8PrivateKey(privKeyRaw)
	if err != nil {
		return
	}

	etcdCert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:         commonName,
			OrganizationalUnit: []string{"etcd"},
		},
		IsCA:                  false,
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              unknownNotAfter,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		DNSNames:              []string{commonName},
	}
	cert, err = x509.CreateCertificate(rand.Reader, etcdCert, ca.CACert, pubKey, ca.PrivateKey)
	return
}

func (ca *CA) ReissueCRL() error {
	newCRL, err := ca.CACert.CreateCRL(rand.Reader, ca.PrivateKey, ca.Revoked, time.Now(), unknownNotAfter)
	if err != nil {
		return err
	}
	ca.CRLRaw = newCRL
	return nil
}

func (ca *CA) Revoke(serial *big.Int) error {
	for _, revokedCert := range ca.Revoked {
		if revokedCert.SerialNumber.Cmp(serial) == 0 {
			return nil // Already revoked
		}
	}
	ca.Revoked = append(ca.Revoked, pkix.RevokedCertificate{
		SerialNumber:   serial,
		RevocationTime: time.Now(),
	})
	return ca.ReissueCRL()
}
