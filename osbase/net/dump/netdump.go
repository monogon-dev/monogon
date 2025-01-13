package netdump

import (
	"bytes"
	"fmt"
	"math"
	"net/netip"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/vishvananda/netlink"
	"go4.org/netipx"
	"golang.org/x/sys/unix"

	netapi "source.monogon.dev/osbase/net/proto"
)

var vlanProtoMap = map[netlink.VlanProtocol]netapi.VLAN_Protocol{
	netlink.VLAN_PROTOCOL_8021Q:  netapi.VLAN_PROTOCOL_CVLAN,
	netlink.VLAN_PROTOCOL_8021AD: netapi.VLAN_PROTOCOL_SVLAN,
}

// From iproute2's rt_protos
const (
	protoUnspec = 0
	protoKernel = 2
	protoBoot   = 3
	protoStatic = 4
)

type ifaceAddrRef struct {
	iface   *netapi.Interface
	addrIdx int
}

// Dump dumps the network configuration of the current network namespace into
// a osbase.net.proto.Net structure. This is currently only expected to work for
// systems which do not use a dynamic routing protocol to establish basic
// internet connectivity.
// The second return value is a list of warnings, i.e. things which might be
// problematic, but might still result in a working (though less complete)
// configuration. The Net in the first return value is set to a non-nil value
// even if there are warnings. The third return value is for a hard error,
// the Net value will be nil in that case.
func Dump() (*netapi.Net, []error, error) {
	var n netapi.Net
	var warnings []error
	links, err := netlink.LinkList()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list network links: %w", err)
	}
	// Map interface index -> interface pointer
	ifIdxMap := make(map[int]*netapi.Interface)
	// Map interface index -> names of children
	ifChildren := make(map[int][]string)
	// Interface address implied on-link routes
	impliedOnLinkRoutes := make(map[netip.Prefix]ifaceAddrRef)
	// Map interface index -> number of reverse dependencies
	ifNRevDeps := make(map[int]int)
	for _, link := range links {
		linkAttrs := link.Attrs()
		// Ignore loopback interfaces. The default one will always be
		// created, and we don't have support for additional loopbacks.
		if linkAttrs.EncapType == "loopback" {
			continue
		}
		// Gather interface-type-specific data into a netapi interface.
		var iface netapi.Interface
		switch l := link.(type) {
		case *netlink.Device:
			mac := link.Attrs().PermHWAddr
			if len(mac) == 0 {
				// Try legacy method for old kernels
				mac, err = getPermanentHWAddrLegacy(l.Name)
				// Errors are expected, not all interfaces support this.
				// If a permanent hardware address could not be obtained, fall
				// back to the configured hardware address.
				if err != nil {
					mac = link.Attrs().HardwareAddr
				}
			}
			iface.Type = &netapi.Interface_Device{Device: &netapi.Device{
				HardwareAddress: mac.String(),
			}}
		case *netlink.Bond:
			bond := netapi.Bond{
				MinLinks:           int32(l.MinLinks),
				TransmitHashPolicy: netapi.Bond_TransmitHashPolicy(l.XmitHashPolicy),
			}
			switch l.Mode {
			case netlink.BOND_MODE_802_3AD:
				lacp := netapi.Bond_LACP{
					Rate:                netapi.Bond_LACP_Rate(l.LacpRate),
					ActorSystemPriority: int32(l.AdActorSysPrio),
					UserPortKey:         int32(l.AdUserPortKey),
					SelectionLogic:      netapi.Bond_LACP_SelectionLogic(l.AdSelect),
				}
				if len(bytes.TrimLeft(l.AdActorSystem, "\x00")) != 0 {
					lacp.ActorSystemMac = l.AdActorSystem.String()
				}
				bond.Mode = &netapi.Bond_Lacp{Lacp: &lacp}
			default:
			}
			iface.Type = &netapi.Interface_Bond{Bond: &bond}
		case *netlink.Vlan:
			parentLink, err := netlink.LinkByIndex(l.ParentIndex)
			if err != nil {
				warnings = append(warnings, fmt.Errorf("unable to get parent for VLAN interface %q, interface ignored: %w", iface.Name, err))
				continue
			}
			iface.Type = &netapi.Interface_Vlan{Vlan: &netapi.VLAN{
				Id:       int32(l.VlanId),
				Protocol: vlanProtoMap[l.VlanProtocol],
				Parent:   parentLink.Attrs().Name,
			}}
		default:
			continue
		}
		// Append common interface data to netapi interface.
		iface.Name = linkAttrs.Name
		iface.Mtu = int32(linkAttrs.MTU)
		// Collect addresses into interface.
		addrs, err := netlink.AddrList(link, netlink.FAMILY_ALL)
		if err != nil {
			warnings = append(warnings, fmt.Errorf("unable to get addresses for interface %q, interface ignored: %w", iface.Name, err))
			continue
		}
		for _, a := range addrs {
			// Ignore IPv6 link-local addresses
			if a.IP.IsLinkLocalUnicast() && a.IP.To4() == nil {
				continue
			}
			// Sadly it's not possible to reliably determine if a DHCP client is
			// running. Good clients usually either don't set the permanent flag
			// and/or a lifetime.
			if a.Flags&unix.IFA_F_PERMANENT == 0 || (a.ValidLft > 0 && a.ValidLft < math.MaxUint32) {
				if a.IP.To4() == nil {
					// Enable IPv6 Autoconfig
					if iface.Ipv6Autoconfig == nil {
						iface.Ipv6Autoconfig = &netapi.IPv6Autoconfig{}
						iface.Ipv6Autoconfig.Privacy, err = getIPv6IfaceAutoconfigPrivacy(linkAttrs.Name)
						if err != nil {
							warnings = append(warnings, err)
						}
					}
				} else {
					if iface.Ipv4Autoconfig == nil {
						iface.Ipv4Autoconfig = &netapi.IPv4Autoconfig{}
					}
				}
				// Dynamic address, ignore
				continue
			}
			if a.Peer != nil {
				// Add an interface route for the peer
				iface.Route = append(iface.Route, &netapi.Interface_Route{
					Destination: a.Peer.String(),
					SourceIp:    a.IP.String(),
				})
			}
			ones, bits := a.Mask.Size()
			baseAddr, ok := netip.AddrFromSlice(a.IP.Mask(a.Mask))
			prefix := netip.PrefixFrom(baseAddr, ones)
			if ok && bits != 0 {
				if !prefix.IsSingleIP() {
					impliedOnLinkRoutes[prefix] = ifaceAddrRef{iface: &iface, addrIdx: len(iface.Address)}
				}
				iface.Address = append(iface.Address, a.IPNet.String())
			} else {
				warnings = append(warnings, fmt.Errorf("address %v on %q is invalid, ignoring", a.IPNet, iface.Name))
			}
		}
		if linkAttrs.MasterIndex != 0 {
			ifChildren[linkAttrs.MasterIndex] = append(ifChildren[linkAttrs.MasterIndex], iface.Name)
			ifNRevDeps[linkAttrs.Index]++
		}
		if linkAttrs.ParentIndex != 0 {
			ifNRevDeps[linkAttrs.ParentIndex]++
		}
		ifIdxMap[link.Attrs().Index] = &iface
	}
	routes, err := netlink.RouteList(nil, netlink.FAMILY_ALL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list routes: %w", err)
	}
	// Collect all routes into routes assigned to exact netapi interfaces.
	for _, r := range routes {
		if r.Family != netlink.FAMILY_V4 && r.Family != netlink.FAMILY_V6 {
			continue
		}
		var route netapi.Interface_Route
		// Ignore all dynamic routes
		if r.Protocol != protoUnspec && r.Protocol != protoBoot &&
			r.Protocol != protoStatic {
			continue
		}
		if r.LinkIndex == 0 {
			// Only for "exotic" routes like "unreachable" which are not
			// necessary for connectivity, skip for now
			continue
		}
		if r.Dst == nil {
			switch r.Family {
			case netlink.FAMILY_V4:
				route.Destination = "0.0.0.0/0"
			case netlink.FAMILY_V6:
				route.Destination = "::/0"
			default:
				// Switch is complete, all other families get ignored at the start
				// of the loop.
				panic("route family changed under us")
			}
		} else {
			dst, ok := netipx.FromStdIPNet(r.Dst)
			if !ok {
				warnings = append(warnings, fmt.Errorf("route %v invalid, ignoring", r.Dst))
			}
			if ref, ok := impliedOnLinkRoutes[dst]; ok && !r.Gw.IsUnspecified() && len(r.Gw) != 0 {
				// Address is not on-link, remove prefix from address to not
				// get an improper on-link route.
				prefix := netip.MustParsePrefix(ref.iface.Address[ref.addrIdx])
				ref.iface.Address[ref.addrIdx] = netip.PrefixFrom(prefix.Addr(), prefix.Addr().BitLen()).String()
			}
			route.Destination = r.Dst.String()
		}
		if !r.Gw.IsUnspecified() && len(r.Gw) != 0 {
			route.GatewayIp = r.Gw.String()
		}
		if !r.Src.IsUnspecified() && len(r.Src) != 0 {
			route.SourceIp = r.Src.String()
		}
		// Linux calls the metric RTA_PRIORITY even though it behaves as lower-
		// is-better. Note that RTA_METRICS is NOT the metric.
		route.Metric = int32(r.Priority)
		iface, ok := ifIdxMap[r.LinkIndex]
		if !ok {
			continue
		}

		iface.Route = append(iface.Route, &route)
	}
	// Finally, gather all interface into a list, filtering out unused ones.
	for ifIdx, iface := range ifIdxMap {
		switch i := iface.Type.(type) {
		case *netapi.Interface_Bond:
			// Add children here, as now they are all known
			i.Bond.MemberInterface = ifChildren[ifIdx]
		case *netapi.Interface_Device:
			// Drop physical interfaces from the config if they have no IPs and
			// no reverse dependencies.
			if len(iface.Address) == 0 && iface.Ipv4Autoconfig == nil &&
				iface.Ipv6Autoconfig == nil && ifNRevDeps[ifIdx] == 0 {
				continue
			}
		}
		n.Interface = append(n.Interface, iface)
	}
	// Make the output stable
	sort.Slice(n.Interface, func(i, j int) bool { return n.Interface[i].Name < n.Interface[j].Name })
	return &n, warnings, nil
}

func getIPv6IfaceAutoconfigPrivacy(name string) (netapi.IPv6Autoconfig_Privacy, error) {
	useTempaddrRaw, err := os.ReadFile(fmt.Sprintf("/proc/sys/net/ipv6/conf/%s/use_tempaddr", name))
	if err != nil {
		return netapi.IPv6Autoconfig_PRIVACY_DISABLE, fmt.Errorf("failed to read use_tempaddr sysctl for interface %q: %w", name, err)
	}
	useTempaddr, err := strconv.ParseInt(strings.TrimSpace(string(useTempaddrRaw)), 10, 64)
	if err != nil {
		return netapi.IPv6Autoconfig_PRIVACY_DISABLE, fmt.Errorf("failed to parse use_tempaddr sysctl for interface %q: %w", name, err)
	}
	switch {
	case useTempaddr <= 0:
		return netapi.IPv6Autoconfig_PRIVACY_DISABLE, nil
	case useTempaddr == 1:
		return netapi.IPv6Autoconfig_PRIVACY_AVOID, nil
	case useTempaddr > 1:
		return netapi.IPv6Autoconfig_PRIVACY_PREFER, nil
	default:
		panic("switch is complete but hit default case")
	}
}
