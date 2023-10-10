package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/netip"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/shepherd"
)

// provider represents a shepherd.Provider that works entirely on a
// static device list. It requires a provider type and a device list.
type provider struct {
	providerType model.Provider
	machines     map[shepherd.ProviderID]machine
}

type machine struct {
	ProviderID shepherd.ProviderID `json:"ID"`
	Address    netip.Addr          `json:"Addr"`
	Location   string              `json:"Location"`
}

func (d machine) ID() shepherd.ProviderID {
	return d.ProviderID
}

func (d machine) Addr() netip.Addr {
	return d.Address
}

func (d machine) State() shepherd.State {
	return shepherd.StatePossiblyUsed
}

func (p *provider) ListMachines(ctx context.Context) ([]shepherd.Machine, error) {
	machines := make([]shepherd.Machine, 0, len(p.machines))
	for _, m := range p.machines {
		machines = append(machines, m)
	}

	return machines, nil
}

func (p *provider) GetMachine(ctx context.Context, id shepherd.ProviderID) (shepherd.Machine, error) {
	// If the provided machine is not inside our known machines,
	// bail-out early as this is unsupported.
	if _, ok := p.machines[id]; !ok {
		return nil, fmt.Errorf("unknown provided machine requested")
	}

	return p.machines[id], nil
}

func (p *provider) CreateMachine(ctx context.Context, session *bmdb.Session, request shepherd.CreateMachineRequest) (shepherd.Machine, error) {
	if request.UnusedMachine == nil {
		return nil, fmt.Errorf("parameter UnusedMachine is missing")
	}

	//TODO: Do we just trust the implementation to be correct?
	m, ok := request.UnusedMachine.(machine)
	if !ok {
		return nil, fmt.Errorf("invalid type for parameter UnusedMachine")
	}

	if err := p.assimilate(ctx, session, m); err != nil {
		klog.Errorf("Failed to provision machine %s: %v", m.ProviderID, err)
		return nil, err
	}

	return m, nil
}

func (p *provider) assimilate(ctx context.Context, sess *bmdb.Session, machine machine) error {
	return sess.Transact(ctx, func(q *model.Queries) error {
		// Create a new machine record within BMDB.
		m, err := q.NewMachine(ctx)
		if err != nil {
			return fmt.Errorf("while creating a new machine record in BMDB: %w", err)
		}

		// Link the new machine with the device, and tag it "provided".
		addParams := model.MachineAddProvidedParams{
			MachineID:  m.MachineID,
			ProviderID: string(machine.ProviderID),
			Provider:   p.providerType,
		}
		klog.Infof("Setting \"provided\" tag (ID: %s, PID: %s, Provider: %s).", addParams.MachineID, addParams.ProviderID, addParams.Provider)
		if err := q.MachineAddProvided(ctx, addParams); err != nil {
			return fmt.Errorf("while tagging machine active: %w", err)
		}

		upParams := model.MachineUpdateProviderStatusParams{
			ProviderID: string(machine.ProviderID),
			Provider:   p.providerType,
			ProviderIpAddress: sql.NullString{
				String: machine.Address.String(),
				Valid:  true,
			},
			ProviderLocation: sql.NullString{
				String: machine.Location,
				Valid:  machine.Location != "",
			},
			ProviderStatus: model.NullProviderStatus{
				ProviderStatus: model.ProviderStatusUnknown,
				Valid:          true,
			},
		}

		klog.Infof("Setting \"provided\" tag status parameter (ID: %s, PID: %s, Provider: %s).", addParams.MachineID, upParams.ProviderID, upParams.Provider)
		if err := q.MachineUpdateProviderStatus(ctx, upParams); err != nil {
			return fmt.Errorf("while setting machine params: %w", err)
		}

		return nil
	})
}

func (p *provider) Type() model.Provider {
	return p.providerType
}
