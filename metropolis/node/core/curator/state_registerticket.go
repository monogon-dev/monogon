// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package curator

import (
	"context"
	"crypto/rand"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/node/core/rpc"
)

// ensureRegisterTicket returns the cluster's current RegisterTicket, creating
// one if not yet present in the cluster state.
func (l *leadership) ensureRegisterTicket(ctx context.Context) ([]byte, error) {
	l.muRegisterTicket.Lock()
	defer l.muRegisterTicket.Unlock()

	rpc.Trace(ctx).Printf("ensureRegisterTicket()...")

	// Retrieve existing ticket, if any.
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(registerTicketEtcdPath))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not retrieve register ticket: %v", err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	if len(kvs) > 0 {
		// Ticket already generated, return.
		rpc.Trace(ctx).Printf("ensureRegisterTicket(): ticket already exists")
		return kvs[0].Value, nil
	}

	// No ticket, generate one.
	ticket := &ppb.RegisterTicket{
		Opaque: make([]byte, registerTicketSize),
	}
	_, err = rand.Read(ticket.Opaque)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not generate new ticket: %v", err)
	}
	ticketBytes, err := proto.Marshal(ticket)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not marshal new ticket: %v", err)
	}

	// Commit new ticket to etcd.
	_, err = l.txnAsLeader(ctx, clientv3.OpPut(registerTicketEtcdPath, string(ticketBytes)))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not save new ticket: %v", err)
	}

	rpc.Trace(ctx).Printf("ensureRegisterTicket(): generated and saved new ticket")

	return ticketBytes, nil
}
