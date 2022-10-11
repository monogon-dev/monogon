package server

import (
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
	apb "source.monogon.dev/cloud/bmaas/server/api"
	"source.monogon.dev/metropolis/node/core/rpc"
)

type agentCallbackService struct {
	s *Server
}

var (
	errAgentUnauthenticated = errors.New("machine id or public key unknown")
)

func (a *agentCallbackService) Heartbeat(ctx context.Context, req *apb.AgentHeartbeatRequest) (*apb.AgentHeartbeatResponse, error) {
	// Extract ED25519 self-signed certificate from client connection.
	cert, err := rpc.GetPeerCertificate(ctx)
	if err != nil {
		return nil, err
	}
	pk := cert.PublicKey.(ed25519.PublicKey)
	machineId, err := uuid.Parse(req.MachineId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "machine_id invalid")
	}

	// TODO(q3k): don't start a session for every RPC.
	session, err := a.s.bmdb.StartSession(ctx)
	if err != nil {
		klog.Errorf("Could not start session: %v", err)
		return nil, status.Error(codes.Unavailable, "could not start session")
	}

	// Verify that machine ID and connection public key match up to a machine in the
	// BMDB. Prevent leaking information about a machine's existence to unauthorized
	// agents.
	err = session.Transact(ctx, func(q *model.Queries) error {
		agents, err := q.AuthenticateAgentConnection(ctx, model.AuthenticateAgentConnectionParams{
			MachineID:      machineId,
			AgentPublicKey: pk,
		})
		if err != nil {
			return fmt.Errorf("AuthenticateAgentConnection: %w", err)
		}
		if len(agents) < 1 {
			return errAgentUnauthenticated
		}
		return nil
	})
	if err != nil {
		if errors.Is(err, errAgentUnauthenticated) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		klog.Errorf("Could not authenticate agent: %v", err)
		return nil, status.Error(codes.Unavailable, "could not authenticate agent")
	}

	// Request is now authenticated.

	// Serialize hardware report if submitted alongside heartbeat.
	var hwraw []byte
	if req.HardwareReport != nil {
		hwraw, err = proto.Marshal(req.HardwareReport)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "could not serialize harcware report: %v", err)
		}
	}

	// Upsert heartbeat time and hardware report.
	err = session.Transact(ctx, func(q *model.Queries) error {
		// Upsert hardware report if submitted.
		if hwraw != nil {
			err = q.MachineSetHardwareReport(ctx, model.MachineSetHardwareReportParams{
				MachineID:         machineId,
				HardwareReportRaw: hwraw,
			})
			if err != nil {
				return fmt.Errorf("hardware report upsert: %w", err)
			}
		}
		return q.MachineSetAgentHeartbeat(ctx, model.MachineSetAgentHeartbeatParams{
			MachineID:        machineId,
			AgentHeartbeatAt: time.Now(),
		})
	})
	if err != nil {
		klog.Errorf("Could not submit heartbeat: %v", err)
		return nil, status.Error(codes.Unavailable, "could not submit heartbeat")
	}

	return &apb.AgentHeartbeatResponse{}, nil
}
