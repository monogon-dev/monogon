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

	"github.com/bazelbuild/rules_go/go/runfiles"

	"source.monogon.dev/metropolis/node/build/mkimage/osimage"
	"source.monogon.dev/osbase/blkio"
	"source.monogon.dev/osbase/blockdev"
)

var (
	// These are filled by bazel at linking time with the canonical path of
	// their corresponding file. Inside the init function we resolve it
	// with the rules_go runfiles package to the real path.
	xBundleYPath  string
	xBundleZPath  string
	xOvmfVarsPath string
	xOvmfCodePath string
	xBootPath     string
	xSystemXPath  string
	xAbloaderPath string
)

func init() {
	var err error
	for _, path := range []*string{
		&xBundleYPath, &xBundleZPath, &xOvmfVarsPath,
		&xOvmfCodePath, &xBootPath, &xSystemXPath,
		&xAbloaderPath,
	} {
		*path, err = runfiles.Rlocation(*path)
		if err != nil {
			panic(err)
		}
	}
}

const Mi = 1024 * 1024

var variantRegexp = regexp.MustCompile(`TESTOS_VARIANT=([A-Z])`)

func stdoutHandler(t *testing.T, cmd *exec.Cmd, cancel context.CancelFunc, testosStarted chan string) {
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	s := bufio.NewScanner(stdoutPipe)
	go func() {
		for s.Scan() {
			if strings.HasPrefix(s.Text(), "[") {
				continue
			}
			errIdx := strings.Index(s.Text(), "Error installing new bundle")
			if errIdx != -1 {
				cancel()
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
}

func stderrHandler(t *testing.T, cmd *exec.Cmd) {
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		t.Fatal(err)
	}
	s := bufio.NewScanner(stderrPipe)
	go func() {
		for s.Scan() {
			if strings.HasPrefix(s.Text(), "[") {
				continue
			}
			t.Log("qemu: " + s.Text())
		}
	}()
}

func runAndCheckVariant(t *testing.T, expectedVariant string, qemuArgs []string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	qemuCmdLaunch := exec.CommandContext(ctx, "qemu-system-x86_64", qemuArgs...)
	testosStarted := make(chan string, 1)
	stdoutHandler(t, qemuCmdLaunch, cancel, testosStarted)
	stderrHandler(t, qemuCmdLaunch)
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
		case <-ctx.Done():
			t.Fatal("Timed out waiting for VM to exit")
		}
	case err := <-procExit:
		t.Fatalf("QEMU exited unexpectedly: %v", err)
	case <-ctx.Done():
		t.Fatalf("Waiting for TestOS variant %s launch timed out", expectedVariant)
	}
}

type bundleServing struct {
	t              *testing.T
	bundlePaths    map[string]string
	bundleFilePath string
	// Protects bundleFilePath above
	m sync.Mutex
}

func (b *bundleServing) setNextBundle(variant string) {
	b.m.Lock()
	defer b.m.Unlock()
	p, ok := b.bundlePaths[variant]
	if !ok {
		b.t.Fatalf("no bundle for variant %s available", variant)
	}
	b.bundleFilePath = p
}

