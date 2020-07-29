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

// Package DNS provides a Kubernetes DNS server using CoreDNS.
package dns

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"text/template"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/pkg/fileargs"
)

type corefileSpec struct {
	KubeconfigPath string
	ClusterDomain  string
}

var corefileTemplate = template.Must(template.New("corefile").Parse(`
.:53 {
    errors
    health {
        lameduck 5s
    }
    kubernetes {{.ClusterDomain}} in-addr.arpa ip6.arpa {
		kubeconfig {{.KubeconfigPath}} default
        pods insecure
        fallthrough in-addr.arpa ip6.arpa
        ttl 30
    }
    forward . /etc/resolv.conf
    cache 30
    loadbalance
}
`))

type Service struct {
	Output        io.Writer
	Kubeconfig    []byte
	ClusterDomain string
}

func (s *Service) Run(ctx context.Context) error {
	args, err := fileargs.New()
	if err != nil {
		return fmt.Errorf("failed to create fileargs: %w", err)
	}
	defer args.Close()

	var corefile bytes.Buffer
	if err := corefileTemplate.Execute(&corefile, &corefileSpec{
		KubeconfigPath: args.ArgPath("kubeconfig", s.Kubeconfig),
		ClusterDomain:  s.ClusterDomain,
	}); err != nil {
		return fmt.Errorf("failed to execute Corefile template: %w", err)
	}

	cmd := exec.CommandContext(ctx, "/kubernetes/bin/coredns",
		args.FileOpt("-conf", "Corefile", corefile.Bytes()),
	)

	if args.Error() != nil {
		return fmt.Errorf("failed to use fileargs: %w", err)
	}

	cmd.Stdout = s.Output
	cmd.Stderr = s.Output

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	err = cmd.Run()
	fmt.Fprintf(s.Output, "coredns stopped: %v\n", err)
	return err
}
