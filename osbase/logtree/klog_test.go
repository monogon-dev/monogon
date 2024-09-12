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

package logtree

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"source.monogon.dev/go/logging"
)

func TestParse(t *testing.T) {
	// Injected 'now'. Used to make these tests reproducible and to allow for
	// testing the log-across-year edgecase.
	// Fri 12 Mar 2021 03:46:26 PM UTC
	now := time.Unix(1615563986, 123456789)
	// Sat 01 Jan 2000 12:00:01 AM UTC
	nowNewYear := time.Unix(946684801, 0)

	for i, te := range []struct {
		now  time.Time
		line string
		want *LeveledPayload
	}{
		// 0: Simple case: everything should parse correctly.
		{now, "E0312 14:20:04.240540    204 shared_informer.go:247] Caches are synced for attach detach", &LeveledPayload{
			messages:  []string{"Caches are synced for attach detach"},
			timestamp: time.Date(2021, 03, 12, 14, 20, 4, 240540000, time.UTC),
			severity:  logging.ERROR,
			file:      "shared_informer.go",
			line:      247,
		}},
		// 1: Mumbling line, should fail.
		{now, "Application starting up...", nil},
		// 2: Empty line, should fail.
		{now, "", nil},
		// 3: Line from the future, should fail.
		{now, "I1224 14:20:04.240540    204 john_titor.go:247] I'm sorry, what day is it today? Uuuh, and what year?", nil},
		// 4: Log-across-year edge case. The log was emitted right before a year
		//    rollover, and parsed right after it. It should be attributed to the
		//    previous year.
		{nowNewYear, "I1231 23:59:43.123456    123 fry.go:123] Here's to another lousy millenium!", &LeveledPayload{
			messages:  []string{"Here's to another lousy millenium!"},
			timestamp: time.Date(1999, 12, 31, 23, 59, 43, 123456000, time.UTC),
			severity:  logging.INFO,
			file:      "fry.go",
			line:      123,
		}},
		// 5: Invalid severity, should fail.
		{now, "D0312 14:20:04.240540    204 shared_informer.go:247] Caches are synced for attach detach", nil},
		// 6: Invalid time, should fail.
		{now, "D0312 25:20:04.240540    204 shared_informer.go:247] Caches are synced for attach detach", nil},
		// 7: Simple case without sub-second timing: everything should parse correctly
		{now, "E0312 14:20:04 204 shared_informer.go:247] Caches are synced for attach detach", &LeveledPayload{
			messages:  []string{"Caches are synced for attach detach"},
			timestamp: time.Date(2021, 03, 12, 14, 20, 4, 0, time.UTC),
			severity:  logging.ERROR,
			file:      "shared_informer.go",
			line:      247,
		}},
	} {
		got := parse(te.now, te.line)
		if diff := cmp.Diff(te.want, got, cmp.AllowUnexported(LeveledPayload{})); diff != "" {
			t.Errorf("%d: mismatch (-want +got):\n%s", i, diff)
		}
	}
}
