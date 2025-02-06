// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package efivarfs provides functions to read and manipulate UEFI runtime
// variables. It uses Linux's efivarfs [1] to access the variables and all
// functions generally require that this is mounted at
// "/sys/firmware/efi/efivars".
//
// [1] https://www.kernel.org/doc/html/latest/filesystems/efivarfs.html
package efivarfs

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/text/encoding/unicode"
)

const (
	Path = "/sys/firmware/efi/efivars"
)

var (
	// ScopeGlobal is the scope of variables defined by the EFI specification
	// itself.
	ScopeGlobal = uuid.MustParse("8be4df61-93ca-11d2-aa0d-00e098032b8c")
	// ScopeSystemd is the scope of variables defined by Systemd/bootspec.
	ScopeSystemd = uuid.MustParse("4a67b082-0a4c-41cf-b6c7-440b29bb8c4f")
)

// Encoding defines the Unicode encoding used by UEFI, which is UCS-2 Little
// Endian. For BMP characters UTF-16 is equivalent to UCS-2. See the UEFI
// Spec 2.9, Sections 33.2.6 and 1.8.1.
var Encoding = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)

// Attribute contains a bitset of EFI variable attributes.
type Attribute uint32

const (
	// If set the value of the variable is is persistent across resets and
	// power cycles. Variables without this set cannot be created or modified
	// after UEFI boot services are terminated.
	AttrNonVolatile Attribute = 1 << iota
	// If set allows access to this variable from UEFI boot services.
	AttrBootserviceAccess
	// If set allows access to this variable from an operating system after
	// UEFI boot services are terminated. Variables setting this must also
	// set AttrBootserviceAccess. This is automatically taken care of by Write
	// in this package.
	AttrRuntimeAccess
	// Marks a variable as being a hardware error record. See UEFI 2.10 section
	// 8.2.8 for more information about this.
	AttrHardwareErrorRecord
	// Deprecated, should not be used for new variables.
	AttrAuthenticatedWriteAccess
	// Variable requires special authentication to write. These variables
	// cannot be written with this package.
	AttrTimeBasedAuthenticatedWriteAccess
	// If set in a Write() call, tries to append the data instead of replacing
	// it completely.
	AttrAppendWrite
	// Variable requires special authentication to access and write. These
	// variables cannot be accessed with this package.
	AttrEnhancedAuthenticatedAccess
)

func varPath(scope uuid.UUID, varName string) string {
	return fmt.Sprintf("/sys/firmware/efi/efivars/%s-%s", varName, scope.String())
}

// Write writes the value of the named variable in the given scope.
func Write(scope uuid.UUID, varName string, attrs Attribute, value []byte) error {
	// Write attributes, see @linux//Documentation/filesystems:efivarfs.rst for format
	f, err := os.OpenFile(varPath(scope, varName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		e := err
		// Unwrap PathError here as we wrap our own parameter message around it
		var perr *fs.PathError
		if errors.As(err, &perr) {
			e = perr.Err
		}
		return fmt.Errorf("writing %q in scope %s: %w", varName, scope, e)
	}
	// Required by UEFI 2.10 Section 8.2.3:
	// Runtime access to a data variable implies boot service access. Attributes
	// that have EFI_VARIABLE_RUNTIME_ACCESS set must also have
	// EFI_VARIABLE_BOOTSERVICE_ACCESS set. The caller is responsible for
	// following this rule.
	if attrs&AttrRuntimeAccess != 0 {
		attrs |= AttrBootserviceAccess
	}
	// Linux wants everything in on write, so assemble an intermediate buffer
	buf := make([]byte, len(value)+4)
	binary.LittleEndian.PutUint32(buf[:4], uint32(attrs))
	copy(buf[4:], value)
	_, err = f.Write(buf)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

// Read reads the value of the named variable in the given scope.
func Read(scope uuid.UUID, varName string) ([]byte, Attribute, error) {
	val, err := os.ReadFile(varPath(scope, varName))
	if err != nil {
		e := err
		// Unwrap PathError here as we wrap our own parameter message around it
		var perr *fs.PathError
		if errors.As(err, &perr) {
			e = perr.Err
		}
		return nil, Attribute(0), fmt.Errorf("reading %q in scope %s: %w", varName, scope, e)
	}
	if len(val) < 4 {
		return nil, Attribute(0), fmt.Errorf("reading %q in scope %s: malformed, less than 4 bytes long", varName, scope)
	}
	return val[4:], Attribute(binary.LittleEndian.Uint32(val[:4])), nil
}

// List lists all variable names present for a given scope sorted by their names
// in Go's "native" string sort order.
func List(scope uuid.UUID) ([]string, error) {
	vars, err := os.ReadDir(Path)
	if err != nil {
		return nil, fmt.Errorf("failed to list variable directory: %w", err)
	}
	var outVarNames []string
	suffix := fmt.Sprintf("-%v", scope)
	for _, v := range vars {
		if v.IsDir() {
			continue
		}
		if !strings.HasSuffix(v.Name(), suffix) {
			continue
		}
		outVarNames = append(outVarNames, strings.TrimSuffix(v.Name(), suffix))
	}
	return outVarNames, nil
}

// Delete deletes the given variable name in the given scope. Use with care,
// some firmware fails to boot if variables it uses are deleted.
func Delete(scope uuid.UUID, varName string) error {
	return os.Remove(varPath(scope, varName))
}
