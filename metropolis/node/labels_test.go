// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
		{"_internal.example.com/annotation", errDomainNameInvalid},
		{"./annotation", errDomainNameInvalid},
		{"../annotation", errDomainNameInvalid},
		{"tcp:80.example.com/annotation", errDomainNameInvalid},
		{"80/annotation", errDomainNameEndsInNumber},
		{"github.com/monogon/monogon/annotation", ErrLabelInvalidPrefix},
	} {
		if got := ValidateLabelKey(te.in); !errors.Is(got, te.want) {
			t.Errorf("%d (%q): wanted %v, got %v", i, te.in, te.want, got)
		}
		if ValidateLabelKey(te.in) == nil {
			if errs := validation.IsQualifiedName(te.in); len(errs) != 0 {
				t.Errorf("%d (%q): is not a valid Kubernetes qualified name: %v", i, te.in, errs)
			}
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
		{"github.com/monogon/monogon/annotation", ErrLabelInvalidCharacter},
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
