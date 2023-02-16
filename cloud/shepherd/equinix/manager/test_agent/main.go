// test_agent is used by the Equinix Metal Manager test code. Its only role
// is to ensure successful delivery of the BMaaS agent executable to the test
// hosts, together with its subsequent execution.
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/cloud/agent/api"
)

func main() {
	// The agent initialization message will arrive from Shepherd on Agent's
	// standard input.
	aimb, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "while reading AgentInit message: %v\n", err)
		return
	}
	var aim apb.TakeoverInit
	if err := proto.Unmarshal(aimb, &aim); err != nil {
		fmt.Fprintf(os.Stderr, "while unmarshaling TakeoverInit message: %v\n", err)
		return
	}

	// Agent should send back apb.TakeoverResponse on its standard output.
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "while generating agent public key: %v\n", err)
		return
	}
	arsp := apb.TakeoverResponse{
		InitMessage: &aim,
		Key:         pub,
	}
	arspb, err := proto.Marshal(&arsp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "while marshaling TakeoverResponse message: %v\n", err)
		return
	}
	if _, err := os.Stdout.Write(arspb); err != nil {
		fmt.Fprintf(os.Stderr, "while writing TakeoverResponse message: %v\n", err)
	}
	// The agent must detach and/or terminate after sending back the reply.
	// Failure to do so will leave the session hanging.
}
