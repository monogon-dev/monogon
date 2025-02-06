// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Taken from go-attestation under Apache 2.0
package internal

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"unicode/utf16"

	"github.com/google/certificate-transparency-go/asn1"
	"github.com/google/certificate-transparency-go/x509"
)

const (
	// maxNameLen is the maximum accepted byte length for a name field.
	// This value should be larger than any reasonable value.
	maxNameLen = 2048
	// maxDataLen is the maximum size in bytes of a variable data field.
	// This value should be larger than any reasonable value.
	maxDataLen = 1024 * 1024 // 1 Megabyte.
)

// GUIDs representing the contents of an UEFI_SIGNATURE_LIST.
var (
	hashSHA256SigGUID        = efiGUID{0xc1c41626, 0x504c, 0x4092, [8]byte{0xac, 0xa9, 0x41, 0xf9, 0x36, 0x93, 0x43, 0x28}}
	hashSHA1SigGUID          = efiGUID{0x826ca512, 0xcf10, 0x4ac9, [8]byte{0xb1, 0x87, 0xbe, 0x01, 0x49, 0x66, 0x31, 0xbd}}
	hashSHA224SigGUID        = efiGUID{0x0b6e5233, 0xa65c, 0x44c9, [8]byte{0x94, 0x07, 0xd9, 0xab, 0x83, 0xbf, 0xc8, 0xbd}}
	hashSHA384SigGUID        = efiGUID{0xff3e5307, 0x9fd0, 0x48c9, [8]byte{0x85, 0xf1, 0x8a, 0xd5, 0x6c, 0x70, 0x1e, 0x01}}
	hashSHA512SigGUID        = efiGUID{0x093e0fae, 0xa6c4, 0x4f50, [8]byte{0x9f, 0x1b, 0xd4, 0x1e, 0x2b, 0x89, 0xc1, 0x9a}}
	keyRSA2048SigGUID        = efiGUID{0x3c5766e8, 0x269c, 0x4e34, [8]byte{0xaa, 0x14, 0xed, 0x77, 0x6e, 0x85, 0xb3, 0xb6}}
	certRSA2048SHA256SigGUID = efiGUID{0xe2b36190, 0x879b, 0x4a3d, [8]byte{0xad, 0x8d, 0xf2, 0xe7, 0xbb, 0xa3, 0x27, 0x84}}
	certRSA2048SHA1SigGUID   = efiGUID{0x67f8444f, 0x8743, 0x48f1, [8]byte{0xa3, 0x28, 0x1e, 0xaa, 0xb8, 0x73, 0x60, 0x80}}
	certX509SigGUID          = efiGUID{0xa5c059a1, 0x94e4, 0x4aa7, [8]byte{0x87, 0xb5, 0xab, 0x15, 0x5c, 0x2b, 0xf0, 0x72}}
	certHashSHA256SigGUID    = efiGUID{0x3bd2a492, 0x96c0, 0x4079, [8]byte{0xb4, 0x20, 0xfc, 0xf9, 0x8e, 0xf1, 0x03, 0xed}}
	certHashSHA384SigGUID    = efiGUID{0x7076876e, 0x80c2, 0x4ee6, [8]byte{0xaa, 0xd2, 0x28, 0xb3, 0x49, 0xa6, 0x86, 0x5b}}
	certHashSHA512SigGUID    = efiGUID{0x446dbf63, 0x2502, 0x4cda, [8]byte{0xbc, 0xfa, 0x24, 0x65, 0xd2, 0xb0, 0xfe, 0x9d}}
)

// EventType describes the type of event signalled in the event log.
type EventType uint32

// BIOS Events (TCG PC Client Specific Implementation Specification for
// Conventional BIOS 1.21)
const (
	PrebootCert          EventType = 0x00000000
	PostCode             EventType = 0x00000001
	unused               EventType = 0x00000002
	NoAction             EventType = 0x00000003
	Separator            EventType = 0x00000004
	Action               EventType = 0x00000005
	EventTag             EventType = 0x00000006
	SCRTMContents        EventType = 0x00000007
	SCRTMVersion         EventType = 0x00000008
	CpuMicrocode         EventType = 0x00000009
	PlatformConfigFlags  EventType = 0x0000000A
	TableOfDevices       EventType = 0x0000000B
	CompactHash          EventType = 0x0000000C
	Ipl                  EventType = 0x0000000D
	IplPartitionData     EventType = 0x0000000E
	NonhostCode          EventType = 0x0000000F
	NonhostConfig        EventType = 0x00000010
	NonhostInfo          EventType = 0x00000011
	OmitBootDeviceEvents EventType = 0x00000012
)

