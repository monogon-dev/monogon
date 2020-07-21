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

package main

import (
	"context"
	"math"

	"git.monogon.dev/source/nexantic.git/core/internal/cluster"
	"git.monogon.dev/source/nexantic.git/core/internal/containerd"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes"
	apb "git.monogon.dev/source/nexantic.git/core/proto/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// debugService implements the Smalltown node debug API.
// TODO(q3k): this should probably be implemented somewhere else way once we have a better
// supervision introspection/status API.
type debugService struct {
	cluster    *cluster.Manager
	kubernetes *kubernetes.Service
	containerd *containerd.Service
}

func (s *debugService) GetDebugKubeconfig(ctx context.Context, req *apb.GetDebugKubeconfigRequest) (*apb.GetDebugKubeconfigResponse, error) {
	return s.kubernetes.GetDebugKubeconfig(ctx, req)
}

// GetComponentLogs gets various logbuffers from binaries we call. This function just deals with the first path component,
// delegating the rest to the service-specific handlers.
func (s *debugService) GetComponentLogs(ctx context.Context, req *apb.GetComponentLogsRequest) (*apb.GetComponentLogsResponse, error) {
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
			lines = s.containerd.Log.ReadLinesTruncated(linesToRead, "...")
		} else if req.ComponentPath[1] == "runsc" {
			lines = s.containerd.RunscLog.ReadLinesTruncated(linesToRead, "...")
		}
	case "kube":
		if len(req.ComponentPath) < 2 {
			return nil, status.Error(codes.NotFound, "Component not found")
		}
		lines, err = s.kubernetes.GetComponentLogs(req.ComponentPath[1], linesToRead)
		if err != nil {
			return nil, status.Error(codes.NotFound, "Component not found")
		}
	default:
		return nil, status.Error(codes.NotFound, "component not found")
	}
	return &apb.GetComponentLogsResponse{Line: lines}, nil
}
