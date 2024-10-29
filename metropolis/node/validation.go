package node

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	// domainNameMaxLength is the maximum length of a domain name supported by DNS
	// when represented without a trailing dot.
	domainNameMaxLength = 253

	// clusterDomainMaxLength is the maximum length of a cluster domain. Limiting
	// this to 80 allows for constructing subdomains of the cluster domain, where
	// the subdomain part can have length up to 172. With the joining dot, this
	// adds up to 253.
	clusterDomainMaxLength = 80
)

var (
	fmtDomainNameLabel       = `[a-z0-9]([-a-z0-9]{0,61}[a-z0-9])?`
	reDomainName             = regexp.MustCompile(`^` + fmtDomainNameLabel + `(\.` + fmtDomainNameLabel + `)*$`)
	reDomainNameEndsInNumber = regexp.MustCompile(`(^|\.)([0-9]+|0x[0-9a-f]*)$`)

	errDomainNameTooLong      = fmt.Errorf("too long, must have length at most %d", domainNameMaxLength)
	errDomainNameInvalid      = errors.New("must consist of labels separated by '.', where each label has between 1 and 63 lowercase letters, digits or '-', and must not start or end with '-'")
	errDomainNameEndsInNumber = errors.New("must not end in a number")

	errClusterDomainTooLong = fmt.Errorf("too long, must have length at most %d", clusterDomainMaxLength)
)

// validateDomainName returns an error if the passed string is not a valid
// domain name, according to these rules: The name must be a valid DNS name
// without a trailing dot. Labels must only consist of lowercase letters, digits
// or '-', and must not start or end with '-'. Additionally, the name must not
// end in a number, so that it won't be parsed as an IPv4 address.
func validateDomainName(d string) error {
	if len(d) > domainNameMaxLength {
		return errDomainNameTooLong
	}
	// This implements RFC 1123 domain validation. Additionally, it does not allow
	// uppercase, so that we don't need to implement case-insensitive matching.
	if !reDomainName.MatchString(d) {
		return errDomainNameInvalid
	}
	// This implements https://url.spec.whatwg.org/#ends-in-a-number-checker
	if reDomainNameEndsInNumber.MatchString(d) {
		return errDomainNameEndsInNumber
	}
	return nil
}

// ValidateClusterDomain returns an error if the passed string is not a valid
// cluster domain.
func ValidateClusterDomain(d string) error {
	if len(d) > clusterDomainMaxLength {
		return errClusterDomainTooLong
	}
	return validateDomainName(d)
}
