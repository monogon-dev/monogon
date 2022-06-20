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
	"fmt"
	"io"
	"net"
	"os/exec"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfig "k8s.io/kubelet/config/v1beta1"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/node/kubernetes/reconciler"
	"source.monogon.dev/metropolis/pkg/fileargs"
	opki "source.monogon.dev/metropolis/pkg/pki"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

type kubeletService struct {
	NodeName           string
	ClusterDNS         []net.IP
	ClusterDomain      string
	KubeletDirectory   *localstorage.DataKubernetesKubeletDirectory
	EphemeralDirectory *localstorage.EphemeralDirectory
	Output             io.Writer
	KPKI               *pki.PKI

	mount               *opki.FilesystemCertificate
	mountKubeconfigPath string
}

func (s *kubeletService) createCertificates(ctx context.Context) error {
	server, client, err := s.KPKI.VolatileKubelet(ctx, s.NodeName)
	if err != nil {
		return fmt.Errorf("when generating local kubelet credentials: %w", err)
	}

	clientKubeconfig, err := pki.Kubeconfig(ctx, s.KPKI.KV, client)
	if err != nil {
		return fmt.Errorf("when generating kubeconfig: %w", err)
	}

	// Use a single fileargs mount for server certificate and client kubeconfig.
	mounted, err := server.Mount(ctx, s.KPKI.KV)
	if err != nil {
		return fmt.Errorf("could not mount kubelet cert dir: %w", err)
	}
	// mounted is closed by Run() on process exit.

	s.mount = mounted
	s.mountKubeconfigPath = mounted.ArgPath("kubeconfig", clientKubeconfig)

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
		TLSCertFile:       s.mount.CertPath,
		TLSPrivateKeyFile: s.mount.KeyPath,
		TLSMinVersion:     "VersionTLS13",
		ClusterDNS:        clusterDNS,
		Authentication: kubeletconfig.KubeletAuthentication{
			X509: kubeletconfig.KubeletX509Authentication{
				ClientCAFile: s.mount.CACertPath,
			},
		},
		// TODO(q3k): move reconciler.False to a generic package, fix the following references.
		ClusterDomain:                s.ClusterDomain,
		EnableControllerAttachDetach: reconciler.False(),
		HairpinMode:                  "none",
		MakeIPTablesUtilChains:       reconciler.False(), // We don't have iptables
		FailSwapOn:                   reconciler.False(), // Our kernel doesn't have swap enabled which breaks Kubelet's detection
		KubeReserved: map[string]string{
			"cpu":    "200m",
			"memory": "300Mi",
		},

		// We're not going to use this, but let's make it point to a
		// known-empty directory in case anybody manages to trigger it.
		VolumePluginDir: s.EphemeralDirectory.FlexvolumePlugins.FullPath(),
	}
}

func (s *kubeletService) Run(ctx context.Context) error {
	if err := s.createCertificates(ctx); err != nil {
		return fmt.Errorf("when creating certificates: %w", err)
	}
	defer s.mount.Close()

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
		fmt.Sprintf("--kubeconfig=%s", s.mountKubeconfigPath),
		fmt.Sprintf("--root-dir=%s", s.KubeletDirectory.FullPath()),
	)
	cmd.Env = []string{"PATH=/kubernetes/bin"}
	return supervisor.RunCommand(ctx, cmd, supervisor.ParseKLog())
}
