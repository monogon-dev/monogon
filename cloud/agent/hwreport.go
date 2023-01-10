package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/mdlayher/ethtool"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"source.monogon.dev/cloud/agent/api"
	"source.monogon.dev/metropolis/pkg/nvme"
	"source.monogon.dev/metropolis/pkg/scsi"
	"source.monogon.dev/metropolis/pkg/smbios"
)

type hwReportContext struct {
	node   *api.Node
	errors []error
}

func (c *hwReportContext) gatherSMBIOS() {
	smbiosFile, err := os.Open("/sys/firmware/dmi/tables/DMI")
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("unable to open SMBIOS table: %w", err))
		return
	}
	defer smbiosFile.Close()
	smbTbl, err := smbios.Unmarshal(bufio.NewReader(smbiosFile))
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("unable to parse SMBIOS table: %w", err))
		return
	}
	if smbTbl.SystemInformationRaw != nil {
		c.node.Manufacturer = smbTbl.SystemInformationRaw.Manufacturer
		c.node.Product = smbTbl.SystemInformationRaw.ProductName
		c.node.SerialNumber = smbTbl.SystemInformationRaw.SerialNumber
	}
	for _, d := range smbTbl.MemoryDevicesRaw {
		if d.StructureVersion.AtLeast(3, 2) && d.MemoryTechnology != 0x03 {
			// If MemoryTechnology is available, only count DRAM
			continue
		}
		size, ok := d.SizeBytes()
		if !ok {
			continue
		}
		c.node.MemoryInstalledBytes += int64(size)
	}
	return
}

var memoryBlockRegexp = regexp.MustCompile("^memory[0-9]+$")

func (c *hwReportContext) gatherMemorySysfs() {
	blockSizeRaw, err := os.ReadFile("/sys/devices/system/memory/block_size_bytes")
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("unable to read memory block size, CONFIG_MEMORY_HOTPLUG disabled or sandbox?: %w", err))
		return
	}
	blockSize, err := strconv.ParseInt(strings.TrimSpace(string(blockSizeRaw)), 16, 64)
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("failed to parse memory block size (%q): %w", string(blockSizeRaw), err))
		return
	}
	dirEntries, err := os.ReadDir("/sys/devices/system/memory")
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("unable to read sysfs memory devices list: %w", err))
		return
	}
	c.node.MemoryInstalledBytes = 0
	for _, e := range dirEntries {
		if memoryBlockRegexp.MatchString(e.Name()) {
			// This is safe as the regexp does not allow for any dots
			state, err := os.ReadFile("/sys/devices/system/memory/%s/state")
			if os.IsNotExist(err) {
				// Memory hotplug operation raced us
				continue
			} else if err != nil {
				c.errors = append(c.errors, fmt.Errorf("failed to read memory block state for %s: %w", e.Name(), err))
				continue
			}
			if strings.TrimSpace(string(state)) != "online" {
				// Only count online memory
				continue
			}
			// Each block is one blockSize of memory
			c.node.MemoryInstalledBytes += blockSize
		}
	}
	return
}

