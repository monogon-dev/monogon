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
	"os"
	"os/exec"
	"strings"
	"sync"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/pkg/fileargs"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

const corefileBase = `
.:53 {
    errors
	hosts {
		fallthrough
	}
	
    cache 30
    loadbalance
    reload 10s
`

type Service struct {
	directiveRegistration chan *ExtraDirective
	directives            map[string]ExtraDirective
	cmd                   *exec.Cmd
	args                  *fileargs.FileArgs
	signalChan            chan os.Signal
	// stateMu guards access to the directives, cmd and args fields
	stateMu sync.Mutex
}

// New creates a new CoreDNS service.
// The given channel can then be used to dynamically register and unregister
// directives in the configuaration.
// To register a new directive, send an ExtraDirective on the channel. To
// remove it again, use CancelDirective() to create a removal message.
func New(directiveRegistration chan *ExtraDirective) *Service {
	return &Service{
		directives:            map[string]ExtraDirective{},
		directiveRegistration: directiveRegistration,
		signalChan:            make(chan os.Signal),
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

// Run runs the DNS service consisting of the CoreDNS process and the directive
// registration process
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
	return supervisor.RunCommand(ctx, s.cmd, supervisor.SignalChan(s.signalChan))
}

// runRegistration runs the background registration runnable which has a
// different lifecycle from the CoreDNS runnable. It is responsible for
// managing dynamic directives.
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
	s.args.ArgPath("Corefile", s.makeCorefile(s.args))
	if s.args.Error() != nil {
		supervisor.Logger(ctx).Errorf("error creating new Corefile: %v", s.args.Error())
	}
	// If the signal sending thread is not ready, do nothing. Sending signals is
	// unreliable anyways as the handler might not be installed yet or another
	// reload might be in progress. Doing it this way saves a significant amount
	// of complexity.
	select {
	case s.signalChan <- unix.SIGUSR1:
	default:
		supervisor.Logger(ctx).Infof("Reload signal could not be sent, relying on restart/reload to pick up changes")
	}
}
