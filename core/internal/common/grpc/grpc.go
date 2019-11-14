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

package grpc

import (
	"git.monogon.dev/source/nexantic.git/core/generated/api"

	"google.golang.org/grpc"
)

type (
	SmalltownClient struct {
		conn *grpc.ClientConn

		Cluster api.ClusterManagementClient
		Setup   api.SetupServiceClient
	}
)

func NewSmalltownAPIClient(address string) (*SmalltownClient, error) {
	s := &SmalltownClient{}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	s.conn = conn

	// Setup all client connections
	s.Cluster = api.NewClusterManagementClient(conn)
	s.Setup = api.NewSetupServiceClient(conn)

	return s, nil
}

func (s *SmalltownClient) Close() error {
	return s.conn.Close()
}
