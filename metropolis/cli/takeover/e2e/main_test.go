// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"bufio"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	xssh "golang.org/x/crypto/ssh"
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/proto/api"

	"source.monogon.dev/go/net/ssh"
	"source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/osbase/fat32"
	"source.monogon.dev/osbase/freeport"
)

var (
	// These are filled by bazel at linking time with the canonical path of
	// their corresponding file. Inside the init function we resolve it
	// with the rules_go runfiles package to the real path.
	xBundleFilePath string
	xOvmfVarsPath   string
	xOvmfCodePath   string
	xCloudImagePath string
	xTakeoverPath   string
)

func init() {
	var err error
	for _, path := range []*string{
		&xCloudImagePath, &xOvmfVarsPath, &xOvmfCodePath,
		&xTakeoverPath, &xBundleFilePath,
	} {
		*path, err = runfiles.Rlocation(*path)
		if err != nil {
			panic(err)
		}
	}
}

const GiB = 1024 * 1024 * 1024

func TestE2E(t *testing.T) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	sshPubKey, err := xssh.NewPublicKey(pubKey)
	if err != nil {
		t.Fatal(err)
	}

	sshPrivkey, err := xssh.NewSignerFromKey(privKey)
	if err != nil {
		t.Fatal(err)
	}

	// CloudConfig doesn't really have a rigid spec, so just put things into it
	cloudConfig := make(map[string]any)
	cloudConfig["ssh_authorized_keys"] = []string{
		strings.TrimSuffix(string(xssh.MarshalAuthorizedKey(sshPubKey)), "\n"),
	}

	userData, err := json.Marshal(cloudConfig)
	if err != nil {
		t.Fatal(err)
	}

	rootInode := fat32.Inode{
		Attrs: fat32.AttrDirectory,
		Children: []*fat32.Inode{
			{
				Name:    "user-data",
				Content: strings.NewReader("#cloud-config\n" + string(userData)),
			},
			{
				Name:    "meta-data",
				Content: strings.NewReader(""),
			},
		},
	}
	cloudInitDataFile, err := os.CreateTemp("", "cidata*.img")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(cloudInitDataFile.Name())
	if err := fat32.WriteFS(cloudInitDataFile, rootInode, fat32.Options{Label: "cidata"}); err != nil {
		t.Fatal(err)
	}

	rootDisk, err := os.CreateTemp("", "rootdisk")
	if err != nil {
		t.Fatal(err)
	}
	// Create a 10GiB sparse root disk
	if err := unix.Ftruncate(int(rootDisk.Fd()), 10*GiB); err != nil {
		t.Fatalf("ftruncate failed: %v", err)
	}

	defer os.Remove(rootDisk.Name())

	sshPort, sshPortCloser, err := freeport.AllocateTCPPort()
	if err != nil {
		t.Fatal(err)
	}

	qemuArgs := []string{
		"-machine", "q35", "-accel", "kvm", "-nographic", "-nodefaults", "-m", "1024",
		"-cpu", "host", "-smp", "sockets=1,cpus=1,cores=2,threads=2,maxcpus=4",
		"-drive", "if=pflash,format=raw,readonly=on,file=" + xOvmfCodePath,
		"-drive", "if=pflash,format=raw,snapshot=on,file=" + xOvmfVarsPath,
		"-drive", "if=none,format=raw,cache=unsafe,id=root,file=" + rootDisk.Name(),
		"-drive", "if=none,format=qcow2,snapshot=on,id=cloud,cache=unsafe,file=" + xCloudImagePath,
		"-device", "virtio-blk-pci,drive=root,bootindex=1",
		"-device", "virtio-blk-pci,drive=cloud,bootindex=2",
		"-drive", "if=virtio,format=raw,snapshot=on,file=" + cloudInitDataFile.Name(),
		"-netdev", fmt.Sprintf("user,id=net0,net=10.42.0.0/24,dhcpstart=10.42.0.10,hostfwd=tcp::%d-:22", sshPort),
		"-device", "virtio-net-pci,netdev=net0,mac=22:d5:8e:76:1d:07",
		"-device", "virtio-rng-pci",
		"-serial", "stdio",
	}
	qemuCmd := exec.Command("qemu-system-x86_64", qemuArgs...)
	stdoutPipe, err := qemuCmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	installSucceed := make(chan struct{})
	go func() {
		s := bufio.NewScanner(stdoutPipe)
		for s.Scan() {
			t.Log("kernel: " + s.Text())
			if strings.Contains(s.Text(), "_TESTOS_LAUNCH_SUCCESS_") {
				installSucceed <- struct{}{}
				break
			}
		}
		qemuCmd.Wait()
	}()
	qemuCmd.Stderr = os.Stderr
	sshPortCloser.Close()
	if err := qemuCmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer qemuCmd.Process.Kill()

	cl := ssh.DirectClient{
		Username:    "debian",
		AuthMethods: []xssh.AuthMethod{xssh.PublicKeys(sshPrivkey)},
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	var conn ssh.Connection
	for {
		conn, err = cl.Dial(ctx, net.JoinHostPort("localhost", fmt.Sprintf("%d", sshPort)), 5*time.Second)
		if err != nil {
			t.Logf("error connecting via SSH, retrying: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	takeover, err := os.Open(xTakeoverPath)
	if err != nil {
		t.Fatal(err)
	}

	const takeoverTargetPath = "/tmp/takeover"
	if err := conn.Upload(ctx, takeoverTargetPath, takeover); err != nil {
		t.Fatalf("error while uploading takeover: %v", err)
	}

	bundleFile, err := os.Open(xBundleFilePath)
	if err != nil {
		t.Fatal(err)
	}

	const bundleTargetPath = "/tmp/bundle.zip"
	if err := conn.Upload(ctx, bundleTargetPath, bundleFile); err != nil {
		t.Fatalf("error while uploading bundle: %v", err)
	}

	params := &api.NodeParameters{
		Cluster: &api.NodeParameters_ClusterBootstrap_{
			ClusterBootstrap: &api.NodeParameters_ClusterBootstrap{
				OwnerPublicKey: launch.InsecurePublicKey,
			},
		},
		NetworkConfig: nil,
	}
	rawParams, err := proto.Marshal(params)
	if err != nil {
		t.Fatalf("error while marshaling node params: %v", err)
	}

	// Start the agent and wait for the agent's output to arrive.
	t.Logf("Starting the takeover executable at path %q.", takeoverTargetPath)
	_, stderr, err := conn.Execute(ctx, fmt.Sprintf("sudo %s -disk %s", takeoverTargetPath, "vda"), rawParams)
	stderrStr := strings.TrimSpace(string(stderr))
	if stderrStr != "" {
		t.Logf("Agent stderr: %q", stderrStr)
	}
	if err != nil {
		t.Fatalf("while starting the takeover executable: %v", err)
	}

	select {
	case <-installSucceed:
		// Done, test passed
	case <-time.After(30 * time.Second):
		t.Fatal("Waiting for installation timed out")
	}
}
