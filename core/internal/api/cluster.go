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
	"git.monogon.dev/source/nexantic.git/core/internal/common"

	"errors"

	"go.uber.org/zap"
)

var (
	ErrAttestationFailed = errors.New("attestation_failed")
)

func (s *Server) AddNode(ctx context.Context, req *schema.AddNodeRequest) (*schema.AddNodeResponse, error) {
	// Setup API client
	c, err := common.NewSmalltownAPIClient(fmt.Sprintf("%s:%d", req.Host, req.ApiPort))
	if err != nil {
		return nil, err
	}

	// Check attestation
	nonce := make([]byte, 20)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	hexNonce := hex.EncodeToString(nonce)

	aRes, err := c.Setup.Attest(ctx, &schema.AttestRequest{
		Challenge: hexNonce,
	})
	if err != nil {
		return nil, err
	}

	//TODO(hendrik): Verify response
	if aRes.Response != hexNonce {
		return nil, ErrAttestationFailed
	}

	// Provision cluster info locally
	memberID, err := s.consensusService.AddMember(ctx, req.Name, fmt.Sprintf("http://%s:%d", req.Host, req.ConsensusPort))
	if err != nil {
		return nil, err
	}

	s.Logger.Info("Added new node to consensus cluster; provisioning external node now",
		zap.String("host", req.Host), zap.Uint32("port", req.ApiPort),
		zap.Uint32("consensus_port", req.ConsensusPort), zap.String("name", req.Name))

	// Provision cluster info externally
	_, err = c.Setup.ProvisionCluster(ctx, &schema.ProvisionClusterRequest{
		InitialCluster:    s.consensusService.GetInitialClusterString(),
		ProvisioningToken: req.Token,
		ExternalHost:      req.Host,
		NodeName:          req.Name,
		TrustBackend:      req.TrustBackend,
	})
	if err != nil {
		// Revert Consensus add member - might fail if consensus cannot be established
		err2 := s.consensusService.RemoveMember(ctx, memberID)
		if err2 != nil {
			return nil, fmt.Errorf("Rollback failed after failed provisioning; err=%v; err_rb=%v", err, err2)
		}
		return nil, err
	}
	s.Logger.Info("Fully provisioned new node",
		zap.String("host", req.Host), zap.Uint32("port", req.ApiPort),
		zap.Uint32("consensus_port", req.ConsensusPort), zap.String("name", req.Name),
		zap.Uint64("member_id", memberID))

	return &schema.AddNodeResponse{}, nil
}

func (s *Server) RemoveNode(context.Context, *schema.RemoveNodeRequest) (*schema.RemoveNodeRequest, error) {
	panic("implement me")
}

func (s *Server) GetNodes(context.Context, *schema.GetNodesRequest) (*schema.GetNodesResponse, error) {
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

	return &schema.GetNodesResponse{
		Nodes: resNodes,
	}, nil
}
