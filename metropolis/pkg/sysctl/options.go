package sysctl

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// Options contains sysctl options to apply
type Options map[string]string

// Apply attempts to apply all options in Options. It aborts on the first
// one which returns an error when applying.
func (o Options) Apply() error {
	for name, value := range o {
		filePath := path.Join("/proc/sys/", strings.ReplaceAll(name, ".", "/"))
		optionFile, err := os.OpenFile(filePath, os.O_WRONLY, 0)
		if err != nil {
			return fmt.Errorf("failed to set option %v: %w", name, err)
		}
		if _, err := optionFile.WriteString(value + "\n"); err != nil {
			optionFile.Close()
			return fmt.Errorf("failed to set option %v: %w", name, err)
		}
		optionFile.Close() // In a loop, defer'ing could open a lot of FDs
	}
	return nil
}
