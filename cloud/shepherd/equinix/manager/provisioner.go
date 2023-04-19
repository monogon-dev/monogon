package manager

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/packethost/packngo"
	"golang.org/x/time/rate"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/sinbin"
	ecl "source.monogon.dev/cloud/shepherd/equinix/wrapngo"
)

// ProvisionerConfig configures the provisioning process.
type ProvisionerConfig struct {
	// OS defines the operating system new devices are created with. Its format
	// is specified by Equinix API.
	OS string
	// MaxCount is the maximum count of managed servers. No new devices will be
	// created after reaching the limit. No attempt will be made to reduce the
	// server count.
	MaxCount uint

	// ReconcileLoopLimiter limits the rate of the main reconciliation loop
	// iterating. As new machines are being provisioned, each loop will cause one
	// 'long' ListHardwareReservations call to Equinix.
	ReconcileLoopLimiter *rate.Limiter

	// DeviceCreation limits the rate at which devices are created within
	// Equinix through use of appropriate API calls.
	DeviceCreationLimiter *rate.Limiter

	// Assimilate Equinix machines that match the configured device prefix into the
	// BMDB as Provided. This should only be used for manual testing with
	// -bmdb_eat_my_data.
	Assimilate bool

	// ReservationChunkSize is how many Equinix machines will try to be spawned in a
	// single reconciliation loop. Higher numbers allow for faster initial
	// provisioning, but lower numbers decrease potential raciness with other systems
	// and make sure that other parts of the reconciliation logic are ran regularly.
	//
	// 20 is decent starting point.
	ReservationChunkSize uint

	// UseProjectKeys defines if the provisioner adds all ssh keys defined inside
	// the used project to every new machine. This is only used for debug purposes.
	UseProjectKeys bool
}

func (p *ProvisionerConfig) RegisterFlags() {
	flag.StringVar(&p.OS, "provisioner_os", "ubuntu_20_04", "OS that provisioner will deploy on Equinix machines. Not the target OS for cluster customers.")
	flag.UintVar(&p.MaxCount, "provisioner_max_machines", 50, "Limit of machines that the provisioner will attempt to pull into the BMDB. Zero for no limit.")
	flagLimiter(&p.ReconcileLoopLimiter, "provisioner_reconciler_rate", "1m,1", "Rate limiting for main provisioner reconciliation loop")
	flagLimiter(&p.DeviceCreationLimiter, "provisioner_device_creation_rate", "5s,1", "Rate limiting for Equinix device/machine creation")
	flag.BoolVar(&p.Assimilate, "provisioner_assimilate", false, "Assimilate matching machines in Equinix project into BMDB as Provided. Only to be used when manually testing.")
	flag.UintVar(&p.ReservationChunkSize, "provisioner_reservation_chunk_size", 20, "How many machines will the provisioner attempt to create in a single reconciliation loop iteration")
	flag.BoolVar(&p.UseProjectKeys, "provisioner_use_project_keys", false, "Add all Equinix project keys to newly provisioned machines, not just the provisioner's managed key. Debug/development only.")
}

// Provisioner implements the server provisioning logic. Provisioning entails
// bringing all available hardware reservations (subject to limits) into BMDB as
// machines provided by Equinix.
type Provisioner struct {
	config       *ProvisionerConfig
	sharedConfig *SharedConfig

	// cl is the wrapngo client instance used.
	cl ecl.Client

	// badReservations is a holiday resort for Equinix hardware reservations which
	// failed to be provisioned for some reason or another. We keep a list of them in
	// memory just so that we don't repeatedly try to provision the same known bad
	// machines.
	badReservations sinbin.Sinbin[string]
}

// New creates a Provisioner instance, checking ProvisionerConfig and
// SharedConfig for errors.
func (c *ProvisionerConfig) New(cl ecl.Client, sc *SharedConfig) (*Provisioner, error) {
	// If these are unset, it's probably because someone is using us as a library.
	// Provide error messages useful to code users instead of flag names.
	if c.OS == "" {
		return nil, fmt.Errorf("OS must be set")
	}
	if c.ReconcileLoopLimiter == nil {
		return nil, fmt.Errorf("ReconcileLoopLimiter must be set")
	}
	if c.DeviceCreationLimiter == nil {
		return nil, fmt.Errorf("DeviceCreationLimiter must be set")
	}
	if c.ReservationChunkSize == 0 {
		return nil, fmt.Errorf("ReservationChunkSize must be set")
	}
	return &Provisioner{
		config:       c,
		sharedConfig: sc,

		cl: cl,
	}, nil
}

