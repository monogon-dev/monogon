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
	"source.monogon.dev/cloud/lib/component"
)

// TestPublicSimple ensures the public grpc-web listener is working.
func TestPublicSimple(t *testing.T) {
	s := Server{
		Config: Config{
			Configuration: component.Configuration{
				GRPCListenAddress: ":0",
				DevCerts:          true,
				DevCertsPath:      "/tmp/foo",
			},
			PublicListenAddress: ":0",
		},
	}

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
