// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package launch

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"google.golang.org/grpc"

	"git.monogon.dev/source/nexantic.git/core/internal/common"
)

// This is more of a best-effort solution and not guaranteed to give us unused ports (since we're not immediately using
// them), but AFAIK qemu cannot dynamically select hostfwd ports
func getFreePort() (uint16, io.Closer, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, nil, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, nil, err
	}
	return uint16(l.Addr().(*net.TCPAddr).Port), l, nil
}

type qemuValue map[string][]string

// qemuValueToOption encodes structured data into a QEMU option.
// Example: "test", {"key1": {"val1"}, "key2": {"val2", "val3"}} returns "test,key1=val1,key2=val2,key2=val3"
func qemuValueToOption(name string, value qemuValue) string {
	var optionValues []string
	optionValues = append(optionValues, name)
	for name, values := range value {
		if len(values) == 0 {
			optionValues = append(optionValues, name)
		}
		for _, val := range values {
			optionValues = append(optionValues, fmt.Sprintf("%v=%v", name, val))
		}
	}
	return strings.Join(optionValues, ",")
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// PortMap represents where VM ports are mapped to on the host. It maps from the VM port number to the host port number.
type PortMap map[uint16]uint16

// toQemuForwards generates QEMU hostfwd values (https://qemu.weilnetz.de/doc/qemu-doc.html#:~:text=hostfwd=) for all
// mapped ports.
func (p PortMap) toQemuForwards() []string {
	var hostfwdOptions []string
	for vmPort, hostPort := range p {
		hostfwdOptions = append(hostfwdOptions, fmt.Sprintf("tcp::%v-:%v", hostPort, vmPort))
	}
	return hostfwdOptions
}

// DialGRPC creates a gRPC client for a VM port that's forwarded/mapped to the host. The given port is automatically
// resolved to the host-mapped port.
func (p PortMap) DialGRPC(port uint16, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	mappedPort, ok := p[port]
	if !ok {
		return nil, fmt.Errorf("cannot dial port: port %v is not mapped/forwarded", port)
	}
	grpcClient, err := grpc.Dial(fmt.Sprintf("localhost:%v", mappedPort), opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial port %v: %w", port, err)
	}
	return grpcClient, nil
}

// Options contains all options that can be passed to Launch()
type Options struct {
	// Ports contains the port mapping where to expose the internal ports of the VM to the host. See IdentityPortMap()
	// and ConflictFreePortMap()
	Ports PortMap

	// If set to true, reboots are honored. Otherwise all reboots exit the Launch() command. Smalltown generally restarts
	// on almost all errors, so unless you want to test reboot behavior this should be false.
	AllowReboot bool
}

var requiredPorts = []uint16{common.ConsensusPort, common.NodeServicePort, common.MasterServicePort,
	common.ExternalServicePort, common.DebugServicePort, common.KubernetesAPIPort}

// IdentityPortMap returns a port map where each VM port is mapped onto itself on the host. This is mainly useful
// for development against Smalltown. The dbg command requires this mapping.
func IdentityPortMap() PortMap {
	portMap := make(PortMap)
	for _, port := range requiredPorts {
		portMap[port] = port
	}
	return portMap
}

// ConflictFreePortMap returns a port map where each VM port is mapped onto a random free port on the host. This is
// intended for automated testing where multiple instances of Smalltown might be running. Please call this function for
// each Launch command separately and as close to it as possible since it cannot guarantee that the ports will remain
// free.
func ConflictFreePortMap() (PortMap, error) {
	portMap := make(PortMap)
	for _, port := range requiredPorts {
		mappedPort, listenCloser, err := getFreePort()
		if err != nil {
			return portMap, fmt.Errorf("failed to get free host port: %w", err)
		}
		// Defer closing of the listening port until the function is done and all ports are allocated
		defer listenCloser.Close()
		portMap[port] = mappedPort
	}
	return portMap, nil
}

// Launch launches a Smalltown instance with the given options. The instance runs mostly paravirtualized but with some
// emulated hardware similar to how a cloud provider might set up its VMs. The disk is fully writable but is run
// in snapshot mode meaning that changes are not kept beyond a single invocation.
func Launch(ctx context.Context, options Options) error {
	// Pin temp directory to /tmp until we can use abstract socket namespace in QEMU (next release after 5.0,
	// https://github.com/qemu/qemu/commit/776b97d3605ed0fc94443048fdf988c7725e38a9). swtpm accepts already-open FDs
	// so we can pass in an abstract socket namespace FD that we open and pass the name of it to QEMU. Not pinning this
	// crashes both swtpm and qemu because we run into UNIX socket length limitations (for legacy reasons 108 chars).
	tempDir, err := ioutil.TempDir("/tmp", "launch*")
	if err != nil {
		return fmt.Errorf("Failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy TPM state into a temporary directory since it's being modified by the emulator
	tpmTargetDir := filepath.Join(tempDir, "tpm")
	tpmSrcDir := "core/tpm"
	if err := os.Mkdir(tpmTargetDir, 0644); err != nil {
		return fmt.Errorf("Failed to create TPM state directory: %w", err)
	}
	tpmFiles, err := ioutil.ReadDir(tpmSrcDir)
	if err != nil {
		return fmt.Errorf("Failed to read TPM directory: %w", err)
	}
	for _, file := range tpmFiles {
		name := file.Name()
		if err := copyFile(filepath.Join(tpmSrcDir, name), filepath.Join(tpmTargetDir, name)); err != nil {
			return fmt.Errorf("Failed to copy TPM directory: %w", err)
		}
	}

	qemuNetConfig := qemuValue{
		"id":        {"net0"},
		"net":       {"10.42.0.0/24"},
		"dhcpstart": {"10.42.0.10"},
		"hostfwd":   options.Ports.toQemuForwards(),
	}

	tpmSocketPath := filepath.Join(tempDir, "tpm-socket")

	qemuArgs := []string{"-machine", "q35", "-accel", "kvm", "-nographic", "-nodefaults", "-m", "2048",
		"-cpu", "host", "-smp", "sockets=1,cpus=1,cores=2,threads=2,maxcpus=4",
		"-drive", "if=pflash,format=raw,readonly,file=external/edk2/OVMF_CODE.fd",
		"-drive", "if=pflash,format=raw,snapshot=on,file=external/edk2/OVMF_VARS.fd",
		"-drive", "if=virtio,format=raw,snapshot=on,cache=unsafe,file=core/smalltown.img",
		"-netdev", qemuValueToOption("user", qemuNetConfig),
		"-device", "virtio-net-pci,netdev=net0",
		"-chardev", "socket,id=chrtpm,path=" + tpmSocketPath,
		"-tpmdev", "emulator,id=tpm0,chardev=chrtpm",
		"-device", "tpm-tis,tpmdev=tpm0",
		"-device", "virtio-rng-pci",
		"-serial", "stdio"}

	if !options.AllowReboot {
		qemuArgs = append(qemuArgs, "-no-reboot")
	}

	// Start TPM emulator as a subprocess
	tpmCtx, tpmCancel := context.WithCancel(ctx)
	defer tpmCancel()

	tpmEmuCmd := exec.CommandContext(tpmCtx, "swtpm", "socket", "--tpm2", "--tpmstate", "dir="+tpmTargetDir, "--ctrl", "type=unixio,path="+tpmSocketPath)
	tpmEmuCmd.Stderr = os.Stderr
	tpmEmuCmd.Stdout = os.Stdout

	err = tpmEmuCmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start TPM emulator: %w", err)
	}

	// Start the main qemu binary
	systemCmd := exec.CommandContext(ctx, "qemu-system-x86_64", qemuArgs...)
	systemCmd.Stderr = os.Stderr
	systemCmd.Stdout = os.Stdout

	err = systemCmd.Run()

	// Stop TPM emulator and wait for it to exit to properly reap the child process
	tpmCancel()
	log.Print("Waiting for TPM emulator to exit")
	// Wait returns a SIGKILL error because we just cancelled its context.
	// We still need to call it to avoid creating zombies.
	_ = tpmEmuCmd.Wait()

	return nil
}
