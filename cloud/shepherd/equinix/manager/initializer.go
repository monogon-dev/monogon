package manager

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/packethost/packngo"
	"golang.org/x/crypto/ssh"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	apb "source.monogon.dev/cloud/agent/api"
	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	ecl "source.monogon.dev/cloud/shepherd/equinix/wrapngo"
)

// AgentConfig configures how the Initializer will deploy Agents on machines. In
// CLI scenarios, this should be populated from flags via RegisterFlags.
type AgentConfig struct {
	// Executable is the contents of the agent binary created and run
	// at the provisioned servers. Must be set.
	Executable []byte

	// TargetPath is a filesystem destination path used while uploading the BMaaS
	// agent executable to hosts as part of the initialization process. Must be set.
	TargetPath string

	// Endpoint is the address Agent will use to contact the BMaaS
	// infrastructure. Must be set.
	Endpoint string

	// SSHTimeout is the amount of time set aside for the initializing
	// SSH session to run its course. Upon timeout, the iteration would be
	// declared a failure. Must be set.
	SSHConnectTimeout time.Duration
	// SSHExecTimeout is the amount of time set aside for executing the agent and
	// getting its output once the SSH connection has been established. Upon timeout,
	// the iteration would be declared as failure. Must be set.
	SSHExecTimeout time.Duration
}

func (a *AgentConfig) RegisterFlags() {
	flag.Func("agent_executable_path", "Local filesystem path of agent binary to be uploaded", func(val string) error {
		if val == "" {
			return nil
		}
		data, err := os.ReadFile(val)
		if err != nil {
			return fmt.Errorf("could not read -agent_executable_path: %w", err)
		}
		a.Executable = data
		return nil
	})
	flag.StringVar(&a.TargetPath, "agent_target_path", "/root/agent", "Filesystem path where the agent will be uploaded to and ran from")
	flag.StringVar(&a.Endpoint, "agent_endpoint", "", "Address of BMDB Server to which the agent will attempt to connect")
	flag.DurationVar(&a.SSHConnectTimeout, "agent_ssh_connect_timeout", 2*time.Second, "Timeout for connecting over SSH to a machine")
	flag.DurationVar(&a.SSHExecTimeout, "agent_ssh_exec_timeout", 60*time.Second, "Timeout for connecting over SSH to a machine")
}

// InitializerConfig configures the broad agent initialization process. The
// specifics of how an agent is started are instead configured in Agent Config. In
// CLI scenarios, this should be populated from flags via RegisterFlags.
type InitializerConfig struct {
	// DBQueryLimiter limits the rate at which BMDB is queried for servers ready
	// for BMaaS agent initialization. Must be set.
	DBQueryLimiter *rate.Limiter
}

// flagLimiter configures a *rate.Limiter as a flag.
func flagLimiter(l **rate.Limiter, name, defval, help string) {
	syntax := "'duration,count' eg. '2m,10' for a 10-sized bucket refilled at one token every 2 minutes"
	help = help + fmt.Sprintf(" (default: %q, syntax: %s)", defval, syntax)
	flag.Func(name, help, func(val string) error {
		if val == "" {
			val = defval
		}
		parts := strings.Split(val, ",")
		if len(parts) != 2 {
			return fmt.Errorf("invalid syntax, want: %s", syntax)
		}
		duration, err := time.ParseDuration(parts[0])
		if err != nil {
			return fmt.Errorf("invalid duration: %w", err)
		}
		refill, err := strconv.ParseUint(parts[1], 10, 31)
		if err != nil {
			return fmt.Errorf("invalid refill rate: %w", err)
		}
		*l = rate.NewLimiter(rate.Every(duration), int(refill))
		return nil
	})
	flag.Set(name, defval)
}

func (i *InitializerConfig) RegisterFlags() {
	flagLimiter(&i.DBQueryLimiter, "initializer_db_query_rate", "250ms,8", "Rate limiting for BMDB queries")
}

// Initializer implements the BMaaS agent initialization process. Initialization
// entails asking the BMDB for machines that need the agent started
// (or-restarted) and acting upon that.
type Initializer struct {
	config       *InitializerConfig
	agentConfig  *AgentConfig
	sharedConfig *SharedConfig
	sshClient    SSHClient
	// cl is the packngo wrapper used by the initializer.
	cl ecl.Client
}

