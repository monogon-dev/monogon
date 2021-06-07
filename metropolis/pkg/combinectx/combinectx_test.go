package combinectx

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	a, aC := context.WithCancel(context.Background())
	b, bC := context.WithCancel(context.Background())

	c := Combine(a, b)
	if want, got := error(nil), c.Err(); want != got {
		t.Fatalf("Newly combined context should return %v, got %v", want, got)
	}
	if _, ok := c.Deadline(); ok {
		t.Errorf("Newly combined context should have no deadline")
	}

	// Cancel A.
	aC()
	// Cancels are not synchronous - wait for it to propagate...
	<-c.Done()
	// ...then cancel B (no-op).
	bC()

	if c.Err() == nil {
		t.Fatalf("After cancel, ctx.Err() should be non-nil")
	}
	if !errors.Is(c.Err(), a.Err()) {
		t.Errorf("After cancel, ctx.Err() should be a.Err()")
	}
	if !errors.Is(c.Err(), c.Err()) {
		t.Errorf("After cancel, ctx.Err() should be ctx.Err()")
	}
	if !errors.Is(c.Err(), context.Canceled) {
		t.Errorf("After cancel, ctx.Err() should be context.Canceled")
	}
	if !errors.Is(c.Err(), &Error{}) {
		t.Errorf("After cancel, ctx.Err() should be a Error pointer")
	}
	cerr := &Error{}
	if !errors.As(c.Err(), &cerr) {
		t.Fatalf("After cancel, ctx.Err() should be usable as *Error")
	}
	if !cerr.First() {
		t.Errorf("ctx.Err().First() should be true")
	}
	if cerr.Second() {
		t.Errorf("ctx.Err().Second() should be false")
	}
	if want, got := a.Err(), cerr.Unwrap(); want != got {
		t.Errorf("ctx.Err().Unwrap() should be %v, got %v", want, got)
	}
}

func TestDeadline(t *testing.T) {
	now := time.Now()
	aD := now.Add(100 * time.Millisecond)
	bD := now.Add(10 * time.Millisecond)

	a, aC := context.WithDeadline(context.Background(), aD)
	b, bC := context.WithDeadline(context.Background(), bD)

	defer aC()
	defer bC()

	c := Combine(a, b)
	if want, got := error(nil), c.Err(); want != got {
		t.Fatalf("Newly combined context should return %v, got %v", want, got)
	}
	if d, ok := c.Deadline(); !ok || !d.Equal(bD) {
		t.Errorf("Newly combined context should have deadline %v, got %v", bD, d)
	}

	<-c.Done()

	if c.Err() == nil {
		t.Fatalf("After deadline, ctx.Err() should be non-nil")
	}
	if !errors.Is(c.Err(), b.Err()) {
		t.Errorf("After deadline, ctx.Err() should be b.Err()")
	}
	if !errors.Is(c.Err(), context.DeadlineExceeded) {
		t.Errorf("After cancel, ctx.Err() should be context.DeadlineExceeded")
	}
	if !errors.Is(c.Err(), &Error{}) {
		t.Errorf("After cancel, ctx.Err() should be a Error pointer")
	}
	cerr := &Error{}
	if !errors.As(c.Err(), &cerr) {
		t.Fatalf("After cancel, ctx.Err() should be usable as *Error")
	}
	if cerr.First() {
		t.Errorf("ctx.Err().First() should be false")
	}
	if !cerr.Second() {
		t.Errorf("ctx.Err().Second() should be true")
	}
	if want, got := b.Err(), cerr.Unwrap(); want != got {
		t.Errorf("ctx.Err().Unwrap() should be %v, got %v", want, got)
	}
}

