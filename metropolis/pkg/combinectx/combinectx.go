// package combinectx implements context.Contexts that 'combine' two other
// 'parent' contexts. These can be used to deal with cases where you want to
// cancel a method call whenever any of two pre-existing contexts expires first.
//
// For example, if you want to tie a method call to some incoming request
// context and an active leader lease, then this library is for you.
package combinectx

import (
	"context"
	"sync"
	"time"
)

// Combine 'joins' two existing 'parent' contexts into a single context. This
// context will be Done() whenever any of the parent context is Done().
// Combining contexts spawns a goroutine that will be cleaned when any of the
// parent contexts is Done().
func Combine(a, b context.Context) context.Context {
	c := &Combined{
		a:     a,
		b:     b,
		doneC: make(chan struct{}),
	}
	go c.run()
	return c
}

type Combined struct {
	// a is the first parent context.
	a context.Context
	// b is the second parent context.
	b context.Context

	// mu guards done.
	mu sync.Mutex
	// done is an Error if either parent context is Done(), or nil otherwise.
	done *Error
	// doneC is closed when either parent context is Done() and Error is set.
	doneC chan struct{}
}

// Error wraps errors returned by parent contexts.
type Error struct {
	// underlyingA points to an error returned by the first parent context if the
	// combined context was Done() as a result of the first parent context being
	// Done().
	underlyingA *error
	// underlyingB points to an error returned by the second parent context if the
	// combined context was Done() as a result of the second parent context being
	// Done().
	underlyingB *error
}

func (e *Error) Error() string {
	if e.underlyingA != nil {
		return (*e.underlyingA).Error()
	}
	if e.underlyingB != nil {
		return (*e.underlyingB).Error()
	}
	return ""
}

// First returns true if the Combined context's first parent was Done().
func (e *Error) First() bool {
	return e.underlyingA != nil
}

// Second returns true if the Combined context's second parent was Done().
func (e *Error) Second() bool {
	return e.underlyingB != nil
}

// Unwrap returns the underlying error of either parent context that is Done().
func (e *Error) Unwrap() error {
	if e.underlyingA != nil {
		return *e.underlyingA
	}
	if e.underlyingB != nil {
		return *e.underlyingB
	}
	return nil
}

// Is allows errors.Is to be true against any *Error.
func (e *Error) Is(target error) bool {
	if _, ok := target.(*Error); ok {
		return true
	}
	return false
}

// As allows errors.As to be true against any *Error.
func (e *Error) As(target interface{}) bool {
	if v, ok := target.(**Error); ok {
		*v = e
		return true
	}
	return false
}

// run is the main logic that ties together the two parent contexts. It exits
// when either parent context is canceled.
func (c *Combined) run() {
	mark := func(first bool, err error) {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.done = &Error{}
		if first {
			c.done.underlyingA = &err
		} else {
			c.done.underlyingB = &err
		}
		close(c.doneC)
	}
	select {
	case <-c.a.Done():
		mark(true, c.a.Err())
	case <-c.b.Done():
		mark(false, c.b.Err())
	}
}

// Deadline returns the earlier Deadline from the two parent contexts, if any.
func (c *Combined) Deadline() (deadline time.Time, ok bool) {
	d1, ok1 := c.a.Deadline()
	d2, ok2 := c.b.Deadline()

	if ok1 && !ok2 {
		return d1, true
	}
	if ok2 && !ok1 {
		return d2, true
	}
	if !ok1 && !ok2 {
		return time.Time{}, false
	}

	if d1.Before(d2) {
		return d1, true
	}
	return d2, true
}

func (c *Combined) Done() <-chan struct{} {
	return c.doneC
}

// Err returns nil if neither parent context is Done() yet, or an error otherwise.
// The returned errors will have the following properties:
//   1) errors.Is(err, Error{}) will always return true.
//   2) errors.Is(err, ctx.Err()) will return true if the combined context was
//      canceled with the same error as ctx.Err().
//      However, this does NOT mean that the combined context was Done() because
//      of the ctx being Done() - to ensure this is the case, use errors.As() to
//      retrieve an Error and its First()/Second() methods.
//   3) errors.Is(err, context.{Canceled,DeadlineExceeded}) will return true if
//      the combined context is Canceled or DeadlineExceeded.
//   4) errors.Is will return false otherwise.
//   5) errors.As(err, &&Error{})) will always return true. The Error object can
//      then be used to check the cause of the combined context's error.
func (c *Combined) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.done == nil {
		return nil
	}
	return c.done
}

// Value returns the value located under the given key by checking the first and
// second parent context in order.
func (c *Combined) Value(key interface{}) interface{} {
	if v := c.a.Value(key); v != nil {
		return v
	}
	if v := c.b.Value(key); v != nil {
		return v
	}
	return nil
}
