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
	"fmt"
	"strings"
	"testing"
	"time"
)

func expect(tree *LogTree, t *testing.T, dn DN, entries ...string) string {
	t.Helper()
	res, err := tree.Read(dn, WithChildren(), WithBacklog(BacklogAllAvailable))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want, got := len(entries), len(res.Backlog); want != got {
		t.Fatalf("wanted %v backlog entries, got %v", want, got)
	}
	got := make(map[string]bool)
	for _, entry := range res.Backlog {
		if entry.Leveled != nil {
			got[entry.Leveled.MessagesJoined()] = true
		}
		if entry.Raw != nil {
			got[entry.Raw.Data] = true
		}
	}
	for _, entry := range entries {
		if !got[entry] {
			return fmt.Sprintf("missing entry %q", entry)
		}
	}
	return ""
}

func TestMultiline(t *testing.T) {
	tree := New()
	// Two lines in a single message.
	tree.MustLeveledFor("main").Info("foo\nbar")
	// Two lines in a single message with a hanging newline that should get stripped.
	tree.MustLeveledFor("main").Info("one\ntwo\n")

	if res := expect(tree, t, "main", "foo\nbar", "one\ntwo"); res != "" {
		t.Errorf("retrieval at main failed: %s", res)
	}
}

func TestBacklog(t *testing.T) {
	tree := New()
	tree.MustLeveledFor("main").Info("hello, main!")
	tree.MustLeveledFor("main.foo").Info("hello, main.foo!")
	tree.MustLeveledFor("main.bar").Info("hello, main.bar!")
	tree.MustLeveledFor("aux").Info("hello, aux!")
	// No newline at the last entry - shouldn't get propagated to the backlog.
	fmt.Fprintf(tree.MustRawFor("aux.process"), "processing foo\nprocessing bar\nbaz")

	if res := expect(tree, t, "main", "hello, main!", "hello, main.foo!", "hello, main.bar!"); res != "" {
		t.Errorf("retrieval at main failed: %s", res)
	}
	if res := expect(tree, t, "", "hello, main!", "hello, main.foo!", "hello, main.bar!", "hello, aux!", "processing foo", "processing bar"); res != "" {
		t.Errorf("retrieval at root failed: %s", res)
	}
	if res := expect(tree, t, "aux", "hello, aux!", "processing foo", "processing bar"); res != "" {
		t.Errorf("retrieval at aux failed: %s", res)
	}
}

func TestStream(t *testing.T) {
	tree := New()
	tree.MustLeveledFor("main").Info("hello, backlog")
	fmt.Fprintf(tree.MustRawFor("main.process"), "hello, raw backlog\n")

	res, err := tree.Read("", WithBacklog(BacklogAllAvailable), WithChildren(), WithStream())
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	defer res.Close()
	if want, got := 2, len(res.Backlog); want != got {
		t.Errorf("wanted %d backlog item, got %d", want, got)
	}

	tree.MustLeveledFor("main").Info("hello, stream")
	fmt.Fprintf(tree.MustRawFor("main.raw"), "hello, raw stream\n")

	entries := make(map[string]bool)
	timeout := time.After(time.Second * 1)
	for {
		done := false
		select {
		case <-timeout:
			done = true
		case p := <-res.Stream:
			if p.Leveled != nil {
				entries[p.Leveled.MessagesJoined()] = true
			}
			if p.Raw != nil {
				entries[p.Raw.Data] = true
			}
		}
		if done {
			break
		}
	}
	if entry := "hello, stream"; !entries[entry] {
		t.Errorf("Missing entry %q", entry)
	}
	if entry := "hello, raw stream"; !entries[entry] {
		t.Errorf("Missing entry %q", entry)
	}
}

