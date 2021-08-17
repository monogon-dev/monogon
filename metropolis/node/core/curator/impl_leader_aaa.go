package curator

import (
	"context"
	"crypto/ed25519"
	"crypto/subtle"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/pkg/pki"
	apb "source.monogon.dev/metropolis/proto/api"
)

const (
	// initialOwnerPath is the etcd key under which private.InitialOwner is stored.
	initialOwnerEtcdPath = "/global/initial_owner"
)

type leaderAAA struct {
	leadership
}

// pubkeyFromGRPC returns the ed25519 public key presented by the client in any
// client certificate for a gRPC call. If no certificate is presented, nil is
// returned. If the connection is insecure or the client presented some invalid
// certificate configuration, a gRPC status is returned that can be directly
// passed to the client. Otherwise, the public key is returned.
//
// SECURITY: the public key is not verified to be authorized to perform any
// action,just to be a valid ed25519 key.
func pubkeyFromGRPC(ctx context.Context) (ed25519.PublicKey, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unavailable, "could not retrieve peer info")
	}
	tlsInfo, ok := p.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "connection not secure")
	}
	count := len(tlsInfo.State.PeerCertificates)
	if count == 0 {
		return nil, nil
	}
	if count > 1 {
		return nil, status.Errorf(codes.Unauthenticated, "exactly one client certificate must be sent (got %d)", count)
	}
	pk, ok := tlsInfo.State.PeerCertificates[0].PublicKey.(ed25519.PublicKey)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "client certificate must be for ed25519 key")
	}
	return pk, nil
}

// getOwnerPubkey returns the public key of the configured owner of the cluster.
//
// MVP: this should be turned into a proper user/entity system.
func (a *leaderAAA) getOwnerPubkey(ctx context.Context) (ed25519.PublicKey, error) {
	res, err := a.etcd.Get(ctx, initialOwnerEtcdPath)
	if err != nil {
		if !errors.Is(err, ctx.Err()) {
			// TODO(q3k): log
			return nil, status.Error(codes.Unavailable, "could not retrieve initial owner status in etcd")
		}
		return nil, err
	}
	if len(res.Kvs) != 1 {
		return nil, status.Error(codes.FailedPrecondition, "no initial owner set for cluster")
	}
	var iom ppb.InitialOwner
	if err := proto.Unmarshal(res.Kvs[0].Value, &iom); err != nil {
		return nil, status.Error(codes.FailedPrecondition, "initial owner data could not be unmarshaled")
	}

	if len(iom.PublicKey) != ed25519.PublicKeySize {
		return nil, status.Error(codes.FailedPrecondition, "initial owner publickey has invalid length")
	}
	return iom.PublicKey, nil
}

// Escrow implements the AAA Escrow gRPC method, but currently only for the
// initial cluster owner exchange workflow. That is, the client presents a
// self-signed certificate for the public key of the InitialClusterOwner public
// key defined in the cluster bootstrap configuration, and receives a
// certificate which can be used to perform further management actions.
func (a *leaderAAA) Escrow(srv apb.AAA_EscrowServer) error {
	ctx := srv.Context()

	// Receive Parameters from client. This tells us what identity the client wants
	// from us.
	msg, err := srv.Recv()
	if err != nil {
		return err
	}
	if msg.Parameters == nil {
		return status.Errorf(codes.InvalidArgument, "client parameters must be set")
	}

	// MVP: only support authenticating as 'owner' identity.
	if msg.Parameters.RequestedIdentityName != "owner" {
		return status.Errorf(codes.Unimplemented, "only owner escrow is currently implemented")
	}

	if len(msg.Parameters.PublicKey) != ed25519.PublicKeySize {
		return status.Errorf(codes.InvalidArgument, "client parameters public_key must be set and valid")
	}

	// The owner is authenticated by the InitialOwnerKey set during cluster
	// bootstrap, whose ownership is proven to the cluster by presenting a
	// self-signed certificate emitted for that key.
	//
	// TODO(q3k) The AAA proto doesn't really have a proof kind for this, for now we
	// go with REFRESH_CERTIFICATE. We should either make the AAA proto explicitly
	// handle this as a special KIND.
	pk, err := pubkeyFromGRPC(ctx)
	if err != nil {
		// If an error occurred, it's either because the connection is not secured by
		// TLS, or an invalid certificate was presented (ie. more then one cert, or a
		// non-ed25519 cert). Fail as per AAA proto.
		return err
	}
	if pk == nil {
		// No cert was presented, respond with REFRESH_CERTIFICATE request.
		err := srv.Send(&apb.EscrowFromServer{
			Needed: []*apb.EscrowFromServer_ProofRequest{
				{
					Kind: apb.EscrowFromServer_ProofRequest_KIND_REFRESH_CERTIFICATE,
				},
			},
		})
		if err != nil {
			return err
		}
		return status.Error(codes.Unauthenticated, "cannot proceed without refresh certificate proof at transport layer")
	}

	// MVP: only support parameters public_key == TLS public key.
	if subtle.ConstantTimeCompare(pk, msg.Parameters.PublicKey) != 1 {
		return status.Errorf(codes.Unimplemented, "client parameters public_key different from transport public key unimplemented")
	}

	// Check client public key is the same as the cluster owner pubkey.
	opk, err := a.getOwnerPubkey(ctx)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(pk, opk) != 1 {
		return status.Errorf(codes.PermissionDenied, "public key not authorized to escrow owner credentials")
	}

	// Everything okay, send response with certificate.
	//
	// MVP: The emitted certificate is valid forever.
	oc := pki.Certificate{
		Namespace: &pkiNamespace,
		Issuer:    pkiCA,
		Template:  pki.Client("owner", nil),
		Name:      "owner",
		Mode:      pki.CertificateExternal,
		PublicKey: pk,
	}
	ocBytes, err := oc.Ensure(ctx, a.etcd)
	if err != nil {
		return status.Errorf(codes.Unavailable, "ensuring new certificate failed: %v", err)
	}

	return srv.Send(&apb.EscrowFromServer{
		Fulfilled: []*apb.EscrowFromServer_ProofRequest{
			{
				Kind: apb.EscrowFromServer_ProofRequest_KIND_REFRESH_CERTIFICATE,
			},
		},
		EmittedCertificate: ocBytes,
	})
}
