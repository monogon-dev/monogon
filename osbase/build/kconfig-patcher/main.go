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
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	inPath  = flag.String("in", "", "Path to input Kconfig")
	outPath = flag.String("out", "", "Path to output Kconfig")
)

func main() {
	flag.Parse()
	if *inPath == "" || *outPath == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}
	inFile, err := os.Open(*inPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input Kconfig: %v\n", err)
		os.Exit(1)
	}
	outFile, err := os.Create(*outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output Kconfig: %v\n", err)
		os.Exit(1)
	}
	var config struct {
		Overrides map[string]string `json:"overrides"`
	}
	if err := json.Unmarshal([]byte(flag.Arg(0)), &config); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse overrides: %v\n", err)
		os.Exit(1)
	}
	err = patchKconfig(inFile, outFile, config.Overrides)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to patch: %v\n", err)
		os.Exit(1)
	}
}

func patchKconfig(inFile io.Reader, outFile io.Writer, overrides map[string]string) error {
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return scanner.Err()
		}
		line := scanner.Text()
		cleanLine := strings.TrimSpace(line)
		if strings.HasPrefix(cleanLine, "#") || cleanLine == "" {
			// Pass through comments and empty lines
			fmt.Fprintln(outFile, line)
		} else {
			// Line contains a configuration option
			parts := strings.SplitN(line, "=", 2)
			keyName := parts[0]
			if overrideVal, ok := overrides[strings.TrimSpace(keyName)]; ok {
				// Override it
				if overrideVal == "" {
					fmt.Fprintf(outFile, "# %v is not set\n", keyName)
				} else {
					fmt.Fprintf(outFile, "%v=%v\n", keyName, overrideVal)
				}
				delete(overrides, keyName)
			} else {
				// Pass through unchanged
				fmt.Fprintln(outFile, line)
			}
		}
	}
	// Process left over overrides
	for key, val := range overrides {
		fmt.Fprintf(outFile, "%v=%v\n", key, val)
	}
	return nil
}