func TestVerbose(t *testing.T) {
	tree := New()

	tree.MustLeveledFor("main").V(10).Info("this shouldn't get logged")

	reader, err := tree.Read("", WithBacklog(BacklogAllAvailable), WithChildren())
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want, got := 0, len(reader.Backlog); want != got {
		t.Fatalf("expected nothing to be logged, got %+v", reader.Backlog)
	}

	tree.SetVerbosity("main", 10)
	tree.MustLeveledFor("main").V(10).Info("this should get logged")

	reader, err = tree.Read("", WithBacklog(BacklogAllAvailable), WithChildren())
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want, got := 1, len(reader.Backlog); want != got {
		t.Fatalf("expected %d entries to get logged, got %d", want, got)
	}
}

func TestMetadata(t *testing.T) {
	tree := New()
	tree.MustLeveledFor("main").Error("i am an error")
	tree.MustLeveledFor("main").Warning("i am a warning")
	tree.MustLeveledFor("main").Info("i am informative")
	tree.MustLeveledFor("main").V(0).Info("i am a zero-level debug")

	reader, err := tree.Read("", WithChildren(), WithBacklog(BacklogAllAvailable))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want, got := 4, len(reader.Backlog); want != got {
		t.Fatalf("expected %d entries, got %d", want, got)
	}

	for _, te := range []struct {
		ix       int
		severity Severity
		message  string
	}{
		{0, ERROR, "i am an error"},
		{1, WARNING, "i am a warning"},
		{2, INFO, "i am informative"},
		{3, INFO, "i am a zero-level debug"},
	} {
		p := reader.Backlog[te.ix]
		if want, got := te.severity, p.Leveled.Severity(); want != got {
			t.Errorf("wanted element %d to have severity %s, got %s", te.ix, want, got)
		}
		if want, got := te.message, p.Leveled.MessagesJoined(); want != got {
			t.Errorf("wanted element %d to have message %q, got %q", te.ix, want, got)
		}
		if want, got := "logtree_test.go", strings.Split(p.Leveled.Location(), ":")[0]; want != got {
			t.Errorf("wanted element %d to have file %q, got %q", te.ix, want, got)
		}
	}
}

func TestSeverity(t *testing.T) {
	tree := New()
	tree.MustLeveledFor("main").Error("i am an error")
	tree.MustLeveledFor("main").Warning("i am a warning")
	tree.MustLeveledFor("main").Info("i am informative")
	tree.MustLeveledFor("main").V(0).Info("i am a zero-level debug")

	reader, err := tree.Read("main", WithBacklog(BacklogAllAvailable), LeveledWithMinimumSeverity(WARNING))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want, got := 2, len(reader.Backlog); want != got {
		t.Fatalf("wanted %d entries, got %d", want, got)
	}
	if want, got := "i am an error", reader.Backlog[0].Leveled.MessagesJoined(); want != got {
		t.Fatalf("wanted entry %q, got %q", want, got)
	}
	if want, got := "i am a warning", reader.Backlog[1].Leveled.MessagesJoined(); want != got {
		t.Fatalf("wanted entry %q, got %q", want, got)
	}
}

func TestAddedStackDepth(t *testing.T) {
	tree := New()
	helper := func(msg string) {
		tree.MustLeveledFor("main").WithAddedStackDepth(1).Infof("oh no: %s", msg)
	}

	// The next three lines are tested to be next to each other.
	helper("it failed")
	tree.MustLeveledFor("main").Infof("something else")

	reader, err := tree.Read("main", WithBacklog(BacklogAllAvailable))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want, got := 2, len(reader.Backlog); want != got {
		t.Fatalf("wanted %d entries, got %d", want, got)
	}
	if want, got := "oh no: it failed", reader.Backlog[0].Leveled.MessagesJoined(); want != got {
		t.Errorf("wanted entry %q, got %q", want, got)
	}
	if want, got := "something else", reader.Backlog[1].Leveled.MessagesJoined(); want != got {
		t.Errorf("wanted entry %q, got %q", want, got)
	}
	if first, second := reader.Backlog[0].Leveled.line, reader.Backlog[1].Leveled.line; first+1 != second {
		t.Errorf("first entry at %d, second at %d, wanted one after the other", first, second)
	}
}
