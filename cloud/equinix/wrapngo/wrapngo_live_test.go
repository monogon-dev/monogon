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

type liveTestClient struct {
	cl  *client
	ctx context.Context

	apipid string
	apios  string

	sshKeyLabel        string
	testDeviceHostname string
}

func newLiveTestClient(t *testing.T) *liveTestClient {
	t.Helper()

	apiuser := os.Getenv("EQUINIX_USER")
	apikey := os.Getenv("EQUINIX_APIKEY")
	apipid := os.Getenv("EQUINIX_PROJECT_ID")
	apios := os.Getenv("EQUINIX_DEVICE_OS")

	if apiuser == "" {
		t.Skip("EQUINIX_USER must be set.")
	}
	if apikey == "" {
		t.Skip("EQUINIX_APIKEY must be set.")
	}
	if apipid == "" {
		t.Skip("EQUINIX_PROJECT_ID must be set.")
	}
	if apios == "" {
		t.Skip("EQUINIX_DEVICE_OS must be set.")
	}
	ctx, ctxC := context.WithCancel(context.Background())
	t.Cleanup(ctxC)
	return &liveTestClient{
		cl: new(&Opts{
			User:   apiuser,
			APIKey: apikey,
		}),
		ctx: ctx,

		apipid: apipid,
		apios:  apios,

		sshKeyLabel:        "shepherd-livetest-client",
		testDeviceHostname: "shepherd-livetest-device",
	}
}

// awaitDeviceState returns nil after device matching the id reaches one of the
// provided states. It will return a non-nil value in case of an API error, and
// particularly if there exists no device matching id.
func (l *liveTestClient) awaitDeviceState(t *testing.T, id string, states ...string) error {
	t.Helper()

	for {
		d, err := l.cl.GetDevice(l.ctx, l.apipid, id, nil)
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
		t.Logf("Waiting for device to be provisioned (ID: %s, current state: %q)", id, d.State)
		time.Sleep(time.Second)
	}
}

