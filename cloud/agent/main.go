package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

func main() {
	setupMounts()

	// Set up logger for the Agent. Currently logs everything to /dev/tty0 and
	// /dev/ttyS0.
	consoles := []string{"/dev/tty0", "/dev/ttyS0"}
	lt := logtree.New()
	for _, p := range consoles {
		f, err := os.OpenFile(p, os.O_WRONLY, 0)
		if err != nil {
			continue
		}
		reader, err := lt.Read("", logtree.WithChildren(), logtree.WithStream())
		if err != nil {
			panic(fmt.Errorf("could not set up root log reader: %v", err))
		}
		go func(path string, f io.Writer) {
			for {
				p := <-reader.Stream
				fmt.Fprintf(f, "%s\n", p.String())
			}
		}(p, f)
	}

	sCtx := context.Background()
	supervisor.New(sCtx, agentRunnable, supervisor.WithExistingLogtree(lt))
	select {}
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
func setupMounts() error {
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
		if err := mkdirAndMount(el.dir, el.fs, el.flags); err != nil {
			return err
		}
	}
	return nil
}
