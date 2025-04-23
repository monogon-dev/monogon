// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	statusFile = flag.String("status", "", "Path to bazel workspace status file")
	outFile    = flag.String("out", "", "Output stamp file path")
)

func main() {
	flag.Parse()

	vars := make(map[string]bool)
	for _, variable := range flag.Args() {
		vars[variable] = true
	}

	statusFileContent, err := os.ReadFile(*statusFile)
	if err != nil {
		log.Fatalf("Failed to open bazel workspace status file: %v\n", err)
	}

	var filtered []string
	for line := range strings.SplitSeq(string(statusFileContent), "\n") {
		if line == "" {
			continue
		}
		variable, value, ok := strings.Cut(line, " ")
		if !ok {
			log.Fatalf("Invalid line in status file: %q\n", line)
		}
		variable, ok = strings.CutPrefix(variable, "STABLE_")
		if ok && vars[variable] {
			filtered = append(filtered, fmt.Sprintf("STABLER_%s %s\n", variable, value))
		}
	}

	filteredContent := []byte(strings.Join(filtered, ""))
	err = os.WriteFile(*outFile, filteredContent, 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}
}
