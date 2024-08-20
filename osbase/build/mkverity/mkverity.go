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
// - source.monogon.dev/osbase/verity
// - https://gitlab.com/cryptsetup/cryptsetup/wikis/DMVerity
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"source.monogon.dev/osbase/verity"
)

// createImage creates a dm-verity target image by combining the input image
// with Verity metadata. Contents of the data image are copied to the output
// image. Then, the same contents are verity-encoded and appended to the
// output image. The verity superblock is written only if wsb is true. It
// returns either a dm-verity target table, or an error.
func createImage(dataImagePath, outputImagePath string, wsb bool) (*verity.MappingTable, error) {
	// Hardcode both the data block size and the hash block size as 4096 bytes.
	bs := uint32(4096)

	// Open the data image for reading.
	dataImage, err := os.Open(dataImagePath)
	if err != nil {
		return nil, fmt.Errorf("while opening the data image: %w", err)
	}
	defer dataImage.Close()

	// Check that the data image is well-formed.
	ds, err := dataImage.Stat()
	if err != nil {
		return nil, fmt.Errorf("while stat-ing the data image: %w", err)
	}
	if !ds.Mode().IsRegular() {
		return nil, fmt.Errorf("the data image must be a regular file")
	}
	if ds.Size()%int64(bs) != 0 {
		return nil, fmt.Errorf("the data image must end on a %d-byte block boundary", bs)
	}

	// Create an empty hash image file.
	outputImage, err := os.OpenFile(outputImagePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return nil, fmt.Errorf("while opening the hash image for writing: %w", err)
	}
	defer outputImage.Close()

	// Copy the input data into the output file, then rewind dataImage to be read
	// again by the Verity encoder.
	_, err = io.Copy(outputImage, dataImage)
	if err != nil {
		return nil, err
	}
	_, err = dataImage.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}

	// Write outputImage contents. Start with initializing a verity encoder,
	// seting outputImage as its output.
	v, err := verity.NewEncoder(outputImage, bs, bs, wsb)
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

	// Return an encoder-generated verity mapping table, containing the salt and
	// the root hash. First, calculate the starting hash block by dividing the
	// data image size by the encoder data block size.
	hashStart := ds.Size() / int64(bs)
	mt, err := v.MappingTable(dataImagePath, outputImagePath, hashStart)
	if err != nil {
		return nil, fmt.Errorf("while querying for the mapping table: %w", err)
	}
	return mt, nil
}

var (
	input           = flag.String("input", "", "input disk image (required)")
	output          = flag.String("output", "", "output disk image with Verity metadata appended (required)")
	dataDeviceAlias = flag.String("data_alias", "", "data device alias used in the mapping table")
	hashDeviceAlias = flag.String("hash_alias", "", "hash device alias used in the mapping table")
	table           = flag.String("table", "", "a file the mapping table will be saved to; disables stdout")
)

func main() {
	flag.Parse()

	// Ensure that required parameters were provided before continuing.
	if *input == "" {
		log.Fatalf("-input must be set.")
	}
	if *output == "" {
		log.Fatalf("-output must be set.")
	}

	// Build the image.
	mt, err := createImage(*input, *output, false)
	if err != nil {
		log.Fatal(err)
	}

	// Patch the device names, if alternatives were provided.
	if *dataDeviceAlias != "" {
		mt.DataDevicePath = *dataDeviceAlias
	}
	if *hashDeviceAlias != "" {
		mt.HashDevicePath = *hashDeviceAlias
	}

	// Print a DeviceMapper target table, or save it to a file, if the table
	// parameter was specified.
	if *table != "" {
		if err := os.WriteFile(*table, []byte(mt.String()), 0644); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(mt)
	}
}
