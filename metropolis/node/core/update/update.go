package update

import (
	"archive/zip"
	"bytes"
	"context"
	"debug/pe"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/build/mkimage/osimage"
	abloaderpb "source.monogon.dev/metropolis/node/core/abloader/spec"
	"source.monogon.dev/metropolis/pkg/blockdev"
	"source.monogon.dev/metropolis/pkg/efivarfs"
	"source.monogon.dev/metropolis/pkg/gpt"
	"source.monogon.dev/metropolis/pkg/kexec"
	"source.monogon.dev/metropolis/pkg/logtree"
)

// Service contains data and functionality to perform A/B updates on a
// Metropolis node.
type Service struct {
	// Path to the mount point of the EFI System Partition (ESP).
	ESPPath string
	// gpt.Partition of the ESP System Partition.
	ESPPart *gpt.Partition
	// Partition number (1-based) of the ESP in the GPT partitions array.
	ESPPartNumber uint32

	// Logger service for the update service.
	Logger logtree.LeveledLogger
}

type Slot int

const (
	SlotInvalid Slot = 0
	SlotA       Slot = 1
	SlotB       Slot = 2
)

// Other returns the "other" slot, i.e. returns slot A for B and B for A.
// It returns SlotInvalid for any s which is not SlotA or SlotB.
func (s Slot) Other() Slot {
	switch s {
	case SlotA:
		return SlotB
	case SlotB:
		return SlotA
	default:
		return SlotInvalid
	}
}

func (s Slot) String() string {
	switch s {
	case SlotA:
		return "A"
	case SlotB:
		return "B"
	default:
		return "<invalid slot>"
	}
}

func (s Slot) EFIBootPath() string {
	switch s {
	case SlotA:
		return osimage.EFIBootAPath
	case SlotB:
		return osimage.EFIBootBPath
	default:
		return ""
	}
}

var slotRegexp = regexp.MustCompile(`PARTLABEL=METROPOLIS-SYSTEM-([AB])`)

// ProvideESP is a convenience function for providing information about the
// ESP after the update service has been instantiated.
func (s *Service) ProvideESP(path string, partNum uint32, part *gpt.Partition) {
	s.ESPPath = path
	s.ESPPartNumber = partNum
	s.ESPPart = part
}

// CurrentlyRunningSlot returns the slot the current system is booted from.
func (s *Service) CurrentlyRunningSlot() Slot {
	cmdline, err := os.ReadFile("/proc/cmdline")
	if err != nil {
		return SlotInvalid
	}
	slotMatches := slotRegexp.FindStringSubmatch(string(cmdline))
	if len(slotMatches) != 2 {
		return SlotInvalid
	}
	switch slotMatches[1] {
	case "A":
		return SlotA
	case "B":
		return SlotB
	default:
		panic("unreachable")
	}
}

var bootVarRegexp = regexp.MustCompile(`^Boot([0-9A-Fa-f]{4})$`)

func (s *Service) getAllBootEntries() (map[int]*efivarfs.LoadOption, error) {
	res := make(map[int]*efivarfs.LoadOption)
	varNames, err := efivarfs.List(efivarfs.ScopeGlobal)
	if err != nil {
		return nil, fmt.Errorf("failed to list EFI variables: %w", err)
	}
	for _, varName := range varNames {
		m := bootVarRegexp.FindStringSubmatch(varName)
		if m == nil {
			continue
		}
		idx, err := strconv.ParseUint(m[1], 16, 16)
		if err != nil {
			// This cannot be hit as all regexp matches are parseable.
			panic(err)
		}
		e, err := efivarfs.GetBootEntry(int(idx))
		if err != nil {
			s.Logger.Warningf("Unable to get boot entry %d, skipping: %v", idx, err)
			continue
		}
		res[int(idx)] = e
	}
	return res, nil
}

