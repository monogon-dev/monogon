// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package node

// Implements a debug gRPC service for testing and introspection
// This is attached to the SmalltownNode because most other services are instantiated there and thus are accessible
// from there. Have a look at //core/cmd/dbg if you need to interact with this from a CLI.

import (
	"context"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	schema "git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/storage"
)

func (s *SmalltownNode) GetDebugKubeconfig(ctx context.Context, req *schema.GetDebugKubeconfigRequest) (*schema.GetDebugKubeconfigResponse, error) {
	return s.Kubernetes.GetDebugKubeconfig(ctx, req)
}

// GetComponentLogs gets various logbuffers from binaries we call. This function just deals with the first path component,
// delegating the rest to the service-specific handlers.
func (s *SmalltownNode) GetComponentLogs(ctx context.Context, req *schema.GetComponentLogsRequest) (*schema.GetComponentLogsResponse, error) {
	if len(req.ComponentPath) < 1 {
		return nil, status.Error(codes.InvalidArgument, "component_path needs to contain at least one part")
	}
	linesToRead := int(req.TailLines)
	if linesToRead == 0 {
		linesToRead = math.MaxInt32
	}
	var lines []string
	var err error
	switch req.ComponentPath[0] {
	case "containerd":
		if len(req.ComponentPath) < 2 {
			lines = s.Containerd.Log.ReadLinesTruncated(linesToRead, "...")
		} else if req.ComponentPath[1] == "runsc" {
			lines = s.Containerd.RunscLog.ReadLinesTruncated(linesToRead, "...")
		}
	case "kube":
		if len(req.ComponentPath) < 2 {
			return nil, status.Error(codes.NotFound, "Component not found")
		}
		lines, err = s.Kubernetes.GetComponentLogs(req.ComponentPath[1], linesToRead)
		if err != nil {
			return nil, status.Error(codes.NotFound, "Component not found")
		}
	default:
		return nil, status.Error(codes.NotFound, "component not found")
	}
	return &schema.GetComponentLogsResponse{Line: lines}, nil
}

// GetCondition checks for various conditions exposed by different services. Mostly intended for testing. If you need
// to make sure something is available in an E2E test, consider adding a condition here.
func (s *SmalltownNode) GetCondition(ctx context.Context, req *schema.GetConditionRequest) (*schema.GetConditionResponse, error) {
	var ok bool
	switch req.Name {
	case "IPAssigned":
		ip, err := s.Network.GetIP(ctx, false)
		if err == nil && ip != nil {
			ok = true
		}
	case "DataAvailable":
		_, err := s.Storage.GetPathInPlace(storage.PlaceData, "test")
		if err == nil {
			ok = true
		}
	default:
		return nil, status.Errorf(codes.NotFound, "condition %v not found", req.Name)
	}
	return &schema.GetConditionResponse{
		Ok: ok,
	}, nil
}
