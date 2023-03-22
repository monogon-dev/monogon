package main

import (
	"fmt"
	"os"
)

var logC = make(chan string)

// logPiper pipes log entries submitted via logf and panicf into whatever
// consoles are available to the system.
func logPiper() {
	var consoles []*os.File
	for _, p := range []string{"/dev/tty0", "/dev/ttyS0"} {
		f, err := os.OpenFile(p, os.O_WRONLY, 0)
		if err != nil {
			continue
		}
		consoles = append(consoles, f)
	}

	for {
		s := <-logC
		for _, c := range consoles {
			fmt.Fprintf(c, "%s\n", s)
		}
	}
}

// logf logs some format/args into the active consoles.
func logf(format string, args ...any) {
	s := fmt.Sprintf(format, args...)
	logC <- s
}

// panicf aborts the installation process with a given format/args.
func panicf(format string, args ...any) {
	s := fmt.Sprintf(format, args...)
	// We don't need to print `s` here, as it's gonna get printed by the recovery
	// code in main.
	panic(s)
}
