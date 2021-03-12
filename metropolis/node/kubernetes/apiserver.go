// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	"context"
	"encoding/pem"
	"fmt"
	"net"
	"os/exec"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/pkg/fileargs"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

type apiserverService struct {
	KPKI                        *pki.PKI
	AdvertiseAddress            net.IP
	ServiceIPRange              net.IPNet
	EphemeralConsensusDirectory *localstorage.EphemeralConsensusDirectory

	// All PKI-related things are in DER
	idCA                  []byte
	kubeletClientCert     []byte
	kubeletClientKey      []byte
	aggregationCA         []byte
	aggregationClientCert []byte
	aggregationClientKey  []byte
	serviceAccountPrivKey []byte // In PKIX form
	serverCert            []byte
	serverKey             []byte
}

func (s *apiserverService) loadPKI(ctx context.Context) error {
	for _, el := range []struct {
		targetCert *[]byte
		targetKey  *[]byte
		name       pki.KubeCertificateName
	}{
		{&s.idCA, nil, pki.IdCA},
		{&s.kubeletClientCert, &s.kubeletClientKey, pki.APIServerKubeletClient},
		{&s.aggregationCA, nil, pki.AggregationCA},
		{&s.aggregationClientCert, &s.aggregationClientKey, pki.FrontProxyClient},
		{&s.serverCert, &s.serverKey, pki.APIServer},
	} {
		cert, key, err := s.KPKI.Certificate(ctx, el.name)
		if err != nil {
			return fmt.Errorf("could not load certificate %q from PKI: %w", el.name, err)
		}
		if el.targetCert != nil {
			*el.targetCert = cert
		}
		if el.targetKey != nil {
			*el.targetKey = key
		}
	}

	var err error
	s.serviceAccountPrivKey, err = s.KPKI.ServiceAccountKey(ctx)
	if err != nil {
		return fmt.Errorf("could not load serviceaccount privkey: %w", err)
	}
	return nil
}

func (s *apiserverService) Run(ctx context.Context) error {
	if err := s.loadPKI(ctx); err != nil {
		return fmt.Errorf("loading PKI data failed: %w", err)
	}
	args, err := fileargs.New()
	if err != nil {
		panic(err) // If this fails, something is very wrong. Just crash.
	}
	defer args.Close()

	cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kube-apiserver",
		fmt.Sprintf("--advertise-address=%v", s.AdvertiseAddress.String()),
		"--authorization-mode=Node,RBAC",
		args.FileOpt("--client-ca-file", "client-ca.pem",
			pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: s.idCA})),
		"--enable-admission-plugins=NodeRestriction,PodSecurityPolicy",
		"--enable-aggregator-routing=true",
		"--insecure-port=0",
		fmt.Sprintf("--secure-port=%v", common.KubernetesAPIPort),
		fmt.Sprintf("--etcd-servers=unix:///%s:0", s.EphemeralConsensusDirectory.ClientSocket.FullPath()),
		args.FileOpt("--kubelet-client-certificate", "kubelet-client-cert.pem",
			pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: s.kubeletClientCert})),
		args.FileOpt("--kubelet-client-key", "kubelet-client-key.pem",
			pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: s.kubeletClientKey})),
		"--kubelet-preferred-address-types=InternalIP",
		args.FileOpt("--proxy-client-cert-file", "aggregation-client-cert.pem",
			pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: s.aggregationClientCert})),
		args.FileOpt("--proxy-client-key-file", "aggregation-client-key.pem",
			pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: s.aggregationClientKey})),
		"--requestheader-allowed-names=front-proxy-client",
		args.FileOpt("--requestheader-client-ca-file", "aggregation-ca.pem",
			pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: s.aggregationCA})),
		"--requestheader-extra-headers-prefix=X-Remote-Extra-",
		"--requestheader-group-headers=X-Remote-Group",
		"--requestheader-username-headers=X-Remote-User",
		args.FileOpt("--service-account-key-file", "service-account-pubkey.pem",
			pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: s.serviceAccountPrivKey})),
		fmt.Sprintf("--service-cluster-ip-range=%v", s.ServiceIPRange.String()),
		args.FileOpt("--tls-cert-file", "server-cert.pem",
			pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: s.serverCert})),
		args.FileOpt("--tls-private-key-file", "server-key.pem",
			pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: s.serverKey})),
	)
	if args.Error() != nil {
		return err
	}
	return supervisor.RunCommand(ctx, cmd, supervisor.ParseKLog())
}
