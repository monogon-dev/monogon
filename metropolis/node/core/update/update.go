package update

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/build/mkimage/osimage"
	"source.monogon.dev/metropolis/pkg/blockdev"
	"source.monogon.dev/metropolis/pkg/efivarfs"
	"source.monogon.dev/metropolis/pkg/logtree"
)

// Service contains data and functionality to perform A/B updates on a
// Metropolis node.
type Service struct {
	// Path to the mount point of the EFI System Partition (ESP).
	ESPPath string
	// UUID of the ESP System Partition.
	ESPUUID uuid.UUID
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
func (s *Service) ProvideESP(path string, partUUID uuid.UUID, partNum uint32) {
	s.ESPPath = path
	s.ESPPartNumber = partNum
	s.ESPUUID = partUUID
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
			return nil, fmt.Errorf("failed to get boot entry %d: %w", idx, err)
		}
		res[int(idx)] = e
	}
	return res, nil
}

func (s *Service) getOrMakeBootEntry(existing map[int]*efivarfs.LoadOption, slot Slot) (int, error) {
	for idx, e := range existing {
		if len(e.FilePath) != 2 {
			// Not our entry
			continue
		}
		switch p := e.FilePath[0].(type) {
		case *efivarfs.HardDrivePath:
			gptMatch, ok := p.PartitionMatch.(*efivarfs.PartitionGPT)
			if ok && gptMatch.PartitionUUID != s.ESPUUID {
				// Not related to our ESP
				continue
			}
		default:
			continue
		}
		switch p := e.FilePath[1].(type) {
		case efivarfs.FilePath:
			if string(p) == slot.EFIBootPath() {
				return idx, nil
			}
		default:
			continue
		}
	}
	newEntry := &efivarfs.LoadOption{
		Description: fmt.Sprintf("Metropolis Slot %s", slot),
		FilePath: efivarfs.DevicePath{
			&efivarfs.HardDrivePath{
				PartitionNumber: s.ESPPartNumber,
				PartitionMatch: efivarfs.PartitionGPT{
					PartitionUUID: s.ESPUUID,
				},
			},
			efivarfs.FilePath(slot.EFIBootPath()),
		},
	}
	newIdx, err := efivarfs.AddBootEntry(newEntry)
	if err == nil {
		existing[newIdx] = newEntry
	}
	return newIdx, err
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
	bootEntries, err := s.getAllBootEntries()
	if err != nil {
		return fmt.Errorf("while getting boot entries: %w", err)
	}
	aIdx, err := s.getOrMakeBootEntry(bootEntries, SlotA)
	if err != nil {
		return fmt.Errorf("while ensuring slot A boot entry: %w", err)
	}
	bIdx, err := s.getOrMakeBootEntry(bootEntries, SlotB)
	if err != nil {
		return fmt.Errorf("while ensuring slot B boot entry: %w", err)
	}

	activeSlot := s.CurrentlyRunningSlot()
	firstSlot := SlotInvalid

	ord, err := efivarfs.GetBootOrder()
	if err != nil {
		return fmt.Errorf("failed to get boot order: %w", err)
	}

	for _, e := range ord {
		if int(e) == aIdx {
			firstSlot = SlotA
			break
		}
		if int(e) == bIdx {
			firstSlot = SlotB
			break
		}
	}

	if firstSlot == SlotInvalid {
		bootOrder := make(efivarfs.BootOrder, 2)
		switch activeSlot {
		case SlotA:
			bootOrder[0], bootOrder[1] = uint16(aIdx), uint16(bIdx)
		case SlotB:
			bootOrder[0], bootOrder[1] = uint16(bIdx), uint16(aIdx)
		default:
			return fmt.Errorf("invalid active slot")
		}
		efivarfs.SetBootOrder(bootOrder)
		s.Logger.Infof("Metropolis missing from boot order, recreated it")
	} else if activeSlot != firstSlot {
		var aPos, bPos int
		for i, e := range ord {
			if int(e) == aIdx {
				aPos = i
			}
			if int(e) == bIdx {
				bPos = i
			}
		}
		// swap A and B slots in boot order
		ord[aPos], ord[bPos] = ord[bPos], ord[aPos]
		if err := efivarfs.SetBootOrder(ord); err != nil {
			return fmt.Errorf("failed to set boot order to permanently switch slot: %w", err)
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

// InstallBundle installs the bundle at the given HTTP(S) URL into the currently
// inactive slot and sets that slot to boot next. If it doesn't return an error,
// a reboot boots into the new slot.
func (s *Service) InstallBundle(ctx context.Context, bundleURL string) error {
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

	bootEntries, err := s.getAllBootEntries()
	if err != nil {
		return fmt.Errorf("while getting boot entries: %w", err)
	}
	targetSlotBootEntryIdx, err := s.getOrMakeBootEntry(bootEntries, targetSlot)
	if err != nil {
		return fmt.Errorf("while ensuring target slot boot entry: %w", err)
	}
	targetSlotBootEntry := bootEntries[targetSlotBootEntryIdx]

	// Disable boot entry while the corresponding slot is being modified.
	targetSlotBootEntry.Inactive = true
	if err := efivarfs.SetBootEntry(targetSlotBootEntryIdx, targetSlotBootEntry); err != nil {
		return fmt.Errorf("failed setting boot entry %d inactive: %w", targetSlotBootEntryIdx, err)
	}

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

	// Reenable target slot boot entry after boot and system have been written
	// fully. The slot should now be bootable again.
	targetSlotBootEntry.Inactive = false
	if err := efivarfs.SetBootEntry(targetSlotBootEntryIdx, targetSlotBootEntry); err != nil {
		return fmt.Errorf("failed setting boot entry %d active: %w", targetSlotBootEntryIdx, err)
	}

	if err := efivarfs.SetBootNext(uint16(targetSlotBootEntryIdx)); err != nil {
		return fmt.Errorf("failed to set BootNext variable: %w", err)
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
