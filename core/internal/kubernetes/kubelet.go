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
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net"
	"os/exec"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/pki"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/reconciler"
	"git.monogon.dev/source/nexantic.git/core/internal/localstorage"
	"git.monogon.dev/source/nexantic.git/core/internal/localstorage/declarative"
	"git.monogon.dev/source/nexantic.git/core/pkg/fileargs"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfig "k8s.io/kubelet/config/v1beta1"
)

type kubeletService struct {
	NodeName           string
	ClusterDNS         []net.IP
	KubeletDirectory   *localstorage.DataKubernetesKubeletDirectory
	EphemeralDirectory *localstorage.EphemeralDirectory
	Output             io.Writer
	KPKI               *pki.KubernetesPKI
}

func (s *kubeletService) createCertificates(ctx context.Context) error {
	identity := fmt.Sprintf("system:node:%s", s.NodeName)

	ca := s.KPKI.Certificates[pki.IdCA]
	cacert, _, err := ca.Ensure(ctx, s.KPKI.KV)
	if err != nil {
		return fmt.Errorf("could not ensure ca certificate: %w", err)
	}

	kubeconfig, err := pki.New(ca, "", pki.Client(identity, []string{"system:nodes"})).Kubeconfig(ctx, s.KPKI.KV)
	if err != nil {
		return fmt.Errorf("could not create volatile kubelet client cert: %w", err)
	}

	cert, key, err := pki.New(ca, "", pki.Server([]string{s.NodeName}, nil)).Ensure(ctx, s.KPKI.KV)
	if err != nil {
		return fmt.Errorf("could not create volatile kubelet server cert: %w", err)
	}

	// TODO(q3k): this should probably become its own function //core/internal/kubernetes/pki.
	for _, el := range []struct {
		target declarative.FilePlacement
		data   []byte
	}{
		{s.KubeletDirectory.Kubeconfig, kubeconfig},
		{s.KubeletDirectory.PKI.CACertificate, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cacert})},
		{s.KubeletDirectory.PKI.Certificate, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})},
		{s.KubeletDirectory.PKI.Key, pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})},
	} {
		if err := el.target.Write(el.data, 0400); err != nil {
			return fmt.Errorf("could not write %v: %w", el.target, err)
		}
	}

	return nil
}

func (s *kubeletService) configure() *kubeletconfig.KubeletConfiguration {
	var clusterDNS []string
	for _, dnsIP := range s.ClusterDNS {
		clusterDNS = append(clusterDNS, dnsIP.String())
	}

	return &kubeletconfig.KubeletConfiguration{
		TypeMeta: v1.TypeMeta{
			Kind:       "KubeletConfiguration",
			APIVersion: kubeletconfig.GroupName + "/v1beta1",
		},
		TLSCertFile:       s.KubeletDirectory.PKI.Certificate.FullPath(),
		TLSPrivateKeyFile: s.KubeletDirectory.PKI.Key.FullPath(),
		TLSMinVersion:     "VersionTLS13",
		ClusterDNS:        clusterDNS,
		Authentication: kubeletconfig.KubeletAuthentication{
			X509: kubeletconfig.KubeletX509Authentication{
				ClientCAFile: s.KubeletDirectory.PKI.CACertificate.FullPath(),
			},
		},
		// TODO(q3k): move reconciler.False to a generic package, fix the following references.
		ClusterDomain:                "cluster.local", // cluster.local is hardcoded in the certificate too currently
		EnableControllerAttachDetach: reconciler.False(),
		HairpinMode:                  "none",
		MakeIPTablesUtilChains:       reconciler.False(), // We don't have iptables
		FailSwapOn:                   reconciler.False(), // Our kernel doesn't have swap enabled which breaks Kubelet's detection
		KubeReserved: map[string]string{
			"cpu":    "200m",
			"memory": "300Mi",
		},

		// We're not going to use this, but let's make it point to a known-empty directory in case anybody manages to
		// trigger it.
		VolumePluginDir: s.EphemeralDirectory.FlexvolumePlugins.FullPath(),
	}
}

func (s *kubeletService) Run(ctx context.Context) error {
	if err := s.createCertificates(ctx); err != nil {
		return fmt.Errorf("when creating certificates: %w", err)
	}

	configRaw, err := json.Marshal(s.configure())
	if err != nil {
		return fmt.Errorf("when marshaling kubelet configuration: %w", err)
	}

	fargs, err := fileargs.New()
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kubelet",
		fargs.FileOpt("--config", "config.json", configRaw),
		"--container-runtime=remote",
		fmt.Sprintf("--container-runtime-endpoint=unix://%s", s.EphemeralDirectory.Containerd.ClientSocket.FullPath()),
		fmt.Sprintf("--kubeconfig=%s", s.KubeletDirectory.Kubeconfig.FullPath()),
		fmt.Sprintf("--root-dir=%s", s.KubeletDirectory.FullPath()),
	)
	cmd.Env = []string{"PATH=/kubernetes/bin"}
	return supervisor.RunCommand(ctx, cmd)
}
