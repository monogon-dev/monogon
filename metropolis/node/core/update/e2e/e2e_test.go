package e2e

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"source.monogon.dev/metropolis/cli/pkg/datafile"
	"source.monogon.dev/metropolis/node/build/mkimage/osimage"
	"source.monogon.dev/metropolis/pkg/blkio"
	"source.monogon.dev/metropolis/pkg/blockdev"
)

const Mi = 1024 * 1024

var variantRegexp = regexp.MustCompile(`TESTOS_VARIANT=([A-Z])`)

func runAndCheckVariant(t *testing.T, expectedVariant string, qemuArgs []string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	qemuCmdLaunch := exec.CommandContext(ctx, "qemu-system-x86_64", qemuArgs...)
	stdoutPipe, err := qemuCmdLaunch.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	stderrPipe, err := qemuCmdLaunch.StderrPipe()
	if err != nil {
		t.Fatal(err)
	}
	testosStarted := make(chan string, 1)
	go func() {
		s := bufio.NewScanner(stdoutPipe)
		for s.Scan() {
			if strings.HasPrefix(s.Text(), "[") {
				continue
			}
			errIdx := strings.Index(s.Text(), "Error installing new bundle")
			if errIdx != -1 {
				t.Error(s.Text()[errIdx:])
			}
			t.Log("vm: " + s.Text())
			if m := variantRegexp.FindStringSubmatch(s.Text()); len(m) == 2 {
				select {
				case testosStarted <- m[1]:
				default:
				}
			}
		}
	}()
	go func() {
		s := bufio.NewScanner(stderrPipe)
		for s.Scan() {
			if strings.HasPrefix(s.Text(), "[") {
				continue
			}
			t.Log("qemu: " + s.Text())
		}
	}()
	if err := qemuCmdLaunch.Start(); err != nil {
		t.Fatal(err)
	}
	procExit := make(chan error)
	go func() {
		procExit <- qemuCmdLaunch.Wait()
		close(procExit)
	}()
	select {
	case variant := <-testosStarted:
		if variant != expectedVariant {
			t.Fatalf("expected variant %s to launch, got %s", expectedVariant, variant)
		}
		select {
		case <-procExit:
			return
		case <-ctx.Done():
			t.Log("Canceled VM")
			cancel()
			<-procExit
			return
		}
	case err := <-procExit:
		t.Fatalf("QEMU exited unexpectedly: %v", err)
		return
	case <-ctx.Done():
		t.Fatalf("Waiting for TestOS variant %s launch timed out", expectedVariant)
	}
}

func TestABUpdateSequence(t *testing.T) {
	blobAddr := net.TCPAddr{
		IP:   net.IPv4(10, 42, 0, 5),
		Port: 80,
	}

	var nextBundlePathToInstall string
	var nbpMutex sync.Mutex

	m := http.NewServeMux()
	bundleYPath, err := datafile.ResolveRunfile("metropolis/node/core/update/e2e/testos/testos_bundle_y.zip")
	if err != nil {
		t.Fatal(err)
	}
	bundleZPath, err := datafile.ResolveRunfile("metropolis/node/core/update/e2e/testos/testos_bundle_z.zip")
	if err != nil {
		t.Fatal(err)
	}
	m.HandleFunc("/bundle.bin", func(w http.ResponseWriter, req *http.Request) {
		nbpMutex.Lock()
		bundleFilePath := nextBundlePathToInstall
		nbpMutex.Unlock()
		if bundleFilePath == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No next bundle set in the test harness"))
		}
		http.ServeFile(w, req, bundleFilePath)
	})
	blobLis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	blobListenAddr := blobLis.Addr().(*net.TCPAddr)
	go http.Serve(blobLis, m)

	rootDevPath := filepath.Join(t.TempDir(), "root.img")
	// Make a 512 bytes * 2Mi = 1Gi file-backed block device
	rootDisk, err := blockdev.CreateFile(rootDevPath, 512, 2097152)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(rootDevPath)
	defer rootDisk.Close()

	ovmfVarsPath, err := datafile.ResolveRunfile("external/edk2/OVMF_VARS.fd")
	if err != nil {
		t.Fatal(err)
	}
	ovmfCodePath, err := datafile.ResolveRunfile("external/edk2/OVMF_CODE.fd")
	if err != nil {
		t.Fatal(err)
	}
	bootPath, err := datafile.ResolveRunfile("metropolis/node/core/update/e2e/testos/kernel_efi_x.efi")
	if err != nil {
		t.Fatal(err)
	}
	boot, err := blkio.NewFileReader(bootPath)
	if err != nil {
		t.Fatal(err)
	}
	defer boot.Close()
	systemXPath, err := datafile.ResolveRunfile("metropolis/node/core/update/e2e/testos/verity_rootfs_x.img")
	if err != nil {
		t.Fatal(err)
	}
	system, err := os.Open(systemXPath)
	if err != nil {
		t.Fatal(err)
	}
	defer system.Close()

	if _, err := osimage.Create(&osimage.Params{
		Output:      rootDisk,
		EFIPayload:  boot,
		SystemImage: system,
		PartitionSize: osimage.PartitionSizeInfo{
			ESP:    128,
			System: 256,
			Data:   10,
		},
	}); err != nil {
		t.Fatalf("unable to generate starting point image: %v", err)
	}
	rootDisk.Close()

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
		"-cpu", "max", "-smp", "sockets=1,cpus=1,cores=2,threads=2,maxcpus=4",
		"-drive", "if=pflash,format=raw,readonly=on,file=" + ovmfCodePath,
		"-drive", "if=pflash,format=raw,file=" + ovmfVars.Name(),
		"-drive", "if=virtio,format=raw,cache=unsafe,file=" + rootDevPath,
		"-netdev", fmt.Sprintf("user,id=net0,net=10.42.0.0/24,dhcpstart=10.42.0.10,%s", blobGuestFwd),
		"-device", "virtio-net-pci,netdev=net0,mac=22:d5:8e:76:1d:07",
		"-device", "virtio-rng-pci",
		"-serial", "stdio",
		"-trace", "pflash*",
		"-no-reboot",
	}
	// Install Bundle Y next
	nbpMutex.Lock()
	nextBundlePathToInstall = bundleYPath
	nbpMutex.Unlock()

	t.Log("Launching X image to install Y")
	runAndCheckVariant(t, "X", qemuArgs)

	// Install Bundle Z next
	nbpMutex.Lock()
	nextBundlePathToInstall = bundleZPath
	nbpMutex.Unlock()

	t.Log("Launching Y on slot B to install Z on slot A")
	runAndCheckVariant(t, "Y", qemuArgs)

	t.Log("Launching Z on slot A")
	runAndCheckVariant(t, "Z", qemuArgs)
}
