package dns

import (
	"net/netip"
	"strings"
)

// IsSubDomain returns true if child is the same as or a subdomain of parent.
// Both names should be in canonical form.
func IsSubDomain(parent, child string) bool {
	offset := len(child) - len(parent)
	if offset < 0 || child[offset:] != parent {
		return false
	}
	if offset == 0 || parent == "." {
		return true
	}
	if child[offset-1] != '.' {
		return false
	}
	j := offset - 2
	for j >= 0 && child[j] == '\\' {
		j--
	}
	return (offset-j)%2 == 0
}

// SplitLastLabel splits off the last label of a domain name. For example,
// "www.example.com." is split into "www.example." and "com".
func SplitLastLabel(name string) (rest string, label string) {
	labelEnd := len(name)
	if labelEnd != 0 && name[labelEnd-1] == '.' {
		labelEnd--
	}
	labelStart := labelEnd
	for ; labelStart > 0; labelStart-- {
		if name[labelStart-1] != '.' {
			continue
		}
		j := labelStart - 2
		for j >= 0 && name[j] == '\\' {
			j--
		}
		if (labelStart-j)%2 != 0 {
			continue
		}
		break
	}
	return name[:labelStart], name[labelStart:labelEnd]
}

// ParseReverse parses name as a reverse lookup name. If name is not a reverse
// name, the returned IP is invalid. The second return value indicates how many
// bits of the address are present. The third return value is true if there are
// extra labels before the reverse name.
func ParseReverse(name string) (ip netip.Addr, bits int, extra bool) {
	if strings.HasSuffix(name, "in-addr.arpa.") {
		var ip [4]uint8
		field := 0
		pos := len(name) - len("in-addr.arpa.") - 1
		for pos >= 0 && field < 4 {
			if name[pos] != '.' {
				break
			}
			nextPos := pos - 1
			for nextPos >= 0 && name[nextPos] >= '0' && name[nextPos] <= '9' {
				nextPos--
			}
			val := 0
			for valPos := nextPos + 1; valPos < pos; valPos++ {
				val = val*10 + int(name[valPos]) - '0'
			}
			valLen := pos - nextPos - 1
			if valLen == 0 || valLen > 3 || (valLen != 1 && name[nextPos+1] == '0') || val > 255 {
				// Number is empty, or too long, or has leading zero, or is too large.
				break
			}
			ip[field] = uint8(val)
			field++
			pos = nextPos
		}
		if pos >= 0 {
			// We did not parse the entire name.
			j := pos - 1
			for j >= 0 && name[j] == '\\' {
				j--
			}
			if name[pos] != '.' || (pos-j)%2 == 0 {
				// The last label we parsed was not terminated by a non-escaped dot.
				field--
				if field < 0 {
					return netip.Addr{}, 0, false
				}
				ip[field] = 0
			}
		}
		return netip.AddrFrom4(ip), field * 8, pos >= 0
	}

	if strings.HasSuffix(name, "ip6.arpa.") {
		var ip [16]uint8
		field := 0
		half := false

		pos := len(name) - len("ip6.arpa.") - 1
		for pos > 0 && field < 16 {
			if name[pos] != '.' {
				break
			}
			var nibble uint8
			if name[pos-1] >= '0' && name[pos-1] <= '9' {
				nibble = name[pos-1] - '0'
			} else if name[pos-1] >= 'a' && name[pos-1] <= 'f' {
				nibble = name[pos-1] - 'a' + 10
			} else {
				break
			}
			if half {
				ip[field] |= nibble
				field++
				half = false
			} else {
				ip[field] = nibble << 4
				half = true
			}
			pos -= 2
		}
		if pos >= 0 {
			// We did not parse the entire name.
			j := pos - 1
			for j >= 0 && name[j] == '\\' {
				j--
			}
			if name[pos] != '.' || (pos-j)%2 == 0 {
				// The last label we parsed was not terminated by a non-escaped dot.
				if half {
					half = false
					ip[field] = 0
				} else {
					half = true
					field--
					if field < 0 {
						return netip.Addr{}, 0, false
					}
					ip[field] &= 0xf0
				}
			}
		}
		bits := field * 8
		if half {
			bits += 4
		}
		return netip.AddrFrom16(ip), bits, pos >= 0
	}

	return netip.Addr{}, 0, false
}
