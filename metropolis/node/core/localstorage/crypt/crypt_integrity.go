// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package crypt

import (
	"encoding/binary"
	"fmt"
	"os"

	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/devicemapper"
)

func integrityDevPath(name string) string {
	return fmt.Sprintf("/dev/%s-integrity", name)
}

func integrityDMName(name string) string {
	return fmt.Sprintf("%s-integrity", name)
}

// readIntegrityDataSectors parses the number of available integrity data sectors
// from a raw dm-integrity formatted device. This is needed to then map the
// device.
//
// This is described in further detail in
// https://docs.kernel.org/admin-guide/device-mapper/dm-integrity.html.
func readIntegrityDataSectors(path string) (uint64, error) {
	integrityPartition, err := blockdev.Open(path)
	if err != nil {
		return 0, err
	}
	defer integrityPartition.Close()

	firstBlock := make([]byte, integrityPartition.BlockSize())
	if _, err = integrityPartition.ReadAt(firstBlock, 0); err != nil {
		return 0, err
	}
	// Based on structure defined in
	//   https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/drivers/md/dm-integrity.c#n59
	providedDataSectors := binary.LittleEndian.Uint64(firstBlock[16:24])

	// Let's perform some simple checks on the read value to make sure the returned
	// data isn't corrupted or has been tampered with.

	if providedDataSectors == 0 {
		return 0, fmt.Errorf("invalid data sector count of zero")
	}

	if providedDataSectors > uint64(integrityPartition.BlockCount()) {
		return 0, fmt.Errorf("device claims %d data sectors but underlying device only has %d", providedDataSectors, integrityPartition.BlockCount())
	}
	return providedDataSectors, nil
}

// initializeIntegrity performs the initialization steps outlined in
// https://docs.kernel.org/admin-guide/device-mapper/dm-integrity.html.
func initializeIntegrity(name, baseName string) error {
	// Zero out superblock.
	integrityPartition, err := os.OpenFile(baseName, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	zeroedBuf := make([]byte, 4096)
	if _, err := integrityPartition.Write(zeroedBuf); err != nil {
		integrityPartition.Close()
		return fmt.Errorf("failed to wipe header: %w", err)
	}
	integrityPartition.Close()

	// Load target with one-sector size. The kernel will format the device.
	_, err = devicemapper.CreateActiveDevice(integrityDMName(name), false, []devicemapper.Target{
		{
			Length:     1,
			Type:       "integrity",
			Parameters: []string{baseName, "0", "28", "J", "1", "journal_sectors:1024"},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create initial integrity device: %w", err)
	}
	// Unload the target.
	if err := devicemapper.RemoveDevice(integrityDMName(name)); err != nil {
		return fmt.Errorf("failed to remove initial integrity device: %w", err)
	}

	return nil
}

func mapIntegrity(name, baseName string, enableJournal bool) (string, error) {
	integritySectors, err := readIntegrityDataSectors(baseName)
	if err != nil {
		return "", fmt.Errorf("failed to read the number of usable sectors on the integrity device: %w", err)
	}

	mode := "D"
	if enableJournal {
		mode = "J"
	}
	integrityDev, err := devicemapper.CreateActiveDevice(integrityDMName(name), false, []devicemapper.Target{
		{
			Length:     integritySectors,
			Type:       "integrity",
			Parameters: []string{baseName, "0", "28", mode, "1", "journal_sectors:1024"},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create Integrity device: %w", err)
	}
	if err := unix.Mknod(integrityDevPath(name), 0600|unix.S_IFBLK, int(integrityDev)); err != nil {
		unmapIntegrity(name)
		return "", fmt.Errorf("failed to create integrity device node: %w", err)
	}

	return integrityDevPath(name), nil
}

func unmapIntegrity(name string) error {
	// Remove /dev node if present.
	if _, err := os.Stat(integrityDevPath(name)); err == nil {
		if err := unix.Unlink(integrityDevPath(name)); err != nil {
			return fmt.Errorf("unlinking integrity device failed: %w", err)
		}
	}

	if err := devicemapper.RemoveDevice(integrityDMName(name)); err != nil {
		return fmt.Errorf("removing integrity DM device failed: %w", err)
	}
	return nil
}
