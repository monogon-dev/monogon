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
// The CA is backed by etcd storage, and can also bootstrap itself without a
// yet running etcd storage (and commit in-memory secrets to etcd at a later
// date).
//
// This is different from //metropolis/pkg/pki in that it has to solve the
// certs-for-etcd-on-etcd bootstrap problem. Perhaps it should be rewritten to
// implement the Issuer/Ceritifcate interface available there.
//
// CA and certificates successfully pass https://github.com/zmap/zlint
// (minus the CA/B rules that a public CA would adhere to, which requires
// things like OCSP servers, Certificate Policies and ECDSA/RSA-only keys).
package ca

// TODO(leo): add zlint test

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const (
	// TODO(q3k): move this to a declarative storage layer
	pathCACertificate      = "/etcd-ca/ca.der"
	pathCAKey              = "/etcd-ca/ca-key.der"
	pathCACRL              = "/etcd-ca/crl.der"
	pathIssuedCertificates = "/etcd-ca/certs/"
)

func pathIssuedCertificate(serial *big.Int) string {
	return pathIssuedCertificates + hex.EncodeToString(serial.Bytes())
}

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

type CA struct {
	// TODO: Potentially protect the key with memguard
	privateKey *ed25519.PrivateKey
	CACert     *x509.Certificate
	CACertRaw  []byte

	// bootstrapIssued are certificates that have been issued by the CA before
	// it has been successfully Saved to etcd.
	bootstrapIssued [][]byte
	// canBootstrapIssue is set on CAs that have been created by New and not
	// yet stored to etcd. If not set, certificates cannot be issued in-memory.
	canBootstrapIssue bool
}

// Workaround for https://github.com/golang/go/issues/26676 in Go's
// crypto/x509. Specifically Go violates Section 4.2.1.2 of RFC 5280 without
// this.
// Fixed for 1.15 in https://go-review.googlesource.com/c/go/+/227098/.
//
// Taken from https://github.com/FiloSottile/mkcert/blob/master/cert.go#L295
// written by one of Go's crypto engineers (BSD 3-clause).
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

// New creates a new certificate authority with the given common name. The
// newly created CA will be stored in memory until committed to etcd by calling
// .Save.
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
		privateKey: &privKey,
		CACertRaw:  caCertRaw,
		CACert:     caCert,

		canBootstrapIssue: true,
	}

	return ca, nil
}

// Load restores CA state from etcd.
func Load(ctx context.Context, kv clientv3.KV) (*CA, error) {
	resp, err := kv.Txn(ctx).Then(
		clientv3.OpGet(pathCACertificate),
		clientv3.OpGet(pathCAKey),
		// We only read the CRL to ensure it exists on etcd (and early fail on
		// inconsistency)
		clientv3.OpGet(pathCACRL)).Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve CA from etcd: %w", err)
	}

	var caCert, caKey, caCRL []byte
	for _, el := range resp.Responses {
		for _, kv := range el.GetResponseRange().GetKvs() {
			switch string(kv.Key) {
			case pathCACertificate:
				caCert = kv.Value
			case pathCAKey:
				caKey = kv.Value
			case pathCACRL:
				caCRL = kv.Value
			}
		}
	}
	if caCert == nil || caKey == nil || caCRL == nil {
		return nil, fmt.Errorf("failed to retrieve CA from etcd, missing at least one of {ca key, ca crt, ca crl}")
	}

	if len(caKey) != ed25519.PrivateKeySize {
		return nil, errors.New("invalid CA private key size")
	}
	privateKey := ed25519.PrivateKey(caKey)

	caCertVal, err := x509.ParseCertificate(caCert)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CA certificate: %w", err)
	}
	return &CA{
		privateKey: &privateKey,
		CACertRaw:  caCert,
		CACert:     caCertVal,
	}, nil
}

