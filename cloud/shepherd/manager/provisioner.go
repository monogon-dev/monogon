package manager

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/netip"
	"sort"
	"time"

	"github.com/google/uuid"
	"golang.org/x/time/rate"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/shepherd"
	"source.monogon.dev/go/mflags"
)

// Provisioner implements the server provisioning logic. Provisioning entails
// bringing all available machines (subject to limits) into BMDB.
type Provisioner struct {
	ProvisionerConfig
	p shepherd.Provider
}

// ProvisionerConfig configures the provisioning process.
type ProvisionerConfig struct {
	// MaxCount is the maximum count of managed servers. No new devices will be
	// created after reaching the limit. No attempt will be made to reduce the
	// server count.
	MaxCount uint

	// ReconcileLoopLimiter limits the rate of the main reconciliation loop
	// iterating.
	ReconcileLoopLimiter *rate.Limiter

	// DeviceCreation limits the rate at which devices are created.
	DeviceCreationLimiter *rate.Limiter

	// ChunkSize is how many machines will try to be spawned in a
	// single reconciliation loop. Higher numbers allow for faster initial
	// provisioning, but lower numbers decrease potential raciness with other systems
	// and make sure that other parts of the reconciliation logic are ran regularly.
	//
	// 20 is decent starting point.
	ChunkSize uint
}

func (pc *ProvisionerConfig) RegisterFlags() {
	flag.UintVar(&pc.MaxCount, "provisioner_max_machines", 50, "Limit of machines that the provisioner will attempt to pull into the BMDB. Zero for no limit.")
	mflags.Limiter(&pc.ReconcileLoopLimiter, "provisioner_reconciler_rate", "1m,1", "Rate limiting for main provisioner reconciliation loop")
	mflags.Limiter(&pc.DeviceCreationLimiter, "provisioner_device_creation_rate", "5s,1", "Rate limiting for machine creation")
	flag.UintVar(&pc.ChunkSize, "provisioner_reservation_chunk_size", 20, "How many machines will the provisioner attempt to create in a single reconciliation loop iteration")
}

func (pc *ProvisionerConfig) check() error {
	// If these are unset, it's probably because someone is using us as a library.
	// Provide error messages useful to code users instead of flag names.
	if pc.ReconcileLoopLimiter == nil {
		return fmt.Errorf("ReconcileLoopLimiter must be set")
	}
	if pc.DeviceCreationLimiter == nil {
		return fmt.Errorf("DeviceCreationLimiter must be set")
	}
	if pc.ChunkSize == 0 {
		return fmt.Errorf("ChunkSize must be set")
	}
	return nil
}

// NewProvisioner creates a Provisioner instance, checking ProvisionerConfig and
// providerConfig for errors.
func NewProvisioner(p shepherd.Provider, pc ProvisionerConfig) (*Provisioner, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}

	return &Provisioner{
		ProvisionerConfig: pc,
		p:                 p,
	}, nil
}

