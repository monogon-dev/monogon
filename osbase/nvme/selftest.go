// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package nvme

import (
	"bytes"
	"encoding/binary"
)

type SelfTestOp uint8

const (
	SelfTestNone     SelfTestOp = 0x0
	SelfTestShort    SelfTestOp = 0x1
	SelfTestExtended SelfTestOp = 0x2
	SelfTestAbort    SelfTestOp = 0xF
)

func (d *Device) StartSelfTest(ns uint32, action SelfTestOp) error {
	return d.RawCommand(&Command{
		Opcode:      0x14,
		NamespaceID: ns,
		CDW10:       uint32(action & 0xF),
	})
}

// Figure 99
type selfTestResult struct {
	SelfTestStatus             uint8
	SegmentNumber              uint8
	ValidDiagnosticInformation uint8
	_                          byte
	PowerOnHours               uint64
	NamespaceID                uint32
	FailingLBA                 uint64
	StatusCodeType             uint8
	StatusCode                 uint8
	VendorSpecific             [2]byte
}

// Figure 98
type selfTestLogPage struct {
	CurrentSelfTestOp         uint8
	CurrentSelfTestCompletion uint8
	_                         [2]byte
	SelfTestResults           [20]selfTestResult
}

type SelfTestResult struct {
	// Op contains the self test type
	Op            SelfTestOp
	Result        uint8
	SegmentNumber uint8
	PowerOnHours  uint64
	NamespaceID   uint32
	FailingLBA    uint64
	Error         Error
}

type SelfTestResults struct {
	// CurrentOp contains the currently in-progress self test type (or
	// SelfTestTypeNone if no self test is in progress).
	CurrentOp SelfTestOp
	// CurrentCompletion contains the progress from 0 to 1 of the currently
	// in-progress self-test. Only valid if CurrentOp is not SelfTestTypeNone.
	CurrentSelfTestCompletion float32
	// PastResults contains a list of up to 20 previous self test results,
	// sorted from the most recent to the oldest.
	PastResults []SelfTestResult
}

func (d *Device) GetSelfTestResults(ns uint32) (*SelfTestResults, error) {
	var buf [564]byte
	if err := d.GetLogPage(ns, 0x06, 0, 0, buf[:]); err != nil {
		return nil, err
	}
	var page selfTestLogPage
	binary.Read(bytes.NewReader(buf[:]), binary.LittleEndian, &page)
	var res SelfTestResults
	res.CurrentOp = SelfTestOp(page.CurrentSelfTestOp & 0xF)
	res.CurrentSelfTestCompletion = float32(page.CurrentSelfTestCompletion&0x7F) / 100.
	for _, r := range page.SelfTestResults {
		var t SelfTestResult
		t.Op = SelfTestOp((r.SelfTestStatus >> 4) & 0xF)
		t.Result = r.SelfTestStatus & 0xF
		if t.Result == 0xF {
			continue
		}
		t.SegmentNumber = r.SegmentNumber
		t.PowerOnHours = r.PowerOnHours
		t.NamespaceID = r.NamespaceID
		t.FailingLBA = r.FailingLBA
		t.Error.StatusCode = r.StatusCode
		t.Error.StatusCodeType = r.StatusCodeType
		res.PastResults = append(res.PastResults, t)
	}
	return &res, nil
}
