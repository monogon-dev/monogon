// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/cloud/api"
	"source.monogon.dev/cloud/apigw/model"
	"source.monogon.dev/cloud/lib/component"
)

func dut() *Server {
	return &Server{
		Config: Config{
			Component: component.ComponentConfig{
				GRPCListenAddress: ":0",
				DevCerts:          true,
				DevCertsPath:      "/tmp/foo",
			},
			Database: component.CockroachConfig{
				InMemory: true,
			},
			PublicListenAddress: ":0",
		},
	}
}

// TestPublicSimple ensures the public grpc-web listener is working.
func TestPublicSimple(t *testing.T) {
	s := dut()
	ctx := context.Background()
	s.Start(ctx)

	// Craft a gRPC-Web request from scratch. There doesn't seem to be a
	// well-supported library to do this.

	// The request is \0 ++ uint32be(len(req)) ++ req.
	msgBytes, err := proto.Marshal(&apb.WhoAmIRequest{})
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.BigEndian, byte(0))
	binary.Write(buf, binary.BigEndian, uint32(len(msgBytes)))
	buf.Write(msgBytes)

	// Perform the request. Set minimum headers required for gRPC-Web to recognize
	// this as a gRPC-Web request.
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/cloud.api.IAM/WhoAmI", s.ListenPublic), buf)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/grpc-web+proto")
	req.Header.Set("X-Grpc-Web", "1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not perform request: %v", err)
	}
	// Regardless for RPC status, 200 should always be returned.
	if want, got := 200, res.StatusCode; want != got {
		t.Errorf("Wanted code %d, got %d", want, got)
	}

	// Expect endpoint to return 'unimplemented'.
	code, _ := strconv.Atoi(res.Header.Get("Grpc-Status"))
	if want, got := uint32(codes.Unimplemented), uint32(code); want != got {
		t.Errorf("Wanted code %d, got %d", want, got)
	}
	if want, got := "unimplemented", res.Header.Get("Grpc-Message"); want != got {
		t.Errorf("Wanted message %q, got %q", want, got)
	}
}

// TestUserSimple makes sure we can add and retrieve users. This is a low-level
// test which mostly exercises the machinery to bring up a working database in
// tests.
func TestUserSimple(t *testing.T) {
	s := dut()
	ctx := context.Background()
	s.Start(ctx)

	db, err := s.Config.Database.Connect()
	if err != nil {
		t.Fatalf("Connecting to the database failed: %v", err)
	}
	q := model.New(db)

	// Start out with no account by sub 'test'.
	accounts, err := q.GetAccountByOIDC(ctx, "test")
	if err != nil {
		t.Fatalf("Retrieving accounts failed: %v", err)
	}
	if want, got := 0, len(accounts); want != got {
		t.Fatalf("Expected no accounts at first, got %d", got)
	}

	// Create a new test account for sub 'test'.
	_, err = q.InitializeAccountFromOIDC(ctx, model.InitializeAccountFromOIDCParams{
		AccountOidcSub:     "test",
		AccountDisplayName: "Test User",
	})
	if err != nil {
		t.Fatalf("Creating new account failed: %v", err)
	}

	// Expect this account to be available now.
	accounts, err = q.GetAccountByOIDC(ctx, "test")
	if err != nil {
		t.Fatalf("Retrieving accounts failed: %v", err)
	}
	if want, got := 1, len(accounts); want != got {
		t.Fatalf("Expected exactly one account after creation, got %d", got)
	}
	if want, got := "Test User", accounts[0].AccountDisplayName; want != got {
		t.Fatalf("Expected to read back display name %q, got %q", want, got)
	}
}
