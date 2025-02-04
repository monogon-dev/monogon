// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"

	"source.monogon.dev/cloud/equinix/wrapngo"
	"source.monogon.dev/cloud/shepherd/manager"
)

var (
	ErrNoSuchKey = errors.New("no such key")
)

// providerConfig contains configuration options used by both the Initializer and
// Provisioner components of the Shepherd. In CLI scenarios, RegisterFlags should
// be called to configure this struct from CLI flags. Otherwise, this structure
// should be explicitly configured, as the default values are not valid.
type providerConfig struct {
	// ProjectId is the Equinix project UUID used by the manager. See Equinix API
	// documentation for details. Must be set.
	ProjectId string

	// KeyLabel specifies the ID to use when handling the Equinix-registered SSH
	// key used to authenticate to newly created servers. Must be set.
	KeyLabel string

	// DevicePrefix applied to all devices (machines) created by the Provisioner,
	// and used by the Provisioner to identify machines which it managed.
	// Must be set.
	DevicePrefix string

	// OS defines the operating system new devices are created with. Its format
	// is specified by Equinix API.
	OS string

	// UseProjectKeys defines if the provisioner adds all ssh keys defined inside
	// the used project to every new machine. This is only used for debug purposes.
	UseProjectKeys bool

	// RebootWaitSeconds defines how many seconds to sleep after a reboot call
	// to ensure a reboot actually happened.
	RebootWaitSeconds int

	// ReservationCacheTimeout defines how after which time the reservations should be
	// refreshed.
	ReservationCacheTimeout time.Duration
}

func (pc *providerConfig) check() error {
	if pc.ProjectId == "" {
		return fmt.Errorf("-equinix_project_id must be set")
	}
	if pc.KeyLabel == "" {
		return fmt.Errorf("-equinix_ssh_key_label must be set")
	}
	if pc.DevicePrefix == "" {
		return fmt.Errorf("-equinix_device_prefix must be set")
	}

	// These variables are _very_ important to configure correctly, otherwise someone
	// running this locally with prod creds will actually destroy production
	// data.
	if strings.Contains(pc.KeyLabel, "FIXME") {
		return fmt.Errorf("refusing to run with -equinix_ssh_key_label %q, please set it to something unique", pc.KeyLabel)
	}
	if strings.Contains(pc.DevicePrefix, "FIXME") {
		return fmt.Errorf("refusing to run with -equinix_device_prefix %q, please set it to something unique", pc.DevicePrefix)
	}

	return nil
}

func (pc *providerConfig) RegisterFlags() {
	flag.StringVar(&pc.ProjectId, "equinix_project_id", "", "Equinix project ID where resources will be managed")
	flag.StringVar(&pc.KeyLabel, "equinix_ssh_key_label", "shepherd-FIXME", "Label used to identify managed SSH key in Equinix project")
	flag.StringVar(&pc.DevicePrefix, "equinix_device_prefix", "shepherd-FIXME-", "Prefix applied to all devices (machines) in Equinix project, used to identify managed machines")
	flag.StringVar(&pc.OS, "equinix_os", "ubuntu_20_04", "OS that provisioner will deploy on Equinix machines. Not the target OS for cluster customers.")
	flag.BoolVar(&pc.UseProjectKeys, "equinix_use_project_keys", false, "Add all Equinix project keys to newly provisioned machines, not just the provisioner's managed key. Debug/development only.")
	flag.IntVar(&pc.RebootWaitSeconds, "equinix_reboot_wait_seconds", 30, "How many seconds to sleep to ensure a reboot happend")
	flag.DurationVar(&pc.ReservationCacheTimeout, "equinix_reservation_cache_timeout", time.Minute*15, "Reservation cache validity timeo")
}

func (pc *providerConfig) New(sshKey *manager.SSHKey, api wrapngo.Client) (*equinixProvider, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}

	return &equinixProvider{
		config: pc,
		sshKey: sshKey,
		api:    api,
	}, nil
}