// Run the provisioner blocking the current goroutine until the given context
// expires.
func (p *Provisioner) Run(ctx context.Context, conn *bmdb.Connection) error {

	var sess *bmdb.Session
	var err error
	for {
		if sess == nil {
			sess, err = conn.StartSession(ctx)
			if err != nil {
				return fmt.Errorf("could not start BMDB session: %w", err)
			}
		}
		err = p.runInSession(ctx, sess)

		switch {
		case err == nil:
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

type machineListing struct {
	machines []uuid.UUID
	err      error
}

// runInSession executes one iteration of the provisioner's control loop within a
// BMDB session. This control loop attempts to bring all Equinix hardware
// reservations into machines in the BMDB, subject to limits.
func (p *Provisioner) runInSession(ctx context.Context, sess *bmdb.Session) error {
	if err := p.config.ReconcileLoopLimiter.Wait(ctx); err != nil {
		return err
	}

	providerC := make(chan *machineListing, 1)
	bmdbC := make(chan *machineListing, 1)

	klog.Infof("Getting provider and bmdb machines...")

	// Make sub-context for two parallel operations, and so that we can cancel one
	// immediately if the other fails.
	subCtx, subCtxC := context.WithCancel(ctx)
	defer subCtxC()

	go func() {
		machines, err := p.listInProvider(subCtx)
		providerC <- &machineListing{
			machines: machines,
			err:      err,
		}
	}()
	go func() {
		machines, err := p.listInBMDB(subCtx, sess)
		bmdbC <- &machineListing{
			machines: machines,
			err:      err,
		}
	}()
	var inProvider, inBMDB *machineListing
	for {
		select {
		case inProvider = <-providerC:
			if err := inProvider.err; err != nil {
				return fmt.Errorf("listing provider machines failed: %w", err)
			}
			klog.Infof("Got %d machines managed in provider.", len(inProvider.machines))
		case inBMDB = <-bmdbC:
			if err := inBMDB.err; err != nil {
				return fmt.Errorf("listing BMDB machines failed: %w", err)
			}
			klog.Infof("Got %d machines in BMDB.", len(inBMDB.machines))
		}
		if inProvider != nil && inBMDB != nil {
			break
		}
	}

	subCtxC()
	if err := p.reconcile(ctx, sess, inProvider.machines, inBMDB.machines); err != nil {
		return fmt.Errorf("reconciliation failed: %w", err)
	}
	return nil
}

// listInProviders returns all machines that the provider thinks we should be
// managing.
func (p *Provisioner) listInProvider(ctx context.Context) ([]uuid.UUID, error) {
	devices, err := p.sharedConfig.managedDevices(ctx, p.cl)
	if err != nil {
		return nil, fmt.Errorf("while fetching managed machines: %w", err)
	}
	var pvr []uuid.UUID
	for _, dev := range devices {
		id, err := uuid.Parse(dev.ID)
		if err != nil {
			klog.Errorf("Device ID %q is not UUID, skipping", dev.ID)
		} else {
			pvr = append(pvr, id)
		}
	}
	sort.Slice(pvr, func(i, j int) bool {
		return pvr[i].String() < pvr[j].String()
	})
	return pvr, nil
}

// listInBMDB returns all the machines that the BMDB thinks we should be managing.
func (p *Provisioner) listInBMDB(ctx context.Context, sess *bmdb.Session) ([]uuid.UUID, error) {
	var res []uuid.UUID
	err := sess.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetProvidedMachines(ctx, model.ProviderEquinix)
		if err != nil {
			return err
		}
		res = make([]uuid.UUID, len(machines))
		for i, machine := range machines {
			id, err := uuid.Parse(machine.ProviderID)
			if err != nil {
				klog.Errorf("BMDB machine %s has unparseable provider ID %q", machine.MachineID, machine.ProviderID)
			} else {
				res[i] = id
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].String() < res[j].String()
	})
	return res, nil
}

// reconcile takes a list of machines that the provider thinks we should be
// managing and that the BMDB thinks we should be managing, and tries to make
// sense of that. First, some checks are performed across the two lists to make
// sure we haven't dropped anything. Then, additional machines are deployed from
// hardware reservations as needed.
func (p *Provisioner) reconcile(ctx context.Context, sess *bmdb.Session, inProvider, inBMDB []uuid.UUID) error {
	klog.Infof("Reconciling...")

	bmdb := make(map[string]bool)
	provider := make(map[string]bool)
	for _, machine := range inProvider {
		provider[machine.String()] = true
	}
	for _, machine := range inBMDB {
		bmdb[machine.String()] = true
	}

	managed := make(map[string]bool)

	// Some desynchronization between the BMDB and Provider point of view might be so
	// bad we shouldn't attempt to do any work, at least not any time soon.
	badbadnotgood := false

	// Find any machines supposedly managed by us in the provider, but not in the
	// BMDB, and assimilate them if so configured.
	for machine, _ := range provider {
		if bmdb[machine] {
			managed[machine] = true
			continue
		}
		if p.config.Assimilate {
			klog.Warningf("Provider machine %s has no corresponding machine in BMDB. Assimilating it.", machine)
			if err := p.assimilate(ctx, sess, machine); err != nil {
				klog.Errorf("Failed to assimilate: %v", err)
			} else {
				managed[machine] = true
			}
		} else {
			klog.Errorf("Provider machine %s has no corresponding machine in BMDB.", machine)
			badbadnotgood = true
		}
	}

	// Find any machines in the BMDB but not in the provider.
	for machine, _ := range bmdb {
		if !provider[machine] {
			klog.Errorf("Provider device ID %s referred to in BMDB (from TODO) but missing in provider.", machine)
			badbadnotgood = true
		}
	}

	// Bail if things are weird.
	if badbadnotgood {
		klog.Errorf("Something's very wrong. Bailing early and refusing to do any work.")
		return fmt.Errorf("fatal discrepency between BMDB and provider")
	}

	// Summarize all managed machines, which is the intersection of BMDB and
	// Provisioner machines, usually both of these sets being equal.
	nmanaged := len(managed)
	klog.Infof("Total managed machines: %d", nmanaged)

	if p.config.MaxCount != 0 && p.config.MaxCount <= uint(nmanaged) {
		klog.Infof("Not bringing up more machines (at limit of %d machines)", p.config.MaxCount)
		return nil
	}

	limitName := "no limit"
	if p.config.MaxCount != 0 {
		limitName = fmt.Sprintf("%d", p.config.MaxCount)
	}
	klog.Infof("Below managed machine limit (%s), bringing up more...", limitName)
	klog.Infof("Retrieving hardware reservations, this will take a while...")
	reservations, err := p.cl.ListReservations(ctx, p.sharedConfig.ProjectId)
	if err != nil {
		return fmt.Errorf("failed to list reservations: %w", err)
	}

	// Collect all reservations.
	var toProvision []packngo.HardwareReservation
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
		if p.badReservations.Penalized(reservation.ID) {
			penalized++
			continue
		}
		toProvision = append(toProvision, reservation)
	}
	klog.Infof("Retrieved hardware reservations: %d (total), %d (available), %d (in use), %d (not provisionable), %d (penalized)", len(reservations), len(toProvision), inUse, notProvisionable, penalized)

	// Limit them to MaxCount, if applicable.
	if p.config.MaxCount != 0 {
		needed := int(p.config.MaxCount) - nmanaged
		if len(toProvision) < needed {
			needed = len(toProvision)
		}
		toProvision = toProvision[:needed]
	}

	// Limit them to an arbitrary 'chunk' size so that we don't do too many things in
	// a single reconciliation operation.
	if uint(len(toProvision)) > p.config.ReservationChunkSize {
		toProvision = toProvision[:p.config.ReservationChunkSize]
	}

	if len(toProvision) == 0 {
		klog.Infof("No more hardware reservations available, or all filtered out.")
		return nil
	}

	klog.Infof("Bringing up %d machines...", len(toProvision))
	for _, res := range toProvision {
		p.config.DeviceCreationLimiter.Wait(ctx)
		if err := p.provision(ctx, sess, res); err != nil {
			klog.Errorf("Failed to provision reservation %s: %v", res.ID, err)
			until := time.Now().Add(time.Hour)
			klog.Errorf("Adding hardware reservation %s to sinbin until %s", res.ID, until)
			p.badReservations.Add(res.ID, until)
		}
	}

	return nil
}

// provision attempts to create a device within Equinix using given Hardware
// Reservation rsv. The resulting device is registered with BMDB, and tagged as
// "provided" in the process.
func (pr *Provisioner) provision(ctx context.Context, sess *bmdb.Session, rsv packngo.HardwareReservation) error {
	klog.Infof("Creating a new device using reservation ID %s.", rsv.ID)
	hostname := pr.sharedConfig.DevicePrefix + rsv.ID[:18]
	kid, err := pr.sharedConfig.sshEquinixId(ctx, pr.cl)
	if err != nil {
		return err
	}
	req := &packngo.DeviceCreateRequest{
		Hostname:              hostname,
		OS:                    pr.config.OS,
		Plan:                  rsv.Plan.Slug,
		ProjectID:             pr.sharedConfig.ProjectId,
		HardwareReservationID: rsv.ID,
		ProjectSSHKeys:        []string{kid},
	}
	if pr.config.UseProjectKeys {
		klog.Warningf("INSECURE: Machines will be created with ALL PROJECT SSH KEYS!")
		req.ProjectSSHKeys = nil
	}

	nd, err := pr.cl.CreateDevice(ctx, req)
	if err != nil {
		return fmt.Errorf("while creating new device within Equinix: %w", err)
	}
	klog.Infof("Created a new device within Equinix (PID: %s).", nd.ID)

	err = pr.assimilate(ctx, sess, nd.ID)
	if err != nil {
		// TODO(mateusz@monogon.tech) at this point the device at Equinix isn't
		// matched by a BMDB record. Schedule device deletion or make sure this
		// case is being handled elsewhere.
		return err
	}
	return nil
}

// assimilate brings in an already existing machine from Equinix into the BMDB.
// This is only used in manual testing.
func (pr *Provisioner) assimilate(ctx context.Context, sess *bmdb.Session, deviceID string) error {
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
