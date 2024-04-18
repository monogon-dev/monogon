// takeover is a self-contained executable which when executed loads the BMaaS
// agent via kexec. It is intended to be called over SSH, given a binary
// TakeoverInit message over standard input and (if all preparation work
// completed successfully) will respond with a TakeoverResponse on standard
// output. At that point the new kernel and agent initramfs are fully staged
// by the current kernel.
// The second stage which is also part of this binary, selected by an
// environment variable, is then executed in detached mode and the main
// takeover binary called over SSH terminates.
// The second stage waits for 5 seconds for the main binary to exit, the SSH
// session to be torn down and various other things before issuing the final
// non-returning syscall which jumps into the new kernel.

package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/cavaliergopher/cpio"
	"github.com/klauspost/compress/zstd"
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/cloud/agent/api"
	"source.monogon.dev/metropolis/pkg/bootparam"
	"source.monogon.dev/metropolis/pkg/kexec"
	netdump "source.monogon.dev/net/dump"
	netapi "source.monogon.dev/net/proto"
)

//go:embed third_party/linux/bzImage
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

func setupTakeover() (*api.TakeoverSuccess, error) {
	// Read init specification from stdin.
	initRaw, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("failed to read TakeoverInit message from stdin: %w", err)
	}
	var takeoverInit api.TakeoverInit
	if err := proto.Unmarshal(initRaw, &takeoverInit); err != nil {
		return nil, fmt.Errorf("failed to parse TakeoverInit messag from stdin: %w", err)
	}

	// Sanity check for empty TakeoverInit messages
	if takeoverInit.BmaasEndpoint == "" {
		return nil, errors.New("BMaaS endpoint is empty, check that a proper TakeoverInit message has been provided")
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

	// Generate agent private key
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("unable to generate Ed25519 key: %w", err)
	}

	agentInit := api.AgentInit{
		TakeoverInit:  &takeoverInit,
		PrivateKey:    privKey,
		NetworkConfig: netconf,
	}
	agentInitRaw, err := proto.Marshal(&agentInit)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal AgentInit message: %v", err)
	}

	// Append AgentInit spec to initramfs
	compressedW, err := zstd.NewWriter(initramfsFile, zstd.WithEncoderLevel(1))
	if err != nil {
		return nil, fmt.Errorf("while creating zstd writer: %w", err)
	}
	cpioW := cpio.NewWriter(compressedW)
	cpioW.WriteHeader(&cpio.Header{
		Name: "/init.pb",
		Size: int64(len(agentInitRaw)),
		Mode: cpio.TypeReg | 0o644,
	})
	cpioW.Write(agentInitRaw)
	cpioW.Close()
	compressedW.Close()

	agentParams := bootparam.Params{
		bootparam.Param{Param: "quiet"},
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
					agentParams = append(agentParams, p)
					customConsoles = true
				}
			}
		}
	}
	if !customConsoles {
		// Add the "default" console on x86
		agentParams = append(agentParams, bootparam.Param{Param: "console", Value: "ttyS0,115200"})
	}
	agentCmdline, err := bootparam.Marshal(agentParams, "")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal agent params: %w", err)
	}
	// Stage agent payload into kernel memory
	if err := kexec.FileLoad(kernelFile, initramfsFile, agentCmdline); err != nil {
		return nil, fmt.Errorf("failed to load kexec payload: %w", err)
	}
	var warningsStrs []string
	for _, w := range warnings {
		warningsStrs = append(warningsStrs, w.Error())
	}
	return &api.TakeoverSuccess{
		InitMessage: &takeoverInit,
		Key:         pubKey,
		Warning:     warningsStrs,
	}, nil
}

// Environment variable which tells the takeover binary to run the second stage
const detachedLaunchEnv = "TAKEOVER_DETACHED_LAUNCH"

func main() {
	// Check if the second stage should be executed
	if os.Getenv(detachedLaunchEnv) == "1" {
		// Wait 5 seconds for data to be sent, connections to be closed and
		// syncs to be executed
		time.Sleep(5 * time.Second)
		// Perform kexec, this will not return unless it fails
		err := unix.Reboot(unix.LINUX_REBOOT_CMD_KEXEC)
		var msg = "takeover: reboot succeeded, but we're still runing??"
		if err != nil {
			msg = err.Error()
		}
		// We have no standard output/error anymore, if this fails it's
		// just borked. Attempt to dump the error into kmesg for manual
		// debugging.
		kmsg, err := os.OpenFile("/dev/kmsg", os.O_WRONLY, 0)
		if err != nil {
			os.Exit(2)
		}
		kmsg.WriteString(msg)
		kmsg.Close()
		os.Exit(1)
	}

	var takeoverResp api.TakeoverResponse
	res, err := setupTakeover()
	if err != nil {
		takeoverResp.Result = &api.TakeoverResponse_Error{Error: &api.TakeoverError{
			Message: err.Error(),
		}}
	} else {
		takeoverResp.Result = &api.TakeoverResponse_Success{Success: res}
	}
	// Respond to stdout
	takeoverRespRaw, err := proto.Marshal(&takeoverResp)
	if err != nil {
		log.Fatalf("failed to marshal response: %v", err)
	}
	if _, err := os.Stdout.Write(takeoverRespRaw); err != nil {
		log.Fatalf("failed to write response to stdout: %v", err)
	}
	// Close stdout, we're done responding
	os.Stdout.Close()

	// Start second stage which waits for 5 seconds while performing
	// final cleanup.
	detachedCmd := exec.Command("/proc/self/exe")
	detachedCmd.Env = []string{detachedLaunchEnv + "=1"}
	if err := detachedCmd.Start(); err != nil {
		log.Fatalf("failed to launch final stage: %v", err)
	}
	// Release the second stage so that the first stage can cleanly terminate.
	if err := detachedCmd.Process.Release(); err != nil {
		log.Fatalf("error releasing final stage process: %v", err)
	}
}
