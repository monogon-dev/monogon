// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package lacp contains an integration test for our custom LACP patches.
// It tests relevant behavior that other parts of the Monogon network stack
// rely on, like proper carrier state indications.
package lacp

import (
	"hash/fnv"
	"math"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

// createVethPair is a helper creating a pair of virtual network interfaces
// acting as a cable (i.e. traffic going in one interface comes back out on
// the other one). Both ends are returned.
func createVethPair(t *testing.T, name string) (*netlink.Veth, *netlink.Veth) {
	t.Helper()
	vethLink := netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name:    name + "a",
			NetNsID: -1,
			TxQLen:  -1,
		},
		PeerName: name + "b",
	}
	if err := netlink.LinkAdd(&vethLink); err != nil {
		t.Fatalf("while creating veth pair: %v", err)
	}
	vethLinkB, err := netlink.LinkByName(name + "b")
	if err != nil {
		t.Fatalf("while creating veth pair: while getting veth peer: %v", err)
	}

	return &vethLink, vethLinkB.(*netlink.Veth)
}

// setupNetem is a helper for setting up Linux's network emulation queuing
// discipline on a network interface, which can simulate various network
// imperfections like extra latency, reordering or packet loss on packets
// transmitted on the interface specified. As it is internally implemented
// as a queuing discipline it only affects transmitted packets.
func setupNetem(t *testing.T, link netlink.Link, conf netlink.NetemQdiscAttrs) *netlink.Netem {
	t.Helper()
	h := fnv.New32a()
	h.Write([]byte(link.Attrs().Name))
	qdisc := netlink.NewNetem(netlink.QdiscAttrs{
		LinkIndex: link.Attrs().Index,
		Handle:    netlink.MakeHandle(uint16(h.Sum32()%math.MaxUint16), 0),
		Parent:    netlink.HANDLE_ROOT,
	}, conf)
	if err := netlink.QdiscAdd(qdisc); err != nil {
		t.Fatalf("while setting up qdisc netem for %q: %v", link.Attrs().Name, err)
	}
	return qdisc
}

// changeNetem is a helper for reconfiguring an existing netem instance
// on-the-fly with new parameters.
func changeNetem(t *testing.T, qdisc *netlink.Netem, conf netlink.NetemQdiscAttrs) {
	t.Helper()
	changedQd := netlink.NewNetem(qdisc.QdiscAttrs, conf)
	if err := netlink.QdiscChange(changedQd); err != nil {
		t.Fatalf("while changing qdisc netem for link index %v: %v", qdisc.LinkIndex, err)
	}
}

func createBond(t *testing.T, name string, links ...netlink.Link) *netlink.Bond {
	t.Helper()
	bondLink := netlink.NewLinkBond(netlink.LinkAttrs{
		Name:    name,
		NetNsID: -1,
		TxQLen:  -1,
		Flags:   net.FlagUp,
	})
	bondLink.Mode = netlink.BOND_MODE_802_3AD
	bondLink.LacpRate = netlink.BOND_LACP_RATE_FAST
	bondLink.MinLinks = 1
	bondLink.AdSelect = netlink.BOND_AD_SELECT_BANDWIDTH
	if err := netlink.LinkAdd(bondLink); err != nil {
		t.Fatalf("while creating bond: %v", err)
	}
	for _, l := range links {
		if err := netlink.LinkSetBondSlave(l, bondLink); err != nil {
			t.Fatalf("while enslaving link to bond %q: %v", name, err)
		}
	}
	return bondLink
}

// assertRunning is a helper for asserting an interface's IFF_RUNNING state.
func assertRunning(t *testing.T, l netlink.Link, expected bool) {
	t.Helper()
	linkCurrent, err := netlink.LinkByIndex(l.Attrs().Index)
	if err != nil {
		t.Fatalf("while checking if link %q is running: %v", l.Attrs().Name, err)
	}
	is := linkCurrent.Attrs().RawFlags&unix.IFF_RUNNING != 0
	if expected != is {
		t.Errorf("expected interface %q running state to be %v, is %v", l.Attrs().Name, expected, is)
	}
}

func TestLACP(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	if err := os.WriteFile("/sys/kernel/debug/dynamic_debug/control", []byte("module bonding +p"), 0); err != nil {
		t.Fatal(err)
	}
	// Log dynamic debug to console
	if err := os.WriteFile("/proc/sys/kernel/printk", []byte("8"), 0); err != nil {
		t.Fatal(err)
	}

	link1a, link1b := createVethPair(t, "link1")
	link2a, link2b := createVethPair(t, "link2")

	// Drop all traffic
	l1aq := setupNetem(t, link1a, netlink.NetemQdiscAttrs{Loss: 100.0})
	l1bq := setupNetem(t, link1b, netlink.NetemQdiscAttrs{Loss: 100.0})
	l2aq := setupNetem(t, link2a, netlink.NetemQdiscAttrs{Loss: 100.0})
	l2bq := setupNetem(t, link2b, netlink.NetemQdiscAttrs{Loss: 100.0})

	bondA := createBond(t, "bonda", link1a, link2a)
	bondB := createBond(t, "bondb", link1b, link2b)

	time.Sleep(5 * time.Second)

	// Bonds should not come up with links dropping all traffic
	assertRunning(t, bondA, false)
	assertRunning(t, bondB, false)

	changeNetem(t, l1aq, netlink.NetemQdiscAttrs{Loss: 0.0})
	changeNetem(t, l1bq, netlink.NetemQdiscAttrs{Loss: 0.0})
	t.Log("Enabled L1")

	time.Sleep(5 * time.Second)

	// Bonds should come up with one link working
	assertRunning(t, bondA, true)
	assertRunning(t, bondB, true)

	changeNetem(t, l2aq, netlink.NetemQdiscAttrs{Loss: 0.0})
	changeNetem(t, l2bq, netlink.NetemQdiscAttrs{Loss: 0.0})
	t.Log("Enabled L2")

	time.Sleep(3 * time.Second)

	// Bonds be up with both links
	assertRunning(t, bondA, true)
	assertRunning(t, bondB, true)

	bondAState, err := os.ReadFile("/proc/net/bonding/bonda")
	if err != nil {
		panic(err)
	}
	t.Log(string(bondAState))
	if !strings.Contains(string(bondAState), "Number of ports: 2") {
		t.Errorf("bonda aggregator should contain two ports")
	}
	if !strings.Contains(string(bondAState), "port state: 63") {
		t.Errorf("bonda port state should be 63")
	}
	t.Log("------------")
	bondBState, err := os.ReadFile("/proc/net/bonding/bondb")
	if err != nil {
		panic(err)
	}
	t.Log(string(bondBState))
	if !strings.Contains(string(bondBState), "Number of ports: 2") {
		t.Errorf("bondb aggregator should contain two ports")
	}
	if !strings.Contains(string(bondBState), "port state: 63") {
		t.Errorf("bondb port state should be 63")
	}
	changeNetem(t, l1aq, netlink.NetemQdiscAttrs{Loss: 100.0})
	changeNetem(t, l1bq, netlink.NetemQdiscAttrs{Loss: 100.0})
	changeNetem(t, l2aq, netlink.NetemQdiscAttrs{Loss: 100.0})
	changeNetem(t, l2bq, netlink.NetemQdiscAttrs{Loss: 100.0})
	t.Log("Disabled both links")

	time.Sleep(5 * time.Second)

	// Bonds should be back down
	assertRunning(t, bondA, false)
	assertRunning(t, bondB, false)
}
