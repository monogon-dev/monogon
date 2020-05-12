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
	"go.uber.org/zap"
	"net"
	"os/exec"

	"go.etcd.io/etcd/clientv3"

	"git.monogon.dev/source/nexantic.git/core/pkg/fileargs"
)

type controllerManagerConfig struct {
	clusterNet net.IPNet
	// All PKI-related things are in DER
	kubeConfig            []byte
	rootCA                []byte
	serviceAccountPrivKey []byte // In PKCS#8 form
	serverCert            []byte
	serverKey             []byte
}

func getPKIControllerManagerConfig(consensusKV clientv3.KV) (*controllerManagerConfig, error) {
	var config controllerManagerConfig
	var err error
	config.rootCA, _, err = getCert(consensusKV, "id-ca")
	if err != nil {
		return nil, fmt.Errorf("failed to get ID root CA: %w", err)
	}
	config.serverCert, config.serverKey, err = getCert(consensusKV, "controller-manager")
	if err != nil {
		return nil, fmt.Errorf("failed to get controller-manager serving certificate: %w", err)
	}
	config.serviceAccountPrivKey, err = getSingle(consensusKV, "service-account-privkey.der")
	if err != nil {
		return nil, fmt.Errorf("failed to get serviceaccount privkey: %w", err)
	}
	config.kubeConfig, err = getSingle(consensusKV, "controller-manager.kubeconfig")
	if err != nil {
		return nil, fmt.Errorf("failed to get controller-manager kubeconfig: %w", err)
	}
	return &config, nil
}

func (s *Service) runControllerManager(ctx context.Context, config controllerManagerConfig) error {
	args, err := fileargs.New()
	if err != nil {
		panic(err) // If this fails, something is very wrong. Just crash.
	}
	defer args.Close()
	cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kube-controller-manager",
		args.FileOpt("--kubeconfig", "kubeconfig", config.kubeConfig),
		args.FileOpt("--service-account-private-key-file", "service-account-privkey.pem",
			pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.serviceAccountPrivKey})),
		args.FileOpt("--root-ca-file", "root-ca.pem",
			pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.rootCA})),
		"--port=0",                               // Kill insecure serving
		"--use-service-account-credentials=true", // Enables things like PSP enforcement
		fmt.Sprintf("--cluster-cidr=%v", config.clusterNet.String()),
		args.FileOpt("--tls-cert-file", "server-cert.pem",
			pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.serverCert})),
		args.FileOpt("--tls-private-key-file", "server-key.pem",
			pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.serverKey})),
	)
	if args.Error() != nil {
		return fmt.Errorf("failed to use fileargs: %w", err)
	}
	cmd.Stdout = s.controllerManagerLogs
	cmd.Stderr = s.controllerManagerLogs
	err = cmd.Run()
	fmt.Fprintf(s.controllerManagerLogs, "controller-manager stopped: %v\n", err)
	if ctx.Err() == context.Canceled {
		s.logger.Info("controller-manager stopped", zap.Error(err))
	} else {
		s.logger.Warn("controller-manager stopped unexpectedly", zap.Error(err))
	}
	return err
}
