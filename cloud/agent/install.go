// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/protobuf/proto"

	bpb "source.monogon.dev/cloud/bmaas/server/api"
	"source.monogon.dev/go/logging"
	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/build/mkimage/osimage"
	"source.monogon.dev/osbase/efivarfs"
	npb "source.monogon.dev/osbase/net/proto"
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

// install dispatches OSInstallationRequests to the appropriate installer
// method
func install(req *bpb.OSInstallationRequest, netConfig *npb.Net, l logging.Leveled) error {
	switch reqT := req.Type.(type) {
	case *bpb.OSInstallationRequest_Metropolis:
		return installMetropolis(reqT.Metropolis, netConfig, l)
	default:
		return errors.New("unknown installation request type")
	}
}

func installMetropolis(req *bpb.MetropolisInstallationRequest, netConfig *npb.Net, l logging.Leveled) error {
	// Validate we are running via EFI.
	if _, err := os.Stat("/sys/firmware/efi"); os.IsNotExist(err) {
		// nolint:ST1005
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
	efiPayload, err := zipBlob(bundle, "kernel_efi.efi")
	if err != nil {
		return fmt.Errorf("invalid bundle: %w", err)
	}
	systemImage, err := zipBlob(bundle, "verity_rootfs.img")
	if err != nil {
		return fmt.Errorf("invalid bundle: %w", err)
	}

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
		EFIPayload:     efiPayload,
		ABLoader:       structfs.Bytes(abloader),
		NodeParameters: structfs.Bytes(nodeParamsRaw),
		Output:         rootDev,
	}

	be, err := osimage.Write(&installParams)
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
