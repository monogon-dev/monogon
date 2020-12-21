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
	"crypto/x509"
	"fmt"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common "git.monogon.dev/source/nexantic.git/metropolis/node"
	"git.monogon.dev/source/nexantic.git/metropolis/node/core/cluster"
	"git.monogon.dev/source/nexantic.git/metropolis/node/core/consensus/ca"
	"git.monogon.dev/source/nexantic.git/metropolis/node/core/logtree"
	"git.monogon.dev/source/nexantic.git/metropolis/node/kubernetes"
	apb "git.monogon.dev/source/nexantic.git/metropolis/proto/api"
)

const (
	logFilterMax = 1000
)

// debugService implements the Metropolis node debug API.
type debugService struct {
	cluster    *cluster.Manager
	kubernetes *kubernetes.Service
	logtree    *logtree.LogTree
}

func (s *debugService) GetGoldenTicket(ctx context.Context, req *apb.GetGoldenTicketRequest) (*apb.GetGoldenTicketResponse, error) {
	ip := net.ParseIP(req.ExternalIp)
	if ip == nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not parse IP %q", req.ExternalIp)
	}
	this := s.cluster.Node()

	certRaw, key, err := s.nodeCertificate()
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to generate node certificate: %v", err)
	}
	cert, err := x509.ParseCertificate(certRaw)
	if err != nil {
		panic(err)
	}
	kv := s.cluster.ConsensusKVRoot()
	ca, err := ca.Load(ctx, kv)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not load CA: %v", err)
	}
	etcdCert, etcdKey, err := ca.Issue(ctx, kv, cert.Subject.CommonName, ip)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not generate etcd peer certificate: %v", err)
	}
	etcdCRL, err := ca.GetCurrentCRL(ctx, kv)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not get etcd CRL: %v", err)
	}

	// Add new etcd member to etcd cluster.
	etcd := s.cluster.ConsensusCluster()
	etcdAddr := fmt.Sprintf("https://%s:%d", ip.String(), common.ConsensusPort)
	_, err = etcd.MemberAddAsLearner(ctx, []string{etcdAddr})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not add as new etcd consensus member: %v", err)
	}

	return &apb.GetGoldenTicketResponse{
		Ticket: &apb.GoldenTicket{
			EtcdCaCert:     ca.CACertRaw,
			EtcdClientCert: etcdCert,
			EtcdClientKey:  etcdKey,
			EtcdCrl:        etcdCRL,
			Peers: []*apb.GoldenTicket_EtcdPeer{
				{Name: this.ID(), Address: this.Address().String()},
			},
			This: &apb.GoldenTicket_EtcdPeer{Name: cert.Subject.CommonName, Address: ip.String()},

			NodeId:   cert.Subject.CommonName,
			NodeCert: certRaw,
			NodeKey:  key,
		},
	}, nil
}

func (s *debugService) GetDebugKubeconfig(ctx context.Context, req *apb.GetDebugKubeconfigRequest) (*apb.GetDebugKubeconfigResponse, error) {
	return s.kubernetes.GetDebugKubeconfig(ctx, req)
}

