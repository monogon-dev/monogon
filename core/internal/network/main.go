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

package network

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/google/nftables"
	"github.com/google/nftables/expr"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/network/dhcp"
	"git.monogon.dev/source/nexantic.git/core/internal/network/dns"
	"git.monogon.dev/source/nexantic.git/core/pkg/logtree"
)

const (
	resolvConfPath     = "/etc/resolv.conf"
	resolvConfSwapPath = "/etc/resolv.conf.new"
)

type Service struct {
	config Config
	dhcp   *dhcp.Client

	logger logtree.LeveledLogger
}

type Config struct {
	CorednsRegistrationChan chan *dns.ExtraDirective
}

func New(config Config) *Service {
	return &Service{
		config: config,
		dhcp:   dhcp.New(),
	}
}

func setResolvconf(nameservers []net.IP, searchDomains []string) error {
	_ = os.Mkdir("/etc", 0755)
	newResolvConf, err := os.Create(resolvConfSwapPath)
	if err != nil {
		return err
	}
	defer newResolvConf.Close()
	defer os.Remove(resolvConfSwapPath)
	for _, ns := range nameservers {
		if _, err := newResolvConf.WriteString(fmt.Sprintf("nameserver %v\n", ns)); err != nil {
			return err
		}
	}
	for _, searchDomain := range searchDomains {
		if _, err := newResolvConf.WriteString(fmt.Sprintf("search %v", searchDomain)); err != nil {
			return err
		}
	}
	newResolvConf.Close()
	// Atomically swap in new config
	return unix.Rename(resolvConfSwapPath, resolvConfPath)
}

func (s *Service) addNetworkRoutes(link netlink.Link, addr net.IPNet, gw net.IP) error {
	if err := netlink.AddrReplace(link, &netlink.Addr{IPNet: &addr}); err != nil {
		return fmt.Errorf("failed to add DHCP address to network interface \"%v\": %w", link.Attrs().Name, err)
	}

	if gw.IsUnspecified() {
		s.logger.Infof("No default route set, only local network %s will be reachable", addr.String())
		return nil
	}

	route := &netlink.Route{
		Dst:   &net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPv4Mask(0, 0, 0, 0)},
		Gw:    gw,
		Scope: netlink.SCOPE_UNIVERSE,
	}
	if err := netlink.RouteAdd(route); err != nil {
		return fmt.Errorf("could not add default route: netlink.RouteAdd(%+v): %v", route, err)
	}
	return nil
}

// nfifname converts an interface name into 16 bytes padded with zeroes (for nftables)
func nfifname(n string) []byte {
	b := make([]byte, 16)
	copy(b, []byte(n+"\x00"))
	return b
}

func (s *Service) useInterface(ctx context.Context, iface netlink.Link) error {
	err := supervisor.Run(ctx, "dhcp", s.dhcp.Run(iface))
	if err != nil {
		return err
	}
	status, err := s.dhcp.Status(ctx, true)
	if err != nil {
		return fmt.Errorf("could not get DHCP Status: %w", err)
	}

	// We're currently never removing this directive just like we're not removing routes and IPs
	s.config.CorednsRegistrationChan <- dns.NewUpstreamDirective(status.DNS)

	if err := s.addNetworkRoutes(iface, status.Address, status.Gateway); err != nil {
		s.logger.Warning("Failed to add routes: %v", err)
	}

	c := nftables.Conn{}

	nat := c.AddTable(&nftables.Table{
		Family: nftables.TableFamilyIPv4,
		Name:   "nat",
	})

	postrouting := c.AddChain(&nftables.Chain{
		Name:     "postrouting",
		Hooknum:  nftables.ChainHookPostrouting,
		Priority: nftables.ChainPriorityNATSource,
		Table:    nat,
		Type:     nftables.ChainTypeNAT,
	})

	// Masquerade/SNAT all traffic going out of the external interface
	c.AddRule(&nftables.Rule{
		Table: nat,
		Chain: postrouting,
		Exprs: []expr.Any{
			&expr.Meta{Key: expr.MetaKeyOIFNAME, Register: 1},
			&expr.Cmp{
				Op:       expr.CmpOpEq,
				Register: 1,
				Data:     nfifname(iface.Attrs().Name),
			},
			&expr.Masq{},
		},
	})

	if err := c.Flush(); err != nil {
		panic(err)
	}

	return nil
}

// GetIP returns the current IP (and optionally waits for one to be assigned)
func (s *Service) GetIP(ctx context.Context, wait bool) (*net.IP, error) {
	status, err := s.dhcp.Status(ctx, wait)
	if err != nil {
		return nil, err
	}
	return &status.Address.IP, nil
}

func (s *Service) Run(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	dnsSvc := dns.New(s.config.CorednsRegistrationChan)
	supervisor.Run(ctx, "dns", dnsSvc.Run)
	supervisor.Run(ctx, "interfaces", s.runInterfaces)

	if err := ioutil.WriteFile("/proc/sys/net/ipv4/ip_forward", []byte("1\n"), 0644); err != nil {
		logger.Fatalf("Failed to enable IPv4 forwarding: %v", err)
	}

	// We're handling all DNS requests with CoreDNS, including local ones
	if err := setResolvconf([]net.IP{{127, 0, 0, 1}}, []string{}); err != nil {
		logger.Fatalf("Failed to set resolv.conf: %v", err)
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}

func (s *Service) runInterfaces(ctx context.Context) error {
	s.logger = supervisor.Logger(ctx)
	s.logger.Info("Starting network interface management")

	links, err := netlink.LinkList()
	if err != nil {
		s.logger.Fatalf("Failed to list network links: %s", err)
	}

	var ethernetLinks []netlink.Link
	for _, link := range links {
		attrs := link.Attrs()
		if link.Type() == "device" && len(attrs.HardwareAddr) > 0 {
			if len(attrs.HardwareAddr) == 6 { // Ethernet
				if attrs.Flags&net.FlagUp != net.FlagUp {
					netlink.LinkSetUp(link) // Attempt to take up all ethernet links
				}
				ethernetLinks = append(ethernetLinks, link)
			} else {
				s.logger.Infof("Ignoring non-Ethernet interface %s", attrs.Name)
			}
		} else if link.Attrs().Name == "lo" {
			if err := netlink.LinkSetUp(link); err != nil {
				s.logger.Errorf("Failed to bring up loopback interface: %v", err)
			}
		}
	}
	if len(ethernetLinks) != 1 {
		s.logger.Warningf("Network service needs exactly one link, bailing")
	} else {
		link := ethernetLinks[0]
		if err := s.useInterface(ctx, link); err != nil {
			return fmt.Errorf("failed to bring up link %s: %w", link.Attrs().Name, err)
		}
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
