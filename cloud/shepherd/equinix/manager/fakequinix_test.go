package manager

import (
	"context"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/packethost/packngo"
)

// fakequinix implements a wrapngo.Client for testing. It starts out with a
// number of made up hardware reservations, and allows for creating devices and
// SSH keys.
type fakequinix struct {
	mu sync.Mutex

	pid          string
	devices      map[string]*packngo.Device
	reservations map[string]*packngo.HardwareReservation
	sshKeys      map[string]*packngo.SSHKey
	reboots      map[string]int
}

// newFakequinix makes a fakequinix with a given fake project ID and number of
// hardware reservations to create.
func newFakequinix(pid string, numReservations int) *fakequinix {
	f := fakequinix{
		pid:          pid,
		devices:      make(map[string]*packngo.Device),
		reservations: make(map[string]*packngo.HardwareReservation),
		sshKeys:      make(map[string]*packngo.SSHKey),
		reboots:      make(map[string]int),
	}

	for i := 0; i < numReservations; i++ {
		uid := uuid.New()
		f.reservations[uid.String()] = &packngo.HardwareReservation{
			ID:            uid.String(),
			ShortID:       uid.String(),
			Provisionable: true,
		}
	}

	return &f
}

func (f *fakequinix) notFound() error {
	return &packngo.ErrorResponse{
		Response: &http.Response{
			StatusCode: http.StatusNotFound,
		},
	}
}

func (f *fakequinix) GetDevice(_ context.Context, pid, did string, _ *packngo.ListOptions) (*packngo.Device, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	val := f.devices[did]
	if val == nil {
		return nil, f.notFound()
	}
	return val, nil
}

func (f *fakequinix) ListDevices(_ context.Context, pid string) ([]packngo.Device, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if pid != f.pid {
		return nil, nil
	}
	var res []packngo.Device
	for _, dev := range f.devices {
		res = append(res, *dev)
	}
	return res, nil
}

func (f *fakequinix) CreateDevice(_ context.Context, request *packngo.DeviceCreateRequest) (*packngo.Device, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	rid := request.HardwareReservationID
	res := f.reservations[rid]
	if res == nil {
		return nil, f.notFound()
	}
	if res.Device != nil {
		return nil, f.notFound()
	}

	dev := &packngo.Device{
		ID:    uuid.New().String(),
		State: "active",
		HardwareReservation: &packngo.HardwareReservation{
			ID: rid,
		},
		Network: []*packngo.IPAddressAssignment{
			{
				IpAddressCommon: packngo.IpAddressCommon{
					Public:  true,
					Address: "1.2.3.4",
				},
			},
		},
		Facility: &packngo.Facility{
			Code: "wad",
		},
		Hostname: request.Hostname,
		OS: &packngo.OS{
			Name: request.OS,
			Slug: request.OS,
		},
	}
	res.Device = dev
	res.Provisionable = false

	f.devices[dev.ID] = dev
	return dev, nil
}

func (f *fakequinix) ListReservations(_ context.Context, pid string) ([]packngo.HardwareReservation, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	var res []packngo.HardwareReservation
	for _, r := range f.reservations {
		res = append(res, *r)
	}

	return res, nil
}

func (f *fakequinix) ListSSHKeys(_ context.Context) ([]packngo.SSHKey, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	var res []packngo.SSHKey
	for _, key := range f.sshKeys {
		res = append(res, *key)
	}

	return res, nil
}

func (f *fakequinix) CreateSSHKey(_ context.Context, req *packngo.SSHKeyCreateRequest) (*packngo.SSHKey, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	for _, k := range f.sshKeys {
		if k.Key == req.Key {
			return nil, f.notFound()
		}
		if k.Label == req.Label {
			return nil, f.notFound()
		}
	}

	uid := uuid.New().String()
	f.sshKeys[uid] = &packngo.SSHKey{
		ID:    uid,
		Label: req.Label,
		Key:   req.Key,
	}

	return f.sshKeys[uid], nil
}

func (f *fakequinix) UpdateSSHKey(_ context.Context, kid string, req *packngo.SSHKeyUpdateRequest) (*packngo.SSHKey, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	key := f.sshKeys[kid]
	if key == nil {
		return nil, f.notFound()
	}
	key.Key = *req.Key

	return key, nil
}

func (f *fakequinix) RebootDevice(_ context.Context, did string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.reboots[did]++

	return nil
}

func (f *fakequinix) Close() {
}
