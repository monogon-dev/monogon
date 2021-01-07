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
	"strings"

	"source.monogon.dev/metropolis/pkg/logbuffer"
	apb "source.monogon.dev/metropolis/proto/api"
)

// LogEntry contains a log entry, combining both leveled and raw logging into a single stream of events. A LogEntry
// will contain exactly one of either LeveledPayload or RawPayload.
type LogEntry struct {
	// If non-nil, this is a leveled logging entry.
	Leveled *LeveledPayload
	// If non-nil, this is a raw logging entry line.
	Raw *logbuffer.Line
	// DN from which this entry was logged.
	DN DN
}

// String returns a canonical representation of this payload as a single string prefixed with metadata. If the entry is
// a leveled log entry that originally was logged with newlines this representation will also contain newlines, with
// each original message part prefixed by the metadata.
// For an alternative call that will instead return a canonical prefix and a list of lines in the message, see Strings().
func (l *LogEntry) String() string {
	if l.Leveled != nil {
		prefix, messages := l.Leveled.Strings()
		res := make([]string, len(messages))
		for i, m := range messages {
			res[i] = fmt.Sprintf("%-32s %s%s", l.DN, prefix, m)
		}
		return strings.Join(res, "\n")
	}
	if l.Raw != nil {
		return fmt.Sprintf("%-32s R %s", l.DN, l.Raw)
	}
	return "INVALID"
}

// Strings returns the canonical representation of this payload split into a prefix and all lines that were contained in
// the original message. This is meant to be displayed to the user by showing the prefix before each line, concatenated
// together - possibly in a table form with the prefixes all unified with a rowspan-like mechanism.
//
// For example, this function can return:
//   prefix = "root.foo.bar                    I1102 17:20:06.921395     0 foo.go:42] "
//   lines = []string{"current tags:", " - one", " - two"}
//
// With this data, the result should be presented to users this way in text form:
// root.foo.bar                    I1102 17:20:06.921395 foo.go:42] current tags:
// root.foo.bar                    I1102 17:20:06.921395 foo.go:42]  - one
// root.foo.bar                    I1102 17:20:06.921395 foo.go:42]  - two
//
// Or, in a table layout:
// .-------------------------------------------------------------------------------------.
// | root.foo.bar                    I1102 17:20:06.921395 foo.go:42] : current tags:    |
// |                                                                  :------------------|
// |                                                                  :  - one           |
// |                                                                  :------------------|
// |                                                                  :  - two           |
// '-------------------------------------------------------------------------------------'
//
func (l *LogEntry) Strings() (prefix string, lines []string) {
	if l.Leveled != nil {
		prefix, messages := l.Leveled.Strings()
		prefix = fmt.Sprintf("%-32s %s", l.DN, prefix)
		return prefix, messages
	}
	if l.Raw != nil {
		return fmt.Sprintf("%-32s R ", l.DN), []string{l.Raw.Data}
	}
	return "INVALID ", []string{"INVALID"}
}

// Convert this LogEntry to proto. Returned value may be nil if given LogEntry is invalid, eg. contains neither a Raw
// nor Leveled entry.
func (l *LogEntry) Proto() *apb.LogEntry {
	p := &apb.LogEntry{
		Dn: string(l.DN),
	}
	switch {
	case l.Leveled != nil:
		leveled := l.Leveled
		p.Kind = &apb.LogEntry_Leveled_{
			Leveled: leveled.Proto(),
		}
	case l.Raw != nil:
		raw := l.Raw
		p.Kind = &apb.LogEntry_Raw_{
			Raw: raw.ProtoLog(),
		}
	default:
		return nil
	}
	return p
}

// Parse a proto LogEntry back into internal structure. This can be used in log proto API consumers to easily print
// received log entries.
func LogEntryFromProto(l *apb.LogEntry) (*LogEntry, error) {
	dn := DN(l.Dn)
	if _, err := dn.Path(); err != nil {
		return nil, fmt.Errorf("could not convert DN: %w", err)
	}
	res := &LogEntry{
		DN: dn,
	}
	switch inner := l.Kind.(type) {
	case *apb.LogEntry_Leveled_:
		leveled, err := LeveledPayloadFromProto(inner.Leveled)
		if err != nil {
			return nil, fmt.Errorf("could not convert leveled entry: %w", err)
		}
		res.Leveled = leveled
	case *apb.LogEntry_Raw_:
		line, err := logbuffer.LineFromLogProto(inner.Raw)
		if err != nil {
			return nil, fmt.Errorf("could not convert raw entry: %w", err)
		}
		res.Raw = line
	default:
		return nil, fmt.Errorf("proto has neither Leveled nor Raw set")
	}
	return res, nil
}
