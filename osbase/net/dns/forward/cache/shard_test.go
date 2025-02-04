// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package cache

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"testing"
)

// newShard returns a new shard with size.
func newShard[T any](size int) *shard[T] { return &shard[T]{items: make(map[uint64]T), size: size} }

func TestShardAddAndGet(t *testing.T) {
	s := newShard[int](1)
	s.Put(1, 1)

	if _, found := s.Get(1); !found {
		t.Fatal("Failed to find inserted record")
	}

	s.Put(2, 1)
	if _, found := s.Get(1); found {
		t.Fatal("Failed to evict record")
	}
	if _, found := s.Get(2); !found {
		t.Fatal("Failed to find inserted record")
	}
}

func TestGetOrPut(t *testing.T) {
	s := newShard[int](1)
	el, exists := s.GetOrPut(1, 2)
	if exists {
		t.Fatalf("Element should not have existed")
	}
	if el != 2 {
		t.Fatalf("Expected element %d, got %d ", 2, el)
	}

	el, exists = s.GetOrPut(1, 3)
	if !exists {
		t.Fatalf("Element should have existed")
	}
	if el != 2 {
		t.Fatalf("Expected element %d, got %d ", 2, el)
	}
}

func TestShardRemove(t *testing.T) {
	s := newShard[int](2)
	s.Put(1, 1)
	s.Put(2, 2)

	s.Remove(1)

	if _, found := s.Get(1); found {
		t.Fatal("Failed to remove record")
	}
	if _, found := s.Get(2); !found {
		t.Fatal("Failed to find inserted record")
	}
}

func TestAddEvict(t *testing.T) {
	const size = 1024
	s := newShard[int](size)

	for i := uint64(0); i < size; i++ {
		s.Put(i, 1)
	}
	for i := uint64(0); i < size; i++ {
		s.Put(i, 1)
		if len(s.items) != size {
			t.Fatal("A item was unnecessarily evicted from the cache")
		}
	}
}

func TestShardLen(t *testing.T) {
	s := newShard[int](4)

	s.Put(1, 1)
	if l := len(s.items); l != 1 {
		t.Fatalf("Shard size should %d, got %d", 1, l)
	}

	s.Put(1, 1)
	if l := len(s.items); l != 1 {
		t.Fatalf("Shard size should %d, got %d", 1, l)
	}

	s.Put(2, 2)
	if l := len(s.items); l != 2 {
		t.Fatalf("Shard size should %d, got %d", 2, l)
	}
}

func TestShardEvict(t *testing.T) {
	s := newShard[int](1)
	s.Put(1, 1)
	s.Put(2, 2)
	// 1 should be gone

	if _, found := s.Get(1); found {
		t.Fatal("Found item that should have been evicted")
	}
}

func TestShardLenEvict(t *testing.T) {
	s := newShard[int](4)
	s.Put(1, 1)
	s.Put(2, 1)
	s.Put(3, 1)
	s.Put(4, 1)

	if l := len(s.items); l != 4 {
		t.Fatalf("Shard size should %d, got %d", 4, l)
	}

	// This should evict one element
	s.Put(5, 1)
	if l := len(s.items); l != 4 {
		t.Fatalf("Shard size should %d, got %d", 4, l)
	}

	// Make sure we don't accidentally evict an element when
	// we the key is already stored.
	for i := 0; i < 4; i++ {
		s.Put(5, 1)
		if l := len(s.items); l != 4 {
			t.Fatalf("Shard size should %d, got %d", 4, l)
		}
	}
}

func BenchmarkShard(b *testing.B) {
	s := newShard[int](shardSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := uint64(i) % shardSize * 2
		s.Put(k, 1)
		s.Get(k)
	}
}

func BenchmarkShardParallel(b *testing.B) {
	s := newShard[int](shardSize)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := uint64(0); pb.Next(); i++ {
			k := i % shardSize * 2
			s.Put(k, 1)
			s.Get(k)
		}
	})
}
