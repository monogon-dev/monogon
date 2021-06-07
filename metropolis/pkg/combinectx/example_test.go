package combinectx_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"source.monogon.dev/metropolis/pkg/combinectx"
)

// ExampleCombine shows how to combine two contexts for use with a contextful
// method.
func ExampleCombine() {
	// Let's say you're trying to combine two different contexts: the first one is
	// some long-term local worker context, while the second is a context from some
	// incoming request.
	ctxA, cancelA := context.WithCancel(context.Background())
	ctxB, cancelB := context.WithTimeout(context.Background(), time.Millisecond * 100)
	defer cancelA()
	defer cancelB()

	// doIO is some contextful, black box IO-heavy function. You want it to return
	// early when either the long-term context or the short-term request context
	// are Done().
	doIO := func(ctx context.Context) (string, error) {
		t := time.NewTicker(time.Second)
		defer t.Stop()

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-t.C:
			return "successfully reticulated splines", nil
		}
	}

	// Combine the two given contexts into one...
	ctx := combinectx.Combine(ctxA, ctxB)
	// ... and call doIO with it.
	v, err := doIO(ctx)
	if err == nil {
		fmt.Printf("doIO success: %v\n", v)
		return
	}

	fmt.Printf("doIO failed: %v\n", err)

	// The returned error will always be equal to the combined context's Err() call
	// if the error is due to the combined context being Done().
	if err == ctx.Err() {
		fmt.Printf("doIO err == ctx.Err()\n")
	}

	// The returned error will pass any errors.Is(err, context....) checks. This
	// ensures compatibility with blackbox code that performs special actions on
	// the given context.
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("doIO err is context.DeadlineExceeded\n")
	}

	// The returned error can be inspected by converting it to a *Error and calling
	// .First()/.Second()/.Unwrap(). This lets the caller figure out which of the
	// parent contexts caused the combined context to expires.
	cerr := &combinectx.Error{}
	if errors.As(err, &cerr) {
		fmt.Printf("doIO err is *combinectx.Error\n")
		fmt.Printf("doIO first failed: %v, second failed: %v\n", cerr.First(), cerr.Second())
	}

	// Output:
	// doIO failed: context deadline exceeded
	// doIO err == ctx.Err()
	// doIO err is context.DeadlineExceeded
	// doIO err is *combinectx.Error
	// doIO first failed: false, second failed: true
}
