// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	apb "source.monogon.dev/cloud/bmaas/server/api"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
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

	session, err := a.s.session(ctx)
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
			klog.Errorf("No agent for %s/%s", machineId.String(), hex.EncodeToString(pk))
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
			return nil, status.Errorf(codes.InvalidArgument, "could not serialize hardware report: %v", err)
		}
	}

	var installRaw []byte
	if req.InstallationReport != nil {
		installRaw, err = proto.Marshal(req.InstallationReport)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "could not serialize installation report: %v", err)
		}
	}

	// Upsert heartbeat time and hardware report.
	err = session.Transact(ctx, func(q *model.Queries) error {
		// Upsert hardware report if submitted.
		if len(hwraw) != 0 {
			err = q.MachineSetHardwareReport(ctx, model.MachineSetHardwareReportParams{
				MachineID:         machineId,
				HardwareReportRaw: hwraw,
			})
			if err != nil {
				return fmt.Errorf("hardware report upsert: %w", err)
			}
		}
		// Upsert os installation report if submitted.
		if len(installRaw) != 0 {
			var result model.MachineOsInstallationResult
			switch req.InstallationReport.Result.(type) {
			case *apb.OSInstallationReport_Success_:
				result = model.MachineOsInstallationResultSuccess
			case *apb.OSInstallationReport_Error_:
				result = model.MachineOsInstallationResultError
			default:
				return fmt.Errorf("unknown installation report result: %T", req.InstallationReport.Result)
			}
			err = q.MachineSetOSInstallationReport(ctx, model.MachineSetOSInstallationReportParams{
				MachineID:               machineId,
				Generation:              req.InstallationReport.Generation,
				OsInstallationResult:    result,
				OsInstallationReportRaw: installRaw,
			})
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
	klog.Infof("Heartbeat from %s/%s", machineId.String(), hex.EncodeToString(pk))

	// Get installation request for machine if present.
	var installRequest *apb.OSInstallationRequest
	err = session.Transact(ctx, func(q *model.Queries) error {
		reqs, err := q.GetExactMachineForOSInstallation(ctx, model.GetExactMachineForOSInstallationParams{
			MachineID: machineId,
			Limit:     1,
		})
		if err != nil {
			return fmt.Errorf("GetExactMachineForOSInstallation: %w", err)
		}
		if len(reqs) > 0 {
			raw := reqs[0].OsInstallationRequestRaw
			var preq apb.OSInstallationRequest
			if err := proto.Unmarshal(raw, &preq); err != nil {
				return fmt.Errorf("could not decode stored OS installation request: %w", err)
			}
			installRequest = &preq
		}
		return nil
	})
	if err != nil {
		// Do not fail entire request. Instead, just log an error.
		// TODO(q3k): alert on this
		klog.Errorf("Failure during OS installation request retrieval: %v", err)
	}

	return &apb.AgentHeartbeatResponse{
		InstallationRequest: installRequest,
	}, nil
}
