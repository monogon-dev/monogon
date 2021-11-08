// MIT License
//
// Copyright (c) 2021 Philippe Voinov (philippevoinov@gmail.com)
// Copyright 2021 The Monogon Project Authors.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package efivarfs

import (
	"fmt"
	"math"

	"golang.org/x/text/transform"
)

// Note on binary format of EFI variables:
// This code follows Section 3 "Boot Manager" of version 2.6 of the UEFI Spec:
// http://www.uefi.org/sites/default/files/resources/UEFI%20Spec%202_6.pdf
// It uses the binary representation from the Linux "efivars" filesystem.
// Specifically, all binary data that is marshaled and unmarshaled is preceded by
// 4 bytes of "Variable Attributes".
// All binary data must have exactly the correct length and may not be padded
// with extra bytes while reading or writing.

// Note on EFI variable attributes:
// This code ignores all EFI variable attributes when reading.
// This code always writes variables with the following attributes:
//   - EFI_VARIABLE_NON_VOLATILE (0x00000001)
//   - EFI_VARIABLE_BOOTSERVICE_ACCESS (0x00000002)
//   - EFI_VARIABLE_RUNTIME_ACCESS (0x00000004)
const defaultAttrsByte0 uint8 = 7

// BootEntry represents a subset of the contents of a Boot#### EFI variable.
type BootEntry struct {
	Description     string // eg. "Linux Boot Manager"
	Path            string // eg. `\EFI\systemd\systemd-bootx64.efi`
	PartitionGUID   string
	PartitionNumber uint32 // Starts with 1
	PartitionStart  uint64 // LBA
	PartitionSize   uint64 // LBA
}

// Marshal generates the binary representation of a BootEntry (EFI_LOAD_OPTION).
// Description, DiskGUID and Path must be set.
// Attributes of the boot entry (EFI_LOAD_OPTION.Attributes, not the same
// as attributes of an EFI variable) are always set to LOAD_OPTION_ACTIVE.
func (t *BootEntry) Marshal() ([]byte, error) {
	if t.Description == "" ||
		t.PartitionGUID == "00000000-0000-0000-0000-000000000000" ||
		t.Path == "" ||
		t.PartitionNumber == 0 ||
		t.PartitionStart == 0 ||
		t.PartitionSize == 0 {
		return nil, fmt.Errorf("missing field, all are required: %+v", *t)
	}

	// EFI_LOAD_OPTION.FilePathList
	var dp []byte

	// EFI_LOAD_OPTION.FilePathList[0]
	dp = append(dp,
		0x04,       // Type ("Media Device Path")
		0x01,       // Sub-Type ("Hard Drive")
		0x2a, 0x00, // Length (always 42 bytes for this type)
	)
	dp = append32(dp, t.PartitionNumber)
	dp = append64(dp, t.PartitionStart)
	dp = append64(dp, t.PartitionSize)
	dp = append(dp, t.PartitionGUID[0:16]...) // Partition Signature
	dp = append(dp,
		0x02, // Partition Format ("GUID Partition Table")
		0x02, // Signature Type ("GUID signature")
	)

	// EFI_LOAD_OPTION.FilePathList[1]
	enc := Encoding.NewEncoder()
	path, _, e := transform.Bytes(enc, []byte(t.Path))
	if e != nil {
		return nil, fmt.Errorf("while encoding Path: %v", e)
	}
	path = append16(path, 0) // null terminate string
	filePathLen := len(path) + 4
	dp = append(dp,
		0x04, // Type ("Media Device Path")
		0x04, // Sub-Type ("File Path")
	)
	dp = append16(dp, uint16(filePathLen))
	dp = append(dp, path...)

	// EFI_LOAD_OPTION.FilePathList[2] ("Device Path End Structure")
	dp = append(dp,
		0x7F,       // Type ("End of Hardware Device Path")
		0xFF,       // Sub-Type ("End Entire Device Path")
		0x04, 0x00, // Length (always 4 bytes for this type)
	)

	out := []byte{
		// EFI variable attributes
		defaultAttrsByte0, 0x00, 0x00, 0x00,

		// EFI_LOAD_OPTION.Attributes (only LOAD_OPTION_ACTIVE)
		0x01, 0x00, 0x00, 0x00,
	}

	// EFI_LOAD_OPTION.FilePathListLength
	if len(dp) > math.MaxUint16 {
		// No need to also check for overflows for Path length field explicitly,
		// since if that overflows, this field will definitely overflow as well.
		// There is no explicit length field for Description, so no special
		// handling is required.
		return nil, fmt.Errorf("variable too large, use shorter strings")
	}
	out = append16(out, uint16(len(dp)))

	// EFI_LOAD_OPTION.Description
	desc, _, e := transform.Bytes(enc, []byte(t.Description))
	if e != nil {
		return nil, fmt.Errorf("while encoding Description: %v", e)
	}
	desc = append16(desc, 0) // null terminate string
	out = append(out, desc...)

	// EFI_LOAD_OPTION.FilePathList
	out = append(out, dp...)

	// EFI_LOAD_OPTION.OptionalData is always empty

	return out, nil
}