// Save stores a newly created CA into etcd, committing both the CA data and
// any certificates issued until then.
func (c *CA) Save(ctx context.Context, kv clientv3.KV) error {
	crl, err := c.makeCRL(nil)
	if err != nil {
		return fmt.Errorf("failed to generate initial CRL: %w", err)
	}

	ops := []clientv3.Op{
		clientv3.OpPut(pathCACertificate, string(c.CACertRaw)),
		clientv3.OpPut(pathCAKey, string([]byte(*c.privateKey))),
		clientv3.OpPut(pathCACRL, string(crl)),
	}
	for i, certRaw := range c.bootstrapIssued {
		cert, err := x509.ParseCertificate(certRaw)
		if err != nil {
			return fmt.Errorf("failed to parse in-memory certificate %d", i)
		}
		ops = append(ops, clientv3.OpPut(pathIssuedCertificate(cert.SerialNumber), string(certRaw)))
	}

	res, err := kv.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(pathCAKey), "=", 0),
	).Then(ops...).Commit()
	if err != nil {
		return fmt.Errorf("failed to store CA to etcd: %w", err)
	}
	if !res.Succeeded {
		// This should pretty much never happen, but we want to catch it just in case.
		return fmt.Errorf("failed to store CA to etcd: CA already present - cluster-level data inconsistency")
	}
	c.bootstrapIssued = nil
	c.canBootstrapIssue = false
	return nil
}

// Issue issues a certificate. If kv is non-nil, the newly issued certificate
// will be immediately stored to etcd, otherwise it will be kept in memory
// (until .Save is called). Certificates can only be issued to memory on
// newly-created CAs that have not been saved to etcd yet.
func (c *CA) Issue(ctx context.Context, kv clientv3.KV, commonName string) (cert []byte, privkey []byte, err error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("failed to generate serial number: %w", err)
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
	cert, err = x509.CreateCertificate(rand.Reader, etcdCert, c.CACert, pubKey, c.privateKey)
	if err != nil {
		err = fmt.Errorf("failed to sign new certificate: %w", err)
		return
	}

	if kv != nil {
		path := pathIssuedCertificate(serialNumber)
		_, err = kv.Put(ctx, path, string(cert))
		if err != nil {
			err = fmt.Errorf("failed to commit new certificate to etcd: %w", err)
			return
		}
	} else {
		if !c.canBootstrapIssue {
			err = fmt.Errorf("cannot issue new certificate to memory on existing, etcd-backed CA")
			return
		}
		c.bootstrapIssued = append(c.bootstrapIssued, cert)
	}
	return
}

func (c *CA) makeCRL(revoked []pkix.RevokedCertificate) ([]byte, error) {
	crl, err := c.CACert.CreateCRL(rand.Reader, c.privateKey, revoked, time.Now(), unknownNotAfter)
	if err != nil {
		return nil, fmt.Errorf("failed to generate CRL: %w", err)
	}
	return crl, nil
}

