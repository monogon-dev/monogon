package curator

import (
	"context"
	"crypto/rand"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
)

// ensureRegisterTicket returns the cluster's current RegisterTicket, creating
// one if not yet present in the cluster state.
func (l *leadership) ensureRegisterTicket(ctx context.Context) ([]byte, error) {
	l.muRegisterTicket.Lock()
	defer l.muRegisterTicket.Unlock()

	// Retrieve existing ticket, if any.
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(registerTicketEtcdPath))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not retrieve register ticket: %v", err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	if len(kvs) > 0 {
		// Ticket already generated, return.
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

	return ticketBytes, nil
}
