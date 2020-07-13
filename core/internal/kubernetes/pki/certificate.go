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
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"net"

	"go.etcd.io/etcd/clientv3"
)

// Certificate is the promise of a Certificate being available to the caller. In this case, Certificate refers to a
// pair of x509 certificate and corresponding private key.
// Certificates can be stored in etcd, and their issuers might also be store on etcd. As such, this type's methods
// contain references to an etcd KV client.
// This Certificate type is agnostic to usage, but mostly geared towards Kubernetes certificates.
type Certificate struct {
	// issuer is the Issuer that will generate this certificate if one doesn't yet exist or etcd, or the requested
	// certificate is volatile (not to be stored on etcd).
	issuer Issuer
	// name is a unique key for storing the certificate in etcd. If empty, certificate is 'volatile', will not be stored
	// on etcd, and every .Ensure() call will generate a new pair.
	name string
	// template is an x509 certificate definition that will be used to generate the certificate when issuing it.
	template x509.Certificate
}

const (
	// etcdPrefix is where all the PKI data is stored in etcd.
	etcdPrefix = "/kube-pki/"
)

func etcdPath(f string, args ...interface{}) string {
	return etcdPrefix + fmt.Sprintf(f, args...)
}

// New creates a new Certificate, or to be more precise, a promise that a certificate will exist once Ensure is called.
// Issuer must be a valid certificate issuer (SelfSigned or another Certificate). Name must be unique among all
// certificates, or empty (which will cause the certificate to be volatile, ie. not stored in etcd).
func New(issuer Issuer, name string, template x509.Certificate) *Certificate {
	return &Certificate{
		issuer:   issuer,
		name:     name,
		template: template,
	}
}

// Client makes a Kubernetes PKI-compatible client certificate template.
// Directly derived from Kubernetes PKI requirements documented at
// https://kubernetes.io/docs/setup/best-practices/certificates/#configure-certificates-manually
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

func (c *Certificate) etcdPaths() (cert, key string) {
	return etcdPath("%s-cert.der", c.name), etcdPath("%s-key.der", c.name)
}

// ensure returns a DER-encoded x509 certificate and internally encoded bare ed25519 key for a given Certificate,
// in memory (if volatile), loading it from etcd, or creating and saving it on etcd if needed.
// This function is safe to call in parallel from multiple etcd clients (including across machines), but it will error
// in case a concurrent certificate generation happens. These errors are, however, safe to retry - as long as all the
// certificate creators (ie., Smalltown nodes) run the same version of this code.
// TODO(q3k): in the future, this should be handled better - especially as we introduce new certificates, or worse,
// change the issuance chain. As a stopgap measure, an explicit per-certificate or even global lock can be implemented.
// And, even before that, we can handle concurrency errors in a smarter way.
func (c *Certificate) ensure(ctx context.Context, kv clientv3.KV) (cert, key []byte, err error) {
	if c.name == "" {
		// Volatile certificate - generate.
		// TODO(q3k): cache internally?
		cert, key, err = c.issuer.Issue(ctx, c.template, kv)
		if err != nil {
			err = fmt.Errorf("failed to issue: %w", err)
			return
		}
		return
	}

	certPath, keyPath := c.etcdPaths()

	// Try loading certificate and key from etcd.
	certRes, err := kv.Get(ctx, certPath)
	if err != nil {
		err = fmt.Errorf("failed to get certificate from etcd: %w", err)
		return
	}
	keyRes, err := kv.Get(ctx, keyPath)
	if err != nil {
		err = fmt.Errorf("failed to get key from etcd: %w", err)
		return
	}

	if len(certRes.Kvs) == 1 && len(keyRes.Kvs) == 1 {
		// Certificate and key exists in etcd, return that.
		cert = certRes.Kvs[0].Value
		key = keyRes.Kvs[0].Value

		err = nil
		// TODO(q3k): check for expiration
		return
	}

	// No certificate found - issue one.
	cert, key, err = c.issuer.Issue(ctx, c.template, kv)
	if err != nil {
		err = fmt.Errorf("failed to issue: %w", err)
		return
	}

	// Save to etcd in transaction. This ensures that no partial writes happen, and that we haven't been raced to the
	// save.
	res, err := kv.Txn(ctx).
		If(
			clientv3.Compare(clientv3.CreateRevision(certPath), "=", 0),
			clientv3.Compare(clientv3.CreateRevision(keyPath), "=", 0),
		).
		Then(
			clientv3.OpPut(certPath, string(cert)),
			clientv3.OpPut(keyPath, string(key)),
		).Commit()
	if err != nil {
		err = fmt.Errorf("failed to write newly issued certificate: %w", err)
	} else if !res.Succeeded {
		err = fmt.Errorf("certificate issuance transaction failed: concurrent write")
	}

	return
}

// Ensure returns an x509 DER-encoded (but not PEM-encoded) certificate and key for a given Certificate.
// If the certificate is volatile, each call to Ensure will cause a new certificate to be generated.
// Otherwise, it will be retrieved from etcd, or generated and stored there if needed.
func (c *Certificate) Ensure(ctx context.Context, kv clientv3.KV) (cert, key []byte, err error) {
	cert, key, err = c.ensure(ctx, kv)
	if err != nil {
		return nil, nil, err
	}
	key, err = x509.MarshalPKCS8PrivateKey(ed25519.PrivateKey(key))
	if err != nil {
		err = fmt.Errorf("could not marshal private key (data corruption?): %w", err)
		return
	}
	return cert, key, err
}
