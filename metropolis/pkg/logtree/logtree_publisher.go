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
	"io"
	"runtime"
	"strings"
	"time"

	"source.monogon.dev/metropolis/pkg/logbuffer"
)

type leveledPublisher struct {
	node  *node
	depth int
}

// LeveledFor returns a LeveledLogger publishing interface for a given DN. An error
// may be returned if the DN is malformed.
func (l *LogTree) LeveledFor(dn DN) (LeveledLogger, error) {
	node, err := l.nodeByDN(dn)
	if err != nil {
		return nil, err
	}
	return &leveledPublisher{
		node:  node,
		depth: 0,
	}, nil
}

func (l *LogTree) RawFor(dn DN) (io.Writer, error) {
	node, err := l.nodeByDN(dn)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve raw logger: %w", err)
	}
	return node.rawLineBuffer, nil
}

// MustLeveledFor returns a LeveledLogger publishing interface for a given DN, or
// panics if the given DN is invalid.
func (l *LogTree) MustLeveledFor(dn DN) LeveledLogger {
	leveled, err := l.LeveledFor(dn)
	if err != nil {
		panic(fmt.Errorf("LeveledFor returned: %w", err))
	}
	return leveled
}

func (l *LogTree) MustRawFor(dn DN) io.Writer {
	raw, err := l.RawFor(dn)
	if err != nil {
		panic(fmt.Errorf("RawFor returned: %w", err))
	}
	return raw
}

// SetVerbosity sets the verbosity for a given DN (non-recursively, ie. for that DN
// only, not its children).
func (l *LogTree) SetVerbosity(dn DN, level VerbosityLevel) error {
	node, err := l.nodeByDN(dn)
	if err != nil {
		return err
	}
	node.verbosity = level
	return nil
}

// logRaw is called by this node's LineBuffer any time a raw log line is completed.
// It will create a new entry, append it to the journal, and notify all pertinent
// subscribers.
func (n *node) logRaw(line *logbuffer.Line) {
	e := &entry{
		origin: n.dn,
		raw:    line,
	}
	n.tree.journal.append(e)
	n.tree.journal.notify(e)
}

// LogExternalLeveled injects a ExternalLeveledPayload into a given
// LeveledLogger. This should only be used by systems which translate external
// data sources into leveled logging - see ExternelLeveledPayload for more
// information.
func LogExternalLeveled(l LeveledLogger, e *ExternalLeveledPayload) error {
	publisher, ok := l.(*leveledPublisher)
	if !ok {
		return fmt.Errorf("the given LeveledLogger is not a *leveledPublisher")
	}
	p := e.sanitize()
	entry := &entry{
		origin:  publisher.node.dn,
		leveled: p,
	}
	publisher.node.tree.journal.append(entry)
	publisher.node.tree.journal.notify(entry)
	return nil
}

// log builds a LeveledPayload and entry for a given message, including all related
// metadata. It will create a new entry append it to the journal, and notify all
// pertinent subscribers.
func (l *leveledPublisher) logLeveled(depth int, severity Severity, msg string) {
	_, file, line, ok := runtime.Caller(2 + depth)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	// Remove leading/trailing newlines and split.
	messages := strings.Split(strings.Trim(msg, "\n"), "\n")

	p := &LeveledPayload{
		timestamp: time.Now(),
		severity:  severity,
		messages:  messages,
		file:      file,
		line:      line,
	}
	e := &entry{
		origin:  l.node.dn,
		leveled: p,
	}
	l.node.tree.journal.append(e)
	l.node.tree.journal.notify(e)
}

// Info implements the LeveledLogger interface.
func (l *leveledPublisher) Info(args ...interface{}) {
	l.logLeveled(l.depth, INFO, fmt.Sprint(args...))
}

// Infof implements the LeveledLogger interface.
func (l *leveledPublisher) Infof(format string, args ...interface{}) {
	l.logLeveled(l.depth, INFO, fmt.Sprintf(format, args...))
}

// Warning implements the LeveledLogger interface.
func (l *leveledPublisher) Warning(args ...interface{}) {
	l.logLeveled(l.depth, WARNING, fmt.Sprint(args...))
}

// Warningf implements the LeveledLogger interface.
func (l *leveledPublisher) Warningf(format string, args ...interface{}) {
	l.logLeveled(l.depth, WARNING, fmt.Sprintf(format, args...))
}

// Error implements the LeveledLogger interface.
func (l *leveledPublisher) Error(args ...interface{}) {
	l.logLeveled(l.depth, ERROR, fmt.Sprint(args...))
}

// Errorf implements the LeveledLogger interface.
func (l *leveledPublisher) Errorf(format string, args ...interface{}) {
	l.logLeveled(l.depth, ERROR, fmt.Sprintf(format, args...))
}

// Fatal implements the LeveledLogger interface.
func (l *leveledPublisher) Fatal(args ...interface{}) {
	l.logLeveled(l.depth, FATAL, fmt.Sprint(args...))
}

// Fatalf implements the LeveledLogger interface.
func (l *leveledPublisher) Fatalf(format string, args ...interface{}) {
	l.logLeveled(l.depth, FATAL, fmt.Sprintf(format, args...))
}

// WithAddedStackDepth impleemnts the LeveledLogger interface.
func (l *leveledPublisher) WithAddedStackDepth(depth int) LeveledLogger {
	l2 := *l
	l2.depth += depth
	return &l2
}

// V implements the LeveledLogger interface.
func (l *leveledPublisher) V(v VerbosityLevel) VerboseLeveledLogger {
	return &verbose{
		publisher: l,
		enabled:   l.node.verbosity >= v,
	}
}

// verbose implements the VerboseLeveledLogger interface. It is a thin wrapper
// around node, with an 'enabled' bool. This means that V(n)-returned
// VerboseLeveledLoggers must be short lived, as a changed in verbosity will not
// affect all already existing VerboseLeveledLoggers.
type verbose struct {
	publisher *leveledPublisher
	node      *node
	enabled   bool
}

func (v *verbose) Enabled() bool {
	return v.enabled
}

func (v *verbose) Info(args ...interface{}) {
	if !v.enabled {
		return
	}
	v.publisher.logLeveled(v.publisher.depth, INFO, fmt.Sprint(args...))
}

func (v *verbose) Infof(format string, args ...interface{}) {
	if !v.enabled {
		return
	}
	v.publisher.logLeveled(v.publisher.depth, INFO, fmt.Sprintf(format, args...))
}
