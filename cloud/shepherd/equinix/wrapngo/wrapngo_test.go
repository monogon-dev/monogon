package wrapngo

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/packethost/packngo"
	"golang.org/x/crypto/ssh"
)

var (
	ctx context.Context

	// apiuser and apikey are the Equinix credentials necessary to test the
	// client package.

	apiuser string = os.Getenv("EQUINIX_USER")
	apikey  string = os.Getenv("EQUINIX_APIKEY")

	// apipid references the Equinix Metal project used. It's recommended to use
	// non-production projects in the context of testing.
	apipid string = os.Getenv("EQUINIX_PROJECT_ID")
	// apios specifies the operating system installed on newly provisioned
	// devices. See Equinix Metal API documentation for details.
	apios string = os.Getenv("EQUINIX_DEVICE_OS")

	// sshKeyID identifies the SSH public key registered with Equinix.
	sshKeyID string
	// sshKeyLabel is the label used to register the SSH key with Equinix.
	sshKeyLabel string = "shepherd-client-testkey"

	// testDevice is the device created in TestCreateDevice, and later
	// referenced to exercise implementation operating on Equinix Metal device
	// objects.
	testDevice *packngo.Device
	// testDeviceHostname is the hostname used to register and reference the
	// test device.
	testDeviceHostname string = "shepherd-client-testdev"
)

// ensureParams returns false if any of the required environment variable
// parameters are missing.
func ensureParams() bool {
	if apiuser == "" {
		log.Print("EQUINIX_USER must be set.")
		return false
	}
	if apikey == "" {
		log.Print("EQUINIX_APIKEY must be set.")
		return false
	}
	if apipid == "" {
		log.Print("EQUINIX_PROJECT_ID must be set.")
		return false
	}
	if apios == "" {
		log.Print("EQUINIX_DEVICE_OS must be set.")
		return false
	}
	return true
}

// awaitDeviceState returns nil after device matching the id reaches one of the
// provided states. It will return a non-nil value in case of an API error, and
// particularly if there exists no device matching id.
func awaitDeviceState(ctx context.Context, t *testing.T, cl *client, id string, states ...string) error {
	if t != nil {
		t.Helper()
	}

	for {
		d, err := cl.GetDevice(ctx, apipid, id)
		if err != nil {
			if errors.Is(err, os.ErrDeadlineExceeded) {
				continue
			}
			return fmt.Errorf("while fetching device info: %w", err)
		}
		if d == nil {
			return fmt.Errorf("expected the test device (ID: %s) to exist.", id)
		}
		for _, s := range states {
			if d.State == s {
				return nil
			}
		}
		log.Printf("Waiting for device to be provisioned (ID: %s, current state: %q)", id, d.State)
		time.Sleep(time.Second)
	}
}

// cleanup ensures both the test device and the test key are deleted at
// Equinix.
func cleanup(ctx context.Context, cl *client) {
	log.Print("Cleaning up.")

	// Ensure the device matching testDeviceHostname is deleted.
	ds, err := cl.ListDevices(ctx, apipid)
	if err != nil {
		log.Fatalf("while listing devices: %v", err)
	}
	var td *packngo.Device
	for _, d := range ds {
		if d.Hostname == testDeviceHostname {
			td = &d
			break
		}
	}
	if td != nil {
		log.Printf("Found a test device (ID: %s) that needs to be deleted before progressing further.", td.ID)

		// Devices currently being provisioned can't be deleted. After it's
		// provisioned, device's state will match either "active", or "failed".
		if err := awaitDeviceState(ctx, nil, cl, td.ID, "active", "failed"); err != nil {
			log.Fatalf("while waiting for device to be provisioned: %v", err)
		}
		if err := cl.deleteDevice(ctx, td.ID); err != nil {
			log.Fatalf("while deleting test device: %v", err)
		}
	}

	// Ensure the key matching sshKeyLabel is deleted.
	ks, err := cl.ListSSHKeys(ctx)
	if err != nil {
		log.Fatalf("while listing SSH keys: %v", err)
	}
	for _, k := range ks {
		if k.Label == sshKeyLabel {
			log.Printf("Found a SSH test key (ID: %s) - deleting...", k.ID)
			if err := cl.deleteSSHKey(ctx, k.ID); err != nil {
				log.Fatalf("while deleting an SSH key: %v", err)
			}
			log.Printf("Deleted a SSH test key (ID: %s).", k.ID)
		}
	}
}

func TestMain(m *testing.M) {
	if !ensureParams() {
		log.Print("Skipping due to missing parameters.")
		return
	}
	ctx = context.Background()

	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})
	defer cl.Close()

	cleanup(ctx, cl)
	code := m.Run()
	cleanup(ctx, cl)
	os.Exit(code)
}

// Most test cases depend on the preceding cases having been executed. The
// test cases can't be run in parallel.

func TestListReservations(t *testing.T) {
	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	_, err := cl.ListReservations(ctx, apipid)
	if err != nil {
		t.Errorf("while listing hardware reservations: %v", err)
	}
}

// createSSHAuthKey returns an SSH public key in OpenSSH authorized_keys
// format.
func createSSHAuthKey(t *testing.T) string {
	t.Helper()
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("while generating SSH key: %v", err)
	}

	sshpub, err := ssh.NewPublicKey(pub)
	if err != nil {
		t.Errorf("while generating SSH public key: %v", err)
	}
	return string(ssh.MarshalAuthorizedKey(sshpub))
}

