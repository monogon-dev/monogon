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
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"
	"time"

	"github.com/google/nftables"
	"github.com/google/nftables/expr"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/network/dns"
	"git.monogon.dev/source/nexantic.git/core/pkg/dhcp4c"
	dhcpcb "git.monogon.dev/source/nexantic.git/core/pkg/dhcp4c/callback"
	"git.monogon.dev/source/nexantic.git/core/pkg/logtree"
)

const (
	resolvConfPath     = "/etc/resolv.conf"
	resolvConfSwapPath = "/etc/resolv.conf.new"
)

type Service struct {
	config Config
	dhcp   *dhcp4c.Client

	// nftConn is a shared file descriptor handle to nftables, automatically initialized on first use.
	nftConn             nftables.Conn
	natTable            *nftables.Table
	natPostroutingChain *nftables.Chain

	// These are a temporary hack pending the removal of the GetIP interface
	ipLock       sync.Mutex
	currentIPTmp net.IP

	logger logtree.LeveledLogger
}

type Config struct {
	CorednsRegistrationChan chan *dns.ExtraDirective
}

func New(config Config) *Service {
	return &Service{
		config: config,
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

// nfifname converts an interface name into 16 bytes padded with zeroes (for nftables)
func nfifname(n string) []byte {
	b := make([]byte, 16)
	copy(b, []byte(n+"\x00"))
	return b
}

func (s *Service) dhcpDNSCallback(old, new *dhcp4c.Lease) error {
	oldServers := old.DNSServers()
	newServers := new.DNSServers()
	if newServers.Equal(oldServers) {
		return nil // nothing to do
	}
	s.logger.Infof("Setting upstream DNS servers to %v", newServers)
	s.config.CorednsRegistrationChan <- dns.NewUpstreamDirective(newServers)
	return nil
}

// TODO(lorenz): Get rid of this once we have robust node resolution
func (s *Service) getIPCallbackHack(old, new *dhcp4c.Lease) error {
	if old == nil && new != nil {
		s.ipLock.Lock()
		s.currentIPTmp = new.AssignedIP
		s.ipLock.Unlock()
	}
	return nil
}

func (s *Service) useInterface(ctx context.Context, iface netlink.Link) error {
	netIface, err := net.InterfaceByIndex(iface.Attrs().Index)
	if err != nil {
		return fmt.Errorf("cannot create Go net.Interface from netlink.Link: %w", err)
	}
	s.dhcp, err = dhcp4c.NewClient(netIface)
	if err != nil {
		return fmt.Errorf("failed to create DHCP client on interface %v: %w", iface.Attrs().Name, err)
	}
	s.dhcp.VendorClassIdentifier = "com.nexantic.smalltown.v1"
	s.dhcp.RequestedOptions = []dhcpv4.OptionCode{dhcpv4.OptionRouter, dhcpv4.OptionNameServer}
	s.dhcp.LeaseCallback = dhcpcb.Compose(dhcpcb.ManageIP(iface), dhcpcb.ManageDefaultRoute(iface), s.dhcpDNSCallback, s.getIPCallbackHack)
	err = supervisor.Run(ctx, "dhcp", s.dhcp.Run)
	if err != nil {
		return err
	}

	// Masquerade/SNAT all traffic going out of the external interface
	s.nftConn.AddRule(&nftables.Rule{
		Table: s.natTable,
		Chain: s.natPostroutingChain,
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

	if err := s.nftConn.Flush(); err != nil {
		panic(err)
	}

	return nil
}

// GetIP returns the current IP (and optionally waits for one to be assigned)
func (s *Service) GetIP(ctx context.Context, wait bool) (*net.IP, error) {
	for {
		var currentIP net.IP
		s.ipLock.Lock()
		currentIP = s.currentIPTmp
		s.ipLock.Unlock()
		if currentIP == nil {
			if !wait {
				return nil, errors.New("no IP available")
			}
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(1 * time.Second):
				continue
			}
		}
		return &currentIP, nil
	}
}

func (s *Service) Run(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	dnsSvc := dns.New(s.config.CorednsRegistrationChan)
	supervisor.Run(ctx, "dns", dnsSvc.Run)
	supervisor.Run(ctx, "interfaces", s.runInterfaces)

	s.natTable = s.nftConn.AddTable(&nftables.Table{
		Family: nftables.TableFamilyIPv4,
		Name:   "nat",
	})

	s.natPostroutingChain = s.nftConn.AddChain(&nftables.Chain{
		Name:     "postrouting",
		Hooknum:  nftables.ChainHookPostrouting,
		Priority: nftables.ChainPriorityNATSource,
		Table:    s.natTable,
		Type:     nftables.ChainTypeNAT,
	})
	if err := s.nftConn.Flush(); err != nil {
		logger.Fatalf("Failed to set up nftables base chains: %v", err)
	}

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
