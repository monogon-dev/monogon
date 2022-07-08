package main

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var noCredentialsError = errors.New("owner certificate or key does not exist")

// getCredentials returns Metropolis credentials (if any) from the current
// metroctl config directory.
func getCredentials() (cert *x509.Certificate, key ed25519.PrivateKey, err error) {
	ownerPrivateKeyPEM, err := os.ReadFile(filepath.Join(flags.configPath, "owner-key.pem"))
	if os.IsNotExist(err) {
		return nil, nil, noCredentialsError
	} else if err != nil {
		return nil, nil, fmt.Errorf("failed to load owner private key: %w", err)
	}
	block, _ := pem.Decode(ownerPrivateKeyPEM)
	if block == nil {
		return nil, nil, errors.New("owner-key.pem contains invalid PEM armoring")
	}
	if block.Type != ownerKeyType {
		return nil, nil, fmt.Errorf("owner-key.pem contains a PEM block that's not a %v", ownerKeyType)
	}
	if len(block.Bytes) != ed25519.PrivateKeySize {
		return nil, nil, errors.New("owner-key.pem contains a non-Ed25519 key")
	}
	key = block.Bytes
	ownerCertPEM, err := os.ReadFile(filepath.Join(flags.configPath, "owner.pem"))
	if os.IsNotExist(err) {
		return nil, nil, noCredentialsError
	} else if err != nil {
		return nil, nil, fmt.Errorf("failed to load owner certificate: %w", err)
	}
	block, _ = pem.Decode(ownerCertPEM)
	if block == nil {
		return nil, nil, errors.New("owner.pem contains invalid PEM armoring")
	}
	if block.Type != "CERTIFICATE" {
		return nil, nil, fmt.Errorf("owner.pem contains a PEM block that's not a CERTIFICATE")
	}
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("owner.pem contains an invalid X.509 certificate: %w", err)
	}
	return
}
