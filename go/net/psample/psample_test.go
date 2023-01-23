// This test requires the following Linux kernel configuration options to be
// set:
//
// CONFIG_NET_CLS_ACT
// CONFIG_NET_CLS_MATCHALL
// CONFIG_NET_SCHED
// CONFIG_NET_SCH_INGRESS
// CONFIG_PSAMPLE
// CONFIG_NET_ACT_SAMPLE
package psample

import (
	"errors"
	"net"
	"os"
	"strings"
	"syscall"
	"testing"

	"golang.org/x/sys/unix"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/vishvananda/netlink"
)

// setupLink adds and brings up a named link suited for packet capture.
func setupLink(t *testing.T, name string) netlink.Link {
	t.Helper()

	lnk := &netlink.Dummy{
		LinkAttrs: netlink.LinkAttrs{Name: name},
	}
	if err := netlink.LinkAdd(lnk); err != nil {
		t.Fatalf("while adding link: %v", err)
	}
	if err := netlink.LinkSetUp(lnk); err != nil {
		t.Fatalf("while setting up link: %v", err)
	}
	return lnk
}

// setupQdisc registers a clsact qdisc, which works on both the ingress and
// the egress. This is important as we'll be sampling packets leaving the test
// interface 'lk'. The qdisc is registered with the link lk.
// More on clsact: https://lwn.net/Articles/671458/
func setupQdisc(t *testing.T, lk netlink.Link) netlink.GenericQdisc {
	t.Helper()

	qdisc := netlink.GenericQdisc{
		QdiscAttrs: netlink.QdiscAttrs{
			LinkIndex: lk.Attrs().Index,
			Handle:    netlink.MakeHandle(0xffff, 0),
			Parent:    netlink.HANDLE_CLSACT,
		},
		QdiscType: "clsact",
	}
	if err := netlink.QdiscAdd(&qdisc); err != nil {
		t.Fatalf("while adding qdisc: %v", err)
	}
	return qdisc
}

// setupSamplingFilter adds a filter on 'lk' link that will sample packets
// exiting the interface.
func setupSamplingFilter(t *testing.T, lk netlink.Link) netlink.Filter {
	t.Helper()

	sa := netlink.NewSampleAction()
	// Sampled packets can be assigned their distinct group, allowing for
	// multiple flows to be sampled and analyzed separately at the same time.
	sa.Group = 7
	// Every 10th packet will be sampled.
	sa.Rate = 10
	// Packet samples will be truncated to sa.TruncSize.
	sa.TruncSize = 1500

	fcid := netlink.MakeHandle(1, 1)
	filter := &netlink.MatchAll{
		FilterAttrs: netlink.FilterAttrs{
			LinkIndex: lk.Attrs().Index,
			Parent:    netlink.HANDLE_MIN_EGRESS,
			Priority:  1,
			Protocol:  unix.ETH_P_ALL,
		},
		ClassId: fcid,
		Actions: []netlink.Action{
			sa,
		},
	}
	if err := netlink.FilterAdd(filter); err != nil {
		t.Fatalf("while adding filter: %v", err)
	}
	return filter
}

// packetAttrs contains the test attributes looked for in a packet sample.
type packetAttrs struct {
	// magic is the string expected to be found in the packet's application layer
	// contents.
	magic string
	// oifIdx identifies the egress interface the packet is exiting.
	oifIdx uint16
}

// match returns true if packet 'raw' matches attributes specified in 'pa'.
func (pa packetAttrs) match(t *testing.T, raw Packet) bool {
	t.Helper()

	// Check the packet's indicated egress interface.
	if raw.OutgoingInterfaceIndex != pa.oifIdx {
		return false
	}

	// Check the packet's payload.
	if raw.Data == nil {
		t.Fatalf("missing payload")
	}
	p := gopacket.NewPacket(raw.Data, layers.LayerTypeEthernet, gopacket.Default)
	if app := p.ApplicationLayer(); app != nil {
		if strings.Contains(string(app.Payload()), pa.magic) {
			return true
		}
	}
	return false
}

// TestSampling ascertains that packet samples can be obtained through use of
// this package's Subscribe(), and Receive().
func TestSampling(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	// Make sure 'psample' module is loaded.
	if _, err := netlink.GenlFamilyGet("psample"); err != nil {
		t.Fatalf("psample genetlink family unavailable - is CONFIG_PSAMPLE enabled?")
	}

	// Set up the test link/interface, and supply it with a network address,
	// which will enable routing for the test packets.
	var localA = netlink.Addr{
		IPNet: &net.IPNet{
			IP:   net.IPv4(10, 0, 0, 4),
			Mask: net.CIDRMask(24, 32),
		},
	}
	lk := setupLink(t, "if1")
	if err := netlink.AddrAdd(lk, &localA); err != nil {
		t.Fatalf("while adding network address: %v", err)
	}

	// Set up sampling.

	setupQdisc(t, lk)
	setupSamplingFilter(t, lk)

	c, err := Subscribe()
	if err != nil {
		t.Fatalf("while subscribing to psample notifications: %v", err)
	}

	// Test case: we'll send UDP datagrams to a remote address within the network
	// associated with link lk, expecting these packets to show up in the sampled
	// egress traffic.
	var dstA = netlink.Addr{
		IPNet: &net.IPNet{
			IP:   net.IPv4(10, 0, 0, 5),
			Mask: net.CIDRMask(24, 32),
		},
	}

	// The sampled packets are expected to:
	// - egress from interface 'if1'
	// - contain a magic payload
	sa := packetAttrs{
		magic:  "test datagram",
		oifIdx: uint16(lk.Attrs().Index),
	}

	// Look for packets matching attributes defined in 'sa'. Signal on 'dC'
	// immediately after the expected packet has been received, then return.
	dC := make(chan struct{})
	go func() {
		for {
			pkts, err := Receive(c)
			// Receiving ENOBUFS is expected in this case. It signals that some of
			// the sampled traffic could not have been captured, and had been
			// dropped instead.
			if err != nil && !errors.Is(err, syscall.ENOBUFS) {
				t.Fatalf("while receiving psamples: %v", err)
			}
			for _, raw := range pkts {
				if sa.match(t, raw) {
					t.Logf("Got the expected packet sample.")
					dC <- struct{}{}
					return
				}
			}
		}
	}()

	// Send out the test datagrams. Return as soon as the test succeeds.

	dstE := net.JoinHostPort(dstA.IP.String(), "1234")
	conn, err := net.Dial("udp", dstE)
	if err != nil {
		t.Fatalf("while dialing UDP address: %v", err)
	}
	defer conn.Close()

	for {
		select {
		case <-dC:
			return
		default:
			conn.Write([]byte(sa.magic))
		}
	}
}
