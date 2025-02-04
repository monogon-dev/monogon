// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package nvme

// SecureEraseType specifices what type of secure erase should be performed by
// by the controller. The zero value requests no secure erase.
type SecureEraseType uint8

const (
	// SecureEraseTypeNone specifies that no secure erase operation is
	// requested.
	SecureEraseTypeNone SecureEraseType = 0
	// SecureEraseTypeUserData specifies that all user data should be securely
	// erased. The controller is allowed to perform a cryptographic erase
	// instead.
	SecureEraseTypeUserData SecureEraseType = 1
	// SecureEraseTypeCryptographic specifies that the encryption key for user
	// data should be erased. This in turn causes all current user data to
	// become unreadable.
	SecureEraseTypeCryptographic SecureEraseType = 2
)

// ProtectionInformationType selects the type of end-to-end protection tags to
// use. NVMe supports the same types as T10 DIF (SCSI).
type ProtectionInformationType uint8

const (
	ProtectionInformationTypeNone ProtectionInformationType = 0
	ProtectionInformationType1    ProtectionInformationType = 1
	ProtectionInformationType2    ProtectionInformationType = 2
	ProtectionInformationType3    ProtectionInformationType = 3
)

type FormatRequest struct {
	// NamespaceID contains the ID of the namespace to format.
	// NamespaceGlobal formats all namespaces.
	NamespaceID uint32
	// SecureEraseSettings specifies the type of secure erase to perform.
	SecureEraseSettings SecureEraseType
	// ProtectionInformationLocation selects where protection information is
	// transmitted. If true, it is transmitted as the first 8 bytes of metadata.
	// If false, it is transmitted as the last 8 bytes of metadata.
	ProtectionInformationLocation bool
	// ProtectionInformation specifies the type of T10 DIF Protection
	// Information to use.
	ProtectionInformation ProtectionInformationType
	// MetadataInline selects whether metadata is transferred as part of an
	// extended data LBA. If false, metadata is returned in a separate buffer.
	// If true, metadata is appended to the data buffer.
	MetadataInline bool
	// LBAFormat specifies the LBA format to use. This needs to be selected
	// from the list of supported LBA formats in the Identify response.
	LBAFormat uint8
}

// Format performs a low-level format of the NVM media. This is used for
// changing the block and/or metadata size. This command causes all data
// on the specified namespace to be lost. By setting SecureEraseSettings
// to the appropriate value it can also be used to securely erase data.
// See also the Sanitize command for just wiping the device.
func (d *Device) Format(req *FormatRequest) error {
	var cdw10 uint32
	cdw10 |= uint32(req.SecureEraseSettings&0x7) << 9
	cdw10 |= uint32(req.ProtectionInformation&0x7) << 5
	cdw10 |= uint32(req.LBAFormat & 0x7)
	if req.ProtectionInformationLocation {
		cdw10 |= 1 << 8
	}
	if req.MetadataInline {
		cdw10 |= 1 << 4
	}
	return d.RawCommand(&Command{
		Opcode:      0x80,
		NamespaceID: req.NamespaceID,
		CDW10:       cdw10,
	})
}
