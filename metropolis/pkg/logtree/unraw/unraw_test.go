package unraw

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
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
	dir, err := ioutil.TempDir("", "metropolis-test-named-pipe-reader")
	if err != nil {
		t.Fatalf("could not create tempdir: %v", err)
	}
	//defer os.RemoveAll(dir)
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

	// Wait until NamedPipeReader returns to make sure the fifo was created..
	<-started

	// Start reading all logs.
	reader, err := lt.Read("root", logtree.WithChildren(), logtree.WithStream())
	if err != nil {
		t.Fatalf("could not get logtree reader: %v", err)
	}
	defer reader.Close()

	// Write two lines to the fifo.
	f, err := os.OpenFile(fifoPath, os.O_RDWR, 0)
	if err != nil {
		t.Fatalf("could not open fifo: %v", err)
	}
	fmt.Fprintf(f, "foo\nbar\n")
	f.Close()

	// Expect lines to end up in logtree.
	if got, want := (<-reader.Stream).Leveled.MessagesJoined(), "foo"; want != got {
		t.Errorf("expected first message to be %q, got %q", want, got)
	}
	if got, want := (<-reader.Stream).Leveled.MessagesJoined(), "bar"; want != got {
		t.Errorf("expected second message to be %q, got %q", want, got)
	}

	// Fully restart the entire hypervisor and pipe reader, redo test, things
	// should continue to work.
	stop()

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

	// Start reading all logs.
	reader, err = lt.Read("root", logtree.WithChildren(), logtree.WithStream())
	if err != nil {
		t.Fatalf("could not get logtree reader: %v", err)
	}
	defer reader.Close()

	<-started

	// Write line to the fifo.
	// Write two lines to the fifo.
	f, err = os.OpenFile(fifoPath, os.O_RDWR, 0)
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
