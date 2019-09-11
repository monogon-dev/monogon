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

package sysfs

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadUevents(filename string) (map[string]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	ueventMap := make(map[string]string)
	reader := bufio.NewReader(f)
	for {
		name, err := reader.ReadString(byte('='))
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		value, err := reader.ReadString(byte('\n'))
		if err == io.EOF {
			continue
		} else if err != nil {
			return nil, err
		}
		ueventMap[strings.Trim(name, "=")] = strings.TrimSpace(value)
	}
	return ueventMap, nil
}
