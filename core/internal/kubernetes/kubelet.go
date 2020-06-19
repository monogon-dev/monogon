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
	"io/ioutil"
	"net"
	"os"
	"os/exec"

	"go.etcd.io/etcd/clientv3"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/pki"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/reconciler"
	"git.monogon.dev/source/nexantic.git/core/pkg/fileargs"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeletconfig "k8s.io/kubelet/config/v1beta1"
)

var (
	kubeletRoot       = "/data/kubernetes"
	kubeletKubeconfig = kubeletRoot + "/kubelet.kubeconfig"
	kubeletCACert     = kubeletRoot + "/ca.crt"
	kubeletCert       = kubeletRoot + "/kubelet.crt"
	kubeletKey        = kubeletRoot + "/kubelet.key"
)

type KubeletSpec struct {
	clusterDNS []net.IP
}

func createKubeletConfig(ctx context.Context, kv clientv3.KV, kpki *pki.KubernetesPKI, nodeName string) error {
	identity := fmt.Sprintf("system:node:%s", nodeName)

	ca := kpki.Certificates[pki.IdCA]
	cacert, _, err := ca.Ensure(ctx, kv)
	if err != nil {
		return fmt.Errorf("could not ensure ca certificate: %w", err)
	}

	kubeconfig, err := pki.New(ca, "", pki.Client(identity, []string{"system:nodes"})).Kubeconfig(ctx, kv)
	if err != nil {
		return fmt.Errorf("could not create volatile kubelet client cert: %w", err)
	}

	cert, key, err := pki.New(ca, "volatile", pki.Server([]string{nodeName}, nil)).Ensure(ctx, kv)
	if err != nil {
		return fmt.Errorf("could not create volatile kubelet server cert: %w", err)
	}

	if err := os.MkdirAll(kubeletRoot, 0755); err != nil {
		return fmt.Errorf("could not create kubelet root directory: %w", err)
	}
	// TODO(q3k): this should probably become its own function //core/internal/kubernetes/pki.
	for _, el := range []struct {
		target string
		data   []byte
	}{
		{kubeletKubeconfig, kubeconfig},
		{kubeletCACert, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cacert})},
		{kubeletCert, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})},
		{kubeletKey, pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})},
	} {
		if err := ioutil.WriteFile(el.target, el.data, 0400); err != nil {
			return fmt.Errorf("could not write %q: %w", el.target, err)
		}
	}

	return nil
}

func runKubelet(spec *KubeletSpec, output io.Writer) supervisor.Runnable {
	return func(ctx context.Context) error {
		fargs, err := fileargs.New()
		if err != nil {
			return err
		}
		var clusterDNS []string
		for _, dnsIP := range spec.clusterDNS {
			clusterDNS = append(clusterDNS, dnsIP.String())
		}

		kubeletConf := &kubeletconfig.KubeletConfiguration{
			TypeMeta: v1.TypeMeta{
				Kind:       "KubeletConfiguration",
				APIVersion: kubeletconfig.GroupName + "/v1beta1",
			},
			TLSCertFile:       "/data/kubernetes/kubelet.crt",
			TLSPrivateKeyFile: "/data/kubernetes/kubelet.key",
			TLSMinVersion:     "VersionTLS13",
			ClusterDNS:        clusterDNS,
			Authentication: kubeletconfig.KubeletAuthentication{
				X509: kubeletconfig.KubeletX509Authentication{
					ClientCAFile: "/data/kubernetes/ca.crt",
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
		cmd.Stdout = output
		cmd.Stderr = output

		supervisor.Signal(ctx, supervisor.SignalHealthy)
		err = cmd.Run()
		fmt.Fprintf(output, "kubelet stopped: %v\n", err)
		return err
	}
}
