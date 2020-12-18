// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dhcp4c

import (
	"net"
	"testing"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/stretchr/testify/assert"
)

func TestLeaseDHCPServers(t *testing.T) {
	var tests = []struct {
		name     string
		lease    *Lease
		expected DNSServers
	}{{
		name:     "ReturnsNilWithNoLease",
		lease:    nil,
		expected: nil,
	}, {
		name: "DiscardsInvalidIPs",
		lease: &Lease{
			Options: dhcpv4.OptionsFromList(dhcpv4.OptDNS(net.IP{0, 0, 0, 0})),
		},
		expected: nil,
	}, {
		name: "DeduplicatesIPs",
		lease: &Lease{
			Options: dhcpv4.OptionsFromList(dhcpv4.OptDNS(net.IP{192, 0, 2, 1}, net.IP{192, 0, 2, 2}, net.IP{192, 0, 2, 1})),
		},
		expected: DNSServers{net.IP{192, 0, 2, 1}, net.IP{192, 0, 2, 2}},
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := test.lease.DNSServers()
			assert.Equal(t, test.expected, res)
		})
	}
}
