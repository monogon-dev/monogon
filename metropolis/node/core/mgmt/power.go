package mgmt

import (
	"context"
	"os"
	"time"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apb "source.monogon.dev/metropolis/proto/api"
	"source.monogon.dev/osbase/efivarfs"
)

func (s *Service) Reboot(ctx context.Context, req *apb.RebootRequest) (*apb.RebootResponse, error) {
	var method int
	// Do not yet perform any system-wide actions here as the request might
	// still get rejected. There is another switch statement for that below.
	switch req.Type {
	case apb.RebootRequest_KEXEC:
		method = unix.LINUX_REBOOT_CMD_KEXEC
	case apb.RebootRequest_FIRMWARE:
		method = unix.LINUX_REBOOT_CMD_RESTART
	case apb.RebootRequest_POWER_OFF:
		method = unix.LINUX_REBOOT_CMD_POWER_OFF
	default:
		return nil, status.Error(codes.Unimplemented, "unimplemented type value")
	}
	switch req.NextBoot {
	case apb.RebootRequest_START_NORMAL:
	case apb.RebootRequest_START_ROLLBACK:
		if err := s.UpdateService.Rollback(); err != nil {
			return nil, status.Errorf(codes.Unavailable, "performing rollback failed: %v", err)
		}
	case apb.RebootRequest_START_FIRMWARE_UI:
		if req.Type == apb.RebootRequest_KEXEC {
			return nil, status.Error(codes.InvalidArgument, "START_FIRMWARE_UI cannot be used with KEXEC type")
		}
		supp, err := efivarfs.OSIndicationsSupported()
		if err != nil || supp&efivarfs.BootToFirmwareUI == 0 {
			return nil, status.Error(codes.Unimplemented, "Unable to boot into firmware UI on this platform")
		}
		if err := efivarfs.SetOSIndications(efivarfs.BootToFirmwareUI); err != nil {
			return nil, status.Errorf(codes.Unavailable, "Unable to set UEFI boot to UI indication: %v", err)
		}
	default:
		return nil, status.Error(codes.Unimplemented, "unimplemented next_boot value")
	}

	switch req.Type {
	case apb.RebootRequest_KEXEC:
		if err := s.UpdateService.KexecLoadNext(); err != nil {
			return nil, status.Errorf(codes.Unavailable, "failed to stage kexec kernel: %v", err)
		}
	case apb.RebootRequest_FIRMWARE:
		// Best-effort, if it fails this will still be a firmware reboot.
		os.WriteFile("/sys/kernel/reboot/mode", []byte("cold"), 0644)
	}
	s.initiateReboot(method)
	return &apb.RebootResponse{}, nil
}

func (s *Service) initiateReboot(method int) {
	s.LogTree.MustLeveledFor("root.mgmt").Warning("Reboot requested, rebooting in 2s")
	// TODO(#253): Tell Supervisor to shut down gracefully and reboot
	go func() {
		time.Sleep(2 * time.Second)
		unix.Unmount(s.UpdateService.ESPPath, 0)
		unix.Sync()
		s.disableNetworkInterfaces()
		unix.Reboot(method)
	}()
}

// For kexec it's recommended to disable all physical network interfaces
// before doing it. This function doesn't return any errors as it's best-
// effort anyways as we cannot reliably log the error anymore.
func (s *Service) disableNetworkInterfaces() {
	links, err := netlink.LinkList()
	if err != nil {
		return
	}
	for _, link := range links {
		d, ok := link.(*netlink.Device)
		if !ok {
			continue
		}
		if err := netlink.LinkSetDown(d); err != nil {
			s.LogTree.MustLeveledFor("root.mgmt").Errorf("Error taking link %q down: %v", link.Attrs().Name, err)
			continue
		}
	}
}
