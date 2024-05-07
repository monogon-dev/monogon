package curator

import (
	"context"
	"crypto/ed25519"
	"crypto/subtle"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	apb "source.monogon.dev/metropolis/proto/api"
	"source.monogon.dev/osbase/pki"
)

const (
	// initialOwnerPath is the etcd key under which private.InitialOwner is stored.
	initialOwnerEtcdPath = "/global/initial_owner"
)

type leaderAAA struct {
	*leadership
}

// getOwnerPubkey returns the public key of the configured owner of the cluster.
//
// MVP: this should be turned into a proper user/entity system.
func (a *leaderAAA) getOwnerPubkey(ctx context.Context) (ed25519.PublicKey, error) {
	res, err := a.etcd.Get(ctx, initialOwnerEtcdPath)
	if err != nil {
		if !errors.Is(err, ctx.Err()) {
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
	peerInfo := rpc.GetPeerInfo(ctx)
	if peerInfo == nil {
		return status.Error(codes.Unauthenticated, "no PeerInfo available")
	}
	if peerInfo.Unauthenticated == nil {
		return status.Error(codes.InvalidArgument, "connection is already authenticated")
	}

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
	pk := peerInfo.Unauthenticated.SelfSignedPublicKey
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
		Template:  identity.UserCertificate("owner"),
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