func parseCpuinfoAMD64(cpuinfoRaw []byte) (*api.CPU, []error) {
	// Parse line-by-line, each segment is separated by a line with no colon
	// character, a  segment describes a logical processor if it contains
	// the key "processor". Keep track of all seen core IDs (physical
	// processors) and processor IDs (logical processors) in a map to fill
	// into the structure.
	s := bufio.NewScanner(bytes.NewReader(cpuinfoRaw))
	var cpu api.CPU
	scannedVals := make(map[string]string)
	seenCoreIDs := make(map[string]bool)
	seenProcessorIDs := make(map[string]bool)
	processItem := func() error {
		if _, ok := scannedVals["processor"]; !ok {
			// Not a cpu, clear data and return
			scannedVals = make(map[string]string)
			return nil
		}
		seenProcessorIDs[scannedVals["processor"]] = true
		seenCoreIDs[scannedVals["core id"]] = true
		cpu.Model = scannedVals["model name"]
		cpu.Vendor = scannedVals["vendor_id"]
		family, err := strconv.Atoi(scannedVals["cpu family"])
		if err != nil {
			return fmt.Errorf("unable to parse CPU family to int: %v", err)
		}
		model, err := strconv.Atoi(scannedVals["model"])
		if err != nil {
			return fmt.Errorf("unable to parse CPU model to int: %v", err)
		}
		stepping, err := strconv.Atoi(scannedVals["stepping"])
		if err != nil {
			return fmt.Errorf("unable to parse CPU stepping to int: %v", err)
		}
		cpu.Architecture = &api.CPU_X86_64_{
			X86_64: &api.CPU_X86_64{
				Family:   int32(family),
				Model:    int32(model),
				Stepping: int32(stepping),
			},
		}
		scannedVals = make(map[string]string)
		return nil
	}
	var errs []error
	for s.Scan() {
		k, v, ok := strings.Cut(s.Text(), ":")
		// If there is a colon, add property to scannedVals.
		if ok {
			scannedVals[strings.TrimSpace(k)] = strings.TrimSpace(v)
			continue
		}
		// Otherwise this is a segment boundary, process the segment.
		if err := processItem(); err != nil {
			errs = append(errs, fmt.Errorf("error parsing cpuinfo block: %w", err))
		}
	}
	// Parse the last segment.
	if err := processItem(); err != nil {
		errs = append(errs, fmt.Errorf("error parsing cpuinfo block: %w", err))
	}
	cpu.Cores = int32(len(seenCoreIDs))
	cpu.HardwareThreads = int32(len(seenProcessorIDs))
	return &cpu, errs
}

func (c *hwReportContext) gatherCPU() {
	switch runtime.GOARCH {
	case "amd64":
		// Currently a rather simple gatherer with no special NUMA handling
		cpuinfoRaw, err := os.ReadFile("/proc/cpuinfo")
		if err != nil {
			c.errors = append(c.errors, fmt.Errorf("unable to read cpuinfo: %w", err))
			return
		}
		cpu, errs := parseCpuinfoAMD64(cpuinfoRaw)
		c.errors = append(c.errors, errs...)
		c.node.Cpu = append(c.node.Cpu, cpu)
	default:
		// Currently unimplemented, do nothing
		c.errors = append(c.errors, fmt.Errorf("architecture %v unsupported by CPU gatherer", runtime.GOARCH))
	}
	return
}

