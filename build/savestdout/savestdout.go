// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s output_file program <args...>", os.Args[0])
	}

	f, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("Create(%q): %v", os.Args[1], err)
	}
	defer f.Close()

	args := os.Args[3:]
	cmd := exec.Command(os.Args[2], args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = f

	err = cmd.Run()
	if err == nil {
		return
	}

	var e *exec.ExitError
	if errors.As(err, &e) {
		os.Exit(e.ExitCode())
	}

	log.Fatalf("Could not start command: %v", err)
}