// setup sets up an an HTTP server for serving bundles which can be controlled
// through the returned bundleServing struct as well as the initial boot disk
// and EFI variable storage. It also returns the required QEMU arguments to
// boot the initial TestOS.
func setup(t *testing.T) (*bundleServing, []string) {
	t.Helper()
	blobAddr := net.TCPAddr{
		IP:   net.IPv4(10, 42, 0, 5),
		Port: 80,
	}

	b := bundleServing{
		t:           t,
		bundlePaths: make(map[string]string),
	}

	m := http.NewServeMux()
	b.bundlePaths["Y"] = xBundleYPath
	b.bundlePaths["Z"] = xBundleZPath
	m.HandleFunc("/bundle.bin", func(w http.ResponseWriter, req *http.Request) {
		b.m.Lock()
		bundleFilePath := b.bundleFilePath
		b.m.Unlock()
		if bundleFilePath == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No next bundle set in the test harness"))
			return
		}
		http.ServeFile(w, req, bundleFilePath)
	})
	blobLis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { blobLis.Close() })
	blobListenAddr := blobLis.Addr().(*net.TCPAddr)
	go http.Serve(blobLis, m)

	rootDevPath := filepath.Join(t.TempDir(), "root.img")
	// Make a 512 bytes * 2Mi = 1Gi file-backed block device
	rootDisk, err := blockdev.CreateFile(rootDevPath, 512, 2097152)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Remove(rootDevPath) })
	defer rootDisk.Close()

	boot, err := blkio.NewFileReader(xBootPath)
	if err != nil {
		t.Fatal(err)
	}
	defer boot.Close()
	system, err := os.Open(xSystemXPath)
	if err != nil {
		t.Fatal(err)
	}
	defer system.Close()

	loader, err := blkio.NewFileReader(xAbloaderPath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := osimage.Create(&osimage.Params{
		Output:      rootDisk,
		ABLoader:    loader,
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

	blobGuestFwd := fmt.Sprintf("guestfwd=tcp:%s-tcp:127.0.0.1:%d", blobAddr.String(), blobListenAddr.Port)

	ovmfVars, err := os.CreateTemp("", "ab-ovmf-vars")
	if err != nil {
		t.Fatal(err)
	}
	defer ovmfVars.Close()
	t.Cleanup(func() { os.Remove(ovmfVars.Name()) })
	ovmfVarsTmpl, err := os.Open(xOvmfVarsPath)
	if err != nil {
		t.Fatal(err)
	}
	defer ovmfVarsTmpl.Close()
	if _, err := io.Copy(ovmfVars, ovmfVarsTmpl); err != nil {
		t.Fatal(err)
	}

	qemuArgs := []string{
		"-machine", "q35", "-accel", "kvm", "-nographic", "-nodefaults", "-m", "1024",
		"-cpu", "max", "-smp", "sockets=1,cpus=1,cores=2,threads=2,maxcpus=4",
		"-drive", "if=pflash,format=raw,readonly=on,file=" + xOvmfCodePath,
		"-drive", "if=pflash,format=raw,file=" + ovmfVars.Name(),
		"-drive", "if=virtio,format=raw,cache=unsafe,file=" + rootDevPath,
		"-netdev", fmt.Sprintf("user,id=net0,net=10.42.0.0/24,dhcpstart=10.42.0.10,%s", blobGuestFwd),
		"-device", "virtio-net-pci,netdev=net0,mac=22:d5:8e:76:1d:07",
		"-device", "virtio-rng-pci",
		"-serial", "stdio",
		"-no-reboot",
	}
	return &b, qemuArgs
}

func TestABUpdateSequenceReboot(t *testing.T) {
	bsrv, qemuArgs := setup(t)

	t.Log("Launching X image to install Y")
	bsrv.setNextBundle("Y")
	runAndCheckVariant(t, "X", qemuArgs)

	t.Log("Launching Y on slot B to install Z on slot A")
	bsrv.setNextBundle("Z")
	runAndCheckVariant(t, "Y", qemuArgs)

	t.Log("Launching Z on slot A")
	runAndCheckVariant(t, "Z", qemuArgs)
}

func TestABUpdateSequenceKexec(t *testing.T) {
	bsrv, qemuArgs := setup(t)
	qemuArgs = append(qemuArgs, "-fw_cfg", "name=use_kexec,string=1")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	qemuCmdLaunch := exec.CommandContext(ctx, "qemu-system-x86_64", qemuArgs...)
	testosStarted := make(chan string, 1)
	stdoutHandler(t, qemuCmdLaunch, cancel, testosStarted)
	stderrHandler(t, qemuCmdLaunch)
	if err := qemuCmdLaunch.Start(); err != nil {
		t.Fatal(err)
	}
	procExit := make(chan error)
	go func() {
		procExit <- qemuCmdLaunch.Wait()
		close(procExit)
	}()
	var expectedVariant = "X"
	for {
		select {
		case variant := <-testosStarted:
			if variant != expectedVariant {
				t.Fatalf("expected variant %s to launch, got %s", expectedVariant, variant)
			}
			switch expectedVariant {
			case "X":
				expectedVariant = "Y"
			case "Y":
				expectedVariant = "Z"
			case "Z":
				// We're done, wait for everything to wind down and return
				select {
				case <-procExit:
					return
				case <-ctx.Done():
					t.Error("Timed out waiting for VM to exit")
					cancel()
					<-procExit
					return
				}
			}
			bsrv.setNextBundle(expectedVariant)
			t.Logf("Got %s, installing %s", variant, expectedVariant)
		case err := <-procExit:
			t.Fatalf("QEMU exited unexpectedly: %v", err)
		case <-ctx.Done():
			t.Fatalf("Waiting for TestOS variant %s launch timed out", expectedVariant)
		}
	}
}
