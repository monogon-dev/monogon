package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/packethost/packngo"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	ecl "source.monogon.dev/cloud/equinix/wrapngo"
	"source.monogon.dev/cloud/lib/sinbin"
)

type UpdaterConfig struct {
	// Enable starts the updater.
	Enable bool
	// IterationRate is the minimu mtime taken between subsequent iterations of the
	// updater.
	IterationRate time.Duration
}

func (u *UpdaterConfig) RegisterFlags() {
	flag.BoolVar(&u.Enable, "updater_enable", true, "Enable the updater, which periodically scans equinix machines and updates their status in the BMDB")
	flag.DurationVar(&u.IterationRate, "updater_iteration_rate", time.Minute, "Rate limiting for updater iteration loop")
}

// The Updater periodically scans all machines backed by the equinix provider and
// updaters their Provided status fields based on data retrieved from the Equinix
// API.
type Updater struct {
	config *UpdaterConfig
	sinbin sinbin.Sinbin[string]

	cl ecl.Client
}

func (c *UpdaterConfig) New(cl ecl.Client) (*Updater, error) {
	return &Updater{
		config: c,
		cl:     cl,
	}, nil
}

func (u *Updater) Run(ctx context.Context, conn *bmdb.Connection) error {
	var sess *bmdb.Session
	var err error

	if !u.config.Enable {
		return nil
	}

	for {
		if sess == nil {
			sess, err = conn.StartSession(ctx, bmdb.SessionOption{Processor: metrics.ProcessorShepherdUpdater})
			if err != nil {
				return fmt.Errorf("could not start BMDB session: %w", err)
			}
		}
		limit := time.After(u.config.IterationRate)

		err = u.runInSession(ctx, sess)
		switch {
		case err == nil:
			<-limit
		case errors.Is(err, ctx.Err()):
			return err
		case errors.Is(err, bmdb.ErrSessionExpired):
			klog.Errorf("Session expired, restarting...")
			sess = nil
			time.Sleep(time.Second)
		case err != nil:
			klog.Errorf("Processing failed: %v", err)
			// TODO(q3k): close session
			time.Sleep(time.Second)
		}
	}
}

// applyNullStringUpdate returns true if 'up' supersedes 'cur'. Otherwise, it
// returns false and zeroes out up.
func applyNullStringUpdate(up, cur *sql.NullString) bool {
	if up.Valid {
		if !cur.Valid {
			return true
		}
		if up.String != cur.String {
			return true
		}
	}
	up.String = ""
	up.Valid = false
	return false
}

// applyNullProviderStatusUpdate returns true if 'up' supersedes 'cur'.
// Otherwise, it returns false and zeroes out up.
func applyNullProviderStatusUpdate(up, cur *model.NullProviderStatus) bool {
	if up.Valid {
		if !cur.Valid {
			return true
		}
		if up.ProviderStatus != cur.ProviderStatus {
			return true
		}
	}
	up.ProviderStatus = model.ProviderStatusUnknown
	up.Valid = false
	return false
}

// applyUpdate returns true if 'up' supersedes 'cur'. Otherwise, it returns false
// and zeroes out up.
func applyUpdate(up *model.MachineUpdateProviderStatusParams, cur *model.MachineProvided) bool {
	res := false
	res = res || applyNullStringUpdate(&up.ProviderReservationID, &cur.ProviderReservationID)
	res = res || applyNullStringUpdate(&up.ProviderIpAddress, &cur.ProviderIpAddress)
	res = res || applyNullStringUpdate(&up.ProviderLocation, &cur.ProviderLocation)
	res = res || applyNullProviderStatusUpdate(&up.ProviderStatus, &cur.ProviderStatus)
	return res
}

// updateLog logs information about the given update as calculated by applyUpdate.
func updateLog(up *model.MachineUpdateProviderStatusParams) {
	if up.ProviderReservationID.Valid {
		klog.Infof("   Machine %s: new reservation ID %s", up.ProviderID, up.ProviderReservationID.String)
	}
	if up.ProviderIpAddress.Valid {
		klog.Infof("   Machine %s: new IP address %s", up.ProviderID, up.ProviderIpAddress.String)
	}
	if up.ProviderLocation.Valid {
		klog.Infof("   Machine %s: new location %s", up.ProviderID, up.ProviderLocation.String)
	}
	if up.ProviderStatus.Valid {
		klog.Infof("   Machine %s: new status %s", up.ProviderID, up.ProviderStatus.ProviderStatus)
	}
}

