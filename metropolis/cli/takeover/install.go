package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"source.monogon.dev/go/logging"
	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/build/mkimage/osimage"
	"source.monogon.dev/osbase/efivarfs"
)

//go:embed metropolis/node/core/abloader/abloader_bin.efi
var abloader []byte

// FileSizedReader is a small adapter from fs.File to fs.SizedReader
// Panics on Stat() failure, so should only be used with sources where Stat()
// cannot fail.
type FileSizedReader struct {
	fs.File
}

func (f FileSizedReader) Size() int64 {
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return stat.Size()
}

// EnvInstallTarget environment variable which tells the takeover binary where
// to install to
const EnvInstallTarget = "TAKEOVER_INSTALL_TARGET"

func installMetropolis(l logging.Leveled) error {
	// Validate we are running via EFI.
	if _, err := os.Stat("/sys/firmware/efi"); os.IsNotExist(err) {
		//nolint:ST1005
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

	efiPayload, err := bundle.Open("kernel_efi.efi")
	if err != nil {
		return nil, fmt.Errorf("invalid bundle: %w", err)
	}

	systemImage, err := bundle.Open("verity_rootfs.img")
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
		EFIPayload:     FileSizedReader{efiPayload},
		ABLoader:       bytes.NewReader(abloader),
		NodeParameters: bytes.NewReader(metropolisSpecRaw),
		Output:         rootDev,
	}, nil
}
