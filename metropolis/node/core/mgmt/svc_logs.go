package mgmt

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/proto/api"
)

func (s *Service) Logs(_ *api.GetLogsRequest, _ api.NodeManagement_LogsServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}
