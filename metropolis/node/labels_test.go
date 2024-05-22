package node

import (
	"errors"
	"testing"
)

func TestValidateLabelKeyValue(t *testing.T) {
	for i, te := range []struct {
		in   string
		want error
	}{
		{"foo", nil},
		{"foo-bar.baz_barfoo", nil},
		{"-", ErrLabelInvalidFirstCharacter},
		{"-invalid", ErrLabelInvalidFirstCharacter},
		{"invalid-", ErrLabelInvalidLastCharacter},
		{"", ErrLabelEmpty},
		{"accordingtoallknownlawsofaviationthereisnowaythatabeeshouldbeabletofly", ErrLabelTooLong},
		{"example.com/annotation", ErrLabelInvalidCharacter},
	} {
		if got := ValidateLabel(te.in); !errors.Is(got, te.want) {
			t.Errorf("%d: wanted %v, got %v", i, te.want, got)
		}
	}
}
