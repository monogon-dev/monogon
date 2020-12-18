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

package logbuffer

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	apb "git.monogon.dev/source/nexantic.git/metropolis/proto/api"
)

// Line is a line stored in the log buffer - a string, that has been perhaps truncated (due to exceeded limits).
type Line struct {
	Data           string
	OriginalLength int
}

// Truncated returns whether this line has been truncated to fit limits.
func (l *Line) Truncated() bool {
	return l.OriginalLength > len(l.Data)
}

// String returns the line with an ellipsis at the end (...) if the line has been truncated, or the original line
// otherwise.
func (l *Line) String() string {
	if l.Truncated() {
		return l.Data + "..."
	}
	return l.Data
}

// ProtoLog returns a Logging-specific protobuf structure.
func (l *Line) ProtoLog() *apb.LogEntry_Raw {
	return &apb.LogEntry_Raw{
		Data:           l.Data,
		OriginalLength: int64(l.OriginalLength),
	}
}

// LineFromLogProto converts a Logging-specific protobuf message back into a Line.
func LineFromLogProto(raw *apb.LogEntry_Raw) (*Line, error) {
	if raw.OriginalLength < int64(len(raw.Data)) {
		return nil, fmt.Errorf("original_length smaller than length of data")
	}
	originalLength := int(raw.OriginalLength)
	if int64(originalLength) < raw.OriginalLength {
		return nil, fmt.Errorf("original_length larger than native int size")
	}
	return &Line{
		Data:           raw.Data,
		OriginalLength: originalLength,
	}, nil
}

// LineBuffer is a io.WriteCloser that will call a given callback every time a line is completed.
type LineBuffer struct {
	maxLineLength int
	cb            LineBufferCallback

	mu  sync.Mutex
	cur strings.Builder
	// length is the length of the line currently being written - this will continue to increase, even if the string
	// exceeds maxLineLength.
	length int
	closed bool
}

// LineBufferCallback is a callback that will get called any time the line is completed. The function must not cause another
// write to the LineBuffer, or the program will deadlock.
type LineBufferCallback func(*Line)

// NewLineBuffer creates a new LineBuffer with a given line length limit and callback.
func NewLineBuffer(maxLineLength int, cb LineBufferCallback) *LineBuffer {
	return &LineBuffer{
		maxLineLength: maxLineLength,
		cb:            cb,
	}
}

// writeLimited writes to the internal buffer, making sure that its size does not exceed the maxLineLength.
func (l *LineBuffer) writeLimited(data []byte) {
	l.length += len(data)
	if l.cur.Len()+len(data) > l.maxLineLength {
		data = data[:l.maxLineLength-l.cur.Len()]
	}
	l.cur.Write(data)
}

// comitLine calls the callback and resets the builder.
func (l *LineBuffer) commitLine() {
	l.cb(&Line{
		Data:           l.cur.String(),
		OriginalLength: l.length,
	})
	l.cur.Reset()
	l.length = 0
}

func (l *LineBuffer) Write(data []byte) (int, error) {
	var pos = 0

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.closed {
		return 0, fmt.Errorf("closed")
	}

	for {
		nextNewline := bytes.IndexRune(data[pos:], '\n')

		// No newline in the data, write everything to the current line
		if nextNewline == -1 {
			l.writeLimited(data[pos:])
			break
		}

		// Write this line and update position
		l.writeLimited(data[pos : pos+nextNewline])
		l.commitLine()
		pos += nextNewline + 1

		// Data ends with a newline, stop now without writing an empty line
		if nextNewline == len(data)-1 {
			break
		}
	}
	return len(data), nil
}

// Close will emit any leftover data in the buffer to the callback. Subsequent calls to Write will fail. Subsequent calls to Close
// will also fail.
func (l *LineBuffer) Close() error {
	if l.closed {
		return fmt.Errorf("already closed")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.closed = true
	if l.length > 0 {
		l.commitLine()
	}
	return nil
}
