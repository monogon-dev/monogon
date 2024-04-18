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
	defer res.Close()
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

func readBacklog(tree *LogTree, t *testing.T, dn DN, backlog int, recursive bool) []string {
	t.Helper()
	opts := []LogReadOption{
		WithBacklog(backlog),
	}
	if recursive {
		opts = append(opts, WithChildren())
	}
	res, err := tree.Read(dn, opts...)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	defer res.Close()

	var lines []string
	for _, e := range res.Backlog {
		lines = append(lines, e.Leveled.Messages()...)
	}
	return lines
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

func TestBacklogAll(t *testing.T) {
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

func TestBacklogExact(t *testing.T) {
	tree := New()
	tree.MustLeveledFor("main").Info("hello, main!")
	tree.MustLeveledFor("main.foo").Info("hello, main.foo!")
	tree.MustLeveledFor("main.bar").Info("hello, main.bar!")
	tree.MustLeveledFor("main.bar.chatty").Info("hey there how are you")
	tree.MustLeveledFor("main.bar.quiet").Info("fine how are you")
	tree.MustLeveledFor("main.bar.chatty").Info("i've been alright myself")
	tree.MustLeveledFor("main.bar.chatty").Info("but to tell you honestly...")
	tree.MustLeveledFor("main.bar.chatty").Info("i feel like i'm stuck?")
	tree.MustLeveledFor("main.bar.quiet").Info("mhm")
	tree.MustLeveledFor("main.bar.chatty").Info("like you know what i'm saying, stuck in like")
	tree.MustLeveledFor("main.bar.chatty").Info("like a go test?")
	tree.MustLeveledFor("main.bar.quiet").Info("yeah totally")
	tree.MustLeveledFor("main.bar.chatty").Info("it's hard to put my finger on it")
	tree.MustLeveledFor("main.bar.chatty").Info("anyway, how's the wife doing?")

	check := func(a []string, b ...string) {
		t.Helper()
		if len(a) != len(b) {
			t.Errorf("Legth mismatch: wanted %d, got %d", len(b), len(a))
		}
		count := len(a)
		if len(b) < count {
			count = len(b)
		}
		for i := 0; i < count; i++ {
			if want, got := b[i], a[i]; want != got {
				t.Errorf("Message %d: wanted %q, got %q", i, want, got)
			}
		}
	}

	check(readBacklog(tree, t, "main", 3, true), "yeah totally", "it's hard to put my finger on it", "anyway, how's the wife doing?")
	check(readBacklog(tree, t, "main.foo", 3, false), "hello, main.foo!")
	check(readBacklog(tree, t, "main.bar.quiet", 2, true), "mhm", "yeah totally")
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

func TestLogEntry_ConciseString(t *testing.T) {
	trim := func(s string) string {
		return strings.Trim(s, "\n")
	}
	for i, te := range []struct {
		entry    *LogEntry
		maxWidth int
		want     string
	}{
		{
			&LogEntry{
				Leveled: &LeveledPayload{
					messages: []string{"Hello there!"},
					severity: WARNING,
				},
				DN: "root.role.kubernetes.run.kubernetes.apiserver",
			},
			120,
			"       k8s apiserver W Hello there!",
		},
		{
			&LogEntry{
				Leveled: &LeveledPayload{
					messages: []string{"Hello there!", "I am multiline."},
					severity: WARNING,
				},
				DN: "root.role.kubernetes.run.kubernetes.apiserver",
			},
			120,
			trim(`
       k8s apiserver W Hello there!
                     | I am multiline.
`),
		},
		{
			&LogEntry{
				Leveled: &LeveledPayload{
					messages: []string{"Hello there! I am a very long string, and I will get wrapped to 120 columns because that's just how life is for long strings."},
					severity: WARNING,
				},
				DN: "root.role.kubernetes.run.kubernetes.apiserver",
			},
			120,
			trim(`
       k8s apiserver W Hello there! I am a very long string, and I will get wrapped to 120 columns because that's just
                     | how life is for long strings.
`),
		},
		{
			&LogEntry{
				Leveled: &LeveledPayload{
					messages: []string{"Hello there!"},
					severity: WARNING,
				},
				DN: "root.role.kubernetes.run.kubernetes.apiserver",
			},
			60,
			trim(`
   k8s apiserver W Hello there!
`),
		},
		{
			&LogEntry{
				Leveled: &LeveledPayload{
					messages: []string{"Hello there!"},
					severity: WARNING,
				},
				DN: "root.role.kubernetes.run.kubernetes.apiserver",
			},
			40,
			"W Hello there!",
		},
	} {
		got := te.entry.ConciseString(MetropolisShortenDict, te.maxWidth)
		for _, line := range strings.Split(got, "\n") {
			if want, got := te.maxWidth, len(line); got > want {
				t.Errorf("Case %d, line %q too long (%d bytes, wanted at most %d)", i, line, got, want)
			}
		}
		if te.want != got {
			t.Errorf("Case %d, message diff", i)
			t.Logf("Wanted:\n%s", te.want)
			t.Logf("Got:\n%s", got)
		}
	}
}
