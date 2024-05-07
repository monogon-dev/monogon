package scsi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

type LogSenseRequest struct {
	// PageCode contains the identifier of the requested page
	PageCode uint8
	// SubpageCode contains the identifier of the requested subpage
	// or the zero value if no subpage is requested.
	SubpageCode uint8
	// PageControl specifies what type of values should be returned for bounded
	// and unbounded log parameters. See also Table 156 in the standard.
	PageControl uint8
	// ParameterPointer allows requesting parameter data beginning from a
	// specific parameter code. The zero value starts from the beginning.
	ParameterPointer uint16
	// SaveParameters requests the device to save all parameters without
	// DisableUpdate set to non-volatile storage.
	SaveParameters bool
	// InitialSize is an optional hint how big the buffer for the log page
	// should be for the initial request. The zero value sets this to 4096.
	InitialSize uint16
}

// LogSenseRaw requests a raw log page. For log pages with parameters
// LogSenseParameters is better-suited.
func (d *Device) LogSenseRaw(r LogSenseRequest) ([]byte, error) {
	var bufferSize uint16 = 4096
	for {
		data := make([]byte, bufferSize)
		var req [8]byte
		if r.SaveParameters {
			req[0] = 0b1
		}
		req[1] = r.PageControl<<6 | r.PageCode
		req[2] = r.SubpageCode
		binary.BigEndian.PutUint16(req[4:6], r.ParameterPointer)
		binary.BigEndian.PutUint16(req[6:8], uint16(len(data)))
		if err := d.RawCommand(&CommandDataBuffer{
			OperationCode:         LogSenseOp,
			Request:               req[:],
			Data:                  data,
			DataTransferDirection: DataTransferFromDevice,
		}); err != nil {
			return nil, fmt.Errorf("error during LOG SENSE: %w", err)
		}
		if data[0]&0b111111 != r.PageCode {
			return nil, fmt.Errorf("requested log page %x, got %x", r.PageCode, data[1])
		}
		if data[1] != r.SubpageCode {
			return nil, fmt.Errorf("requested log subpage %x, got %x", r.SubpageCode, data[1])
		}
		pageLength := binary.BigEndian.Uint16(data[2:4])
		if pageLength > math.MaxUint16-4 {
			// Guard against uint16 overflows, this cannot be requested anyways
			return nil, fmt.Errorf("device log page is too long (%d bytes)", pageLength)
		}
		if pageLength > uint16(len(data)-4) {
			bufferSize = pageLength + 4
			continue
		}
		return data[4 : pageLength+4], nil
	}
}

// SupportedLogPages returns a map with all supported log pages.
// This can return an error if the device does not support logs at all.
func (d *Device) SupportedLogPages() (map[uint8]bool, error) {
	raw, err := d.LogSenseRaw(LogSenseRequest{PageCode: 0})
	if err != nil {
		return nil, err
	}
	res := make(map[uint8]bool)
	for _, r := range raw {
		res[r] = true
	}
	return res, nil
}

// PageAndSubpage identifies a log page uniquely.
type PageAndSubpage uint16

func NewPageAndSubpage(page, subpage uint8) PageAndSubpage {
	return PageAndSubpage(uint16(page)<<8 | uint16(subpage))
}

func (p PageAndSubpage) Page() uint8 {
	return uint8(p >> 8)
}
func (p PageAndSubpage) Subpage() uint8 {
	return uint8(p & 0xFF)
}

func (p PageAndSubpage) String() string {
	return fmt.Sprintf("Page %xh Subpage %xh", p.Page(), p.Subpage())
}

// SupportedLogPagesAndSubpages returns the list of supported pages and subpages.
// This can return an error if the device does not support logs at all.
func (d *Device) SupportedLogPagesAndSubpages() (map[PageAndSubpage]bool, error) {
	raw, err := d.LogSenseRaw(LogSenseRequest{PageCode: 0x00, SubpageCode: 0xff})
	if err != nil {
		return nil, err
	}
	res := make(map[PageAndSubpage]bool)
	for i := 0; i < len(raw)/2; i++ {
		res[NewPageAndSubpage(raw[i*2], raw[(i*2)+1])] = true
	}
	return res, nil
}

// SupportedLogSubPages returns the list of subpages supported in a log page.
func (d *Device) SupportedLogSubPages(pageCode uint8) (map[uint8]bool, error) {
	raw, err := d.LogSenseRaw(LogSenseRequest{PageCode: pageCode, SubpageCode: 0xff})
	if err != nil {
		return nil, err
	}
	res := make(map[uint8]bool)
	for _, r := range raw {
		res[r] = true
	}
	return res, nil
}

type LogParameter struct {
	// DisableUpdate indicates if the device is updating this parameter.
	// If this is true the parameter has either overflown or updating has been
	// manually disabled.
	DisableUpdate bool
	// TargetSaveDisable indicates if automatic saving of this parameter has
	// been disabled.
	TargetSaveDisable bool
	// FormatAndLinking contains the format of the log parameter.
	FormatAndLinking uint8
	// Data contains the payload of the log parameter.
	Data []byte
}

// LogSenseParameters returns the parameters of a log page. The returned map
// contains one entry per parameter ID in the result. The order of parameters
// of the same ID is kept.
func (d *Device) LogSenseParameters(r LogSenseRequest) (map[uint16][]LogParameter, error) {
	raw, err := d.LogSenseRaw(r)
	if err != nil {
		return nil, err
	}
	res := make(map[uint16][]LogParameter)
	for {
		if len(raw) == 0 {
			break
		}
		if len(raw) < 4 {
			return nil, errors.New("not enough data left to read full parameter metadata")
		}
		var param LogParameter
		parameterCode := binary.BigEndian.Uint16(raw[0:2])
		param.DisableUpdate = raw[2]&(1<<7) != 0
		param.TargetSaveDisable = raw[2]&(1<<5) != 0
		param.FormatAndLinking = raw[2] & 0b11
		if int(raw[3]) > len(raw)-4 {
			fmt.Println(raw[3], len(raw))
			return nil, errors.New("unable to read parameter, not enough data for indicated length")
		}
		param.Data = raw[4 : int(raw[3])+4]
		raw = raw[int(raw[3])+4:]
		res[parameterCode] = append(res[parameterCode], param)
	}
	return res, nil
}
