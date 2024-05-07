package e2e

import (
	"bufio"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/cavaliergopher/cpio"
	"github.com/klauspost/compress/zstd"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/cloud/agent/api"
	bpb "source.monogon.dev/cloud/bmaas/server/api"
	mpb "source.monogon.dev/metropolis/proto/api"
	"source.monogon.dev/osbase/pki"
)

type fakeServer struct {
	hardwareReport      *bpb.AgentHardwareReport
	installationRequest *bpb.OSInstallationRequest
	installationReport  *bpb.OSInstallationReport
}

func (f *fakeServer) Heartbeat(ctx context.Context, req *bpb.AgentHeartbeatRequest) (*bpb.AgentHeartbeatResponse, error) {
	var res bpb.AgentHeartbeatResponse
	if req.HardwareReport != nil {
		f.hardwareReport = req.HardwareReport
	}
	if req.InstallationReport != nil {
		f.installationReport = req.InstallationReport
	}
	if f.installationRequest != nil {
		res.InstallationRequest = f.installationRequest
	}
	return &res, nil
}

const GiB = 1024 * 1024 * 1024

// TestMetropolisInstallE2E exercises the agent communicating against a test cloud
// API server. This server requests the installation of the Metropolis 'TestOS',
// which we then validate by looking for a string it outputs on boot.
func TestMetropolisInstallE2E(t *testing.T) {
	var f fakeServer

	// Address inside fake QEMU userspace networking
	grpcAddr := net.TCPAddr{
		IP:   net.IPv4(10, 42, 0, 5),
		Port: 3000,
	}

	blobAddr := net.TCPAddr{
		IP:   net.IPv4(10, 42, 0, 6),
		Port: 80,
	}

	f.installationRequest = &bpb.OSInstallationRequest{
		Generation: 5,
		Type: &bpb.OSInstallationRequest_Metropolis{Metropolis: &bpb.MetropolisInstallationRequest{
			BundleUrl:      (&url.URL{Scheme: "http", Host: blobAddr.String(), Path: "/bundle.bin"}).String(),
			NodeParameters: &mpb.NodeParameters{},
			RootDevice:     "vda",
		}},
	}

	caPubKey, caPrivKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	caCertTmpl := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "Agent E2E Test CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              pki.UnknownNotAfter,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}
	caCertRaw, err := x509.CreateCertificate(rand.Reader, &caCertTmpl, &caCertTmpl, caPubKey, caPrivKey)
	if err != nil {
		t.Fatal(err)
	}
	caCert, err := x509.ParseCertificate(caCertRaw)
	if err != nil {
		t.Fatal(err)
	}

	serverPubKey, serverPrivKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	serverCertTmpl := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{},
		NotBefore:             time.Now(),
		NotAfter:              pki.UnknownNotAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:           []net.IP{grpcAddr.IP},
		BasicConstraintsValid: true,
	}
	serverCert, err := x509.CreateCertificate(rand.Reader, &serverCertTmpl, caCert, serverPubKey, caPrivKey)
	if err != nil {
		t.Fatal(err)
	}

	s := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&tls.Certificate{
		Certificate: [][]byte{serverCert},
		PrivateKey:  serverPrivKey,
	})))
	bpb.RegisterAgentCallbackServer(s, &f)
	grpcLis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	go s.Serve(grpcLis)
	grpcListenAddr := grpcLis.Addr().(*net.TCPAddr)

	m := http.NewServeMux()
	bundleFilePath, err := runfiles.Rlocation("_main/metropolis/installer/test/testos/testos_bundle.zip")
	if err != nil {
		t.Fatal(err)
	}
	m.HandleFunc("/bundle.bin", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, bundleFilePath)
	})
	blobLis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	blobListenAddr := blobLis.Addr().(*net.TCPAddr)
	go http.Serve(blobLis, m)

	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	init := apb.AgentInit{
		TakeoverInit: &apb.TakeoverInit{
			MachineId:     "testbox1",
			BmaasEndpoint: grpcAddr.String(),
			CaCertificate: caCertRaw,
		},
		PrivateKey: privateKey,
	}

	rootDisk, err := os.CreateTemp("", "rootdisk")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(rootDisk.Name())
	// Create a 10GiB sparse root disk
	if err := unix.Ftruncate(int(rootDisk.Fd()), 10*GiB); err != nil {
		t.Fatalf("ftruncate failed: %v", err)
	}

	ovmfVarsPath, err := runfiles.Rlocation("edk2/OVMF_VARS.fd")
	if err != nil {
		t.Fatal(err)
	}
	ovmfCodePath, err := runfiles.Rlocation("edk2/OVMF_CODE.fd")
	if err != nil {
		t.Fatal(err)
	}
	kernelPath, err := runfiles.Rlocation("_main/third_party/linux/bzImage")
	if err != nil {
		t.Fatal(err)
	}
	initramfsOrigPath, err := runfiles.Rlocation("_main/cloud/agent/takeover/initramfs.cpio.zst")
	if err != nil {
		t.Fatal(err)
	}
	initramfsOrigFile, err := os.Open(initramfsOrigPath)
	if err != nil {
		t.Fatal(err)
	}
	defer initramfsOrigFile.Close()

	initramfsFile, err := os.CreateTemp("", "agent-initramfs")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(initramfsFile.Name())
	if _, err := initramfsFile.ReadFrom(initramfsOrigFile); err != nil {
		t.Fatal(err)
	}

	// Append AgentInit spec to initramfs
	agentInitRaw, err := proto.Marshal(&init)
	if err != nil {
		t.Fatal(err)
	}
	compressedW, err := zstd.NewWriter(initramfsFile, zstd.WithEncoderLevel(1))
	if err != nil {
		t.Fatal(err)
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

	grpcGuestFwd := fmt.Sprintf("guestfwd=tcp:%s-tcp:127.0.0.1:%d", grpcAddr.String(), grpcListenAddr.Port)
	blobGuestFwd := fmt.Sprintf("guestfwd=tcp:%s-tcp:127.0.0.1:%d", blobAddr.String(), blobListenAddr.Port)

	ovmfVars, err := os.CreateTemp("", "agent-ovmf-vars")
	if err != nil {
		t.Fatal(err)
	}
	ovmfVarsTmpl, err := os.Open(ovmfVarsPath)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.Copy(ovmfVars, ovmfVarsTmpl); err != nil {
		t.Fatal(err)
	}

	qemuArgs := []string{
		"-machine", "q35", "-accel", "kvm", "-nographic", "-nodefaults", "-m", "1024",
		"-cpu", "host", "-smp", "sockets=1,cpus=1,cores=2,threads=2,maxcpus=4",
		"-drive", "if=pflash,format=raw,readonly=on,file=" + ovmfCodePath,
		"-drive", "if=pflash,format=raw,file=" + ovmfVars.Name(),
		"-drive", "if=virtio,format=raw,cache=unsafe,file=" + rootDisk.Name(),
		"-netdev", fmt.Sprintf("user,id=net0,net=10.42.0.0/24,dhcpstart=10.42.0.10,%s,%s", grpcGuestFwd, blobGuestFwd),
		"-device", "virtio-net-pci,netdev=net0,mac=22:d5:8e:76:1d:07",
		"-device", "virtio-rng-pci",
		"-serial", "stdio",
		"-no-reboot",
	}
	stage1Args := append(qemuArgs,
		"-kernel", kernelPath,
		"-initrd", initramfsFile.Name(),
		"-append", "console=ttyS0 quiet")
	qemuCmdAgent := exec.Command("qemu-system-x86_64", stage1Args...)
	qemuCmdAgent.Stdout = os.Stdout
	qemuCmdAgent.Stderr = os.Stderr
	qemuCmdAgent.Run()
	qemuCmdLaunch := exec.Command("qemu-system-x86_64", qemuArgs...)
	stdoutPipe, err := qemuCmdLaunch.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	testosStarted := make(chan struct{})
	go func() {
		s := bufio.NewScanner(stdoutPipe)
		for s.Scan() {
			if strings.HasPrefix(s.Text(), "[") {
				continue
			}
			t.Log("vm: " + s.Text())
			if strings.Contains(s.Text(), "_TESTOS_LAUNCH_SUCCESS_") {
				testosStarted <- struct{}{}
				break
			}
		}
		qemuCmdLaunch.Wait()
	}()
	if err := qemuCmdLaunch.Start(); err != nil {
		t.Fatal(err)
	}
	defer qemuCmdLaunch.Process.Kill()
	select {
	case <-testosStarted:
		// Done, test passed
	case <-time.After(30 * time.Second):
		t.Fatal("Waiting for TestOS launch timed out")
	}
}
