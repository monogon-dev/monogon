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

// package pki implements an x509 PKI (Public Key Infrastructure) system backed
// on etcd.
package pki

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"net"

	clientv3 "go.etcd.io/etcd/client/v3"

	"source.monogon.dev/osbase/fileargs"
)

// Namespace represents some path in etcd where certificate/CA data will be
// stored. Creating a namespace via Namespaced then permits the consumer of
// this library to start creating certificates within this namespace.
type Namespace struct {
	prefix string
}

// Namespaced creates a namespace for storing certificate data in etcd at a
// given 'path' prefix.
func Namespaced(prefix string) Namespace {
	return Namespace{
		prefix: prefix,
	}
}

type CertificateMode int

const (
	// CertificateManaged is a certificate whose key material is fully managed by
	// the Certificate code. When set, PublicKey and PrivateKey must not be set by
	// the user, and instead will be populated by the Ensure call. Name must be set,
	// and will be used to store this Certificate and its keys within etcd. After
	// the initial generation during Ensure, other Certificates with the same Name
	// will be retrieved (including key material) from etcd.
	CertificateManaged CertificateMode = iota

	// CertificateExternal is a certificate whose key material is not managed by
	// Certificate or stored in etcd, but the X509 certificate itself is. PublicKey
	// must be set while PrivateKey must not be set. Name must be set, and will be
	// used to store the emitted X509 certificate in etcd on Ensure. After the
	// initial generation during Ensure, other Certificates with the same Name will
	// be retrieved (without key material) from etcd.
	CertificateExternal

	// CertificateEphemeral is a certificate whose data (X509 certificate and
	// possibly key material) is generated on demand each time Ensure is called.
	// Nothing is stored in etcd or loaded from etcd. PrivateKey or PublicKey can be
	// set, if both are nil then a new keypair will be generated. Name is ignored.
	CertificateEphemeral
)

// Certificate is the promise of a Certificate being available to the caller.
// In this case, Certificate refers to a pair of x509 certificate and
// corresponding private key.  Certificates can be stored in etcd, and their
// issuers might also be store on etcd. As such, this type's methods contain
// references to an etcd KV client.
type Certificate struct {
	Namespace *Namespace

	// Issuer is the Issuer that will generate this certificate if one doesn't
	// yet exist or etcd, or the requested certificate is ephemeral (not to be
	// stored on etcd).
	Issuer Issuer
	// Name is a unique key for storing the certificate in etcd (if the requested
	// certificate is not ephemeral).
	Name string
	// Template is an x509 certificate definition that will be used to generate
	// the certificate when issuing it.
	Template x509.Certificate

	// Mode in which this Certificate will operate. This influences the behaviour of
	// the Ensure call.
	Mode CertificateMode

	// PrivateKey is the private key for this Certificate. It should never be set by
	// the user, and instead will be populated by the Ensure call for Managed
	// Certificates.
	PrivateKey ed25519.PrivateKey

	// PublicKey is the public key for this Certificate. It should only be set by
	// the user for External or Ephemeral certificates, and will be populated by the
	// next Ensure call if missing.
	PublicKey ed25519.PublicKey
}

func (n *Namespace) etcdPath(f string, args ...interface{}) string {
	return n.prefix + fmt.Sprintf(f, args...)
}

// Client makes a Kubernetes PKI-compatible client certificate template.
// Directly derived from Kubernetes PKI requirements documented at
//   https://kubernetes.io/docs/setup/best-practices/certificates/#configure-certificates-manually
func Client(identity string, groups []string) x509.Certificate {
	return x509.Certificate{
		Subject: pkix.Name{
			CommonName:   identity,
			Organization: groups,
		},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
}

// Server makes a Kubernetes PKI-compatible server certificate template.
func Server(dnsNames []string, ips []net.IP) x509.Certificate {
	return x509.Certificate{
		Subject:     pkix.Name{},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    dnsNames,
		IPAddresses: ips,
	}
}

// CA makes a Certificate that can sign other certificates.
func CA(cn string) x509.Certificate {
	return x509.Certificate{
		Subject: pkix.Name{
			CommonName: cn,
		},
		IsCA:        true,
		KeyUsage:    x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageOCSPSigning},
	}
}

