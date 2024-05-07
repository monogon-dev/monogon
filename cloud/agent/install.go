package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/protobuf/proto"

	bpb "source.monogon.dev/cloud/bmaas/server/api"
	"source.monogon.dev/metropolis/node/build/mkimage/osimage"
	npb "source.monogon.dev/net/proto"
	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/efivarfs"
	"source.monogon.dev/osbase/logtree"
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

// install dispatches OSInstallationRequests to the appropriate installer
// method
func install(req *bpb.OSInstallationRequest, netConfig *npb.Net, l logtree.LeveledLogger, isEFIBoot bool) error {
	switch reqT := req.Type.(type) {
	case *bpb.OSInstallationRequest_Metropolis:
		return installMetropolis(reqT.Metropolis, netConfig, l, isEFIBoot)
	default:
		return errors.New("unknown installation request type")
	}
}

func installMetropolis(req *bpb.MetropolisInstallationRequest, netConfig *npb.Net, l logtree.LeveledLogger, isEFIBoot bool) error {
	if !isEFIBoot {
		//nolint:ST1005
		return errors.New("Monogon OS can only be installed on EFI-booted machines, this one is not")
	}

	// Override the NodeParameters.NetworkConfig with the current NetworkConfig
	// if it's missing.
	if req.NodeParameters.NetworkConfig == nil {
		req.NodeParameters.NetworkConfig = netConfig
	}

	// Download into a buffer as ZIP files cannot efficiently be read from
	// HTTP in Go as the ReaderAt has no way of indicating continuous sections,
	// thus a ton of small range requests would need to be used, causing
	// a huge latency penalty as well as costing a lot of money on typical
	// object storages. This should go away when we switch to a better bundle
	// format which can be streamed.
	var bundleRaw bytes.Buffer
	b := backoff.NewExponentialBackOff()
	err := backoff.Retry(func() error {
		bundleRes, err := http.Get(req.BundleUrl)
		if err != nil {
			l.Warningf("Metropolis bundle request failed: %v", err)
			return fmt.Errorf("HTTP request failed: %w", err)
		}
		defer bundleRes.Body.Close()
		switch bundleRes.StatusCode {
		case http.StatusTooEarly, http.StatusTooManyRequests,
			http.StatusInternalServerError, http.StatusBadGateway,
			http.StatusServiceUnavailable, http.StatusGatewayTimeout:
			l.Warningf("Metropolis bundle request HTTP %d error, retrying", bundleRes.StatusCode)
			return fmt.Errorf("HTTP error %d", bundleRes.StatusCode)
		default:
			// Non-standard code range used for proxy-related issue by various
			// vendors. Treat as non-permanent error.
			if bundleRes.StatusCode >= 520 && bundleRes.StatusCode < 599 {
				l.Warningf("Metropolis bundle request HTTP %d error, retrying", bundleRes.StatusCode)
				return fmt.Errorf("HTTP error %d", bundleRes.StatusCode)
			}
			if bundleRes.StatusCode != 200 {
				l.Errorf("Metropolis bundle request permanent HTTP %d error, aborting", bundleRes.StatusCode)
				return backoff.Permanent(fmt.Errorf("HTTP error %d", bundleRes.StatusCode))
			}
		}
		if _, err := bundleRaw.ReadFrom(bundleRes.Body); err != nil {
			l.Warningf("Metropolis bundle download failed, retrying: %v", err)
			bundleRaw.Reset()
			return err
		}
		return nil
	}, b)
	if err != nil {
		return fmt.Errorf("error downloading Metropolis bundle: %w", err)
	}
	l.Info("Metropolis Bundle downloaded")
	bundle, err := zip.NewReader(bytes.NewReader(bundleRaw.Bytes()), int64(bundleRaw.Len()))
	if err != nil {
		return fmt.Errorf("failed to open node bundle: %w", err)
	}
	efiPayload, err := bundle.Open("kernel_efi.efi")
	if err != nil {
		return fmt.Errorf("invalid bundle: %w", err)
	}
	defer efiPayload.Close()
	systemImage, err := bundle.Open("verity_rootfs.img")
	if err != nil {
		return fmt.Errorf("invalid bundle: %w", err)
	}
	defer systemImage.Close()

	nodeParamsRaw, err := proto.Marshal(req.NodeParameters)
	if err != nil {
		return fmt.Errorf("failed marshaling: %w", err)
	}

	rootDev, err := blockdev.Open(filepath.Join("/dev", req.RootDevice))
	if err != nil {
		return fmt.Errorf("failed to open root device: %w", err)
	}

	installParams := osimage.Params{
		PartitionSize: osimage.PartitionSizeInfo{
			ESP:    384,
			System: 4096,
			Data:   128,
		},
		SystemImage:    systemImage,
		EFIPayload:     FileSizedReader{efiPayload},
		ABLoader:       bytes.NewReader(abloader),
		NodeParameters: bytes.NewReader(nodeParamsRaw),
		Output:         rootDev,
	}

	be, err := osimage.Create(&installParams)
	if err != nil {
		return err
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
