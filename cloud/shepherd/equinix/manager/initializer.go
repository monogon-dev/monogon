package manager

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/packethost/packngo"
	"golang.org/x/crypto/ssh"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	apb "source.monogon.dev/cloud/agent/api"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	ecl "source.monogon.dev/cloud/shepherd/equinix/wrapngo"
)

// InitializerConfig configures how the Initializer will deploy Agents on
// machines. In CLI scenarios, this should be populated from flags via
// RegisterFlags.
type InitializerConfig struct {
	ControlLoopConfig

	// Executable is the contents of the agent binary created and run
	// at the provisioned servers. Must be set.
	Executable []byte

	// TargetPath is a filesystem destination path used while uploading the BMaaS
	// agent executable to hosts as part of the initialization process. Must be set.
	TargetPath string

	// Endpoint is the address Agent will use to contact the BMaaS
	// infrastructure. Must be set.
	Endpoint string
	// EndpointCACertificate is an optional DER-encoded (but not PEM-armored) X509
	// certificate used to populate the trusted CA store of the agent. It should be
	// set to the CA certificate of the endpoint if not using a system-trusted CA
	// certificate.
	EndpointCACertificate []byte

	// SSHTimeout is the amount of time set aside for the initializing
	// SSH session to run its course. Upon timeout, the iteration would be
	// declared a failure. Must be set.
	SSHConnectTimeout time.Duration
	// SSHExecTimeout is the amount of time set aside for executing the agent and
	// getting its output once the SSH connection has been established. Upon timeout,
	// the iteration would be declared as failure. Must be set.
	SSHExecTimeout time.Duration
}

func (a *InitializerConfig) RegisterFlags() {
	a.ControlLoopConfig.RegisterFlags("initializer")

	flag.Func("agent_executable_path", "Local filesystem path of agent binary to be uploaded", func(val string) error {
		if val == "" {
			return nil
		}
		data, err := os.ReadFile(val)
		if err != nil {
			return fmt.Errorf("could not read: %w", err)
		}
		a.Executable = data
		return nil
	})
	flag.StringVar(&a.TargetPath, "agent_target_path", "/root/agent", "Filesystem path where the agent will be uploaded to and ran from")
	flag.StringVar(&a.Endpoint, "agent_endpoint", "", "Address of BMDB Server to which the agent will attempt to connect")
	flag.Func("agent_endpoint_ca_certificate_path", "Path to PEM X509 CA certificate that the agent endpoint is serving with. If not set, the agent will attempt to use system CA certificates to authenticate the endpoint.", func(val string) error {
		if val == "" {
			return nil
		}
		data, err := os.ReadFile(val)
		if err != nil {
			return fmt.Errorf("could not read: %w", err)
		}
		block, _ := pem.Decode(data)
		if block.Type != "CERTIFICATE" {
			return fmt.Errorf("not a certificate")
		}
		_, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			return fmt.Errorf("invalid certificate: %w", err)
		}
		a.EndpointCACertificate = block.Bytes
		return nil
	})
	flag.DurationVar(&a.SSHConnectTimeout, "agent_ssh_connect_timeout", 2*time.Second, "Timeout for connecting over SSH to a machine")
	flag.DurationVar(&a.SSHExecTimeout, "agent_ssh_exec_timeout", 60*time.Second, "Timeout for connecting over SSH to a machine")
}

// The Initializer starts the agent on machines that aren't yet running it.
type Initializer struct {
	InitializerConfig

	sharedConfig *SharedConfig
	cl           ecl.Client
	signer       ssh.Signer
	sshClient    SSHClient
}

// NewInitializer creates an Initializer instance, checking the
// InitializerConfig, SharedConfig and AgentConfig for errors.
func NewInitializer(cl ecl.Client, ic InitializerConfig, sc *SharedConfig) (*Initializer, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	if err := ic.ControlLoopConfig.Check(); err != nil {
		return nil, err
	}

	if len(ic.Executable) == 0 {
		return nil, fmt.Errorf("agent executable not configured")
	}
	if ic.TargetPath == "" {
		return nil, fmt.Errorf("agent target path must be set")
	}
	if ic.Endpoint == "" {
		return nil, fmt.Errorf("agent endpoint must be set")
	}
	if ic.SSHConnectTimeout == 0 {
		return nil, fmt.Errorf("agent SSH connection timeout must be set")
	}
	if ic.SSHExecTimeout == 0 {
		return nil, fmt.Errorf("agent SSH execution timeout must be set")
	}

	signer, err := sc.sshSigner()
	if err != nil {
		return nil, fmt.Errorf("could not initialize signer: %w", err)
	}

	return &Initializer{
		InitializerConfig: ic,

		sharedConfig: sc,
		cl:           cl,
		sshClient:    &PlainSSHClient{},
		signer:       signer,
	}, nil
}

