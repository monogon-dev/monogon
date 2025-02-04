// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"net/netip"
	"slices"
	"strings"
	"time"

	"github.com/packethost/packngo"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/equinix/wrapngo"
	"source.monogon.dev/cloud/lib/sinbin"
	"source.monogon.dev/cloud/shepherd"
	"source.monogon.dev/cloud/shepherd/manager"
)

type equinixProvider struct {
	config *providerConfig
	api    wrapngo.Client
	sshKey *manager.SSHKey

	// badReservations is a holiday resort for Equinix hardware reservations which
	// failed to be provisioned for some reason or another. We keep a list of them in
	// memory just so that we don't repeatedly try to provision the same known bad
	// machines.
	badReservations sinbin.Sinbin[string]

	reservationDeadline time.Time
	reservationCache    []packngo.HardwareReservation
}

func (ep *equinixProvider) RebootMachine(ctx context.Context, id shepherd.ProviderID) error {
	if err := ep.api.RebootDevice(ctx, string(id)); err != nil {
		return fmt.Errorf("failed to reboot device: %w", err)
	}

	// TODO(issue/215): replace this
	// This is required as Equinix doesn't reboot the machines synchronously
	// during the API call.
	select {
	case <-time.After(time.Duration(ep.config.RebootWaitSeconds) * time.Second):
	case <-ctx.Done():
		return fmt.Errorf("while waiting for reboot: %w", ctx.Err())
	}
	return nil
}

func (ep *equinixProvider) ReinstallMachine(ctx context.Context, id shepherd.ProviderID) error {
	return shepherd.ErrNotImplemented
}

func (ep *equinixProvider) GetMachine(ctx context.Context, id shepherd.ProviderID) (shepherd.Machine, error) {
	machines, err := ep.ListMachines(ctx)
	if err != nil {
		return nil, err
	}

	for _, machine := range machines {
		if machine.ID() == id {
			return machine, nil
		}
	}

	return nil, shepherd.ErrMachineNotFound
}

func (ep *equinixProvider) ListMachines(ctx context.Context) ([]shepherd.Machine, error) {
	if ep.reservationDeadline.Before(time.Now()) {
		reservations, err := ep.listReservations(ctx)
		if err != nil {
			return nil, err
		}
		ep.reservationCache = reservations
		ep.reservationDeadline = time.Now().Add(ep.config.ReservationCacheTimeout)
	}

	devices, err := ep.managedDevices(ctx)
	if err != nil {
		return nil, err
	}

	machines := make([]shepherd.Machine, 0, len(ep.reservationCache)+len(devices))
	for _, device := range devices {
		machines = append(machines, &machine{device})
	}

	for _, res := range ep.reservationCache {
		machines = append(machines, reservation{res})
	}

	return machines, nil
}

func (ep *equinixProvider) CreateMachine(ctx context.Context, session *bmdb.Session, request shepherd.CreateMachineRequest) (shepherd.Machine, error) {
	if request.UnusedMachine == nil {
		return nil, fmt.Errorf("parameter UnusedMachine is missing")
	}

	//TODO: Do we just trust the implementation to be correct?
	res, ok := request.UnusedMachine.(reservation)
	if !ok {
		return nil, fmt.Errorf("invalid type for parameter UnusedMachine")
	}

	d, err := ep.provision(ctx, session, res.HardwareReservation)
	if err != nil {
		klog.Errorf("Failed to provision reservation %s: %v", res.HardwareReservation.ID, err)
		until := time.Now().Add(time.Hour)
		klog.Errorf("Adding hardware reservation %s to sinbin until %s", res.HardwareReservation.ID, until)
		ep.badReservations.Add(res.HardwareReservation.ID, until)
		return nil, err
	}

	return &machine{*d}, nil
}

func (ep *equinixProvider) Type() model.Provider {
	return model.ProviderEquinix
}

type reservation struct {
	packngo.HardwareReservation
}

func (e reservation) Failed() bool {
	return false
}

func (e reservation) ID() shepherd.ProviderID {
	return shepherd.InvalidProviderID
}

func (e reservation) Addr() netip.Addr {
	return netip.Addr{}
}

