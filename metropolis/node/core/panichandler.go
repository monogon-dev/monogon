// This uses the unstable overrideWrite interface to also emit all runtime
// writes to a dedicated runtime file descriptor to catch and debug crash dumps.
// See https://go-review.googlesource.com/c/go/+/278792 for details about the
// interface. This interface is relatively special, refrain from using most Go
// features in here as it might cause unexpected behavior. Especially yielding
// is a bad idea as the scheduler might be in an inconsistent state. But using
// this interface was judged to be vastly more maintenance-friendly than
// attempting to parse out this information from a combined stderr.
package main

import (
	"io"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/pkg/logtree"
)

// This hooks into a global variable which is checked by runtime.write and used
// instead of runtime.write1 if populated.
//go:linkname overrideWrite runtime.overrideWrite
var overrideWrite func(fd uintptr, p unsafe.Pointer, n int32) int32

// Contains the file into which runtime logs and crashes are written.
var runtimeFd os.File

// This is essentially a reimplementation of the assembly function
// runtime.write1, just with a hardcoded file descriptor and using the assembly
// function unix.RawSyscall to not get a dependency on Go's calling convention
// and needing an implementation for every architecture.
//go:nosplit
func runtimeWrite(fd uintptr, p unsafe.Pointer, n int32) int32 {
	_, _, err := unix.RawSyscall(unix.SYS_WRITE, runtimeFd.Fd(), uintptr(p), uintptr(n))
	if err != 0 {
		return int32(err)
	}
	// Also write to original FD
	_, _, err = unix.RawSyscall(unix.SYS_WRITE, fd, uintptr(p), uintptr(n))
	return int32(err)
}

const runtimeLogPath = "/esp/core_runtime.log"

func initPanicHandler(lt *logtree.LogTree) {
	rl := lt.MustRawFor("panichandler")
	l := lt.MustLeveledFor("panichandler")
	runtimeLogFile, err := os.Open(runtimeLogPath)
	if err != nil && !os.IsNotExist(err) {
		l.Errorf("Failed to open runtimeLogFile: %v", err)
	}
	if err == nil {
		if _, err := io.Copy(rl, runtimeLogFile); err != nil {
			l.Errorf("Failed to log old persistent crash: %v", err)
		}
		runtimeLogFile.Close()
		if err := os.Remove(runtimeLogPath); err != nil {
			l.Errorf("Failed to delete old persistent runtime crash log: %v", err)
		}
	}

	file, err := os.Create(runtimeLogPath)
	if err != nil {
		l.Errorf("Failed to open core runtime log file: %w", err)
		l.Warningf("Continuing without persistent panic storage.")
		return
	}
	runtimeFd = *file
	// Make sure the Fd is in blocking mode. Go's runtime opens all FDs in non-
	// blocking mode by default and switches them back once you get a reference
	// to the raw file descriptor to not break existing code. This switching
	// back is done on the first Fd() call and involves calls into the runtime
	// scheduler as it issues non-raw syscalls. Calling Fd() here makes sure
	// that these calls happen in a sane environment before any actual panic.
	// After this Fd() performs only memory accesses which is safe even when
	// panicing the runtime.
	// Keeping the raw fd is not possible as Go's runtime would eventually
	// garbage-collect the backing os.File and close it, so we must keep around
	// the actual os.File.
	_ = runtimeFd.Fd()
	// This could cause a data race if the runtime crashed while we're
	// initializing the crash handler, but there is no locking infrastructure
	// for this so we have to take that risk.
	overrideWrite = runtimeWrite
	return
}
