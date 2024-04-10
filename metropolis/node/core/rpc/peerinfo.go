package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/identity"
	epb "source.monogon.dev/metropolis/proto/ext"
)

type Permissions map[epb.Permission]bool

func (p Permissions) String() string {
	var res []string

	for k := range p {
		res = append(res, k.String())
	}

	sort.Strings(res)
	return strings.Join(res, ", ")
}

// PeerInfo represents the Metropolis-level information about the remote side
// of a gRPC RPC, ie. about the calling client in server handlers and about the
// handling server in client code.
//
// Exactly one of {Node, User, Unauthenticated} will be non-nil.
type PeerInfo struct {
	// Node is the information about a peer Node, and identifies that the other side
	// of the connection is either a Node servicng gRPC requests for a cluster, or a
	// Node connecting to a gRPC service.
	Node *PeerInfoNode
	// User is the information about a peer User, and identifies that the other side
	// of the connection is a Metropolis user or manager (eg. owner). This will only
	// be set in service handlers, as users cannot serve gRPC connections.
	User *PeerInfoUser
	// Unauthenticated is set for incoming gRPC connections which that have the
	// Unauthenticated authorization extension set to true, and mark that the other
	// side of the connection has not been verified at all.
	Unauthenticated *PeerInfoUnauthenticated
}

// PeerInfoNode contains information about a Node on the other side of a gRPC
// connection.
type PeerInfoNode struct {
	// PublicKey is the ED25519 public key bytes of the node.
	PublicKey []byte

	// Permissions are the set of permissions this node has.
	Permissions Permissions
}

// PeerInfoUser contains information about a user on the other side of a gRPC
// connection.
type PeerInfoUser struct {
	// Identity is an opaque identifier for the user. MVP: Currently this is always
	// "manager".
	Identity string
}

type PeerInfoUnauthenticated struct {
	// SelfSignedPublicKey is the ED25519 public key bytes of the other side of the
	// connection, if that side presented a self-signed certificate to prove control
	// of a private key corresponding to this public key. If it did not present a
	// self-signed certificate that can be parsed for such a key, this will be nil.
	//
	// This can be used by code with expects Unauthenticated RPCs but wants to
	// authenticate the connection based on ownership of some keypair, for example
	// in the AAA.Escrow method.
	SelfSignedPublicKey []byte
}

// GetPeerInfo returns the PeerInfo of the peer of a gRPC connection, or nil if
// this connection does not carry any PeerInfo.
func GetPeerInfo(ctx context.Context) *PeerInfo {
	if pi, ok := ctx.Value(peerInfoKey).(*PeerInfo); ok {
		return pi
	}
	return nil
}

func (p *PeerInfo) CheckPermissions(need Permissions) error {
	if p.Unauthenticated != nil {
		// This generally shouldn't happen, as unauthenticated users shouldn't be
		// allowed to reach this part of the code - methods with Need != nil will not be
		// processed as unauthenticated for security, and will instead act as
		// authenticated methods and reject unauthenticated connections.
		for _, v := range need {
			if v {
				return status.Error(codes.Unauthenticated, "unauthenticated connection")
			}
		}
		return nil
	} else if p.User != nil {
		// MVP: all permissions are granted to all users.
		// TODO(q3k): check authz.Need once we have a user/identity system implemented.
		return nil
	} else if p.Node != nil {
		for n, v := range need {
			if v && !p.Node.Permissions[n] {
				return status.Errorf(codes.PermissionDenied, "node missing %s permission", n.String())
			}
		}
		return nil
	}

	return fmt.Errorf("invalid PeerInfo: neither Unauthenticated, User nor Node is set")
}

func (p *PeerInfo) String() string {
	if p == nil {
		return "nil"
	}
	switch {
	case p.Node != nil:
		return fmt.Sprintf("node: %s, %s", identity.NodeID(p.Node.PublicKey), p.Node.Permissions)
	case p.User != nil:
		return fmt.Sprintf("user: %s", p.User.Identity)
	case p.Unauthenticated != nil:
		return fmt.Sprintf("unauthenticated: pubkey %s", hex.EncodeToString(p.Unauthenticated.SelfSignedPublicKey))
	default:
		return "invalid"
	}
}

type peerInfoKeyType string

// peerInfoKey is the context key for storing PeerInfo.
const peerInfoKey = peerInfoKeyType("peerInfo")

// apply returns the given context with itself stored under a unique key, that
// can be later retrieved via GetPeerInfo.
func (p *PeerInfo) apply(ctx context.Context) context.Context {
	return context.WithValue(ctx, peerInfoKey, p)
}

// peerInfoServerStream is a grpc.ServerStream wrapper which contains some
// PeerInfo, and returns it as part of the Context() of the ServerStream.
type peerInfoServerStream struct {
	grpc.ServerStream
	pi *PeerInfo
}

func (p *peerInfoServerStream) Context() context.Context {
	return p.pi.apply(p.ServerStream.Context())
}

// serverStream wraps a grpc.ServerStream with a structure that attaches this
// PeerInfo in all contexts returned by Context().
func (p *PeerInfo) serverStream(ss grpc.ServerStream) grpc.ServerStream {
	return &peerInfoServerStream{ss, p}
}
