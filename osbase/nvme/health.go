// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package nvme

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

// healthPage represents the raw data from a NVMe Health/SMART page.
// See Figure 93 in the spec.
type healthPage struct {
	CriticalWarning         uint8
	CompositeTemperature    uint16
	AvailableSpare          uint8
	AvailableSpareThreshold uint8
	PercentageUsed          uint8

	_ [26]byte

	DataUnitsRead               uint128le
	DataUnitsWritten            uint128le
	HostReadCommands            uint128le
	HostWriteCommands           uint128le
	ControllerBusyTime          uint128le
	PowerCycles                 uint128le
	PowerOnHours                uint128le
	UnsafeSHutdowns             uint128le
	MediaAndDataIntegrityErrors uint128le
	ErrorInformationLogEntries  uint128le

	WarningCompositeTemperatureTime  uint32
	CriticalCompositeTemperatureTime uint32

	TemperatureSensors [8]uint16

	ThermalMgmtTemperature1TransitionCount uint32
	ThermalMgmtTemperature2TransitionCount uint32

	_ [8]byte

	TotalTimeForThermalMgmtTemperature1 uint32
	TotalTimeForThermalMgmtTemperature2 uint32
}

// HealthInfo contains information related to the health of the NVMe device.
//
// Note that some values might be clamped under highly abnormal circumstances
// as they are reported as 128-bit integers which Go doesn't support.
// For easier handling values which are very unlikely to exceed 64 bits are
// exposed as 64 bit integers.
type HealthInfo struct {
	// AvailableSpareSpaceCritical is set if the avilable spare threshold has
	// fallen below the critical threshold.
	AvailableSpareSpaceCritical bool
	// TemperatureCritical is set if a temperature is outside the acceptable
	// operating thresholds.
	TemperatureCritical bool
	// MediaCritical is set if significant media or internal issues affect the
	// operation of the device.
	MediaCritical bool
	// ForcedReadOnly is set if the device is forced into read-only mode due
	// to an error.
	ForcedReadOnly bool
	// VolatileMemoryBackupFailed is set if the volatile memory backup device
	// has failed.
	VolatileMemoryBackupFailed bool
	// CompositeTemperatureKelvin contains a derived value representing the
	// composite state of controller and namespace/flash temperature.
	// The exact mechanism used to derive it is vendor-specific.
	CompositeTemperatureKelvin uint16
	// AvailableSpare represents the relative amount (0-1) of spare capacity
	// still unnused.
	AvailableSpare float32
	// AvailableSpareThreshold represents the vendor-defined threshold which
	// AvailableSpare shuld not fall under.
	AvailableSpareThreshold float32
	// LifeUsed represents vendor-defined relative estimate of the life of
	// the device which has been used up. It is allowed to exceed 1 and will
	// be clamped by the device somewhere between 1.0 and 2.55.
	LifeUsed float32
	// BytesRead contains the number of bytes read from the device.
	// This value is only updated in 512KiB increments.
	BytesRead *big.Int
	// BytesWritten contains the number of bytes written to the device.
	// This value is only updated in 512KiB increments.
	BytesWritten *big.Int
	// HostReadCommands contains the number of read commands completed by the
	// controller.
	HostReadCommands *big.Int
	// HostWriteCommands contains the number of write commands completed by the
	// controller.
	HostWriteCommands *big.Int
	// ControllerBusyTime contains the cumulative amount of time the controller
	// has spent being busy (i.e. having at least one command outstanding on an
	// I/O queue). This value is only updated in 1m increments.
	ControllerBusyTime time.Duration
	// PowerCycles contains the number of power cycles.
	PowerCycles uint64
	// PowerOnHours contains the number of hours the controller has been
	// powered on. Depending on the vendor implementation it may or may
	// not contain time spent in a non-operational power state.
	PowerOnHours uint64
	// UnsafeShutdown contains the number of power loss events without
	// a prior shutdown notification from the host.
	UnsafeShutdowns uint64
	// MediaAndDataIntegrityErrors contains the number of occurrences where the
	// controller detecte an unrecovered data integrity error.
	MediaAndDataIntegrityErrors uint64
	// ErrorInformationLogEntriesCount contains the number of Error
	// Information log entries over the life of the controller.
	ErrorInformationLogEntriesCount uint64
	// WarningCompositeTemperatureTime contains the amount of time the
	// controller is operational while the composite temperature is greater
	// than the warning composite threshold.
	WarningCompositeTemperatureTime time.Duration
	// CriticalCompositeTemperatureTime contains the amount of time the
	// controller is operational while the composite temperature is greater
	// than the critical composite threshold.
	CriticalCompositeTemperatureTime time.Duration
	// TemperatureSensorValues contains the current temperature in Kelvin as
	// reported by up to 8 sensors on the device. A value of zero means that
	// the given sensor is not available.
	TemperatureSensorValues [8]uint16
	// ThermalMgmtTemperature1TransitionCount contains the number of times the
	// controller transitioned to lower power active power states or performed
	// vendor-specific thermal management actions to reduce temperature.
	ThermalMgmtTemperature1TransitionCount uint32
	// ThermalMgmtTemperature2TransitionCount is the same as above, but
	// for "heavier" thermal management actions including heavy throttling.
	// The actual difference is vendor-specific.
	ThermalMgmtTemperature2TransitionCount uint32
	// TotalTimeForThermalMgmtTemperature1 contains the total time the
	// controller spent under "light" thermal management.
	TotalTimeForThermalMgmtTemperature1 time.Duration
	// TotalTimeForThermalMgmtTemperature2 contains the total time the
	// controller spent under "heavy" thermal management.
	TotalTimeForThermalMgmtTemperature2 time.Duration
}

