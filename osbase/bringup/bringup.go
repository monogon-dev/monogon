// Package bringup implements a simple wrapper which configures all default
// mounts, logging and the corresponding forwarders to tty0 and ttyS0. It
// then configures a new logtree and starts a supervisor to run the provided
// supervisor.Runnable. Said Runnable is expected to never return. If it does,
// the supervisor will exit, an error will be printed and the system will
// reboot after five seconds.
package bringup

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.uber.org/multierr"
	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/bootparam"
	"source.monogon.dev/osbase/efivarfs"
	"source.monogon.dev/osbase/logtree"
	"source.monogon.dev/osbase/supervisor"
)

type Runnable supervisor.Runnable

func (r Runnable) Run() {
	// Pause execution on panic to require manual intervention.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Fatal error: %v\n", r)
			fmt.Printf("This node could not be started. Rebooting...\n")

			time.Sleep(5 * time.Second)
			unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)
		}
	}()

	if err := setupMounts(); err != nil {
		// We cannot do anything if we fail to mount.
		panic(err)
	}

	// Set up logger. Parse consoles from the kernel command line
	// as well as adding the two standard tty0/ttyS0 consoles.
	consoles := make(map[string]bool)
	cmdline, err := os.ReadFile("/proc/cmdline")
	if err == nil {
		params, _, err := bootparam.Unmarshal(string(cmdline))
		if err == nil {
			consoles = params.Consoles()
		}
	}
	consoles["tty0"] = true
	consoles["ttyS0"] = true

	lt := logtree.New()
	for consolePath := range consoles {
		f, err := os.OpenFile("/dev/"+consolePath, os.O_WRONLY, 0)
		if err != nil {
			continue
		}
		reader, err := lt.Read("", logtree.WithChildren(), logtree.WithStream())
		if err != nil {
			panic(fmt.Errorf("could not set up root log reader: %v", err))
		}
		go func() {
			for {
				p := <-reader.Stream
				fmt.Fprintf(f, "%s\n", p.String())
			}
		}()
	}

	sCtx, cancel := context.WithCancelCause(context.Background())

	// Don't reschedule the root runnable...
	supervisor.New(sCtx, func(ctx context.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("root runnable paniced: %v", r)
				cancel(err)
			}
		}()

		err = r(ctx)
		if err == nil {
			err = fmt.Errorf("root runnable exited without any error")
		}

		cancel(err)
		return nil
	}, supervisor.WithExistingLogtree(lt))

	<-sCtx.Done()
	panic(context.Cause(sCtx))
}

func mkdirAndMount(dir, fs string, flags uintptr) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("could not make %s: %w", dir, err)
	}
	if err := unix.Mount(fs, dir, fs, flags, ""); err != nil {
		return fmt.Errorf("could not mount %s on %s: %w", fs, dir, err)
	}
	return nil
}

// setupMounts sets up basic mounts like sysfs, procfs, devtmpfs and cgroups.
// This should be called early during init as a lot of processes depend on this
// being available.
func setupMounts() (err error) {
	// Set up target filesystems.
	for _, el := range []struct {
		dir   string
		fs    string
		flags uintptr
	}{
		{"/sys", "sysfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/sys/kernel/tracing", "tracefs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/sys/fs/pstore", "pstore", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/proc", "proc", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/dev/pts", "devpts", unix.MS_NOEXEC | unix.MS_NOSUID},
	} {
		err = multierr.Append(err, mkdirAndMount(el.dir, el.fs, el.flags))
	}

	// We try to mount efivarfs but ignore any error,
	// as we don't want to crash on non-EFI systems.
	_ = mkdirAndMount(efivarfs.Path, "efivarfs", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV)
	return
}
