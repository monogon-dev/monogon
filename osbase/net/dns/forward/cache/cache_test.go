// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package cache

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"testing"
)

func TestCacheAddAndGet(t *testing.T) {
	const N = shardSize * 4
	c := New[int](N)
	c.Put(1, 1)

	if _, found := c.Get(1); !found {
		t.Fatal("Failed to find inserted record")
	}
}

func TestCacheSharding(t *testing.T) {
	c := New[int](shardSize)
	for i := 0; i < shardSize*2; i++ {
		c.Put(uint64(i), 1)
	}
	for i := range c.shards {
		if len(c.shards[i].items) == 0 {
			t.Errorf("Failed to populate shard: %d", i)
		}
	}
}

func BenchmarkCache(b *testing.B) {
	b.ReportAllocs()

	c := New[int](4)
	for n := 0; n < b.N; n++ {
		c.Put(1, 1)
		c.Get(1)
	}
}
