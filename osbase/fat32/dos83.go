// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package fat32

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// By default, DOS names would be encoded as what Microsoft calls the OEM
// code page. This is however dependant on the code page settings of the
// OS reading the file name as it's not mentioned in FAT32 metadata.
// To get maximum compatibility and make it easy to read in hex editors
// this only encodes ASCII characters and not any specific code page.
// This can still result in garbled data when using a non-latin code page,
// but this is unavoidable.
// This is legal as there is no specific requirements for generating these
// DOS names and any semi-modern system should use the unicode filenames
// anyways.

var invalidDOSNameChar = regexp.MustCompile("^[^A-Z0-9!#$%&'()@^_\x60{}~-]$")

// validDOSName matches names which are valid and unique DOS 8.3 file names as
// well as valid ASCII
var validDOSName = regexp.MustCompile(`^^([A-Z0-9!#$%&'()@^_\x60{}~-]{0,8})(\.[A-Z0-9!#$%&'()-@^_\x60{}~-]{1,3})?$`)

func makeUniqueDOSNames(nodes []*node) error {
	taken := make(map[[11]byte]bool)
	var lossyNameNodes []*node
	// Make two passes to ensure that names can always be passed through even
	// if they would conflict with a generated name.
	for _, i := range nodes {
		for j := range i.dosName {
			i.dosName[j] = ' '
		}
		nameUpper := strings.ToUpper(i.Name)
		dosParts := validDOSName.FindStringSubmatch(nameUpper)
		if dosParts != nil {
			// Name is pass-through
			copy(i.dosName[:8], dosParts[1])
			if len(dosParts[2]) > 0 {
				// Skip the dot, it is implicit
				copy(i.dosName[8:], dosParts[2][1:])
			}
			if taken[i.dosName] {
				// Mapping is unique, complain about the actual file name, not
				// the 8.3 one
				return fmt.Errorf("name %q occurs more than once in the same directory", i.Name)
			}
			taken[i.dosName] = true
			continue
		}
		lossyNameNodes = append(lossyNameNodes, i)
	}
	// Willfully ignore the recommended short name generation algorithm as it
	// requires tons of bookkeeping and doesn't result in stable names so
	// cannot be relied on anyway.
	// A FAT32 directory is limited to 2^16 entries (in practice less than half
	// of that because of long file name entries), so 4 hex characters
	// guarantee uniqueness, regardless of the rest of name.
	var nameIdx int
	for _, i := range lossyNameNodes {
		nameUpper := strings.ToUpper(i.Name)
		dotParts := strings.Split(nameUpper, ".")
		for j := range dotParts {
			// Remove all invalid chars
			dotParts[j] = invalidDOSNameChar.ReplaceAllString(dotParts[j], "")
		}
		var fileName string
		lastDotPart := dotParts[len(dotParts)-1]
		if len(dotParts) > 1 && len(dotParts[0]) > 0 && len(lastDotPart) > 0 {
			// We have a valid 8.3 extension
			copy(i.dosName[8:], lastDotPart)
			fileName = strings.Join(dotParts[:len(dotParts)-1], "")
		} else {
			fileName = strings.Join(dotParts[:], "")
		}
		copy(i.dosName[:4], fileName)

		for {
			copy(i.dosName[4:], fmt.Sprintf("%04X", nameIdx))
			nameIdx++
			if nameIdx >= math.MaxUint16 {
				return errors.New("invariant violated: unable to find unique name with 16 bit counter in 16 bit space")
			}
			if !taken[i.dosName] {
				break
			}
		}
	}
	return nil
}
