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
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/golang/protobuf/proto"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/pkg/freeport"
	apb "source.monogon.dev/metropolis/proto/api"
)

type qemuValue map[string][]string

// toOption encodes structured data into a QEMU option.
// Example: "test", {"key1": {"val1"}, "key2": {"val2", "val3"}} returns "test,key1=val1,key2=val2,key2=val3"
func (value qemuValue) toOption(name string) string {
	var optionValues []string
	if name != "" {
		optionValues = append(optionValues, name)
	}
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
	// and ConflictFreePortMap(). Ignored when ConnectToSocket is set.
	Ports PortMap

	// If set to true, reboots are honored. Otherwise all reboots exit the Launch() command. Metropolis nodes
	// generally restarts on almost all errors, so unless you want to test reboot behavior this should be false.
	AllowReboot bool

	// By default the VM is connected to the Host via SLIRP. If ConnectToSocket is set, it is instead connected
	// to the given file descriptor/socket. If this is set, all port maps from the Ports option are ignored.
	// Intended for networking this instance together with others for running  more complex network configurations.
	ConnectToSocket *os.File

	// SerialPort is a io.ReadWriter over which you can communicate with the serial port of the machine
	// It can be set to an existing file descriptor (like os.Stdout/os.Stderr) or any Go structure implementing this interface.
	SerialPort io.ReadWriter

	// EnrolmentConfig is passed into the VM and subsequently used for bootstrapping if no enrolment config is built-in
	EnrolmentConfig *apb.EnrolmentConfig
}

// NodePorts is the list of ports a fully operational Metropolis node listens on
var NodePorts = []uint16{node.ConsensusPort, node.NodeServicePort, node.MasterServicePort,
	node.ExternalServicePort, node.DebugServicePort, node.KubernetesAPIPort, node.DebuggerPort}

// IdentityPortMap returns a port map where each given port is mapped onto itself on the host. This is mainly useful
// for development against Metropolis. The dbg command requires this mapping.
func IdentityPortMap(ports []uint16) PortMap {
	portMap := make(PortMap)
	for _, port := range ports {
		portMap[port] = port
	}
	return portMap
}

// ConflictFreePortMap returns a port map where each given port is mapped onto a random free port on the host. This is
// intended for automated testing where multiple instances of Metropolis nodes might be running. Please call this
// function for each Launch command separately and as close to it as possible since it cannot guarantee that the ports
// will remain free.
func ConflictFreePortMap(ports []uint16) (PortMap, error) {
	portMap := make(PortMap)
	for _, port := range ports {
		mappedPort, listenCloser, err := freeport.AllocateTCPPort()
		if err != nil {
			return portMap, fmt.Errorf("failed to get free host port: %w", err)
		}
		// Defer closing of the listening port until the function is done and all ports are allocated
		defer listenCloser.Close()
		portMap[port] = mappedPort
	}
	return portMap, nil
}

// Gets a random EUI-48 Ethernet MAC address
func generateRandomEthernetMAC() (*net.HardwareAddr, error) {
	macBuf := make([]byte, 6)
	_, err := rand.Read(macBuf)
	if err != nil {
		return nil, fmt.Errorf("failed to read randomness for MAC: %v", err)
	}

	// Set U/L bit and clear I/G bit (locally administered individual MAC)
	// Ref IEEE 802-2014 Section 8.2.2
	macBuf[0] = (macBuf[0] | 2) & 0xfe
	mac := net.HardwareAddr(macBuf)
	return &mac, nil
}

