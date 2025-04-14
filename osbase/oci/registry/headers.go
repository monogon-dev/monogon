// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package registry

import "strings"

type authenticateChallenge struct {
	scheme string
	info   string
	params map[string]string
}

// parseAuthenticateHeader parses the values of a WWW-Authenticate HTTP header.
// parameter names are converted to lower case.
// If any value fails to parse, it returns nil.
func parseAuthenticateHeader(authenticate []string) []authenticateChallenge {
	// From RFC 9110:
	// WWW-Authenticate = #challenge
	// challenge = auth-scheme [ 1*SP ( token68 / #auth-param ) ]
	// auth-scheme = token
	// token68 = 1*( ALPHA / DIGIT / "-" / "." / "_" / "~" / "+" / "/" ) *"="
	// auth-param = token BWS "=" BWS ( token / quoted-string )
	// #element => [ element ] *( OWS "," OWS [ element ] )
	// token = 1*tchar
	// tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*"
	//       / "+" / "-" / "." / "^" / "_" / "`" / "|" / "~"
	//       / DIGIT / ALPHA
	// quoted-string = DQUOTE *( qdtext / quoted-pair ) DQUOTE
	// qdtext = HTAB / SP / %x21 / %x23-5B / %x5D-7E / obs-text
	// quoted-pair = "\" ( HTAB / SP / VCHAR / obs-text )
	// obs-text = %x80-FF
	// OWS = *( SP / HTAB )
	// BWS = OWS
	// VCHAR = %x21-7E

	var challenges []authenticateChallenge
	for _, a := range authenticate {
		for {
			a = strings.TrimLeft(a, " \t,") // Consume commas and OWS
			if a == "" {
				break
			}
			var scheme string
			scheme, a = scanToken(a) // Consume auth-scheme
			if scheme == "" {
				return nil
			}
			challenge := authenticateChallenge{
				scheme: scheme,
			}
			if !strings.HasPrefix(a, " ") { // Check for 1*SP
				a = strings.TrimLeft(a, " \t") // Consume OWS
				if a != "" && a[0] != ',' {    // Check for mandatory comma
					return nil
				}
				challenges = append(challenges, challenge)
				continue
			}
			a = strings.TrimLeft(a, " ") // Consume 1*SP

			// Check for token68
			i := 0
			for i < len(a) && charType[a[i]]&charTypeToken68 != 0 {
				i++
			}
			if i != 0 {
				for i < len(a) && a[i] == '=' { // Consume *"="
					i++
				}
				remain := strings.TrimLeft(a[i:], " \t") // Consume OWS
				if remain == "" || remain[0] == ',' {    // Check for mandatory comma
					// Confirmed token68
					challenge.info = a[:i]
					challenges = append(challenges, challenge)
					a = remain
					continue
				}
			}

			challenge.params = make(map[string]string)
			for {
				// Check for auth-param
				remain := strings.TrimLeft(a, " \t,") // Consume commas and OWS
				var name string
				name, remain = scanToken(remain) // Consume token
				if name == "" {
					break
				}
				remain = strings.TrimLeft(remain, " \t") // Consume BWS
				var ok bool
				if remain, ok = strings.CutPrefix(remain, "="); !ok { // Consume "="
					break
				}
				remain = strings.TrimLeft(remain, " \t") // Consume BWS
				var value string
				if remain, ok = strings.CutPrefix(remain, `"`); ok { // Check for quoted-string
					i := 0
					for i < len(remain) {
						if charType[remain[i]]&charTypeQdtext != 0 {
							i++
						} else if remain[i] == '\\' && i+1 < len(remain) && charType[remain[i+1]]&charTypeQuotedPair != 0 {
							value += remain[:i]
							remain = remain[i+1:] // Drop the backslash to unescape the string
							i = 1
						} else {
							break
						}
					}
					value += remain[:i]
					remain = remain[i:]
					if remain, ok = strings.CutPrefix(remain, `"`); !ok { // Consume quote
						break
					}
				} else {
					value, remain = scanToken(remain) // Consume token
					if value == "" {
						break
					}
				}
				// Confirmed auth-param
				name = strings.ToLower(name) // name is case-insensitive
				if _, ok := challenge.params[name]; ok {
					return nil // each parameter name MUST only occur once
				}
				challenge.params[name] = value
				a = remain
				a = strings.TrimLeft(a, " \t") // Consume OWS
				if a != "" && a[0] != ',' {    // Check for mandatory comma
					return nil
				}
			}
			challenges = append(challenges, challenge)
			a = strings.TrimLeft(a, " \t") // Consume OWS
			if a != "" && a[0] != ',' {    // Check for mandatory comma
				return nil
			}
		}
	}
	return challenges
}

var charType [256]uint8

const (
	charTypeToken = 1 << iota
	charTypeToken68
	charTypeQdtext
	charTypeQuotedPair
)

func init() {
	for _, c := range "!#$%&'*+-.^_`|~" {
		charType[c] |= charTypeToken
	}
	for _, c := range "-._~+/" {
		charType[c] |= charTypeToken68
	}
	for c := range 256 {
		if '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
			charType[c] |= charTypeToken | charTypeToken68
		}
		if c == '\t' || c == ' ' || 0x21 <= c && c != 0x7f {
			if c != '"' && c != '\\' {
				charType[c] |= charTypeQdtext
			}
			charType[c] |= charTypeQuotedPair
		}
	}
}

func scanToken(s string) (token string, remain string) {
	for i := range len(s) {
		if charType[s[i]]&charTypeToken == 0 {
			return s[:i], s[i:]
		}
	}
	return s, ""
}
