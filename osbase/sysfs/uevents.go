// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
