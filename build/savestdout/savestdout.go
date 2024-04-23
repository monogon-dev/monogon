// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
