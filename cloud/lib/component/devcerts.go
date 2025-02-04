// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"k8s.io/klog/v2"

	"source.monogon.dev/osbase/pki"
)

// GetDevCerts returns paths to this component's development certificate, key
// and CA, or panics if unavailable.
func (c *ComponentConfig) GetDevCerts() (certPath, keyPath, caPath string) {
	klog.Infof("Using developer certificates at %s", c.DevCertsPath)

	caPath = c.ensureDevCA()
	certPath, keyPath = c.ensureDevComponent()
	return
}

// ensureDevComponent ensures that a development certificate/key exists for this
// component and returns paths to them. This data is either read from disk if it
// already exists, or is generated when this function is called. If any problem
// occurs, the code panics.
func (c *ComponentConfig) ensureDevComponent() (certPath, keyPath string) {
	caKeyPath := c.DevCertsPath + "/ca.key"
	caCertPath := c.DevCertsPath + "/ca.cert"

	// Load CA. By convention, we are always called after ensureDevCA.
	ca, err := tls.LoadX509KeyPair(caCertPath, caKeyPath)
	if err != nil {
		klog.Exitf("Could not load Dev CA: %v", err)
	}
	caCert, err := x509.ParseCertificate(ca.Certificate[0])
	if err != nil {
		klog.Exitf("Could not parse Dev CA: %v", err)
	}

	// Check if we have keys already.
	keyPath = c.DevCertsPath + fmt.Sprintf("/%s.key", c.ComponentName)
	certPath = c.DevCertsPath + fmt.Sprintf("/%s.crt", c.ComponentName)
	noKey := false
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		noKey = true
	}
	noCert := false
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		noCert = true
	}

	if noKey || noCert {
		klog.Infof("Generating developer %s certificate...", c.ComponentName)
	} else {
		return
	}

	// Generate key/certificate.
	cert := pki.Server([]string{
		fmt.Sprintf("%s.local", c.ComponentName),
	}, nil)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		klog.Exitf("Failed to generate %s serial number: %v", c.ComponentName, err)
	}
	cert.ExtKeyUsage = append(cert.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
	cert.SerialNumber = serialNumber
	cert.NotBefore = time.Now()
	cert.NotAfter = pki.UnknownNotAfter
	cert.BasicConstraintsValid = true

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		klog.Exitf("Failed to generate %s key: %v", c.ComponentName, err)
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &cert, caCert, pub, ca.PrivateKey)
	if err != nil {
		klog.Exitf("Failed to generate %s certificate: %v", c.ComponentName, err)
	}

	// And marshal them to disk.
	privPKCS, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		klog.Exitf("Failed to marshal %s private key: %v", c.ComponentName, err)
	}
	err = os.WriteFile(keyPath, pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privPKCS,
	}), 0600)
	if err != nil {
		klog.Exitf("Failed to write %s private key: %v", c.ComponentName, err)
	}
	err = os.WriteFile(certPath, pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}), 0644)
	if err != nil {
		klog.Exitf("Failed to write %s certificate: %v", c.ComponentName, err)
	}

	return
}

// ensureDevCA ensures that a development CA certificate/key exists and returns
// paths to them. This data is either read from disk if it already exists, or is
// generated when this function is called. If any problem occurs, the code
// panics.
func (c *ComponentConfig) ensureDevCA() (caCertPath string) {
	caKeyPath := c.DevCertsPath + "/ca.key"
	caCertPath = c.DevCertsPath + "/ca.cert"

	if err := os.MkdirAll(c.DevCertsPath, 0700); err != nil {
		klog.Exitf("Failed to make developer certificate directory: %v", err)
	}

	// Check if we already have a key/certificate.
	noKey := false
	if _, err := os.Stat(caKeyPath); os.IsNotExist(err) {
		noKey = true
	}
	noCert := false
	if _, err := os.Stat(caCertPath); os.IsNotExist(err) {
		noCert = true
	}

	if noKey || noCert {
		klog.Infof("Generating developer CA certificate...")
	} else {
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	// No key/certificate, generate them.
	ca := pki.CA(fmt.Sprintf("monogon dev certs CA (%s)", hostname))

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		klog.Exitf("Failed to generate CA serial number: %v", err)
	}
	ca.SerialNumber = serialNumber
	ca.NotBefore = time.Now()
	ca.NotAfter = pki.UnknownNotAfter
	ca.BasicConstraintsValid = true

	caPub, caPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		klog.Exitf("Failed to generate CA key: %v", err)
	}
	caBytes, err := x509.CreateCertificate(rand.Reader, &ca, &ca, caPub, caPriv)
	if err != nil {
		klog.Exitf("Failed to generate CA certificate: %v", err)
	}

	// And marshal them to disk.
	caPrivPKCS, err := x509.MarshalPKCS8PrivateKey(caPriv)
	if err != nil {
		klog.Exitf("Failed to marshal %s private key: %v", c.ComponentName, err)
	}
	err = os.WriteFile(caKeyPath, pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: caPrivPKCS,
	}), 0600)
	if err != nil {
		klog.Exitf("Failed to write CA private key: %v", err)
	}
	err = os.WriteFile(caCertPath, pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	}), 0644)
	if err != nil {
		klog.Exitf("Failed to write CA certificate: %v", err)
	}

	return
}
