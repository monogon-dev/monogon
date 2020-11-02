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

	apb "git.monogon.dev/source/nexantic.git/core/proto/api"
)

// LeveledPayload is a log entry for leveled logs (as per leveled.go). It contains not only the log message itself and
// its severity, but also additional metadata that would be usually seen in a text representation of a leveled log entry.
type LeveledPayload struct {
	// message is the log message, rendered from a leveled log call like Infof(), Warningf(), ...
	message string
	// timestamp is the time at which this message was emitted.
	timestamp time.Time
	// severity is the leveled Severity at which this message was emitted.
	severity Severity
	// file is the filename of the caller that emitted this message.
	file string
	// line is the line number within the file of the caller that emitted this message.
	line int
}

func (p *LeveledPayload) String() string {
	// Same format as in glog:
	// Lmmdd hh:mm:ss.uuuuuu threadid file:line]
	// Except, threadid is (currently) always zero. In the future this field might be used for something different.

	_, month, day := p.timestamp.Date()
	hour, minute, second := p.timestamp.Clock()
	nsec := p.timestamp.Nanosecond() / 1000

	// TODO(q3k): rewrite this to printf-less code.
	return fmt.Sprintf("%s%02d%02d %02d:%02d:%02d.%06d % 7d %s:%d] %s", p.severity, month, day, hour, minute, second,
		nsec, 0, p.file, p.line, p.message)
}

// Message returns the inner message of this entry, ie. what was passed to the actual logging method.
func (p *LeveledPayload) Message() string { return p.message }

// Timestamp returns the time at which this entry was logged.
func (p *LeveledPayload) Timestamp() time.Time { return p.timestamp }

// Location returns a string in the form of file_name:line_number that shows the origin of the log entry in the
// program source.
func (p *LeveledPayload) Location() string { return fmt.Sprintf("%s:%d", p.file, p.line) }

// Severity returns the Severity with which this entry was logged.
func (p *LeveledPayload) Severity() Severity { return p.severity }

// Proto converts a LeveledPayload to protobuf format.
func (p *LeveledPayload) Proto() *apb.LogEntry_Leveled {
	return &apb.LogEntry_Leveled{
		Message:   p.Message(),
		Timestamp: p.Timestamp().UnixNano(),
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
		message:   p.Message,
		timestamp: time.Unix(0, p.Timestamp),
		severity:  severity,
		file:      file,
		line:      line,
	}, nil
}
