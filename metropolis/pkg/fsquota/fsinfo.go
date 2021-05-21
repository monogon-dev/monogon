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

package fsquota

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

// This requires fsinfo() support, which is not yet in any stable kernel. Our
// kernel has that syscall backported. This would otherwise be an extremely
// expensive operation and also involve lots of logic from our side.

// From syscall_64.tbl
const sys_fsinfo = 441

// From uapi/linux/fsinfo.h
const fsinfo_attr_source = 0x09
const fsinfo_flags_query_path = 0x0000
const fsinfo_flags_query_fd = 0x0001

type fsinfoParams struct {
	resolveFlags uint64
	atFlags      uint32
	flags        uint32
	request      uint32
	nth          uint32
	mth          uint32
}

func fsinfoGetSource(dir *os.File) (string, error) {
	buf := make([]byte, 256)
	params := fsinfoParams{
		flags:   fsinfo_flags_query_fd,
		request: fsinfo_attr_source,
	}
	n, _, err := unix.Syscall6(sys_fsinfo, dir.Fd(), 0, uintptr(unsafe.Pointer(&params)), unsafe.Sizeof(params), uintptr(unsafe.Pointer(&buf[0])), 256)
	if err != unix.Errno(0) {
		return "", fmt.Errorf("failed to call fsinfo: %w", err)
	}
	return string(buf[:n-1]), nil
}
