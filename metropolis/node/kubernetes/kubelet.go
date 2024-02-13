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
	"crypto/ed25519"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net"
	"os/exec"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfig "k8s.io/kubelet/config/v1beta1"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/node/kubernetes/reconciler"
	"source.monogon.dev/metropolis/pkg/fileargs"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

type kubeletService struct {
	ClusterDNS         []net.IP
	ClusterDomain      string
	KubeletDirectory   *localstorage.DataKubernetesKubeletDirectory
	EphemeralDirectory *localstorage.EphemeralDirectory

	kubeconfig   []byte
	serverCACert []byte
	serverCert   []byte
}

func (s *kubeletService) getPubkey(ctx context.Context) (ed25519.PublicKey, error) {
	// First make sure we have a local ED25519 private key, and generate one if not.
	if err := s.KubeletDirectory.PKI.GeneratePrivateKey(); err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}
	priv, err := s.KubeletDirectory.PKI.ReadPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("could not read keypair: %w", err)
	}
	pubkey := priv.Public().(ed25519.PublicKey)
	return pubkey, nil
}

func (s *kubeletService) setCertificates(kw *ipb.IssueCertificateResponse_KubernetesWorker) error {
	key, err := s.KubeletDirectory.PKI.ReadPrivateKey()
	if err != nil {
		return fmt.Errorf("could not read private key from disk: %w", err)
	}

	s.kubeconfig, err = pki.KubeconfigRaw(kw.IdentityCaCertificate, kw.KubeletClientCertificate, key, pki.KubernetesAPIEndpointForWorker)
	if err != nil {
		return fmt.Errorf("when generating kubeconfig: %w", err)
	}
	s.serverCACert = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: kw.IdentityCaCertificate})
	s.serverCert = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: kw.KubeletServerCertificate})
	return nil
}

func (s *kubeletService) configure(fargs *fileargs.FileArgs) *kubeletconfig.KubeletConfiguration {
	var clusterDNS []string
	for _, dnsIP := range s.ClusterDNS {
		clusterDNS = append(clusterDNS, dnsIP.String())
	}

	return &kubeletconfig.KubeletConfiguration{
		TypeMeta: v1.TypeMeta{
			Kind:       "KubeletConfiguration",
			APIVersion: kubeletconfig.GroupName + "/v1beta1",
		},
		TLSCertFile:       fargs.ArgPath("server.crt", s.serverCert),
		TLSPrivateKeyFile: s.KubeletDirectory.PKI.Key.FullPath(),
		TLSMinVersion:     "VersionTLS13",
		ClusterDNS:        clusterDNS,
		Authentication: kubeletconfig.KubeletAuthentication{
			X509: kubeletconfig.KubeletX509Authentication{
				ClientCAFile: fargs.ArgPath("ca.crt", s.serverCACert),
			},
		},
		// TODO(q3k): move reconciler.False to a generic package, fix the following references.
		ClusterDomain:                s.ClusterDomain,
		EnableControllerAttachDetach: reconciler.False(),
		HairpinMode:                  "none",
		MakeIPTablesUtilChains:       reconciler.False(), // We don't have iptables
		FailSwapOn:                   reconciler.False(), // Our kernel doesn't have swap enabled which breaks Kubelet's detection
		CgroupRoot:                   "/",
		KubeReserved: map[string]string{
			"cpu":    "200m",
			"memory": "300Mi",
		},

		// We're not going to use this, but let's make it point to a
		// known-empty directory in case anybody manages to trigger it.
		VolumePluginDir: s.EphemeralDirectory.FlexvolumePlugins.FullPath(),
		// Currently we allocate a /24 per node, so we can have a maximum of
		// 253 pods per node.
		MaxPods: 253,
	}
}

func (s *kubeletService) Run(ctx context.Context) error {
	if len(s.serverCert) == 0 || len(s.serverCACert) == 0 || len(s.kubeconfig) == 0 {
		return fmt.Errorf("setCertificates was not called")
	}

	fargs, err := fileargs.New()
	if err != nil {
		return err
	}
	defer fargs.Close()

	configRaw, err := json.Marshal(s.configure(fargs))
	if err != nil {
		return fmt.Errorf("when marshaling kubelet configuration: %w", err)
	}

	cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kubelet",
		fargs.FileOpt("--config", "config.json", configRaw),
		fmt.Sprintf("--container-runtime-endpoint=unix://%s", s.EphemeralDirectory.Containerd.ClientSocket.FullPath()),
		//TODO: Remove with k8s 1.29 (https://github.com/kubernetes/kubernetes/pull/118544)
		"--pod-infra-container-image", "preseed.metropolis.internal/node/kubernetes/pause:latest",
		fargs.FileOpt("--kubeconfig", "kubeconfig", s.kubeconfig),
		fmt.Sprintf("--root-dir=%s", s.KubeletDirectory.FullPath()),
	)
	cmd.Env = []string{"PATH=/kubernetes/bin"}
	return supervisor.RunCommand(ctx, cmd, supervisor.ParseKLog())
}