func (u *Updater) runInSession(ctx context.Context, sess *bmdb.Session) error {
	// Get all machines provided by us into the BMDB.
	// TODO(q3k): do not load all machines into memory.

	var machines []model.MachineProvided
	err := sess.Transact(ctx, func(q *model.Queries) error {
		var err error
		machines, err = q.GetProvidedMachines(ctx, model.ProviderEquinix)
		return err
	})
	if err != nil {
		return fmt.Errorf("when fetching provided machines: %w", err)
	}

	// Limit how many machines we check by timing them out if they're likely to not
	// get updated soon.
	penalized := 0
	var check []model.MachineProvided
	for _, m := range machines {
		if u.sinbin.Penalized(m.ProviderID) {
			penalized += 1
		} else {
			check = append(check, m)
		}
	}

	klog.Infof("Machines to check %d, skipping: %d", len(check), penalized)
	for _, m := range check {
		dev, err := u.cl.GetDevice(ctx, "", m.ProviderID, &packngo.ListOptions{
			Includes: []string{
				"hardware_reservation",
			},
			Excludes: []string{
				"created_by", "customdata", "network_ports", "operating_system", "actions",
				"plan", "provisioning_events", "ssh_keys", "tags", "volumes",
			},
		})
		if err != nil {
			klog.Warningf("Fetching device %s failed: %v", m.ProviderID, err)
			continue
		}

		// nextCheck will be used to sinbin the machine for some given time if there is
		// no difference between the current state and new state.
		//
		// Some conditions override this to be shorter (when the machine doesn't yet have
		// all data available or is in an otherwise unstable state).
		nextCheck := time.Minute * 30

		up := model.MachineUpdateProviderStatusParams{
			Provider:   m.Provider,
			ProviderID: m.ProviderID,
		}

		if dev.HardwareReservation != nil {
			up.ProviderReservationID.Valid = true
			up.ProviderReservationID.String = dev.HardwareReservation.ID
		} else {
			nextCheck = time.Minute
		}

		for _, addr := range dev.Network {
			if !addr.Public {
				continue
			}
			up.ProviderIpAddress.Valid = true
			up.ProviderIpAddress.String = addr.Address
			break
		}
		if !up.ProviderIpAddress.Valid {
			nextCheck = time.Minute
		}

		if dev.Facility != nil {
			up.ProviderLocation.Valid = true
			up.ProviderLocation.String = dev.Facility.Code
		} else {
			nextCheck = time.Minute
		}

		up.ProviderStatus.Valid = true
		switch dev.State {
		case "active":
			up.ProviderStatus.ProviderStatus = model.ProviderStatusRunning
		case "deleted":
			up.ProviderStatus.ProviderStatus = model.ProviderStatusMissing
		case "failed":
			up.ProviderStatus.ProviderStatus = model.ProviderStatusProvisioningFailedPermanent
		case "inactive":
			up.ProviderStatus.ProviderStatus = model.ProviderStatusStopped
		case "powering_on", "powering_off":
			nextCheck = time.Minute
			up.ProviderStatus.ProviderStatus = model.ProviderStatusStopped
		case "queued", "provisioning", "reinstalling", "post_provisioning":
			nextCheck = time.Minute
			up.ProviderStatus.ProviderStatus = model.ProviderStatusProvisioning
		default:
			klog.Warningf("Device %s has unexpected status: %q", m.ProviderID, dev.State)
			nextCheck = time.Minute
			up.ProviderStatus.ProviderStatus = model.ProviderStatusUnknown
		}

		if !applyUpdate(&up, &m) {
			u.sinbin.Add(m.ProviderID, time.Now().Add(nextCheck))
			continue
		}

		klog.Infof("Device %s has new data:", m.ProviderID)
		updateLog(&up)
		err = sess.Transact(ctx, func(q *model.Queries) error {
			return q.MachineUpdateProviderStatus(ctx, up)
		})
		if err != nil {
			klog.Warningf("Device %s failed to update: %v", m.ProviderID, err)
		}
		u.sinbin.Add(m.ProviderID, time.Now().Add(time.Minute))
	}
	return nil
}
