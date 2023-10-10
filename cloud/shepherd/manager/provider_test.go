package manager

import (
	"context"
	"fmt"
	"net/netip"
	"time"

	"github.com/google/uuid"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/shepherd"
)

type dummyMachine struct {
	id           shepherd.ProviderID
	addr         netip.Addr
	state        shepherd.State
	agentStarted bool
}

func (dm *dummyMachine) ID() shepherd.ProviderID {
	return dm.id
}

func (dm *dummyMachine) Addr() netip.Addr {
	return dm.addr
}

func (dm *dummyMachine) State() shepherd.State {
	return dm.state
}

type dummySSHClient struct {
	SSHClient
	dp *dummyProvider
}

type dummySSHConnection struct {
	SSHConnection
	m *dummyMachine
}

func (dsc *dummySSHConnection) Execute(ctx context.Context, command string, stdin []byte) ([]byte, []byte, error) {
	stdout, stderr, err := dsc.SSHConnection.Execute(ctx, command, stdin)
	if err != nil {
		return nil, nil, err
	}

	dsc.m.agentStarted = true
	return stdout, stderr, nil
}

func (dsc *dummySSHClient) Dial(ctx context.Context, address string, timeout time.Duration) (SSHConnection, error) {
	conn, err := dsc.SSHClient.Dial(ctx, address, timeout)
	if err != nil {
		return nil, err
	}

	addrPort := netip.MustParseAddrPort(address)
	uid, err := uuid.FromBytes(addrPort.Addr().AsSlice())
	if err != nil {
		return nil, err
	}

	m := dsc.dp.machines[shepherd.ProviderID(uid.String())]
	if m == nil {
		return nil, fmt.Errorf("failed finding machine in map")
	}

	return &dummySSHConnection{conn, m}, nil
}

func (dp *dummyProvider) sshClient() SSHClient {
	return &dummySSHClient{
		SSHClient: &FakeSSHClient{},
		dp:        dp,
	}
}

func newDummyProvider(cap int) *dummyProvider {
	return &dummyProvider{
		capacity: cap,
		machines: make(map[shepherd.ProviderID]*dummyMachine),
	}
}

type dummyProvider struct {
	capacity int
	machines map[shepherd.ProviderID]*dummyMachine
}

func (dp *dummyProvider) createDummyMachines(ctx context.Context, session *bmdb.Session, count int) ([]shepherd.Machine, error) {
	if len(dp.machines)+count > dp.capacity {
		return nil, fmt.Errorf("no capacity left")
	}

	var machines []shepherd.Machine
	for i := 0; i < count; i++ {
		uid := uuid.Must(uuid.NewRandom())
		m, err := dp.CreateMachine(ctx, session, shepherd.CreateMachineRequest{
			UnusedMachine: &dummyMachine{
				id:    shepherd.ProviderID(uid.String()),
				state: shepherd.StateKnownUsed,
				addr:  netip.AddrFrom16(uid),
			},
		})
		if err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}

	return machines, nil
}

func (dp *dummyProvider) ListMachines(ctx context.Context) ([]shepherd.Machine, error) {
	var machines []shepherd.Machine
	for _, m := range dp.machines {
		machines = append(machines, m)
	}

	unusedMachineCount := dp.capacity - len(machines)
	for i := 0; i < unusedMachineCount; i++ {
		uid := uuid.Must(uuid.NewRandom())
		machines = append(machines, &dummyMachine{
			id:    shepherd.ProviderID(uid.String()),
			state: shepherd.StateKnownUnused,
			addr:  netip.AddrFrom16(uid),
		})
	}

	return machines, nil
}

func (dp *dummyProvider) GetMachine(ctx context.Context, id shepherd.ProviderID) (shepherd.Machine, error) {
	for _, m := range dp.machines {
		if m.ID() == id {
			return m, nil
		}
	}

	return nil, shepherd.ErrMachineNotFound
}

func (dp *dummyProvider) CreateMachine(ctx context.Context, session *bmdb.Session, request shepherd.CreateMachineRequest) (shepherd.Machine, error) {
	dm := request.UnusedMachine.(*dummyMachine)

	err := session.Transact(ctx, func(q *model.Queries) error {
		// Create a new machine record within BMDB.
		m, err := q.NewMachine(ctx)
		if err != nil {
			return fmt.Errorf("while creating a new machine record in BMDB: %w", err)
		}

		p := model.MachineAddProvidedParams{
			MachineID:  m.MachineID,
			ProviderID: string(dm.id),
			Provider:   dp.Type(),
		}
		klog.Infof("Setting \"provided\" tag (ID: %s, PID: %s, Provider: %s).", p.MachineID, p.ProviderID, p.Provider)
		if err := q.MachineAddProvided(ctx, p); err != nil {
			return fmt.Errorf("while tagging machine active: %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	dm.state = shepherd.StateKnownUsed
	dp.machines[dm.id] = dm

	return dm, nil
}

func (dp *dummyProvider) Type() model.Provider {
	return model.ProviderNone
}
