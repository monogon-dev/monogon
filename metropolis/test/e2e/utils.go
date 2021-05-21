// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	"context"
	"errors"
	"testing"
	"time"
)

// testEventual creates a new subtest looping the given function until it
// either doesn't return an error anymore or the timeout is exceeded. The last
// returned non-context-related error is being used as the test error.
func testEventual(t *testing.T, name string, ctx context.Context, timeout time.Duration, f func(context.Context) error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	t.Helper()
	t.Run(name, func(t *testing.T) {
		defer cancel()
		var lastErr = errors.New("test didn't run to completion at least once")
		t.Parallel()
		for {
			err := f(ctx)
			if err == nil {
				return
			}
			if err == ctx.Err() {
				t.Fatal(lastErr)
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
