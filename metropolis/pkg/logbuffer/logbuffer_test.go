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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleLine(t *testing.T) {
	buf := New(1, 16000)
	buf.Write([]byte("Hello World\n"))
	out := buf.ReadLines(1)
	require.Len(t, out, 1, "Invalid number of lines read")
	require.Equal(t, "Hello World", out[0].Data, "Read bad log line")
	require.Equal(t, 11, out[0].OriginalLength, "Invalid line length")
}

func TestPartialWritesAndReads(t *testing.T) {
	buf := New(2, 16000)
	buf.Write([]byte("Hello "))
	buf.Write([]byte("World\nTest "))
	buf.Write([]byte("2\n"))

	out := buf.ReadLines(1)
	require.Len(t, out, 1, "Invalid number of lines for partial read")
	require.Equal(t, "Test 2", out[0].Data, "Read bad log line")

	out2 := buf.ReadLines(2)
	require.Len(t, out2, 2, "Invalid number of lines read")
	require.Equal(t, "Hello World", out2[0].Data, "Read bad log line")
	require.Equal(t, "Test 2", out2[1].Data, "Read bad log line")
}

func TestBufferOverwrite(t *testing.T) {
	buf := New(3, 16000)
	buf.Write([]byte("Test1\nTest2\nTest3\nTest4\n"))

	out := buf.ReadLines(3)
	require.Equal(t, out[0].Data, "Test2", "Read bad log line")
	require.Equal(t, out[1].Data, "Test3", "Read bad log line")
	require.Equal(t, out[2].Data, "Test4", "Overwritten data is invalid")
}

func TestTooLargeRequests(t *testing.T) {
	buf := New(1, 16000)
	outEmpty := buf.ReadLines(1)
	require.Len(t, outEmpty, 0, "Returned more data than there is")

	buf.Write([]byte("1\n2\n"))
	out := buf.ReadLines(2)
	require.Len(t, out, 1, "Returned more data than the ring buffer can hold")
}

func TestSpecialCases(t *testing.T) {
	buf := New(2, 16000)
	buf.Write([]byte("Test1"))
	buf.Write([]byte("\nTest2\n"))
	out := buf.ReadLines(2)
	require.Len(t, out, 2, "Too many lines written")
	require.Equal(t, out[0].Data, "Test1", "Read bad log line")
	require.Equal(t, out[1].Data, "Test2", "Read bad log line")
}

func TestLineLengthLimit(t *testing.T) {
	buf := New(2, 6)

	testStr := "Just Testing"

	buf.Write([]byte(testStr + "\nShort\n"))

	out := buf.ReadLines(2)
	require.Equal(t, len(testStr), out[0].OriginalLength, "Line is over length limit")
	require.Equal(t, "Just T", out[0].Data, "Log line not properly truncated")

	out2 := buf.ReadLinesTruncated(2, "...")
	require.Equal(t, out2[0], "Just T...", "Line is over length limit")
	require.Equal(t, out2[1], "Short", "Truncated small enough line")
}
