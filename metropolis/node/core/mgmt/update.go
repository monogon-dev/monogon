package mgmt

import (
	"context"
	"time"

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
		// TODO(#253): Tell Supervisor to shut down gracefully and reboot
		go func() {
			time.Sleep(10 * time.Second)
			unix.Sync()
			if req.ActivationMode == apb.ActivationMode_ACTIVATION_KEXEC {
				unix.Reboot(unix.LINUX_REBOOT_CMD_KEXEC)
			} else {
				unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)
			}
		}()
	}

	return &apb.UpdateNodeResponse{}, nil
}
