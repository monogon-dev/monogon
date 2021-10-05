// cluster builds on the launch package and implements launching Metropolis
// nodes and clusters in a virtualized environment using qemu. It's kept in a
// separate package as it depends on a Metropolis node image, which might not be
// required for some use of the launch library.
package cluster

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/cenkalti/backoff/v4"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/rpc"
	apb "source.monogon.dev/metropolis/proto/api"
	"source.monogon.dev/metropolis/test/launch"
)

// Options contains all options that can be passed to Launch()
type NodeOptions struct {
	// Ports contains the port mapping where to expose the internal ports of the VM to
	// the host. See IdentityPortMap() and ConflictFreePortMap(). Ignored when
	// ConnectToSocket is set.
	Ports launch.PortMap

	// If set to true, reboots are honored. Otherwise all reboots exit the Launch()
	// command. Metropolis nodes generally restarts on almost all errors, so unless you
	// want to test reboot behavior this should be false.
	AllowReboot bool

	// By default the VM is connected to the Host via SLIRP. If ConnectToSocket is set,
	// it is instead connected to the given file descriptor/socket. If this is set, all
	// port maps from the Ports option are ignored. Intended for networking this
	// instance together with others for running  more complex network configurations.
	ConnectToSocket *os.File

	// SerialPort is a io.ReadWriter over which you can communicate with the serial
	// port of the machine It can be set to an existing file descriptor (like
	// os.Stdout/os.Stderr) or any Go structure implementing this interface.
	SerialPort io.ReadWriter

	// NodeParameters is passed into the VM and subsequently used for bootstrapping or
	// registering into a cluster.
	NodeParameters *apb.NodeParameters
}

// NodePorts is the list of ports a fully operational Metropolis node listens on
var NodePorts = []uint16{
	node.ConsensusPort,

	node.CuratorServicePort,
	node.DebugServicePort,

	node.KubernetesAPIPort,
	node.MasterServicePort,
	node.CuratorServicePort,
	node.DebuggerPort,
}