var (
	FRUUnavailable = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

func (c *hwReportContext) gatherNVMe(bd *api.BlockDevice, bde os.DirEntry) error {
	bd.Protocol = api.BlockDevice_NVME
	nvmeDev, err := nvme.Open("/dev/" + bde.Name())
	if err != nil {
		return fmt.Errorf("unable to open NVMe device: %w", err)
	}
	defer nvmeDev.Close()
	identifyData, err := nvmeDev.Identify()
	if err != nil {
		return fmt.Errorf("calling Identify failed: %w", err)
	}
	bd.DeviceModel = identifyData.ModelNumber
	bd.SerialNumber = identifyData.SerialNumber
	if identifyData.FRUGloballyUniqueIdentifier != FRUUnavailable {
		bd.Wwn = identifyData.FRUGloballyUniqueIdentifier[:]
	}
	if healthInfo, err := nvmeDev.GetHealthInfo(); err == nil {
		bd.AvailableSpareRatio = &healthInfo.AvailableSpare
		bd.CriticalWarning = healthInfo.HasCriticalWarning()
		var mediaErrors = int64(healthInfo.MediaAndDataIntegrityErrors)
		bd.MediaErrors = &mediaErrors
		bd.UsageRatio = &healthInfo.LifeUsed
	}
	return nil
}

func (c *hwReportContext) gatherSCSI(bd *api.BlockDevice, bde os.DirEntry) error {
	bd.Protocol = api.BlockDevice_SCSI
	scsiDev, err := scsi.Open("/dev/" + bde.Name())
	if err != nil {
		return fmt.Errorf("unable to open SCSI device: %w", err)
	}
	defer scsiDev.Close()
	inquiryData, err := scsiDev.Inquiry()
	if err != nil {
		return fmt.Errorf("failed calling INQUIRY: %w", err)
	}
	if serial, err := scsiDev.UnitSerialNumber(); err == nil {
		bd.SerialNumber = serial
	}

	// SAT-5 R8 Table 14
	if inquiryData.Vendor == "ATA" { // ATA device behind SAT
		bd.Protocol = api.BlockDevice_ATA
		// TODO: ATA Vendor from WWN if available
	} else { // Normal SCSI device
		bd.Vendor = inquiryData.Vendor
		// Attempt to read defect list to populate media error count
		var mediaErrors int64
		if defectsLBA, err := scsiDev.ReadDefectDataLBA(false, true); err == nil {
			mediaErrors = int64(len(defectsLBA))
			bd.MediaErrors = &mediaErrors
		} else if defectsPhysical, err := scsiDev.ReadDefectDataPhysical(false, true); err == nil {
			mediaErrors = int64(len(defectsPhysical))
			bd.MediaErrors = &mediaErrors
		}
		if mediaHealth, err := scsiDev.SolidStateMediaHealth(); err == nil {
			used := float32(mediaHealth.PercentageUsedEnduranceIndicator) / 100.
			bd.UsageRatio = &used
		}
		if informationalExceptions, err := scsiDev.GetInformationalExceptions(); err == nil {
			// Only consider FailurePredictionThresholdExceeded-class sense codes critical.
			// The second commonly reported error here according to random forums are
			// Warning-class errors, but looking through these they don't indicate imminent
			// or even permanent errors.
			bd.CriticalWarning = informationalExceptions.InformationalSenseCode.IsKey(scsi.FailurePredictionThresholdExceeded)
		}
		// SCSI has no reporting of available spares, so this will never be populated
	}
	bd.DeviceModel = inquiryData.Product
	return nil
}

func (c *hwReportContext) gatherBlockDevices() {
	blockDeviceEntries, err := os.ReadDir("/sys/class/block")
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("unable to read sysfs block device list: %w", err))
		return
	}
	for _, bde := range blockDeviceEntries {
		sysfsDir := fmt.Sprintf("/sys/class/block/%s", bde.Name())
		if _, err := os.Stat(sysfsDir + "/partition"); err == nil {
			// Ignore partitions, we only care about their parents
			continue
		}
		var bd api.BlockDevice
		if rotational, err := os.ReadFile(sysfsDir + "/queue/rotational"); err == nil {
			if strings.TrimSpace(string(rotational)) == "1" {
				bd.Rotational = true
			}
		}
		if sizeRaw, err := os.ReadFile(sysfsDir + "/size"); err == nil {
			size, err := strconv.ParseInt(strings.TrimSpace(string(sizeRaw)), 10, 64)
			if err != nil {
				c.errors = append(c.errors, fmt.Errorf("unable to parse block device %v size: %w", bde.Name(), err))
			} else {
				// Linux always defines size in terms of 512 byte blocks regardless
				// of what the configured logical and physical block sizes are.
				bd.CapacityBytes = size * 512
			}
		}
		if lbsRaw, err := os.ReadFile(sysfsDir + "/queue/logical_block_size"); err == nil {
			lbs, err := strconv.ParseInt(strings.TrimSpace(string(lbsRaw)), 10, 32)
			if err != nil {
				c.errors = append(c.errors, fmt.Errorf("unable to parse block device %v logical block size: %w", bde.Name(), err))
			} else {
				bd.LogicalBlockSizeBytes = int32(lbs)
			}
		}
		if pbsRaw, err := os.ReadFile(sysfsDir + "/queue/physical_block_size"); err == nil {
			pbs, err := strconv.ParseInt(strings.TrimSpace(string(pbsRaw)), 10, 32)
			if err != nil {
				c.errors = append(c.errors, fmt.Errorf("unable to parse physical block size: %w", err))
			} else {
				bd.PhysicalBlockSizeBytes = int32(pbs)
			}
		}
		if strings.HasPrefix(bde.Name(), "nvme") {
			err := c.gatherNVMe(&bd, bde)
			if err != nil {
				c.errors = append(c.errors, fmt.Errorf("block device %v: %w", bde.Name(), err))
			} else {
				c.node.BlockDevice = append(c.node.BlockDevice, &bd)
			}
		}
		if strings.HasPrefix(bde.Name(), "sd") {
			err := c.gatherSCSI(&bd, bde)
			if err != nil {
				c.errors = append(c.errors, fmt.Errorf("block device %v: %w", bde.Name(), err))
			} else {
				c.node.BlockDevice = append(c.node.BlockDevice, &bd)
			}
		}
		if strings.HasPrefix(bde.Name(), "mmcblk") {
			// TODO: MMC information
			bd.Protocol = api.BlockDevice_MMC
			c.node.BlockDevice = append(c.node.BlockDevice, &bd)
		}
	}
	return
}

