package unraw

import (
	"context"
	"errors"
	"fmt"
	"os"
	"syscall"
	"testing"

	"source.monogon.dev/metropolis/pkg/logbuffer"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

func testParser(l *logbuffer.Line, w LeveledWriter) {
	w(&logtree.ExternalLeveledPayload{
		Message: l.Data,
	})
}

func TestNamedPipeReader(t *testing.T) {
	dir, err := os.MkdirTemp("/tmp", "metropolis-test-named-pipe-reader")
	if err != nil {
		t.Fatalf("could not create tempdir: %v", err)
	}
	defer os.RemoveAll(dir)
	fifoPath := dir + "/fifo"

	// Start named pipe reader.
	started := make(chan struct{})
	stop, lt := supervisor.TestHarness(t, func(ctx context.Context) error {
		converter := Converter{
			Parser:        testParser,
			LeveledLogger: supervisor.Logger(ctx),
		}

		r, err := converter.NamedPipeReader(fifoPath)
		if err != nil {
			return fmt.Errorf("could not create pipe reader: %w", err)
		}
		close(started)
		return r(ctx)
	})

	<-started

	// Open FIFO...
	f, err := os.OpenFile(fifoPath, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("could not open fifo: %v", err)
	}

	// Start reading all logs.
	reader, err := lt.Read("root", logtree.WithChildren(), logtree.WithStream())
	if err != nil {
		t.Fatalf("could not get logtree reader: %v", err)
	}
	defer reader.Close()

	// Write two lines to the fifo.
	fmt.Fprintf(f, "foo\nbar\n")
	f.Close()

	// Expect lines to end up in logtree.
	if got, want := (<-reader.Stream).Leveled.MessagesJoined(), "foo"; want != got {
		t.Errorf("expected first message to be %q, got %q", want, got)
	}
	if got, want := (<-reader.Stream).Leveled.MessagesJoined(), "bar"; want != got {
		t.Errorf("expected second message to be %q, got %q", want, got)
	}

	// Fully restart the entire supervisor and pipe reader, redo test, things
	// should continue to work.
	stop()

	// Block until FIFO isn't being read anymore. This ensures that the
	// NamedPipeReader actually stopped running, otherwise the following write to
	// the fifo can race by writing to the old NamedPipeReader and making the test
	// time out. This can also happen in production, but that will just cause us to
	// lose piped data in the very small race window when this can happen
	// (statistically in this test, <0.1%).
	//
	// The check is being done by opening the FIFO in 'non-blocking mode', which
	// returns ENXIO immediately if the FIFO has no corresponding writer, and
	// succeeds otherwise.
	for {
		ft, err := os.OpenFile(fifoPath, os.O_WRONLY|syscall.O_NONBLOCK, 0)
		if err == nil {
			// There's still a writer, keep trying.
			ft.Close()
		} else if errors.Is(err, syscall.ENXIO) {
			// No writer, break.
			break
		} else {
			// Something else?
			t.Fatalf("OpenFile(%q): %v", fifoPath, err)
		}
	}

	started = make(chan struct{})
	stop, lt = supervisor.TestHarness(t, func(ctx context.Context) error {
		converter := Converter{
			Parser:        testParser,
			LeveledLogger: supervisor.Logger(ctx),
		}

		r, err := converter.NamedPipeReader(fifoPath)
		if err != nil {
			return fmt.Errorf("could not create pipe reader: %w", err)
		}
		close(started)
		return r(ctx)
	})

	<-started

	// Start reading all logs.
	reader, err = lt.Read("root", logtree.WithChildren(), logtree.WithStream())
	if err != nil {
		t.Fatalf("could not get logtree reader: %v", err)
	}
	defer reader.Close()

	// Write line to the fifo.
	f, err = os.OpenFile(fifoPath, os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("could not open fifo: %v", err)
	}
	fmt.Fprintf(f, "baz\n")
	f.Close()

	// Expect lines to end up in logtree.
	if got, want := (<-reader.Stream).Leveled.MessagesJoined(), "baz"; want != got {
		t.Errorf("expected first message to be %q, got %q", want, got)
	}
}
