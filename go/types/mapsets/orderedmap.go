package mapsets

import (
	"cmp"
	"sort"
)

// OrderedMap is a map from K to V which provides total ordering as defined by
// the keys used.
//
// The K type used must implement total ordering and comparability per the Key
// interface, and must not be mutated after being inserted (or at least, must
// keep the same ordering expressed as part of its Key implementation). A string
// (or string wrapping type) is a good key, as strings in Go are immutable and
// act like pointers.
//
// The values used will be copied by the map implementation. For structured
// types, defining V to be a pointer to a structure is probably the correct
// choice.
//
// An empty OrderedMap is ready to use and will contain no elements.
//
// It is not safe to be used by multiple goroutines. It can be locked behind a
// sync.RWMutex, with a read lock taken for any method that does not mutate the
// map, and with a write lock taken for any method that does mutate it.
type OrderedMap[K Key, V any] struct {
	keys   []K
	values map[K]V
}

// Key must be implemented by keys used by OrderedMap. Most 'typical' key types
// (string, integers, etc.) already implement it.
type Key interface {
	comparable
	cmp.Ordered
}

func (s *OrderedMap[K, V]) sort() {
	sort.Slice(s.keys, func(i, j int) bool {
		return s.keys[i] < s.keys[j]
	})
}

// Insert a given value at a given key. If a value at this key already exists, it
// will be overwritten with the new value.
//
// This method mutates the map.
func (s *OrderedMap[K, V]) Insert(k K, v V) {
	if s.values == nil {
		s.values = make(map[K]V)
	}

	if _, ok := s.values[k]; !ok {
		s.keys = append(s.keys, k)
		s.sort()
	}
	s.values[k] = v
}

// Get a value at a given key. If there is no value for the given key, an empty V
// and false will be returned.
//
// This returns a copy of the stored value.
func (s *OrderedMap[K, V]) Get(k K) (V, bool) {
	var zero V

	if s.values == nil {
		return zero, false
	}

	v, ok := s.values[k]
	return v, ok
}

// Keys returns a copy of the keys of this map, ordered according to the Key
// implementation used.
func (s *OrderedMap[K, V]) Keys() []K {
	if s.values == nil {
		return nil
	}

	keys := make([]K, len(s.keys))
	for i, k := range s.keys {
		keys[i] = k
	}
	return keys
}

// Values returns a copy of all the keys and values of this map, ordered
// according to the key implementation used.
func (s *OrderedMap[K, V]) Values() []KeyValue[K, V] {
	var res []KeyValue[K, V]
	for _, k := range s.keys {
		res = append(res, KeyValue[K, V]{
			Key:   k,
			Value: s.values[k],
		})
	}
	return res
}

// KeyValue represents a value V at a key K of an OrderedMap.
type KeyValue[K Key, V any] struct {
	Key   K
	Value V
}

// Delete the value at the given key, if present.
//
// This method mutates the map.
func (s *OrderedMap[K, V]) Delete(k K) {
	// Short-circuit delete in empty map.
	if s.values == nil {
		return
	}

	// Key not set? Just return.
	if _, ok := s.values[k]; !ok {
		return
	}

	// Iterate over keys to find index.
	toRemove := -1
	for i, k2 := range s.keys {
		if k2 == k {
			toRemove = i
			break
		}
	}
	if toRemove == -1 {
		panic("programming error: keys and values out of sync in OrderedMap")
	}

	// No need to re-sort, as the keys were already sorted.
	left := s.keys[0:toRemove]
	right := s.keys[toRemove+1:]
	s.keys = append(left, right...)
	delete(s.values, k)
}

// Clear removes all keys/values from the OrderedMap.
func (s *OrderedMap[K, V]) Clear() {
	s.keys = nil
	s.values = nil
}

// Count returns the number of keys/values in this OrderedMap.
func (s *OrderedMap[K, V]) Count() int {
	return len(s.keys)
}

// Clone (perform a deep copy) of the OrderedMap, copying all keys and values.
func (s *OrderedMap[K, V]) Clone() OrderedMap[K, V] {
	// Short-circuit clone of empty map.
	if s.values == nil {
		return OrderedMap[K, V]{}
	}

	keys := make([]K, len(s.keys))
	values := make(map[K]V)
	for i, k := range s.keys {
		keys[i] = k
		values[k] = s.values[k]
	}

	return OrderedMap[K, V]{
		keys:   keys,
		values: values,
	}
}

// Replace all contents of this map with the contents of another map. The other
// map must not be concurrently accessed.
//
// This method mutates the map.
func (s *OrderedMap[K, V]) Replace(o *OrderedMap[K, V]) {
	if s.values == nil {
		s.values = make(map[K]V)
	}

	cloned := o.Clone()
	s.keys = cloned.keys
	s.values = cloned.values
}

// Cycle returns a CycleIterator referencing this OrderedMap.
//
// This iterator will cycle through all the keys/values of the OrderedMap, then
// wrap around again when it reaches the end of the OrderedMap. This behaviour is
// designed to be used in round-robin mechanisms.
//
// As the OrderedMap changes, the returned CycleIterator will also change:
//   - If a key is removed or added, the CycleIterator is guaranteed to return it
//     within one entire cycle through all the values.
//   - If a value is changed, the new value is guaranteed to be returned on the
//     next iteration when the previous value would have been returned.
//
// Generally, it is safe to hold on to OrderedMap pointers and mutate the
// underlying OrderedMap.
//
// However, concurrent access to the OrderedMap and CycleIterator from different
// goroutines is not safe. Access to the CycleIterator should be guarded behind
// the same mutual exclusion mechanism as access to the OrderedMap.
func (s *OrderedMap[K, V]) Cycle() CycleIterator[K, V] {
	return CycleIterator[K, V]{
		set: s,
	}
}

// CycleIterator for an OrderedMap. See OrderedMap.Cycle for more details.
type CycleIterator[K Key, V any] struct {
	set *OrderedMap[K, V]
	ix  int
}

// Reset the iteration to the beginning of the ordered map.
func (c *CycleIterator[K, V]) Reset() {
	c.ix = 0
}

// Next returns the next key and value of the OrderedMap and a bool describing
// the iteration state. If the OrderedMap is empty (and thus can't provide any
// keys/values, including for this iteration), false will be returned in the
// third value.
func (c *CycleIterator[K, V]) Next() (K, V, bool) {
	var zeroK K
	var zeroV V

	if len(c.set.keys) == 0 {
		return zeroK, zeroV, false
	}

	if c.ix >= len(c.set.keys) {
		c.ix = 0
	}

	key := c.set.keys[c.ix]
	val := c.set.values[key]
	c.ix++
	return key, val, true
}
