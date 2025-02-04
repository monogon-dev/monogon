// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package curator

import (
	"context"
	"crypto/ed25519"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/rpc"
	kpki "source.monogon.dev/metropolis/node/kubernetes/pki"
)

func issueKubernetesWorkerCertificates(ctx context.Context, kp *kpki.PKI, nodeID string, req *ipb.IssueCertificateRequest_KubernetesWorker) (*ipb.IssueCertificateResponse, error) {
	idca, err := kp.Certificates[kpki.IdCA].Ensure(ctx, kp.KV)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not ensure CA certificate: %v", err)
	}

	if len(req.KubeletPubkey) != ed25519.PublicKeySize {
		return nil, status.Error(codes.InvalidArgument, "kubelet pubkey must be set and valid")
	}
	if len(req.CsiProvisionerPubkey) != ed25519.PublicKeySize {
		return nil, status.Error(codes.InvalidArgument, "CSI provisioner pubkey must be set and valid")
	}
	if len(req.NetservicesPubkey) != ed25519.PublicKeySize {
		return nil, status.Error(codes.InvalidArgument, "network services pubkey must be set and valid")
	}

	kubeletServer, kubeletClient, err := kp.Kubelet(ctx, nodeID, req.KubeletPubkey)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not generate kubelet certificates: %v", err)
	}

	kubeletServerCert, err := kubeletServer.Ensure(ctx, kp.KV)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not ensure kubelet server certificate: %v", err)
	}
	kubeletClientCert, err := kubeletClient.Ensure(ctx, kp.KV)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not ensure kubelet client certificate: %v", err)
	}

	csiClient, err := kp.CSIProvisioner(ctx, nodeID, req.CsiProvisionerPubkey)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not generate CSI provisioner certificates: %v", err)
	}

	csiClientCert, err := csiClient.Ensure(ctx, kp.KV)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not ensure CSI provisioner client certificate: %v", err)
	}

	netservClient, err := kp.NetServices(ctx, nodeID, req.NetservicesPubkey)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not generate netservices client certificates: %v", err)
	}

	netservClientCert, err := netservClient.Ensure(ctx, kp.KV)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not ensure netservices client certificate: %v", err)
	}

	return &ipb.IssueCertificateResponse{
		Kind: &ipb.IssueCertificateResponse_KubernetesWorker_{
			KubernetesWorker: &ipb.IssueCertificateResponse_KubernetesWorker{
				IdentityCaCertificate:     idca,
				KubeletServerCertificate:  kubeletServerCert,
				KubeletClientCertificate:  kubeletClientCert,
				CsiProvisionerCertificate: csiClientCert,
				NetservicesCertificate:    netservClientCert,
			},
		},
	}, nil
}

func (l *leaderCurator) IssueCertificate(ctx context.Context, req *ipb.IssueCertificateRequest) (*ipb.IssueCertificateResponse, error) {
	// Get remote node.
	pi := rpc.GetPeerInfo(ctx)
	if pi == nil || pi.Node == nil {
		return nil, status.Error(codes.PermissionDenied, "only nodes can request certificates")
	}
	id := pi.Node.ID
	node, err := nodeLoad(ctx, l.leadership, id)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not load node info: %v", err)
	}

	pki, err := kpki.FromLocalConsensus(ctx, l.consensus)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "could not get kube PKI: %v", err)
	}

	// Issue certificate if appropriate.
	switch kind := req.Kind.(type) {
	case *ipb.IssueCertificateRequest_KubernetesWorker_:
		if node.kubernetesWorker == nil {
			rpc.Trace(ctx).Printf("refusing to issue kube worker certificates for node %s", id)
			return nil, status.Errorf(codes.PermissionDenied, "node %s cannot request a kubelet certificate", id)
		}
		return issueKubernetesWorkerCertificates(ctx, pki, node.ID(), kind.KubernetesWorker)
	default:
		return nil, status.Error(codes.InvalidArgument, "certificate kind must be set")
	}
}
