// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"context"
	"fmt"
	"net"
	"time"
)

// NetworkStatus is example data that will be stored in a Value.
type NetworkStatus struct {
	ExternalAddress net.IP
	DefaultGateway  net.IP
}

// NetworkService is a fake/example network service that is responsible for
// communicating the newest information about a machine's network configuration
// to consumers/watchers.
type NetworkService struct {
	Provider Value[NetworkStatus]
}

// Run pretends to execute the network service's main logic loop, in which it
// pretends to have received an IP address over DHCP, and communicates that to
// consumers/watchers.
func (s *NetworkService) Run(ctx context.Context) {
	s.Provider.Set(NetworkStatus{
		ExternalAddress: nil,
		DefaultGateway:  nil,
	})

	select {
	case <-time.After(100 * time.Millisecond):
	case <-ctx.Done():
		return
	}

	fmt.Printf("NS: Got DHCP Lease\n")
	s.Provider.Set(NetworkStatus{
		ExternalAddress: net.ParseIP("203.0.113.24"),
		DefaultGateway:  net.ParseIP("203.0.113.1"),
	})

	select {
	case <-time.After(100 * time.Millisecond):
	case <-ctx.Done():
		return
	}

	fmt.Printf("NS: DHCP Address changed\n")
	s.Provider.Set(NetworkStatus{
		ExternalAddress: net.ParseIP("203.0.113.103"),
		DefaultGateway:  net.ParseIP("203.0.113.1"),
	})

	time.Sleep(100 * time.Millisecond)
}

// ExampleValue_full demonstrates a typical usecase for Event Values, in which
// a mock network service lets watchers know that the machine on which the code
// is running has received a new network configuration.
// It also shows the typical boilerplate required in order to wrap a Value (eg.
// MemoryValue) within a typesafe wrapper.
func ExampleValue_full() {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Create a fake NetworkService.
	var ns NetworkService

	// Run an /etc/hosts updater. It will watch for updates from the NetworkService
	// about the current IP address of the node.
	go func() {
		w := ns.Provider.Watch()
		for {
			status, err := w.Get(ctx)
			if err != nil {
				break
			}
			if status.ExternalAddress == nil {
				continue
			}
			// Pretend to write /etc/hosts with the newest ExternalAddress.
			// In production code, you would also check for whether ExternalAddress has
			// changed from the last written value, if writing to /etc/hosts is expensive.
			fmt.Printf("/etc/hosts: foo.example.com is now %s\n", status.ExternalAddress.String())
		}
	}()

	// Run fake network service.
	ns.Run(ctx)

	// Output:
	// NS: Got DHCP Lease
	// /etc/hosts: foo.example.com is now 203.0.113.24
	// NS: DHCP Address changed
	// /etc/hosts: foo.example.com is now 203.0.113.103
}