// MarkBootSuccessful must be called after each boot if some implementation-
// defined criteria for a successful boot are met. If an update has been
// installed and booted and this function is called, the updated version is
// marked as default. If an issue occurs during boot and so this function is
// not called the old version will be started again on next boot.
func (s *Service) MarkBootSuccessful() error {
	if s.ESPPath == "" {
		return errors.New("no ESP information provided to update service, cannot continue")
	}
	activeSlot := s.CurrentlyRunningSlot()
	abState, err := s.getABState()
	if err != nil {
		s.Logger.Warningf("Error while getting A/B loader state, recreating: %v", err)
		abState = &abloaderpb.ABLoaderData{
			ActiveSlot: abloaderpb.Slot(activeSlot),
		}
		err := s.setABState(abState)
		if err != nil {
			return fmt.Errorf("while recreating A/B loader state: %w", err)
		}
	}
	if Slot(abState.ActiveSlot) != activeSlot {
		err := s.setABState(&abloaderpb.ABLoaderData{
			ActiveSlot: abloaderpb.Slot(activeSlot),
		})
		if err != nil {
			return fmt.Errorf("while setting next A/B slot: %w", err)
		}
		s.Logger.Infof("Permanently activated slot %v", activeSlot)
	} else {
		s.Logger.Infof("Normal boot from slot %v", activeSlot)
	}

	return nil
}

func openSystemSlot(slot Slot) (*blockdev.Device, error) {
	switch slot {
	case SlotA:
		return blockdev.Open("/dev/system-a")
	case SlotB:
		return blockdev.Open("/dev/system-b")
	default:
		return nil, errors.New("invalid slot identifier given")
	}
}

func (s *Service) getABState() (*abloaderpb.ABLoaderData, error) {
	abDataRaw, err := os.ReadFile(filepath.Join(s.ESPPath, "EFI/metropolis/loader_state.pb"))
	if err != nil {
		return nil, err
	}
	var abData abloaderpb.ABLoaderData
	if err := proto.Unmarshal(abDataRaw, &abData); err != nil {
		return nil, err
	}
	return &abData, nil
}

func (s *Service) setABState(d *abloaderpb.ABLoaderData) error {
	abDataRaw, err := proto.Marshal(d)
	if err != nil {
		return fmt.Errorf("while marshaling: %w", err)
	}
	if err := os.WriteFile(filepath.Join(s.ESPPath, "EFI/metropolis/loader_state.pb"), abDataRaw, 0666); err != nil {
		return err
	}
	return nil
}

// InstallBundle installs the bundle at the given HTTP(S) URL into the currently
// inactive slot and sets that slot to boot next. If it doesn't return an error,
// a reboot boots into the new slot.
func (s *Service) InstallBundle(ctx context.Context, bundleURL string, withKexec bool) error {
	if s.ESPPath == "" {
		return errors.New("no ESP information provided to update service, cannot continue")
	}
	// Download into a buffer as ZIP files cannot efficiently be read from
	// HTTP in Go as the ReaderAt has no way of indicating continuous sections,
	// thus a ton of small range requests would need to be used, causing
	// a huge latency penalty as well as costing a lot of money on typical
	// object storages. This should go away when we switch to a better bundle
	// format which can be streamed.
	var bundleRaw bytes.Buffer
	b := backoff.NewExponentialBackOff()
	err := backoff.Retry(func() error {
		return s.tryDownloadBundle(ctx, bundleURL, &bundleRaw)
	}, backoff.WithContext(b, ctx))
	if err != nil {
		return fmt.Errorf("error downloading Metropolis bundle: %v", err)
	}
	bundle, err := zip.NewReader(bytes.NewReader(bundleRaw.Bytes()), int64(bundleRaw.Len()))
	if err != nil {
		return fmt.Errorf("failed to open node bundle: %w", err)
	}
	efiPayload, err := bundle.Open("kernel_efi.efi")
	if err != nil {
		return fmt.Errorf("invalid bundle: %w", err)
	}
	defer efiPayload.Close()
	systemImage, err := bundle.Open("verity_rootfs.img")
	if err != nil {
		return fmt.Errorf("invalid bundle: %w", err)
	}
	defer systemImage.Close()
	activeSlot := s.CurrentlyRunningSlot()
	if activeSlot == SlotInvalid {
		return errors.New("unable to determine active slot, cannot continue")
	}
	targetSlot := activeSlot.Other()

	systemPart, err := openSystemSlot(targetSlot)
	if err != nil {
		return status.Errorf(codes.Internal, "Inactive system slot unavailable: %v", err)
	}
	defer systemPart.Close()
	if _, err := io.Copy(blockdev.NewRWS(systemPart), systemImage); err != nil {
		return status.Errorf(codes.Unavailable, "Failed to copy system image: %v", err)
	}

	bootFile, err := os.Create(filepath.Join(s.ESPPath, targetSlot.EFIBootPath()))
	if err != nil {
		return fmt.Errorf("failed to open boot file: %w", err)
	}
	defer bootFile.Close()
	if _, err := io.Copy(bootFile, efiPayload); err != nil {
		return fmt.Errorf("failed to write boot file: %w", err)
	}

	if withKexec {
		if err := s.stageKexec(bootFile, targetSlot); err != nil {
			return fmt.Errorf("while kexec staging: %w", err)
		}
	} else {
		err := s.setABState(&abloaderpb.ABLoaderData{
			ActiveSlot: abloaderpb.Slot(activeSlot),
			NextSlot:   abloaderpb.Slot(targetSlot),
		})
		if err != nil {
			return fmt.Errorf("while setting next A/B slot: %w", err)
		}
	}

	return nil
}

