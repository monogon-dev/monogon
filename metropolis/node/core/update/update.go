package update

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"debug/pe"
	_ "embed"
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

	"source.monogon.dev/go/logging"
	abloaderpb "source.monogon.dev/metropolis/node/core/abloader/spec"
	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/build/mkimage/osimage"
	"source.monogon.dev/osbase/efivarfs"
	"source.monogon.dev/osbase/gpt"
	"source.monogon.dev/osbase/kexec"
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
	Logger logging.Leveled
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

// MarkBootSuccessful must be called after each boot if some implementation-
// defined criteria for a successful boot are met. If an update has been
// installed and booted and this function is called, the updated version is
// marked as default. If an issue occurs during boot and so this function is
// not called the old version will be started again on next boot.
func (s *Service) MarkBootSuccessful() error {
	if s.ESPPath == "" {
		return errors.New("no ESP information provided to update service, cannot continue")
	}
	if err := s.fixupEFI(); err != nil {
		s.Logger.Errorf("Error when checking boot entry configuration: %v", err)
	}
	if err := s.fixupPreloader(); err != nil {
		s.Logger.Errorf("Error when fixing A/B preloader: %v", err)
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

// Rollback sets the currently-inactive slot as the next boot slot. This is
// intended to recover from scenarios where roll-forward fixing is difficult.
// Only the next boot slot is set to make sure that the node is not
// made unbootable accidentally. On successful bootup that code can switch the
// active slot to itself.
func (s *Service) Rollback() error {
	if s.ESPPath == "" {
		return errors.New("no ESP information provided to update service, cannot continue")
	}
	activeSlot := s.CurrentlyRunningSlot()
	abState, err := s.getABState()
	if err != nil {
		return fmt.Errorf("no valid A/B loader state, cannot rollback: %w", err)
	}
	nextSlot := activeSlot.Other()
	err = s.setABState(&abloaderpb.ABLoaderData{
		ActiveSlot: abState.ActiveSlot,
		NextSlot:   abloaderpb.Slot(nextSlot),
	})
	if err != nil {
		return fmt.Errorf("while setting next A/B slot: %w", err)
	}
	s.Logger.Warningf("Rollback requested, NextSlot set to %v", nextSlot)
	return nil
}

// KexecLoadNext loads the slot to be booted next into the kexec staging area.
// The next slot can then be launched by executing kexec via the reboot
// syscall. Calling this function counts as a next boot for the purposes of
// A/B state tracking, so it should not be called without kexecing afterwards.
func (s *Service) KexecLoadNext() error {
	state, err := s.getABState()
	if err != nil {
		return fmt.Errorf("bad A/B state: %w", err)
	}
	slotToLoad := Slot(state.ActiveSlot)
	if state.NextSlot != abloaderpb.Slot_SLOT_NONE {
		slotToLoad = Slot(state.NextSlot)
		state.NextSlot = abloaderpb.Slot_SLOT_NONE
		err = s.setABState(state)
		if err != nil {
			return fmt.Errorf("while updating A/B state: %w", err)
		}
	}
	boot, err := os.Open(filepath.Join(s.ESPPath, slotToLoad.EFIBootPath()))
	if err != nil {
		return fmt.Errorf("failed to open boot file for slot %v: %w", slotToLoad, err)
	}
	defer boot.Close()
	if err := s.stageKexec(boot, slotToLoad); err != nil {
		return fmt.Errorf("failed to stage next slot for kexec: %w", err)
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
		return fmt.Errorf("error downloading Metropolis bundle: %w", err)
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
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
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

//go:embed metropolis/node/core/abloader/abloader.efi
var abloader []byte

func (s *Service) fixupPreloader() error {
	abLoaderFile, err := os.Open(filepath.Join(s.ESPPath, osimage.EFIPayloadPath))
	if err != nil {
		s.Logger.Warningf("A/B preloader not available, attempting to restore: %v", err)
	} else {
		expectedSum := sha256.Sum256(abloader)
		h := sha256.New()
		_, err := io.Copy(h, abLoaderFile)
		abLoaderFile.Close()
		if err == nil {
			if bytes.Equal(h.Sum(nil), expectedSum[:]) {
				// A/B Preloader is present and has correct hash
				return nil
			} else {
				s.Logger.Infof("Replacing A/B preloader with current version: %x %x", h.Sum(nil), expectedSum[:])
			}
		} else {
			s.Logger.Warningf("Error while reading A/B preloader, restoring: %v", err)
		}
	}
	preloader, err := os.Create(filepath.Join(s.ESPPath, "preloader.swp"))
	if err != nil {
		return fmt.Errorf("while creating preloader swap file: %w", err)
	}
	if _, err := preloader.Write(abloader); err != nil {
		return fmt.Errorf("while writing preloader swap file: %w", err)
	}
	if err := preloader.Sync(); err != nil {
		return fmt.Errorf("while sync'ing preloader swap file: %w", err)
	}
	preloader.Close()
	if err := os.Rename(filepath.Join(s.ESPPath, "preloader.swp"), filepath.Join(s.ESPPath, osimage.EFIPayloadPath)); err != nil {
		return fmt.Errorf("while swapping preloader: %w", err)
	}
	s.Logger.Info("Successfully wrote current preloader")
	return nil
}

// fixupEFI checks for the existence and correctness of the EFI boot entry
// repairs/recreates it if needed.
func (s *Service) fixupEFI() error {
	varNames, err := efivarfs.List(efivarfs.ScopeGlobal)
	if err != nil {
		return fmt.Errorf("failed to list EFI variables: %w", err)
	}
	var validBootEntryIdx = -1
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
		if len(e.FilePath) != 2 {
			// Not our entry, ours always have two parts
			continue
		}
		switch p := e.FilePath[0].(type) {
		case *efivarfs.HardDrivePath:
			gptMatch, ok := p.PartitionMatch.(*efivarfs.PartitionGPT)
			if ok && gptMatch.PartitionUUID != s.ESPPart.ID {
				// Not related to our ESP
				continue
			}
		default:
			continue
		}
		switch p := e.FilePath[1].(type) {
		case efivarfs.FilePath:
			if string(p) == osimage.EFIPayloadPath {
				if validBootEntryIdx == -1 {
					validBootEntryIdx = int(idx)
				} else {
					// Another valid boot entry already exists, delete this one
					err := efivarfs.DeleteBootEntry(int(idx))
					if err == nil {
						s.Logger.Infof("Deleted duplicate boot entry %q", e.Description)
					} else {
						s.Logger.Warningf("Error while deleting duplicate boot entry %q: %v", e.Description, err)
					}
				}
			} else if strings.Contains(e.Description, "Metropolis") {
				err := efivarfs.DeleteBootEntry(int(idx))
				if err == nil {
					s.Logger.Infof("Deleted orphaned boot entry %q", e.Description)
				} else {
					s.Logger.Warningf("Error while deleting orphaned boot entry %q: %v", e.Description, err)
				}
			}
		default:
			continue
		}
	}
	if validBootEntryIdx == -1 {
		validBootEntryIdx, err = efivarfs.AddBootEntry(&efivarfs.LoadOption{
			Description: "Metropolis",
			FilePath: efivarfs.DevicePath{
				&efivarfs.HardDrivePath{
					PartitionNumber:     1,
					PartitionStartBlock: s.ESPPart.FirstBlock,
					PartitionSizeBlocks: s.ESPPart.SizeBlocks(),
					PartitionMatch: efivarfs.PartitionGPT{
						PartitionUUID: s.ESPPart.ID,
					},
				},
				efivarfs.FilePath(osimage.EFIPayloadPath),
			},
		})
		if err == nil {
			s.Logger.Infof("Restored missing EFI boot entry for Metropolis")
		} else {
			return fmt.Errorf("while restoring missing EFI boot entry for Metropolis: %w", err)
		}
	}
	bootOrder, err := efivarfs.GetBootOrder()
	if err != nil {
		return fmt.Errorf("failed to get EFI boot order: %w", err)
	}
	for _, bentry := range bootOrder {
		if bentry == uint16(validBootEntryIdx) {
			// Our boot entry is in the boot order, everything's ok
			return nil
		}
	}
	newBootOrder := append(efivarfs.BootOrder{uint16(validBootEntryIdx)}, bootOrder...)
	if err := efivarfs.SetBootOrder(newBootOrder); err != nil {
		return fmt.Errorf("while setting EFI boot order: %w", err)
	}
	return nil
}