// LaunchNode launches a single Metropolis node instance with the given options.
// The instance runs mostly paravirtualized but with some emulated hardware
// similar to how a cloud provider might set up its VMs. The disk is fully
// writable but is run in snapshot mode meaning that changes are not kept beyond
// a single invocation.
func LaunchNode(ctx context.Context, options NodeOptions) error {
	// Pin temp directory to /tmp until we can use abstract socket namespace in QEMU
	// (next release after 5.0,
	// https://github.com/qemu/qemu/commit/776b97d3605ed0fc94443048fdf988c7725e38a9).
	// swtpm accepts already-open FDs so we can pass in an abstract socket namespace FD
	// that we open and pass the name of it to QEMU. Not pinning this crashes both
	// swtpm and qemu because we run into UNIX socket length limitations (for legacy
	// reasons 108 chars).
	tempDir, err := ioutil.TempDir("/tmp", "launch*")
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy TPM state into a temporary directory since it's being modified by the
	// emulator
	tpmTargetDir := filepath.Join(tempDir, "tpm")
	tpmSrcDir := "metropolis/node/tpm"
	if err := os.Mkdir(tpmTargetDir, 0755); err != nil {
		return fmt.Errorf("failed to create TPM state directory: %w", err)
	}
	tpmFiles, err := ioutil.ReadDir(tpmSrcDir)
	if err != nil {
		return fmt.Errorf("failed to read TPM directory: %w", err)
	}
	for _, file := range tpmFiles {
		name := file.Name()
		src := filepath.Join(tpmSrcDir, name)
		target := filepath.Join(tpmTargetDir, name)
		if err := copyFile(src, target); err != nil {
			return fmt.Errorf("failed to copy TPM directory: file %q to %q: %w", src, target, err)
		}
	}

	var qemuNetType string
	var qemuNetConfig launch.QemuValue
	if options.ConnectToSocket != nil {
		qemuNetType = "socket"
		qemuNetConfig = launch.QemuValue{
			"id": {"net0"},
			"fd": {"3"},
		}
	} else {
		qemuNetType = "user"
		qemuNetConfig = launch.QemuValue{
			"id":        {"net0"},
			"net":       {"10.42.0.0/24"},
			"dhcpstart": {"10.42.0.10"},
			"hostfwd":   options.Ports.ToQemuForwards(),
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
		"-netdev", qemuNetConfig.ToOption(qemuNetType),
		"-device", "virtio-net-pci,netdev=net0,mac=" + mac.String(),
		"-chardev", "socket,id=chrtpm,path=" + tpmSocketPath,
		"-tpmdev", "emulator,id=tpm0,chardev=chrtpm",
		"-device", "tpm-tis,tpmdev=tpm0",
		"-device", "virtio-rng-pci",
		"-serial", "stdio"}

	if !options.AllowReboot {
		qemuArgs = append(qemuArgs, "-no-reboot")
	}

	if options.NodeParameters != nil {
		parametersPath := filepath.Join(tempDir, "parameters.pb")
		parametersRaw, err := proto.Marshal(options.NodeParameters)
		if err != nil {
			return fmt.Errorf("failed to encode node paraeters: %w", err)
		}
		if err := ioutil.WriteFile(parametersPath, parametersRaw, 0644); err != nil {
			return fmt.Errorf("failed to write node parameters: %w", err)
		}
		qemuArgs = append(qemuArgs, "-fw_cfg", "name=dev.monogon.metropolis/parameters.pb,file="+parametersPath)
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
	log.Print("Node: Waiting for TPM emulator to exit")
	// Wait returns a SIGKILL error because we just cancelled its context.
	// We still need to call it to avoid creating zombies.
	_ = tpmEmuCmd.Wait()
	log.Print("Node: TPM emulator done")

	var exerr *exec.ExitError
	if err != nil && errors.As(err, &exerr) {
		status := exerr.ProcessState.Sys().(syscall.WaitStatus)
		if status.Signaled() && status.Signal() == syscall.SIGKILL {
			// Process was killed externally (most likely by our context being canceled).
			// This is a normal exit for us, so return nil
			return nil
		}
		exerr.Stderr = stdErrBuf.Bytes()
		newErr := launch.QEMUError(*exerr)
		return &newErr
	}
	return err
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("when opening source: %w", err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("when creating destination: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("when copying file: %w", err)
	}
	return out.Close()
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

// ClusterPorts contains all ports forwarded by Nanoswitch to the first VM in a
// launched Metropolis cluster.
var ClusterPorts = []uint16{
	node.CuratorServicePort,
	node.DebugServicePort,

	node.KubernetesAPIPort,
}

// ClusterOptions contains all options for launching a Metropolis cluster.
type ClusterOptions struct {
	// The number of nodes this cluster should be started with.
	NumNodes int
}

// Cluster is the running Metropolis cluster launched using the LaunchCluster
// function.
type Cluster struct {
	// Debug is the NodeDebugService gRPC service, allowing for debug
	// unauthenticated cluster to the access.
	Debug apb.NodeDebugServiceClient
	// Management is the Management gRPC service, authenticated as the owner of the
	// cluster.
	Management apb.ManagementClient
	// Owner is the TLS Certificate of the owner of the test cluster. This can be
	// used to authenticate further clients to the running cluster.
	Owner tls.Certificate
	// Ports is the PortMap used to access the first nodes' services (defined in
	// ClusterPorts).
	Ports launch.PortMap

	// nodesDone is a list of channels populated with the return codes from all the
	// nodes' qemu instances. It's used by Close to ensure all nodes have
	// succesfully been stopped.
	nodesDone []chan error
	// ctxC is used by Close to cancel the context under which the nodes are
	// running.
	ctxC context.CancelFunc
}

// LaunchCluster launches a cluster of Metropolis node VMs together with a
// Nanoswitch instance to network them all together.
//
// The given context will be used to run all qemu instances in the cluster, and
// canceling the context or calling Close() will terminate them.
func LaunchCluster(ctx context.Context, opts ClusterOptions) (*Cluster, error) {
	if opts.NumNodes == 0 {
		return nil, errors.New("refusing to start cluster with zero nodes")
	}
	if opts.NumNodes > 1 {
		return nil, errors.New("unimplemented")
	}

	ctxT, ctxC := context.WithCancel(ctx)

	// Prepare links between nodes and nanoswitch.
	var switchPorts []*os.File
	var vmPorts []*os.File
	for i := 0; i < opts.NumNodes; i++ {
		switchPort, vmPort, err := launch.NewSocketPair()
		if err != nil {
			ctxC()
			return nil, fmt.Errorf("failed to get socketpair: %w", err)
		}
		switchPorts = append(switchPorts, switchPort)
		vmPorts = append(vmPorts, vmPort)
	}

	// Make a list of channels that will be populated and closed by all running node
	// qemu processes.
	done := make([]chan error, opts.NumNodes)
	for i, _ := range done {
		done[i] = make(chan error, 1)
	}

	// Start first node.
	go func() {
		err := LaunchNode(ctxT, NodeOptions{
			ConnectToSocket: vmPorts[0],
			NodeParameters: &apb.NodeParameters{
				Cluster: &apb.NodeParameters_ClusterBootstrap_{
					ClusterBootstrap: &apb.NodeParameters_ClusterBootstrap{
						OwnerPublicKey: InsecurePublicKey,
					},
				},
			},
		})
		done[0] <- err
	}()

	portMap, err := launch.ConflictFreePortMap(ClusterPorts)
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("failed to allocate ephemeral ports: %w", err)
	}

	go func() {
		if err := launch.RunMicroVM(ctxT, &launch.MicroVMOptions{
			KernelPath:             "metropolis/test/ktest/vmlinux",
			InitramfsPath:          "metropolis/test/nanoswitch/initramfs.lz4",
			ExtraNetworkInterfaces: switchPorts,
			PortMap:                portMap,
		}); err != nil {
			if !errors.Is(err, ctxT.Err()) {
				log.Printf("Failed to launch nanoswitch: %v", err)
			}
		}
	}()

	// Dial debug service.
	copts := []grpcretry.CallOption{
		grpcretry.WithBackoff(grpcretry.BackoffExponential(100 * time.Millisecond)),
	}
	debugConn, err := portMap.DialGRPC(node.DebugServicePort, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(copts...)))
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("failed to dial debug service: %w", err)
	}

	// Dial external service.
	remote := fmt.Sprintf("localhost:%v", portMap[node.CuratorServicePort])
	initClient, err := rpc.NewEphemeralClient(remote, InsecurePrivateKey, nil)
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("NewInitialClient: %w", err)
	}

	// Retrieve owner certificate - this can take a while because the node is still
	// coming up, so do it in a backoff loop.
	log.Printf("Cluster: retrieving owner certificate...")
	aaa := apb.NewAAAClient(initClient)
	var cert *tls.Certificate
	err = backoff.Retry(func() error {
		cert, err = rpc.RetrieveOwnerCertificate(ctxT, aaa, InsecurePrivateKey)
		return err
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctxT))
	if err != nil {
		ctxC()
		return nil, err
	}
	log.Printf("Cluster: retrieved owner certificate.")

	authClient, err := rpc.NewAuthenticatedClient(remote, *cert, nil)
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("NewAuthenticatedClient: %w", err)
	}

	return &Cluster{
		Debug:      apb.NewNodeDebugServiceClient(debugConn),
		Management: apb.NewManagementClient(authClient),
		Owner:      *cert,
		Ports:      portMap,
		nodesDone:  done,

		ctxC: ctxC,
	}, nil
}

// Close cancels the running clusters' context and waits for all virtualized
// nodes to stop. It returns an error if stopping the nodes failed, or one of
// the nodes failed to fully start in the first place.
func (c *Cluster) Close() error {
	log.Printf("Cluster: stopping...")
	c.ctxC()

	var errors []error
	log.Printf("Cluster: waiting for nodes to exit...")
	for _, c := range c.nodesDone {
		err := <-c
		if err != nil {
			errors = append(errors, err)
		}
	}
	log.Printf("Cluster: done")
	return multierr.Combine(errors...)
}
