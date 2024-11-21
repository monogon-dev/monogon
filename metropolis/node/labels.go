package node

import (
	"fmt"
	"regexp"
	"strings"

	cpb "source.monogon.dev/metropolis/proto/common"
)

var (
	reLabelFirstLast = regexp.MustCompile(`^[a-zA-Z0-9]$`)
	reLabelBody      = regexp.MustCompile(`^[a-zA-Z0-9\-._]*$`)

	// ErrLabelEmpty is returned by ValidateLabel if the label key/value is not at
	// least one character long.
	ErrLabelEmpty = fmt.Errorf("empty")
	// ErrLabelTooLong is returned by ValidateLabel if the label key/value is more
	// than 63 characters long.
	ErrLabelTooLong = fmt.Errorf("too long")
	// ErrLabelInvalidFirstCharacter is returned by ValidateLabel if the label
	// key/value contains an invalid character on the first position.
	ErrLabelInvalidFirstCharacter = fmt.Errorf("first character not a letter or number")
	// ErrLabelInvalidLastCharacter is returned by ValidateLabel if the label
	// key/value contains an invalid character on the last position.
	ErrLabelInvalidLastCharacter = fmt.Errorf("last character not a letter or number")
	// ErrLabelInvalidCharacter is returned by ValidateLabel if the label key/value
	// contains an invalid character.
	ErrLabelInvalidCharacter = fmt.Errorf("invalid character")
	ErrLabelEmptyPrefix      = fmt.Errorf("empty prefix")
	ErrLabelInvalidPrefix    = fmt.Errorf("invalid prefix")
)

const (
	// MaxLabelsPerNode is the absolute maximum of labels that can be attached to a
	// node.
	MaxLabelsPerNode = 128
)

func validatePrefix(prefix string) error {
	if prefix == "" {
		return ErrLabelEmptyPrefix
	}
	if err := validateDomainName(prefix); err != nil {
		return fmt.Errorf("invalid prefix: %w", err)
	}
	return nil
}

// ValidateLabelKey ensures that a given node label key is valid:
//
//  1. 1 to 63 characters long (inclusive);
//  2. Characters are all ASCII a-z A-Z 0-9 '_', '-' or '.';
//  3. The first character is ASCII a-z A-Z or 0-9.
//  4. The last character is ASCII a-z A-Z or 0-9.
//  5. Optional slash-delimited prefix which contains a valid 'DNS subdomain'.
//
// If it's valid, nil is returned. Otherwise, one of ErrLabelEmpty,
// ErrLabelTooLong, ErrLabelInvalidFirstCharacter or ErrLabelInvalidCharacter is
// returned.
func ValidateLabelKey(v string) error {
	// Split away prefix.
	parts := strings.Split(v, "/")
	switch len(parts) {
	case 1:
	case 2:
		prefix := parts[0]
		if err := validatePrefix(prefix); err != nil {
			return err
		}
		v = parts[1]
	default:
		return ErrLabelInvalidPrefix
	}

	if len(v) == 0 {
		return ErrLabelEmpty
	}

	if len(v) > 63 {
		return ErrLabelTooLong
	}
	if !reLabelFirstLast.MatchString(string(v[0])) {
		return ErrLabelInvalidFirstCharacter
	}
	if !reLabelFirstLast.MatchString(string(v[len(v)-1])) {
		return ErrLabelInvalidLastCharacter
	}
	// Body characters are a superset of the first/last characters, and we've already
	// checked those so we can check the entire string here.
	if !reLabelBody.MatchString(v) {
		return ErrLabelInvalidCharacter
	}
	return nil
}

// ValidateLabelValue ensures that a given node label value is valid:
//
//  1. 0 to 63 characters long (inclusive);
//  2. Characters are all ASCII a-z A-Z 0-9 '_', '-' or '.';
//  3. The first character is ASCII a-z A-Z or 0-9.
//  4. The last character is ASCII a-z A-Z or 0-9.
//
// If it's valid, nil is returned. Otherwise, one of ErrLabelTooLong,
// ErrLabelInvalidFirstCharacter, or ErrLabelInvalidLastCharacter,
// ErrLabelInvalidCharacter is returned.
func ValidateLabelValue(v string) error {
	if len(v) == 0 {
		return nil
	}
	if len(v) > 63 {
		return ErrLabelTooLong
	}
	if !reLabelFirstLast.MatchString(string(v[0])) {
		return ErrLabelInvalidFirstCharacter
	}
	if !reLabelFirstLast.MatchString(string(v[len(v)-1])) {
		return ErrLabelInvalidLastCharacter
	}
	// Body characters are a superset of the first character, and we've already
	// checked that so we can check the entire string here.
	if !reLabelBody.MatchString(v) {
		return ErrLabelInvalidCharacter
	}
	return nil
}

// GetNodeLabel retrieves a node label by key, returning its value or an empty
// string if no labels with this key is set on the node.
func GetNodeLabel(labels *cpb.NodeLabels, key string) string {
	for _, pair := range labels.Pairs {
		if pair.Key == key {
			return pair.Value
		}
	}
	return ""
}

// Labels on a node, a map from label key to value.
type Labels map[string]string

// Equals returns true if these Labels are equal to some others. Equality is
// defined by having the same set of keys and corresponding values.
func (l Labels) Equals(others Labels) bool {
	for k, v := range l {
		if v2, ok := others[k]; !ok || v != v2 {
			return false
		}
	}
	for k, v := range others {
		if v2, ok := l[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

// Filter returns a subset of labels for which pred returns true.
func (l Labels) Filter(pred func(k, v string) bool) Labels {
	res := make(Labels)
	for k, v := range l {
		if pred(k, v) {
			res[k] = v
		}
	}
	return res
}
