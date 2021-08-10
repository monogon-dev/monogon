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

// This package implements a command line tool that creates dm-verity hash
// images at a selected path, given an existing data image. The tool
// outputs a Verity mapping table on success.
//
// For more information, see:
// - source.monogon.dev/metropolis/pkg/verity
// - https://gitlab.com/cryptsetup/cryptsetup/wikis/DMVerity
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"source.monogon.dev/metropolis/pkg/verity"
)

// createHashImage creates a complete dm-verity hash image at
// hashImagePath. Contents of the file at dataImagePath are accessed
// read-only, hashed and written to the hash image in the process.
// The verity superblock is written only if wsb is true.
// It returns a string-convertible VerityMappingTable, or an error.
func createHashImage(dataImagePath, hashImagePath string, wsb bool) (*verity.MappingTable, error) {
	// Open the data image for reading.
	dataImage, err := os.Open(dataImagePath)
	if err != nil {
		return nil, fmt.Errorf("while opening the data image: %w", err)
	}
	defer dataImage.Close()
	// Create an empty hash image file.
	hashImage, err := os.OpenFile(hashImagePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("while opening the hash image for writing: %w", err)
	}
	defer hashImage.Close()

	// Write hashImage contents. Start with initializing a verity encoder,
	// seting hashImage as its output.
	v, err := verity.NewEncoder(hashImage, wsb)
	if err != nil {
		return nil, fmt.Errorf("while initializing a verity encoder: %w", err)
	}
	// Hash the contents of dataImage, block by block.
	_, err = io.Copy(v, dataImage)
	if err != nil {
		return nil, fmt.Errorf("while reading the data image: %w", err)
	}
	// The resulting hash tree won't be written until Close is called.
	err = v.Close()
	if err != nil {
		return nil, fmt.Errorf("while writing the hash image: %w", err)
	}

	// Return an encoder-generated verity mapping table, containing the salt
	// and the root hash.
	mt, err := v.MappingTable(dataImagePath, hashImagePath)
	if err != nil {
		return nil, fmt.Errorf("while querying for the mapping table: %w", err)
	}
	return mt, nil
}

// usage prints program usage information.
func usage(executable string) {
	fmt.Println("Usage: ", executable, " <data image> <hash image>")
}

func main() {
	if len(os.Args) != 3 {
		usage(os.Args[0])
		os.Exit(2)
	}
	dataImagePath := os.Args[1]
	hashImagePath := os.Args[2]

	// Attempt to build a new Verity hash Image at hashImagePath, based on
	// the data image at dataImagePath. Include the Verity superblock.
	mt, err := createHashImage(dataImagePath, hashImagePath, true)
	if err != nil {
		log.Fatal(err)
	}
	// Print a Device Mapper compatible mapping table.
	fmt.Println(mt)
}
