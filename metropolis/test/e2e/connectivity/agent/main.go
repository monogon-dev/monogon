// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

// connectivity agent hosts test runner-defined network listeners and performs
// connectivity tests to other instances of itself.
// It runs in an OCI image and a test runner communicates with it over
// stdin/stdout with delimited protobufs. See the spec directory for the
// request/response definitions.

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/protobuf/encoding/protodelim"

	"source.monogon.dev/metropolis/test/e2e/connectivity/spec"
)

func main() {
	t := tester{
		servers: make(map[uint64]net.Listener),
	}
	stdinReader := bufio.NewReader(os.Stdin)
	for {
		var req spec.Request
		if err := protodelim.UnmarshalFrom(stdinReader, &req); err != nil {
			log.Fatalf("Unable to unmarshal request: %v", err)
		}
		var res spec.Response
		switch r := req.Req.(type) {
		case *spec.Request_Test:
			res.Res = &spec.Response_Test{Test: t.runTest(r.Test)}
		case *spec.Request_StartServer:
			res.Res = &spec.Response_StartServer{StartServer: t.startServer(r.StartServer)}
		case *spec.Request_StopServer:
			res.Res = &spec.Response_StopServer{StopServer: t.stopServer(r.StopServer)}
		default:
			log.Fatalf("Unknown request type: %T", r)
		}
		if _, err := protodelim.MarshalTo(os.Stdout, &res); err != nil {
			log.Fatalf("Unable to marshal response: %v", err)
		}
	}
}

type tester struct {
	servers map[uint64]net.Listener
}

func errToResponse(err error) *spec.TestResponse {
	switch {
	case errors.Is(err, os.ErrDeadlineExceeded) || errors.Is(err, context.DeadlineExceeded):
		return &spec.TestResponse{
			Result:           spec.TestResponse_CONNECTION_TIMEOUT,
			ErrorDescription: err.Error(),
		}
	default:
		return &spec.TestResponse{
			Result:           spec.TestResponse_CONNECTION_REJECTED,
			ErrorDescription: err.Error(),
		}
	}
}

func (t *tester) runTest(req *spec.TestRequest) *spec.TestResponse {
	conn, err := net.DialTimeout("tcp", req.Address, req.Timeout.AsDuration())
	if err != nil {
		return errToResponse(err)
	}
	defer conn.Close()
	var tokenRaw [8]byte
	conn.SetReadDeadline(time.Now().Add(req.Timeout.AsDuration()))
	if _, err := io.ReadFull(conn, tokenRaw[:]); err != nil {
		return errToResponse(err)
	}
	receivedToken := binary.LittleEndian.Uint64(tokenRaw[:])
	if receivedToken != req.Token {
		return &spec.TestResponse{
			Result:           spec.TestResponse_WRONG_TOKEN,
			ErrorDescription: fmt.Sprintf("Received token %d, wanted %d", receivedToken, req.Token),
		}
	}
	return &spec.TestResponse{
		Result: spec.TestResponse_SUCCESS,
	}
}

func (t *tester) startServer(req *spec.StartServerRequest) *spec.StartServerResponse {
	l, err := net.Listen("tcp", req.Address)
	if err != nil {
		return &spec.StartServerResponse{ErrorDescription: err.Error()}
	}
	t.servers[req.Token] = l
	go tokenServer(l, req.Token)
	return &spec.StartServerResponse{Ok: true}
}

func tokenServer(l net.Listener, token uint64) {
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		conn.Write(binary.LittleEndian.AppendUint64(nil, token))
		conn.Close()
	}
}

func (t *tester) stopServer(req *spec.StopServerRequest) *spec.StopServerResponse {
	t.servers[req.Token].Close()
	delete(t.servers, req.Token)
	return &spec.StopServerResponse{Ok: true}
}
