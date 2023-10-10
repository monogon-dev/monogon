package manager

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/go/mflags"
)

// task describes a single server currently being processed by a control loop.
type task struct {
	// machine is the machine data (including provider and provider ID) retrieved
	// from the BMDB.
	machine *model.MachineProvided
	// work is a machine lock facilitated by BMDB that prevents machines from
	// being processed by multiple workers at the same time.
	work *bmdb.Work
	// backoff is configured from processInfo.defaultBackoff but can be overridden by
	// processMachine to set a different backoff policy for specific failure modes.
	backoff bmdb.Backoff
}

// controlLoop is implemented by any component which should act as a BMDB-based
// control loop. Implementing these methods allows the given component to be
// started using RunControlLoop.
type controlLoop interface {
	getProcessInfo() processInfo

	// getMachines must return the list of machines ready to be processed by the
	// control loop for a given control loop implementation.
	getMachines(ctx context.Context, q *model.Queries, limit int32) ([]model.MachineProvided, error)
	// processMachine will be called within the scope of an active task/BMDB work by
	// the control loop logic.
	processMachine(ctx context.Context, t *task) error

	// getControlLoopConfig is implemented by ControlLoopConfig which should be
	// embedded by the control loop component. If not embedded, this method will have
	// to be implemented, too.
	getControlLoopConfig() *ControlLoopConfig
}

type processInfo struct {
	process        model.Process
	processor      metrics.Processor
	defaultBackoff bmdb.Backoff
}

// ControlLoopConfig should be embedded the every component which acts as a
// control loop. RegisterFlags should be called by the component whenever it is
// registering its own flags. Check should be called whenever the component is
// instantiated, after RegisterFlags has been called.
type ControlLoopConfig struct {
	// DBQueryLimiter limits the rate at which BMDB is queried for servers ready
	// for BMaaS agent initialization. Must be set.
	DBQueryLimiter *rate.Limiter

	// Parallelism is how many instances of the Initializer will be allowed to run in
	// parallel against the BMDB. This speeds up the process of starting/restarting
	// agents significantly, as one initializer instance can handle at most one agent
	// (re)starting process.
	//
	// If not set (ie. 0), default to 1. A good starting value for production
	// deployments is 10 or so.
	Parallelism int
}

func (c *ControlLoopConfig) getControlLoopConfig() *ControlLoopConfig {
	return c
}

// RegisterFlags should be called on this configuration whenever the embeddeding
// component/configuration is registering its own flags. The prefix should be the
// name of the component.
func (c *ControlLoopConfig) RegisterFlags(prefix string) {
	mflags.Limiter(&c.DBQueryLimiter, prefix+"_db_query_rate", "250ms,8", "Rate limiting for BMDB queries")
	flag.IntVar(&c.Parallelism, prefix+"_loop_parallelism", 1, "How many initializer instances to run in parallel, ie. how many agents to attempt to (re)start at once")
}

// Check should be called after RegisterFlags but before the control loop is ran.
// If an error is returned, the control loop cannot start.
func (c *ControlLoopConfig) Check() error {
	if c.DBQueryLimiter == nil {
		return fmt.Errorf("DBQueryLimiter must be configured")
	}
	if c.Parallelism == 0 {
		c.Parallelism = 1
	}
	return nil
}

// RunControlLoop runs the given controlLoop implementation against the BMDB. The
// loop will be run with the parallelism and rate configured by the
// ControlLoopConfig embedded or otherwise returned by the controlLoop.
func RunControlLoop(ctx context.Context, conn *bmdb.Connection, loop controlLoop) error {
	clr := &controlLoopRunner{
		loop:   loop,
		config: loop.getControlLoopConfig(),
	}
	return clr.run(ctx, conn)
}

// controlLoopRunner is a configured control loop with an underlying control loop
// implementation.
type controlLoopRunner struct {
	config *ControlLoopConfig
	loop   controlLoop
}

// run the control loops(s) (depending on opts.Parallelism) blocking the current
// goroutine until the given context expires and all provisioners quit.
func (r *controlLoopRunner) run(ctx context.Context, conn *bmdb.Connection) error {
	pinfo := r.loop.getProcessInfo()

	eg := errgroup.Group{}
	for j := 0; j < r.config.Parallelism; j += 1 {
		eg.Go(func() error {
			return r.runOne(ctx, conn, &pinfo)
		})
	}
	return eg.Wait()
}

// run the control loop blocking the current goroutine until the given context
// expires.
func (r *controlLoopRunner) runOne(ctx context.Context, conn *bmdb.Connection, pinfo *processInfo) error {
	var err error

	// Maintain a BMDB session as long as possible.
	var sess *bmdb.Session
	for {
		if sess == nil {
			sess, err = conn.StartSession(ctx, bmdb.SessionOption{Processor: pinfo.processor})
			if err != nil {
				return fmt.Errorf("could not start BMDB session: %w", err)
			}
		}
		// Inside that session, run the main logic.
		err := r.runInSession(ctx, sess, pinfo)

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

// runInSession executes one iteration of the control loop within a BMDB session.
// This control loop attempts to start or re-start the agent on any machines that
// need this per the BMDB.
func (r *controlLoopRunner) runInSession(ctx context.Context, sess *bmdb.Session, pinfo *processInfo) error {
	t, err := r.source(ctx, sess, pinfo)
	if err != nil {
		return fmt.Errorf("could not source machine: %w", err)
	}
	if t == nil {
		return nil
	}
	defer t.work.Cancel(ctx)

	if err := r.loop.processMachine(ctx, t); err != nil {
		klog.Errorf("Failed to process machine %s: %v", t.machine.MachineID, err)
		err = t.work.Fail(ctx, &t.backoff, fmt.Sprintf("failed to process: %v", err))
		return err
	}
	return nil
}

// source supplies returns a BMDB-locked server ready for processing by the
// control loop, locked by a work item. If both task and error are nil, then
// there are no machines needed to be initialized. The returned work item in task
// _must_ be canceled or finished by the caller.
func (r *controlLoopRunner) source(ctx context.Context, sess *bmdb.Session, pinfo *processInfo) (*task, error) {
	r.config.DBQueryLimiter.Wait(ctx)

	var machine *model.MachineProvided
	work, err := sess.Work(ctx, pinfo.process, func(q *model.Queries) ([]uuid.UUID, error) {
		machines, err := r.loop.getMachines(ctx, q, 1)
		if err != nil {
			return nil, err
		}
		if len(machines) < 1 {
			return nil, bmdb.ErrNothingToDo
		}
		machine = &machines[0]
		return []uuid.UUID{machines[0].MachineID}, nil
	})

	if errors.Is(err, bmdb.ErrNothingToDo) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("while querying BMDB agent candidates: %w", err)
	}

	return &task{
		machine: machine,
		work:    work,
		backoff: pinfo.defaultBackoff,
	}, nil
}
