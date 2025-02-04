// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package manager

import (
	"context"
	"fmt"
	"time"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/shepherd"
)

type RecovererConfig struct {
	ControlLoopConfig
}

func (r *RecovererConfig) RegisterFlags() {
	r.ControlLoopConfig.RegisterFlags("recoverer")
}

// The Recoverer reboots machines whose agent has stopped sending heartbeats or
// has not sent any heartbeats at all.
type Recoverer struct {
	RecovererConfig
	r shepherd.Recoverer
}

func NewRecoverer(r shepherd.Recoverer, rc RecovererConfig) (*Recoverer, error) {
	if err := rc.ControlLoopConfig.Check(); err != nil {
		return nil, err
	}
	return &Recoverer{
		RecovererConfig: rc,
		r:               r,
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
	return q.GetMachineForAgentRecovery(ctx, model.GetMachineForAgentRecoveryParams{
		Limit:    limit,
		Provider: r.r.Type(),
	})
}

func (r *Recoverer) processMachine(ctx context.Context, t *task) error {
	klog.Infof("Starting recovery of machine (ID: %s, PID %s)", t.machine.MachineID, t.machine.ProviderID)

	if err := r.r.RebootMachine(ctx, shepherd.ProviderID(t.machine.ProviderID)); err != nil {
		return fmt.Errorf("failed to reboot machine: %w", err)
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
