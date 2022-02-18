package rpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/pkg/logtree"
)

// Span implements a compatible subset of
// go.opentelemetry.io/otel/trace.Span.

// It is used in place of trace.Span until opentelemetry support
// is fully implemented and thus the library is pulled in. Once
// that happens, all relevant methods will be replace with an
// embedding of the trace.Span interface.
type Span interface {
	// End() not implemented.

	// AddEvent adds an event with the provided name.
	//
	// Changed from otel/trace.Span: no options.
	AddEvent(name string)

	// IsRecording returns the recording state of the Span. It will return true if
	// the Span is active and events can be recorded.
	IsRecording() bool

	// RecordError() not implemented.

	// SpanContext() not implemented.

	// SetStatus() not implemented.

	// SetName() not implemented.

	// SetAttributes() not implemented.

	// TraceProvider() not implemented.

	// Monogon extensions follow. These call into standard otel.Span methods
	// (and effectively underlying model), but provide tighter API for
	// Metropolis.

	// Printf adds an event via AddEvent after performing a string format expansion
	// via fmt.Sprintf. The formatting is performed during the call if the span is
	// recording, or never if it isn't.
	Printf(format string, a ...interface{})
}

// logtreeSpan is an implementation of Span which just forwards events into a
// local logtree LeveledLogger. All spans are always recording.
//
// This is a stop-gap implementation to introduce gRPC trace-based
// logging/metrics into Metropolis which can then be evolved into a full-blown
// opentelemetry implementation.
type logtreeSpan struct {
	// logger is the logtree LeveledLogger backing this span. All Events added into
	// the Span will go straight into that logger. If the logger is nil, all events
	// will be dropped instead.
	logger logtree.LeveledLogger
	// uid is the span ID of this logtreeSpan. Currently this is a monotonic counter
	// based on the current nanosecond epoch, but this might change in the future.
	// This field is ignored if logger is nil.
	uid uint64
}

func newLogtreeSpan(l logtree.LeveledLogger) *logtreeSpan {
	uid := uint64(time.Now().UnixNano())
	return &logtreeSpan{
		logger: l,
		uid:    uid,
	}
}

func (l *logtreeSpan) AddEvent(name string) {
	if l.logger == nil {
		return
	}
	l.logger.WithAddedStackDepth(1).Infof("Span %x: %s", l.uid, name)
}

func (l *logtreeSpan) Printf(format string, a ...interface{}) {
	if l.logger == nil {
		return
	}
	msg := fmt.Sprintf(format, a...)
	l.logger.WithAddedStackDepth(1).Infof("Span %x: %s", l.uid, msg)
}

func (l *logtreeSpan) IsRecording() bool {
	return l.logger != nil
}

type spanKey string

var spanKeyValue spanKey = "metropolis-trace-span"

// contextWithSpan wraps a given context with a given logtreeSpan. This
// logtreeSpan will be returned by Trace() calls on the returned context.
func contextWithSpan(ctx context.Context, s *logtreeSpan) context.Context {
	return context.WithValue(ctx, spanKeyValue, s)
}

// Trace returns the active Span for the current Go context. If no Span was set
// up for this context, an inactive/empty span object is returned, on which
// every operation is a no-op.
func Trace(ctx context.Context) Span {
	v := ctx.Value(spanKeyValue)
	if v == nil {
		return &logtreeSpan{}
	}
	if s, ok := v.(*logtreeSpan); ok {
		return s
	}
	return &logtreeSpan{}
}

// spanServerStream is a grpc.ServerStream wrapper which contains some
// logtreeSpan, and returns it as part of the Context() of the ServerStream. It
// also intercepts SendMsg/RecvMsg and logs them to the same span.
type spanServerStream struct {
	grpc.ServerStream
	span *logtreeSpan
}

func (s *spanServerStream) Context() context.Context {
	return contextWithSpan(s.ServerStream.Context(), s.span)
}

func (s *spanServerStream) SendMsg(m interface{}) error {
	s.span.Printf("RPC send: %s", protoMessagePretty(m))
	return s.ServerStream.SendMsg(m)
}

func (s *spanServerStream) RecvMsg(m interface{}) error {
	err := s.ServerStream.RecvMsg(m)
	s.span.Printf("RPC recv: %s", protoMessagePretty(m))
	return err
}

// protoMessagePretty attempts to pretty-print a given proto message into a
// one-line string. The returned format is not guaranteed to be stable, and is
// only intended to be used for debug purposes by operators.
//
// TODO(q3k): make this not print any confidential fields (once we have any),
// eg. via extensions/annotations.
func protoMessagePretty(m interface{}) string {
	if m == nil {
		return "nil"
	}
	v, ok := m.(proto.Message)
	if !ok {
		return "invalid"
	}
	name := string(v.ProtoReflect().Type().Descriptor().Name())
	bytes, err := prototext.Marshal(v)
	if err != nil {
		return name
	}
	return fmt.Sprintf("%s: %s", name, strings.ReplaceAll(string(bytes), "\n", " "))
}
