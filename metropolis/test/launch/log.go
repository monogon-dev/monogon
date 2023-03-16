package launch

import (
	"fmt"
	"os"
	"strings"
)

// Log is compatible with the output of ConciseString as used in the Metropolis
// console log, making the output more readable in unified test logs.
func Log(f string, args ...any) {
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

func Fatal(f string, args ...any) {
	Log(f, args...)
	os.Exit(1)
}
