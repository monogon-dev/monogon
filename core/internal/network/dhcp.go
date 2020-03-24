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

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
)

type dhcpClient struct {
	reqC chan *dhcpStatusReq
}

func newDHCPClient() *dhcpClient {
	return &dhcpClient{
		reqC: make(chan *dhcpStatusReq),
	}
}

type dhcpStatusReq struct {
	resC chan *dhcpStatus
	wait bool
}

func (r *dhcpStatusReq) fulfill(s *dhcpStatus) {
	go func() {
		r.resC <- s
	}()
}

// dhcpStatus is the IPv4 configuration provisioned via DHCP for a given interface. It does not necessarily represent
// a configuration that is active or even valid.
type dhcpStatus struct {
	// address is 'our' (the node's) IPv4 address on the network.
	address net.IPNet
	// gateway is the default gateway/router of this network, or 0.0.0.0 if none was given.
	gateway net.IP
	// dns is a list of IPv4 DNS servers to use.
	dns []net.IP
}

func (s *dhcpStatus) String() string {
	return fmt.Sprintf("Address: %s, Gateway: %s, DNS: %v", s.address.String(), s.gateway.String(), s.dns)
}

func (c *dhcpClient) run(iface netlink.Link) supervisor.Runnable {
	return func(ctx context.Context) error {
		logger := supervisor.Logger(ctx)

		// Channel updated with address once one gets assigned/updated
		newC := make(chan *dhcpStatus)
		// Status requests waiting for configuration
		waiters := []*dhcpStatusReq{}

		// Start lease acquisition
		// TODO(q3k): actually maintain the lease instead of hoping we never get
		// kicked off.

		client, err := nclient4.New(iface.Attrs().Name)
		if err != nil {
			return fmt.Errorf("nclient4.New: %w", err)
		}

		err = supervisor.Run(ctx, "client", func(ctx context.Context) error {
			supervisor.Signal(ctx, supervisor.SignalHealthy)
			_, ack, err := client.Request(ctx)
			if err != nil {
				// TODO(q3k): implement retry logic with full state machine
				logger.Error("DHCP lease request failed", zap.Error(err))
				return err
			}
			newC <- parseAck(ack)
			supervisor.Signal(ctx, supervisor.SignalDone)
			return nil
		})
		if err != nil {
			return err
		}

		supervisor.Signal(ctx, supervisor.SignalHealthy)

		// State machine
		// Two implicit states: WAITING -> ASSIGNED
		// We start at WAITING, once we get a current config we move to ASSIGNED
		// Once this becomes more complex (ie. has to handle link state changes)
		// this should grow into a real state machine.
		var current *dhcpStatus
		logger.Info("DHCP client WAITING")
		for {
			select {
			case <-ctx.Done():
				// TODO(q3k): don't leave waiters hanging
				return err

			case cfg := <-newC:
				current = cfg
				logger.Info("DHCP client ASSIGNED", zap.String("ip", current.String()))
				for _, w := range waiters {
					w.fulfill(current)
				}
				waiters = []*dhcpStatusReq{}

			case r := <-c.reqC:
				if current != nil || !r.wait {
					r.fulfill(current)
				} else {
					waiters = append(waiters, r)
				}
			}
		}
	}
}

// parseAck turns an internal status (from the dhcpv4 library) into a dhcpStatus
func parseAck(ack *dhcpv4.DHCPv4) *dhcpStatus {
	address := net.IPNet{IP: ack.YourIPAddr, Mask: ack.SubnetMask()}

	// DHCP routers are optional - if none are provided, assume no router and set gateway to 0.0.0.0
	// (this makes gateway.IsUnspecified() == true)
	gateway, _, _ := net.ParseCIDR("0.0.0.0/0")
	if routers := ack.Router(); len(routers) > 0 {
		gateway = routers[0]
	}
	return &dhcpStatus{
		address: address,
		gateway: gateway,
		dns:     ack.DNS(),
	}
}

// status returns the DHCP configuration requested from us by the local DHCP server.
// If wait is true, this function will block until a DHCP configuration is available. Otherwise, a nil status may be
// returned to indicate that no configuration has been received yet.
func (c *dhcpClient) status(ctx context.Context, wait bool) (*dhcpStatus, error) {
	resC := make(chan *dhcpStatus)
	c.reqC <- &dhcpStatusReq{
		resC: resC,
		wait: wait,
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r := <-resC:
		return r, nil
	}
}
