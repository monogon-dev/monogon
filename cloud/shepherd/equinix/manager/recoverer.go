package manager

import (
	"context"
	"flag"
	"fmt"
	"time"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	ecl "source.monogon.dev/cloud/shepherd/equinix/wrapngo"
)

type RecovererConfig struct {
	ControlLoopConfig
	RebootWaitSeconds int
}

func (r *RecovererConfig) RegisterFlags() {
	r.ControlLoopConfig.RegisterFlags("recoverer")
	flag.IntVar(&r.RebootWaitSeconds, "recoverer_reboot_wait_seconds", 30, "How many seconds to sleep to ensure a reboot happend")
}

// The Recoverer reboots machines whose agent has stopped sending heartbeats or
// has not sent any heartbeats at all.
type Recoverer struct {
	RecovererConfig

	cl ecl.Client
}

func NewRecoverer(cl ecl.Client, rc RecovererConfig) (*Recoverer, error) {
	if err := rc.ControlLoopConfig.Check(); err != nil {
		return nil, err
	}
	return &Recoverer{
		RecovererConfig: rc,
		cl:              cl,
	}, nil
}

func (r *Recoverer) getProcessInfo() processInfo {
	return processInfo{
		process: model.ProcessShepherdRecovery,
		defaultBackoff: bmdb.Backoff{
			Initial:  1 * time.Minute,
			Maximum:  1 * time.Hour,
			Exponent: 1.2,
		},
		processor: metrics.ProcessorShepherdRecoverer,
	}
}

func (r *Recoverer) getMachines(ctx context.Context, q *model.Queries, limit int32) ([]model.MachineProvided, error) {
	return q.GetMachineForAgentRecovery(ctx, limit)
}

func (r *Recoverer) processMachine(ctx context.Context, t *task) error {
	klog.Infof("Starting recovery of device (ID: %s, PID %s)", t.machine.MachineID, t.machine.ProviderID)

	if err := r.cl.RebootDevice(ctx, t.machine.ProviderID); err != nil {
		return fmt.Errorf("failed to reboot device: %w", err)
	}

	// TODO(issue/215): replace this
	// This is required as Equinix doesn't reboot the machines synchronously
	// during the API call.
	select {
	case <-time.After(time.Duration(r.RebootWaitSeconds) * time.Second):
	case <-ctx.Done():
		return fmt.Errorf("while waiting for reboot: %w", ctx.Err())
	}

	klog.Infof("Removing AgentStarted/AgentHeartbeat (ID: %s, PID: %s)...", t.machine.MachineID, t.machine.ProviderID)
	err := t.work.Finish(ctx, func(q *model.Queries) error {
		if err := q.MachineDeleteAgentStarted(ctx, t.machine.MachineID); err != nil {
			return fmt.Errorf("while deleting AgentStarted: %w", err)
		}
		if err := q.MachineDeleteAgentHeartbeat(ctx, t.machine.MachineID); err != nil {
			return fmt.Errorf("while deleting AgentHeartbeat: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("while deleting AgentStarted/AgentHeartbeat tags: %w", err)
	}
	return nil
}
