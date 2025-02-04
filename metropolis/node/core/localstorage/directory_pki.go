// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package localstorage

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
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
			return false, fmt.Errorf("failed to check %q: %w", d.FullPath(), err)
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
			return false, fmt.Errorf("failed to check %q: %w", d.FullPath(), err)
		}
		if exists {
			return false, nil
		}
	}
	return true, nil
}

// GeneratePrivateKey will generate an ED25519 private key for this PKIDirectory
// if it doesn't yet exist.
func (p *PKIDirectory) GeneratePrivateKey() error {
	// Do nothing if key already exists.
	_, err := p.Key.Read()
	if err == nil {
		return nil
	}
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	return p.WritePrivateKey(priv)
}

// WritePrivateKey serializes the given private key (PKCS8 + PEM) and writes it
// to the PKIDirectory, overwriting whatever might already be present there.
func (p *PKIDirectory) WritePrivateKey(key ed25519.PrivateKey) error {
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
	return nil
}

// ReadPrivateKey loads an ED25519 private key from the PKIDirectory and
// deserializes it (PEM + PKCS).
func (p *PKIDirectory) ReadPrivateKey() (ed25519.PrivateKey, error) {
	raw, err := p.Key.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read key: %w", err)
	}
	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, errors.New("not PEM")
	}
	keyType := "PRIVATE KEY"
	if block.Type != keyType {
		return nil, fmt.Errorf("contains a PEM block that's not a %v", keyType)
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse key: %w", err)
	}
	switch k := key.(type) {
	case ed25519.PrivateKey:
		return k, nil
	default:
		return nil, fmt.Errorf("PCKS8 contains invalid key type")
	}
}

// WriteCertificates serializes (PEM) and saves the given certificates into the
// PKIDirectory, overwriting whatever might already be present there.
//
// For ease of use, the accepted certificates are expected to already be in
// DER-encoded form (eg. from the Raw field in a x509.Certificate).
func (p *PKIDirectory) WriteCertificates(ca, cert []byte) error {
	if err := p.MkdirAll(0700); err != nil {
		return fmt.Errorf("failed to make PKI directory: %w", err)
	}
	if err := p.Certificate.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert}), 0600); err != nil {
		return fmt.Errorf("failed to write certificate: %w", err)
	}
	if err := p.CACertificate.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ca}), 0600); err != nil {
		return fmt.Errorf("failed to write CA certificate: %w", err)
	}
	return nil
}

// readPEMX509 reads a file and parses it as a PEM-encoded X509 certificate.
func readPEMX509(p *declarative.File) (*x509.Certificate, error) {
	bytes, err := p.Read()
	if err != nil {
		return nil, fmt.Errorf("couldn't read: %w", err)
	}
	block, _ := pem.Decode(bytes)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, errors.New("invalid PEM armoring")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("invalid X509: %w", err)
	}
	return cert, nil
}

// ReadCertificates reads and parses (PEM + X509) the certificates from a
// PKIDirectory.
func (p *PKIDirectory) ReadCertificates() (ca, cert *x509.Certificate, err error) {
	ca, err = readPEMX509(&p.CACertificate)
	if err != nil {
		return nil, nil, fmt.Errorf("when loading CA certificate: %w", err)
	}
	cert, err = readPEMX509(&p.Certificate)
	if err != nil {
		return nil, nil, fmt.Errorf("when loading certificate: %w", err)
	}
	return ca, cert, nil
}

// WriteAll (over)writes the PKI data in this directory with the given private
// key, certificate and CA certificate.
//
// For ease of use, the accepted certificates are expected to already be in
// DER-encoded form (eg. from the Raw field in a x509.Certificate).
func (p *PKIDirectory) WriteAll(cert []byte, key ed25519.PrivateKey, ca []byte) error {
	if err := p.WritePrivateKey(key); err != nil {
		return err
	}
	if err := p.WriteCertificates(ca, cert); err != nil {
		return err
	}
	return nil
}

// ReadAll reads and parses (PEM + PKCS8/X509) the stored certificates and key of
// this PKIDirectory.
func (p *PKIDirectory) ReadAll() (ca, cert *x509.Certificate, key ed25519.PrivateKey, err error) {
	ca, cert, err = p.ReadCertificates()
	if err != nil {
		return
	}
	key, err = p.ReadPrivateKey()
	return
}
