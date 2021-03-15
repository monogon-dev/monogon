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
	"fmt"
	"io"
	"log"
	"os"

	"source.monogon.dev/metropolis/pkg/logbuffer"
)

// prefixedStdout is a os.Stdout proxy that prefixes every line with a constant
// prefix. This is used to show logs from two Metropolis nodes without getting
// them confused.
// TODO(q3k): move to logging API instead of relying on qemu stdout, and remove
// this function.
func prefixedStdout(prefix string) io.ReadWriter {
	lb := logbuffer.NewLineBuffer(2048, func(l *logbuffer.Line) {
		fmt.Fprintf(os.Stdout, "%s%s\n", prefix, l.Data)
	})
	// Make a ReaderWriter from LineBuffer (a Reader), by combining into an
	// anonymous struct with a io.MultiReader() (which will always return EOF
	// on every Read if given no underlying readers).
	return struct {
		io.Reader
		io.Writer
	}{
		Reader: io.MultiReader(),
		Writer: lb,
	}
}

func main() {
	log.Fatal("unimplemented")
}
