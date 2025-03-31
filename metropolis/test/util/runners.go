// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// This file implements test helper functions that augment the way any given
// test is run.
package util

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

// TestEventual creates a new subtest looping the given function until it either
// doesn't return an error anymore, the timeout is exceeded or PermanentError is
// returned. The last returned non-context-related error is being used as the
// test error.
func TestEventual(t *testing.T, name string, ctx context.Context, timeout time.Duration, f func(context.Context) error) bool {
	start := time.Now()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	t.Helper()
	fmt.Printf("Test: %s: starting...\n", name)
	return t.Run(name, func(t *testing.T) {
		defer cancel()
		var lastErr = errors.New("test didn't run to completion at least once")
		for {
			err := f(ctx)
			if err == nil {
				fmt.Printf("Test: %s: okay after %.1f seconds\n", name, time.Since(start).Seconds())
				return
			}
			if errors.Is(err, ctx.Err()) {
				t.Fatal(lastErr)
			}
			if errors.Is(err, &PermanentError{}) {
				t.Fatal(err)
			}
			lastErr = err
			select {
			case <-ctx.Done():
				t.Fatal(lastErr)
			case <-time.After(1 * time.Second):
			}
		}
	})
}

// MustTestEventual is like TestEventual, but aborts the `t` test with Fatal if a
// timeout occurred or PermanentError was returned.
func MustTestEventual(t *testing.T, name string, ctx context.Context, timeout time.Duration, f func(context.Context) error) {
	if !TestEventual(t, name, ctx, timeout, f) {
		t.Fatalf("Test: %s: fatal failure", name)
	}
}

// PermanentError can be returned inside TestEventual to indicate that the test
// is 'stuck', that it will not make progress anymore and that it should be
// failed immediately.
type PermanentError struct {
	Err error
}

func (p *PermanentError) Error() string {
	return fmt.Sprintf("test permanently failed: %v", p.Err)
}

func (p *PermanentError) Unwrap() error {
	return p.Err
}

func (p *PermanentError) Is(o error) bool {
	var op *PermanentError
	if !errors.As(o, &op) {
		return false
	}
	if p.Err == nil || op.Err == nil {
		return true
	}
	return errors.Is(p.Err, op.Err)
}

// Permanent wraps the given error into a PermanentError, which will cause
// TestEventual to immediately fail the test it's returned within.
func Permanent(err error) error {
	return &PermanentError{
		Err: err,
	}
}
