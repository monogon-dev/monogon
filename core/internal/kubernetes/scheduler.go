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
	"io"
	"os/exec"

	"go.etcd.io/etcd/clientv3"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/pki"
	"git.monogon.dev/source/nexantic.git/core/pkg/fileargs"
)

type schedulerConfig struct {
	kubeConfig []byte
	serverCert []byte
	serverKey  []byte
}

func getPKISchedulerConfig(consensusKV clientv3.KV) (*schedulerConfig, error) {
	var config schedulerConfig
	var err error
	config.serverCert, config.serverKey, err = pki.GetCert(consensusKV, "scheduler")
	if err != nil {
		return nil, fmt.Errorf("failed to get scheduler serving certificate: %w", err)
	}
	config.kubeConfig, err = pki.GetSingle(consensusKV, "scheduler.kubeconfig")
	if err != nil {
		return nil, fmt.Errorf("failed to get scheduler kubeconfig: %w", err)
	}
	return &config, nil
}

func runScheduler(config schedulerConfig, output io.Writer) supervisor.Runnable {
	return func(ctx context.Context) error {
		args, err := fileargs.New()
		if err != nil {
			panic(err) // If this fails, something is very wrong. Just crash.
		}
		defer args.Close()
		cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kube-scheduler",
			args.FileOpt("--kubeconfig", "kubeconfig", config.kubeConfig),
			"--port=0", // Kill insecure serving
			args.FileOpt("--tls-cert-file", "server-cert.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.serverCert})),
			args.FileOpt("--tls-private-key-file", "server-key.pem",
				pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.serverKey})),
		)
		if args.Error() != nil {
			return fmt.Errorf("failed to use fileargs: %w", err)
		}
		cmd.Stdout = output
		cmd.Stderr = output
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		err = cmd.Run()
		fmt.Fprintf(output, "scheduler stopped: %v\n", err)
		return err
	}
}
