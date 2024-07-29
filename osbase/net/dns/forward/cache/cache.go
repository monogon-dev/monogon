// Package cache implements a cache. The cache hold 256 shards, each shard
// holds a cache: a map with a mutex. There is no fancy expunge algorithm, it
// just randomly evicts elements when it gets full.
package cache

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"sync"

	"golang.org/x/sys/cpu"
)

const shardSize = 256

// Cache is cache.
type Cache[T any] struct {
	shards [shardSize]shard[T]
}

// shard is a cache with random eviction.
type shard[T any] struct {
	items map[uint64]T
	size  int

	sync.RWMutex

	_ cpu.CacheLinePad
}

// New returns a new cache.
func New[T any](size int) *Cache[T] {
	ssize := size / shardSize
	if ssize < 4 {
		ssize = 4
	}

	c := &Cache[T]{}

	// Initialize all the shards
	for i := 0; i < shardSize; i++ {
		c.shards[i] = shard[T]{items: make(map[uint64]T), size: ssize}
	}
	return c
}

// Get returns the element under key, and whether the element exists.
func (c *Cache[T]) Get(key uint64) (el T, exists bool) {
	shard := key % shardSize
	return c.shards[shard].Get(key)
}

// Put adds a new element to the cache. If the element already exists,
// it is overwritten.
func (c *Cache[T]) Put(key uint64, el T) {
	shard := key % shardSize
	c.shards[shard].Put(key, el)
}

// GetOrPut returns the element under key if it exists,
// or else stores newEl there. This operation is atomic.
func (c *Cache[T]) GetOrPut(key uint64, newEl T) (el T, exists bool) {
	shard := key % shardSize
	return c.shards[shard].GetOrPut(key, newEl)
}

// Remove removes the element indexed with key.
func (c *Cache[T]) Remove(key uint64) {
	shard := key % shardSize
	c.shards[shard].Remove(key)
}

func (s *shard[T]) Get(key uint64) (el T, exists bool) {
	s.RLock()
	el, exists = s.items[key]
	s.RUnlock()
	return
}

func (s *shard[T]) Put(key uint64, el T) {
	s.Lock()
	if len(s.items) >= s.size {
		if _, ok := s.items[key]; !ok {
			for k := range s.items {
				delete(s.items, k)
				break
			}
		}
	}
	s.items[key] = el
	s.Unlock()
}

func (s *shard[T]) GetOrPut(key uint64, newEl T) (el T, exists bool) {
	s.Lock()
	el, exists = s.items[key]
	if !exists {
		if len(s.items) >= s.size {
			for k := range s.items {
				delete(s.items, k)
				break
			}
		}
		s.items[key] = newEl
		el = newEl
	}
	s.Unlock()
	return
}

func (s *shard[T]) Remove(key uint64) {
	s.Lock()
	delete(s.items, key)
	s.Unlock()
}