// ensure returns a DER-encoded x509 certificate and internally encoded bare
// ed25519 key for a given Certificate, in memory (if ephemeral), loading it
// from etcd, or creating and saving it on etcd if needed.
// This function is safe to call in parallel from multiple etcd clients
// (including across machines), but it will error in case a concurrent
// certificate generation happens. These errors are, however, safe to retry -
// as long as all the certificate creators (ie., Metropolis nodes) run the same
// version of this code.
func (c *Certificate) ensure(ctx context.Context, kv clientv3.KV) (cert []byte, err error) {
	// Ensure key is available.
	if err := c.ensureKey(ctx, kv); err != nil {
		return nil, err
	}

	switch c.Mode {
	case CertificateEphemeral:
		// TODO(q3k): cache internally?
		cert, err = c.Issuer.Issue(ctx, c, kv)
		if err != nil {
			return nil, fmt.Errorf("failed to issue: %w", err)
		}
		return cert, nil
	case CertificateManaged, CertificateExternal:
	default:
		return nil, fmt.Errorf("invalid certificate mode %v", c.Mode)
	}

	if c.Name == "" {
		if c.Mode == CertificateExternal {
			return nil, fmt.Errorf("external certificate must have name set")
		} else {
			return nil, fmt.Errorf("managed certificate must have name set")
		}
	}

	certPath := c.Namespace.etcdPath("issued/%s-cert.der", c.Name)

	// Try loading certificate from etcd.
	certRes, err := kv.Get(ctx, certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get certificate from etcd: %w", err)
	}

	if len(certRes.Kvs) == 1 {
		certBytes := certRes.Kvs[0].Value
		cert, err := x509.ParseCertificate(certBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate retrieved from etcd: %w", err)
		}
		pk, ok := cert.PublicKey.(ed25519.PublicKey)
		if !ok {
			return nil, fmt.Errorf("unexpected non-ed25519 certificate found in etcd")
		}
		if !bytes.Equal(pk, c.PublicKey) {
			return nil, fmt.Errorf("certificate stored in etcd emitted for different public key")
		}
		// TODO(q3k): ensure issuer and template haven't changed
		return certBytes, nil
	}

	// No certificate found - issue one and save to etcd.
	cert, err = c.Issuer.Issue(ctx, c, kv)
	if err != nil {
		return nil, fmt.Errorf("failed to issue: %w", err)
	}

	res, err := kv.Txn(ctx).
		If(
			clientv3.Compare(clientv3.CreateRevision(certPath), "=", 0),
		).
		Then(
			clientv3.OpPut(certPath, string(cert)),
		).Commit()
	if err != nil {
		err = fmt.Errorf("failed to write newly issued certificate: %w", err)
	} else if !res.Succeeded {
		err = fmt.Errorf("certificate issuance transaction failed: concurrent write")
	}

	return
}

// ensureKey retrieves or creates PublicKey as needed based on the Certificate
// Mode. For Managed Certificates and Ephemeral Certificates with no PrivateKey
// it will also populate PrivateKay.
func (c *Certificate) ensureKey(ctx context.Context, kv clientv3.KV) error {
	// If we have a public key then we're all set.
	if c.PublicKey != nil {
		return nil
	}

	// For ephemeral keys, we just generate them.
	// For external keys, we can't do anything - not having the keys set means
	// a programming error.

	switch c.Mode {
	case CertificateEphemeral:
		pub, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return fmt.Errorf("when generating ephemeral key: %w", err)
		}
		c.PublicKey = pub
		c.PrivateKey = priv
		return nil
	case CertificateExternal:
		if c.PrivateKey != nil {
			// We prohibit having PrivateKey set in External Certificates to simplify the
			// different logic paths this library implements. Being able to assume External
			// == PublicKey only makes things easier elsewhere.
			return fmt.Errorf("external certificate must not have PrivateKey set")
		}
		return fmt.Errorf("external certificate must have PublicKey set")
	case CertificateManaged:
	default:
		return fmt.Errorf("invalid certificate mode %v", c.Mode)
	}

	// For managed keys, synchronize with etcd.
	if c.Name == "" {
		return fmt.Errorf("managed certificate must have Name set")
	}

	// First, try loading.
	privPath := c.Namespace.etcdPath("keys/%s-privkey.bin", c.Name)
	privRes, err := kv.Get(ctx, privPath)
	if err != nil {
		return fmt.Errorf("failed to get private key from etcd: %w", err)
	}
	if len(privRes.Kvs) == 1 {
		privBytes := privRes.Kvs[0].Value
		if len(privBytes) != ed25519.PrivateKeySize {
			return fmt.Errorf("stored private key has invalid size")
		}
		c.PrivateKey = privBytes
		c.PublicKey = c.PrivateKey.Public().(ed25519.PublicKey)
		return nil
	}

	// No key in etcd? Generate and save.
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("while generating keypair: %w", err)
	}

	res, err := kv.Txn(ctx).
		If(
			clientv3.Compare(clientv3.CreateRevision(privPath), "=", 0),
		).
		Then(
			clientv3.OpPut(privPath, string(priv)),
		).Commit()
	if err != nil {
		return fmt.Errorf("failed to write newly generated keypair: %w", err)
	} else if !res.Succeeded {
		return fmt.Errorf("key generation transaction failed: concurrent write")
	}

	crlPath := c.crlPath()
	emptyCRL, err := c.makeCRL(ctx, kv, nil)
	if err != nil {
		return fmt.Errorf("failed to generate empty CRL: %w", err)
	}

	// Also attempt to emit an empty CRL if one doesn't exist yet.
	_, err = kv.Txn(ctx).
		If(
			clientv3.Compare(clientv3.CreateRevision(crlPath), "=", 0),
		).
		Then(
			clientv3.OpPut(crlPath, string(emptyCRL)),
		).Commit()
	if err != nil {
		return fmt.Errorf("failed to upsert empty CRL")
	}

	c.PrivateKey = priv
	c.PublicKey = pub
	return nil
}

