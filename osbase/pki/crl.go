// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package pki

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"source.monogon.dev/osbase/event"
	"source.monogon.dev/osbase/event/etcd"
)

// crlPath returns the etcd path under which the marshaled X.509 Certificate
// Revocation List is stored.
//
// TODO(q3k): use etcd keyspace API from
func (c *Certificate) crlPath() string {
	return c.Namespace.etcdPath("%s-crl.der", c.Name)
}

// Revoke performs a CRL-based revocation of a given certificate by this CA,
// looking it up by DNS name. The revocation is immediately written to the
// backing etcd store and will be available to consumers through the WatchCRL
// API.
//
// An error is returned if the CRL could not be emitted (eg. due to an etcd
// communication error, a conflicting CRL write) or if the given hostname
// matches no emitted certificate.
//
// Only Managed and External certificates can be revoked.
func (c Certificate) Revoke(ctx context.Context, kv clientv3.KV, hostname string) error {
	crlPath := c.crlPath()
	issuedCerts := c.Namespace.etcdPath("issued/")

	res, err := kv.Txn(ctx).Then(
		clientv3.OpGet(crlPath),
		clientv3.OpGet(issuedCerts, clientv3.WithPrefix())).Commit()
	if err != nil {
		return fmt.Errorf("failed to retrieve certificates and CRL from etcd: %w", err)
	}

	// Parse certs, CRL and CRL revision from state.
	var certs []*x509.Certificate
	var crlRevision int64
	var crl *pkix.CertificateList
	for _, el := range res.Responses {
		for _, kv := range el.GetResponseRange().GetKvs() {
			if string(kv.Key) == crlPath {
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

	// Otherwise, revoke and save new CRL.
	revoked = append(revoked, pkix.RevokedCertificate{
		SerialNumber:   serial,
		RevocationTime: time.Now(),
	})

	crlRaw, err := c.makeCRL(ctx, kv, revoked)
	if err != nil {
		return fmt.Errorf("when generating new CRL for revocation: %w", err)
	}

	res, err = kv.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(crlPath), "=", crlRevision),
	).Then(
		clientv3.OpPut(crlPath, string(crlRaw)),
	).Commit()
	if err != nil {
		return fmt.Errorf("when saving new CRL: %w", err)
	}
	if !res.Succeeded {
		return fmt.Errorf("CRL save transaction failed, retry possible")
	}

	return nil
}

// makeCRL returns a valid CRL for a given list of certificates to be revoked.
// The given etcd client is used to ensure this CA certificate exists in etcd,
// but is not used to write any CRL to etcd.
func (c *Certificate) makeCRL(ctx context.Context, kv clientv3.KV, revoked []pkix.RevokedCertificate) ([]byte, error) {
	if c.Mode != CertificateManaged {
		return nil, fmt.Errorf("only managed certificates can issue CRLs")
	}
	certBytes, err := c.ensure(ctx, kv)
	if err != nil {
		return nil, fmt.Errorf("when ensuring certificate: %w", err)
	}
	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, fmt.Errorf("when parsing issuing certificate: %w", err)
	}
	crl, err := cert.CreateCRL(rand.Reader, c.PrivateKey, revoked, time.Now(), UnknownNotAfter)
	if err != nil {
		return nil, fmt.Errorf("failed to generate CRL: %w", err)
	}
	return crl, nil
}

// WatchCRL returns and Event Value compatible CRLWatcher which can be used to
// retrieve and watch for the newest CRL available from this CA certificate.
func (c *Certificate) WatchCRL(cl etcd.ThinClient) event.Watcher[*CRL] {
	value := etcd.NewValue(cl, c.crlPath(), func(_, data []byte) (*CRL, error) {
		crl, err := x509.ParseCRL(data)
		if err != nil {
			return nil, fmt.Errorf("could not parse CRL from etcd: %w", err)
		}
		return &CRL{
			Raw:  data,
			List: crl,
		}, nil
	})
	return value.Watch()
}

type CRL struct {
	Raw  []byte
	List *pkix.CertificateList
}
