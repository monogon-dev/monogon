// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package network

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/mdlayher/lldp"
	"github.com/mdlayher/packet"

	"source.monogon.dev/metropolis/node/core/productinfo"
	"source.monogon.dev/osbase/supervisor"
)

var lldpTxAddr = packet.Addr{HardwareAddr: net.HardwareAddr{0x01, 0x80, 0xC2, 0x00, 0x00, 0x0E}}

// System capabilitiy bits. See 802.1AB-2016 Table 8-4.
const capRouter = 1 << 4

type caps struct {
	Supported uint16
	Enabled   uint16
}

// runLLDP transmits LLDP frames with standard timings advertising the
// Monogon node.
func runLLDP(ctx context.Context, iface *net.Interface) error {
	conn, err := packet.Listen(iface, packet.Datagram, int(lldp.EtherType), nil)
	if err != nil {
		return fmt.Errorf("while setting up LLDP listener: %w", err)
	}

	go func() {
		rxBuf := make([]byte, 9000)
		for {
			// Clear the RX queue, but do nothing with received frames.
			// These will need to be processed at some point.
			_, _, err := conn.ReadFrom(rxBuf)
			if err != nil {
				return
			}
		}
	}()
	go func() {
		for {
			var txFrame lldp.Frame
			txFrame.ChassisID = &lldp.ChassisID{
				// Using an interface MAC here is what most network gear does.
				// There isn't really anything better.
				Subtype: lldp.ChassisIDSubtypeMACAddress,
				ID:      []byte(iface.HardwareAddr),
			}
			txFrame.PortID = &lldp.PortID{
				// This is also a bit suboptimal, but getting better data requires
				// parsing and handling lots of VPD and DMI/SMBUS data.
				Subtype: lldp.PortIDSubtypeInterfaceAlias,
				ID:      []byte(iface.Name),
			}
			hostname, err := os.Hostname()
			if err != nil {
				// Should never happen, but if it does, we'll just use a generic string.
				hostname = "<unknown>"
			}
			txFrame.Optional = append(txFrame.Optional, &lldp.TLV{
				Type:   lldp.TLVTypeSystemName,
				Length: uint16(len(hostname)),
				Value:  []byte(hostname),
			})
			systemDesc := productinfo.Get().Info.Name + " " + productinfo.Get().VersionString
			txFrame.Optional = append(txFrame.Optional, &lldp.TLV{
				Type:   lldp.TLVTypeSystemDescription,
				Length: uint16(len(systemDesc)),
				Value:  []byte(systemDesc),
			})
			capsRaw, err := binary.Append(nil, binary.BigEndian, caps{
				Supported: capRouter,
				Enabled:   capRouter,
			})
			if err != nil {
				// All inputs hardcoded
				panic(err)
			}
			txFrame.Optional = append(txFrame.Optional, &lldp.TLV{
				Type:   lldp.TLVTypeSystemCapabilities,
				Length: uint16(len(capsRaw)),
				Value:  capsRaw,
			})

			// Use standard timings (30s interval / 120s TTL)
			txFrame.TTL = 120 * time.Second
			txBuf, err := txFrame.MarshalBinary()
			if err == nil {
				// Make sure this makes progress if an interface dies.
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				conn.WriteTo(txBuf, &lldpTxAddr)
			} else {
				supervisor.Logger(ctx).Warningf("Failed to marshal LLDP frame (interface %v): %v", iface.Name, err)
			}

			select {
			case <-time.After(30 * time.Second):
			case <-ctx.Done():
				// Do not send a zero-TTL clear frame as the context is only
				// canceled if the interface is already down, making it pointless.
				conn.Close()
				return
			}
		}
	}()
	return nil
}
