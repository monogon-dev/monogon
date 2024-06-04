package main

// Minimal GnuTLS certtool-like tool in Go. Implements only what's needed for
// compatibility with `swtpm_localca`.

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

var (
	flagGeneratePrivkey     bool
	flagGenerateSelfSigned  bool
	flagGenerateCertificate bool

	flagOutfile           string
	flagTemplate          string
	flagLoadPrivkey       string
	flagLoadCAPrivkey     string
	flagLoadCACertificate string
)

func main() {
	pflag.BoolVar(&flagGeneratePrivkey, "generate-privkey", false, "Generate RSA private kay")
	pflag.BoolVar(&flagGenerateSelfSigned, "generate-self-signed", false, "Generate self-signed certificate")
	pflag.BoolVar(&flagGenerateCertificate, "generate-certificate", false, "Sign certificate")

	pflag.StringVar(&flagOutfile, "outfile", "", "Output file for operation")
	pflag.StringVar(&flagTemplate, "template", "", "Certificate template file (GnuTLS proprietary)")
	pflag.StringVar(&flagLoadPrivkey, "load-privkey", "", "Path to private key")
	pflag.StringVar(&flagLoadCAPrivkey, "load-ca-privkey", "", "Path to CA private key")
	pflag.StringVar(&flagLoadCACertificate, "load-ca-certificate", "", "Path to CA certificate")
	pflag.Parse()

	modesActive := 0
	for _, mode := range []bool{flagGeneratePrivkey, flagGenerateSelfSigned, flagGenerateCertificate} {
		if mode {
			modesActive++
		}
	}
	if modesActive != 1 {
		log.Fatalf("Exactly one of --generate-privkey, --generate-self-signed, --generate-certificate must be set")
	}

	if flagGeneratePrivkey || flagGenerateSelfSigned || flagGenerateCertificate {
		if flagOutfile == "" {
			log.Fatalf("--outfile must be set")
		}
	}
	if flagGenerateSelfSigned || flagGenerateCertificate {
		if flagTemplate == "" {
			log.Fatalf("--template must be set")
		}
		if flagLoadPrivkey == "" {
			log.Fatalf("--load-privkey must be set")
		}
		if flagOutfile == "" {
			log.Fatalf("--outfile must be set")
		}
	}
	if flagGenerateCertificate {
		if flagLoadCAPrivkey == "" {
			log.Fatalf("--load-ca-privkey must be set")
		}
		if flagLoadCACertificate == "" {
			log.Fatalf("--load-ca-certificate must be set")
		}
	}
	switch {
	case flagGeneratePrivkey:
		generatePrivkey(flagOutfile)
	case flagGenerateSelfSigned:
		generateSelfSigned(flagTemplate, flagLoadPrivkey, flagOutfile)
	case flagGenerateCertificate:
		generateCertificate(flagTemplate, flagLoadPrivkey, flagLoadCAPrivkey, flagLoadCACertificate, flagOutfile)
	}
}

func generatePrivkey(outfile string) {
	priv, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		log.Fatalf("Could not generate RSA key: %v", err)
	}
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}
	if err := os.WriteFile(outfile, pem.EncodeToMemory(&block), 0600); err != nil {
		log.Fatalf("Could not write RSA key: %v", err)
	}
}

// certificateFromTemplate parses a GnuTLS 'template' file. This template file is
// made up of newline-separated stanzas, with an optional 'data' part after a =
// character.
//
// Supported stanzas (on conflict last stanza wins):
//
//	cn=data: set subject CN to data
//	ca: mark certificate as CA
//	cert_signing_key: enable keyCertSign KeyUsage
//	expiration_days: number of days in which cert expires (if -1/default, no expiry date)
func certificateFromTemplate(template string) *x509.Certificate {
	serial, err := rand.Int(rand.Reader, big.NewInt(1).Lsh(big.NewInt(1), 128))
	if err != nil {
		log.Fatalf("Could not generate serial: %v", err)
	}
	res := x509.Certificate{
		SerialNumber:          serial,
		NotBefore:             time.Now().Add(-time.Minute),
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
	}
	for _, line := range strings.Split(template, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		switch parts[0] {
		case "cn":
			res.Subject.CommonName = parts[1]
		case "ca":
			res.IsCA = true
		case "cert_signing_key":
			res.KeyUsage |= x509.KeyUsageCertSign
		case "expiration_days":
			days, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Fatalf("Invalid expiration_days: %q", err)
			}
			if days != -1 {
				res.NotAfter = time.Now().Add(time.Hour * 24 * time.Duration(days))
			} else {
				res.NotAfter = time.Unix(253402300799, 0)
			}
		default:
			log.Fatalf("Unhandled template line %q", line)
		}
	}
	return &res
}

func readPrivkey(path string) *rsa.PrivateKey {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read private key: %v", err)
	}
	block, _ := pem.Decode(bytes)
	if block.Type != "RSA PRIVATE KEY" {
		log.Fatalf("Private key contains invalid PEM data")
	}
	res, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Could not parse private key: %v", err)
	}
	return res
}

func readCertificate(path string) *x509.Certificate {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read certificate: %v", err)
	}
	block, _ := pem.Decode(bytes)
	if block.Type != "CERTIFICATE" {
		log.Fatalf("Certificate contains invalid PEM data")
	}
	res, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("Could not parse certificate: %v", err)
	}
	return res
}

func generateSelfSigned(templatePath, privkeyPath, outfile string) {
	template, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("Could not read template: %v", err)
	}
	priv := readPrivkey(privkeyPath)
	cert := certificateFromTemplate(string(template))

	derBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, priv.Public(), priv)
	if err != nil {
		log.Fatalf("Could not generate self-signed certificate: %v", err)
	}
	block := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	if err := os.WriteFile(outfile, pem.EncodeToMemory(&block), 0600); err != nil {
		log.Fatalf("Could not write self-signed certificate: %v", err)
	}
}

func generateCertificate(templatePath, privkeyPath, caPrivkeyPath, caCertificatePath, outfile string) {
	template, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("Could not read template: %v", err)
	}
	priv := readPrivkey(privkeyPath)
	caPriv := readPrivkey(caPrivkeyPath)
	cert := certificateFromTemplate(string(template))
	ca := readCertificate(caCertificatePath)

	derBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, priv.Public(), caPriv)
	if err != nil {
		log.Fatalf("Could not generate certificate: %v", err)
	}
	block := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	if err := os.WriteFile(outfile, pem.EncodeToMemory(&block), 0600); err != nil {
		log.Fatalf("Could not write certificate: %v", err)
	}
}
