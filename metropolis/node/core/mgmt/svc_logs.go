package mgmt

import (
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	"source.monogon.dev/osbase/logtree"
	lpb "source.monogon.dev/osbase/logtree/proto"
)

const (
	logFilterMax = 10
)

// LogService implements NodeManagement.Logs. This is split away from the rest of
// the Service to allow the debug service to reuse this implementation.
type LogService struct {
	LogTree *logtree.LogTree
}

// sanitizedEntries returns a deep copy of the given log entries, but replaces
// all invalid UTF-8 characters with "<INVALID>".
func sanitizedEntries(entries []*lpb.LogEntry) []*lpb.LogEntry {
	res := make([]*lpb.LogEntry, len(entries))
	for i, entry := range entries {
		res[i] = &lpb.LogEntry{
			Dn:   entry.Dn,
			Kind: nil,
		}
		switch k := entry.Kind.(type) {
		case *lpb.LogEntry_Leveled_:
			leveled := &lpb.LogEntry_Leveled_{
				Leveled: &lpb.LogEntry_Leveled{
					Lines:     make([]string, len(k.Leveled.Lines)),
					Timestamp: k.Leveled.Timestamp,
					Severity:  k.Leveled.Severity,
					Location:  k.Leveled.Location,
				},
			}
			for j, line := range k.Leveled.Lines {
				leveled.Leveled.Lines[j] = strings.ToValidUTF8(line, "<INVALID>")
			}
			res[i].Kind = leveled

		case *lpb.LogEntry_Raw_:
			res[i].Kind = &lpb.LogEntry_Raw_{
				Raw: &lpb.LogEntry_Raw{
					Data:           strings.ToValidUTF8(k.Raw.Data, "<INVALID>"),
					OriginalLength: k.Raw.OriginalLength,
				},
			}
		}
	}
	return res
}

func (s *LogService) Logs(req *api.GetLogsRequest, srv api.NodeManagement_LogsServer) error {
	if len(req.Filters) > logFilterMax {
		return status.Errorf(codes.InvalidArgument, "requested %d filters, maximum permitted is %d", len(req.Filters), logFilterMax)
	}
	dn := logtree.DN(req.Dn)
	_, err := dn.Path()
	switch {
	case err == nil:
	case errors.Is(err, logtree.ErrInvalidDN):
		return status.Errorf(codes.InvalidArgument, "invalid DN")
	default:
		return status.Errorf(codes.Unavailable, "could not parse DN: %v", err)
	}

	var options []logtree.LogReadOption

	// Turn backlog mode into logtree option(s).
	switch req.BacklogMode {
	case api.GetLogsRequest_BACKLOG_MODE_DISABLE:
	case api.GetLogsRequest_BACKLOG_MODE_ALL:
		options = append(options, logtree.WithBacklog(logtree.BacklogAllAvailable))
	case api.GetLogsRequest_BACKLOG_MODE_COUNT:
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
	case api.GetLogsRequest_STREAM_MODE_DISABLE:
	case api.GetLogsRequest_STREAM_MODE_UNBUFFERED:
		streamEnable = true
		options = append(options, logtree.WithStream())
	}

	// Parse proto filters into logtree options.
	for i, filter := range req.Filters {
		switch inner := filter.Filter.(type) {
		case *cpb.LogFilter_WithChildren_:
			options = append(options, logtree.WithChildren())
		case *cpb.LogFilter_OnlyRaw_:
			options = append(options, logtree.OnlyRaw())
		case *cpb.LogFilter_OnlyLeveled_:
			options = append(options, logtree.OnlyLeveled())
		case *cpb.LogFilter_LeveledWithMinimumSeverity_:
			severity, err := logtree.SeverityFromProto(inner.LeveledWithMinimumSeverity.Minimum)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "filter %d has invalid severity: %v", i, err)
			}
			options = append(options, logtree.LeveledWithMinimumSeverity(severity))
		}
	}

	reader, err := s.LogTree.Read(logtree.DN(req.Dn), options...)
	switch {
	case err == nil:
	case errors.Is(err, logtree.ErrRawAndLeveled):
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
	chunk := make([]*lpb.LogEntry, 0, maxChunkSize)
	for _, entry := range reader.Backlog {
		p := entry.Proto()
		if p == nil {
			// TODO(q3k): log this once we have logtree/gRPC compatibility.
			continue
		}
		chunk = append(chunk, p)

		if len(chunk) >= maxChunkSize {
			err := srv.Send(&api.GetLogsResponse{
				BacklogEntries: sanitizedEntries(chunk),
			})
			if err != nil {
				return err
			}
			chunk = make([]*lpb.LogEntry, 0, maxChunkSize)
		}
	}

	// Send last chunk of backlog, if present..
	if len(chunk) > 0 {
		err := srv.Send(&api.GetLogsResponse{
			BacklogEntries: sanitizedEntries(chunk),
		})
		if err != nil {
			return err
		}
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
		err := srv.Send(&api.GetLogsResponse{
			StreamEntries: []*lpb.LogEntry{p},
		})
		if err != nil {
			return err
		}
	}
}
