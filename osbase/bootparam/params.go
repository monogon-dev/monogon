// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package bootparam

import (
	"regexp"
	"strings"
)

var validTTYRegexp = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// Consoles returns the set of consoles passed to the kernel, i.e. the values
// passed to the console= directive. It normalizes away any possibly present
// /dev/ prefix, returning values like ttyS0. It returns an empty set in case
// no valid console parameters exist.
func (p Params) Consoles() map[string]bool {
	consoles := make(map[string]bool)
	for _, pa := range p {
		if pa.Param == "console" {
			consoleParts := strings.Split(pa.Value, ",")
			consoleName := strings.TrimPrefix(consoleParts[0], "/dev/")
			if validTTYRegexp.MatchString(consoleName) {
				consoles[consoleName] = true
			}
		}
	}
	return consoles
}
