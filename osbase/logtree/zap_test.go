// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package logtree

import (
	"testing"

	"go.uber.org/zap"

	"source.monogon.dev/go/logging"
)

func TestZapify(t *testing.T) {
	lt := New()

	z := Zapify(lt.MustLeveledFor("zap"), zap.InfoLevel)
	z.Info("foo", zap.String("strp", "strv"), zap.Int("intp", 42))
	z.Warn("foo!", zap.String("strp", "strv"), zap.Int("intp", 1337))
	z.Error("foo!!")

	res, err := lt.Read("zap", WithBacklog(BacklogAllAvailable))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	defer res.Close()

	if want, got := 3, len(res.Backlog); want != got {
		t.Errorf("Wanted %d entries, got %d", want, got)
	} else {
		for i, te := range []struct {
			msg string
			sev logging.Severity
		}{
			{`foo {"intp":42,"strp":"strv"}`, logging.INFO},
			{`foo! {"intp":1337,"strp":"strv"}`, logging.WARNING},
			{`foo!!`, logging.ERROR},
		} {
			if want, got := te.msg, res.Backlog[i].Leveled.messages[0]; want != got {
				t.Errorf("Line %d: wanted message %q, got %q", i, want, got)
			}
			if want, got := te.sev, res.Backlog[i].Leveled.severity; want != got {
				t.Errorf("Line %d: wanted level %s, got %s", i, want, got)
			}
		}
	}
}
