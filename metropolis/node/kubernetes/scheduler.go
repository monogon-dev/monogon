// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package kubernetes

import (
	"context"
	"encoding/pem"
	"fmt"
	"os/exec"

	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/osbase/fileargs"
	"source.monogon.dev/osbase/supervisor"
)

type schedulerConfig struct {
	kubeConfig []byte
	serverCert []byte
	serverKey  []byte
	rootCA     []byte
}

func getPKISchedulerConfig(ctx context.Context, kpki *pki.PKI) (*schedulerConfig, error) {
	var config schedulerConfig
	var err error
	config.rootCA, _, err = kpki.Certificate(ctx, pki.IdCA)
	if err != nil {
		return nil, fmt.Errorf("failed to get ID root CA: %w", err)
	}
	config.serverCert, config.serverKey, err = kpki.Certificate(ctx, pki.Scheduler)
	if err != nil {
		return nil, fmt.Errorf("failed to get scheduler serving certificate: %w", err)
	}
	config.kubeConfig, err = kpki.Kubeconfig(ctx, pki.SchedulerClient, pki.KubernetesAPIEndpointForController)
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
			args.FileOpt("--client-ca-file", "root-ca.pem",
				pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: config.rootCA})),
			extraFeatureGates.AsFlag(),
		)
		if args.Error() != nil {
			return fmt.Errorf("failed to use fileargs: %w", err)
		}
		return supervisor.RunCommand(ctx, cmd, supervisor.ParseKLog())
	}
}
