// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package consensus

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"source.monogon.dev/go/logging"
	"source.monogon.dev/osbase/logbuffer"
	"source.monogon.dev/osbase/logtree"
	"source.monogon.dev/osbase/logtree/unraw"
)

// etcdLogEntry is a JSON-encoded, structured log entry received from a running
// etcd server. The format comes from the logging library used there,
// github.com/uber-go/zap.
type etcdLogEntry struct {
	Level   string                 `json:"level"`
	TS      time.Time              `json:"ts"`
	Caller  string                 `json:"caller"`
	Message string                 `json:"msg"`
	Extras  map[string]interface{} `json:"-"`
}

// parseEtcdLogEntry is a logtree/unraw compatible parser for etcd log lines.
// It is fairly liberal in what it will accept, falling back to writing a
// message that outlines the given log entry could not have been parsed. This
// ensures that no lines are lost, even if malformed.
func parseEtcdLogEntry(l *logbuffer.Line, write unraw.LeveledWriter) {
	if l.Truncated() {
		write(&logtree.ExternalLeveledPayload{
			Message: "Log line truncated: " + l.Data,
		})
		return
	}

	var e etcdLogEntry
	// Parse constant fields
	if err := json.Unmarshal([]byte(l.Data), &e); err != nil {
		write(&logtree.ExternalLeveledPayload{
			Message: "Log line unparseable: " + l.Data,
		})
		return
	}
	// Parse extra fields.
	if err := json.Unmarshal([]byte(l.Data), &e.Extras); err != nil {
		// Not exactly sure how this could ever happen - the previous parse
		// went fine, so why wouldn't this one? But to be on the safe side,
		// just don't attempt to parse this line any further.
		write(&logtree.ExternalLeveledPayload{
			Message: "Log line unparseable: " + l.Data,
		})
		return
	}
	delete(e.Extras, "level")
	delete(e.Extras, "ts")
	delete(e.Extras, "caller")
	delete(e.Extras, "msg")

	out := logtree.ExternalLeveledPayload{
		Timestamp: e.TS,
	}

	// Attempt to parse caller (eg. raft/raft.go:765) into file/line (eg.
	// raft.go 765).
	if len(e.Caller) > 0 {
		parts := strings.Split(e.Caller, "/")
		fileLine := parts[len(parts)-1]
		parts = strings.Split(fileLine, ":")
		if len(parts) == 2 {
			out.File = parts[0]
			if line, err := strconv.ParseInt(parts[1], 10, 32); err == nil {
				out.Line = int(line)
			}
		}
	}

	// Convert zap level into logtree severity.
	switch e.Level {
	case "info":
		out.Severity = logging.INFO
	case "warn":
		out.Severity = logging.WARNING
	case "error":
		out.Severity = logging.ERROR
	case "fatal", "panic", "dpanic":
		out.Severity = logging.FATAL
	}

	// Sort extra keys alphabetically.
	extraKeys := make([]string, 0, len(e.Extras))
	for k := range e.Extras {
		extraKeys = append(extraKeys, k)
	}
	sort.Strings(extraKeys)

	// Convert structured extras into a human-friendly logline. We will
	// comma-join the received message and any structured logging data after
	// it.
	parts := make([]string, 0, len(e.Extras)+1)
	parts = append(parts, e.Message)
	for _, k := range extraKeys {

		// Format the value for logs. We elect to use JSON for representing
		// each element, as:
		// - this quotes strings
		// - all the data we retrieved must already be representable in JSON,
		//   as we just decoded it from an existing blob.
		// - the extra data might be arbitrarily nested (eg. an array or
		//   object) and we don't want to be in the business of coming up with
		//   our own serialization format in case of such data.
		var v string
		vbytes, err := json.Marshal(e.Extras[k])
		if err != nil {
			// Fall back to +%v just in case. We don't make any API promises
			// that the log line will be machine parseable or in any stable
			// format.
			v = fmt.Sprintf("%+v", v)
		} else {
			v = string(vbytes)
		}
		extra := fmt.Sprintf("%s: %s", k, v)

		parts = append(parts, extra)
	}

	// If the given message was empty and there are some extra data attached,
	// explicitly state that the message was empty (to avoid a mysterious
	// leading comma).
	// Otherwise, if the message was empty and there was no extra structured
	// data, assume that the sender intended to have it represented as an empty
	// line.
	if len(parts) > 1 && parts[0] == "" {
		parts[0] = "<empty>"
	}

	// Finally build the message line to emit in leveled logging and emit it.
	out.Message = strings.Join(parts, ", ")

	write(&out)
}
