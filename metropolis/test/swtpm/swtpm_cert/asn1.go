package main

import (
	"encoding/asn1"
	"log"
)

type manufacturerInfo struct {
	Manufacturer struct {
		Sequence struct {
			OID  asn1.ObjectIdentifier
			Data string `asn1:"utf8"`
		}
	} `asn1:"set"`
	Model struct {
		Sequence struct {
			OID  asn1.ObjectIdentifier
			Data string `asn1:"utf8"`
		}
	} `asn1:"set"`
	Version struct {
		Sequence struct {
			OID  asn1.ObjectIdentifier
			Data string `asn1:"utf8"`
		}
	} `asn1:"set"`
}

// buildManufacturerInfo marshals TPM manufacturer info (TPMManufacturer
// structure from TCG EK Credential Profile For TPM Family 2.0; Level 0; Version
// 2.4; Revision 3; 16 July 2021).
//
// This is embedded as a directoryName GeneralName SubjectAltName in the
// generated X509 certificate for an EK.
func buildManufacturerInfo(manufacturer, model, version string) []byte {
	var v manufacturerInfo
	v.Manufacturer.Sequence.OID = asn1.ObjectIdentifier{2, 23, 133, 2, 1}
	v.Manufacturer.Sequence.Data = manufacturer
	v.Model.Sequence.OID = asn1.ObjectIdentifier{2, 23, 133, 2, 2}
	v.Model.Sequence.Data = model
	v.Version.Sequence.OID = asn1.ObjectIdentifier{2, 23, 133, 2, 3}
	v.Version.Sequence.Data = version

	res, err := asn1.Marshal(v)
	if err != nil {
		log.Fatalf("Failed to marshal manufacturer info: %v", err)
	}
	return res
}

type platformManufacturerInfo struct {
	Manufacturer struct {
		Sequence struct {
			OID  asn1.ObjectIdentifier
			Data string `asn1:"utf8"`
		}
	} `asn1:"set"`
	Model struct {
		Sequence struct {
			OID  asn1.ObjectIdentifier
			Data string `asn1:"utf8"`
		}
	} `asn1:"set"`
	Version struct {
		Sequence struct {
			OID  asn1.ObjectIdentifier
			Data string `asn1:"utf8"`
		}
	} `asn1:"set"`
}

// buildPlatformManufacturerInfo marshals TPM platform manufacturer info.
//
// See: TCG Platform Certificate Profile; Specification Version 1.1; Revision 19;
// 10 April 2020: Section 3.1.2 (Name Attributes
// Platform{ManufacturerStr,Model,Version}) and Section 3.2 (Platform
// Certificate, Extensions Subject Alternative Names).
//
// This is embedded as a directoryName GeneralName SubjectAltName in the
// generated X509 certificate for a Platform.
//
// The spec seems to have missing ASN.1 definitions to tie together the strings
// into a structure that's embedded into the SAN. This corresponds to whatever
// upstream swtpm_cert is doing.
func buildPlatformManufacturerInfo(manufacturer, model, version string) []byte {
	var v platformManufacturerInfo
	v.Manufacturer.Sequence.OID = asn1.ObjectIdentifier{2, 23, 133, 5, 1, 1}
	v.Manufacturer.Sequence.Data = manufacturer
	v.Model.Sequence.OID = asn1.ObjectIdentifier{2, 23, 133, 5, 1, 4}
	v.Model.Sequence.Data = model
	v.Version.Sequence.OID = asn1.ObjectIdentifier{2, 23, 133, 5, 1, 5}
	v.Version.Sequence.Data = version

	res, err := asn1.Marshal(v)
	if err != nil {
		log.Fatalf("Failed to marshal platform manufacturer info: %v", err)
	}
	return res
}

type specificationInfo struct {
	OID asn1.ObjectIdentifier
	Set struct {
		Sequence struct {
			Family   string
			Level    int
			Revision int
		}
	} `asn1:"set"`
}

// buildSpecificationInfo marshals TPM manufacturer info (tPMSpecification
// structure from TCG EK Credential Profile For TPM Family 2.0; Level 0; Version
// 2.4; Revision 3; 16 July 2021).
//
// This is embedded as a directoryName SAN or extension in the generated X509
// certificate for an EK.
func buildSpecificationInfo(family string, level, revision int) []byte {
	var v specificationInfo
	v.OID = asn1.ObjectIdentifier{2, 23, 133, 2, 16}
	v.Set.Sequence.Family = family
	v.Set.Sequence.Level = level
	v.Set.Sequence.Revision = revision
	res, err := asn1.Marshal(v)
	if err != nil {
		log.Fatalf("Failed to marshal specification info: %v", err)
	}
	return res
}