// EFI Events (TCG EFI Platform Specification Version 1.22)
const (
	EFIEventBase               EventType = 0x80000000
	EFIVariableDriverConfig    EventType = 0x80000001
	EFIVariableBoot            EventType = 0x80000002
	EFIBootServicesApplication EventType = 0x80000003
	EFIBootServicesDriver      EventType = 0x80000004
	EFIRuntimeServicesDriver   EventType = 0x80000005
	EFIGPTEvent                EventType = 0x80000006
	EFIAction                  EventType = 0x80000007
	EFIPlatformFirmwareBlob    EventType = 0x80000008
	EFIHandoffTables           EventType = 0x80000009
	EFIHCRTMEvent              EventType = 0x80000010
	EFIVariableAuthority       EventType = 0x800000e0
)

// ErrSigMissingGUID is returned if an EFI_SIGNATURE_DATA structure was parsed
// successfully, however was missing the SignatureOwner GUID. This case is
// handled specially as a workaround for a bug relating to authority events.
var ErrSigMissingGUID = errors.New("signature data was missing owner GUID")

var eventTypeNames = map[EventType]string{
	PrebootCert:          "Preboot Cert",
	PostCode:             "POST Code",
	unused:               "Unused",
	NoAction:             "No Action",
	Separator:            "Separator",
	Action:               "Action",
	EventTag:             "Event Tag",
	SCRTMContents:        "S-CRTM Contents",
	SCRTMVersion:         "S-CRTM Version",
	CpuMicrocode:         "CPU Microcode",
	PlatformConfigFlags:  "Platform Config Flags",
	TableOfDevices:       "Table of Devices",
	CompactHash:          "Compact Hash",
	Ipl:                  "IPL",
	IplPartitionData:     "IPL Partition Data",
	NonhostCode:          "Non-Host Code",
	NonhostConfig:        "Non-HostConfig",
	NonhostInfo:          "Non-Host Info",
	OmitBootDeviceEvents: "Omit Boot Device Events",

	EFIEventBase:               "EFI Event Base",
	EFIVariableDriverConfig:    "EFI Variable Driver Config",
	EFIVariableBoot:            "EFI Variable Boot",
	EFIBootServicesApplication: "EFI Boot Services Application",
	EFIBootServicesDriver:      "EFI Boot Services Driver",
	EFIRuntimeServicesDriver:   "EFI Runtime Services Driver",
	EFIGPTEvent:                "EFI GPT Event",
	EFIAction:                  "EFI Action",
	EFIPlatformFirmwareBlob:    "EFI Platform Firmware Blob",
	EFIVariableAuthority:       "EFI Variable Authority",
	EFIHandoffTables:           "EFI Handoff Tables",
	EFIHCRTMEvent:              "EFI H-CRTM Event",
}

func (e EventType) String() string {
	if s, ok := eventTypeNames[e]; ok {
		return s
	}
	return fmt.Sprintf("EventType(0x%x)", uint32(e))
}

// UntrustedParseEventType returns the event type indicated by
// the provided value.
func UntrustedParseEventType(et uint32) (EventType, error) {
	// "The value associated with a UEFI specific platform event type MUST be in
	// the range between 0x80000000 and 0x800000FF, inclusive."
	if (et < 0x80000000 || et > 0x800000FF) && et > 0x12 {
		return EventType(0), fmt.Errorf("event type not between [0x0, 0x12] or [0x80000000, 0x800000FF]: got %#x", et)
	}
	if _, ok := eventTypeNames[EventType(et)]; !ok {
		return EventType(0), fmt.Errorf("unknown event type %#x", et)
	}
	return EventType(et), nil
}

// efiGUID represents the EFI_GUID type.
// See section "2.3.1 Data Types" in the specification for more information.
// type efiGUID [16]byte
type efiGUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func (d efiGUID) String() string {
	var u [8]byte
	binary.BigEndian.PutUint32(u[:4], d.Data1)
	binary.BigEndian.PutUint16(u[4:6], d.Data2)
	binary.BigEndian.PutUint16(u[6:8], d.Data3)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[:4], u[4:6], u[6:8], d.Data4[:2], d.Data4[2:])
}

