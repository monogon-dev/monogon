// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package logtree

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"source.monogon.dev/go/logging"
	"source.monogon.dev/osbase/logbuffer"
)

// Zapify turns a LeveledLogger into a zap.Logger which pipes its output into the
// LeveledLogger. The message, severity and caller are carried over. Extra fields
// are appended as JSON to the end of the log line.
func Zapify(logger logging.Leveled, minimumLevel zapcore.Level) *zap.Logger {
	p, ok := logger.(*leveledPublisher)
	if !ok {
		// Fail fast, as this is a programming error.
		panic("Expected *leveledPublisher in LeveledLogger from supervisor")
	}

	ec := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.EpochTimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	s := zapSink{
		publisher: p,
	}
	s.buffer = logbuffer.NewLineBuffer(4096, s.consumeLine)
	zc := zapcore.NewCore(zapcore.NewJSONEncoder(ec), s.buffer, minimumLevel)
	return zap.New(zc, zap.AddCaller())
}

type zapSink struct {
	publisher *leveledPublisher
	buffer    *logbuffer.LineBuffer
}

func (z *zapSink) consumeLine(l *logbuffer.Line) {
	ze, err := parseZapJSON(l.Data)
	if err != nil {
		z.publisher.Warningf("failed to parse zap JSON: %v: %q", err, l.Data)
		return
	}
	message := ze.message
	if len(ze.extra) > 0 {
		message += " " + ze.extra
	}
	e := &entry{
		origin: z.publisher.node.dn,
		leveled: &LeveledPayload{
			timestamp: ze.time,
			severity:  ze.severity,
			messages:  []string{message},
			file:      ze.file,
			line:      ze.line,
		},
	}
	z.publisher.node.tree.journal.append(e)
	z.publisher.node.tree.journal.notify(e)
}

type zapEntry struct {
	message  string
	severity logging.Severity
	time     time.Time
	file     string
	line     int
	extra    string
}

func parseZapJSON(s string) (*zapEntry, error) {
	entry := make(map[string]any)
	if err := json.Unmarshal([]byte(s), &entry); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	message, ok := entry["message"].(string)
	if !ok {
		return nil, fmt.Errorf("no message field")
	}
	level, ok := entry["level"].(string)
	if !ok {
		return nil, fmt.Errorf("no level field")
	}
	t, ok := entry["time"].(float64)
	if !ok {
		return nil, fmt.Errorf("no time field")
	}
	caller, ok := entry["caller"].(string)
	if !ok {
		return nil, fmt.Errorf("no caller field")
	}

	callerParts := strings.Split(caller, ":")
	if len(callerParts) != 2 {
		return nil, fmt.Errorf("invalid caller")
	}
	callerDirFile := strings.Split(callerParts[0], "/")
	callerFile := callerDirFile[len(callerDirFile)-1]
	callerLineS := callerParts[1]
	callerLine, _ := strconv.Atoi(callerLineS)

	var severity logging.Severity
	switch level {
	case "warn":
		severity = logging.WARNING
	case "error", "dpanic", "panic", "fatal":
		severity = logging.ERROR
	default:
		severity = logging.INFO
	}

	secs := int64(t)
	nsecs := int64((t - float64(secs)) * 1e9)

	delete(entry, "message")
	delete(entry, "level")
	delete(entry, "time")
	delete(entry, "caller")
	var extra []byte
	if len(entry) > 0 {
		extra, _ = json.Marshal(entry)
	}
	return &zapEntry{
		message:  message,
		severity: severity,
		time:     time.Unix(secs, nsecs),
		file:     callerFile,
		line:     callerLine,
		extra:    string(extra),
	}, nil
}
