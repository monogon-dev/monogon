// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package mgmt

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/testing/protocmp"

	"source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	"source.monogon.dev/osbase/logtree"
	lpb "source.monogon.dev/osbase/logtree/proto"
)

func dut(t *testing.T) (*Service, *grpc.ClientConn) {
	lt := logtree.New()
	s := &Service{
		LogTree: lt,
		LogService: LogService{
			LogTree: lt,
		},
	}

	srv := grpc.NewServer()
	api.RegisterNodeManagementServer(srv, s)
	externalLis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := srv.Serve(externalLis); err != nil {
			t.Errorf("GRPC serve failed: %v", err)
			return
		}
	}()
	withLocalDialer := grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return externalLis.Dial()
	})
	cl, err := grpc.Dial("local", withLocalDialer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Dialing GRPC failed: %v", err)
	}

	return s, cl
}

func cleanLogEntry(e *lpb.LogEntry) {
	// Filter out bits that change too much to test them.
	switch k := e.Kind.(type) {
	case *lpb.LogEntry_Leveled_:
		k.Leveled.Location = ""
		k.Leveled.Timestamp = nil
	}
}

func mkRawEntry(dn string, line string) *lpb.LogEntry {
	return &lpb.LogEntry{
		Dn: dn, Kind: &lpb.LogEntry_Raw_{
			Raw: &lpb.LogEntry_Raw{
				Data:           line,
				OriginalLength: int64(len(line)),
			},
		},
	}
}

func mkLeveledEntry(dn string, severity string, lines string) *lpb.LogEntry {
	var sev lpb.LeveledLogSeverity
	switch severity {
	case "i":
		sev = lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_INFO
	case "w":
		sev = lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_WARNING
	case "e":
		sev = lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_ERROR
	}
	return &lpb.LogEntry{
		Dn: dn, Kind: &lpb.LogEntry_Leveled_{
			Leveled: &lpb.LogEntry_Leveled{
				Lines:    strings.Split(lines, "\n"),
				Severity: sev,
			},
		},
	}
}

func drainLogs(t *testing.T, srv api.NodeManagement_LogsClient) (res []*lpb.LogEntry) {
	t.Helper()
	for {
		ev, err := srv.Recv()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			t.Errorf("Recv: %v", err)
			return
		}
		res = append(res, ev.BacklogEntries...)
	}
}