func (e reservation) Availability() shepherd.Availability {
	return shepherd.AvailabilityKnownUnused
}

type machine struct {
	packngo.Device
}

func (e *machine) Failed() bool {
	return e.State == "failed"
}

func (e *machine) ID() shepherd.ProviderID {
	return shepherd.ProviderID(e.Device.ID)
}

func (e *machine) Addr() netip.Addr {
	ni := e.GetNetworkInfo()

	var addr string
	if ni.PublicIPv4 != "" {
		addr = ni.PublicIPv4
	} else if ni.PublicIPv6 != "" {
		addr = ni.PublicIPv6
	} else {
		klog.Errorf("missing address for machine: %v", e.ID())
		return netip.Addr{}
	}

	a, err := netip.ParseAddr(addr)
	if err != nil {
		klog.Errorf("failed parsing address %q: %v", addr, err)
		return netip.Addr{}
	}

	return a
}

func (e *machine) Availability() shepherd.Availability {
	return shepherd.AvailabilityKnownUsed
}

// listReservations doesn't lock the mutex and expects the caller to lock.
func (ep *equinixProvider) listReservations(ctx context.Context) ([]packngo.HardwareReservation, error) {
	klog.Infof("Retrieving hardware reservations, this will take a while...")
	reservations, err := ep.api.ListReservations(ctx, ep.config.ProjectId)
	if err != nil {
		return nil, fmt.Errorf("failed to list reservations: %w", err)
	}

	var available []packngo.HardwareReservation
	var inUse, notProvisionable, penalized int
	for _, reservation := range reservations {
		if reservation.Device != nil {
			inUse++
			continue
		}
		if !reservation.Provisionable {
			notProvisionable++
			continue
		}
		if ep.badReservations.Penalized(reservation.ID) {
			penalized++
			continue
		}
		available = append(available, reservation)
	}
	klog.Infof("Retrieved hardware reservations: %d (total), %d (available), %d (in use), %d (not provisionable), %d (penalized)", len(reservations), len(available), inUse, notProvisionable, penalized)

	return available, nil
}

// provision attempts to create a device within Equinix using given Hardware
// Reservation rsv. The resulting device is registered with BMDB, and tagged as
// "provided" in the process.
func (ep *equinixProvider) provision(ctx context.Context, sess *bmdb.Session, rsv packngo.HardwareReservation) (*packngo.Device, error) {
	klog.Infof("Creating a new device using reservation ID %s.", rsv.ID)
	hostname := ep.config.DevicePrefix + rsv.ID[:18]
	kid, err := ep.sshEquinixId(ctx)
	if err != nil {
		return nil, err
	}
	req := &packngo.DeviceCreateRequest{
		Hostname:              hostname,
		OS:                    ep.config.OS,
		Plan:                  rsv.Plan.Slug,
		ProjectID:             ep.config.ProjectId,
		HardwareReservationID: rsv.ID,
		ProjectSSHKeys:        []string{kid},
	}
	if ep.config.UseProjectKeys {
		klog.Warningf("INSECURE: Machines will be created with ALL PROJECT SSH KEYS!")
		req.ProjectSSHKeys = nil
	}

	nd, err := ep.api.CreateDevice(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("while creating new device within Equinix: %w", err)
	}
	klog.Infof("Created a new device within Equinix (RID: %s, PID: %s, HOST: %s)", rsv.ID, nd.ID, hostname)

	ep.reservationCache = slices.DeleteFunc(ep.reservationCache, func(v packngo.HardwareReservation) bool {
		return rsv.ID == v.ID
	})

	err = ep.assimilate(ctx, sess, nd.ID)
	if err != nil {
		// TODO(serge@monogon.tech) at this point the device at Equinix isn't
		// matched by a BMDB record. Schedule device deletion or make sure this
		// case is being handled elsewhere.
		return nil, err
	}
	return nd, nil
}

