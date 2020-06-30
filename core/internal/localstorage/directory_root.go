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

package localstorage

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/sys/unix"

	"git.monogon.dev/source/nexantic.git/core/internal/localstorage/crypt"
)

func (r *Root) Start(ctx context.Context) error {
	r.Data.flagLock.Lock()
	defer r.Data.flagLock.Unlock()
	if r.Data.canMount {
		return fmt.Errorf("cannot re-start root storage")
	}
	// TODO(q3k): turn this into an Ensure call
	err := crypt.MakeBlockDevices(ctx)
	if err != nil {
		return fmt.Errorf("MakeBlockDevices: %w", err)
	}

	if err := os.Mkdir(r.ESP.FullPath(), 0755); err != nil {
		return fmt.Errorf("making ESP directory: %w", err)
	}

	if err := unix.Mount(crypt.ESPDevicePath, r.ESP.FullPath(), "vfat", unix.MS_NOEXEC|unix.MS_NODEV|unix.MS_SYNC, ""); err != nil {
		return fmt.Errorf("mounting ESP partition: %w", err)
	}

	r.Data.canMount = true

	return nil
}