// TestLogService_Logs_Backlog exercises the basic log API by requesting
// backlogged leveled log entries.
func TestLogService_Logs_Backlog(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	s, cl := dut(t)

	mgmt := api.NewNodeManagementClient(cl)

	s.LogTree.MustLeveledFor("init").Infof("Hello")
	s.LogTree.MustLeveledFor("main").Infof("Starting roleserver...")
	s.LogTree.MustLeveledFor("main.roleserver").Infof("Waiting for node roles...")
	s.LogTree.MustLeveledFor("main.roleserver.kubernetes").Infof("Starting kubernetes...")
	s.LogTree.MustLeveledFor("main.roleserver.controlplane").Infof("Starting control plane...")
	s.LogTree.MustLeveledFor("main.roleserver.kubernetes").Infof("Kubernetes version: 1.21.37")
	s.LogTree.MustLeveledFor("main.roleserver.controlplane").Infof("Starting etcd...")
	s.LogTree.MustLeveledFor("main.weirdo").Infof("Here comes some invalid utf-8: a\xc5z")

	mkReq := func(dn string, backlog int64) *api.GetLogsRequest {
		var backlogMode api.GetLogsRequest_BacklogMode
		var backlogCount int64
		switch {
		case backlog < 0:
			backlogMode = api.GetLogsRequest_BACKLOG_MODE_ALL
		case backlog > 0:
			backlogMode = api.GetLogsRequest_BACKLOG_MODE_COUNT
			backlogCount = backlog
		case backlog == 0:
			backlogMode = api.GetLogsRequest_BACKLOG_MODE_DISABLE
		}
		return &api.GetLogsRequest{
			Dn:           dn,
			BacklogMode:  backlogMode,
			BacklogCount: backlogCount,
			StreamMode:   api.GetLogsRequest_STREAM_MODE_DISABLE,
		}
	}
	mkRecursive := func(in *api.GetLogsRequest) *api.GetLogsRequest {
		in.Filters = append(in.Filters, &cpb.LogFilter{
			Filter: &cpb.LogFilter_WithChildren_{
				WithChildren: &cpb.LogFilter_WithChildren{},
			},
		})
		return in
	}
	for i, te := range []struct {
		req  *api.GetLogsRequest
		want []*lpb.LogEntry
	}{
		{
			// Test all backlog.
			req: mkReq("main.roleserver.kubernetes", -1),
			want: []*lpb.LogEntry{
				mkLeveledEntry("main.roleserver.kubernetes", "i", "Starting kubernetes..."),
				mkLeveledEntry("main.roleserver.kubernetes", "i", "Kubernetes version: 1.21.37"),
			},
		},
		{
			// Test exact backlog.
			req: mkReq("main.roleserver.kubernetes", 1),
			want: []*lpb.LogEntry{
				mkLeveledEntry("main.roleserver.kubernetes", "i", "Kubernetes version: 1.21.37"),
			},
		},
		{
			// Test no backlog.
			req:  mkReq("main.roleserver.kubernetes", 0),
			want: nil,
		},
		{
			// Test recursion with backlog.
			req: mkRecursive(mkReq("main.roleserver", 2)),
			want: []*lpb.LogEntry{
				mkLeveledEntry("main.roleserver.kubernetes", "i", "Kubernetes version: 1.21.37"),
				mkLeveledEntry("main.roleserver.controlplane", "i", "Starting etcd..."),
			},
		},
		{
			// Test invalid utf-8 in log data
			req: mkReq("main.weirdo", 1),
			want: []*lpb.LogEntry{
				mkLeveledEntry("main.weirdo", "i", "Here comes some invalid utf-8: a<INVALID>z"),
			},
		},
	} {
		srv, err := mgmt.Logs(ctx, te.req)
		if err != nil {
			t.Errorf("Case %d: Logs failed: %v", i, err)
			continue
		}
		logs := drainLogs(t, srv)
		for _, e := range logs {
			cleanLogEntry(e)
		}
		diff := cmp.Diff(te.want, logs, protocmp.Transform())
		if diff != "" {
			t.Errorf("Case %d: diff: \n%s", i, diff)
		}
	}
}

