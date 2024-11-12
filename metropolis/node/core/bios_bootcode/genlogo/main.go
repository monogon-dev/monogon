// SPDX-License-Identifier: Apache-2.0
// Copyright Monogon Project Authors

package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	input := flag.String("input", "", "")
	output := flag.String("output", "", "")
	flag.Parse()

	if *input == "" || *output == "" {
		log.Fatal("missing input or output flag")
	}

	inputFile, err := os.Open(*input)
	if err != nil {
		log.Fatal("Error opening image file:", err)
		return
	}
	defer inputFile.Close()

	img, err := png.Decode(inputFile)
	if err != nil {
		log.Fatal("Error decoding image:", err)
	}

	if img.Bounds().Dx() != 80 || img.Bounds().Dy() != 20 {
		log.Fatal("Image dimensions must be 80x20")
	}

	var linear []uint8
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y
			linear = append(linear, gray)
		}
	}

	// Perform RLE compression
	var rle []uint8
	for len(linear) > 0 {
		val := linear[0]
		l := uint8(1)
		for i := 1; i < len(linear); i++ {
			if linear[i] != val {
				break
			}
			l++
		}

		L := l
		for l > 0 {
			block := l
			if block > 127 {
				block = 127
			}
			rle = append(rle, (val<<7)|block)
			l -= block
		}
		linear = linear[L:]
	}

	rle = append(rle, 0)

	outputFile, err := os.Create(*output)
	if err != nil {
		log.Fatalf("failed creating output file: %v", err)
	}
	defer outputFile.Close()

	outputFile.WriteString("logo: db ")
	for i, r := range rle {
		if i > 0 {
			outputFile.WriteString(", ")
		}
		fmt.Fprintf(outputFile, "0x%02x", r)
	}
	outputFile.WriteString("\n")
}