// Ensure returns an x509 DER-encoded (but not PEM-encoded) certificate for a
// given Certificate.
//
// If the Certificate is ephemeral, each call to Ensure will cause a new
// certificate to be generated. Otherwise, it will be retrieved from etcd, or
// generated and stored there if needed.
func (c *Certificate) Ensure(ctx context.Context, kv clientv3.KV) (cert []byte, err error) {
	return c.ensure(ctx, kv)
}

func (c *Certificate) PrivateKeyX509() ([]byte, error) {
	if c.PrivateKey == nil {
		return nil, fmt.Errorf("certificate has no private key")
	}
	key, err := x509.MarshalPKCS8PrivateKey(c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("could not marshal private key (data corruption?): %w", err)
	}
	return key, nil
}

// FilesystemCertificate is a fileargs.FileArgs wrapper which will contain PEM
// encoded certificate material when Mounted. This construct is useful when
// dealing with services that want to access etcd-backed certificates as files
// available locally.
// Paths to the available files are considered opaque and should not be leaked
// outside of the struct. Further restrictions on access to these files might
// be imposed in the future.
type FilesystemCertificate struct {
	*fileargs.FileArgs
	// CACertPath is the full path at which the CA certificate is available.
	// Read only.
	CACertPath string
	// CertPath is the full path at which the certificate is available. Read
	// only.
	CertPath string
	// KeyPath is the full path at which the private key is available, or an empty
	// string if the Certificate was created without a private key. Read only.
	KeyPath string
}

// Mount returns a locally mounted FilesystemCertificate for this Certificate,
// which allows services to access this Certificate via local filesystem
// access.
// The embeded fileargs.FileArgs can also be used to add additional file-backed
// data under the same mount by calling ArgPath.
// The returned FilesystemCertificate must be Closed in order to prevent a
// system mount leak.
func (c *Certificate) Mount(ctx context.Context, kv clientv3.KV) (*FilesystemCertificate, error) {
	fa, err := fileargs.New()
	if err != nil {
		return nil, fmt.Errorf("when creating fileargs mount: %w", err)
	}
	fs := &FilesystemCertificate{FileArgs: fa}

	cert, err := c.Ensure(ctx, kv)
	if err != nil {
		return nil, fmt.Errorf("when issuing certificate: %w", err)
	}

	cacert, err := c.Issuer.CACertificate(ctx, kv)
	if err != nil {
		return nil, fmt.Errorf("when getting issuer CA: %w", err)
	}
	// cacert will be null if this is a self-signed certificate.
	if cacert == nil {
		cacert = cert
	}

	fs.CACertPath = fs.ArgPath("ca.crt", pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cacert}))
	fs.CertPath = fs.ArgPath("tls.crt", pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert}))
	if c.PrivateKey != nil {
		key, err := c.PrivateKeyX509()
		if err != nil {
			return nil, err
		}
		fs.KeyPath = fs.ArgPath("tls.key", pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key}))
	}

	return fs, nil
}
