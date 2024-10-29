package node

import (
	"errors"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation"
)

func TestValidateDomainName(t *testing.T) {
	for _, te := range []struct {
		in   string
		want error
	}{
		{"example.com", nil},
		{"localhost", nil},
		{"123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.1.example.com", nil},
		{"123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.123456789.12.example.com", errDomainNameTooLong},
		{"ex_ample.com", errDomainNameInvalid},
		{"-.com", errDomainNameInvalid},
		{"example-.com", errDomainNameInvalid},
		{"-example.com", errDomainNameInvalid},
		{"1-1.com", nil},
		{"xn--h-0fa.com", nil},
		{".", errDomainNameInvalid},
		{"example..com", errDomainNameInvalid},
		{"example.com.", errDomainNameInvalid},
		{".example.com", errDomainNameInvalid},
		{"0.example.com", nil},
		{"01.example.com", nil},
		{"012345678901234567890123456789012345678901234567890123456789012.example.com", nil},
		{"0123456789012345678901234567890123456789012345678901234567890123.example.com", errDomainNameInvalid},
		{"1.1.1.1", errDomainNameEndsInNumber},
		{"example.123", errDomainNameEndsInNumber},
		{"0123456789", errDomainNameEndsInNumber},
		{"example.0x", errDomainNameEndsInNumber},
		{"0x0123456789abcdef", errDomainNameEndsInNumber},
		{"1.2.3.1a1", nil},
	} {
		if got := validateDomainName(te.in); !errors.Is(got, te.want) {
			t.Errorf("%q: wanted %v, got %v", te.in, te.want, got)
		}
		if validateDomainName(te.in) == nil {
			if errs := validation.IsDNS1123Subdomain(te.in); len(errs) > 0 {
				t.Errorf("%q: is not a valid Kubernetes domain: %v", te.in, errs)
			}
		}
	}
}
