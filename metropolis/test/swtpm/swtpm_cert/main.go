package main

// swtpm_cert (from swtpm project) reimplemented in Go.
//
// This tool generates a TPM EK or Platform certificate. These certificates have
// to be in a very specific format and include non-standard extensions and
// subjectAltNames.

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"

	"source.monogon.dev/osbase/pki"
)

func getSignkey() *rsa.PrivateKey {
	if strings.HasPrefix(flagSignKey, "tpmkey:") || strings.HasPrefix(flagSignKey, "pkcs11:") {
		log.Fatalf("Loading tpmkey: and pkcs11: sign keys unimplemented")
	}
	bytes, err := os.ReadFile(flagSignKey)
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

func getPubkey() any {
	if flagModulus != "" {
		if flagECCX != "" || flagECCY != "" || flagECCCurveID != "" {
			log.Fatalf("--modulus and --ecc* cannot be set simultaneously")
		}
		var modulus big.Int
		modulusBytes, err := hex.DecodeString(flagModulus)
		if err != nil {
			log.Fatalf("Could not decode modulus: %v", err)
		}
		modulus.SetBytes(modulusBytes)
		return &rsa.PublicKey{
			N: &modulus,
			E: flagExponent,
		}
	}
	if flagECCX != "" && flagECCY != "" && flagECCCurveID != "" {
		if flagModulus != "" {
			log.Fatalf("--modulus and --ecc* cannot be set simultaneously")
		}
		var x, y big.Int
		xBytes, err := hex.DecodeString(flagECCX)
		if err != nil {
			log.Fatalf("Could not decode ECC X: %v", err)
		}
		x.SetBytes(xBytes)
		yBytes, err := hex.DecodeString(flagECCY)
		if err != nil {
			log.Fatalf("Could not decode ECC Y: %v", err)
		}
		y.SetBytes(yBytes)
		res := ecdsa.PublicKey{X: &x, Y: &y}
		switch flagECCCurveID {
		case "secp256r1":
			res.Curve = elliptic.P256()
		case "secp384r1":
			res.Curve = elliptic.P384()
		default:
			log.Fatalf("Unknown ECC curve ID %q", flagECCCurveID)
		}
		return &res
	}
	log.Fatalf("--modulus or --ecc* must be set")
	panic("unreachable")
}

func getIssuerCert() *x509.Certificate {
	bytes, err := os.ReadFile(flagIssuerCert)
	if err != nil {
		log.Fatalf("Could not read issuer certificate: %v", err)
	}
	block, _ := pem.Decode(bytes)
	if block.Type != "CERTIFICATE" {
		log.Fatalf("Issuer certificate contains invalid PEM data")
	}
	res, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("Could not parse issuer certificate: %v", err)
	}
	return res
}

type certType string

const (
	certTypeEK       certType = "ek"
	certTypePlatform certType = "platform"
)

var (
	flagType                 string
	flagSubject              string
	flagPlatformManufacturer string
	flagPlatformVersion      string
	flagPlatformModel        string
	flagTPM2                 bool
	flagTPMSpecFamily        string
	flagTPMSpecLevel         int
	flagTPMSpecRevision      int
	flagTPMManufacturer      string
	flagTPMModel             string
	flagTPMVersion           string
	flagOutCert              string
	flagExponent             int
	flagSignKey              string
	flagIssuerCert           string
	flagDays                 int
	flagSerial               string

	flagModulus string

	flagECCX       string
	flagECCY       string
	flagECCCurveID string
)