// Revoke revokes a certificate by hostname. The selected hostname will be
// added to the CRL stored in etcd. This call might fail (safely) if a
// simultaneous revoke happened that caused the CRL to be bumped. The call can
// be then retried safely.
func (c *CA) Revoke(ctx context.Context, kv clientv3.KV, hostname string) error {
	res, err := kv.Txn(ctx).Then(
		clientv3.OpGet(pathCACRL),
		clientv3.OpGet(pathIssuedCertificates, clientv3.WithPrefix())).Commit()
	if err != nil {
		return fmt.Errorf("failed to retrieve certificates and CRL from etcd: %w", err)
	}

	var certs []*x509.Certificate
	var crlRevision int64
	var crl *pkix.CertificateList
	for _, el := range res.Responses {
		for _, kv := range el.GetResponseRange().GetKvs() {
			if string(kv.Key) == pathCACRL {
				crl, err = x509.ParseCRL(kv.Value)
				if err != nil {
					return fmt.Errorf("could not parse CRL from etcd: %w", err)
				}
				crlRevision = kv.CreateRevision
			} else {
				cert, err := x509.ParseCertificate(kv.Value)
				if err != nil {
					return fmt.Errorf("could not parse certificate %q from etcd: %w", string(kv.Key), err)
				}
				certs = append(certs, cert)
			}
		}
	}

	if crl == nil {
		return fmt.Errorf("could not find CRL in etcd")
	}
	revoked := crl.TBSCertList.RevokedCertificates

	// Find requested hostname in issued certificates.
	var serial *big.Int
	for _, cert := range certs {
		for _, dnsName := range cert.DNSNames {
			if dnsName == hostname {
				serial = cert.SerialNumber
				break
			}
		}
		if serial != nil {
			break
		}
	}
	if serial == nil {
		return fmt.Errorf("could not find requested hostname")
	}

	// Check if certificate has already been revoked.
	for _, revokedCert := range revoked {
		if revokedCert.SerialNumber.Cmp(serial) == 0 {
			return nil // Already revoked
		}
	}

	revoked = append(revoked, pkix.RevokedCertificate{
		SerialNumber:   serial,
		RevocationTime: time.Now(),
	})

	crlRaw, err := c.makeCRL(revoked)
	if err != nil {
		return fmt.Errorf("when generating new CRL for revocation: %w", err)
	}

	res, err = kv.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(pathCACRL), "=", crlRevision),
	).Then(
		clientv3.OpPut(pathCACRL, string(crlRaw)),
	).Commit()
	if err != nil {
		return fmt.Errorf("when saving new CRL: %w", err)
	}
	if !res.Succeeded {
		return fmt.Errorf("CRL save transaction failed, retry possibly")
	}

	return nil
}

// WaitCRLChange returns a channel that will receive a CRLUpdate any time the
// remote CRL changed. Immediately after calling this method, the current CRL
// is retrieved from the cluster and put into the channel.
func (c *CA) WaitCRLChange(ctx context.Context, kv clientv3.KV, w clientv3.Watcher) <-chan CRLUpdate {
	C := make(chan CRLUpdate)

	go func(ctx context.Context) {
		ctxC, cancel := context.WithCancel(ctx)
		defer cancel()

		fail := func(f string, args ...interface{}) {
			C <- CRLUpdate{Err: fmt.Errorf(f, args...)}
			close(C)
		}

		initial, err := kv.Get(ctx, pathCACRL)
		if err != nil {
			fail("failed to retrieve initial CRL: %w", err)
			return
		}

		C <- CRLUpdate{CRL: initial.Kvs[0].Value}

		for wr := range w.Watch(ctxC, pathCACRL, clientv3.WithRev(initial.Kvs[0].CreateRevision)) {
			if wr.Err() != nil {
				fail("failed watching CRL: %w", wr.Err())
				return
			}

			for _, e := range wr.Events {
				if string(e.Kv.Key) != pathCACRL {
					continue
				}

				C <- CRLUpdate{CRL: e.Kv.Value}
			}
		}
	}(ctx)

	return C
}

// CRLUpdate is emitted for every remote CRL change, and spuriously on ever new
// WaitCRLChange.
type CRLUpdate struct {
	// The new (or existing, in the case of the first call) CRL. If nil, Err
	// will be set.
	CRL []byte
	// If set, an error occurred and the WaitCRLChange call must be restarted.
	// If set, CRL will be nil.
	Err error
}

// GetCurrentCRL returns the current CRL for the CA. This should only be used
// for one-shot operations like bootstrapping a new node that doesn't yet have
// access to etcd - otherwise, WaitCRLChange shoulde be used.
func (c *CA) GetCurrentCRL(ctx context.Context, kv clientv3.KV) ([]byte, error) {
	initial, err := kv.Get(ctx, pathCACRL)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve initial CRL: %w", err)
	}
	return initial.Kvs[0].Value, nil
}