// task describes a single server currently being processed either in the
// context of agent initialization or recovery.
type task struct {
	// id is the BMDB-assigned machine identifier.
	id uuid.UUID
	// pid is an identifier assigned by the provider (Equinix).
	pid uuid.UUID
	// work is a machine lock facilitated by BMDB that prevents machines from
	// being processed by multiple workers at the same time.
	work *bmdb.Work
	// dev is a provider machine/device record.
	dev *packngo.Device
}

// New creates an Initializer instance, checking the InitializerConfig,
// SharedConfig and AgentConfig for errors.
func (c *InitializerConfig) New(cl ecl.Client, sc *SharedConfig, ac *AgentConfig) (*Initializer, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	if len(ac.Executable) == 0 {
		return nil, fmt.Errorf("agent executable not configured")
	}
	if ac.TargetPath == "" {
		return nil, fmt.Errorf("agent target path must be set")
	}
	if ac.Endpoint == "" {
		return nil, fmt.Errorf("agent endpoint must be set")
	}
	if ac.SSHConnectTimeout == 0 {
		return nil, fmt.Errorf("agent SSH connection timeout must be set")
	}
	if ac.SSHExecTimeout == 0 {
		return nil, fmt.Errorf("agent SSH execution timeout must be set")
	}
	if c.DBQueryLimiter == nil {
		return nil, fmt.Errorf("DBQueryLimiter must be configured")
	}
	return &Initializer{
		config:       c,
		sharedConfig: sc,
		agentConfig:  ac,
		sshClient:    &PlainSSHClient{},
		cl:           cl,
	}, nil
}

