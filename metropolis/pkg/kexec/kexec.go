//go:build amd64 || arm64 || riscv64
// +build amd64 arm64 riscv64

// Package kexec allows executing subsequent kernels from Linux userspace.
package kexec

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/sys/unix"
)

// FileLoad loads the given kernel as the new kernel with the given initramfs
// and cmdline. The kernel can be started by calling
// unix.Reboot(unix.LINUX_REBOOT_CMD_KEXEC). The underlying syscall is only
// available on x86_64, arm64 and riscv.
// Parts of this function are taken from u-root's kexec package.
func FileLoad(kernel, initramfs *os.File, cmdline string) error {
	var flags int
	var initramfsfd int
	if initramfs != nil {
		initramfsfd = int(initramfs.Fd())
	} else {
		flags |= unix.KEXEC_FILE_NO_INITRAMFS
	}

	if err := unix.KexecFileLoad(int(kernel.Fd()), initramfsfd, cmdline, flags); err != nil {
		return fmt.Errorf("SYS_kexec_file_load(%d, %d, %s, %x) = %v", kernel.Fd(), initramfsfd, cmdline, flags, err)
	}
	runtime.KeepAlive(kernel)
	runtime.KeepAlive(initramfs)
	return nil
}
