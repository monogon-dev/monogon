package server

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	apb "source.monogon.dev/cloud/bmaas/server/api"
	"source.monogon.dev/cloud/lib/component"
	"source.monogon.dev/metropolis/node/core/rpc"
)

func dut() *Server {
	return &Server{
		Config: Config{
			Component: component.ComponentConfig{
				GRPCListenAddress: ":0",
				DevCerts:          true,
				DevCertsPath:      "/tmp/foo",
			},
			BMDB: bmdb.BMDB{
				Config: bmdb.Config{
					Database: component.CockroachConfig{
						InMemory: true,
					},
				},
			},
			PublicListenAddress: ":0",
		},
	}
}

// TestAgentCallbackService exercises the basic flow for submitting an agent
// heartbeat and hardware report.
func TestAgentCallbackService(t *testing.T) {
	s := dut()
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	s.Start(ctx)

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("could not generate keypair: %v", err)
	}

	sess, err := s.bmdb.StartSession(ctx)
	if err != nil {
		t.Fatalf("could not start session")
	}

	heartbeat := func(mid uuid.UUID) error {
		creds, err := rpc.NewEphemeralCredentials(priv, nil)
		if err != nil {
			t.Fatalf("could not generate ephemeral credentials: %v", err)
		}
		conn, err := grpc.Dial(s.ListenPublic, grpc.WithTransportCredentials(creds))
		if err != nil {
			t.Fatalf("Dial failed: %v", err)
		}
		defer conn.Close()

		stub := apb.NewAgentCallbackClient(conn)
		_, err = stub.Heartbeat(ctx, &apb.AgentHeartbeatRequest{
			MachineId:      mid.String(),
			HardwareReport: &apb.AgentHardwareReport{},
		})
		return err
	}

	// First, attempt to heartbeat for some totally made up machine ID. That should
	// fail.
	if err := heartbeat(uuid.New()); err == nil {
		t.Errorf("heartbeat for made up UUID should've failed")
	}

	// Create an actual machine in the BMDB alongside the expected pubkey within an
	// AgentStarted tag.
	var machine model.Machine
	err = sess.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		if err != nil {
			return err
		}
		err = q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   model.ProviderEquinix,
			ProviderID: "123",
		})
		if err != nil {
			return err
		}
		return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
			MachineID:      machine.MachineID,
			AgentStartedAt: time.Now(),
			AgentPublicKey: pub,
		})
	})
	if err != nil {
		t.Fatalf("could not create machine: %v", err)
	}

	// Now heartbeat with correct machine ID and key. This should succeed.
	if err := heartbeat(machine.MachineID); err != nil {
		t.Errorf("heartbeat should've succeeded, got: %v", err)
	}

	// TODO(q3k): test hardware report being attached once we have some debug API
	// for tags.
}

// TestOSInstallationFlow exercises the agent's OS installation request/report
// functionality.
func TestOSInstallationFlow(t *testing.T) {
	s := dut()
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	s.Start(ctx)

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("could not generate keypair: %v", err)
	}

	sess, err := s.bmdb.StartSession(ctx)
	if err != nil {
		t.Fatalf("could not start session")
	}

	heartbeat := func(mid uuid.UUID, report *apb.OSInstallationReport) (*apb.AgentHeartbeatResponse, error) {
		creds, err := rpc.NewEphemeralCredentials(priv, nil)
		if err != nil {
			t.Fatalf("could not generate ephemeral credentials: %v", err)
		}
		conn, err := grpc.Dial(s.ListenPublic, grpc.WithTransportCredentials(creds))
		if err != nil {
			t.Fatalf("Dial failed: %v", err)
		}
		defer conn.Close()

		stub := apb.NewAgentCallbackClient(conn)
		return stub.Heartbeat(ctx, &apb.AgentHeartbeatRequest{
			MachineId:          mid.String(),
			HardwareReport:     &apb.AgentHardwareReport{},
			InstallationReport: report,
		})
	}

	// Create machine with no OS installation request.
	var machine model.Machine
	err = sess.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		if err != nil {
			return err
		}
		err = q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   model.ProviderEquinix,
			ProviderID: "123",
		})
		if err != nil {
			return err
		}
		return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
			MachineID:      machine.MachineID,
			AgentStartedAt: time.Now(),
			AgentPublicKey: pub,
		})
	})
	if err != nil {
		t.Fatalf("could not create machine: %v", err)
	}

	// Expect successful heartbeat, but no OS installation request.
	hbr, err := heartbeat(machine.MachineID, nil)
	if err != nil {
		t.Fatalf("heartbeat: %v", err)
	}
	if hbr.InstallationRequest != nil {
		t.Fatalf("expected no installation request")
	}

	// Now add an OS installation request tag, and expect it to be returned.
	err = sess.Transact(ctx, func(q *model.Queries) error {
		req := apb.OSInstallationRequest{
			Generation: 123,
		}
		raw, _ := proto.Marshal(&req)
		return q.MachineSetOSInstallationRequest(ctx, model.MachineSetOSInstallationRequestParams{
			MachineID:                machine.MachineID,
			Generation:               req.Generation,
			OsInstallationRequestRaw: raw,
		})
	})
	if err != nil {
		t.Fatalf("could not add os installation request to machine: %v", err)
	}

	// Heartbeat a few times just to make sure every response is as expected.
	for i := 0; i < 3; i++ {
		hbr, err = heartbeat(machine.MachineID, nil)
		if err != nil {
			t.Fatalf("heartbeat: %v", err)
		}
		if hbr.InstallationRequest == nil || hbr.InstallationRequest.Generation != 123 {
			t.Fatalf("expected installation request for generation 123, got %+v", hbr.InstallationRequest)
		}
	}

	// Submit a report, expect no more request.
	hbr, err = heartbeat(machine.MachineID, &apb.OSInstallationReport{Generation: 123})
	if err != nil {
		t.Fatalf("heartbeat: %v", err)
	}
	if hbr.InstallationRequest != nil {
		t.Fatalf("expected no installation request")
	}

	// Submit a newer request, expect it to be returned.
	err = sess.Transact(ctx, func(q *model.Queries) error {
		req := apb.OSInstallationRequest{
			Generation: 234,
		}
		raw, _ := proto.Marshal(&req)
		return q.MachineSetOSInstallationRequest(ctx, model.MachineSetOSInstallationRequestParams{
			MachineID:                machine.MachineID,
			Generation:               req.Generation,
			OsInstallationRequestRaw: raw,
		})
	})
	if err != nil {
		t.Fatalf("could not update installation request: %v", err)
	}

	// Heartbeat a few times just to make sure every response is as expected.
	for i := 0; i < 3; i++ {
		hbr, err = heartbeat(machine.MachineID, nil)
		if err != nil {
			t.Fatalf("heartbeat: %v", err)
		}
		if hbr.InstallationRequest == nil || hbr.InstallationRequest.Generation != 234 {
			t.Fatalf("expected installation request for generation 234, got %+v", hbr.InstallationRequest)
		}
	}
}
