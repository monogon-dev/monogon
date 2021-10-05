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

// Package callback contains minimal callbacks for configuring the kernel with
// options received over DHCP.
//
// These directly configure the relevant kernel subsytems and need to own
// certain parts of them as documented on a per- callback basis to make sure
// that they can recover from restarts and crashes of the DHCP client.
// The callbacks in here are not suitable for use in advanced network scenarios
// like running multiple DHCP clients per interface via ClientIdentifier or
// when running an external FIB manager. In these cases it's advised to extract
// the necessary information from the lease in your own callback and
// communicate it directly to the responsible entity.
package callback

import (
	"fmt"
	"math"
	"net"
	"os"
	"time"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node/core/network/dhcp4c"
)

// Compose can be used to chain multiple callbacks
func Compose(callbacks ...dhcp4c.LeaseCallback) dhcp4c.LeaseCallback {
	return func(old, new *dhcp4c.Lease) error {
		for _, cb := range callbacks {
			if err := cb(old, new); err != nil {
				return err
			}
		}
		return nil
	}
}

func isIPNetEqual(a, b *net.IPNet) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	aOnes, aBits := a.Mask.Size()
	bOnes, bBits := b.Mask.Size()
	return a.IP.Equal(b.IP) && aOnes == bOnes && aBits == bBits
}

// ManageIP sets up and tears down the assigned IP address. It takes exclusive
// ownership of all IPv4 addresses on the given interface which do not have
// IFA_F_PERMANENT set, so it's not possible to run multiple dynamic addressing
// clients on a single interface.
func ManageIP(iface netlink.Link) dhcp4c.LeaseCallback {
	return func(old, new *dhcp4c.Lease) error {
		newNet := new.IPNet()

		addrs, err := netlink.AddrList(iface, netlink.FAMILY_V4)
		if err != nil {
			return fmt.Errorf("netlink failed to list addresses: %w", err)
		}

		for _, addr := range addrs {
			if addr.Flags&unix.IFA_F_PERMANENT == 0 {
				// Linux identifies addreses by IP, mask and peer (see
				// net/ipv4/devinet.find_matching_ifa in Linux 5.10).
				// So don't touch addresses which match on these properties as
				// AddrReplace will atomically reconfigure them anyways without
				// interrupting things.
				if isIPNetEqual(addr.IPNet, newNet) && addr.Peer == nil && new != nil {
					continue
				}

				if err := netlink.AddrDel(iface, &addr); !os.IsNotExist(err) && err != nil {
					return fmt.Errorf("failed to delete address: %w", err)
				}
			}
		}

		if new != nil {
			remainingLifetimeSecs := int(math.Ceil(new.ExpiresAt.Sub(time.Now()).Seconds()))
			newBroadcastIP := dhcpv4.GetIP(dhcpv4.OptionBroadcastAddress, new.Options)
			if err := netlink.AddrReplace(iface, &netlink.Addr{
				IPNet:       newNet,
				ValidLft:    remainingLifetimeSecs,
				PreferedLft: remainingLifetimeSecs,
				Broadcast:   newBroadcastIP,
			}); err != nil {
				return fmt.Errorf("failed to update address: %w", err)
			}
		}
		return nil
	}
}

// ManageDefaultRoute manages a default route through the first router offered
// by DHCP. It does nothing if DHCP doesn't provide any routers. It takes
// ownership of all RTPROTO_DHCP routes on the given interface, so it's not
// possible to run multiple DHCP clients on the given interface.
func ManageDefaultRoute(iface netlink.Link) dhcp4c.LeaseCallback {
	return func(old, new *dhcp4c.Lease) error {
		newRouter := new.Router()

		dhcpRoutes, err := netlink.RouteListFiltered(netlink.FAMILY_V4, &netlink.Route{
			Protocol:  unix.RTPROT_DHCP,
			LinkIndex: iface.Attrs().Index,
		}, netlink.RT_FILTER_OIF|netlink.RT_FILTER_PROTOCOL)
		if err != nil {
			return fmt.Errorf("netlink failed to list routes: %w", err)
		}
		ipv4DefaultRoute := net.IPNet{IP: net.IPv4zero, Mask: net.CIDRMask(0, 32)}
		for _, route := range dhcpRoutes {
			// Don't remove routes which can be atomically replaced by
			// RouteReplace to prevent potential traffic disruptions.
			if !isIPNetEqual(&ipv4DefaultRoute, route.Dst) && newRouter != nil {
				continue
			}
			err := netlink.RouteDel(&route)
			if !os.IsNotExist(err) && err != nil {
				return fmt.Errorf("failed to delete DHCP route: %w", err)
			}
		}

		if newRouter != nil {
			err := netlink.RouteReplace(&netlink.Route{
				Protocol:  unix.RTPROT_DHCP,
				Dst:       &ipv4DefaultRoute,
				Gw:        newRouter,
				Src:       new.AssignedIP,
				LinkIndex: iface.Attrs().Index,
				Scope:     netlink.SCOPE_UNIVERSE,
			})
			if err != nil {
				return fmt.Errorf("failed to add default route via %s: %w", newRouter, err)
			}
		}
		return nil
	}
}
