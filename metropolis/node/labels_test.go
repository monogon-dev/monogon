package node

import (
	"errors"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation"
)

func TestValidateLabelKey(t *testing.T) {
	for i, te := range []struct {
		in   string
		want error
	}{
		{"foo", nil},
		{"example.com/", ErrLabelEmpty},
		{"foo-bar.baz_barfoo", nil},
		{"-", ErrLabelInvalidFirstCharacter},
		{"-invalid", ErrLabelInvalidFirstCharacter},
		{"invalid-", ErrLabelInvalidLastCharacter},
		{"", ErrLabelEmpty},
		{"accordingtoallknownlawsofaviationthereisnowaythatabeeshouldbeabletofly", ErrLabelTooLong},
		{"example.com/annotation", nil},
		{"/annotation", ErrLabelEmptyPrefix},
		{"_internal.example.com/annotation", ErrLabelInvalidPrefix},
		{"./annotation", ErrLabelInvalidPrefix},
		{"../annotation", ErrLabelInvalidPrefix},
		{"tcp:80.example.com/annotation", ErrLabelInvalidPrefix},
		{"github.com/monogon-dev/monogon/annotation", ErrLabelInvalidPrefix},
	} {
		if got := ValidateLabelKey(te.in); !errors.Is(got, te.want) {
			t.Errorf("%d (%q): wanted %v, got %v", i, te.in, te.want, got)
		}
	}
}

func TestValidateLabelValue(t *testing.T) {
	for i, te := range []struct {
		in   string
		want error
	}{
		{"foo", nil},
		{"foo-bar.baz_barfoo", nil},
		{"-", ErrLabelInvalidFirstCharacter},
		{"-invalid", ErrLabelInvalidFirstCharacter},
		{"invalid-", ErrLabelInvalidLastCharacter},
		{"", nil},
		{"accordingtoallknownlawsofaviationthereisnowaythatabeeshouldbeabletofly", ErrLabelTooLong},
		{"example.com/annotation", ErrLabelInvalidCharacter},
		{"/annotation", ErrLabelInvalidFirstCharacter},
		{"_internal.example.com/annotation", ErrLabelInvalidFirstCharacter},
		{"./annotation", ErrLabelInvalidFirstCharacter},
		{"../annotation", ErrLabelInvalidFirstCharacter},
		{"tcp:80.example.com/annotation", ErrLabelInvalidCharacter},
		{"github.com/monogon-dev/monogon/annotation", ErrLabelInvalidCharacter},
	} {
		// Test our implementation against test cases.
		if got := ValidateLabelValue(te.in); !errors.Is(got, te.want) {
			t.Errorf("%d (%q): wanted %v, got %v", i, te.in, te.want, got)
		}
		// Validate test cases against Kubernetes.
		if errs := validation.IsValidLabelValue(te.in); (te.want == nil) != (len(errs) == 0) {
			t.Errorf("%d (%q): wanted %v, kubernetes implementation returned %v", i, te.in, te.want, errs)
		}
	}
}