// UnmarshalBootEntry loads a BootEntry from its binary representation.
// WARNING: UnmarshalBootEntry only loads the Description field.
// Everything else is ignored (and not validated if possible)
func UnmarshalBootEntry(d []byte) (*BootEntry, error) {
	descOffset := 4 /* EFI Var Attrs */ + 4 /* EFI_LOAD_OPTION.Attributes */ + 2 /*FilePathListLength*/
	if len(d) < descOffset {
		return nil, fmt.Errorf("too short: %v bytes", len(d))
	}
	descBytes := []byte{}
	var foundNull bool
	for i := descOffset; i+1 < len(d); i += 2 {
		a := d[i]
		b := d[i+1]
		if a == 0 && b == 0 {
			foundNull = true
			break
		}
		descBytes = append(descBytes, a, b)
	}
	if !foundNull {
		return nil, fmt.Errorf("didn't find null terminator for Description")
	}
	descDecoded, _, e := transform.Bytes(Encoding.NewDecoder(), descBytes)
	if e != nil {
		return nil, fmt.Errorf("while decoding Description: %v", e)
	}
	return &BootEntry{Description: string(descDecoded)}, nil
}

// BootOrder represents the contents of the BootOrder EFI variable.
type BootOrder []uint16

// Marshal generates the binary representation of a BootOrder.
func (t *BootOrder) Marshal() []byte {
	out := []byte{defaultAttrsByte0, 0x00, 0x00, 0x00}
	for _, v := range *t {
		out = append16(out, v)
	}
	return out
}

// UnmarshalBootOrder loads a BootOrder from its binary representation.
func UnmarshalBootOrder(d []byte) (*BootOrder, error) {
	if len(d) < 4 || len(d)%2 != 0 {
		return nil, fmt.Errorf("invalid length: %v bytes", len(d))
	}
	l := (len(d) - 4) / 2
	out := make(BootOrder, l)
	for i := 0; i < l; i++ {
		out[i] = uint16(d[4+2*i]) | uint16(d[4+2*i+1])<<8
	}
	return &out, nil
}

func append16(d []byte, v uint16) []byte {
	return append(d,
		byte(v&0xFF),
		byte(v>>8&0xFF),
	)
}

func append32(d []byte, v uint32) []byte {
	return append(d,
		byte(v&0xFF),
		byte(v>>8&0xFF),
		byte(v>>16&0xFF),
		byte(v>>24&0xFF),
	)
}

func append64(d []byte, v uint64) []byte {
	return append(d,
		byte(v&0xFF),
		byte(v>>8&0xFF),
		byte(v>>16&0xFF),
		byte(v>>24&0xFF),
		byte(v>>32&0xFF),
		byte(v>>40&0xFF),
		byte(v>>48&0xFF),
		byte(v>>56&0xFF),
	)
}