// TestLogService_Logs_Strea, exercises the basic log API by requesting
// streaming leveled log entries.
func TestLogService_Logs_Stream(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	s, cl := dut(t)

	// Start streaming all logs.
	mgmt := api.NewNodeManagementClient(cl)
	srv, err := mgmt.Logs(ctx, &api.GetLogsRequest{
		Dn:          "",
		BacklogMode: api.GetLogsRequest_BACKLOG_MODE_ALL,
		StreamMode:  api.GetLogsRequest_STREAM_MODE_UNBUFFERED,
		Filters: []*cpb.LogFilter{
			{
				Filter: &cpb.LogFilter_WithChildren_{
					WithChildren: &cpb.LogFilter_WithChildren{},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("Logs failed: %v", err)
	}

	// Pipe returned logs into a channel for analysis.
	logC := make(chan *lpb.LogEntry)
	go func() {
		for {
			ev, err := srv.Recv()
			if err != nil {
				return
			}
			for _, e := range ev.BacklogEntries {
				logC <- e
			}
			for _, e := range ev.StreamEntries {
				logC <- e
			}
		}
	}()

	// Submit log entry, expect it on the channel.
	s.LogTree.MustLeveledFor("test").Infof("Hello, world")
	select {
	case e := <-logC:
		cleanLogEntry(e)
		if diff := cmp.Diff(mkLeveledEntry("test", "i", "Hello, world"), e, protocmp.Transform()); diff != "" {
			t.Errorf("Diff:\n%s", diff)
		}
	case <-time.After(time.Second * 2):
		t.Errorf("Timeout")
	}

	// That could've made it through the backlog. Do it again to make sure it came
	// through streaming.
	s.LogTree.MustLeveledFor("test").Infof("Hello again, world")
	select {
	case e := <-logC:
		cleanLogEntry(e)
		if diff := cmp.Diff(mkLeveledEntry("test", "i", "Hello again, world"), e, protocmp.Transform()); diff != "" {
			t.Errorf("Diff:\n%s", diff)
		}
	case <-time.After(time.Second * 2):
		t.Errorf("Timeout")
	}
}

// TestLogService_Logs_Filters exercises the rest of the filter functionality.
func TestLogService_Logs_Filters(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	s, cl := dut(t)

	mgmt := api.NewNodeManagementClient(cl)
	s.LogTree.MustLeveledFor("main").Infof("Hello")
	s.LogTree.MustLeveledFor("main").Infof("Starting...")
	s.LogTree.MustLeveledFor("main").Warningf("Something failed!")
	fmt.Fprintln(s.LogTree.MustRawFor("main"), "medium rare")
	s.LogTree.MustLeveledFor("main").Errorf("Something failed very hard!")

	for i, te := range []struct {
		req  *api.GetLogsRequest
		want []*lpb.LogEntry
	}{
		// Case 0: request given level
		{
			req: &api.GetLogsRequest{
				Dn:          "main",
				BacklogMode: api.GetLogsRequest_BACKLOG_MODE_ALL,
				StreamMode:  api.GetLogsRequest_STREAM_MODE_DISABLE,
				Filters: []*cpb.LogFilter{
					{
						Filter: &cpb.LogFilter_LeveledWithMinimumSeverity_{
							LeveledWithMinimumSeverity: &cpb.LogFilter_LeveledWithMinimumSeverity{
								Minimum: lpb.LeveledLogSeverity_LEVELED_LOG_SEVERITY_WARNING,
							},
						},
					},
				},
			},
			want: []*lpb.LogEntry{
				mkLeveledEntry("main", "w", "Something failed!"),
				mkLeveledEntry("main", "e", "Something failed very hard!"),
			},
		},
		// Case 1: request raw only
		{
			req: &api.GetLogsRequest{
				Dn:          "main",
				BacklogMode: api.GetLogsRequest_BACKLOG_MODE_ALL,
				StreamMode:  api.GetLogsRequest_STREAM_MODE_DISABLE,
				Filters: []*cpb.LogFilter{
					{
						Filter: &cpb.LogFilter_OnlyRaw_{
							OnlyRaw: &cpb.LogFilter_OnlyRaw{},
						},
					},
				},
			},
			want: []*lpb.LogEntry{
				mkRawEntry("main", "medium rare"),
			},
		},
		// Case 2: request leveled only
		{
			req: &api.GetLogsRequest{
				Dn:          "main",
				BacklogMode: api.GetLogsRequest_BACKLOG_MODE_ALL,
				StreamMode:  api.GetLogsRequest_STREAM_MODE_DISABLE,
				Filters: []*cpb.LogFilter{
					{
						Filter: &cpb.LogFilter_OnlyLeveled_{
							OnlyLeveled: &cpb.LogFilter_OnlyLeveled{},
						},
					},
				},
			},
			want: []*lpb.LogEntry{
				mkLeveledEntry("main", "i", "Hello"),
				mkLeveledEntry("main", "i", "Starting..."),
				mkLeveledEntry("main", "w", "Something failed!"),
				mkLeveledEntry("main", "e", "Something failed very hard!"),
			},
		},
	} {
		srv, err := mgmt.Logs(ctx, te.req)
		if err != nil {
			t.Errorf("Case %d: Logs failed: %v", i, err)
			continue
		}
		logs := drainLogs(t, srv)
		for _, e := range logs {
			cleanLogEntry(e)
		}
		diff := cmp.Diff(te.want, logs, protocmp.Transform())
		if diff != "" {
			t.Errorf("Case %d: diff: \n%s", i, diff)
		}
	}

}