// Launch launches a Metropolis node instance with the given options. The instance runs mostly paravirtualized but
// with some emulated hardware similar to how a cloud provider might set up its VMs. The disk is fully writable but
// is run in snapshot mode meaning that changes are not kept beyond a single invocation.
func Launch(ctx context.Context, options Options) error {
	// Pin temp directory to /tmp until we can use abstract socket namespace in QEMU (next release after 5.0,
	// https://github.com/qemu/qemu/commit/776b97d3605ed0fc94443048fdf988c7725e38a9). swtpm accepts already-open FDs
	// so we can pass in an abstract socket namespace FD that we open and pass the name of it to QEMU. Not pinning this
	// crashes both swtpm and qemu because we run into UNIX socket length limitations (for legacy reasons 108 chars).
	tempDir, err := ioutil.TempDir("/tmp", "launch*")
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy TPM state into a temporary directory since it's being modified by the emulator
	tpmTargetDir := filepath.Join(tempDir, "tpm")
	tpmSrcDir := "metropolis/node/tpm"
	if err := os.Mkdir(tpmTargetDir, 0644); err != nil {
		return fmt.Errorf("failed to create TPM state directory: %w", err)
	}
	tpmFiles, err := ioutil.ReadDir(tpmSrcDir)
	if err != nil {
		return fmt.Errorf("failed to read TPM directory: %w", err)
	}
	for _, file := range tpmFiles {
		name := file.Name()
		if err := copyFile(filepath.Join(tpmSrcDir, name), filepath.Join(tpmTargetDir, name)); err != nil {
			return fmt.Errorf("failed to copy TPM directory: %w", err)
		}
	}

	var qemuNetType string
	var qemuNetConfig qemuValue
	if options.ConnectToSocket != nil {
		qemuNetType = "socket"
		qemuNetConfig = qemuValue{
			"id": {"net0"},
			"fd": {"3"},
		}
	} else {
		qemuNetType = "user"
		qemuNetConfig = qemuValue{
			"id":        {"net0"},
			"net":       {"10.42.0.0/24"},
			"dhcpstart": {"10.42.0.10"},
			"hostfwd":   options.Ports.toQemuForwards(),
		}
	}

	tpmSocketPath := filepath.Join(tempDir, "tpm-socket")

	mac, err := generateRandomEthernetMAC()
	if err != nil {
		return err
	}

	qemuArgs := []string{"-machine", "q35", "-accel", "kvm", "-nographic", "-nodefaults", "-m", "4096",
		"-cpu", "host", "-smp", "sockets=1,cpus=1,cores=2,threads=2,maxcpus=4",
		"-drive", "if=pflash,format=raw,readonly,file=external/edk2/OVMF_CODE.fd",
		"-drive", "if=pflash,format=raw,snapshot=on,file=external/edk2/OVMF_VARS.fd",
		"-drive", "if=virtio,format=raw,snapshot=on,cache=unsafe,file=metropolis/node/node.img",
		"-netdev", qemuNetConfig.toOption(qemuNetType),
		"-device", "virtio-net-pci,netdev=net0,mac=" + mac.String(),
		"-chardev", "socket,id=chrtpm,path=" + tpmSocketPath,
		"-tpmdev", "emulator,id=tpm0,chardev=chrtpm",
		"-device", "tpm-tis,tpmdev=tpm0",
		"-device", "virtio-rng-pci",
		"-serial", "stdio"}

	if !options.AllowReboot {
		qemuArgs = append(qemuArgs, "-no-reboot")
	}

	if options.EnrolmentConfig != nil {
		enrolmentConfigPath := filepath.Join(tempDir, "enrolment.pb")
		enrolmentConfigRaw, err := proto.Marshal(options.EnrolmentConfig)
		if err != nil {
			return fmt.Errorf("failed to encode enrolment config: %w", err)
		}
		if err := ioutil.WriteFile(enrolmentConfigPath, enrolmentConfigRaw, 0644); err != nil {
			return fmt.Errorf("failed to write enrolment config: %w", err)
		}
		qemuArgs = append(qemuArgs, "-fw_cfg", "name=dev.monogon.metropolis/enrolment.pb,file="+enrolmentConfigPath)
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
	if options.ConnectToSocket != nil {
		systemCmd.ExtraFiles = []*os.File{options.ConnectToSocket}
	}

	var stdErrBuf bytes.Buffer
	systemCmd.Stderr = &stdErrBuf
	systemCmd.Stdout = options.SerialPort

	err = systemCmd.Run()

	// Stop TPM emulator and wait for it to exit to properly reap the child process
	tpmCancel()
	log.Print("Waiting for TPM emulator to exit")
	// Wait returns a SIGKILL error because we just cancelled its context.
	// We still need to call it to avoid creating zombies.
	_ = tpmEmuCmd.Wait()

	var exerr *exec.ExitError
	if err != nil && errors.As(err, &exerr) {
		status := exerr.ProcessState.Sys().(syscall.WaitStatus)
		if status.Signaled() && status.Signal() == syscall.SIGKILL {
			// Process was killed externally (most likely by our context being canceled).
			// This is a normal exit for us, so return nil
			return nil
		}
		exerr.Stderr = stdErrBuf.Bytes()
		newErr := QEMUError(*exerr)
		return &newErr
	}
	return err
}

// NewSocketPair creates a new socket pair. By connecting both ends to different instances you can connect them
// with a virtual "network cable". The ends can be passed into the ConnectToSocket option.
func NewSocketPair() (*os.File, *os.File, error) {
	fds, err := unix.Socketpair(unix.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to call socketpair: %w", err)
	}

	fd1 := os.NewFile(uintptr(fds[0]), "network0")
	fd2 := os.NewFile(uintptr(fds[1]), "network1")
	return fd1, fd2, nil
}

// HostInterfaceMAC is the MAC address the host SLIRP network interface has if it is not disabled (see
// DisableHostNetworkInterface in MicroVMOptions)
var HostInterfaceMAC = net.HardwareAddr{0x02, 0x72, 0x82, 0xbf, 0xc3, 0x56}

// MicroVMOptions contains all options to start a MicroVM
type MicroVMOptions struct {
	// Path to the ELF kernel binary
	KernelPath string

	// Path to the Initramfs
	InitramfsPath string

	// Cmdline contains additional kernel commandline options
	Cmdline string

	// SerialPort is a File(descriptor) over which you can communicate with the serial port of the machine
	// It can be set to an existing file descriptor (like os.Stdout/os.Stderr) or you can use NewSocketPair() to get one
	// end to talk to from Go.
	SerialPort *os.File

	// ExtraChardevs can be used similar to SerialPort, but can contain an arbitrary number of additional serial ports
	ExtraChardevs []*os.File

	// ExtraNetworkInterfaces can contain an arbitrary number of file descriptors which are mapped into the VM as virtio
	// network interfaces. The first interface is always a SLIRP-backed interface for communicating with the host.
	ExtraNetworkInterfaces []*os.File

	// PortMap contains ports that are mapped to the host through the built-in SLIRP network interface.
	PortMap PortMap

	// DisableHostNetworkInterface disables the SLIRP-backed host network interface that is normally the first network
	// interface. If this is set PortMap is ignored. Mostly useful for speeding up QEMU's startup time for tests.
	DisableHostNetworkInterface bool
}

// RunMicroVM launches a tiny VM mostly intended for testing. Very quick to boot (<40ms).
func RunMicroVM(ctx context.Context, opts *MicroVMOptions) error {
	// Generate options for all the file descriptors we'll be passing as virtio "serial ports"
	var extraArgs []string
	for idx, _ := range opts.ExtraChardevs {
		idxStr := strconv.Itoa(idx)
		id := "extra" + idxStr
		// That this works is pretty much a hack, but upstream QEMU doesn't have a bidirectional chardev backend not
		// based around files/sockets on the disk which are a giant pain to work with.
		// We're using QEMU's fdset functionality to make FDs available as pseudo-files and then "ab"using the pipe
		// backend's fallback functionality to get a single bidirectional chardev backend backed by a passed-down
		// RDWR fd.
		// Ref https://lists.gnu.org/archive/html/qemu-devel/2015-12/msg01256.html
		addFdConf := qemuValue{
			"set": {idxStr},
			"fd":  {strconv.Itoa(idx + 3)},
		}
		chardevConf := qemuValue{
			"id":   {id},
			"path": {"/dev/fdset/" + idxStr},
		}
		deviceConf := qemuValue{
			"chardev": {id},
		}
		extraArgs = append(extraArgs, "-add-fd", addFdConf.toOption(""),
			"-chardev", chardevConf.toOption("pipe"), "-device", deviceConf.toOption("virtserialport"))
	}

	for idx, _ := range opts.ExtraNetworkInterfaces {
		id := fmt.Sprintf("net%v", idx)
		netdevConf := qemuValue{
			"id": {id},
			"fd": {strconv.Itoa(idx + 3 + len(opts.ExtraChardevs))},
		}
		extraArgs = append(extraArgs, "-netdev", netdevConf.toOption("socket"), "-device", "virtio-net-device,netdev="+id)
	}

	// This sets up a minimum viable environment for our Linux kernel.
	// It clears all standard QEMU configuration and sets up a MicroVM machine
	// (https://github.com/qemu/qemu/blob/master/docs/microvm.rst) with all legacy emulation turned off. This means
	// the only "hardware" the Linux kernel inside can communicate with is a single virtio-mmio region. Over that MMIO
	// interface we run a paravirtualized RNG (since the kernel in there has nothing to gather that from and it
	// delays booting), a single paravirtualized console and an arbitrary number of extra serial ports for talking to
	// various things that might run inside. The kernel, initramfs and command line are mapped into VM memory at boot
	// time and not loaded from any sort of disk. Booting and shutting off one of these VMs takes <100ms.
	baseArgs := []string{"-nodefaults", "-no-user-config", "-nographic", "-no-reboot",
		"-accel", "kvm", "-cpu", "host",
		// Needed until QEMU updates their bundled qboot version (needs https://github.com/bonzini/qboot/pull/28)
		"-bios", "external/com_github_bonzini_qboot/bios.bin",
		"-M", "microvm,x-option-roms=off,pic=off,pit=off,rtc=off,isa-serial=off",
		"-kernel", opts.KernelPath,
		// We force using a triple-fault reboot strategy since otherwise the kernel first tries others (like ACPI) which
		// are not available in this very restricted environment. Similarly we need to override the boot console since
		// there's nothing on the ISA bus that the kernel could talk to. We also force quiet for performance reasons.
		"-append", "reboot=t console=hvc0 quiet " + opts.Cmdline,
		"-initrd", opts.InitramfsPath,
		"-device", "virtio-rng-device,max-bytes=1024,period=1000",
		"-device", "virtio-serial-device,max_ports=16",
		"-chardev", "stdio,id=con0", "-device", "virtconsole,chardev=con0",
	}

	if !opts.DisableHostNetworkInterface {
		qemuNetType := "user"
		qemuNetConfig := qemuValue{
			"id":        {"usernet0"},
			"net":       {"10.42.0.0/24"},
			"dhcpstart": {"10.42.0.10"},
		}
		if opts.PortMap != nil {
			qemuNetConfig["hostfwd"] = opts.PortMap.toQemuForwards()
		}

		baseArgs = append(baseArgs, "-netdev", qemuNetConfig.toOption(qemuNetType),
			"-device", "virtio-net-device,netdev=usernet0,mac="+HostInterfaceMAC.String())
	}

	var stdErrBuf bytes.Buffer
	cmd := exec.CommandContext(ctx, "qemu-system-x86_64", append(baseArgs, extraArgs...)...)
	cmd.Stdout = opts.SerialPort
	cmd.Stderr = &stdErrBuf

	cmd.ExtraFiles = append(cmd.ExtraFiles, opts.ExtraChardevs...)
	cmd.ExtraFiles = append(cmd.ExtraFiles, opts.ExtraNetworkInterfaces...)

	err := cmd.Run()
	var exerr *exec.ExitError
	if err != nil && errors.As(err, &exerr) {
		exerr.Stderr = stdErrBuf.Bytes()
		newErr := QEMUError(*exerr)
		return &newErr
	}
	return err
}

// QEMUError is a special type of ExitError used when QEMU fails. In addition to normal ExitError features it
// prints stderr for debugging.
type QEMUError exec.ExitError

func (e *QEMUError) Error() string {
	return fmt.Sprintf("%v: %v", e.String(), string(e.Stderr))
}

// NanoswitchPorts contains all ports forwarded by Nanoswitch to the first VM
var NanoswitchPorts = []uint16{
	node.ExternalServicePort,
	node.DebugServicePort,
	node.KubernetesAPIPort,
}

// ClusterOptions contains all options for launching a Metropolis cluster
type ClusterOptions struct {
	// The number of nodes this cluster should be started with initially
	NumNodes int
}

// LaunchCluster launches a cluster of Metropolis node VMs together with a Nanoswitch instance to network them all together.
func LaunchCluster(ctx context.Context, opts ClusterOptions) (apb.NodeDebugServiceClient, PortMap, error) {
	var switchPorts []*os.File
	var vmPorts []*os.File
	for i := 0; i < opts.NumNodes; i++ {
		switchPort, vmPort, err := NewSocketPair()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get socketpair: %w", err)
		}
		switchPorts = append(switchPorts, switchPort)
		vmPorts = append(vmPorts, vmPort)
	}

	if opts.NumNodes == 0 {
		return nil, nil, errors.New("refusing to start cluster with zero nodes")
	}

	if opts.NumNodes > 2 {
		return nil, nil, errors.New("launching more than 2 nodes is unsupported pending replacement of golden tickets")
	}

	go func() {
		if err := Launch(ctx, Options{ConnectToSocket: vmPorts[0]}); err != nil {
			// Launch() only terminates when QEMU has terminated. At that point our function probably doesn't run anymore
			// so we have no way of communicating the error back up, so let's just log it. Also a failure in launching
			// VMs should be very visible by the unavailability of the clients we return.
			log.Printf("Failed to launch vm0: %v", err)
		}
	}()

	portMap, err := ConflictFreePortMap(NanoswitchPorts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to allocate ephemeral ports: %w", err)
	}

	go func() {
		if err := RunMicroVM(ctx, &MicroVMOptions{
			KernelPath:             "metropolis/test/ktest/linux-testing.elf",
			InitramfsPath:          "metropolis/test/nanoswitch/initramfs.lz4",
			ExtraNetworkInterfaces: switchPorts,
			PortMap:                portMap,
		}); err != nil {
			log.Printf("Failed to launch nanoswitch: %v", err)
		}
	}()
	copts := []grpcretry.CallOption{
		grpcretry.WithBackoff(grpcretry.BackoffExponential(100 * time.Millisecond)),
	}
	conn, err := portMap.DialGRPC(node.DebugServicePort, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(copts...)))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial debug service: %w", err)
	}
	debug := apb.NewNodeDebugServiceClient(conn)

	if opts.NumNodes == 2 {
		res, err := debug.GetGoldenTicket(ctx, &apb.GetGoldenTicketRequest{
			// HACK: this is assigned by DHCP, and we assume that everything goes well.
			ExternalIp: "10.1.0.3",
		}, grpcretry.WithMax(10))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get golden ticket: %w", err)
		}

		ec := &apb.EnrolmentConfig{
			GoldenTicket: res.Ticket,
		}

		go func() {
			if err := Launch(ctx, Options{ConnectToSocket: vmPorts[1], EnrolmentConfig: ec}); err != nil {
				log.Printf("Failed to launch vm1: %v", err)
			}
		}()
	}

	return debug, portMap, nil
}
