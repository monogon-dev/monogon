package manager

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/packethost/packngo"
	"golang.org/x/crypto/ssh"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/cloud/agent/api"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
)

// fakeSSHClient is an SSHClient that pretends to start an agent, but in reality
// just responds with what an agent would respond on every execution attempt.
type fakeSSHClient struct{}

type fakeSSHConnection struct{}

func (f *fakeSSHClient) Dial(ctx context.Context, address, username string, sshkey ssh.Signer, timeout time.Duration) (SSHConnection, error) {
	return &fakeSSHConnection{}, nil
}

func (f *fakeSSHConnection) Execute(ctx context.Context, command string, stdin []byte) (stdout []byte, stderr []byte, err error) {
	var aim apb.TakeoverInit
	if err := proto.Unmarshal(stdin, &aim); err != nil {
		return nil, nil, fmt.Errorf("while unmarshaling TakeoverInit message: %v", err)
	}

	// Agent should send back apb.TakeoverResponse on its standard output.
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("while generating agent public key: %v", err)
	}
	arsp := apb.TakeoverResponse{
		Result: &apb.TakeoverResponse_Success{Success: &apb.TakeoverSuccess{
			InitMessage: &aim,
			Key:         pub,
		}},
	}
	arspb, err := proto.Marshal(&arsp)
	if err != nil {
		return nil, nil, fmt.Errorf("while marshaling TakeoverResponse message: %v", err)
	}
	return arspb, nil, nil
}

func (f *fakeSSHConnection) Upload(ctx context.Context, targetPath string, data []byte) error {
	if targetPath != "/fake/path" {
		return fmt.Errorf("unexpected target path in test")
	}
	return nil
}

func (f *fakeSSHConnection) Close() error {
	return nil
}

type initializerDut struct {
	f    *fakequinix
	i    *Initializer
	bmdb *bmdb.Connection
	ctx  context.Context
}

func newInitializerDut(t *testing.T) *initializerDut {
	t.Helper()

	_, key, _ := ed25519.GenerateKey(rand.Reader)
	sc := SharedConfig{
		ProjectId:    "noproject",
		KeyLabel:     "somekey",
		Key:          key,
		DevicePrefix: "test-",
	}
	ic := InitializerConfig{
		ControlLoopConfig: ControlLoopConfig{
			DBQueryLimiter: rate.NewLimiter(rate.Every(time.Second), 10),
		},
		Executable:        []byte("beep boop i'm a real program"),
		TargetPath:        "/fake/path",
		Endpoint:          "example.com:1234",
		SSHConnectTimeout: time.Second,
		SSHExecTimeout:    time.Second,
	}

	f := newFakequinix(sc.ProjectId, 100)
	i, err := NewInitializer(f, ic, &sc)
	if err != nil {
		t.Fatalf("Could not create Initializer: %v", err)
	}

	b := bmdb.BMDB{
		Config: bmdb.Config{
			Database: component.CockroachConfig{
				InMemory: true,
			},
			ComponentName: "test",
			RuntimeInfo:   "test",
		},
	}
	conn, err := b.Open(true)
	if err != nil {
		t.Fatalf("Could not create in-memory BMDB: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	t.Cleanup(ctxC)

	if err := sc.SSHEquinixEnsure(ctx, f); err != nil {
		t.Fatalf("Failed to ensure SSH key: %v", err)
	}

	i.sshClient = &fakeSSHClient{}
	go RunControlLoop(ctx, conn, i)

	return &initializerDut{
		f:    f,
		i:    i,
		bmdb: conn,
		ctx:  ctx,
	}
}

// TestInitializerSmokes makes sure the Initializer doesn't go up in flames on
// the happy path.
func TestInitializerSmokes(t *testing.T) {
	dut := newInitializerDut(t)
	f := dut.f
	ctx := dut.ctx
	conn := dut.bmdb
	sc := dut.i.sharedConfig

	reservations, _ := f.ListReservations(ctx, sc.ProjectId)
	kid, err := sc.sshEquinixId(ctx, f)
	if err != nil {
		t.Fatalf("Failed to retrieve equinix key ID: %v", err)
	}
	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Failed to create BMDB session for verifiaction: %v", err)
	}

	// Create 10 provided machines for testing.
	for i := 0; i < 10; i++ {
		res := reservations[i]
		dev, _ := f.CreateDevice(ctx, &packngo.DeviceCreateRequest{
			Hostname:              fmt.Sprintf("test-%d", i),
			OS:                    "fake",
			ProjectID:             sc.ProjectId,
			HardwareReservationID: res.ID,
			ProjectSSHKeys:        []string{kid},
		})
		f.devices[dev.ID].Network = []*packngo.IPAddressAssignment{
			{
				IpAddressCommon: packngo.IpAddressCommon{
					ID:            "fake",
					Address:       "1.2.3.4",
					Management:    true,
					AddressFamily: 4,
					Public:        true,
				},
			},
		}
		err = sess.Transact(ctx, func(q *model.Queries) error {
			machine, err := q.NewMachine(ctx)
			if err != nil {
				return err
			}
			return q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
				MachineID:  machine.MachineID,
				Provider:   model.ProviderEquinix,
				ProviderID: dev.ID,
			})
		})
		if err != nil {
			t.Fatalf("Failed to create BMDB machine: %v", err)
		}
	}

	// Expect to find 0 machines needing start.
	for {
		time.Sleep(100 * time.Millisecond)

		var machines []model.MachineProvided
		err = sess.Transact(ctx, func(q *model.Queries) error {
			var err error
			machines, err = q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
				Limit:    100,
				Provider: model.ProviderEquinix,
			})
			return err
		})
		if err != nil {
			t.Fatalf("Failed to run Transaction: %v", err)
		}
		if len(machines) == 0 {
			break
		}
	}
}
