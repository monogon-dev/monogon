// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package bringup

// This uses the unstable overrideWrite interface to also emit all runtime
// writes to a dedicated runtime file descriptor to catch and debug crash dumps.
// See https://go-review.googlesource.com/c/go/+/278792 for details about the
// interface. This interface is relatively special, refrain from using most Go
// features in here as it might cause unexpected behavior. Especially yielding
// is a bad idea as the scheduler might be in an inconsistent state. But using
// this interface was judged to be vastly more maintenance-friendly than
// attempting to parse out this information from a combined stderr.

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/logtree"
)

// This hooks into a global variable which is checked by runtime.write and used
// instead of runtime.write1 if populated.
//
//go:linkname overrideWrite runtime.overrideWrite
var overrideWrite func(fd uintptr, p unsafe.Pointer, n int32) int32

// Contains the files into which runtime logs and crashes are written.
var runtimeFds []uintptr

// This is essentially a reimplementation of the assembly function
// runtime.write1, just with a hardcoded file descriptor and using the assembly
// function syscall.RawSyscall to not get a dependency on Go's calling
// convention and needing an implementation for every architecture.
//
// This has to use the syscall package as the x/sys/unix implementation
// has a broken trampoline that calls the abi0 variant which is instrumented
// for some unknown reason.
//
//go:norace
//go:nosplit
func runtimeWrite(fd uintptr, p unsafe.Pointer, n int32) int32 {
	// Only redirect writes to stderr.
	if fd != 2 {
		a, _, err := syscall.RawSyscall(syscall.SYS_WRITE, fd, uintptr(p), uintptr(n))
		if err == 0 {
			return int32(a)
		}
		return int32(err)
	}

	// Write to the runtime panic FDs.
	for _, f := range runtimeFds {
		_, _, _ = syscall.RawSyscall(syscall.SYS_WRITE, f, uintptr(p), uintptr(n))
	}

	// Finally, write to original FD
	a, _, err := syscall.RawSyscall(syscall.SYS_WRITE, fd, uintptr(p), uintptr(n))
	if err == 0 {
		return int32(a)
	}
	return int32(err)
}

func initPanicHandler(lt *logtree.LogTree, consoles []*console) {
	l := lt.MustLeveledFor("panichandler")

	// Setup pstore userspace message buffer
	fd, err := unix.Open("/dev/pmsg0", os.O_WRONLY|os.O_SYNC, 0)
	if err != nil {
		l.Errorf("Failed to open pstore userspace device (pstore probably unavailable): %v", err)
		l.Warningf("Continuing without persistent panic storage.")
	} else {
		runtimeFds = append(runtimeFds, uintptr(fd))
	}

	for _, c := range consoles {
		fd, err := unix.Open(c.path, os.O_WRONLY, 0)
		if err != nil {
			l.Errorf("Failed opening panic console: %s", err)
			continue
		}

		runtimeFds = append(runtimeFds, uintptr(fd))
		l.Infof("Panic console: %s", c.path)
	}

	// This could cause a data race if the runtime crashed while we're
	// initializing the crash handler, but there is no locking infrastructure
	// for this so we have to take that risk.
	overrideWrite = runtimeWrite
}
