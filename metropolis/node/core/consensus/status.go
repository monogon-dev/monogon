package consensus

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"fmt"
	"net"
	"strconv"

	"go.etcd.io/etcd/clientv3"

	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/pki"
)

// Watch returns a Event Value compatible Watcher for accessing the State of the
// consensus Service in a safe manner.
func (s *Service) Watch() Watcher {
	return Watcher{s.value.Watch()}
}

type Watcher struct {
	event.Watcher
}

func (w *Watcher) Get(ctx context.Context, opts ...event.GetOption) (*Status, error) {
	v, err := w.Watcher.Get(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return v.(*Status), nil
}

// Status of the consensus service. It represents either a running consensus
// service to which a client can connect and on which management can be
// performed, or a stopped service.
type Status struct {
	localPeerURL  string
	localMemberID uint64
	cl            *clientv3.Client
	ca            *pki.Certificate
	stopped       bool
}

// Running returns true if this status represents a running consensus service
// which can be connected to or managed. These calls are not guaranteed to
// succeed (as the server might have stopped in the meantime), but the caller
// can use this value as a hint to whether attempts to access the consensus
// service should be done.
func (s *Status) Running() bool {
	return !s.stopped
}

func (s *Status) pkiClient() (client.Namespaced, error) {
	return clientFor(s.cl, "namespaced", "etcd-pki")
}

// MetropolisClient returns a namespaced etcd client for use by the rest of the
// metropolis code (thtough the cluster bootstrap code). This method is
// deprecated, and will be replaced with more granular clients as the cluster
// bootstrap code gets refactored.
func (s *Status) MetropolisClient() (client.Namespaced, error) {
	return clientFor(s.cl, "namespaced", "metropolis")
}

// AddNode creates a new consensus member corresponding to a given Ed25519 node
// public key if one does not yet exist. The member will at first be marked as a
// Learner, ensuring it does not take part in quorum until it has finished
// catching up to the state of the etcd store. As it does, the autopromoter will
// turn it into a 'full' node and it will start taking part in the quorum and be
// able to perform all etcd operations.
func (s *Status) AddNode(ctx context.Context, pk ed25519.PublicKey, opts ...*AddNodeOption) (*JoinCluster, error) {
	clPKI, err := s.pkiClient()
	if err != nil {
		return nil, err
	}

	nodeID := identity.NodeID(pk)
	var extraNames []string
	name := nodeID
	port := int(node.ConsensusPort)
	for _, opt := range opts {
		if opt.externalAddress != "" {
			name = opt.externalAddress
			extraNames = append(extraNames, name)
		}
		if opt.externalPort != 0 {
			port = opt.externalPort
		}
	}

	member := pki.Certificate{
		Name:      nodeID,
		Namespace: &pkiNamespace,
		Issuer:    s.ca,
		Template:  pkiPeerCertificate(pk, extraNames),
		Mode:      pki.CertificateExternal,
		PublicKey: pk,
	}
	caBytes, err := s.ca.Ensure(ctx, clPKI)
	if err != nil {
		return nil, fmt.Errorf("could not ensure CA certificate: %w", err)
	}
	memberBytes, err := member.Ensure(ctx, clPKI)
	if err != nil {
		return nil, fmt.Errorf("could not ensure member certificate: %w", err)
	}
	caCert, err := x509.ParseCertificate(caBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse CA certificate: %w", err)
	}
	memberCert, err := x509.ParseCertificate(memberBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse newly issued member certificate: %w", err)
	}

	members, err := s.cl.MemberList(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve existing members: %w", err)
	}

	var existingNodes []ExistingNode
	var newExists bool
	for _, m := range members.Members {
		if m.Name == nodeID {
			newExists = true
		}
		if m.IsLearner {
			continue
		}
		if len(m.PeerURLs) < 1 {
			continue
		}
		existingNodes = append(existingNodes, ExistingNode{
			Name: m.Name,
			URL:  m.PeerURLs[0],
		})
	}

	crlW := s.ca.WatchCRL(clPKI)
	crl, err := crlW.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve initial CRL: %w", err)
	}

	if !newExists {
		addr := fmt.Sprintf("https://%s", net.JoinHostPort(name, strconv.Itoa(port)))
		if _, err := s.cl.MemberAddAsLearner(ctx, []string{addr}); err != nil {
			return nil, fmt.Errorf("could not add new member as learner: %w", err)
		}
	}

	return &JoinCluster{
		CACertificate:   caCert,
		NodeCertificate: memberCert,
		ExistingNodes:   existingNodes,
		InitialCRL:      crl,
	}, nil
}

// AddNodeOptions can be passed to AddNode to influence the behaviour of the
// function. Currently this is only used internally by tests.
type AddNodeOption struct {
	externalAddress string
	externalPort    int
}
