// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package launch

import (
	"context"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"path"
	"sort"

	"github.com/kballard/go-shellquote"

	metroctl "source.monogon.dev/metropolis/cli/metroctl/core"
)

type acceptall struct{}

func (a *acceptall) Ask(ctx context.Context, _ *metroctl.ConnectOptions, _ *x509.Certificate) (bool, error) {
	return true, nil
}

// ConnectOptions returns metroctl.ConnectOptions that describe connectivity to
// the launched cluster.
func (c *Cluster) ConnectOptions() *metroctl.ConnectOptions {
	// Use all metropolis nodes as endpoints. That's fine, metroctl's resolver will
	// figure out what to actually use.
	var endpoints []string
	for _, n := range c.Nodes {
		endpoints = append(endpoints, n.ManagementAddress)
	}
	sort.Strings(endpoints)
	return &metroctl.ConnectOptions{
		ConfigPath:  c.metroctlDir,
		ProxyServer: net.JoinHostPort("127.0.0.1", fmt.Sprintf("%d", c.Ports[SOCKSPort])),
		Endpoints:   endpoints,
		TOFU:        &acceptall{},
	}
}

// MetroctlFlags return stringified flags to pass to a metroctl binary to connect
// to the launched cluster.
func (c *Cluster) MetroctlFlags() string {
	return shellquote.Join(c.ConnectOptions().ToFlags()...)
}

// MakeMetroctlWrapper builds and returns the path to a shell script which calls
// metroctl (from //metropolis/cli/metroctl, which must be included as a data
// dependency of the built target) with all the required flags to connect to the
// launched cluster.
func (c *Cluster) MakeMetroctlWrapper() (string, error) {
	wpath := path.Join(c.metroctlDir, "metroctl.sh")

	// Don't create wrapper if it already exists.
	if _, err := os.Stat(wpath); err == nil {
		return wpath, nil
	}

	wrapper := fmt.Sprintf("#!/usr/bin/env bash\nexec %s %s \"$@\"", xMetroctlPath, c.MetroctlFlags())
	if err := os.WriteFile(wpath, []byte(wrapper), 0555); err != nil {
		return "", fmt.Errorf("could not write wrapper: %w", err)
	}
	return wpath, nil
}