// HasCriticalWarning returns true if any of the critical warnings
// (AvailableSpareSpaceCritical, TemperatureCritical, MediaCritical,
// ForcedReadOnly, VolatileMemoryBackupFailed) are active.
// If this returns true the NVMe medium has reason to believe that
// data availability or integrity is endangered.
func (h *HealthInfo) HasCriticalWarning() bool {
	return h.AvailableSpareSpaceCritical || h.TemperatureCritical || h.MediaCritical || h.ForcedReadOnly || h.VolatileMemoryBackupFailed
}

// See Figure 93 Data Units Read
var dataUnit = big.NewInt(512 * 1000)

const (
	healthLogPage = 0x02
)

// GetHealthInfo gets health information from the NVMe device's health log page.
func (d *Device) GetHealthInfo() (*HealthInfo, error) {
	var buf [512]byte

	if err := d.GetLogPage(GlobalNamespace, healthLogPage, 0, 0, buf[:]); err != nil {
		return nil, fmt.Errorf("unable to get health log page: %w", err)
	}

	var page healthPage
	binary.Read(bytes.NewReader(buf[:]), binary.LittleEndian, &page)
	var res HealthInfo
	res.AvailableSpareSpaceCritical = page.CriticalWarning&(1<<0) != 0
	res.TemperatureCritical = page.CriticalWarning&(1<<1) != 0
	res.MediaCritical = page.CriticalWarning&(1<<2) != 0
	res.ForcedReadOnly = page.CriticalWarning&(1<<3) != 0
	res.VolatileMemoryBackupFailed = page.CriticalWarning&(1<<4) != 0
	res.CompositeTemperatureKelvin = page.CompositeTemperature
	res.AvailableSpare = float32(page.AvailableSpare) / 100.
	res.AvailableSpareThreshold = float32(page.AvailableSpareThreshold) / 100.
	res.LifeUsed = float32(page.PercentageUsed) / 100.
	res.BytesRead = new(big.Int).Mul(page.DataUnitsRead.BigInt(), dataUnit)
	res.BytesWritten = new(big.Int).Mul(page.DataUnitsWritten.BigInt(), dataUnit)
	res.HostReadCommands = page.HostReadCommands.BigInt()
	res.HostWriteCommands = page.HostWriteCommands.BigInt()
	res.ControllerBusyTime = time.Duration(page.ControllerBusyTime.Uint64()) * time.Minute
	res.PowerCycles = page.PowerCycles.Uint64()
	res.PowerOnHours = page.PowerOnHours.Uint64()
	res.UnsafeShutdowns = page.UnsafeSHutdowns.Uint64()
	res.MediaAndDataIntegrityErrors = page.MediaAndDataIntegrityErrors.Uint64()
	res.ErrorInformationLogEntriesCount = page.ErrorInformationLogEntries.Uint64()
	res.WarningCompositeTemperatureTime = time.Duration(page.WarningCompositeTemperatureTime) * time.Minute
	res.CriticalCompositeTemperatureTime = time.Duration(page.CriticalCompositeTemperatureTime) * time.Minute
	res.TemperatureSensorValues = page.TemperatureSensors
	res.ThermalMgmtTemperature1TransitionCount = page.ThermalMgmtTemperature1TransitionCount
	res.ThermalMgmtTemperature2TransitionCount = page.ThermalMgmtTemperature2TransitionCount
	res.TotalTimeForThermalMgmtTemperature1 = time.Duration(page.TotalTimeForThermalMgmtTemperature1) * time.Second
	res.TotalTimeForThermalMgmtTemperature2 = time.Duration(page.TotalTimeForThermalMgmtTemperature2) * time.Second
	return &res, nil
}
