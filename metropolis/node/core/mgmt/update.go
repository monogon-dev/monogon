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
	if err := s.UpdateService.InstallBundle(ctx, req.BundleUrl); err != nil {
		return nil, status.Errorf(codes.Unavailable, "error installing update: %v", err)
	}
	if !req.NoReboot {
		// TODO(#253): Tell Supervisor to shut down gracefully and reboot
		go func() {
			time.Sleep(10 * time.Second)
			unix.Sync()
			unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)
		}()
	}
	return &apb.UpdateNodeResponse{}, nil
}
