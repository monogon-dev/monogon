package crypt

import (
	"encoding/hex"
	"fmt"
	"os"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/pkg/devicemapper"
)

func encryptionDevPath(name string) string {
	return fmt.Sprintf("/dev/%s-crypt", name)
}

func encryptionDMName(name string) string {
	return fmt.Sprintf("%s-crypt", name)
}

func mapEncryption(name, underlying string, encryptionKey []byte, authenticated bool) (string, error) {
	sizeBytes, err := getSizeBytes(underlying)
	if err != nil {
		return "", fmt.Errorf("getting size of block device failed: %w", err)
	}
	blockSize, err := getBlockSize(underlying)
	if err != nil {
		return "", fmt.Errorf("getting block size failed: %w", err)
	}

	optParams := []string{
		"no_read_workqueue", "no_write_workqueue",
	}
	cipher := "capi:xts(aes)-essiv:sha256"
	if authenticated {
		optParams = append(optParams, "integrity:28:aead")
		cipher = "capi:gcm(aes)-random"
	} else {
		// discard (TRIM/UNMAP) only works without integrity enabled.
		optParams = append(optParams, "allow_discards")
	}
	params := []string{
		// cipher, key, iv_offset, device_path, offset
		cipher, hex.EncodeToString(encryptionKey), "0", underlying, "0",
		// number of opt params
		fmt.Sprintf("%d", len(optParams)),
	}
	params = append(params, optParams...)

	cryptDev, err := devicemapper.CreateActiveDevice(encryptionDMName(name), false, []devicemapper.Target{
		{
			Length:     sizeBytes / uint64(blockSize),
			Type:       "crypt",
			Parameters: params,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create crypt device: %w", err)
	}
	if err := unix.Mknod(encryptionDevPath(name), 0600|unix.S_IFBLK, int(cryptDev)); err != nil {
		// Best-effort cleanup, swallow errors.
		unmapEncryption(name)
		return "", fmt.Errorf("failed to create crypt device node: %w", err)
	}
	return encryptionDevPath(name), nil
}

func unmapEncryption(name string) error {
	// Remove /dev node if present.
	if _, err := os.Stat(encryptionDevPath(name)); err == nil {
		if err := unix.Unlink(encryptionDevPath(name)); err != nil {
			return fmt.Errorf("unlinking encryption device failed: %w", err)
		}
	}

	// Remove dm target.
	if err := devicemapper.RemoveDevice(encryptionDMName(name)); err != nil {
		return fmt.Errorf("removing encryption device failed: %w", err)
	}
	return nil
}
