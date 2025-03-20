// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"archive/zip"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"source.monogon.dev/go/logging"
	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/build/mkimage/osimage"
	"source.monogon.dev/osbase/efivarfs"
	"source.monogon.dev/osbase/structfs"
)

//go:embed metropolis/node/core/abloader/abloader.efi
var abloader []byte

// zipBlob looks up a file in a [zip.Reader] and adapts it to [structfs.Blob].
func zipBlob(reader *zip.Reader, name string) (zipFileBlob, error) {
	for _, file := range reader.File {
		if file.Name == name {
			return zipFileBlob{file}, nil
		}
	}
	return zipFileBlob{}, fmt.Errorf("file %q not found", name)
}

type zipFileBlob struct {
	*zip.File
}

func (f zipFileBlob) Size() int64 {
	return int64(f.File.UncompressedSize64)
}

// EnvInstallTarget environment variable which tells the takeover binary where
// to install to
const EnvInstallTarget = "TAKEOVER_INSTALL_TARGET"

func installMetropolis(l logging.Leveled) error {
	// Validate we are running via EFI.
	if _, err := os.Stat("/sys/firmware/efi"); os.IsNotExist(err) {
		// nolint:ST1005
		return fmt.Errorf("Monogon OS can only be installed on EFI-booted machines, this one is not")
	}

	metropolisSpecRaw, err := os.ReadFile("/params.pb")
	if err != nil {
		return err
	}

	bundleRaw, err := os.Open("/bundle.zip")
	if err != nil {
		return err
	}

	bundleStat, err := bundleRaw.Stat()
	if err != nil {
		return err
	}

	bundle, err := zip.NewReader(bundleRaw, bundleStat.Size())
	if err != nil {
		return fmt.Errorf("failed to open node bundle: %w", err)
	}

	installParams, err := setupOSImageParams(bundle, metropolisSpecRaw, os.Getenv(EnvInstallTarget))
	if err != nil {
		return err
	}

	be, err := osimage.Write(installParams)
	if err != nil {
		return fmt.Errorf("failed to apply installation: %w", err)
	}
	bootEntryIdx, err := efivarfs.AddBootEntry(be)
	if err != nil {
		return fmt.Errorf("error creating EFI boot entry: %w", err)
	}
	if err := efivarfs.SetBootOrder(efivarfs.BootOrder{uint16(bootEntryIdx)}); err != nil {
		return fmt.Errorf("error setting EFI boot order: %w", err)
	}
	l.Info("Metropolis installation completed")
	return nil
}

func setupOSImageParams(bundle *zip.Reader, metropolisSpecRaw []byte, installTarget string) (*osimage.Params, error) {
	rootDev, err := blockdev.Open(filepath.Join("/dev", installTarget))
	if err != nil {
		return nil, fmt.Errorf("failed to open root device: %w", err)
	}

	efiPayload, err := zipBlob(bundle, "kernel_efi.efi")
	if err != nil {
		return nil, fmt.Errorf("invalid bundle: %w", err)
	}

	systemImage, err := zipBlob(bundle, "verity_rootfs.img")
	if err != nil {
		return nil, fmt.Errorf("invalid bundle: %w", err)
	}

	return &osimage.Params{
		PartitionSize: osimage.PartitionSizeInfo{
			ESP:    384,
			System: 4096,
			Data:   128,
		},
		SystemImage:    systemImage,
		EFIPayload:     efiPayload,
		ABLoader:       structfs.Bytes(abloader),
		NodeParameters: structfs.Bytes(metropolisSpecRaw),
		Output:         rootDev,
	}, nil
}
