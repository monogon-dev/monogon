package curator

import (
	"context"
	"crypto/rand"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	apb "source.monogon.dev/metropolis/proto/api"
)

type leaderManagement struct {
	leadership
}

const (
	// registerTicketSize is the size, in bytes, of the RegisterTicket used to
	// perform early perimeter checks for nodes which wish to register into the
	// cluster.
	//
	// The size was picked to offer resistance against on-line bruteforcing attacks
	// in even the worst case scenario (no ratelimiting, no monitoring, zero latency
	// between attacker and cluster). 256 bits of entropy require 3.6e68 requests
	// per second to bruteforce the ticket within 10 years. The ticket doesn't need
	// to be manually copied by humans, so the relatively overkill size also doesn't
	// impact usability.
	registerTicketSize = 32
)

const (
	// registerTicketEtcdPath is the etcd key under which private.RegisterTicket is
	// stored.
	registerTicketEtcdPath = "/global/register_ticket"
)

func (l *leaderManagement) GetRegisterTicket(ctx context.Context, req *apb.GetRegisterTicketRequest) (*apb.GetRegisterTicketResponse, error) {
	// Retrieve existing ticket, if any.
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(registerTicketEtcdPath))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not retrieve register ticket: %v", err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	if len(kvs) > 0 {
		// Ticket already generated, return.
		return &apb.GetRegisterTicketResponse{
			Ticket: kvs[0].Value,
		}, nil
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

	return &apb.GetRegisterTicketResponse{
		Ticket: ticketBytes,
	}, nil
}
