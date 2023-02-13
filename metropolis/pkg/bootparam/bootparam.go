// Package bootparam implements encoding and decoding of Linux kernel command
// lines as documented in
// https://docs.kernel.org/admin-guide/kernel-parameters.html
//
// The format is quite quirky and thus the implementation is mostly based
// on the code in the Linux kernel implementing the decoder and not the
// specification.
package bootparam

import (
	"errors"
	"fmt"
	"strings"
)

// Param represents a single boot parameter with or without a value
type Param struct {
	Param, Value string
	HasValue     bool
}

// Params represents a list of kernel boot parameters
type Params []Param

// Linux has for historical reasons an unusual definition of this function
// Taken from @linux//lib:ctype.c
func isSpace(r byte) bool {
	switch r {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0xa0:
		return true
	default:
		return false
	}
}

// Trim spaces as defined by Linux from the left of the string.
// This is only exported for tests, do not use this. Because of import loops
// as well as cgo restrictions this cannot be an internal function used by
// tests.
func TrimLeftSpace(s string) string {
	start := 0
	for ; start < len(s); start++ {
		c := s[start]
		if !isSpace(c) {
			break
		}
	}

	return s[start:]
}

func containsSpace(s string) bool {
	for i := 0; i < len(s); i++ {
		if isSpace(s[i]) {
			return true
		}
	}
	return false
}

func parseToken(token string) (p Param, err error) {
	if strings.HasPrefix(token, `=`) || strings.HasPrefix(token, `"=`) {
		return Param{}, errors.New("param contains `=` at first position, this causes broken behavior")
	}
	param, value, hasValue := strings.Cut(token, "=")

	if strings.HasPrefix(param, `"`) {
		p.Param = strings.TrimPrefix(param, `"`)
		if !hasValue {
			p.Param = strings.TrimSuffix(p.Param, `"`)
		}
	} else {
		p.Param = param
	}
	if hasValue {
		if strings.HasPrefix(value, `"`) {
			p.Value = strings.TrimSuffix(strings.TrimPrefix(value, `"`), `"`)
		} else if strings.HasPrefix(param, `"`) {
			p.Value = strings.TrimSuffix(value, `"`)
		} else {
			p.Value = value
		}
	}
	return
}

// Unmarshal decodes a Linux kernel command line and returns a list of kernel
// parameters as well as a rest section after the "--" parsing terminator.
func Unmarshal(cmdline string) (params Params, rest string, err error) {
	cmdline = TrimLeftSpace(cmdline)
	if pos := strings.IndexByte(cmdline, 0x00); pos != -1 {
		cmdline = cmdline[:pos]
	}
	var lastIdx int
	var inQuote bool
	var p Param
	for i := 0; i < len(cmdline); i++ {
		if isSpace(cmdline[i]) && !inQuote {
			token := cmdline[lastIdx:i]
			lastIdx = i + 1
			if TrimLeftSpace(token) == "" {
				continue
			}
			p, err = parseToken(token)
			if err != nil {
				return
			}

			// Stop processing and return everything left as rest
			if p.Param == "--" {
				rest = TrimLeftSpace(cmdline[lastIdx:])
				return
			}
			params = append(params, p)
		}
		if cmdline[i] == '"' {
			inQuote = !inQuote
		}
	}
	if len(cmdline)-lastIdx > 0 {
		token := cmdline[lastIdx:]
		if TrimLeftSpace(token) == "" {
			return
		}
		p, err = parseToken(token)
		if err != nil {
			return
		}

		// Stop processing, do not set rest as there is none
		if p.Param == "--" {
			return
		}
		params = append(params, p)
	}
	return
}

// Marshal encodes a set of kernel parameters and an optional rest string into
// a Linux kernel command line. It rejects data which is not encodable, which
// includes null bytes, double quotes in params as well as characters which
// contain 0xa0 in their UTF-8 representation (historical Linux quirk of
// treating that as a space, inherited from Latin-1).
func Marshal(params Params, rest string) (string, error) {
	if strings.IndexByte(rest, 0x00) != -1 {
		return "", errors.New("rest contains 0x00 byte, this is disallowed")
	}
	var strb strings.Builder
	for _, p := range params {
		if strings.ContainsRune(p.Param, '=') {
			return "", fmt.Errorf("invalid '=' character in param %q", p.Param)
		}
		// Technically a weird subset of double quotes can be encoded, but
		// this should probably not be done so just reject them all.
		if strings.ContainsRune(p.Param, '"') {
			return "", fmt.Errorf("invalid '\"' character in param %q", p.Param)
		}
		if strings.ContainsRune(p.Value, '"') {
			return "", fmt.Errorf("invalid '\"' character in value %q", p.Value)
		}
		if strings.IndexByte(p.Param, 0x00) != -1 {
			return "", fmt.Errorf("invalid null byte in param %q", p.Param)
		}
		if strings.IndexByte(p.Value, 0x00) != -1 {
			return "", fmt.Errorf("invalid null byte in value %q", p.Value)
		}
		// Linux treats 0xa0 as a space, even though it is a valid UTF-8
		// surrogate. This is unfortunate, but passing it through would
		// break the whole command line.
		if strings.IndexByte(p.Param, 0xa0) != -1 {
			return "", fmt.Errorf("invalid 0xa0 byte in param %q", p.Param)
		}
		if strings.IndexByte(p.Value, 0xa0) != -1 {
			return "", fmt.Errorf("invalid 0xa0 byte in value %q", p.Value)
		}
		if strings.ContainsRune(p.Param, '"') {
			return "", fmt.Errorf("invalid '\"' character in value %q", p.Value)
		}
		// This should be allowed according to the docs, but is in fact broken.
		if p.Value != "" && containsSpace(p.Param) {
			return "", fmt.Errorf("param %q contains spaces and value, this is unsupported", p.Param)
		}
		if p.Param == "--" {
			return "", errors.New("param '--' is reserved and cannot be used")
		}
		if p.Param == "" {
			return "", errors.New("empty params are not supported")
		}
		if containsSpace(p.Param) {
			strb.WriteRune('"')
			strb.WriteString(p.Param)
			strb.WriteRune('"')
		} else {
			strb.WriteString(p.Param)
		}
		if p.Value != "" {
			strb.WriteRune('=')
			if containsSpace(p.Value) {
				strb.WriteRune('"')
				strb.WriteString(p.Value)
				strb.WriteRune('"')
			} else {
				strb.WriteString(p.Value)
			}
		}
		strb.WriteRune(' ')
	}
	if len(rest) > 0 {
		strb.WriteString("-- ")
		// Starting whitespace will be dropped by the decoder anyways, do it
		// here to make the resulting command line nicer.
		strb.WriteString(TrimLeftSpace(rest))
	}
	return strb.String(), nil
}
