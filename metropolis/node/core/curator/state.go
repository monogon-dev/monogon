package curator

import (
	"fmt"
	"strings"

	"go.etcd.io/etcd/clientv3"
)

// etcdPrefix is the location of some data in etcd, with each data element keyed
// by some unique string key.
//
// Prefixes are /-delimited filesystem-like paths, eg. prefix `/foo/bar/` could
// contain an item with ID `one` at etcd key `/foo/bar/one`, and an item with ID
// `two` at etcd key `/foo/bar/two`.
//
// Given etcd's lexicographic range operation, this allows for retrieval of all
// items under a given prefix.
//
// An etcdPrefix should be built using the newEtcdPrefix function.
type etcdPrefix struct {
	// parts are the parts of an etcd prefix path, eg. for prefix '/foo/bar/' parts
	// would be ["foo", "bar"].
	parts []string
}

// newEtcdPrefix returns an etcdPrefix for the given string representation of a
// prefix. The string representation must start and end with a / character, and
// must contain at least one path component (ie. one /-delimited component) with
// no empty components.
func newEtcdPrefix(p string) (*etcdPrefix, error) {
	parts := strings.Split(p, "/")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid path: must contain at least one component")
	}
	// Expect ["", "foo", "bar", ""] for /foo/bar/. Ie., everything but the first
	// and part path part must be non-empty, and the first and last path parts must
	// be empty.
	for i, p := range parts {
		if i == 0 || i == len(parts)-1 {
			if p != "" {
				return nil, fmt.Errorf("invalid path: must start and end with a /")
			}
		} else {
			if p == "" {
				return nil, fmt.Errorf("invalid path: must not contain repeated / characters")
			}
		}
	}

	// Omit the leading and trailing parts, ie. keep ["foo", "bar"] for /foo/bar/.
	parts = parts[1 : len(parts)-1]
	return &etcdPrefix{parts}, nil
}

// mustNewEtcdPrefix returns an etcdPrefix given the string representation of a
// path, or panics if the string representation is invalid (see newEtcdPrefix
// for more information).
func mustNewEtcdPrefix(p string) etcdPrefix {
	res, err := newEtcdPrefix(p)
	if err != nil {
		panic(err)
	}
	return *res
}

// Key returns the key for an item within an etcdPrefix, or an error if the
// given ID is invalid (ie. contains a / character).
func (p etcdPrefix) Key(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("invalid id: cannot be empty")
	}
	if strings.Contains(id, "/") {
		return "", fmt.Errorf("invalid id: cannot contain / character")
	}
	path := append(p.parts, id)
	return "/" + strings.Join(path, "/"), nil
}

// keyRange returns a pair of [start, end) keys for use with etcd range queries
// to retrieve all keys under a given prefix.
func (p etcdPrefix) KeyRange() (start, end string) {
	// Range from /foo/bar/ ... to /foo/bar0 ('0' is one ASCII codepoint after '/').
	start = "/" + strings.Join(p.parts, "/") + "/"
	end = "/" + strings.Join(p.parts, "/") + "0"
	return
}

// Range returns an etcd clientv3.Op that represents a Range Get request over
// all the keys within this etcdPrefix.
func (p etcdPrefix) Range() clientv3.Op {
	start, end := p.KeyRange()
	return clientv3.OpGet(start, clientv3.WithRange(end))
}

// ExtractID carves out the etcdPrefix ID from an existing etcd key for this
// range, ie. is the reverse of the Key() function.
//
// If the given etcd key does not correspond to a valid ID, an empty string is
// returned.
func (p etcdPrefix) ExtractID(key string) string {
	strPrefix := "/" + strings.Join(p.parts, "/") + "/"
	if !strings.HasPrefix(key, strPrefix) {
		return ""
	}
	id := key[len(strPrefix):]
	if strings.Contains(id, "/") {
		return ""
	}
	return id
}