func main() {
	pflag.BoolVar(&flagTPM2, "tpm2", false, "Enable TPM2 mode (no-op, only mode supported)")
	pflag.StringVar(&flagType, "type", "ek", "Type of certificate to create, ek or platform")

	pflag.StringVar(&flagPlatformManufacturer, "platform-manufacturer", "", "TPM platform manufacturer")
	pflag.StringVar(&flagPlatformVersion, "platform-version", "", "TPM platform version")
	pflag.StringVar(&flagPlatformModel, "platform-model", "", "TPM platform model")

	pflag.StringVar(&flagTPMSpecFamily, "tpm-spec-family", "", "TPM Specification family")
	pflag.IntVar(&flagTPMSpecLevel, "tpm-spec-level", -1, "TPM Specification level")
	pflag.IntVar(&flagTPMSpecRevision, "tpm-spec-revision", -1, "TPM Specification revision")

	pflag.StringVar(&flagTPMManufacturer, "tpm-manufacturer", "", "TPM device manufacturer")
	pflag.StringVar(&flagTPMModel, "tpm-model", "", "TPM device model")
	pflag.StringVar(&flagTPMVersion, "tpm-version", "", "TPM device version")

	pflag.StringVar(&flagSubject, "subject", "", "Certificate subject (only cn=... is implemented)")
	pflag.IntVar(&flagDays, "days", 0, "")
	pflag.StringVar(&flagSerial, "serial", "", "")

	pflag.StringVar(&flagOutCert, "out-cert", "", "Path to generated certificate (.pem)")
	pflag.StringVar(&flagIssuerCert, "issuercert", "", "Path to issuer certificate (.pem)")
	pflag.StringVar(&flagSignKey, "signkey", "", "Path to private key used to sign certificate")

	pflag.IntVar(&flagExponent, "exponent", 0x10001, "RSA key exponent")
	pflag.StringVar(&flagModulus, "modulus", "", "RSA key modulus")

	pflag.StringVar(&flagECCX, "ecc-x", "", "ECC key x component")
	pflag.StringVar(&flagECCY, "ecc-y", "", "ECC key y component")
	pflag.StringVar(&flagECCCurveID, "ecc-curveid", "", "ECC curve id (one of secp256r1, secp384r1)")
	pflag.Parse()

	var ty certType
	switch flagType {
	case "ek":
		ty = certTypeEK
	case "platform":
		ty = certTypePlatform
	default:
		log.Fatalf("Unknown type %q (must be ek or platform)", flagType)
	}
	if ty == certTypeEK || ty == certTypePlatform {
		if flagTPMManufacturer == "" {
			log.Fatalf("--tpm-manufacturer must be set")
		}
		if flagTPMModel == "" {
			log.Fatalf("--tpm-model must be set")
		}
		if flagTPMVersion == "" {
			log.Fatalf("--tpm-version must be set")
		}
	}
	if ty == certTypeEK {
		if flagTPMSpecFamily == "" {
			log.Fatalf("--tpm-spec-family must be set")
		}
		if flagTPMSpecLevel == -1 {
			log.Fatalf("--tpm-spec-level must be set")
		}
		if flagTPMSpecRevision == -1 {
			log.Fatalf("--tpm-spec-revision must be set")
		}
	}
	if ty == certTypePlatform {
		if flagPlatformManufacturer == "" {
			log.Fatalf("--platform-manufacturer must be set")
		}
		if flagPlatformModel == "" {
			log.Fatalf("--platform-model must be set")
		}
		if flagPlatformVersion == "" {
			log.Fatalf("--platform-version must be set")
		}
	}

	pubkey := getPubkey()
	signkey := getSignkey()
	issuercert := getIssuerCert()

	var cert x509.Certificate
	cert.Version = 3
	cert.SerialNumber = big.NewInt(0)
	if _, ok := cert.SerialNumber.SetString(flagSerial, 10); !ok {
		log.Fatalf("Could not parse serial %q", flagSerial)
	}
	cert.NotBefore = time.Now()
	if flagDays > 0 {
		cert.NotAfter = time.Now().Add(time.Hour * 24 * time.Duration(flagDays))
	} else {
		cert.NotAfter = pki.UnknownNotAfter
	}
	if flagSubject != "" {
		parts := strings.Split(flagSubject, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			els := strings.SplitN(part, "=", 2)
			k := strings.ToLower(els[0])
			switch k {
			case "cn":
				cert.Subject.CommonName = els[1]
			default:
				log.Fatalf("Unparseable subject: %q", flagSubject)
			}
		}
	}
	var sanValues = []asn1.RawValue{}
	switch ty {
	case certTypeEK:
		sanValues = append(sanValues, asn1.RawValue{
			Tag: 4, Class: 2, Bytes: buildManufacturerInfo(flagTPMManufacturer, flagTPMModel, flagTPMVersion),
		})
	case certTypePlatform:
		sanValues = append(sanValues, asn1.RawValue{
			Tag: 4, Class: 2, Bytes: buildPlatformManufacturerInfo(flagPlatformManufacturer, flagPlatformModel, flagPlatformVersion),
		})
	}
	sanBytes, err := asn1.Marshal(sanValues)
	if err != nil {
		log.Fatalf("Failed to marshal SAN values: %v", err)
	}
	cert.ExtraExtensions = []pkix.Extension{
		{
			Id:    asn1.ObjectIdentifier{2, 5, 29, 17}, // subjectAltName
			Value: sanBytes,
		},
	}
	if ty == certTypeEK {
		cert.ExtraExtensions = append(cert.ExtraExtensions, pkix.Extension{
			Id:    asn1.ObjectIdentifier{2, 5, 29, 9}, // directoryName
			Value: buildSpecificationInfo(flagTPMSpecFamily, flagTPMSpecLevel, flagTPMSpecRevision),
		})
	}
	cert.BasicConstraintsValid = true
	cert.IsCA = false
	switch ty {
	case certTypeEK:
		// tcg-kp-EKCertificate
		cert.UnknownExtKeyUsage = []asn1.ObjectIdentifier{{2, 23, 133, 8, 1}}
	case certTypePlatform:
		// tcg-kp-PlatformAttributeCertificate
		cert.UnknownExtKeyUsage = []asn1.ObjectIdentifier{{2, 23, 133, 8, 2}}
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &cert, issuercert, pubkey, signkey)
	if err != nil {
		log.Fatalf("Generating certificate failed: %v", err)
	}
	block := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	if err := os.WriteFile(flagOutCert, pem.EncodeToMemory(&block), 0644); err != nil {
		log.Fatalf("Writing certificate failed: %v", err)
	}
}
