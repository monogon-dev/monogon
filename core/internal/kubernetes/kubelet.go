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
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"os/exec"

	"net"

	"git.monogon.dev/source/nexantic.git/core/pkg/fileargs"
	"go.etcd.io/etcd/clientv3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubelet/config/v1beta1"
)

type KubeletSpec struct {
	clusterDNS []net.IP
}

func bootstrapLocalKubelet(consensusKV clientv3.KV, nodeName string) error {
	idCA, idKeyRaw, err := getCert(consensusKV, "id-ca")
	if err != nil {
		return err
	}
	idKey := ed25519.PrivateKey(idKeyRaw)
	cert, key, err := issueCertificate(clientCertTemplate("system:node:"+nodeName, []string{"system:nodes"}), idCA, idKey)
	if err != nil {
		return err
	}
	kubeconfig, err := makeLocalKubeconfig(idCA, cert, key)
	if err != nil {
		return err
	}

	serverCert, serverKey, err := issueCertificate(serverCertTemplate([]string{nodeName}, []net.IP{}), idCA, idKey)
	if err != nil {
		return err
	}
	if err := os.MkdirAll("/data/kubernetes", 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/kubelet.kubeconfig", kubeconfig, 0400); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/ca.crt", pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: idCA}), 0400); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/kubelet.crt", pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: serverCert}), 0400); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/kubelet.key", pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: serverKey}), 0400); err != nil {
		return err
	}

	return nil
}

func (s *Service) runKubelet(ctx context.Context, spec *KubeletSpec) error {
	fargs, err := fileargs.New()
	if err != nil {
		return err
	}
	var clusterDNS []string
	for _, dnsIP := range spec.clusterDNS {
		clusterDNS = append(clusterDNS, dnsIP.String())
	}

	kubeletConf := &v1beta1.KubeletConfiguration{
		TypeMeta: v1.TypeMeta{
			Kind:       "KubeletConfiguration",
			APIVersion: v1beta1.GroupName + "/v1beta1",
		},
		TLSCertFile:       "/data/kubernetes/kubelet.crt",
		TLSPrivateKeyFile: "/data/kubernetes/kubelet.key",
		TLSMinVersion:     "VersionTLS13",
		ClusterDNS:        clusterDNS,
		Authentication: v1beta1.KubeletAuthentication{
			X509: v1beta1.KubeletX509Authentication{
				ClientCAFile: "/data/kubernetes/ca.crt",
			},
		},
		ClusterDomain:                "cluster.local",
		EnableControllerAttachDetach: False(),
		HairpinMode:                  "none",
		MakeIPTablesUtilChains:       False(), // We don't have iptables
		FailSwapOn:                   False(), // Our kernel doesn't have swap enabled which breaks Kubelet's detection
		KubeReserved: map[string]string{
			"cpu":    "200m",
			"memory": "300Mi",
		},
		// We're not going to use this, but let's make it point to a known-empty directory in case anybody manages to
		// trigger it.
		VolumePluginDir: "/kubernetes/conf/flexvolume-plugins",
	}

	configRaw, err := json.Marshal(kubeletConf)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kubelet",
		fargs.FileOpt("--config", "config.json", configRaw),
		"--container-runtime=remote",
		"--container-runtime-endpoint=unix:///containerd/run/containerd.sock",
		"--kubeconfig=/data/kubernetes/kubelet.kubeconfig",
		"--root-dir=/data/kubernetes/kubelet",
	)
	cmd.Env = []string{"PATH=/kubernetes/bin"}
	cmd.Stdout = s.kubeletLogs
	cmd.Stderr = s.kubeletLogs

	err = cmd.Run()
	fmt.Fprintf(s.kubeletLogs, "kubelet stopped: %v\n", err)
	if ctx.Err() == context.Canceled {
		s.logger.Info("kubelet stopped", zap.Error(err))
	} else {
		s.logger.Warn("kubelet stopped unexpectedly", zap.Error(err))
	}
	return err
}