func (*Service) tryDownloadBundle(ctx context.Context, bundleURL string, bundleRaw *bytes.Buffer) error {
	bundleReq, err := http.NewRequestWithContext(ctx, "GET", bundleURL, nil)
	bundleRes, err := http.DefaultClient.Do(bundleReq)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer bundleRes.Body.Close()
	switch bundleRes.StatusCode {
	case http.StatusTooEarly, http.StatusTooManyRequests,
		http.StatusInternalServerError, http.StatusBadGateway,
		http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return fmt.Errorf("HTTP error %d", bundleRes.StatusCode)
	default:
		// Non-standard code range used for proxy-related issue by various
		// vendors. Treat as non-permanent error.
		if bundleRes.StatusCode >= 520 && bundleRes.StatusCode < 599 {
			return fmt.Errorf("HTTP error %d", bundleRes.StatusCode)
		}
		if bundleRes.StatusCode != 200 {
			return backoff.Permanent(fmt.Errorf("HTTP error %d", bundleRes.StatusCode))
		}
	}
	if _, err := bundleRaw.ReadFrom(bundleRes.Body); err != nil {
		bundleRaw.Reset()
		return err
	}
	return nil
}

// newMemfile creates a new file which is not located on a specific filesystem,
// but is instead backed by anonymous memory.
func newMemfile(name string, flags int) (*os.File, error) {
	fd, err := unix.MemfdCreate(name, flags)
	if err != nil {
		return nil, fmt.Errorf("memfd_create: %w", err)
	}
	return os.NewFile(uintptr(fd), name), nil
}

// stageKexec stages the kernel, command line and initramfs if available for
// a future kexec. It extracts the relevant data from the EFI boot executable.
func (s *Service) stageKexec(bootFile io.ReaderAt, targetSlot Slot) error {
	bootPE, err := pe.NewFile(bootFile)
	if err != nil {
		return fmt.Errorf("unable to open bootFile as PE: %w", err)
	}
	var cmdlineRaw []byte
	cmdlineSection := bootPE.Section(".cmdline")
	if cmdlineSection == nil {
		return fmt.Errorf("no .cmdline section in boot PE")
	}
	cmdlineRaw, err = cmdlineSection.Data()
	if err != nil {
		return fmt.Errorf("while reading .cmdline PE section: %w", err)
	}
	cmdline := string(bytes.TrimRight(cmdlineRaw, "\x00"))
	cmdline = strings.ReplaceAll(cmdline, "METROPOLIS-SYSTEM-X", fmt.Sprintf("METROPOLIS-SYSTEM-%s", targetSlot))
	kernelFile, err := newMemfile("kernel", 0)
	if err != nil {
		return fmt.Errorf("failed to create kernel memfile: %w", err)
	}
	defer kernelFile.Close()
	kernelSection := bootPE.Section(".linux")
	if kernelSection == nil {
		return fmt.Errorf("no .linux section in boot PE")
	}
	if _, err := io.Copy(kernelFile, kernelSection.Open()); err != nil {
		return fmt.Errorf("while copying .linux PE section: %w", err)
	}

	initramfsSection := bootPE.Section(".initrd")
	var initramfsFile *os.File
	if initramfsSection != nil && initramfsSection.Size > 0 {
		initramfsFile, err = newMemfile("initramfs", 0)
		if err != nil {
			return fmt.Errorf("failed to create initramfs memfile: %w", err)
		}
		defer initramfsFile.Close()
		if _, err := io.Copy(initramfsFile, initramfsSection.Open()); err != nil {
			return fmt.Errorf("while copying .initrd PE section: %w", err)
		}
	}
	if err := kexec.FileLoad(kernelFile, initramfsFile, cmdline); err != nil {
		return fmt.Errorf("while staging new kexec kernel: %w", err)
	}
	return nil
}