func (c *Initializer) getMachines(ctx context.Context, q *model.Queries, limit int32) ([]model.MachineProvided, error) {
	return q.GetMachinesForAgentStart(ctx, limit)
}

func (c *Initializer) processMachine(ctx context.Context, t *task) error {
	dev, err := c.cl.GetDevice(ctx, c.sharedConfig.ProjectId, t.machine.ProviderID, nil)
	if err != nil {
		return fmt.Errorf("while fetching device %q: %v", t.machine.ProviderID, err)
	}

	// Start the agent.
	klog.Infof("Starting agent on device (ID: %s, PID %s)", t.machine.MachineID, t.machine.ProviderID)
	apk, err := c.startAgent(ctx, c.signer, dev, t.machine.MachineID)
	if err != nil {
		return fmt.Errorf("while starting the agent: %w", err)
	}

	// Agent startup succeeded. Set the appropriate BMDB tag, and release the
	// lock.
	klog.Infof("Setting AgentStarted (ID: %s, PID: %s, Agent public key: %s).", t.machine.MachineID, t.machine.ProviderID, hex.EncodeToString(apk))
	err = t.work.Finish(ctx, func(q *model.Queries) error {
		return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
			MachineID:      t.machine.MachineID,
			AgentStartedAt: time.Now(),
			AgentPublicKey: apk,
		})
	})
	if err != nil {
		return fmt.Errorf("while setting AgentStarted tag: %w", err)
	}
	return nil
}

// startAgent runs the agent executable on the target device d, returning the
// agent's public key on success.
func (i *Initializer) startAgent(ctx context.Context, sgn ssh.Signer, d *packngo.Device, mid uuid.UUID) ([]byte, error) {
	// Provide a bound on execution time in case we get stuck after the SSH
	// connection is established.
	sctx, sctxC := context.WithTimeout(ctx, i.SSHExecTimeout)
	defer sctxC()

	// Use the device's IP address exposed by Equinix API.
	ni := d.GetNetworkInfo()
	var addr string
	if ni.PublicIPv4 != "" {
		addr = net.JoinHostPort(ni.PublicIPv4, "22")
	} else if ni.PublicIPv6 != "" {
		addr = net.JoinHostPort(ni.PublicIPv6, "22")
	} else {
		return nil, fmt.Errorf("device (machine ID: %s) has no available addresses", mid)
	}
	klog.V(1).Infof("Dialing device (machine ID: %s, addr: %s).", mid, addr)

	conn, err := i.sshClient.Dial(sctx, addr, "root", sgn, i.SSHConnectTimeout)
	if err != nil {
		return nil, fmt.Errorf("while dialing the device: %w", err)
	}
	defer conn.Close()

	// Upload the agent executable.

	klog.Infof("Uploading the agent executable (machine ID: %s, addr: %s).", mid, addr)
	if err := conn.Upload(sctx, i.TargetPath, i.Executable); err != nil {
		return nil, fmt.Errorf("while uploading agent executable: %w", err)
	}
	klog.V(1).Infof("Upload successful (machine ID: %s, addr: %s).", mid, addr)

	// The initialization protobuf message will be sent to the agent on its
	// standard input.
	imsg := apb.TakeoverInit{
		MachineId:     mid.String(),
		BmaasEndpoint: i.Endpoint,
		CaCertificate: i.EndpointCACertificate,
	}
	imsgb, err := proto.Marshal(&imsg)
	if err != nil {
		return nil, fmt.Errorf("while marshaling agent message: %w", err)
	}

	// Start the agent and wait for the agent's output to arrive.
	klog.V(1).Infof("Starting the agent executable at path %q (machine ID: %s).", i.TargetPath, mid)
	stdout, stderr, err := conn.Execute(ctx, i.TargetPath, imsgb)
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
	var successResp *apb.TakeoverSuccess
	switch r := arsp.Result.(type) {
	case *apb.TakeoverResponse_Error:
		return nil, fmt.Errorf("agent returned error: %v", r.Error.Message)
	case *apb.TakeoverResponse_Success:
		successResp = r.Success
	default:
		return nil, fmt.Errorf("agent returned unknown result of type %T", arsp.Result)
	}
	if !proto.Equal(&imsg, successResp.InitMessage) {
		return nil, fmt.Errorf("agent did not send back the init message.")
	}
	if len(successResp.Key) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("agent key length mismatch.")
	}
	klog.Infof("Started the agent (machine ID: %s, key: %s).", mid, hex.EncodeToString(successResp.Key))
	return successResp.Key, nil
}