// Run the provisioner blocking the current goroutine until the given context
// expires.
func (p *Provisioner) Run(ctx context.Context, conn *bmdb.Connection) error {

	var sess *bmdb.Session
	var err error
	for {
		if sess == nil {
			sess, err = conn.StartSession(ctx, bmdb.SessionOption{Processor: metrics.ProcessorShepherdProvisioner})
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
	machines []shepherd.Machine
	err      error
}

// runInSession executes one iteration of the provisioner's control loop within a
// BMDB session. This control loop attempts to bring all capacity into machines in
// the BMDB, subject to limits.
func (p *Provisioner) runInSession(ctx context.Context, sess *bmdb.Session) error {
	if err := p.ReconcileLoopLimiter.Wait(ctx); err != nil {
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
			klog.Infof("Got %d machines in provider.", len(inProvider.machines))
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
func (p *Provisioner) listInProvider(ctx context.Context) ([]shepherd.Machine, error) {
	machines, err := p.p.ListMachines(ctx)
	if err != nil {
		return nil, fmt.Errorf("while fetching managed machines: %w", err)
	}
	sort.Slice(machines, func(i, j int) bool {
		return machines[i].ID() < machines[j].ID()
	})
	return machines, nil
}

type providedMachine struct {
	model.MachineProvided
}

func (p providedMachine) Failed() bool {
	if !p.MachineProvided.ProviderStatus.Valid {
		// If we don't have any ProviderStatus to check for, return false
		// to trigger the validation inside the reconciler loop.
		return false
	}
	switch p.MachineProvided.ProviderStatus.ProviderStatus {
	case model.ProviderStatusProvisioningFailedPermanent:
		return true
	}
	return false
}

func (p providedMachine) ID() shepherd.ProviderID {
	return shepherd.ProviderID(p.ProviderID)
}

func (p providedMachine) Addr() netip.Addr {
	if !p.ProviderIpAddress.Valid {
		return netip.Addr{}
	}

	addr, err := netip.ParseAddr(p.ProviderIpAddress.String)
	if err != nil {
		return netip.Addr{}
	}
	return addr
}

func (p providedMachine) State() shepherd.State {
	return shepherd.StateKnownUsed
}

// listInBMDB returns all the machines that the BMDB thinks we should be managing.
func (p *Provisioner) listInBMDB(ctx context.Context, sess *bmdb.Session) ([]shepherd.Machine, error) {
	var res []shepherd.Machine
	err := sess.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetProvidedMachines(ctx, p.p.Type())
		if err != nil {
			return err
		}
		res = make([]shepherd.Machine, 0, len(machines))
		for _, machine := range machines {
			_, err := uuid.Parse(machine.ProviderID)
			if err != nil {
				klog.Errorf("BMDB machine %s has unparseable provider ID %q", machine.MachineID, machine.ProviderID)
				continue
			}

			res = append(res, providedMachine{machine})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID() < res[j].ID()
	})
	return res, nil
}

// resolvePossiblyUsed checks if the state is set to possibly used and finds out
// which state is the correct one.
func (p *Provisioner) resolvePossiblyUsed(machine shepherd.Machine, providedMachines map[shepherd.ProviderID]shepherd.Machine) shepherd.State {
	state, id := machine.State(), machine.ID()

	// Bail out if this isn't a possibly used state.
	if state != shepherd.StatePossiblyUsed {
		return state
	}

	// If a machine does not have a valid id, its always seen as unused.
	if !id.IsValid() {
		return shepherd.StateKnownUnused
	}

	// If the machine is not inside the bmdb, it's seen as unused.
	if _, ok := providedMachines[id]; !ok {
		return shepherd.StateKnownUnused
	}

	return shepherd.StateKnownUsed
}

// reconcile takes a list of machines that the provider thinks we should be
// managing and that the BMDB thinks we should be managing, and tries to make
// sense of that. First, some checks are performed across the two lists to make
// sure we haven't dropped anything. Then, additional machines are deployed from
// hardware reservations as needed.
func (p *Provisioner) reconcile(ctx context.Context, sess *bmdb.Session, inProvider, bmdbMachines []shepherd.Machine) error {
	klog.Infof("Reconciling...")

	bmdb := make(map[shepherd.ProviderID]shepherd.Machine)
	for _, machine := range bmdbMachines {
		// Dont check the state here as its hardcoded to be known used.
		bmdb[machine.ID()] = machine
	}

	var availableMachines []shepherd.Machine
	provider := make(map[shepherd.ProviderID]shepherd.Machine)
	for _, machine := range inProvider {
		state := p.resolvePossiblyUsed(machine, bmdb)

		switch state {
		case shepherd.StateKnownUnused:
			availableMachines = append(availableMachines, machine)

		case shepherd.StateKnownUsed:
			provider[machine.ID()] = machine

		default:
			return fmt.Errorf("machine has invalid state (ID: %s, Addr: %s): %s", machine.ID(), machine.Addr(), state)
		}
	}

	managed := make(map[shepherd.ProviderID]bool)

	// We discovered that a machine mostly fails either when provisioning or
	// deprovisioning. A already deployed and running machine can only switch
	// into failed state if any api interaction happend, e.g. rebooting the
	// machine into recovery mode. If such a machine is returned to the
	// reconciling loop, it will trigger the badbadnotgood safety switch and
	// return with an error. To reduce the manual intervention required we
	// filter out these machines on both sides (bmdb and provider).
	isBadBadNotGood := func(known map[shepherd.ProviderID]shepherd.Machine, machine shepherd.Machine) bool {
		// If the machine is missing and not failed, its a bad case.
		if known[machine.ID()] == nil && !machine.Failed() {
			return true
		}
		return false
	}

	// Some desynchronization between the BMDB and Provider point of view might be so
	// bad we shouldn't attempt to do any work, at least not any time soon.
	badbadnotgood := false

	// Find any machines supposedly managed by us in the provider, but not in the
	// BMDB.
	for id, machine := range provider {
		if isBadBadNotGood(bmdb, machine) {
			klog.Errorf("Provider machine has no corresponding machine in BMDB. (PID: %s)", id)
			badbadnotgood = true
			continue
		}

		managed[id] = true
	}

	// Find any machines in the BMDB but not in the provider.
	for id, machine := range bmdb {
		if isBadBadNotGood(provider, machine) {
			klog.Errorf("Provider machine referred to in BMDB but missing in provider. (PID: %s)", id)
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

	if p.MaxCount != 0 && p.MaxCount <= uint(nmanaged) {
		klog.Infof("Not bringing up more machines (at limit of %d machines)", p.MaxCount)
		return nil
	}

	limitName := "no limit"
	if p.MaxCount != 0 {
		limitName = fmt.Sprintf("%d", p.MaxCount)
	}
	klog.Infof("Below managed machine limit (%s), bringing up more...", limitName)

	if len(availableMachines) == 0 {
		klog.Infof("No more capacity available.")
		return nil
	}

	toProvision := availableMachines
	// Limit them to MaxCount, if applicable.
	if p.MaxCount != 0 {
		needed := int(p.MaxCount) - nmanaged
		if len(toProvision) < needed {
			needed = len(toProvision)
		}
		toProvision = toProvision[:needed]
	}

	// Limit them to an arbitrary 'chunk' size so that we don't do too many things in
	// a single reconciliation operation.
	if uint(len(toProvision)) > p.ChunkSize {
		toProvision = toProvision[:p.ChunkSize]
	}

	if len(toProvision) == 0 {
		klog.Infof("No more unused machines available, or all filtered out.")
		return nil
	}

	klog.Infof("Bringing up %d machines...", len(toProvision))
	for _, machine := range toProvision {
		if err := p.DeviceCreationLimiter.Wait(ctx); err != nil {
			return err
		}

		nd, err := p.p.CreateMachine(ctx, sess, shepherd.CreateMachineRequest{
			UnusedMachine: machine,
		})
		if err != nil {
			klog.Errorf("while creating new device (ID: %s, Addr: %s, State: %s): %w", machine.ID(), machine.Addr(), machine.State(), err)
			continue
		}
		klog.Infof("Created new machine with ID: %s", nd.ID())
	}

	return nil
}
