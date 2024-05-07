package network

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mdlayher/ethtool"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/logtree"
)

// applyQuirks applies settings to drivers and/or hardware to make it work
// better (i.e. with less crashes or faster).
func applyQuirks(l logtree.LeveledLogger) error {
	ethtoolFd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, unix.IPPROTO_IP)
	if err != nil {
		return fmt.Errorf("while creating IP socket for ethtool: %w", err)
	}
	defer unix.Close(ethtoolFd)
	ethtoolC, err := ethtool.New()
	if err != nil {
		return fmt.Errorf("while getting ethtool netlink fd: %w", err)
	}
	defer ethtoolC.Close()
	links, err := netlink.LinkList()
	if err != nil {
		return fmt.Errorf("while getting links for applying quirks: %w", err)
	}
	for _, link := range links {
		linkinfo, err := unix.IoctlGetEthtoolDrvinfo(ethtoolFd, link.Attrs().Name)
		if errors.Is(err, unix.EOPNOTSUPP) {
			// These are normally software/virtual devices which should never
			// need quirking.
			continue
		} else if err != nil {
			l.Warningf("Unexpected error during ioctl(ETHTOOL_GDRVINFO) for device %q, skipping quirks: %v", link.Attrs().Name, err)
			continue
		}
		driver := unix.ByteSliceToString(linkinfo.Driver[:])
		firmwareVersion := strings.TrimSpace(unix.ByteSliceToString(linkinfo.Fw_version[:]))
		opromVersion := strings.TrimSpace(unix.ByteSliceToString(linkinfo.Erom_version[:]))

		// Log firmware version of all NICs which have one as we have currently
		// no better way of accessing these.
		if firmwareVersion != "" {
			if opromVersion != "" {
				l.Infof("Interface %q (driver %v) has firmware version %q with Option ROM version %q", link.Attrs().Name, driver, firmwareVersion, opromVersion)
			}
			l.Infof("Interface %q (driver %v) has firmware version %q", link.Attrs().Name, driver, firmwareVersion)
		}

		switch driver {
		case "i40e":
			err := ethtoolC.SetPrivateFlags(ethtool.PrivateFlags{
				Interface: ethtool.Interface{Index: link.Attrs().Index},
				Flags: map[string]bool{
					// Disable firmware-based LLDP processing as it both makes
					// LLDP unavailable to the OS as well as being suspected of
					// causing fimware crashes. Metropolis currently does not
					// have DCB support anyway and if it gains such support it
					// will proccess the LLDP packets for that in userspace.
					"disable-fw-lldp": true,
				},
			})
			if err != nil {
				l.Warningf("Error when applying quirk for LLDP firmware processing to %q: %v", link.Attrs().Name, err)
			} else {
				l.Infof("Applied LLDP firmware processing quirk to %q", link.Attrs().Name)
			}
		}
	}
	return nil
}
