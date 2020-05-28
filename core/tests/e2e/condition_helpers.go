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
	"time"

	apipb "git.monogon.dev/source/nexantic.git/core/generated/api"
)

func waitForCondition(ctx context.Context, client apipb.NodeDebugServiceClient, condition string) error {
	var lastErr = errors.New("No RPC for checking condition completed")
	for {
		res, err := client.GetCondition(ctx, &apipb.GetConditionRequest{Name: condition})
		if err != nil {
			if err == ctx.Err() {
				return err
			}
			lastErr = err
		}
		if err == nil && res.Ok {
			return nil
		}
		select {
		case <-time.After(1 * time.Second):
		case <-ctx.Done():
			return lastErr
		}
	}
}
