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

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
)

type dhcpClient struct {
	reqC   chan *dhcpStatusReq
	logger *zap.Logger
}

func newDHCPClient(logger *zap.Logger) *dhcpClient {
	return &dhcpClient{
		logger: logger,
		reqC:   make(chan *dhcpStatusReq),
	}
}

type dhcpStatusReq struct {
	resC chan *dhcpv4.DHCPv4
	wait bool
}

func (r *dhcpStatusReq) fulfill(p *dhcpv4.DHCPv4) {
	go func() {
		r.resC <- p
	}()
}

func (c *dhcpClient) run(ctx context.Context, iface netlink.Link) {
	// Channel updated with address once one gets assigned/updated
	newC := make(chan *dhcpv4.DHCPv4)
	// Status requests waiting for configuration
	waiters := []*dhcpStatusReq{}

	// Start lease acquisition
	// TODO(q3k): actually maintain the lease instead of hoping we never get
	// kicked off.
	client, err := nclient4.New(iface.Attrs().Name)
	if err != nil {
		panic(err)
	}
	go func() {
		_, ack, err := client.Request(ctx)
		if err != nil {
			c.logger.Error("DHCP lease request failed", zap.Error(err))
			// TODO(q3k): implement retry logic with full state machine
		}
		newC <- ack
	}()

	// State machine
	// Two implicit states: WAITING -> ASSIGNED
	// We start at WAITING, once we get a current config we move to ASSIGNED
	// Once this becomes more complex (ie. has to handle link state changes)
	// this should grow into a real state machine.
	var current *dhcpv4.DHCPv4
	c.logger.Info("DHCP client WAITING")
	for {
		select {
		case <-ctx.Done():
			// TODO(q3k): don't leave waiters hanging
			return

		case cfg := <-newC:
			current = cfg
			c.logger.Info("DHCP client ASSIGNED", zap.String("ip", current.String()))
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

func (c *dhcpClient) status(ctx context.Context, wait bool) (*dhcpv4.DHCPv4, error) {
	resC := make(chan *dhcpv4.DHCPv4)
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
