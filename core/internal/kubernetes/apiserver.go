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
	"errors"
	"fmt"
	"io"
	"net"
	"os/exec"
	"path"

	"go.etcd.io/etcd/clientv3"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/pkg/fileargs"
)

type apiserverConfig struct {
	advertiseAddress net.IP
	serviceIPRange   net.IPNet
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

func getPKIApiserverConfig(consensusKV clientv3.KV) (*apiserverConfig, error) {
	var config apiserverConfig
	var err error
	config.idCA, _, err = getCert(consensusKV, "id-ca")
	config.kubeletClientCert, config.kubeletClientKey, err = getCert(consensusKV, "kubelet-client")
	config.aggregationCA, _, err = getCert(consensusKV, "aggregation-ca")
	config.aggregationClientCert, config.aggregationClientKey, err = getCert(consensusKV, "front-proxy-client")
	config.serverCert, config.serverKey, err = getCert(consensusKV, "apiserver")
	saPrivkey, err := consensusKV.Get(context.Background(), path.Join(etcdPath, "service-account-privkey.der"))
	if err != nil {
		return nil, fmt.Errorf("failed to get serviceaccount privkey: %w", err)
	}
	if len(saPrivkey.Kvs) != 1 {
		return nil, errors.New("failed to get serviceaccount privkey: not found")
	}
	config.serviceAccountPrivKey = saPrivkey.Kvs[0].Value
	return &config, nil
}

func runAPIServer(config apiserverConfig, output io.Writer) supervisor.Runnable {
	return func(ctx context.Context) error {
		args, err := fileargs.New()
		if err != nil {
			panic(err) // If this fails, something is very wrong. Just crash.
		}
		defer args.Close()
		cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kube-apiserver",
			fmt.Sprintf("--advertise-address=%v", config.advertiseAddress.String()),
			"--authorization-mode=Node,RBAC",
			args.FileOpt("--client-ca-file", "client-ca.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.idCA})),
			"--enable-admission-plugins=NodeRestriction,PodSecurityPolicy",
			"--enable-aggregator-routing=true",
			"--insecure-port=0",
			// Due to the magic of GRPC this really needs four slashes and a :0
			fmt.Sprintf("--etcd-servers=%v", "unix:////consensus/listener.sock:0"),
			args.FileOpt("--kubelet-client-certificate", "kubelet-client-cert.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.kubeletClientCert})),
			args.FileOpt("--kubelet-client-key", "kubelet-client-key.pem",
				pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.kubeletClientKey})),
			"--kubelet-preferred-address-types=Hostname",
			args.FileOpt("--proxy-client-cert-file", "aggregation-client-cert.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.aggregationClientCert})),
			args.FileOpt("--proxy-client-key-file", "aggregation-client-key.pem",
				pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.aggregationClientKey})),
			"--requestheader-allowed-names=front-proxy-client",
			args.FileOpt("--requestheader-client-ca-file", "aggregation-ca.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.aggregationCA})),
			"--requestheader-extra-headers-prefix=X-Remote-Extra-",
			"--requestheader-group-headers=X-Remote-Group",
			"--requestheader-username-headers=X-Remote-User",
			args.FileOpt("--service-account-key-file", "service-account-pubkey.pem",
				pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.serviceAccountPrivKey})),
			fmt.Sprintf("--service-cluster-ip-range=%v", config.serviceIPRange.String()),
			args.FileOpt("--tls-cert-file", "server-cert.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.serverCert})),
			args.FileOpt("--tls-private-key-file", "server-key.pem",
				pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.serverKey})),
		)
		if args.Error() != nil {
			return err
		}
		cmd.Stdout = output
		cmd.Stderr = output
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		err = cmd.Run()
		fmt.Fprintf(output, "apiserver stopped: %v\n", err)
		return err
	}
}