func TestCreateSSHKey(t *testing.T) {
	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	nk, err := cl.CreateSSHKey(ctx, &packngo.SSHKeyCreateRequest{
		Label:     sshKeyLabel,
		Key:       createSSHAuthKey(t),
		ProjectID: apipid,
	})
	if err != nil {
		t.Errorf("while creating an SSH key: %v", err)
	}
	if nk.Label != sshKeyLabel {
		t.Errorf("key labels don't match.")
	}
	t.Logf("Created an SSH key (ID: %s)", nk.ID)
	sshKeyID = nk.ID
}

var (
	// dummySSHPK2 is the alternate key used to exercise TestUpdateSSHKey and
	// TestGetSSHKey.
	dummySSHPK2 string
)

func TestUpdateSSHKey(t *testing.T) {
	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	if sshKeyID == "" {
		t.Skip("SSH key couldn't have been created - skipping...")
	}

	dummySSHPK2 = createSSHAuthKey(t)
	k, err := cl.UpdateSSHKey(ctx, sshKeyID, &packngo.SSHKeyUpdateRequest{
		Key: &dummySSHPK2,
	})
	if err != nil {
		t.Errorf("while updating an SSH key: %v", err)
	}
	if k.Key != dummySSHPK2 {
		t.Errorf("updated SSH key doesn't match the original.")
	}
}

func TestGetSSHKey(t *testing.T) {
	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	if sshKeyID == "" {
		t.Skip("SSH key couldn't have been created - skipping...")
	}

	k, err := cl.getSSHKey(ctx, sshKeyID)
	if err != nil {
		t.Errorf("while getting an SSH key: %v", err)
	}
	if k.Key != dummySSHPK2 {
		t.Errorf("got key contents that don't match the original.")
	}
}

func TestListSSHKeys(t *testing.T) {
	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	if sshKeyID == "" {
		t.Skip("SSH key couldn't have been created - skipping...")
	}

	ks, err := cl.ListSSHKeys(ctx)
	if err != nil {
		t.Errorf("while listing SSH keys: %v", err)
	}

	// Check that our key is part of the list.
	found := false
	for _, k := range ks {
		if k.ID == sshKeyID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("SSH key not listed.")
	}
}

func TestCreateDevice(t *testing.T) {
	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	// Find a provisionable hardware reservation the device will be created with.
	rvs, err := cl.ListReservations(ctx, apipid)
	if err != nil {
		t.Errorf("while listing hardware reservations: %v", err)
	}
	var rv *packngo.HardwareReservation
	for _, r := range rvs {
		if r.Provisionable {
			rv = &r
			break
		}
	}
	if rv == nil {
		t.Skip("could not find a provisionable hardware reservation - skipping...")
	}

	d, err := cl.CreateDevice(ctx, &packngo.DeviceCreateRequest{
		Hostname:              testDeviceHostname,
		OS:                    apios,
		Plan:                  rv.Plan.Slug,
		HardwareReservationID: rv.ID,
		ProjectID:             apipid,
	})
	if err != nil {
		t.Errorf("while creating a device: %v", err)
	}
	t.Logf("Created a new test device (ID: %s)", d.ID)
	testDevice = d
}

func TestGetDevice(t *testing.T) {
	if testDevice == nil {
		t.Skip("the test device couldn't have been created - skipping...")
	}

	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	d, err := cl.GetDevice(ctx, apipid, testDevice.ID)
	if err != nil {
		t.Errorf("while fetching device info: %v", err)
	}
	if d == nil {
		t.Errorf("expected the test device (ID: %s) to exist.", testDevice.ID)
	}
	if d.ID != testDevice.ID {
		t.Errorf("got device ID that doesn't match the original.")
	}
}

func TestListDevices(t *testing.T) {
	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	ds, err := cl.ListDevices(ctx, apipid)
	if err != nil {
		t.Errorf("while listing devices: %v", err)
	}
	if len(ds) == 0 {
		t.Errorf("expected at least one device.")
	}
}

func TestDeleteDevice(t *testing.T) {
	if testDevice == nil {
		t.Skip("the test device couldn't have been created - skipping...")
	}

	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	// Devices currently being provisioned can't be deleted. After it's
	// provisioned, device's state will match either "active", or "failed".
	if err := awaitDeviceState(ctx, t, cl, testDevice.ID, "active", "failed"); err != nil {
		t.Errorf("while waiting for device to be provisioned: %v", err)
	}
	t.Logf("Deleting the test device (ID: %s)", testDevice.ID)
	if err := cl.deleteDevice(ctx, testDevice.ID); err != nil {
		t.Errorf("while deleting a device: %v", err)
	}
	d, err := cl.GetDevice(ctx, apipid, testDevice.ID)
	if err != nil && !IsNotFound(err) {
		t.Errorf("while fetching device info: %v", err)
	}
	if d != nil {
		t.Errorf("device should not exist.")
	}
	t.Logf("Deleted the test device (ID: %s)", testDevice.ID)
}

func TestDeleteSSHKey(t *testing.T) {
	if sshKeyID == "" {
		t.Skip("SSH key couldn't have been created - skipping...")
	}

	cl := new(&Opts{
		User:   apiuser,
		APIKey: apikey,
	})

	t.Logf("Deleting the test SSH key (ID: %s)", sshKeyID)
	if err := cl.deleteSSHKey(ctx, sshKeyID); err != nil {
		t.Errorf("couldn't delete an SSH key: %v", err)
	}
	_, err := cl.getSSHKey(ctx, sshKeyID)
	if err == nil {
		t.Errorf("SSH key should not exist")
	}
	t.Logf("Deleted the test SSH key (ID: %s)", sshKeyID)
}
