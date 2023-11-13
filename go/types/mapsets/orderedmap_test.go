package mapsets

import (
	"strings"
	"testing"
)

// TestOrderedMap exercises the basics of an OrderedMap.
func TestOrderedMap(t *testing.T) {
	m := OrderedMap[string, string]{}

	if want, got := 0, len(m.Keys()); want != got {
		t.Errorf("empty map should have %d keys, got %d", want, got)
	}

	// Test insertion.
	m.Insert("a", "foo")
	if got, ok := m.Get("a"); !ok || got != "foo" {
		t.Errorf("Wanted value 'foo', got %q", got)
	}
	// Test update.
	m.Insert("a", "bar")
	if got, ok := m.Get("a"); !ok || got != "bar" {
		t.Errorf("Wanted value 'bar', got %q", got)
	}

	// Helper to ensure that the keys returned by the map are the same as keystring
	// split by commas.
	keysEq := func(m *OrderedMap[string, string], keystring string) {
		t.Helper()
		keys := strings.Split(keystring, ",")
		// Support keystring "" indicating an empty map is expected.
		if len(keystring) == 0 {
			keys = nil
		}
		if want, got := len(keys), m.Count(); want != got {
			t.Errorf("Wanted count %d, got %d", want, got)
		}
		if want, got := keys, m.Keys(); len(want) == len(got) {
			for i, k := range want {
				if got[i] != k {
					t.Errorf("Wanted key %d %q, got %q", i, want[i], got[i])
				}
			}
		} else {
			t.Errorf("Wanted keys %v, got %v", want, got)
		}
	}

	// Test keys. They should get sorted.
	m.Insert("c", "baz")
	m.Insert("b", "barfoo")
	keysEq(&m, "a,b,c")
	// Test deleting and re-inserting a key.
	m.Delete("a")
	keysEq(&m, "b,c")
	m.Insert("a", "a")
	keysEq(&m, "a,b,c")
	// Test clone.
	n := m.Clone()
	// Inserting into the original OrderedMap shouldn't impact the clone.
	m.Insert("z", "zz")
	keysEq(&m, "a,b,c,z")
	keysEq(&n, "a,b,c")
	// Test replace.
	n.Replace(&m)
	keysEq(&n, "a,b,c,z")
	// Test Values.
	values := n.Values()
	if want, got := 4, len(values); want != got {
		t.Errorf("Expected %d values, got %d", want, got)
	} else {
		for i, k := range []string{"a", "b", "c", "z"} {
			if want, got := k, values[i].Key; want != got {
				t.Errorf("Key %d should've been %q, got %q", i, want, got)
			}
		}
		for i, v := range []string{"a", "barfoo", "baz", "zz"} {
			if want, got := v, values[i].Value; want != got {
				t.Errorf("Value %d should've been %q, got %q", i, want, got)
			}
		}
	}
	// Test Clear
	n.Clear()
	keysEq(&n, "")
}

// TestCycleIterator exercises the CycleIterator.
func TestCycleIterator(t *testing.T) {
	m := OrderedMap[string, string]{}
	it := m.Cycle()

	// Helper to ensure that the iterator returns a list of (key, values) string
	// pairs, as defined by keystring and valuestring, split by commas.
	valuesEq := func(keystring, valuestring string) {
		t.Helper()

		wantKeys := strings.Split(keystring, ",")
		if keystring == "" {
			wantKeys = nil
		}
		wantValues := strings.Split(valuestring, ",")
		if valuestring == "" {
			wantValues = nil
		}
		if len(wantKeys) != len(wantValues) {
			// Test programming error.
			t.Fatalf("keystring length != valuestring length")
		}

		var ix int
		for {
			// Out of test elements? We're done.
			if ix >= len(wantKeys) {
				return
			}

			k, v, ok := it.Next()
			if !ok {
				// Iterator empty while test cases present? Fail.
				if ix < len(wantKeys) {
					t.Errorf("Iterator empty")
				}
				break
			}
			if want, got := wantKeys[ix], k; want != got {
				t.Errorf("Iterator returned key %q at %d, wanted %q", got, ix, want)
			}
			if want, got := wantValues[ix], v; want != got {
				t.Errorf("Iterator returned value %q at %d, wanted %q", got, ix, want)
			}
			ix++
		}
	}

	// Empty iterator should fail immediately.
	it.Reset()
	valuesEq("", "")

	// Expect keys/values in order.
	m.Insert("z", "foo")
	m.Insert("a", "bar")
	m.Insert("c", "baz")
	valuesEq("a,c,z", "bar,baz,foo")

	// The iterator should cycle.
	valuesEq("a,c,z,a,c,z,a,c,z", "bar,baz,foo,bar,baz,foo,bar,baz,foo")

	// Reset iterator and read only first value.
	it.Reset()
	valuesEq("a", "bar")
	// Now insert a new key, the iterator should handle that.
	m.Insert("b", "hello")
	valuesEq("b,c,z", "hello,baz,foo")

	// Reset iterator and only read first value.
	it.Reset()
	valuesEq("a", "bar")
	// Now update an existing key. The iterator should continue.
	m.Insert("b", "goodbye")
	valuesEq("b,c,z", "goodbye,baz,foo")

	// Reset iterator and only read first value.
	it.Reset()
	valuesEq("a", "bar")
	// Now remove all values. The iterator should fail immediately.
	m.Replace(&OrderedMap[string, string]{})
	valuesEq("", "")
}
