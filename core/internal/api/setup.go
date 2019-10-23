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
	"errors"
	"fmt"
	schema "git.monogon.dev/source/nexantic.git/core/generated/api"
)

const (
	MinNameLength = 3
)

var (
	ErrInvalidProvisioningToken = errors.New("invalid provisioning token")
	ErrInvalidNameLength        = fmt.Errorf("name must be at least %d characters long", MinNameLength)
)

func (s *Server) Setup(c context.Context, r *schema.SetupRequest) (*schema.SetupResponse, error) {

	switch r.Request.(type) {
	case *schema.SetupRequest_JoinCluster:
		token, err := s.enterJoinCluster(r.GetJoinCluster())
		if err != nil {
			return nil, err
		}

		return &schema.SetupResponse{
			Response: &schema.SetupResponse_JoinCluster{
				JoinCluster: &schema.JoinClusterResponse{
					ProvisioningToken: token,
				},
			},
		}, nil

	case *schema.SetupRequest_NewCluster:
		return &schema.SetupResponse{
			Response: &schema.SetupResponse_NewCluster{
				NewCluster: &schema.NewClusterResponse{},
			},
		}, s.setupNewCluster(r.GetNewCluster())
	}

	return &schema.SetupResponse{}, nil
}

func (s *Server) enterJoinCluster(r *schema.JoinClusterRequest) (string, error) {
	err := s.setupService.EnterJoinClusterMode()
	if err != nil {
		return "", err
	}

	return s.setupService.GetJoinClusterToken(), nil
}

func (s *Server) setupNewCluster(r *schema.NewClusterRequest) error {
	if len(r.NodeName) < MinNameLength {
		return ErrInvalidNameLength
	}
	return s.setupService.SetupNewCluster(r.NodeName, r.ExternalHost)
}

func (s *Server) ProvisionCluster(ctx context.Context, req *schema.ProvisionClusterRequest) (*schema.ProvisionClusterResponse, error) {
	if len(req.NodeName) < MinNameLength {
		return nil, ErrInvalidNameLength
	}

	// Verify provisioning token
	if s.setupService.GetJoinClusterToken() != req.ProvisioningToken {
		return nil, ErrInvalidProvisioningToken
	}

	// Join cluster
	err := s.setupService.JoinCluster(req.NodeName, req.InitialCluster, req.ExternalHost)
	if err != nil {
		return nil, err
	}

	return &schema.ProvisionClusterResponse{}, nil
}

func (s *Server) Attest(c context.Context, r *schema.AttestRequest) (*schema.AttestResponse, error) {
	// TODO implement
	return &schema.AttestResponse{
		Response: r.Challenge,
	}, nil
}
