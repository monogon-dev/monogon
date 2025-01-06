package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cavaliergopher/cpio"
	"github.com/klauspost/compress/zstd"
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/metropolis/proto/api"
	netapi "source.monogon.dev/osbase/net/proto"

	"source.monogon.dev/osbase/bootparam"
	"source.monogon.dev/osbase/build/mkimage/osimage"
	"source.monogon.dev/osbase/kexec"
	netdump "source.monogon.dev/osbase/net/dump"
)

//go:embed third_party/linux/Image
var kernel []byte

//go:embed third_party/ucode.cpio
var ucode []byte

//go:embed initramfs.cpio.zst
var initramfs []byte

// newMemfile creates a new file which is not located on a specific filesystem,
// but is instead backed by anonymous memory.
func newMemfile(name string, flags int) (*os.File, error) {
	fd, err := unix.MemfdCreate(name, flags)
	if err != nil {
		return nil, fmt.Errorf("memfd_create failed: %w", err)
	}
	return os.NewFile(uintptr(fd), name), nil
}

func setupTakeover(nodeParamsRaw []byte, target string) ([]string, error) {
	// Validate we are running via EFI.
	if _, err := os.Stat("/sys/firmware/efi"); os.IsNotExist(err) {
		//nolint:ST1005
		return nil, fmt.Errorf("Monogon OS can only be installed on EFI-booted machines, this one is not")
	}

	currPath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	bundleRaw, err := os.Open(filepath.Join(filepath.Dir(currPath), "bundle.zip"))
	if err != nil {
		return nil, err
	}

	bundleStat, err := bundleRaw.Stat()
	if err != nil {
		return nil, err
	}

	bundle, err := zip.NewReader(bundleRaw, bundleStat.Size())
	if err != nil {
		return nil, fmt.Errorf("failed to open node bundle: %w", err)
	}

	// Dump the current network configuration
	netconf, warnings, err := netdump.Dump()
	if err != nil {
		return nil, fmt.Errorf("failed to dump network configuration: %w", err)
	}

	if len(netconf.Nameserver) == 0 {
		netconf.Nameserver = []*netapi.Nameserver{{
			Ip: "8.8.8.8",
		}, {
			Ip: "1.1.1.1",
		}}
	}

	var params apb.NodeParameters
	if err := proto.Unmarshal(nodeParamsRaw, &params); err != nil {
		return nil, fmt.Errorf("failed to unmarshal node parameters: %w", err)
	}

	// Override the NodeParameters.NetworkConfig with the current NetworkConfig
	// if it's missing.
	if params.NetworkConfig == nil {
		params.NetworkConfig = netconf
	}

	// Marshal NodeParameters again.
	nodeParamsRaw, err = proto.Marshal(&params)
	if err != nil {
		return nil, fmt.Errorf("failed marshaling: %w", err)
	}

	oParams, err := setupOSImageParams(bundle, nodeParamsRaw, target)
	if err != nil {
		return nil, err
	}

	// Validate that this installation will not fail because of disk issues
	if _, err := osimage.Plan(oParams); err != nil {
		return nil, fmt.Errorf("failed to plan installation: %w", err)
	}

	// Load data from embedded files into memfiles as the kexec load syscall
	// requires file descriptors.
	kernelFile, err := newMemfile("kernel", 0)
	if err != nil {
		return nil, fmt.Errorf("failed to create kernel memfile: %w", err)
	}
	initramfsFile, err := newMemfile("initramfs", 0)
	if err != nil {
		return nil, fmt.Errorf("failed to create initramfs memfile: %w", err)
	}
	if _, err := kernelFile.ReadFrom(bytes.NewReader(kernel)); err != nil {
		return nil, fmt.Errorf("failed to read kernel into memory-backed file: %w", err)
	}
	if _, err := initramfsFile.ReadFrom(bytes.NewReader(ucode)); err != nil {
		return nil, fmt.Errorf("failed to read ucode into memory-backed file: %w", err)
	}
	if _, err := initramfsFile.ReadFrom(bytes.NewReader(initramfs)); err != nil {
		return nil, fmt.Errorf("failed to read initramfs into memory-backed file: %w", err)
	}

	// Append this executable, the bundle and node params to initramfs
	compressedW, err := zstd.NewWriter(initramfsFile, zstd.WithEncoderLevel(1))
	if err != nil {
		return nil, fmt.Errorf("while creating zstd writer: %w", err)
	}
	{
		self, err := os.Open("/proc/self/exe")
		if err != nil {
			return nil, err
		}
		selfStat, err := self.Stat()
		if err != nil {
			return nil, err
		}

		cpioW := cpio.NewWriter(compressedW)
		cpioW.WriteHeader(&cpio.Header{
			Name: "/init",
			Size: selfStat.Size(),
			Mode: cpio.TypeReg | 0o755,
		})
		io.Copy(cpioW, self)
		cpioW.Close()
	}
	{
		cpioW := cpio.NewWriter(compressedW)
		cpioW.WriteHeader(&cpio.Header{
			Name: "/bundle.zip",
			Size: bundleStat.Size(),
			Mode: cpio.TypeReg | 0o644,
		})
		bundleRaw.Seek(0, io.SeekStart)
		io.Copy(cpioW, bundleRaw)
		cpioW.Close()
	}
	{
		cpioW := cpio.NewWriter(compressedW)
		cpioW.WriteHeader(&cpio.Header{
			Name: "/params.pb",
			Size: int64(len(nodeParamsRaw)),
			Mode: cpio.TypeReg | 0o644,
		})
		cpioW.Write(nodeParamsRaw)
		cpioW.Close()
	}
	compressedW.Close()

	initParams := bootparam.Params{
		bootparam.Param{Param: "quiet"},
		bootparam.Param{Param: launchModeEnv, Value: launchModeInit},
		bootparam.Param{Param: EnvInstallTarget, Value: target},
		bootparam.Param{Param: "init", Value: "/init"},
	}

	var customConsoles bool
	cmdline, err := os.ReadFile("/proc/cmdline")
	if err != nil {
		warnings = append(warnings, fmt.Errorf("unable to read current kernel command line: %w", err))
	} else {
		params, _, err := bootparam.Unmarshal(string(cmdline))
		// If the existing command line is well-formed, add all existing console
		// parameters to the console for the agent
		if err == nil {
			for _, p := range params {
				if p.Param == "console" {
					initParams = append(initParams, p)
					customConsoles = true
				}
			}
		}
	}
	if !customConsoles {
		// Add the "default" console on x86
		initParams = append(initParams, bootparam.Param{Param: "console", Value: "ttyS0,115200"})
	}
	agentCmdline, err := bootparam.Marshal(initParams, "")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal bootparams: %w", err)
	}
	// Stage agent payload into kernel memory
	if err := kexec.FileLoad(kernelFile, initramfsFile, agentCmdline); err != nil {
		return nil, fmt.Errorf("failed to load kexec payload: %w", err)
	}
	var warningsStrs []string
	for _, w := range warnings {
		warningsStrs = append(warningsStrs, w.Error())
	}
	return warningsStrs, nil
}
