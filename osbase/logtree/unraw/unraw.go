// unraw implements a facility to convert raw logs from external sources into
// leveled logs.
//
// This is not the same as raw logging inside the logtree, which exists to
// ingest logs that are either fully arbitrary or do not map cleanly to the
// leveled logging concept. The unraw library is instead made to parse logs
// from systems that also use leveled logs internally, but emit them to a
// serialized byte stream that then needs to be turned back into something
// leveled inside metropolis.
//
// Logs converted this way are unfortunately lossy and do not come with the
// same guarantees as logs directly emitted via logtree. For example, there's
// no built-in protection against systems emiting fudged timestamps or file
// locations. Thus, this functionality should be used to interact with trusted
// systems, not fully arbitrary logs.
package unraw

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"syscall"
	"time"

	"source.monogon.dev/osbase/logbuffer"
	"source.monogon.dev/osbase/logtree"
	"source.monogon.dev/osbase/supervisor"
)

// Parser is a user-defined function for converting a log line received from an
// external system into a leveled logging payload.
// The given LeveledWriter should be called for every leveled log entry that
// results from this line. This means that a parser might skip some lines, or
// emit multiple leveled payloads per line.
type Parser func(*logbuffer.Line, LeveledWriter)

// Converter is the main entrypoint of the unraw library. It wraps a
// LeveledLogger in combination with a Parser to create an io.Writer that can
// be sent raw log data.
type Converter struct {
	// Parser is the user-defined parsing function for converting log lines
	// into leveled logging payloads. This must be set.
	Parser Parser
	// MaximumLineLength is the maximum length of a single log line when
	// splitting incoming writes into lines. If a line is longer than this, it
	// will be truncated (and will be sent to the Parser regardless).
	//
	// If not set, this defaults to 1024 bytes.
	MaximumLineLength int
	// LeveledLogger is the logtree leveled logger into which events from the
	// Parser will be sent.
	LeveledLogger logtree.LeveledLogger

	// mu guards lb.
	mu sync.Mutex
	// lb is the underlying line buffer used to split incoming data into lines.
	// It will be initialized on first Write.
	lb *logbuffer.LineBuffer
}

// LeveledWriter is called by a Parser for every ExternelLeveledPayload it
// wishes to emit into a backing LeveledLogger. If the payload is missing some
// fields, these will default to some sensible values - see the
// ExternalLeveledPayload structure definition for more information.
type LeveledWriter func(*logtree.ExternalLeveledPayload)

// Write implements io.Writer. Any write performed into the Converter will
// populate the converter's internal buffer, and any time that buffer contains
// a full line it will be sent over to the Parser for processing.
func (e *Converter) Write(p []byte) (int, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.MaximumLineLength <= 0 {
		e.MaximumLineLength = 1024
	}
	if e.lb == nil {
		e.lb = logbuffer.NewLineBuffer(e.MaximumLineLength, func(l *logbuffer.Line) {
			e.Parser(l, e.insert)
		})
	}
	return e.lb.Write(p)
}

// insert implements LeveledWriter.
func (e *Converter) insert(d *logtree.ExternalLeveledPayload) {
	if err := logtree.LogExternalLeveled(e.LeveledLogger, d); err != nil {
		e.LeveledLogger.Fatal("Could not insert unrawed entry: %v", err)
	}
}

// NamedPipeReader returns a supervisor runnable that continously reads logs
// from the given path and attempts to parse them into leveled logs using this
// Converter.
//
// If the given path doesn't exist, a named pipe will be created there before
// the function exits. This guarantee means that as long as any writing process
// is not started before NamedPipeReader returns ther is no need to
// remove/recreate the named pipe.
//
// TODO(q3k): defer the creation of the FIFO to localstorage so this doesn't
// need to be taken care of in the first place.
func (e *Converter) NamedPipeReader(path string) (supervisor.Runnable, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := syscall.Mkfifo(path, 0666); err != nil {
			return nil, fmt.Errorf("when creating named pipe: %w", err)
		}
	}
	return func(ctx context.Context) error {
		fifo, err := os.OpenFile(path, os.O_RDONLY, os.ModeNamedPipe)
		if err != nil {
			return fmt.Errorf("when opening named pipe: %w", err)
		}
		go func() {
			<-ctx.Done()
			fifo.Close()
		}()
		defer fifo.Close()
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		for {
			// Quit if requested.
			if ctx.Err() != nil {
				return ctx.Err()
			}

			n, err := io.Copy(e, fifo)
			if n == 0 && err == nil {
				// Hack because pipes/FIFOs can return zero reads when nobody
				// is writing. To avoid busy-looping, sleep a bit before
				// retrying. This does not loose data since the FIFO internal
				// buffer will stall writes when it becomes full. 10ms maximum
				// stall in a non-latency critical process (reading debug logs)
				// is not an issue for us.
				time.Sleep(10 * time.Millisecond)
			} else if err != nil {
				// Since we close fifo on context cancel, we'll get a 'file is already closed'
				// io error here. Translate that over to the context error that caused it.
				if ctx.Err() != nil {
					return ctx.Err()
				}
				return fmt.Errorf("log pump failed: %w", err)
			}

		}
	}, nil
}
