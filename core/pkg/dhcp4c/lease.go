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
	"time"

	"github.com/insomniacslk/dhcp/dhcpv4"
)

// Lease represents a DHCPv4 lease. It only consists of an IP, an expiration timestamp and options as all other
// relevant parts of the message have been normalized into their respective options. It also contains some smart
// getters for commonly-used options which extract only valid information from options.
type Lease struct {
	AssignedIP net.IP
	ExpiresAt  time.Time
	Options    dhcpv4.Options
}

// SubnetMask returns the SubnetMask option or the default mask if not set or invalid.
// It returns nil if the lease is nil.
func (l *Lease) SubnetMask() net.IPMask {
	if l == nil {
		return nil
	}
	mask := net.IPMask(dhcpv4.GetIP(dhcpv4.OptionSubnetMask, l.Options))
	if _, bits := mask.Size(); bits != 32 { // If given mask is not valid, use the default mask
		mask = l.AssignedIP.DefaultMask()
	}
	return mask
}

// IPNet returns an IPNet from the assigned IP and subnet mask.
// It returns nil if the lease is nil.
func (l *Lease) IPNet() *net.IPNet {
	if l == nil {
		return nil
	}
	return &net.IPNet{
		IP:   l.AssignedIP,
		Mask: l.SubnetMask(),
	}
}

// Router returns the first valid router from the DHCP router option or nil if none such exists.
// It returns nil if the lease is nil.
func (l *Lease) Router() net.IP {
	if l == nil {
		return nil
	}
	routers := dhcpv4.GetIPs(dhcpv4.OptionRouter, l.Options)
	for _, r := range routers {
		if r.IsGlobalUnicast() || r.IsLinkLocalUnicast() {
			return r
		}
	}
	// No (valid) router found
	return nil
}
