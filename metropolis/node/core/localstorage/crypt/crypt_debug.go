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

package crypt

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// CryptMap implements a debug version of CryptMap from crypt.go. It aliases the given baseName device into name
// without any encryption.
func CryptMap(name string, baseName string, _ []byte) error {
	var stat unix.Stat_t
	if err := unix.Stat(baseName, &stat); err != nil {
		return fmt.Errorf("cannot stat base device: %w", err)
	}
	cryptDevName := fmt.Sprintf("/dev/%v", name)
	if err := unix.Mknod(cryptDevName, 0600|unix.S_IFBLK, int(stat.Rdev)); err != nil {
		return fmt.Errorf("failed to create crypt device node: %w", err)
	}
	return nil
}

// CryptInit implements a debug version of CryptInit from crypt.go. It aliases the given baseName device into name
// without any encryption. As an identity mapping doesn't need any initialization it doesn't do anything else.
func CryptInit(name, baseName string, encryptionKey []byte) error {
	return CryptMap(name, baseName, encryptionKey)
}