// Run the initializer blocking the current goroutine until the given context
// expires.
func (c *Initializer) Run(ctx context.Context, conn *bmdb.Connection) error {
	signer, err := c.sharedConfig.sshSigner()
	if err != nil {
		return fmt.Errorf("could not initialize signer: %w", err)
	}

	// Maintain a BMDB session as long as possible.
	var sess *bmdb.Session
	for {
		if sess == nil {
			sess, err = conn.StartSession(ctx)
			if err != nil {
				return fmt.Errorf("could not start BMDB session: %w", err)
			}
		}
		// Inside that session, run the main logic.
		err = c.runInSession(ctx, sess, signer)

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

// runInSession executes one iteration of the initializer's control loop within a
// BMDB session. This control loop attempts to start or re-start the agent on any
// machines that need this per the BMDB.
func (c *Initializer) runInSession(ctx context.Context, sess *bmdb.Session, signer ssh.Signer) error {
	t, err := c.source(ctx, sess)
	if err != nil {
		return fmt.Errorf("could not source machine: %w", err)
	}
	if t == nil {
		return nil
	}
	defer t.work.Cancel(ctx)

	klog.Infof("Machine %q needs agent start, fetching corresponding packngo device %q...", t.id, t.pid)
	dev, err := c.cl.GetDevice(ctx, c.sharedConfig.ProjectId, t.pid.String())
	if err != nil {
		klog.Errorf("failed to fetch device %q: %v", t.pid, err)
		d := 30 * time.Second
		err = t.work.Fail(ctx, &d, "failed to fetch device from equinix")
		return err
	}
	t.dev = dev

	err = c.init(ctx, signer, t)
	if err != nil {
		klog.Errorf("Failed to initialize: %v", err)
		d := 1 * time.Minute
		err = t.work.Fail(ctx, &d, fmt.Sprintf("failed to initialize machine: %v", err))
		return err
	}
	return nil
}

// startAgent runs the agent executable on the target device d, returning the
// agent's public key on success.
func (i *Initializer) startAgent(ctx context.Context, sgn ssh.Signer, d packngo.Device) ([]byte, error) {
	// Provide a bound on execution time in case we get stuck after the SSH
	// connection is established.
	sctx, sctxC := context.WithTimeout(ctx, i.agentConfig.SSHExecTimeout)
	defer sctxC()

	// Use the device's IP address exposed by Equinix API.
	ni := d.GetNetworkInfo()
	var addr string
	if ni.PublicIPv4 != "" {
		addr = net.JoinHostPort(ni.PublicIPv4, "22")
	} else if ni.PublicIPv6 != "" {
		addr = net.JoinHostPort(ni.PublicIPv6, "22")
	} else {
		return nil, fmt.Errorf("device (ID: %s) has no available addresses", d.ID)
	}
	klog.V(1).Infof("Dialing device (provider ID: %s, addr: %s).", d.ID, addr)

	conn, err := i.sshClient.Dial(sctx, addr, "root", sgn, i.agentConfig.SSHConnectTimeout)
	if err != nil {
		return nil, fmt.Errorf("while dialing the device: %w", err)
	}
	defer conn.Close()

	// Upload the agent executable.

	klog.Infof("Uploading the agent executable (provider ID: %s, addr: %s).", d.ID, addr)
	if err := conn.Upload(sctx, i.agentConfig.TargetPath, i.agentConfig.Executable); err != nil {
		return nil, fmt.Errorf("while uploading agent executable: %w", err)
	}
	klog.V(1).Infof("Upload successful (provider ID: %s, addr: %s).", d.ID, addr)

	// The initialization protobuf message will be sent to the agent on its
	// standard input.
	imsg := apb.TakeoverInit{
		Provider:      "equinix",
		ProviderId:    d.ID,
		BmaasEndpoint: i.agentConfig.Endpoint,
	}
	imsgb, err := proto.Marshal(&imsg)
	if err != nil {
		return nil, fmt.Errorf("while marshaling agent message: %w", err)
	}

	// Start the agent and wait for the agent's output to arrive.
	klog.V(1).Infof("Starting the agent executable at path %q (provider ID: %s).", i.agentConfig.TargetPath, d.ID)
	stdout, stderr, err := conn.Execute(ctx, i.agentConfig.TargetPath, imsgb)
	stderrStr := strings.TrimSpace(string(stderr))
	if stderrStr != "" {
		klog.Warningf("Agent stderr: %q", stderrStr)
	}
	if err != nil {
		return nil, fmt.Errorf("while starting the agent executable: %w", err)
	}

	var arsp apb.TakeoverResponse
	if err := proto.Unmarshal(stdout, &arsp); err != nil {
		return nil, fmt.Errorf("agent reply couldn't be unmarshaled: %w", err)
	}
	if !proto.Equal(&imsg, arsp.InitMessage) {
		return nil, fmt.Errorf("agent did not send back the init message.")
	}
	if len(arsp.Key) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("agent key length mismatch.")
	}
	klog.Infof("Started the agent (provider ID: %s, key: %s).", d.ID, hex.EncodeToString(arsp.Key))
	return arsp.Key, nil
}

// init initializes the server described by t, using BMDB session 'sess' to set
// the relevant BMDB tag on success, and 'sgn' to authenticate to the server.
func (ir *Initializer) init(ctx context.Context, sgn ssh.Signer, t *task) error {
	// Start the agent.
	klog.Infof("Starting agent on device (ID: %s, PID %s)", t.id, t.pid)
	apk, err := ir.startAgent(ctx, sgn, *t.dev)
	if err != nil {
		return fmt.Errorf("while starting the agent: %w", err)
	}

	// Agent startup succeeded. Set the appropriate BMDB tag, and release the
	// lock.
	klog.Infof("Setting AgentStarted (ID: %s, PID: %s, Agent public key: %s).", t.id, t.pid, hex.EncodeToString(apk))
	err = t.work.Finish(ctx, func(q *model.Queries) error {
		return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
			MachineID:      t.id,
			AgentStartedAt: time.Now(),
			AgentPublicKey: apk,
		})
	})
	if err != nil {
		return fmt.Errorf("while setting AgentStarted tag: %w", err)
	}
	return nil
}

// source supplies returns a BMDB-locked server ready for initialization, locked
// by a work item. If both task and error are nil, then there are no machines
// needed to be initialized.
// The returned work item in task _must_ be canceled or finished by the caller.
func (ir *Initializer) source(ctx context.Context, sess *bmdb.Session) (*task, error) {
	ir.config.DBQueryLimiter.Wait(ctx)

	var machine *model.MachineProvided
	work, err := sess.Work(ctx, model.ProcessShepherdAccess, func(q *model.Queries) ([]uuid.UUID, error) {
		machines, err := q.GetMachinesForAgentStart(ctx, 1)
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

	pid, err := uuid.Parse(machine.ProviderID)
	if err != nil {
		t := time.Hour
		work.Fail(ctx, &t, fmt.Sprintf("could not parse provider UUID %q", machine.ProviderID))
		return nil, fmt.Errorf("while parsing provider UUID: %w", err)
	}

	return &task{
		id:   machine.MachineID,
		pid:  pid,
		work: work,
	}, nil
}