// assimilate brings in an already existing machine from Equinix into the BMDB.
// This is only used in manual testing.
func (ep *equinixProvider) assimilate(ctx context.Context, sess *bmdb.Session, deviceID string) error {
	return sess.Transact(ctx, func(q *model.Queries) error {
		// Create a new machine record within BMDB.
		m, err := q.NewMachine(ctx)
		if err != nil {
			return fmt.Errorf("while creating a new machine record in BMDB: %w", err)
		}

		// Link the new machine with the Equinix device, and tag it "provided".
		p := model.MachineAddProvidedParams{
			MachineID:  m.MachineID,
			ProviderID: deviceID,
			Provider:   model.ProviderEquinix,
		}
		klog.Infof("Setting \"provided\" tag (ID: %s, PID: %s, Provider: %s).", p.MachineID, p.ProviderID, p.Provider)
		if err := q.MachineAddProvided(ctx, p); err != nil {
			return fmt.Errorf("while tagging machine active: %w", err)
		}
		return nil
	})
}

// sshEquinixGet looks up the Equinix key matching providerConfig.KeyLabel,
// returning its packngo.SSHKey instance.
func (ep *equinixProvider) sshEquinix(ctx context.Context) (*packngo.SSHKey, error) {
	ks, err := ep.api.ListSSHKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("while listing SSH keys: %w", err)
	}

	for _, k := range ks {
		if k.Label == ep.config.KeyLabel {
			return &k, nil
		}
	}
	return nil, ErrNoSuchKey
}

// sshEquinixId looks up the Equinix key identified by providerConfig.KeyLabel,
// returning its Equinix-assigned UUID.
func (ep *equinixProvider) sshEquinixId(ctx context.Context) (string, error) {
	k, err := ep.sshEquinix(ctx)
	if err != nil {
		return "", err
	}
	return k.ID, nil
}

// sshEquinixUpdate makes sure the existing SSH key registered with Equinix
// matches the one from sshPub.
func (ep *equinixProvider) sshEquinixUpdate(ctx context.Context, kid string) error {
	pub, err := ep.sshKey.PublicKey()
	if err != nil {
		return err
	}
	_, err = ep.api.UpdateSSHKey(ctx, kid, &packngo.SSHKeyUpdateRequest{
		Key: &pub,
	})
	if err != nil {
		return fmt.Errorf("while updating the SSH key: %w", err)
	}
	return nil
}

// sshEquinixUpload registers a new SSH key from sshPub.
func (ep *equinixProvider) sshEquinixUpload(ctx context.Context) error {
	pub, err := ep.sshKey.PublicKey()
	if err != nil {
		return fmt.Errorf("while generating public key: %w", err)
	}
	_, err = ep.api.CreateSSHKey(ctx, &packngo.SSHKeyCreateRequest{
		Label:     ep.config.KeyLabel,
		Key:       pub,
		ProjectID: ep.config.ProjectId,
	})
	if err != nil {
		return fmt.Errorf("while creating an SSH key: %w", err)
	}
	return nil
}

// SSHEquinixEnsure initializes the locally managed SSH key (from a persistence
// path or explicitly set key) and updates or uploads it to Equinix. The key is
// generated as needed The key is generated as needed
func (ep *equinixProvider) SSHEquinixEnsure(ctx context.Context) error {
	k, err := ep.sshEquinix(ctx)
	switch {
	case errors.Is(err, ErrNoSuchKey):
		if err := ep.sshEquinixUpload(ctx); err != nil {
			return fmt.Errorf("while uploading key: %w", err)
		}
		return nil
	case err == nil:
		if err := ep.sshEquinixUpdate(ctx, k.ID); err != nil {
			return fmt.Errorf("while updating key: %w", err)
		}
		return nil
	default:
		return err
	}
}

// managedDevices provides a map of device provider IDs to matching
// packngo.Device instances. It calls Equinix API's ListDevices. The returned
// devices are filtered according to DevicePrefix provided through Opts. The
// returned error value, if not nil, will originate in wrapngo.
func (ep *equinixProvider) managedDevices(ctx context.Context) (map[string]packngo.Device, error) {
	ds, err := ep.api.ListDevices(ctx, ep.config.ProjectId)
	if err != nil {
		return nil, err
	}
	dm := map[string]packngo.Device{}
	for _, d := range ds {
		if strings.HasPrefix(d.Hostname, ep.config.DevicePrefix) {
			dm[d.ID] = d
		}
	}
	return dm, nil
}
