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

package supervisor

import (
	"context"
	"testing"
)

// waitSettle waits until the supervisor reaches a 'settled' state - ie., one
// where no actions have been performed for a number of GC cycles.
// This is used in tests only.
func (s *supervisor) waitSettle(ctx context.Context) error {
	waiter := make(chan struct{})
	s.pReq <- &processorRequest{
		waitSettled: &processorRequestWaitSettled{
			waiter: waiter,
		},
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-waiter:
		return nil
	}
}

// waitSettleError wraps waitSettle to fail a test if an error occurs, eg. the
// context is canceled.
func (s *supervisor) waitSettleError(ctx context.Context, t *testing.T) {
	err := s.waitSettle(ctx)
	if err != nil {
		t.Fatalf("waitSettle: %v", err)
	}
}
