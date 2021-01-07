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

// Package DNS provides a DNS server using CoreDNS.
package dns

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"syscall"

	"git.monogon.dev/source/nexantic.git/metropolis/pkg/fileargs"
	"git.monogon.dev/source/nexantic.git/metropolis/pkg/supervisor"
)

const corefileBase = `
.:53 {
    errors
	hosts {
		fallthrough
	}
	
    cache 30
    loadbalance
`

type Service struct {
	directiveRegistration chan *ExtraDirective
	directives            map[string]ExtraDirective
	cmd                   *exec.Cmd
	args                  *fileargs.FileArgs
	// stateMu guards access to the directives, cmd and args fields
	stateMu sync.Mutex
}

// New creates a new CoreDNS service.
// The given channel can then be used to dynamically register and unregister directives in the configuaration.
// To register a new directive, send an ExtraDirective on the channel. To remove it again, use CancelDirective()
// to create a removal message.
func New(directiveRegistration chan *ExtraDirective) *Service {
	return &Service{
		directives:            map[string]ExtraDirective{},
		directiveRegistration: directiveRegistration,
	}
}

func (s *Service) makeCorefile(fargs *fileargs.FileArgs) []byte {
	corefile := bytes.Buffer{}
	corefile.WriteString(corefileBase)
	for _, dir := range s.directives {
		resolvedDir := dir.directive
		for fname, fcontent := range dir.files {
			resolvedDir = strings.ReplaceAll(resolvedDir, fmt.Sprintf("$FILE(%v)", fname), fargs.ArgPath(fname, fcontent))
		}
		corefile.WriteString(resolvedDir)
		corefile.WriteString("\n")
	}
	corefile.WriteString("\n}")
	return corefile.Bytes()
}

// CancelDirective creates a message to cancel the given directive.
func CancelDirective(d *ExtraDirective) *ExtraDirective {
	return &ExtraDirective{
		ID: d.ID,
	}
}

// Run runs the DNS service consisting of the CoreDNS process and the directive registration process
func (s *Service) Run(ctx context.Context) error {
	supervisor.Run(ctx, "coredns", s.runCoreDNS)
	supervisor.Run(ctx, "registration", s.runRegistration)
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}

// runCoreDNS runs the CoreDNS proceess
func (s *Service) runCoreDNS(ctx context.Context) error {
	s.stateMu.Lock()
	args, err := fileargs.New()
	if err != nil {
		s.stateMu.Unlock()
		return fmt.Errorf("failed to create fileargs: %w", err)
	}
	defer args.Close()
	s.args = args

	s.cmd = exec.CommandContext(ctx, "/kubernetes/bin/coredns",
		args.FileOpt("-conf", "Corefile", s.makeCorefile(args)),
	)

	if args.Error() != nil {
		s.stateMu.Unlock()
		return fmt.Errorf("failed to use fileargs: %w", err)
	}

	s.stateMu.Unlock()
	return supervisor.RunCommand(ctx, s.cmd)
}

// runRegistration runs the background registration runnable which has a different lifecycle from the CoreDNS
// runnable. It is responsible for managing dynamic directives.
func (s *Service) runRegistration(ctx context.Context) error {
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	for {
		select {
		case <-ctx.Done():
			return nil
		case d := <-s.directiveRegistration:
			s.processRegistration(ctx, d)
		}
	}
}

func (s *Service) processRegistration(ctx context.Context, d *ExtraDirective) {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	if d.directive == "" {
		delete(s.directives, d.ID)
	} else {
		s.directives[d.ID] = *d
	}
	// If the process is not currenty running we're relying on corefile regeneration on startup
	if s.cmd != nil && s.cmd.Process != nil && s.cmd.ProcessState == nil {
		s.args.ArgPath("Corefile", s.makeCorefile(s.args))
		if err := s.cmd.Process.Signal(syscall.SIGUSR1); err != nil {
			supervisor.Logger(ctx).Warningf("Failed to send SIGUSR1 to CoreDNS for reload: %v", err)
		}
	}
}