var speedModeRegexp = regexp.MustCompile("^([0-9]+)base")

const mbps = (1000 * 1000) / 8

func (c *hwReportContext) gatherNICs() {
	links, err := netlink.LinkList()
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("failed to list network links: %w", err))
		return
	}
	ethClient, err := ethtool.New()
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("failed to get ethtool netlink client: %w", err))
		return
	}
	defer ethClient.Close()
	for _, l := range links {
		if l.Type() != "device" || len(l.Attrs().HardwareAddr) == 0 {
			// Not a physical device, ignore
			continue
		}
		var nif api.NetworkInterface
		nif.Mac = l.Attrs().HardwareAddr
		mode, err := ethClient.LinkMode(ethtool.Interface{Index: l.Attrs().Index})
		if err == nil {
			if mode.SpeedMegabits < math.MaxInt32 {
				nif.CurrentSpeedBytes = int64(mode.SpeedMegabits) * mbps
			}
			speeds := make(map[int64]bool)
			for _, m := range mode.Ours {
				// Doing this with a regexp is arguably more future-proof as
				// we don't need to add each link mode for the detection to
				// work.
				modeParts := speedModeRegexp.FindStringSubmatch(m.Name)
				if len(modeParts) > 0 {
					speedMegabits, err := strconv.ParseInt(modeParts[1], 10, 64)
					if err != nil {
						c.errors = append(c.errors, fmt.Errorf("nic %v: failed to parse %q as integer: %w", l.Attrs().Name, modeParts[1], err))
						continue
					}
					speeds[int64(speedMegabits)*mbps] = true
				}
			}
			for s := range speeds {
				nif.SupportedSpeedBytes = append(nif.SupportedSpeedBytes, s)
			}
			// Go randomizes the map keys, sort to make the report stable.
			sort.Slice(nif.SupportedSpeedBytes, func(i, j int) bool { return nif.SupportedSpeedBytes[i] > nif.SupportedSpeedBytes[j] })
		}
		state, err := ethClient.LinkState(ethtool.Interface{Index: l.Attrs().Index})
		if err == nil {
			nif.LinkUp = state.Link
		} else {
			// We have no ethtool support, fall back to checking if Linux
			// thinks the link is up.
			nif.LinkUp = l.Attrs().OperState == netlink.OperUp
		}
		// Linux blocks creation of interfaces which conflict with special path
		// characters, so this path assembly is fine.
		driverPath, err := os.Readlink("/sys/class/net/" + l.Attrs().Name + "/device/driver")
		if err == nil {
			nif.Driver = filepath.Base(driverPath)
		}
		c.node.NetworkInterface = append(c.node.NetworkInterface, &nif)
	}
	return
}

func gatherHWReport() (*api.Node, []error) {
	var hwReportCtx hwReportContext

	hwReportCtx.gatherCPU()
	hwReportCtx.gatherSMBIOS()
	if hwReportCtx.node.MemoryInstalledBytes == 0 {
		hwReportCtx.gatherMemorySysfs()
	}
	var sysinfo unix.Sysinfo_t
	if err := unix.Sysinfo(&sysinfo); err != nil {
		hwReportCtx.errors = append(hwReportCtx.errors, fmt.Errorf("unable to execute sysinfo syscall: %w", err))
	} else {
		hwReportCtx.node.MemoryUsableRatio = float32(sysinfo.Totalram) / float32(hwReportCtx.node.MemoryInstalledBytes)
	}
	hwReportCtx.gatherNICs()
	hwReportCtx.gatherBlockDevices()

	return hwReportCtx.node, hwReportCtx.errors
}