// UEFIVariableDataHeader represents the leading fixed-size fields
// within UEFI_VARIABLE_DATA.
type UEFIVariableDataHeader struct {
	VariableName       efiGUID
	UnicodeNameLength  uint64 // uintN
	VariableDataLength uint64 // uintN
}

// UEFIVariableData represents the UEFI_VARIABLE_DATA structure.
type UEFIVariableData struct {
	Header       UEFIVariableDataHeader
	UnicodeName  []uint16
	VariableData []byte // []int8
}

// ParseUEFIVariableData parses the data section of an event structured as a
// UEFI variable.
//
//	https://trustedcomputinggroup.org/wp-content/uploads/TCG_PCClient_Specific_Platform_Profile_for_TPM_2p0_1p04_PUBLIC.pdf#page=100
func ParseUEFIVariableData(r io.Reader) (ret UEFIVariableData, err error) {
	err = binary.Read(r, binary.LittleEndian, &ret.Header)
	if err != nil {
		return
	}
	if ret.Header.UnicodeNameLength > maxNameLen {
		return UEFIVariableData{}, fmt.Errorf("unicode name too long: %d > %d", ret.Header.UnicodeNameLength, maxNameLen)
	}
	ret.UnicodeName = make([]uint16, ret.Header.UnicodeNameLength)
	for i := 0; uint64(i) < ret.Header.UnicodeNameLength; i++ {
		err = binary.Read(r, binary.LittleEndian, &ret.UnicodeName[i])
		if err != nil {
			return
		}
	}
	if ret.Header.VariableDataLength > maxDataLen {
		return UEFIVariableData{}, fmt.Errorf("variable data too long: %d > %d", ret.Header.VariableDataLength, maxDataLen)
	}
	ret.VariableData = make([]byte, ret.Header.VariableDataLength)
	_, err = io.ReadFull(r, ret.VariableData)
	return
}

func (v *UEFIVariableData) VarName() string {
	return string(utf16.Decode(v.UnicodeName))
}

func (v *UEFIVariableData) SignatureData() (certs []x509.Certificate, hashes [][]byte, err error) {
	return parseEfiSignatureList(v.VariableData)
}

// UEFIVariableAuthority describes the contents of a UEFI variable authority
// event.
type UEFIVariableAuthority struct {
	Certs []x509.Certificate
}

// ParseUEFIVariableAuthority parses the data section of an event structured as
// a UEFI variable authority.
//
// https://uefi.org/sites/default/files/resources/UEFI_Spec_2_8_final.pdf#page=1789
func ParseUEFIVariableAuthority(r io.Reader) (UEFIVariableAuthority, error) {
	v, err := ParseUEFIVariableData(r)
	if err != nil {
		return UEFIVariableAuthority{}, err
	}
	certs, err := parseEfiSignature(v.VariableData)
	return UEFIVariableAuthority{Certs: certs}, err
}

// efiSignatureData represents the EFI_SIGNATURE_DATA type.  See section
// "31.4.1 Signature Database" in the specification for more information.
type efiSignatureData struct {
	SignatureOwner efiGUID
	SignatureData  []byte // []int8
}

// efiSignatureList represents the EFI_SIGNATURE_LIST type.
// See section "31.4.1 Signature Database" in the specification for more
// information.
type efiSignatureListHeader struct {
	SignatureType       efiGUID
	SignatureListSize   uint32
	SignatureHeaderSize uint32
	SignatureSize       uint32
}

type efiSignatureList struct {
	Header        efiSignatureListHeader
	SignatureData []byte
	Signatures    []byte
}