func (s *debugService) GetLogs(req *apb.GetLogsRequest, srv apb.NodeDebugService_GetLogsServer) error {
	if len(req.Filters) > logFilterMax {
		return status.Errorf(codes.InvalidArgument, "requested %d filters, maximum permitted is %d", len(req.Filters), logFilterMax)
	}
	dn := logtree.DN(req.Dn)
	_, err := dn.Path()
	switch err {
	case nil:
	case logtree.ErrInvalidDN:
		return status.Errorf(codes.InvalidArgument, "invalid DN")
	default:
		return status.Errorf(codes.Unavailable, "could not parse DN: %v", err)
	}

	var options []logtree.LogReadOption

	// Turn backlog mode into logtree option(s).
	switch req.BacklogMode {
	case apb.GetLogsRequest_BACKLOG_DISABLE:
	case apb.GetLogsRequest_BACKLOG_ALL:
		options = append(options, logtree.WithBacklog(logtree.BacklogAllAvailable))
	case apb.GetLogsRequest_BACKLOG_COUNT:
		count := int(req.BacklogCount)
		if count <= 0 {
			return status.Errorf(codes.InvalidArgument, "backlog_count must be > 0 if backlog_mode is BACKLOG_COUNT")
		}
		options = append(options, logtree.WithBacklog(count))
	default:
		return status.Errorf(codes.InvalidArgument, "unknown backlog_mode %d", req.BacklogMode)
	}

	// Turn stream mode into logtree option(s).
	streamEnable := false
	switch req.StreamMode {
	case apb.GetLogsRequest_STREAM_DISABLE:
	case apb.GetLogsRequest_STREAM_UNBUFFERED:
		streamEnable = true
		options = append(options, logtree.WithStream())
	}

	// Parse proto filters into logtree options.
	for i, filter := range req.Filters {
		switch inner := filter.Filter.(type) {
		case *apb.LogFilter_WithChildren_:
			options = append(options, logtree.WithChildren())
		case *apb.LogFilter_OnlyRaw_:
			options = append(options, logtree.OnlyRaw())
		case *apb.LogFilter_OnlyLeveled_:
			options = append(options, logtree.OnlyLeveled())
		case *apb.LogFilter_LeveledWithMinimumSeverity_:
			severity, err := logtree.SeverityFromProto(inner.LeveledWithMinimumSeverity.Minimum)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "filter %d has invalid severity: %v", i, err)
			}
			options = append(options, logtree.LeveledWithMinimumSeverity(severity))
		}
	}

	reader, err := s.logtree.Read(logtree.DN(req.Dn), options...)
	switch err {
	case nil:
	case logtree.ErrRawAndLeveled:
		return status.Errorf(codes.InvalidArgument, "requested only raw and only leveled logs simultaneously")
	default:
		return status.Errorf(codes.Unavailable, "could not retrieve logs: %v", err)
	}
	defer reader.Close()

	// Default protobuf message size limit is 64MB. We want to limit ourselves
	// to 10MB.
	// Currently each raw log line can be at most 1024 unicode codepoints (or
	// 4096 bytes). To cover extra metadata and proto overhead, let's round
	// this up to 4500 bytes. This in turn means we can store a maximum of
	// (10e6/4500) == 2222 entries.
	// Currently each leveled log line can also be at most 1024 unicode
	// codepoints (or 4096 bytes). To cover extra metadata and proto overhead
	// let's round this up to 2000 bytes. This in turn means we can store a
	// maximum of (10e6/5000) == 2000 entries.
	// The lowever of these numbers, ie the worst case scenario, is 2000
	// maximum entries.
	maxChunkSize := 2000

	// Serve all backlog entries in chunks.
	chunk := make([]*apb.LogEntry, 0, maxChunkSize)
	for _, entry := range reader.Backlog {
		p := entry.Proto()
		if p == nil {
			// TODO(q3k): log this once we have logtree/gRPC compatibility.
			continue
		}
		chunk = append(chunk, p)

		if len(chunk) >= maxChunkSize {
			err := srv.Send(&apb.GetLogsResponse{
				BacklogEntries: chunk,
			})
			if err != nil {
				return err
			}
			chunk = make([]*apb.LogEntry, 0, maxChunkSize)
		}
	}

	// Send last chunk of backlog, if present..
	if len(chunk) > 0 {
		err := srv.Send(&apb.GetLogsResponse{
			BacklogEntries: chunk,
		})
		if err != nil {
			return err
		}
		chunk = make([]*apb.LogEntry, 0, maxChunkSize)
	}

	// Start serving streaming data, if streaming has been requested.
	if !streamEnable {
		return nil
	}

	for {
		entry, ok := <-reader.Stream
		if !ok {
			// Streaming has been ended by logtree - tell the client and return.
			return status.Error(codes.Unavailable, "log streaming aborted by system")
		}
		p := entry.Proto()
		if p == nil {
			// TODO(q3k): log this once we have logtree/gRPC compatibility.
			continue
		}
		err := srv.Send(&apb.GetLogsResponse{
			StreamEntries: []*apb.LogEntry{p},
		})
		if err != nil {
			return err
		}
	}
}
