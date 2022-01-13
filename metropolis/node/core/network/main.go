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
	"net"
	"os"
	"path"
	"strings"

	"github.com/google/nftables"
	"github.com/google/nftables/expr"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/vishvananda/netlink"

	"source.monogon.dev/metropolis/node/core/network/dhcp4c"
	dhcpcb "source.monogon.dev/metropolis/node/core/network/dhcp4c/callback"
	"source.monogon.dev/metropolis/node/core/network/dns"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// Service is the network service for this node. It maintains all
// networking-related functionality, but is generally not aware of the inner
// workings of Metropolis, instead functioning in a generic manner. Once created
// via New, it can be started and restarted arbitrarily, but the service object
// itself must be long-lived.
type Service struct {
	dnsReg chan *dns.ExtraDirective
	dnsSvc *dns.Service

	// dhcp client for the 'main' interface of the node.
	dhcp *dhcp4c.Client

	// nftConn is a shared file descriptor handle to nftables, automatically
	// initialized on first use.
	nftConn             nftables.Conn
	natTable            *nftables.Table
	natPostroutingChain *nftables.Chain

	status memory.Value
}

func New() *Service {
	dnsReg := make(chan *dns.ExtraDirective)
	dnsSvc := dns.New(dnsReg)
	return &Service{
		dnsReg: dnsReg,
		dnsSvc: dnsSvc,
	}
}

// Status is the current network status of the host. It will be updated by the
// network Service whenever the node's network configuration changes. Spurious
// changes might occur, consumers should ensure that the change that occured is
// meaningful to them.
type Status struct {
	ExternalAddress net.IP
	DNSServers      dhcp4c.DNSServers
}

// Watcher allows network Service consumers to watch for updates of the current
// Status.
type Watcher struct {
	watcher event.Watcher
}

// Get returns the newest network Status from a Watcher. It will block until a
// new Status is available.
func (w *Watcher) Get(ctx context.Context) (*Status, error) {
	val, err := w.watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	status := val.(Status)
	return &status, err
}

func (w *Watcher) Close() error {
	return w.watcher.Close()
}

// Watch returns a Watcher, which can be used by consumers of the network
// Service to retrieve the current network status.
// Close must be called on the Watcher when it is not used anymore in order to
// prevent goroutine leaks.
func (s *Service) Watch() Watcher {
	return Watcher{
		watcher: s.status.Watch(),
	}
}

// ConfigureDNS sets a DNS ExtraDirective on the built-in DNS server of the
// network Service. See //metropolis/node/core/network/dns for more
// information.
func (s *Service) ConfigureDNS(d *dns.ExtraDirective) {
	s.dnsReg <- d
}

// nfifname converts an interface name into 16 bytes padded with zeroes (for
// nftables)
func nfifname(n string) []byte {
	b := make([]byte, 16)
	copy(b, []byte(n+"\x00"))
	return b
}

// statusCallback is the main DHCP client callback connecting updates to the
// current lease to the rest of Metropolis. It updates the DNS service's
// configuration to use the received upstream servers, and notifies the rest of
// Metropolis via en event value that the network configuration has changed.
func (s *Service) statusCallback(old, new *dhcp4c.Lease) error {
	// Reconfigure DNS if needed.
	oldServers := old.DNSServers()
	newServers := new.DNSServers()
	if !newServers.Equal(oldServers) {
		s.ConfigureDNS(dns.NewUpstreamDirective(newServers))
	}
	// Notify status waiters.
	s.status.Set(Status{
		ExternalAddress: new.AssignedIP,
		DNSServers:      new.DNSServers(),
	})
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
	s.dhcp.VendorClassIdentifier = "dev.monogon.metropolis.node.v1"
	s.dhcp.RequestedOptions = []dhcpv4.OptionCode{dhcpv4.OptionRouter, dhcpv4.OptionDomainNameServer, dhcpv4.OptionClasslessStaticRoute}
	s.dhcp.LeaseCallback = dhcpcb.Compose(dhcpcb.ManageIP(iface), dhcpcb.ManageRoutes(iface), s.statusCallback)
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

// sysctlOptions contains sysctl options to apply
type sysctlOptions map[string]string

// apply attempts to apply all options in sysctlOptions. It aborts on the first
// one which returns an error when applying.
func (o sysctlOptions) apply() error {
	for name, value := range o {
		filePath := path.Join("/proc/sys/", strings.ReplaceAll(name, ".", "/"))
		optionFile, err := os.OpenFile(filePath, os.O_WRONLY, 0)
		if err != nil {
			return fmt.Errorf("failed to set option %v: %w", name, err)
		}
		if _, err := optionFile.WriteString(value + "\n"); err != nil {
			optionFile.Close()
			return fmt.Errorf("failed to set option %v: %w", name, err)
		}
		optionFile.Close() // In a loop, defer'ing could open a lot of FDs
	}
	return nil
}

func (s *Service) Run(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	supervisor.Run(ctx, "dns", s.dnsSvc.Run)
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

	sysctlOpts := sysctlOptions{
		// Enable IP forwarding for our pods
		"net.ipv4.ip_forward": "1",
		// Enable strict reverse path filtering on all interfaces (important
		// for spoofing prevention from Pods with CAP_NET_ADMIN)
		"net.ipv4.conf.all.rp_filter": "1",
		// Disable source routing
		"net.ipv4.conf.all.accept_source_route": "0",

		// Increase Linux socket kernel buffer sizes to 16MiB (needed for fast
		// datacenter networks)
		"net.core.rmem_max": "16777216",
		"net.core.wmem_max": "16777216",
		"net.ipv4.tcp_rmem": "4096 87380 16777216",
		"net.ipv4.tcp_wmem": "4096 87380 16777216",
	}
	if err := sysctlOpts.apply(); err != nil {
		logger.Fatalf("Failed to set up kernel network config: %v", err)
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}

func (s *Service) runInterfaces(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	logger.Info("Starting network interface management")

	links, err := netlink.LinkList()
	if err != nil {
		logger.Fatalf("Failed to list network links: %s", err)
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
				logger.Infof("Ignoring non-Ethernet interface %s", attrs.Name)
			}
		} else if link.Attrs().Name == "lo" {
			if err := netlink.LinkSetUp(link); err != nil {
				logger.Errorf("Failed to bring up loopback interface: %v", err)
			}
		}
	}
	if len(ethernetLinks) != 1 {
		logger.Warningf("Network service needs exactly one link, bailing")
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