// parseEfiSignatureList parses a EFI_SIGNATURE_LIST structure.
// The structure and related GUIDs are defined at:
// https://uefi.org/sites/default/files/resources/UEFI_Spec_2_8_final.pdf#page=1790
func parseEfiSignatureList(b []byte) ([]x509.Certificate, [][]byte, error) {
	if len(b) < 28 {
		// Being passed an empty signature list here appears to be valid
		return nil, nil, nil
	}
	var signatures efiSignatureList
	buf := bytes.NewReader(b)
	var certificates []x509.Certificate
	var hashes [][]byte

	for buf.Len() > 0 {
		err := binary.Read(buf, binary.LittleEndian, &signatures.Header)
		if err != nil {
			return nil, nil, err
		}

		if signatures.Header.SignatureHeaderSize > maxDataLen {
			return nil, nil, fmt.Errorf("signature header too large: %d > %d", signatures.Header.SignatureHeaderSize, maxDataLen)
		}
		if signatures.Header.SignatureListSize > maxDataLen {
			return nil, nil, fmt.Errorf("signature list too large: %d > %d", signatures.Header.SignatureListSize, maxDataLen)
		}

		signatureType := signatures.Header.SignatureType
		switch signatureType {
		case certX509SigGUID: // X509 certificate
			for sigOffset := 0; uint32(sigOffset) < signatures.Header.SignatureListSize-28; {
				var signature efiSignatureData
				signature.SignatureData = make([]byte, signatures.Header.SignatureSize-16)
				err := binary.Read(buf, binary.LittleEndian, &signature.SignatureOwner)
				if err != nil {
					return nil, nil, err
				}
				err = binary.Read(buf, binary.LittleEndian, &signature.SignatureData)
				if err != nil {
					return nil, nil, err
				}
				cert, err := x509.ParseCertificate(signature.SignatureData)
				if err != nil {
					return nil, nil, err
				}
				sigOffset += int(signatures.Header.SignatureSize)
				certificates = append(certificates, *cert)
			}
		case hashSHA256SigGUID: // SHA256
			for sigOffset := 0; uint32(sigOffset) < signatures.Header.SignatureListSize-28; {
				var signature efiSignatureData
				signature.SignatureData = make([]byte, signatures.Header.SignatureSize-16)
				err := binary.Read(buf, binary.LittleEndian, &signature.SignatureOwner)
				if err != nil {
					return nil, nil, err
				}
				err = binary.Read(buf, binary.LittleEndian, &signature.SignatureData)
				if err != nil {
					return nil, nil, err
				}
				hashes = append(hashes, signature.SignatureData)
				sigOffset += int(signatures.Header.SignatureSize)
			}
		case keyRSA2048SigGUID:
			err = errors.New("unhandled RSA2048 key")
		case certRSA2048SHA256SigGUID:
			err = errors.New("unhandled RSA2048-SHA256 key")
		case hashSHA1SigGUID:
			err = errors.New("unhandled SHA1 hash")
		case certRSA2048SHA1SigGUID:
			err = errors.New("unhandled RSA2048-SHA1 key")
		case hashSHA224SigGUID:
			err = errors.New("unhandled SHA224 hash")
		case hashSHA384SigGUID:
			err = errors.New("unhandled SHA384 hash")
		case hashSHA512SigGUID:
			err = errors.New("unhandled SHA512 hash")
		case certHashSHA256SigGUID:
			err = errors.New("unhandled X509-SHA256 hash metadata")
		case certHashSHA384SigGUID:
			err = errors.New("unhandled X509-SHA384 hash metadata")
		case certHashSHA512SigGUID:
			err = errors.New("unhandled X509-SHA512 hash metadata")
		default:
			err = fmt.Errorf("unhandled signature type %s", signatureType)
		}
		if err != nil {
			return nil, nil, err
		}
	}
	return certificates, hashes, nil
}

// EFISignatureData represents the EFI_SIGNATURE_DATA type.
// See section "31.4.1 Signature Database" in the specification
// for more information.
type EFISignatureData struct {
	SignatureOwner efiGUID
	SignatureData  []byte // []int8
}

func parseEfiSignature(b []byte) ([]x509.Certificate, error) {
	var certificates []x509.Certificate

	if len(b) < 16 {
		return nil, fmt.Errorf("invalid signature: buffer smaller than header (%d < %d)", len(b), 16)
	}

	buf := bytes.NewReader(b)
	var signature EFISignatureData
	signature.SignatureData = make([]byte, len(b)-16)

	if err := binary.Read(buf, binary.LittleEndian, &signature.SignatureOwner); err != nil {
		return certificates, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &signature.SignatureData); err != nil {
		return certificates, err
	}

	cert, err := x509.ParseCertificate(signature.SignatureData)
	if err == nil {
		certificates = append(certificates, *cert)
	} else {
		// A bug in shim may cause an event to be missing the SignatureOwner GUID.
		// We handle this, but signal back to the caller using ErrSigMissingGUID.
		var structuralError asn1.StructuralError
		if errors.As(err, &structuralError) {
			var err2 error
			cert, err2 = x509.ParseCertificate(b)
			if err2 == nil {
				certificates = append(certificates, *cert)
				err = ErrSigMissingGUID
			}
		}
	}
	return certificates, err
}
