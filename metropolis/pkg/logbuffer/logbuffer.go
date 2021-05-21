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

// Package logbuffer implements a fixed-size in-memory ring buffer for
// line-separated logs. It implements io.Writer and splits the data into lines.
// The lines are kept in a ring where the oldest are overwritten once it's
// full. It allows retrieval of the last n lines. There is a built-in line
// length limit to bound the memory usage at maxLineLength * size.
package logbuffer

import (
	"sync"
)

// LogBuffer implements a fixed-size in-memory ring buffer for line-separated logs
type LogBuffer struct {
	mu      sync.RWMutex
	content []Line
	length  int
	*LineBuffer
}

// New creates a new LogBuffer with a given ringbuffer size and maximum line
// length.
func New(size, maxLineLength int) *LogBuffer {
	lb := &LogBuffer{
		content: make([]Line, size),
	}
	lb.LineBuffer = NewLineBuffer(maxLineLength, lb.lineCallback)
	return lb
}

func (b *LogBuffer) lineCallback(line *Line) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.content[b.length%len(b.content)] = *line
	b.length++
}

// capToContentLength caps the number of requested lines to what is actually
// available
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

// ReadLines reads the last n lines from the buffer in chronological order. If
// n is bigger than the ring buffer or the number of available lines only the
// number of stored lines are returned.
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

// ReadLinesTruncated works exactly the same as ReadLines, but adds an ellipsis
// at the end of every line that was truncated because it was over
// MaxLineLength
func (b *LogBuffer) ReadLinesTruncated(n int, ellipsis string) []string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	// This does not use ReadLines() to prevent excessive reference copying and
	// associated GC pressure since it could process a lot of lines.

	n = b.capToContentLength(n)

	outArray := make([]string, n)
	for i := 1; i <= n; i++ {
		line := b.content[(b.length-i)%len(b.content)]
		outArray[n-i] = line.String()
	}
	return outArray
}
