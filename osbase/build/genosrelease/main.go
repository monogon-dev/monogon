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

// genosrelease provides rudimentary support to generate os-release files
// following the freedesktop spec from arguments and stamping
//
// https://www.freedesktop.org/software/systemd/man/os-release.html
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	flagStatusFile = flag.String("status_file", "", "path to bazel workspace status file")
	flagOutFile    = flag.String("out_file", "os-release", "path to os-release output file")
	flagStampVar   = flag.String("stamp_var", "", "variable to use as version from the workspace status file")
	flagName       = flag.String("name", "", "name parameter (see freedesktop spec)")
	flagID         = flag.String("id", "", "id parameter (see freedesktop spec)")
)

func main() {
	flag.Parse()
	statusFileContent, err := os.ReadFile(*flagStatusFile)
	if err != nil {
		fmt.Printf("Failed to open bazel workspace status file: %v\n", err)
		os.Exit(1)
	}
	statusVars := make(map[string]string)
	for _, line := range strings.Split(string(statusFileContent), "\n") {
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		statusVars[parts[0]] = parts[1]
	}

	version, ok := statusVars[*flagStampVar]
	if !ok {
		fmt.Printf("%v key not set in bazel workspace status file\n", *flagStampVar)
		os.Exit(1)
	}
	// As specified by https://www.freedesktop.org/software/systemd/man/os-release.html
	osReleaseVars := map[string]string{
		"NAME":        *flagName,
		"ID":          *flagID,
		"VERSION":     version,
		"VERSION_ID":  version,
		"PRETTY_NAME": *flagName + " " + version,
	}
	osReleaseContent, err := godotenv.Marshal(osReleaseVars)
	if err != nil {
		fmt.Printf("Failed to encode os-release file: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*flagOutFile, []byte(osReleaseContent+"\n"), 0644); err != nil {
		fmt.Printf("Failed to write os-release file: %v\n", err)
		os.Exit(1)
	}
}
