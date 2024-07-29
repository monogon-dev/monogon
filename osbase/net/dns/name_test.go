package dns

import (
	"testing"
)

func TestIsSubDomain(t *testing.T) {
	cases := []struct {
		parent, child string
		expected      bool
	}{
		{".", ".", true},
		{".", "test.", true},
		{"example.com.", "example.com.", true},
		{"example.com.", "www.example.com.", true},
		{"example.com.", "xample.com.", false},
		{"example.com.", "www.axample.com.", false},
		{"example.com.", "wwwexample.com.", false},
		{"example.com.", `www\.example.com.`, false},
		{"example.com.", `www\\.example.com.`, true},
	}
	for _, c := range cases {
		if IsSubDomain(c.parent, c.child) != c.expected {
			t.Errorf("IsSubDomain(%q, %q): expected %v", c.parent, c.child, c.expected)
		}
	}
}

func TestSplitLastLabel(t *testing.T) {
	cases := []struct {
		name, rest, label string
	}{
		{"", "", ""},
		{".", "", ""},
		{"com.", "", "com"},
		{"www.example.com", "www.example.", "com"},
		{"www.example.com.", "www.example.", "com"},
		{`www.example\.com.`, "www.", `example\.com`},
		{`www.example\\.com.`, `www.example\\.`, "com"},
	}
	for _, c := range cases {
		rest, label := SplitLastLabel(c.name)
		if rest != c.rest || label != c.label {
			t.Errorf("SplitLastLabel(%q) = (%q, %q), expected (%q, %q)", c.name, rest, label, c.rest, c.label)
		}
	}
}

func TestParseReverse(t *testing.T) {
	cases := []struct {
		name  string
		ip    string
		bits  int
		extra bool
	}{
		{"example.", "invalid IP", 0, false},
		{"0.10.200.255.in-addr.arpa.", "255.200.10.0", 32, false},
		{"7.6.45.123.in-addr.arpa.", "123.45.6.7", 32, false},
		{"6.45.123.in-addr.arpa.", "123.45.6.0", 24, false},
		{"45.123.in-addr.arpa.", "123.45.0.0", 16, false},
		{"123.in-addr.arpa.", "123.0.0.0", 8, false},
		{"in-addr.arpa.", "0.0.0.0", 0, false},
		{"8.7.6.45.123.in-addr.arpa.", "123.45.6.7", 32, true}, // too many fields
		{".6.45.123.in-addr.arpa.", "123.45.6.0", 24, true},    // empty field
		{"7.06.45.123.in-addr.arpa.", "123.45.0.0", 16, true},  // leading 0
		{"7.256.45.123.in-addr.arpa.", "123.45.0.0", 16, true}, // number too large
		{"a6.45.123.in-addr.arpa.", "123.45.0.0", 16, true},    // invalid character
		{`7\.6.45.123.in-addr.arpa.`, "123.45.0.0", 16, true},  // escaped .
		{"0.6.45.123in-addr.arpa.", "invalid IP", 0, false},    // missing .
		{
			"0.1.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.",
			"::fedc:ba98:7654:3210",
			128,
			false,
		},
		{
			"1.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.",
			"::fedc:ba98:7654:3210",
			124,
			false,
		},
		{
			"2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.",
			"::fedc:ba98:7654:3200",
			120,
			false,
		},
		{
			"3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.",
			"::fedc:ba98:7654:3000",
			116,
			false,
		},
		{
			"2.ip6.arpa.",
			"2000::",
			4,
			false,
		},
		{
			"ip6.arpa.",
			"::",
			0,
			false,
		},
		{
			"0.0.1.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.", // too long
			"::fedc:ba98:7654:3210",
			128,
			true,
		},
		{
			"01.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.", // missing dot
			"::fedc:ba98:7654:3200",
			120,
			true,
		},
		{
			"001.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.", // missing dot
			"::fedc:ba98:7654:3200",
			120,
			true,
		},
		{
			`0.1\.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.`, // escaped dot
			"::fedc:ba98:7654:3000",
			116,
			true,
		},
		{
			"g.1.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.", // invalid character
			"::fedc:ba98:7654:3210",
			124,
			true,
		},
	}
	for _, c := range cases {
		ip, bits, extra := ParseReverse(c.name)
		if ip.String() != c.ip || bits != c.bits || extra != c.extra {
			t.Errorf("ParseReverse(%q) = (%s, %v, %v), expected (%s, %v, %v)", c.name, ip, bits, extra, c.ip, c.bits, c.extra)
		}
	}
}

func BenchmarkParseReverseIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseReverse("7.6.45.123.in-addr.arpa.")
	}
}

func BenchmarkParseReverseIPv6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseReverse("0.1.2.3.4.5.6.7.8.9.a.b.c.d.e.f.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.")
	}
}
