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

package logtree

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tpb "google.golang.org/protobuf/types/known/timestamppb"

	apb "source.monogon.dev/metropolis/proto/api"
)

// LeveledPayload is a log entry for leveled logs (as per leveled.go). It contains
// the input to these calls (severity and message split into newline-delimited
// messages) and additional metadata that would be usually seen in a text
// representation of a leveled log entry.
type LeveledPayload struct {
	// messages is the list of messages contained in this payload. This list is built
	// from splitting up the given message from the user by newline.
	messages []string
	// timestamp is the time at which this message was emitted.
	timestamp time.Time
	// severity is the leveled Severity at which this message was emitted.
	severity Severity
	// file is the filename of the caller that emitted this message.
	file string
	// line is the line number within the file of the caller that emitted this message.
	line int
}

// String returns a canonical representation of this payload as a single string
// prefixed with metadata. If the original message was logged with newlines, this
// representation will also contain newlines, with each original message part
// prefixed by the metadata. For an alternative call that will instead return a
// canonical prefix and a list of lines in the message, see Strings().
func (p *LeveledPayload) String() string {
	prefix, lines := p.Strings()
	res := make([]string, len(p.messages))
	for i, line := range lines {
		res[i] = fmt.Sprintf("%s%s", prefix, line)
	}
	return strings.Join(res, "\n")
}

// Strings returns the canonical representation of this payload split into a
// prefix and all lines that were contained in the original message. This is
// meant to be displayed to the user by showing the prefix before each line,
// concatenated together - possibly in a table form with the prefixes all
// unified with a rowspan- like mechanism.
//
// For example, this function can return:
//   prefix = "I1102 17:20:06.921395 foo.go:42] "
//   lines = []string{"current tags:", " - one", " - two"}
//
// With this data, the result should be presented to users this way in text form:
// I1102 17:20:06.921395 foo.go:42] current tags:
// I1102 17:20:06.921395 foo.go:42]  - one
// I1102 17:20:06.921395 foo.go:42]  - two
//
// Or, in a table layout:
// .-----------------------------------------------------------.
// | I1102 17:20:06.921395     0 foo.go:42] : current tags:    |
// |                                        :------------------|
// |                                        :  - one           |
// |                                        :------------------|
// |                                        :  - two           |
// '-----------------------------------------------------------'
func (p *LeveledPayload) Strings() (prefix string, lines []string) {
	_, month, day := p.timestamp.Date()
	hour, minute, second := p.timestamp.Clock()
	nsec := p.timestamp.Nanosecond() / 1000

	// Same format as in glog, but without treadid.
	// Lmmdd hh:mm:ss.uuuuuu file:line]
	// TODO(q3k): rewrite this to printf-less code.
	prefix = fmt.Sprintf("%s%02d%02d %02d:%02d:%02d.%06d %s:%d] ", p.severity, month, day, hour, minute, second, nsec, p.file, p.line)

	lines = p.messages
	return
}

// Message returns the inner message lines of this entry, ie. what was passed to
// the actual logging method, but split by newlines.
func (p *LeveledPayload) Messages() []string { return p.messages }

func (p *LeveledPayload) MessagesJoined() string { return strings.Join(p.messages, "\n") }

// Timestamp returns the time at which this entry was logged.
func (p *LeveledPayload) Timestamp() time.Time { return p.timestamp }

// Location returns a string in the form of file_name:line_number that shows the
// origin of the log entry in the program source.
func (p *LeveledPayload) Location() string { return fmt.Sprintf("%s:%d", p.file, p.line) }

// Severity returns the Severity with which this entry was logged.
func (p *LeveledPayload) Severity() Severity { return p.severity }

// Proto converts a LeveledPayload to protobuf format.
func (p *LeveledPayload) Proto() *apb.LogEntry_Leveled {
	return &apb.LogEntry_Leveled{
		Lines:     p.Messages(),
		Timestamp: tpb.New(p.Timestamp()),
		Severity:  p.Severity().ToProto(),
		Location:  p.Location(),
	}
}

// LeveledPayloadFromProto parses a protobuf message into the internal format.
func LeveledPayloadFromProto(p *apb.LogEntry_Leveled) (*LeveledPayload, error) {
	severity, err := SeverityFromProto(p.Severity)
	if err != nil {
		return nil, fmt.Errorf("could not convert severity: %w", err)
	}
	parts := strings.Split(p.Location, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid location, must be two :-delimited parts, is %d parts", len(parts))
	}
	file := parts[0]
	line, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid location line number: %w", err)
	}
	return &LeveledPayload{
		messages:  p.Lines,
		timestamp: p.Timestamp.AsTime(),
		severity:  severity,
		file:      file,
		line:      line,
	}, nil
}

// ExternalLeveledPayload is a LeveledPayload received from an external source,
// eg. from parsing the logging output of third-party programs. It can be
// converted into a LeveledPayload and inserted into a leveled logger, but will
// be sanitized before that, ensuring that potentially buggy
// emitters/converters do not end up polluting the leveled logger data.
//
// This type should be used only when inserting data from external systems, not
// by code that just wishes to log things. In the future, data inserted this
// way might be explicitly marked as tainted so operators can understand that
// parts of this data might not give the same guarantees as the log entries
// emitted by the native LeveledLogger API.
type ExternalLeveledPayload struct {
	// Log line. If any newlines are found, they will split the message into
	// multiple messages within LeveledPayload. Empty messages are accepted
	// verbatim.
	Message string
	// Timestamp when this payload was emitted according to its source. If not
	// given, will default to the time of conversion to LeveledPayload.
	Timestamp time.Time
	// Log severity. If invalid or unset will default to INFO.
	Severity Severity
	// File name of originating code. Defaults to "unknown" if not set.
	File string
	// Line in File. Zero indicates the line is not known.
	Line int
}

// sanitize the given ExternalLeveledPayload by creating a corresponding
// LeveledPayload. The original object is unaltered.
func (e *ExternalLeveledPayload) sanitize() *LeveledPayload {
	l := &LeveledPayload{
		messages:  strings.Split(e.Message, "\n"),
		timestamp: e.Timestamp,
		severity:  e.Severity,
		file:      e.File,
		line:      e.Line,
	}
	if l.timestamp.IsZero() {
		l.timestamp = time.Now()
	}
	if !l.severity.Valid() {
		l.severity = INFO
	}
	if l.file == "" {
		l.file = "unknown"
	}
	if l.line < 0 {
		l.line = 0
	}
	return l
}
