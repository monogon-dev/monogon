package mgmt

import (
	"context"
	"time"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apb "source.monogon.dev/metropolis/proto/api"
)

func (s *Service) UpdateNode(ctx context.Context, req *apb.UpdateNodeRequest) (*apb.UpdateNodeResponse, error) {
	ok := s.updateMutex.TryLock()
	if ok {
		defer s.updateMutex.Unlock()
	} else {
		return nil, status.Error(codes.Aborted, "another UpdateNode RPC is in progress on this node")
	}
	if req.ActivationMode == apb.ActivationMode_ACTIVATION_INVALID {
		return nil, status.Errorf(codes.InvalidArgument, "activation_mode needs to be explicitly specified")
	}
	if err := s.UpdateService.InstallBundle(ctx, req.BundleUrl, req.ActivationMode == apb.ActivationMode_ACTIVATION_KEXEC); err != nil {
		return nil, status.Errorf(codes.Unavailable, "error installing update: %v", err)
	}
	if req.ActivationMode != apb.ActivationMode_ACTIVATION_NONE {

		methodString, method := "reboot", unix.LINUX_REBOOT_CMD_RESTART
		if req.ActivationMode == apb.ActivationMode_ACTIVATION_KEXEC {
			methodString = "kexec"
			method = unix.LINUX_REBOOT_CMD_KEXEC
		}

		s.LogTree.MustLeveledFor("update").Infof("activating update with method: %s", methodString)

		go func() {
			// TODO(#253): Tell Supervisor to shut down gracefully and reboot
			time.Sleep(10 * time.Second)
			s.LogTree.MustLeveledFor("update").Info("activating now...")
			unix.Unmount(s.UpdateService.ESPPath, 0)
			unix.Sync()
			disableNetworkInterfaces()
			unix.Reboot(method)
		}()
	}

	return &apb.UpdateNodeResponse{}, nil
}

// For kexec it's recommended to disable all physical network interfaces
// before doing it. This function doesn't return any errors as it's best-
// effort anyways as we cannot reliably log the error anymore.
func disableNetworkInterfaces() {
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
			continue
		}
	}
}
