// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package watchdog provides access to hardware watchdogs. These can be used to
// automatically reset/reboot a system if they are no longer pinged.
package watchdog

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

// Device represents a handle to a hardware watchdog.
type Device struct {
	// Type identifies the type of watchdog device. It corresponds to the Linux
	// driver's watchdog_info.identity value.
	Type string
	// HasConfiguratbleTimeout indicates if the device supports the SetTimeout
	// call.
	HasConfigurableTimeout bool
	// HasPretimeout indicates if the device supports notifying the system of
	// an impending reset and the functions to control this
	// (Get/SetPreTimeout).
	HasPretimeout bool
	// Indicates if the watchdog is capable of reporting that it is responsible
	// for the last system reset.
	ReportsWatchdogReset bool

	raw syscall.RawConn
	f   *os.File
}

// Open opens a watchdog device identified by the path to its device inode.
func Open(name string) (*Device, error) {
	f, err := os.Open(name)
	if err != nil {
		// Already wrapped by PathError
		return nil, err
	}
	raw, err := f.SyscallConn()
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("while obtaining RawConn: %w", err)
	}
	var wdInfo *unix.WatchdogInfo
	ctrlErr := raw.Control(func(fd uintptr) {
		wdInfo, err = unix.IoctlGetWatchdogInfo(int(fd))
	})
	if ctrlErr != nil {
		f.Close()
		return nil, fmt.Errorf("when calling RawConn.Control: %w", err)
	}
	if errors.Is(err, unix.ENOTTY) {
		f.Close()
		return nil, errors.New("device is not a watchdog")
	}
	if err != nil {
		return nil, fmt.Errorf("while getting watchdog metadata: %w", err)
	}
	w := &Device{
		Type:                   string(bytes.Trim(wdInfo.Identity[:], "\x00")),
		f:                      f,
		raw:                    raw,
		HasConfigurableTimeout: wdInfo.Options&unix.WDIOF_SETTIMEOUT != 0,
		HasPretimeout:          wdInfo.Options&unix.WDIOF_PRETIMEOUT != 0,
		ReportsWatchdogReset:   wdInfo.Options&unix.WDIOF_CARDRESET != 0,
	}
	return w, nil
}

// SetTimeout sets the duration since the last ping after which it performs
// a recovery actions (usually a reset or reboot).
// Due to hardware limitations this function may approximate the set duration
// or not be a available at all. GetTimeout returns the active timeout.
func (w *Device) SetTimeout(t time.Duration) error {
	if !w.HasConfigurableTimeout {
		return errors.New("watchdog does not have a configurable timeout, check HasConfigurableTimeout")
	}
	var err error
	ctrlErr := w.raw.Control(func(fd uintptr) {
		err = unix.IoctlSetInt(int(fd), unix.WDIOC_SETTIMEOUT, int(math.Ceil(t.Seconds())))
	})
	if ctrlErr != nil {
		return fmt.Errorf("when calling RawConn.Control: %w", err)
	}
	if err != nil {
		return fmt.Errorf("ioctl(WDIOC_SETTIMEOUT): %w", err)
	}
	return nil
}

// GetTimeout returns the configured timeout duration.
func (w *Device) GetTimeout() (time.Duration, error) {
	var err error
	var t int
	ctrlErr := w.raw.Control(func(fd uintptr) {
		t, err = unix.IoctlGetInt(int(fd), unix.WDIOC_GETTIMEOUT)
	})
	if ctrlErr != nil {
		return 0, fmt.Errorf("when calling RawConn.Control: %w", err)
	}
	if err != nil {
		return 0, fmt.Errorf("ioctl(WDIOC_GETTIMEOUT): %w", err)
	}
	return time.Duration(t) * time.Second, nil
}

// SetPreTimeout sets the minimum duration left on the expiry timer where when
// it drops below that, the system is notified (via some high-priority
// interrupt, usually an NMI). This is only available if HasPretimeout is true.
// This can be used by the system (if it's still in a sem-working state) to
// recover or dump diagnostic information before it gets forcibly reset by the
// watchdog. To disable this functionality, set the duration to zero.
func (w *Device) SetPreTimeout(t time.Duration) error {
	if !w.HasPretimeout {
		return errors.New("watchdog does not have a pretimeout, check HasPretimeout")
	}
	var err error
	ctrlErr := w.raw.Control(func(fd uintptr) {
		err = unix.IoctlSetInt(int(fd), unix.WDIOC_SETPRETIMEOUT, int(math.Ceil(t.Seconds())))
	})
	if ctrlErr != nil {
		return fmt.Errorf("when calling RawConn.Control: %w", err)
	}
	if err != nil {
		return fmt.Errorf("ioctl(WDIOC_SETPRETIMEOUT): %w", err)
	}
	return nil
}

// GetPreTimeout gets the current pre-timeout (see SetPreTimeout for more).
func (w *Device) GetPreTimeout() (time.Duration, error) {
	if !w.HasPretimeout {
		return 0, errors.New("watchdog does not have a pretimeout, check HasPretimeout")
	}
	var err error
	var t int
	ctrlErr := w.raw.Control(func(fd uintptr) {
		t, err = unix.IoctlGetInt(int(fd), unix.WDIOC_GETPRETIMEOUT)
	})
	if ctrlErr != nil {
		return 0, fmt.Errorf("when calling RawConn.Control: %w", err)
	}
	if err != nil {
		return 0, fmt.Errorf("ioctl(WDIOC_GETPRETIMEOUT): %w", err)
	}
	return time.Duration(t) * time.Second, nil

}

// Ping the watchdog. This needs to be called regularly before the
// watchdog timeout expires, otherwise the system resets.
func (w *Device) Ping() error {
	var err error
	ctrlErr := w.raw.Control(func(fd uintptr) {
		err = unix.IoctlWatchdogKeepalive(int(fd))
	})
	if ctrlErr != nil {
		return fmt.Errorf("when calling RawConn.Control: %w", err)
	}
	if err != nil {
		return fmt.Errorf("ioctl(WDIOC_KEEPALIVE): %w", err)
	}
	return nil
}

// LastResetByWatchdog returns true if the last system reset was caused by
// this watchdog. Not all watchdogs report this accurately.
func (w *Device) LastResetByWatchdog() (bool, error) {
	if !w.ReportsWatchdogReset {
		return false, errors.New("watchdog does not report resets, check ReportsWatchdogReset")
	}
	var err error
	var flags int
	ctrlErr := w.raw.Control(func(fd uintptr) {
		flags, err = unix.IoctlGetInt(int(fd), unix.WDIOC_GETBOOTSTATUS)
	})
	if ctrlErr != nil {
		return false, fmt.Errorf("when calling RawConn.Control: %w", err)
	}
	if err != nil {
		return false, fmt.Errorf("ioctl(WDIOC_GETBOOTSTATUS): %w", err)
	}
	return flags&unix.WDIOF_CARDRESET != 0, nil
}

// Close disables the watchdog and releases all associated resources.
func (w *Device) Close() error {
	if w.f != nil {
		_, err := w.f.Write([]byte{'V'})
		errClose := w.f.Close()
		w.f = nil
		if err != nil {
			return err
		}
		return errClose
	}
	return nil
}

// CloseActive releases all resources and file handles, but keeps the
// watchdog active. Another system must reopen it and ping it before
// it expires to avoid a reset.
func (w *Device) CloseActive() error {
	if w.f != nil {
		err := w.f.Close()
		w.f = nil
		return err
	}
	return nil
}
