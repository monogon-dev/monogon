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

// Package logbuffer implements a fixed-size in-memory ring buffer for line-separated logs.
// It implements io.Writer and splits the data into lines. The lines are kept in a ring where the
// oldest are overwritten once it's full. It allows retrieval of the last n lines. There is a built-in
// line length limit to bound the memory usage at maxLineLength * size.
package logbuffer

import (
	"bytes"
	"strings"
	"sync"
)

// LogBuffer implements a fixed-size in-memory ring buffer for line-separated logs
type LogBuffer struct {
	mu            sync.RWMutex
	maxLineLength int
	content       []Line
	length        int

	currentLineBuilder       strings.Builder
	currentLineWrittenLength int
}

type Line struct {
	Data           string
	OriginalLength int
}

func New(size, maxLineLength int) *LogBuffer {
	return &LogBuffer{
		content:       make([]Line, size),
		maxLineLength: maxLineLength,
	}
}

func (b *LogBuffer) writeLimited(newData []byte) {
	builder := &b.currentLineBuilder
	b.currentLineWrittenLength += len(newData)
	if builder.Len()+len(newData) > b.maxLineLength {
		builder.Write(newData[:b.maxLineLength-builder.Len()])
	} else {
		builder.Write(newData)
	}
}

func (b *LogBuffer) commitLine() {
	b.content[b.length%len(b.content)] = Line{
		Data:           b.currentLineBuilder.String(),
		OriginalLength: b.currentLineWrittenLength}
	b.length++
	b.currentLineBuilder.Reset()
	b.currentLineWrittenLength = 0
}

func (b *LogBuffer) Write(data []byte) (int, error) {
	var pos int = 0

	b.mu.Lock()
	defer b.mu.Unlock()

	for {
		nextNewline := bytes.IndexRune(data[pos:], '\n')

		// No newline in the data, write everything to the current line
		if nextNewline == -1 {
			b.writeLimited(data[pos:])
			break
		}

		// Write this line and update position
		b.writeLimited(data[pos : pos+nextNewline])
		b.commitLine()
		pos += nextNewline + 1

		// Data ends with a newline, stop now without writing an empty line
		if nextNewline == len(data)-1 {
			break
		}
	}
	return len(data), nil
}

// capToContentLength caps the number of requested lines to what is actually available
func (b *LogBuffer) capToContentLength(n int) int {
	// If there aren't enough lines to read, reduce the request size
	if n > b.length {
		n = b.length
	}
	// If there isn't enough ringbuffer space, reduce the request size
	if n > len(b.content) {
		n = len(b.content)
	}
	return n
}

// ReadLines reads the last n lines from the buffer in chronological order. If n is bigger than the
// ring buffer or the number of available lines only the number of stored lines are returned.
func (b *LogBuffer) ReadLines(n int) []Line {
	b.mu.RLock()
	defer b.mu.RUnlock()

	n = b.capToContentLength(n)

	// Copy references out to keep them around
	outArray := make([]Line, n)
	for i := 1; i <= n; i++ {
		outArray[n-i] = b.content[(b.length-i)%len(b.content)]
	}
	return outArray
}

// ReadLinesTruncated works exactly the same as ReadLines, but adds an ellipsis at the end of every
// line that was truncated because it was over MaxLineLength
func (b *LogBuffer) ReadLinesTruncated(n int, ellipsis string) []string {
	// This does not use ReadLines() to prevent excessive reference copying and associated GC pressure
	// since it could process a lot of lines.

	n = b.capToContentLength(n)

	outArray := make([]string, n)
	for i := 1; i <= n; i++ {
		line := b.content[(b.length-i)%len(b.content)]
		if line.OriginalLength > b.maxLineLength {
			outArray[n-i] = line.Data + ellipsis
		} else {
			outArray[n-i] = line.Data
		}
	}
	return outArray
}
