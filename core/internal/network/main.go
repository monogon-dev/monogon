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

	"git.monogon.dev/source/nexantic.git/core/internal/common/service"

	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
)

const (
	resolvConfPath     = "/etc/resolv.conf"
	resolvConfSwapPath = "/etc/resolv.conf.new"
)

type Service struct {
	*service.BaseService
	config      Config
	dhcp4Client *nclient4.Client
}

type Config struct {
}

func NewNetworkService(config Config, logger *zap.Logger) (*Service, error) {
	s := &Service{
		config: config,
	}
	s.BaseService = service.NewBaseService("network", logger, s)
	return s, nil
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

func addNetworkRoutes(link netlink.Link, addr net.IPNet, gw net.IP) error {
	if err := netlink.AddrReplace(link, &netlink.Addr{IPNet: &addr}); err != nil {
		return err
	}
	if err := netlink.RouteAdd(&netlink.Route{
		Dst:   &net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPv4Mask(0, 0, 0, 0)},
		Gw:    gw,
		Scope: netlink.SCOPE_UNIVERSE,
	}); err != nil {
		return fmt.Errorf("Failed to add default route: %w", err)
	}
	return nil
}

const (
	stateInitialize = 1
	stateSelect     = 2
	stateBound      = 3
	stateRenew      = 4
	stateRebind     = 5
)

var dhcpBroadcastAddr = &net.UDPAddr{IP: net.IP{255, 255, 255, 255}, Port: 67}

// TODO(lorenz): This is a super terrible DHCP client, but it works for QEMU slirp
func (s *Service) dhcpClient(iface netlink.Link) error {
	client, err := nclient4.New(iface.Attrs().Name)
	if err != nil {
		panic(err)
	}
	_, ack, err := client.Request(context.Background())
	if err != nil {
		panic(err)
	}
	s.Logger.Info("Network service got IP", zap.String("ip", ack.YourIPAddr.String()))
	if err := setResolvconf(ack.DNS(), []string{}); err != nil {
		s.Logger.Warn("Failed to set resolvconf", zap.Error(err))
	}
	if err := addNetworkRoutes(iface, net.IPNet{IP: ack.YourIPAddr, Mask: ack.SubnetMask()}, ack.GatewayIPAddr); err != nil {
		s.Logger.Warn("Failed to add routes", zap.Error(err))
	}
	return nil
}

func (s *Service) OnStart() error {
	s.Logger.Info("Starting network service")
	links, err := netlink.LinkList()
	if err != nil {
		s.Logger.Fatal("Failed to list network links", zap.Error(err))
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
				s.Logger.Info("Ignoring non-Ethernet interface", zap.String("interface", attrs.Name))
			}
		} else if link.Attrs().Name == "lo" {
			if err := netlink.LinkSetUp(link); err != nil {
				s.Logger.Error("Failed to take up loopback interface", zap.Error(err))
			}
		}
	}
	if len(ethernetLinks) == 1 {
		link := ethernetLinks[0]
		go s.dhcpClient(link)

	} else {
		s.Logger.Warn("Network service cannot yet handle more than one interface :(")
	}
	return nil
}

func (s *Service) OnStop() error {
	return nil
}
