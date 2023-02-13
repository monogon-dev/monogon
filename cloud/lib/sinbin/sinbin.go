// Package sinbin implements a sinbin for naughty processed elements that we wish
// to time out for a while. This is kept in memory, and effectively implements a
// simplified version of the Circuit Breaker pattern.
//
// “sin bin”, noun, informal: (in sport) a box or bench to which offending
// players can be sent for a period as a penalty during a game, especially in ice
// hockey.
package sinbin

import (
	"sync"
	"time"
)

type entry struct {
	until time.Time
}

// A Sinbin contains a set of entries T which are added with a deadline, and will
// be automatically collected when that deadline expires.
//
// The zero value of a Sinbin is ready to use, and can be called from multiple
// goroutines.
type Sinbin[T comparable] struct {
	mu    sync.RWMutex
	bench map[T]*entry

	lastSweep time.Time
}

func (s *Sinbin[T]) initializeUnlocked() {
	if s.bench == nil {
		s.bench = make(map[T]*entry)
	}
}

func (s *Sinbin[T]) sweepUnlocked() {
	if s.lastSweep.Add(time.Minute).After(time.Now()) {
		return
	}
	now := time.Now()
	for k, e := range s.bench {
		if now.After(e.until) {
			delete(s.bench, k)
		}
	}
	s.lastSweep = now
}

// Add an element 't' to a Sinbin with a given deadline. From now until that
// deadline Penalized(t) will return true.
func (s *Sinbin[T]) Add(t T, until time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.initializeUnlocked()
	s.sweepUnlocked()

	existing, ok := s.bench[t]
	if ok {
		if until.After(existing.until) {
			existing.until = until
		}
		return
	}
	s.bench[t] = &entry{
		until: until,
	}
}

// Penalized returns whether the given element is currently sitting on the
// time-out bench after having been Added previously.
func (s *Sinbin[T]) Penalized(t T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.bench == nil {
		return false
	}

	existing, ok := s.bench[t]
	if !ok {
		return false
	}
	if time.Now().After(existing.until) {
		delete(s.bench, t)
		return false
	}
	return true
}