// cleanup ensures both the test device and the test key are deleted at
// Equinix.
func (l *liveTestClient) cleanup(t *testing.T) {
	t.Helper()

	t.Logf("Cleaning up.")

	// Ensure the device matching testDeviceHostname is deleted.
	ds, err := l.cl.ListDevices(l.ctx, l.apipid)
	if err != nil {
		log.Fatalf("while listing devices: %v", err)
	}
	var td *packngo.Device
	for _, d := range ds {
		if d.Hostname == l.testDeviceHostname {
			td = &d
			break
		}
	}
	if td != nil {
		t.Logf("Found a test device (ID: %s) that needs to be deleted before progressing further.", td.ID)

		// Devices currently being provisioned can't be deleted. After it's
		// provisioned, device's state will match either "active", or "failed".
		if err := l.awaitDeviceState(t, "active", "failed"); err != nil {
			t.Fatalf("while waiting for device to be provisioned: %v", err)
		}
		if err := l.cl.deleteDevice(l.ctx, td.ID); err != nil {
			t.Fatalf("while deleting test device: %v", err)
		}
	}

	// Ensure the key matching sshKeyLabel is deleted.
	ks, err := l.cl.ListSSHKeys(l.ctx)
	if err != nil {
		t.Fatalf("while listing SSH keys: %v", err)
	}
	for _, k := range ks {
		if k.Label == l.sshKeyLabel {
			t.Logf("Found a SSH test key (ID: %s) - deleting...", k.ID)
			if err := l.cl.deleteSSHKey(l.ctx, k.ID); err != nil {
				t.Fatalf("while deleting an SSH key: %v", err)
			}
			t.Logf("Deleted a SSH test key (ID: %s).", k.ID)
		}
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

// TestLiveAPI performs smoke tests of wrapngo against the real Equinix API. See
// newLiveTestClient to see which environment variables need to be provided in
// order for this test to run.
func TestLiveAPI(t *testing.T) {
	ltc := newLiveTestClient(t)
	ltc.cleanup(t)

	cl := ltc.cl
	ctx := ltc.ctx

	t.Run("ListReservations", func(t *testing.T) {
		_, err := cl.ListReservations(ctx, ltc.apipid)
		if err != nil {
			t.Errorf("while listing hardware reservations: %v", err)
		}
	})

	var sshKeyID string
	t.Run("CreateSSHKey", func(t *testing.T) {
		nk, err := cl.CreateSSHKey(ctx, &packngo.SSHKeyCreateRequest{
			Label:     ltc.sshKeyLabel,
			Key:       createSSHAuthKey(t),
			ProjectID: ltc.apipid,
		})
		if err != nil {
			t.Fatalf("while creating an SSH key: %v", err)
		}
		if nk.Label != ltc.sshKeyLabel {
			t.Errorf("key labels don't match.")
		}
		t.Logf("Created an SSH key (ID: %s)", nk.ID)
		sshKeyID = nk.ID
	})

	var dummySSHPK2 string
	t.Run("UpdateSSHKey", func(t *testing.T) {
		if sshKeyID == "" {
			t.Skip("SSH key couldn't have been created - skipping...")
		}

		dummySSHPK2 = createSSHAuthKey(t)
		k, err := cl.UpdateSSHKey(ctx, sshKeyID, &packngo.SSHKeyUpdateRequest{
			Key: &dummySSHPK2,
		})
		if err != nil {
			t.Fatalf("while updating an SSH key: %v", err)
		}
		if k.Key != dummySSHPK2 {
			t.Errorf("updated SSH key doesn't match the original.")
		}
	})
	t.Run("GetSSHKey", func(t *testing.T) {
		if sshKeyID == "" {
			t.Skip("SSH key couldn't have been created - skipping...")
		}

		k, err := cl.getSSHKey(ctx, sshKeyID)
		if err != nil {
			t.Fatalf("while getting an SSH key: %v", err)
		}
		if k.Key != dummySSHPK2 {
			t.Errorf("got key contents that don't match the original.")
		}
	})
	t.Run("ListSSHKeys", func(t *testing.T) {
		if sshKeyID == "" {
			t.Skip("SSH key couldn't have been created - skipping...")
		}

		ks, err := cl.ListSSHKeys(ctx)
		if err != nil {
			t.Fatalf("while listing SSH keys: %v", err)
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
	})

	var testDevice *packngo.Device
	t.Run("CreateDevice", func(t *testing.T) {
		// Find a provisionable hardware reservation the device will be created with.
		rvs, err := cl.ListReservations(ctx, ltc.apipid)
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
			Hostname:              ltc.testDeviceHostname,
			OS:                    ltc.apios,
			Plan:                  rv.Plan.Slug,
			HardwareReservationID: rv.ID,
			ProjectID:             ltc.apipid,
		})
		if err != nil {
			t.Fatalf("while creating a device: %v", err)
		}
		t.Logf("Created a new test device (ID: %s)", d.ID)
		testDevice = d
	})
	t.Run("GetDevice", func(t *testing.T) {
		if testDevice == nil {
			t.Skip("the test device couldn't have been created - skipping...")
		}

		d, err := cl.GetDevice(ctx, ltc.apipid, testDevice.ID, nil)
		if err != nil {
			t.Fatalf("while fetching device info: %v", err)
		}
		if d == nil {
			t.Fatalf("expected the test device (ID: %s) to exist.", testDevice.ID)
		}
		if d.ID != testDevice.ID {
			t.Errorf("got device ID that doesn't match the original.")
		}
	})
	t.Run("ListDevices", func(t *testing.T) {
		if testDevice == nil {
			t.Skip("the test device couldn't have been created - skipping...")
		}

		ds, err := cl.ListDevices(ctx, ltc.apipid)
		if err != nil {
			t.Errorf("while listing devices: %v", err)
		}
		if len(ds) == 0 {
			t.Errorf("expected at least one device.")
		}
	})
	t.Run("DeleteDevice", func(t *testing.T) {
		if testDevice == nil {
			t.Skip("the test device couldn't have been created - skipping...")
		}

		// Devices currently being provisioned can't be deleted. After it's
		// provisioned, device's state will match either "active", or "failed".
		if err := ltc.awaitDeviceState(t, testDevice.ID, "active", "failed"); err != nil {
			t.Fatalf("while waiting for device to be provisioned: %v", err)
		}
		t.Logf("Deleting the test device (ID: %s)", testDevice.ID)
		if err := cl.deleteDevice(ctx, testDevice.ID); err != nil {
			t.Fatalf("while deleting a device: %v", err)
		}
		d, err := cl.GetDevice(ctx, ltc.apipid, testDevice.ID, nil)
		if err != nil && !IsNotFound(err) {
			t.Fatalf("while fetching device info: %v", err)
		}
		if d != nil {
			t.Fatalf("device should not exist.")
		}
		t.Logf("Deleted the test device (ID: %s)", testDevice.ID)
	})
	t.Run("DeleteSSHKey", func(t *testing.T) {
		if sshKeyID == "" {
			t.Skip("SSH key couldn't have been created - skipping...")
		}

		t.Logf("Deleting the test SSH key (ID: %s)", sshKeyID)
		if err := cl.deleteSSHKey(ctx, sshKeyID); err != nil {
			t.Fatalf("couldn't delete an SSH key: %v", err)
		}
		_, err := cl.getSSHKey(ctx, sshKeyID)
		if err == nil {
			t.Fatalf("SSH key should not exist")
		}
		t.Logf("Deleted the test SSH key (ID: %s)", sshKeyID)
	})

	ltc.cleanup(t)
}
