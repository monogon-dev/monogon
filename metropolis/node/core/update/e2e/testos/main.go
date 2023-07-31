package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node/build/mkimage/osimage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/update"
	"source.monogon.dev/metropolis/pkg/blockdev"
	"source.monogon.dev/metropolis/pkg/gpt"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

var Variant = "U"

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
		{"/sys/firmware/efi/efivars", "efivarfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/proc", "proc", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/dev/pts", "devpts", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/tmp", "tmpfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
	} {
		if err := mkdirAndMount(el.dir, el.fs, el.flags); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := setupMounts(); err != nil {
		fmt.Printf("early init error, stopping: %v\n", err)
		unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
		return
	}
	lt := logtree.New()
	f, err := os.OpenFile("/dev/ttyS0", os.O_WRONLY, 0)
	if err != nil {
		fmt.Printf("early init error, stopping: %v\n", err)
		unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
		return
	}
	reader, err := lt.Read("", logtree.WithChildren(), logtree.WithStream())
	if err != nil {
		fmt.Printf("early init error, stopping: %v\n", err)
		unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
		return
	}

	sCtx := context.Background()
	supervisor.New(sCtx, testosRunnable, supervisor.WithExistingLogtree(lt))

	for {
		p := <-reader.Stream
		fmt.Fprintf(f, "%s\n", p.String())
	}
}

func testosRunnable(ctx context.Context) error {
	supervisor.Logger(ctx).Info("TESTOS_VARIANT=" + Variant)
	networkSvc := network.New(nil)
	networkSvc.DHCPVendorClassID = "dev.monogon.testos.v1"
	supervisor.Run(ctx, "networking", networkSvc.Run)

	vda, err := blockdev.Open("/dev/vda")
	if err != nil {
		return fmt.Errorf("unable to open root device: %w", err)
	}
	defer vda.Close()
	vdaParts, err := gpt.Read(vda)
	if err != nil {
		return fmt.Errorf("unable to read GPT from root device: %w", err)
	}

	updateSvc := update.Service{
		Logger: supervisor.MustSubLogger(ctx, "update"),
	}
	for pn, p := range vdaParts.Partitions {
		switch p.Type {
		case gpt.PartitionTypeEFISystem:
			if err := unix.Mount(fmt.Sprintf("/dev/vda%d", pn+1), "/esp", "vfat", unix.MS_SYNC, ""); err != nil {
				return fmt.Errorf("unable to mkdir ESP mountpoint: %w", err)
			}
			updateSvc.ProvideESP("/esp", uint32(pn+1), p)
		case osimage.SystemAType:
			if err := unix.Symlink(fmt.Sprintf("/dev/vda%d", pn+1), "/dev/system-a"); err != nil {
				return fmt.Errorf("failed to symlink system-a: %w", err)
			}
		case osimage.SystemBType:
			if err := unix.Symlink(fmt.Sprintf("/dev/vda%d", pn+1), "/dev/system-b"); err != nil {
				return fmt.Errorf("failed to symlink system-b: %w", err)
			}
		}
	}
	if err := updateSvc.MarkBootSuccessful(); err != nil {
		supervisor.Logger(ctx).Errorf("error marking boot successful: %w", err)
	}
	if Variant != "Z" {
		if err := updateSvc.InstallBundle(ctx, "http://10.42.0.5:80/bundle.bin"); err != nil {
			supervisor.Logger(ctx).Errorf("Error installing new bundle: %v", err)
		}
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Logger(ctx).Info("Installed bundle successfully, powering off")
	unix.Sync()
	time.Sleep(1 * time.Second)
	unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
	return nil
}
