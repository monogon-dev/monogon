// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package scsi

import "errors"

type InformationalExceptions struct {
	InformationalSenseCode AdditionalSenseCode
	Temperature            uint8
}

func (d *Device) GetInformationalExceptions() (*InformationalExceptions, error) {
	raw, err := d.LogSenseParameters(LogSenseRequest{PageCode: 0x0b})
	if err != nil {
		return nil, err
	}
	if len(raw[0x1]) == 0 {
		return nil, errors.New("mandatory parameter 0001h missing")
	}
	param1 := raw[0x01][0]
	if len(param1.Data) < 3 {
		return nil, errors.New("parameter 0001h too short")
	}
	return &InformationalExceptions{
		InformationalSenseCode: AdditionalSenseCode(uint16(param1.Data[0])<<8 | uint16(param1.Data[1])),
		Temperature:            param1.Data[2],
	}, nil
}
