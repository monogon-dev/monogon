package sinbin

import (
	"testing"
	"time"
)

// TestSinbinBasics performs some basic tests on the Sinbin.
func TestSinbinBasics(t *testing.T) {
	var s Sinbin[string]

	if s.Penalized("foo") {
		t.Errorf("'foo' should not be penalized as it hasn't yet been added")
	}
	s.Add("foo", time.Now().Add(-1*time.Second))
	if s.Penalized("foo") {
		t.Errorf("'foo' should not be penalized as it has been added with an expired time")
	}
	s.Add("bar", time.Now().Add(time.Hour))
	if !s.Penalized("bar") {
		t.Errorf("'bar' should be penalized as it has been added with an hour long penalty")
	}

	// Force sweep.
	s.lastSweep = time.Now().Add(-1 * time.Hour)
	s.sweepUnlocked()

	if len(s.bench) != 1 {
		t.Errorf("there should only be one element on the bench")
	}
	if _, ok := s.bench["bar"]; !ok {
		t.Errorf("the bench should contain 'bar'")
	}
}
