// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"context"
	"fmt"
	"time"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/supervisor"
)

func (s *Service) runLinkState(ctx context.Context) error {
	l := supervisor.Logger(ctx)
	linkUpdates := make(chan netlink.LinkUpdate, 10)
	options := netlink.LinkSubscribeOptions{
		ErrorCallback: func(err error) {
			l.Errorf("netlink subscription error: %v", err)
		},
	}
	if err := netlink.LinkSubscribeWithOptions(linkUpdates, ctx.Done(), options); err != nil {
		return fmt.Errorf("while subscribing to netlink link updates: %w", err)
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	lastIfState := make(map[string]bool)
	ifContexts := make(map[string]context.CancelFunc)
	for {
		select {
		case u, ok := <-linkUpdates:
			if !ok {
				return fmt.Errorf("link update channel closed")
			}
			attrs := u.Link.Attrs()
			if u.Link.Type() == "veth" {
				// Virtual links are not managed by this.
				continue
			}
			before := lastIfState[attrs.Name]
			now := attrs.RawFlags&unix.IFF_RUNNING != 0
			lastIfState[attrs.Name] = now

			if !before && now {
				ifCtx, cancel := context.WithCancel(ctx)
				ifContexts[attrs.Name] = cancel
				// Announces all configured IPs marked as permanent via ARP
				// every time an interface comes up.
				if err := sendARPAnnouncements(u.Link); err != nil {
					l.Warningf("Failed sending ARP announcements for interface %q: %v", attrs.Name, err)
				}
				// Send a second one after 10s if the network infrastructure is
				// slow to configure itself after link up or the first one got
				// lost.
				time.AfterFunc(10*time.Second, func() {
					if err := sendARPAnnouncements(u.Link); err != nil {
						l.Warningf("Failed sending repeated ARP announcements for interface %q: %v", attrs.Name, err)
					}
				})
				if err := runLLDP(ifCtx, netlinkLinkToNetInterface(u.Link)); err != nil {
					l.Warningf("Failed running LLDP for interface %q: %v")
				}
				l.Infof("Interface %q is up", attrs.Name)
			} else if before && !now {
				if cancel, ok := ifContexts[attrs.Name]; ok {
					cancel()
					delete(ifContexts, attrs.Name)
				}
				l.Infof("Interface %q is down", attrs.Name)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
