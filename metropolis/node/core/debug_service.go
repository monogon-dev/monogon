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
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	ctr "github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/roleserve"
	"source.monogon.dev/metropolis/pkg/logtree"
	apb "source.monogon.dev/metropolis/proto/api"
)

const (
	logFilterMax = 1000
)

// debugService implements the Metropolis node debug API.
type debugService struct {
	roleserve       *roleserve.Service
	logtree         *logtree.LogTree
	ephemeralVolume *localstorage.EphemeralContainerdDirectory

	// traceLock provides exclusive access to the Linux tracing infrastructure
	// (ftrace)
	// This is a channel because Go's mutexes can't be cancelled or be acquired
	// in a non-blocking way.
	traceLock chan struct{}
}

func (s *debugService) GetDebugKubeconfig(ctx context.Context, req *apb.GetDebugKubeconfigRequest) (*apb.GetDebugKubeconfigResponse, error) {
	w := s.roleserve.Watch()
	defer w.Close()
	for {
		v, err := w.Get(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, "could not get roleserve status: %v", err)
		}
		if v.Kubernetes == nil {
			continue
		}
		return v.Kubernetes.GetDebugKubeconfig(ctx, req)
	}
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

// Validate property names as they are used in path construction and we really
// don't want a path traversal vulnerability
var safeTracingPropertyNamesRe = regexp.MustCompile("^[a-z0-9_]+$")

func writeTracingProperty(name string, value string) error {
	if !safeTracingPropertyNamesRe.MatchString(name) {
		return fmt.Errorf("disallowed tracing property name received: \"%v\"", name)
	}
	return ioutil.WriteFile("/sys/kernel/tracing/"+name, []byte(value+"\n"), 0)
}

func (s *debugService) Trace(req *apb.TraceRequest, srv apb.NodeDebugService_TraceServer) error {
	// Don't allow more than one trace as the kernel doesn't support this.
	select {
	case s.traceLock <- struct{}{}:
		defer func() {
			<-s.traceLock
		}()
	default:
		return status.Error(codes.FailedPrecondition, "a trace is already in progress")
	}

	if len(req.FunctionFilter) == 0 {
		req.FunctionFilter = []string{"*"} // For reset purposes
	}
	if len(req.GraphFunctionFilter) == 0 {
		req.GraphFunctionFilter = []string{"*"} // For reset purposes
	}

	defer writeTracingProperty("current_tracer", "nop")
	if err := writeTracingProperty("current_tracer", req.Tracer); err != nil {
		return status.Errorf(codes.InvalidArgument, "requested tracer not available: %v", err)
	}

	if err := writeTracingProperty("set_ftrace_filter", strings.Join(req.FunctionFilter, " ")); err != nil {
		return status.Errorf(codes.InvalidArgument, "setting ftrace filter failed: %v", err)
	}
	if err := writeTracingProperty("set_graph_function", strings.Join(req.GraphFunctionFilter, " ")); err != nil {
		return status.Errorf(codes.InvalidArgument, "setting graph filter failed: %v", err)
	}
	tracePipe, err := os.Open("/sys/kernel/tracing/trace_pipe")
	if err != nil {
		return status.Errorf(codes.Unavailable, "cannot open trace output pipe: %v", err)
	}
	defer tracePipe.Close()

	defer writeTracingProperty("tracing_on", "0")
	if err := writeTracingProperty("tracing_on", "1"); err != nil {
		return status.Errorf(codes.InvalidArgument, "requested tracer not available: %v", err)
	}

	go func() {
		<-srv.Context().Done()
		tracePipe.Close()
	}()

	eventScanner := bufio.NewScanner(tracePipe)
	for eventScanner.Scan() {
		if err := eventScanner.Err(); err != nil {
			return status.Errorf(codes.Unavailable, "event pipe read error: %v", err)
		}
		err := srv.Send(&apb.TraceEvent{
			RawLine: eventScanner.Text(),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// imageReader is an adapter converting a gRPC stream into an io.Reader
type imageReader struct {
	srv        apb.NodeDebugService_LoadImageServer
	restOfPart []byte
}

func (i *imageReader) Read(p []byte) (n int, err error) {
	n1 := copy(p, i.restOfPart)
	if len(p) > len(i.restOfPart) {
		part, err := i.srv.Recv()
		if err != nil {
			return n1, err
		}
		n2 := copy(p[n1:], part.DataPart)
		i.restOfPart = part.DataPart[n2:]
		return n1 + n2, nil
	} else {
		i.restOfPart = i.restOfPart[n1:]
		return n1, nil
	}
}

// LoadImage loads an OCI image into the image cache of this node
func (s *debugService) LoadImage(srv apb.NodeDebugService_LoadImageServer) error {
	client, err := ctr.New(s.ephemeralVolume.ClientSocket.FullPath())
	if err != nil {
		return status.Errorf(codes.Unavailable, "failed to connect to containerd: %v", err)
	}
	ctxWithNS := namespaces.WithNamespace(srv.Context(), "k8s.io")
	reader := &imageReader{srv: srv}
	_, err = client.Import(ctxWithNS, reader)
	if err != nil {
		return status.Errorf(codes.Unknown, "failed to import image: %v", err)
	}
	return srv.SendAndClose(&apb.LoadImageResponse{})
}
