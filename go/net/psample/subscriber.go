// Package psample provides a receiver for sampled network packets using the
// Netlink psample interface.
package psample

import (
	"fmt"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
)

// attrId identifies psample netlink message attributes.
// Identifier numbers are based on psample kernel module sources:
// https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/tree/include/uapi/linux/psample.h?h=v5.15.89#n5
type attrId uint16

const (
	aIIfIndex      attrId = iota // u16
	aOIfIndex                    // u16
	aOrigSize                    // u32
	aSampleGroup                 // u32
	aGroupSeq                    // u32
	aSampleRate                  // u32
	aData                        // []byte
	aGroupRefcount               // u32
	aTunnel

	aPad
	aOutTC     // u16
	aOutTCOCC  // u64
	aLatency   // u64, nanoseconds
	aTimestamp // u64, nanoseconds
	aProto     // u16
)

// Packet contains the sampled packet in its raw form, along with its
// 'psample' metadata.
type Packet struct {
	// IncomingInterfaceIndex is the incoming interface index of the packet or 0
	// if not applicable.
	IncomingInterfaceIndex uint16
	// OutgoingInterfaceIndex is the outgoing interface index of the packet or 0
	// if not applicable.
	OutgoingInterfaceIndex uint16
	// OriginalSize is the packet's original size in bytes without any
	// truncation.
	OriginalSize uint32
	// SampleGroup is the sample group to which this packet belongs. This is set
	// by the sampling action and can be used to differentiate different
	// sampling streams.
	SampleGroup uint32
	// GroupSequence is a monotonically-increasing counter of packets sampled
	// for each sample group.
	GroupSequence uint32
	// SampleRate is the sampling rate (1 in SampleRate packets) used to capture
	// this packet.
	SampleRate uint32
	// Data contains the packet data up to the specified size for truncation.
	Data []byte

	// The following attributes are only available on kernel versions 5.13+

	// Latency is the sampled packet's latency as indicated by psample. It's
	// expressed in nanoseconds.
	Latency uint64
	// Timestamp marks time of the packet's sampling. It's set by the kernel, and
	// expressed in Unix nanoseconds.
	Timestamp uint64
}

// decode converts raw generic netlink message attributes into a Packet. In
// cases where some of the known psample attributes were left unspecified in
// the message, appropriate Packet member variables will be left with their
// zero values.
func decode(b []byte) (*Packet, error) {
	ad, err := netlink.NewAttributeDecoder(b)
	if err != nil {
		return nil, err
	}

	var p Packet
	for ad.Next() {
		switch attrId(ad.Type()) {
		case aIIfIndex:
			p.IncomingInterfaceIndex = ad.Uint16()
		case aOIfIndex:
			p.OutgoingInterfaceIndex = ad.Uint16()
		case aOrigSize:
			p.OriginalSize = ad.Uint32()
		case aSampleGroup:
			p.SampleGroup = ad.Uint32()
		case aGroupSeq:
			p.GroupSequence = ad.Uint32()
		case aSampleRate:
			p.SampleRate = ad.Uint32()
		case aData:
			p.Data = ad.Bytes()
		case aLatency:
			p.Latency = ad.Uint64()
		case aTimestamp:
			p.Timestamp = ad.Uint64()
		default:
		}
	}
	return &p, nil
}

// Subscribe returns a NetlinkSocket that's already subscribed to "packets"
// psample multicast group, which makes it ready to receive packet samples.
// Close should be called on the returned socket.
func Subscribe() (*genetlink.Conn, error) {
	// Create a netlink socket.
	c, err := genetlink.Dial(nil)
	if err != nil {
		return nil, fmt.Errorf("while dialing netlink socket: %w", err)
	}

	// Lookup the netlink family id associated with psample kernel module.
	f, err := c.GetFamily("psample")
	if err != nil {
		c.Close()
		return nil, fmt.Errorf("couldn't lookup \"psample\" netlink family: %w", err)
	}

	// Lookup psample's packet sampling netlink multicast group.
	var pktGrpId uint32
	for _, mgrp := range f.Groups {
		if mgrp.Name == "packets" {
			pktGrpId = mgrp.ID
			break
		}
	}
	if pktGrpId == 0 {
		c.Close()
		return nil, fmt.Errorf("packets multicast group not found")
	}

	// Subscribe to 'packets' multicast group in order to receive packet
	// samples.
	if err := c.JoinGroup(pktGrpId); err != nil {
		c.Close()
		return nil, fmt.Errorf("couldn't join multicast group: %w", err)
	}
	return c, nil
}

// Receive returns one or more of the sampled packets as soon as they're
// available. It may return a syscall.ENOBUFS error which indicates that the
// kernel-side buffer of the netlink connection has overflowed and lost
// packets. This is a transient error, calling Receive again will retrieve
// future packet samples.
func Receive(c *genetlink.Conn) ([]Packet, error) {
	// Wait for the samples to arrive over generic netlink connection c.
	gnms, nms, err := c.Receive()
	if err != nil {
		return nil, fmt.Errorf("while receiving netlink notifications: %w", err)
	}

	var pkts []Packet
	for i := 0; i < len(nms); i++ {
		// Only process multicast notifications.
		if nms[i].Header.PID != 0 {
			continue
		}

		// PSAMPLE_CMD_SAMPLE should be zero in multicast notifications.
		if gnms[i].Header.Command != 0 {
			continue
		}

		// Iterate over the Generic Netlink attributes present in the message,
		// extracting any relating to the sampled packet.
		pkt, err := decode(gnms[i].Data)
		if err != nil {
			return nil, fmt.Errorf("while decoding netlink notification: %w", err)
		}
		pkts = append(pkts, *pkt)
	}
	return pkts, nil
}
