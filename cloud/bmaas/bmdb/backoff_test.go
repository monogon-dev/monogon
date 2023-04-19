package bmdb

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// TestBackoffMath exercises the rules of Backoff.
func TestBackoffMath(t *testing.T) {
	for _, te := range []struct {
		name     string
		b        *Backoff
		existing *existingBackoff
		wantSecs []int64
	}{
		{"NoBackoffSet", nil, nil, []int64{1, 1, 1}},
		{"EmptyBackoff", &Backoff{}, nil, []int64{1, 1, 1}},
		{"SimpleBackoff", &Backoff{Initial: time.Minute}, nil, []int64{60, 60, 60}},
		{"ExponentialWithMax",
			&Backoff{Initial: time.Minute, Exponent: 1.1, Maximum: time.Minute * 2},
			nil,
			[]int64{60, 66, 73, 81, 90, 99, 109, 120, 120},
		},

		{"SimpleOverridePrevious",
			&Backoff{Initial: time.Minute},
			&existingBackoff{lastInterval: time.Second * 2},
			[]int64{60, 60, 60},
		},
		{"ExponentialOverridePrevious",
			&Backoff{Initial: time.Minute, Exponent: 2.0, Maximum: time.Minute * 2},
			&existingBackoff{lastInterval: time.Second * 2},
			[]int64{4, 8, 16, 32, 64, 120, 120},
		},

		{"ContinueExisting", nil, &existingBackoff{lastInterval: time.Minute}, []int64{60, 60, 60}},
		{"ContinueExistingInvalid1", nil, &existingBackoff{lastInterval: 0}, []int64{1, 1, 1}},
		{"ContinueExistingInvalid2", nil, &existingBackoff{lastInterval: time.Millisecond}, []int64{1, 1, 1}},

		{"InvalidBackoff1", &Backoff{Exponent: 0.2}, nil, []int64{1, 1, 1}},
		{"InvalidBackoff2", &Backoff{Maximum: time.Millisecond, Initial: time.Millisecond}, nil, []int64{1, 1, 1}},
	} {
		t.Run(te.name, func(t *testing.T) {
			existing := te.existing

			gotSecs := make([]int64, len(te.wantSecs))
			for j := 0; j < len(te.wantSecs); j++ {
				gotSecs[j] = te.b.next(existing)
				existing = &existingBackoff{
					lastInterval: time.Duration(gotSecs[j]) * time.Second,
				}
			}

			if diff := cmp.Diff(te.wantSecs, gotSecs); diff != "" {
				t.Errorf("Difference: %s", diff)
			}
		})
	}
}
