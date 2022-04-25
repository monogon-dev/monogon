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
	"os/exec"

	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/pkg/fileargs"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

type schedulerConfig struct {
	kubeConfig []byte
	serverCert []byte
	serverKey  []byte
}

func getPKISchedulerConfig(ctx context.Context, kpki *pki.PKI) (*schedulerConfig, error) {
	var config schedulerConfig
	var err error
	config.serverCert, config.serverKey, err = kpki.Certificate(ctx, pki.Scheduler)
	if err != nil {
		return nil, fmt.Errorf("failed to get scheduler serving certificate: %w", err)
	}
	config.kubeConfig, err = kpki.Kubeconfig(ctx, pki.SchedulerClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get scheduler kubeconfig: %w", err)
	}
	return &config, nil
}

func runScheduler(config schedulerConfig) supervisor.Runnable {
	return func(ctx context.Context) error {
		args, err := fileargs.New()
		if err != nil {
			panic(err) // If this fails, something is very wrong. Just crash.
		}
		defer args.Close()
		cmd := exec.CommandContext(ctx, "/kubernetes/bin/kube", "kube-scheduler",
			args.FileOpt("--kubeconfig", "kubeconfig", config.kubeConfig),
			args.FileOpt("--tls-cert-file", "server-cert.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.serverCert})),
			args.FileOpt("--tls-private-key-file", "server-key.pem",
				pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: config.serverKey})),
		)
		if args.Error() != nil {
			return fmt.Errorf("failed to use fileargs: %w", err)
		}
		return supervisor.RunCommand(ctx, cmd, supervisor.ParseKLog())
	}
}
