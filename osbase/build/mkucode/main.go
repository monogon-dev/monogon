// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// This assembles standalone microcode files into the format expected by the
// Linux microcode loader. See
// https://www.kernel.org/doc/html/latest/x86/microcode.html for further
// information.
package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/cavaliergopher/cpio"
	"google.golang.org/protobuf/encoding/prototext"

	"source.monogon.dev/osbase/build/mkucode/spec"
)

var (
	specPath = flag.String("spec", "", "Path to prototext specification (osbase.build.mkucode.UCode)")
	outPath  = flag.String("out", "", "Output path for cpio to be prepend to initrd")
)

// Usage: -spec <ucode.prototxt> -out <ucode.cpio>
func main() {
	flag.Parse()
	specRaw, err := os.ReadFile(*specPath)
	if err != nil {
		log.Fatalf("Failed to read spec: %v", err)
	}
	var ucodeSpec spec.UCode
	if err := prototext.Unmarshal(specRaw, &ucodeSpec); err != nil {
		log.Fatalf("Failed unmarshaling ucode spec: %v", err)
	}
	out, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("Failed to create cpio: %v", err)
	}
	defer out.Close()
	cpioWriter := cpio.NewWriter(out)
	for _, vendor := range ucodeSpec.Vendor {
		var totalSize int64
		for _, file := range vendor.File {
			data, err := os.Stat(file)
			if err != nil {
				log.Fatalf("Failed to stat file for vendor %q: %v", vendor.Id, err)
			}
			totalSize += data.Size()
		}
		if err := cpioWriter.WriteHeader(&cpio.Header{
			Mode: 0444,
			Name: "kernel/x86/microcode/" + vendor.Id + ".bin",
			Size: totalSize,
		}); err != nil {
			log.Fatalf("Failed to write cpio header for vendor %q: %v", vendor.Id, err)
		}
		for _, file := range vendor.File {
			f, err := os.Open(file)
			if err != nil {
				log.Fatalf("Failed to open file for vendor %q: %v", vendor.Id, err)
			}
			if _, err := io.Copy(cpioWriter, f); err != nil {
				log.Fatalf("Failed to copy data for file %q: %v", file, err)
			}
			f.Close()
		}
	}
	if err := cpioWriter.Close(); err != nil {
		log.Fatalf("Failed writing cpio: %v", err)
	}
}
