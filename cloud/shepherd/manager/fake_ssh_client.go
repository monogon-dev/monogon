// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package manager

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/cloud/agent/api"

	"source.monogon.dev/go/net/ssh"
)

// FakeSSHClient is an Client that pretends to start an agent, but in reality
// just responds with what an agent would respond on every execution attempt.
type FakeSSHClient struct{}

type fakeSSHConnection struct{}

func (f *FakeSSHClient) Dial(ctx context.Context, address string, timeout time.Duration) (ssh.Connection, error) {
	return &fakeSSHConnection{}, nil
}

func (f *fakeSSHConnection) Execute(ctx context.Context, command string, stdin []byte) (stdout []byte, stderr []byte, err error) {
	var aim apb.TakeoverInit
	if err := proto.Unmarshal(stdin, &aim); err != nil {
		return nil, nil, fmt.Errorf("while unmarshaling TakeoverInit message: %w", err)
	}

	// Agent should send back apb.TakeoverResponse on its standard output.
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("while generating agent public key: %w", err)
	}
	arsp := apb.TakeoverResponse{
		Result: &apb.TakeoverResponse_Success{Success: &apb.TakeoverSuccess{
			InitMessage: &aim,
			Key:         pub,
		}},
	}
	arspb, err := proto.Marshal(&arsp)
	if err != nil {
		return nil, nil, fmt.Errorf("while marshaling TakeoverResponse message: %w", err)
	}
	return arspb, nil, nil
}

func (f *fakeSSHConnection) Upload(ctx context.Context, targetPath string, _ io.Reader) error {
	if targetPath != "/fake/path" {
		return fmt.Errorf("unexpected target path in test")
	}
	return nil
}

func (f *fakeSSHConnection) Close() error {
	return nil
}
