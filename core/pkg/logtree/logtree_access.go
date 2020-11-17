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
	"errors"
	"fmt"
	"strings"
	"sync/atomic"

	"git.monogon.dev/source/nexantic.git/core/pkg/logbuffer"
	apb "git.monogon.dev/source/nexantic.git/core/proto/api"
)

// LogReadOption describes options for the LogTree.Read call.
type LogReadOption struct {
	withChildren               bool
	withStream                 bool
	withBacklog                int
	onlyLeveled                bool
	onlyRaw                    bool
	leveledWithMinimumSeverity Severity
}

// WithChildren makes Read return/stream data for both a given DN and all its children.
func WithChildren() LogReadOption { return LogReadOption{withChildren: true} }

// WithStream makes Read return a stream of data. This works alongside WithBacklog to create a read-and-stream
// construct.
func WithStream() LogReadOption { return LogReadOption{withStream: true} }

// WithBacklog makes Read return already recorded log entries, up to count elements.
func WithBacklog(count int) LogReadOption { return LogReadOption{withBacklog: count} }

// BacklogAllAvailable makes WithBacklog return all backlogged log data that logtree possesses.
const BacklogAllAvailable int = -1

func OnlyRaw() LogReadOption { return LogReadOption{onlyRaw: true} }

func OnlyLeveled() LogReadOption { return LogReadOption{onlyLeveled: true} }

// LeveledWithMinimumSeverity makes Read return only log entries that are at least at a given Severity. If only leveled
// entries are needed, OnlyLeveled must be used. This is a no-op when OnlyRaw is used.
func LeveledWithMinimumSeverity(s Severity) LogReadOption {
	return LogReadOption{leveledWithMinimumSeverity: s}
}

// LogReader permits reading an already existing backlog of log entries and to stream further ones.
type LogReader struct {
	// Backlog are the log entries already logged by LogTree. This will only be set if WithBacklog has been passed to
	// Read.
	Backlog []*LogEntry
	// Stream is a channel of new entries as received live by LogTree. This will only be set if WithStream has been
	// passed to Read. In this case, entries from this channel must be read as fast as possible by the consumer in order
	// to prevent missing entries.
	Stream <-chan *LogEntry
	// done is channel used to signal (by closing) that the log consumer is not interested in more Stream data.
	done chan<- struct{}
	// missed is an atomic integer pointer that tells the subscriber how many messages in Stream they missed. This
	// pointer is nil if no streaming has been requested.
	missed *uint64
}

// Missed returns the amount of entries that were missed from Stream (as the channel was not drained fast enough).
func (l *LogReader) Missed() uint64 {
	// No Stream.
	if l.missed == nil {
		return 0
	}
	return atomic.LoadUint64(l.missed)
}

// Close closes the LogReader's Stream. This must be called once the Reader does not wish to receive streaming messages
// anymore.
func (l *LogReader) Close() {
	if l.done != nil {
		close(l.done)
	}
}

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

var (
	ErrRawAndLeveled = errors.New("cannot return logs that are simultaneously OnlyRaw and OnlyLeveled")
)

// Read and/or stream entries from a LogTree. The returned LogReader is influenced by the LogReadOptions passed, which
// influence whether the Read will return existing entries, a stream, or both. In addition the options also dictate
// whether only entries for that particular DN are returned, or for all sub-DNs as well.
func (l *LogTree) Read(dn DN, opts ...LogReadOption) (*LogReader, error) {
	l.journal.mu.RLock()
	defer l.journal.mu.RUnlock()

	var backlog int
	var stream bool
	var recursive bool
	var leveledSeverity Severity
	var onlyRaw, onlyLeveled bool

	for _, opt := range opts {
		if opt.withBacklog > 0 || opt.withBacklog == BacklogAllAvailable {
			backlog = opt.withBacklog
		}
		if opt.withStream {
			stream = true
		}
		if opt.withChildren {
			recursive = true
		}
		if opt.leveledWithMinimumSeverity != "" {
			leveledSeverity = opt.leveledWithMinimumSeverity
		}
		if opt.onlyLeveled {
			onlyLeveled = true
		}
		if opt.onlyRaw {
			onlyRaw = true
		}
	}

	if onlyLeveled && onlyRaw {
		return nil, ErrRawAndLeveled
	}

	var filters []filter
	if onlyLeveled {
		filters = append(filters, filterOnlyLeveled)
	}
	if onlyRaw {
		filters = append(filters, filterOnlyRaw)
	}
	if recursive {
		filters = append(filters, filterSubtree(dn))
	} else {
		filters = append(filters, filterExact(dn))
	}
	if leveledSeverity != "" {
		filters = append(filters, filterSeverity(leveledSeverity))
	}

	var entries []*entry
	if backlog > 0 || backlog == BacklogAllAvailable {
		// TODO(q3k): pass over the backlog count to scanEntries/getEntries, instead of discarding them here.
		if recursive {
			entries = l.journal.scanEntries(filters...)
		} else {
			entries = l.journal.getEntries(dn, filters...)
		}
		if backlog != BacklogAllAvailable && len(entries) > backlog {
			entries = entries[:backlog]
		}
	}

	var sub *subscriber
	if stream {
		sub = &subscriber{
			// TODO(q3k): make buffer size configurable
			dataC:   make(chan *LogEntry, 128),
			doneC:   make(chan struct{}),
			filters: filters,
		}
		l.journal.subscribe(sub)
	}

	lr := &LogReader{}
	lr.Backlog = make([]*LogEntry, len(entries))
	for i, entry := range entries {
		lr.Backlog[i] = entry.external()
	}
	if stream {
		lr.Stream = sub.dataC
		lr.done = sub.doneC
		lr.missed = &sub.missed
	}
	return lr, nil
}
