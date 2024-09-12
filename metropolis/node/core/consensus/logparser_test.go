package consensus

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"source.monogon.dev/go/logging"
	"source.monogon.dev/osbase/logbuffer"
	"source.monogon.dev/osbase/logtree"
)

// TestParsing exercises the parseEtcdLogEntry function.
func TestParsing(t *testing.T) {
	timeParse := func(s string) time.Time {
		res, err := time.Parse(time.RFC3339, s)
		if err != nil {
			t.Fatalf("could not parse time: %v", err)
		}
		return res
	}
	for _, te := range []struct {
		// Name of subtest.
		name string
		// Data to be parsed.
		raw string
		// The expected parsed data. The parser does not attempt to set any
		// 'default' values in case any are missing, instead the logtree's
		// external leveled payload functionality does that.
		want *logtree.ExternalLeveledPayload
	}{
		{
			"Parse configuring peer listeners message",
			`{"level":"info","ts":"2021-07-06T17:18:24.368Z","caller":"embed/etcd.go:117","msg":"configuring peer listeners","listen-peer-urls":["https://[::]:7834"]}`,
			&logtree.ExternalLeveledPayload{
				Message:   `configuring peer listeners, listen-peer-urls: ["https://[::]:7834"]`,
				Timestamp: timeParse("2021-07-06T17:18:24.368Z"),
				Severity:  logging.INFO,
				File:      "etcd.go",
				Line:      117,
			},
		},
		{
			"Parse added member message",
			`{"level":"info","ts":"2021-07-06T17:21:49.462Z","caller":"membership/cluster.go:392","msg":"added member","cluster-id":"137c8e19524788c1","local-member-id":"9642132f5d0d99e2","added-peer-id":"9642132f5d0d99e2","added-peer-peer-urls":["https://metropolis-eb8d68cfb52711ad04c339abdeea74ed:7834"]}`,
			&logtree.ExternalLeveledPayload{
				Message:   `added member, added-peer-id: "9642132f5d0d99e2", added-peer-peer-urls: ["https://metropolis-eb8d68cfb52711ad04c339abdeea74ed:7834"], cluster-id: "137c8e19524788c1", local-member-id: "9642132f5d0d99e2"`,
				Timestamp: timeParse("2021-07-06T17:21:49.462Z"),
				Severity:  logging.INFO,
				File:      "cluster.go",
				Line:      392,
			},
		},
		{
			"Parse empty message",
			`{}`,
			&logtree.ExternalLeveledPayload{},
		},
		{
			"Parse invalid message",
			`PANIC`,
			&logtree.ExternalLeveledPayload{
				Message: "Log line unparseable: PANIC",
			},
		},
	} {
		te := te
		t.Run(te.name, func(t *testing.T) {
			t.Parallel()

			gotC := make(chan *logtree.ExternalLeveledPayload, 1)
			parseEtcdLogEntry(&logbuffer.Line{
				Data:           te.raw,
				OriginalLength: len(te.raw),
			}, func(d *logtree.ExternalLeveledPayload) {
				gotC <- d
			})

			got := <-gotC
			if diff := cmp.Diff(te.want, got); diff != "" {
				t.Fatalf("diff: %s", diff)
			}
		})
	}
}
