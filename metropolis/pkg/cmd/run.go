// Package cmd contains helpers that abstract away the chore of starting new
// processes, tracking their lifetime, inspecting their output, etc.
package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"source.monogon.dev/metropolis/pkg/logbuffer"
)

// RunCommand starts a new process and waits until either its completion, or
// until the supplied predicate function returns true. The function is called
// for each line produced by the new process.
//
// The process will be killed both in the event the context is cancelled, and
// when expectedOutput is found.
func RunCommand(ctx context.Context, path string, args []string, pf func(string) bool) (bool, error) {
	// Make a sub-context to ensure the process exits when this function is done.
	ctx, ctxC := context.WithCancel(ctx)
	defer ctxC()

	// Copy the stdout and stderr output to a single channel of lines so that they
	// can then be matched against expectedOutput.

	// Since LineBuffer can write its buffered contents on a deferred Close,
	// after the reader loop is broken, avoid deadlocks by making lineC a
	// buffered channel.
	lineC := make(chan string, 2)
	outBuffer := logbuffer.NewLineBuffer(1024, func(l *logbuffer.Line) {
		lineC <- l.Data
	})
	defer outBuffer.Close()
	errBuffer := logbuffer.NewLineBuffer(1024, func(l *logbuffer.Line) {
		lineC <- l.Data
	})
	defer errBuffer.Close()

	// Prepare the command context, and start the process.
	cmd := exec.CommandContext(ctx, path, args...)
	// Tee std{out,err} into the linebuffers above and the process' std{out,err}, to
	// allow easier debugging.
	cmd.Stdout = io.MultiWriter(os.Stdout, outBuffer)
	cmd.Stderr = io.MultiWriter(os.Stderr, errBuffer)
	if err := cmd.Start(); err != nil {
		return false, fmt.Errorf("couldn't start the process: %w", err)
	}

	// Try matching against expectedOutput and return the result.
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		case line := <-lineC:
			if pf(line) {
				cmd.Process.Kill()
				cmd.Wait()
				return true, nil
			}
		}
	}
}

// TerminateIfFound creates RunCommand predicates that instantly terminate
// program execution in the event the given string is found in any line
// produced. RunCommand will return true, if the string searched for was found,
// and false otherwise.
func TerminateIfFound(needle string) func(string) bool {
	return func(haystack string) bool {
		return strings.Contains(haystack, needle)
	}
}
