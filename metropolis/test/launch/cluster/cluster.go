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
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/multierr"
	"golang.org/x/net/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node"
	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
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
var NodePorts = []node.Port{
	node.ConsensusPort,

	node.CuratorServicePort,
	node.DebugServicePort,

	node.KubernetesAPIPort,
	node.KubernetesAPIWrappedPort,
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
	tempDir, err := os.MkdirTemp("/tmp", "launch*")
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
	tpmFiles, err := os.ReadDir(tpmSrcDir)
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
		if err := os.WriteFile(parametersPath, parametersRaw, 0644); err != nil {
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

// getNodes wraps around Management.GetNodes to return a list of nodes in a
// cluster.
func getNodes(ctx context.Context, mgmt apb.ManagementClient) ([]*apb.Node, error) {
	var res []*apb.Node
	bo := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	err := backoff.Retry(func() error {
		res = nil
		srvN, err := mgmt.GetNodes(ctx, &apb.GetNodesRequest{})
		if err != nil {
			return fmt.Errorf("GetNodes: %w", err)
		}
		for {
			node, err := srvN.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("GetNodes.Recv: %w", err)
			}
			res = append(res, node)
		}
		return nil
	}, bo)
	if err != nil {
		return nil, err
	}
	return res, nil
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

const SOCKSPort uint16 = 1080

// ClusterPorts contains all ports handled by Nanoswitch.
var ClusterPorts = []uint16{
	// Forwarded to the first node.
	uint16(node.CuratorServicePort),
	uint16(node.DebugServicePort),
	uint16(node.KubernetesAPIPort),
	uint16(node.KubernetesAPIWrappedPort),

	// SOCKS proxy to the switch network
	SOCKSPort,
}

// ClusterOptions contains all options for launching a Metropolis cluster.
type ClusterOptions struct {
	// The number of nodes this cluster should be started with.
	NumNodes int
}

// Cluster is the running Metropolis cluster launched using the LaunchCluster
// function.
type Cluster struct {
	// Owner is the TLS Certificate of the owner of the test cluster. This can be
	// used to authenticate further clients to the running cluster.
	Owner tls.Certificate
	// Ports is the PortMap used to access the first nodes' services (defined in
	// ClusterPorts) and the SOCKS proxy (at SOCKSPort).
	Ports launch.PortMap

	// Nodes is a map from Node ID to its runtime information.
	Nodes map[string]*NodeInCluster
	// NodeIDs is a list of node IDs that are backing this cluster, in order of
	// creation.
	NodeIDs []string

	// nodesDone is a list of channels populated with the return codes from all the
	// nodes' qemu instances. It's used by Close to ensure all nodes have
	// succesfully been stopped.
	nodesDone []chan error
	// ctxC is used by Close to cancel the context under which the nodes are
	// running.
	ctxC context.CancelFunc
	// socksDialer is used by DialNode to establish connections to nodes via the
	// SOCKS server ran by nanoswitch.
	socksDialer proxy.Dialer
}

// NodeInCluster represents information about a node that's part of a Cluster.
type NodeInCluster struct {
	// ID of the node, which can be used to dial this node's services via DialNode.
	ID string
	// Address of the node on the network ran by nanoswitch. Not reachable from the
	// host unless dialed via DialNode or via the nanoswitch SOCKS proxy (reachable
	// on Cluster.Ports[SOCKSPort]).
	ManagementAddress string
}

// firstConnection performs the initial owner credential escrow with a newly
// started nanoswitch-backed cluster over SOCKS. It expects the first node to be
// running at 10.1.0.2, which is always the case with the current nanoswitch
// implementation.
//
// It returns the newly escrowed credentials as well as the firt node's
// information as NodeInCluster.
func firstConnection(ctx context.Context, socksDialer proxy.Dialer) (*tls.Certificate, *NodeInCluster, error) {
	// Dial external service.
	remote := fmt.Sprintf("10.1.0.2:%s", node.CuratorServicePort.PortString())
	initCreds, err := rpc.NewEphemeralCredentials(InsecurePrivateKey, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("NewEphemeralCredentials: %w", err)
	}
	initDialer := func(_ context.Context, addr string) (net.Conn, error) {
		return socksDialer.Dial("tcp", addr)
	}
	initClient, err := grpc.Dial(remote, grpc.WithContextDialer(initDialer), grpc.WithTransportCredentials(initCreds))
	if err != nil {
		return nil, nil, fmt.Errorf("dialing with ephemeral credentials failed: %w", err)
	}
	defer initClient.Close()

	// Retrieve owner certificate - this can take a while because the node is still
	// coming up, so do it in a backoff loop.
	log.Printf("Cluster: retrieving owner certificate (this can take a few seconds while the first node boots)...")
	aaa := apb.NewAAAClient(initClient)
	var cert *tls.Certificate
	err = backoff.Retry(func() error {
		cert, err = rpc.RetrieveOwnerCertificate(ctx, aaa, InsecurePrivateKey)
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.Unavailable {
				return err
			}
		}
		return backoff.Permanent(err)
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		return nil, nil, err
	}
	log.Printf("Cluster: retrieved owner certificate.")

	// Now connect authenticated and get the node ID.
	creds := rpc.NewAuthenticatedCredentials(*cert, nil)
	authClient, err := grpc.Dial(remote, grpc.WithContextDialer(initDialer), grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, nil, fmt.Errorf("dialing with owner credentials failed: %w", err)
	}
	defer authClient.Close()
	mgmt := apb.NewManagementClient(authClient)

	var node *NodeInCluster
	err = backoff.Retry(func() error {
		nodes, err := getNodes(ctx, mgmt)
		if err != nil {
			return fmt.Errorf("retrieving nodes failed: %w", err)
		}
		if len(nodes) != 1 {
			return fmt.Errorf("expected one node, got %d", len(nodes))
		}
		n := nodes[0]
		if n.Status == nil || n.Status.ExternalAddress == "" {
			return fmt.Errorf("node has no status and/or address")
		}
		node = &NodeInCluster{
			ID:                identity.NodeID(n.Pubkey),
			ManagementAddress: n.Status.ExternalAddress,
		}
		return nil
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		return nil, nil, err
	}

	return cert, node, nil
}

// LaunchCluster launches a cluster of Metropolis node VMs together with a
// Nanoswitch instance to network them all together.
//
// The given context will be used to run all qemu instances in the cluster, and
// canceling the context or calling Close() will terminate them.
func LaunchCluster(ctx context.Context, opts ClusterOptions) (*Cluster, error) {
	if opts.NumNodes <= 0 {
		return nil, errors.New("refusing to start cluster with zero nodes")
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

	// Make a list of channels that will be populated by all running node qemu
	// processes.
	done := make([]chan error, opts.NumNodes)
	for i, _ := range done {
		done[i] = make(chan error, 1)
	}

	// Start first node.
	log.Printf("Cluster: Starting node %d...", 1)
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
			SerialPort: newPrefixedStdio(0),
		})
		done[0] <- err
	}()

	// Launch nanoswitch.
	portMap, err := launch.ConflictFreePortMap(ClusterPorts)
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("failed to allocate ephemeral ports: %w", err)
	}

	go func() {
		if err := launch.RunMicroVM(ctxT, &launch.MicroVMOptions{
			KernelPath:             "metropolis/test/ktest/vmlinux",
			InitramfsPath:          "metropolis/test/nanoswitch/initramfs.cpio.lz4",
			ExtraNetworkInterfaces: switchPorts,
			PortMap:                portMap,
		}); err != nil {
			if !errors.Is(err, ctxT.Err()) {
				log.Fatalf("Failed to launch nanoswitch: %v", err)
			}
		}
	}()

	// Build SOCKS dialer.
	socksRemote := fmt.Sprintf("localhost:%v", portMap[SOCKSPort])
	socksDialer, err := proxy.SOCKS5("tcp", socksRemote, nil, proxy.Direct)
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("failed to build SOCKS dialer: %w", err)
	}

	// Retrieve owner credentials and first node.
	cert, firstNode, err := firstConnection(ctxT, socksDialer)
	if err != nil {
		ctxC()
		return nil, err
	}

	cluster := &Cluster{
		Owner: *cert,
		Ports: portMap,
		Nodes: map[string]*NodeInCluster{
			firstNode.ID: firstNode,
		},
		NodeIDs: []string{
			firstNode.ID,
		},

		nodesDone:   done,
		socksDialer: socksDialer,

		ctxC: ctxC,
	}

	// Now start the rest of the nodes and register them into the cluster.

	// Build authenticated owner client to first node.
	authCreds := rpc.NewAuthenticatedCredentials(*cert, nil)
	remote := net.JoinHostPort(cluster.NodeIDs[0], common.CuratorServicePort.PortString())
	authClient, err := grpc.Dial(remote, grpc.WithTransportCredentials(authCreds), grpc.WithContextDialer(cluster.DialNode))
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("dialing with owner credentials failed: %w", err)
	}
	defer authClient.Close()
	mgmt := apb.NewManagementClient(authClient)

	// Retrieve register ticket to register further nodes.
	log.Printf("Cluster: retrieving register ticket...")
	resT, err := mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("GetRegisterTicket: %w", err)
	}
	ticket := resT.Ticket
	log.Printf("Cluster: retrieved register ticket (%d bytes).", len(ticket))

	// Retrieve cluster info (for directory and ca public key) to register further
	// nodes.
	resI, err := mgmt.GetClusterInfo(ctx, &apb.GetClusterInfoRequest{})
	if err != nil {
		ctxC()
		return nil, fmt.Errorf("GetClusterInfo: %w", err)
	}

	// TODO(q3k): parallelize this
	for i := 1; i < opts.NumNodes; i++ {
		log.Printf("Cluster: Starting node %d...", i+1)
		go func(i int) {
			err := LaunchNode(ctxT, NodeOptions{
				ConnectToSocket: vmPorts[i],
				NodeParameters: &apb.NodeParameters{
					Cluster: &apb.NodeParameters_ClusterRegister_{
						ClusterRegister: &apb.NodeParameters_ClusterRegister{
							RegisterTicket:   ticket,
							ClusterDirectory: resI.ClusterDirectory,
							CaCertificate:    resI.CaCertificate,
						},
					},
				},
				SerialPort: newPrefixedStdio(i),
			})
			done[i] <- err
		}(i)
		var newNode *apb.Node

		log.Printf("Cluster: waiting for node %d to appear as NEW...", i)
		for {
			nodes, err := getNodes(ctx, mgmt)
			if err != nil {
				ctxC()
				return nil, fmt.Errorf("could not get nodes: %w", err)
			}
			for _, n := range nodes {
				if n.State == cpb.NodeState_NODE_STATE_NEW {
					newNode = n
					break
				}
			}
			if newNode != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		id := identity.NodeID(newNode.Pubkey)
		log.Printf("Cluster: node %d is %s", i, id)

		log.Printf("Cluster: approving node %d", i)
		_, err := mgmt.ApproveNode(ctx, &apb.ApproveNodeRequest{
			Pubkey: newNode.Pubkey,
		})
		if err != nil {
			ctxC()
			return nil, fmt.Errorf("ApproveNode(%s): %w", id, err)
		}
		log.Printf("Cluster: node %d approved, waiting for it to appear as UP and with a network address...", i)
		for {
			nodes, err := getNodes(ctx, mgmt)
			if err != nil {
				ctxC()
				return nil, fmt.Errorf("could not get nodes: %w", err)
			}
			found := false
			for _, n := range nodes {
				if !bytes.Equal(n.Pubkey, newNode.Pubkey) {
					continue
				}
				if n.Status == nil || n.Status.ExternalAddress == "" {
					break
				}
				if n.State != cpb.NodeState_NODE_STATE_UP {
					break
				}
				found = true
				cluster.Nodes[identity.NodeID(n.Pubkey)] = &NodeInCluster{
					ID:                identity.NodeID(n.Pubkey),
					ManagementAddress: n.Status.ExternalAddress,
				}
				cluster.NodeIDs = append(cluster.NodeIDs, identity.NodeID(n.Pubkey))
				break
			}
			if found {
				break
			}
			time.Sleep(time.Second)
		}
		log.Printf("Cluster: node %d (%s) UP!", i, id)
	}

	log.Printf("Cluster: all nodes up:")
	for _, node := range cluster.Nodes {
		log.Printf("Cluster:  - %s at %s", node.ID, node.ManagementAddress)
	}

	return cluster, nil
}

// Close cancels the running cluster's context and waits for all virtualized
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

// DialNode is a grpc.WithContextDialer compatible dialer which dials nodes by
// their ID. This is performed by connecting to the cluster nanoswitch via its
// SOCKS proxy, and using the cluster node list for name resolution.
//
// For example:
//
//   grpc.Dial("metropolis-deadbeef:1234", grpc.WithContextDialer(c.DialNode))
//
func (c *Cluster) DialNode(_ context.Context, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid host:port: %w", err)
	}
	// Already an IP address?
	if net.ParseIP(host) != nil {
		return c.socksDialer.Dial("tcp", addr)
	}

	// Otherwise, expect a node name.
	node, ok := c.Nodes[host]
	if !ok {
		return nil, fmt.Errorf("unknown node %q", host)
	}
	addr = net.JoinHostPort(node.ManagementAddress, port)
	return c.socksDialer.Dial("tcp", addr)
}
