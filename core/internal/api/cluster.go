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
	"encoding/base64"
	"io"

	"github.com/gogo/protobuf/proto"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
	schema "git.monogon.dev/source/nexantic.git/core/generated/api"

	"go.uber.org/zap"
)

func (s *Server) AddNode(ctx context.Context, req *schema.AddNodeRequest) (*schema.AddNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Server) RemoveNode(ctx context.Context, req *schema.RemoveNodeRequest) (*schema.RemoveNodeRequest, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Server) ListNodes(ctx context.Context, req *schema.ListNodesRequest) (*schema.ListNodesResponse, error) {
	store := s.getStore()
	res, err := store.Get(ctx, "nodes/", clientv3.WithPrefix())
	if err != nil {
		return nil, status.Error(codes.Unavailable, "Consensus unavailable")
	}
	var resNodes []*api.Node
	for _, nodeEntry := range res.Kvs {
		var node api.Node
		if err := proto.Unmarshal(nodeEntry.Value, &node); err != nil {
			s.Logger.Error("Encountered invalid node data", zap.Error(err))
			return nil, status.Error(codes.Internal, "Invalid data")
		}
		// Zero out Global Unlock Key, it's never supposed to leave the cluster
		node.GlobalUnlockKey = []byte{}

		resNodes = append(resNodes, &node)
	}

	return &schema.ListNodesResponse{
		Nodes: resNodes,
	}, nil
}

func (s *Server) ListEnrolmentConfigs(ctx context.Context, req *api.ListEnrolmentConfigsRequest) (*api.ListEnrolmentConfigsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Server) NewEnrolmentConfig(ctx context.Context, req *api.NewEnrolmentConfigRequest) (*api.NewEnrolmentConfigResponse, error) {
	store := s.getStore()
	token := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, token); err != nil {
		return nil, status.Error(codes.Unavailable, "failed to get randonmess")
	}
	nodes, err := store.Get(ctx, "nodes/", clientv3.WithPrefix())
	if err != nil {
		return nil, status.Error(codes.Unavailable, "consensus unavailable")
	}
	var masterIPs [][]byte
	for _, nodeKV := range nodes.Kvs {
		var node api.Node
		if err := proto.Unmarshal(nodeKV.Value, &node); err != nil {
			return nil, status.Error(codes.Internal, "invalid node")
		}
		if node.State == api.Node_MASTER {
			masterIPs = append(masterIPs, node.Address)
		}
	}
	masterCert, err := s.GetMasterCert()
	if err != nil {
		return nil, status.Error(codes.Unavailable, "consensus unavailable")
	}

	enrolmentConfig := &api.EnrolmentConfig{
		EnrolmentSecret: token,
		MasterIps:       masterIPs,
		MastersCert:     masterCert,
	}
	enrolmentConfigRaw, err := proto.Marshal(enrolmentConfig)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to encode config")
	}
	if _, err := store.Put(ctx, "enrolments/"+base64.RawURLEncoding.EncodeToString(token), string(enrolmentConfigRaw)); err != nil {
		return nil, status.Error(codes.Unavailable, "consensus unavailable")
	}
	return &schema.NewEnrolmentConfigResponse{
		EnrolmentConfig: enrolmentConfig,
	}, nil
}

func (s *Server) RemoveEnrolmentConfig(ctx context.Context, req *api.RemoveEnrolmentConfigRequest) (*api.RemoveEnrolmentConfigResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
