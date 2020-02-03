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

package api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	schema "git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.uber.org/zap"
)

var (
	ErrAttestationFailed = status.Error(codes.PermissionDenied, "attestation failed")
)

func (s *Server) AddNode(ctx context.Context, req *schema.AddNodeRequest) (*schema.AddNodeResponse, error) {
	// Setup API client
	c, err := grpc.NewSmalltownAPIClient(fmt.Sprintf("%s:%d", req.Addr, s.config.Port))
	if err != nil {
		return nil, err
	}

	// Check attestation
	nonce := make([]byte, 20)
	_, err = rand.Read(nonce)
	if err != nil {
		s.Logger.Error("Nonce generation failed", zap.Error(err))
		return nil, status.Error(codes.Unavailable, "nonce generation failed")
	}
	hexNonce := hex.EncodeToString(nonce)

	aRes, err := c.Setup.Attest(ctx, &schema.AttestRequest{
		Challenge: hexNonce,
	})
	if err != nil {
		s := status.Convert(err)
		return nil, status.Errorf(s.Code(), "attestation failed: %v", s.Message())
	}

	//TODO(hendrik): Verify response
	if aRes.Response != hexNonce {
		return nil, ErrAttestationFailed
	}

	consensusCerts, err := s.consensusService.IssueCertificate(req.Addr)
	if err != nil {
		// Errors from IssueCertificate are always treated as internal
		s.Logger.Error("Node certificate issuance failed", zap.String("addr", req.Addr), zap.Error(err))
		return nil, status.Error(codes.Internal, "could not issue node certificate")
	}

	// TODO(leo): fetch remote hostname rather than using the addr
	name := req.Addr

	// Add new node to local etcd cluster.
	memberID, err := s.consensusService.AddMember(ctx, name, fmt.Sprintf("https://%s:%d", req.Addr, s.config.Port))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to add node to etcd cluster: %v", err)
	}

	s.Logger.Info("Added new node to consensus cluster; sending cluster join request to node",
		zap.String("addr", req.Addr), zap.Uint16("port", s.config.Port))

	// Send JoinCluster request to new node to make it join.
	_, err = c.Setup.JoinCluster(ctx, &schema.JoinClusterRequest{
		InitialCluster:    s.consensusService.GetInitialClusterString(),
		ProvisioningToken: req.ProvisioningToken,
		Certs:             consensusCerts,
	})
	if err != nil {
		errRevoke := s.consensusService.RevokeCertificate(req.Addr)
		if errRevoke != nil {
			s.Logger.Error("Failed to revoke a certificate after rollback - potential security risk", zap.Error(errRevoke))
		}
		// Revert etcd add member - might fail if consensus cannot be established.
		errRemove := s.consensusService.RemoveMember(ctx, memberID)
		if errRemove != nil || errRevoke != nil {
			return nil, fmt.Errorf("rollback failed after failed provisioning; err=%v; err_rb=%v; err_revoke=%v", err, errRemove, errRevoke)
		}
		return nil, status.Errorf(codes.Unavailable, "failed to join etcd cluster with node: %v", err)
	}
	s.Logger.Info("Fully provisioned new node",
		zap.String("host", req.Addr),
		zap.Uint16("apiPort", s.config.Port),
		zap.Uint64("member_id", memberID))

	return &schema.AddNodeResponse{}, nil
}

func (s *Server) RemoveNode(context.Context, *schema.RemoveNodeRequest) (*schema.RemoveNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Server) ListNodes(context.Context, *schema.ListNodesRequest) (*schema.ListNodesResponse, error) {
	nodes := s.consensusService.GetNodes()
	resNodes := make([]*schema.Node, len(nodes))

	for i, node := range nodes {
		resNodes[i] = &schema.Node{
			Id:      node.ID,
			Name:    node.Name,
			Address: node.Address,
			Synced:  node.Synced,
		}
	}

	return &schema.ListNodesResponse{
		Nodes: resNodes,
	}, nil
}
