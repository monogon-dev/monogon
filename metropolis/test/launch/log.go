// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package launch

import (
	"fmt"
	"strings"
)

// logf is compatible with the output of ConciseString as used in the Metropolis
// console log, making the output more readable in unified test logs.
func logf(f string, args ...any) {
	formatted := fmt.Sprintf(f, args...)
	for i, line := range strings.Split(formatted, "\n") {
		if len(line) == 0 {
			continue
		}
		if i == 0 {
			fmt.Printf("TT| %20s ! %s\n", "test launch", line)
		} else {
			fmt.Printf("TT| %20s | %s\n", "", line)
		}
	}
}
