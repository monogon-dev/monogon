package curator

import "testing"

func TestEtcdPrefixParse(t *testing.T) {
	for i, te := range []struct {
		p  string
		ok bool
	}{
		{"/foo/", true},
		{"/foo/bar/", true},

		{"/foo//", false},
		{"/foo//bar/", false},
		{"/foo/bar", false},
		{"foo", false},
		{"foo/", false},
		{"foo/bar", false},
	} {
		_, err := newEtcdPrefix(te.p)
		if te.ok {
			if err != nil {
				t.Errorf("Case %d: wanted nil, got err %v", i, err)
			}
		} else {
			if err == nil {
				t.Errorf("Case %d: wanted err, got nil", i)
			}
		}
	}
}

func TestEtcdPrefixKeyRange(t *testing.T) {
	p := mustNewEtcdPrefix("/foo/")

	// Test Key() functionality.
	key, err := p.Key("bar")
	if err != nil {
		t.Fatalf("Key(bar): %v", err)
	}
	if want, got := "/foo/bar", key; want != got {
		t.Errorf("Wrong key, wanted %q, got %q", want, got)
	}

	// Test Key() with invalid IDs.
	_, err = p.Key("")
	if err == nil {
		t.Error("Key(bar/baz) returned nil, wanted error")
	}
	_, err = p.Key("bar/baz")
	if err == nil {
		t.Error("Key(bar/baz) returned nil, wanted error")
	}

	// Test Range() functionality.
	op := p.Range()
	if want, got := "/foo/", string(op.KeyBytes()); want != got {
		t.Errorf("Wrong start key, wanted %q, got %q", want, got)
	}
	if want, got := "/foo0", string(op.RangeBytes()); want != got {
		t.Errorf("Wrong end key, wanted %q, got %q", want, got)
	}
}

func TestEtcdPrefixExtractID(t *testing.T) {
	p := mustNewEtcdPrefix("/foo/")

	for i, te := range []struct {
		key  string
		want string
	}{
		{"/foo/", ""},
		{"/foo0", ""},
		{"/foo", ""},
		{"bar", ""},

		{"/foo/bar", "bar"},
		{"/foo/bar/baz", ""},
	} {
		got := p.ExtractID(te.key)
		if te.want != got {
			t.Errorf("%d: ExtractID(%q) should have returned %q, got %q", i, te.key, te.want, got)
		}
	}
}
