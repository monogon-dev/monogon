// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/webug"
	"source.monogon.dev/cloud/equinix/wrapngo"
	"source.monogon.dev/cloud/lib/component"
	"source.monogon.dev/cloud/shepherd/manager"
)

type Config struct {
	Component   component.ComponentConfig
	BMDB        bmdb.BMDB
	WebugConfig webug.Config

	SSHKey            manager.SSHKey
	InitializerConfig manager.InitializerConfig
	ProvisionerConfig manager.ProvisionerConfig
	RecovererConfig   manager.RecovererConfig

	API           wrapngo.Opts
	Provider      providerConfig
	UpdaterConfig UpdaterConfig
}

// TODO(q3k): factor this out to BMDB library?
func runtimeInfo() string {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "UNKNOWN"
	}
	return fmt.Sprintf("host %s", hostname)
}

func (c *Config) RegisterFlags() {
	c.Component.RegisterFlags("shepherd")
	c.BMDB.ComponentName = "shepherd-equinix"
	c.BMDB.RuntimeInfo = runtimeInfo()
	c.BMDB.Database.RegisterFlags("bmdb")
	c.WebugConfig.RegisterFlags()

	c.SSHKey.RegisterFlags()
	c.InitializerConfig.RegisterFlags()
	c.ProvisionerConfig.RegisterFlags()
	c.RecovererConfig.RegisterFlags()

	c.API.RegisterFlags()
	c.Provider.RegisterFlags()
	c.UpdaterConfig.RegisterFlags()
}

func main() {
	var c Config
	c.RegisterFlags()

	flag.Parse()
	if flag.NArg() > 0 {
		klog.Exitf("unexpected positional arguments: %v", flag.Args())
	}

	registry := c.Component.PrometheusRegistry()
	c.BMDB.EnableMetrics(registry)

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	c.Component.StartPrometheus(ctx)

	if c.API.APIKey == "" || c.API.User == "" {
		klog.Exitf("-equinix_api_username and -equinix_api_key must be set")
	}
	c.API.MetricsRegistry = registry
	api := wrapngo.New(&c.API)

	provider, err := c.Provider.New(&c.SSHKey, api)
	if err != nil {
		klog.Exitf("%v", err)
	}

	sshSigner, err := c.SSHKey.Signer()
	if err != nil {
		klog.Exitf("%v", err)
	}

	c.InitializerConfig.SSHConfig.Auth = []ssh.AuthMethod{ssh.PublicKeys(sshSigner)}
	// Equinix OS installations always use root.
	c.InitializerConfig.SSHConfig.User = "root"
	// Ignore the host key, since it's likely the first time anything logs into
	// this device, and also because there's no way of knowing its fingerprint.
	c.InitializerConfig.SSHConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	provisioner, err := manager.NewProvisioner(provider, c.ProvisionerConfig)
	if err != nil {
		klog.Exitf("%v", err)
	}

	initializer, err := manager.NewInitializer(provider, c.InitializerConfig)
	if err != nil {
		klog.Exitf("%v", err)
	}

	recoverer, err := manager.NewRecoverer(provider, c.RecovererConfig)
	if err != nil {
		klog.Exitf("%v", err)
	}

	updater, err := c.UpdaterConfig.New(api)
	if err != nil {
		klog.Exitf("%v", err)
	}

	conn, err := c.BMDB.Open(true)
	if err != nil {
		klog.Exitf("Failed to open BMDB connection: %v", err)
	}

	go func() {
		err = provisioner.Run(ctx, conn)
		if err != nil {
			klog.Exit(err)
		}
	}()
	go func() {
		err = manager.RunControlLoop(ctx, conn, initializer)
		if err != nil {
			klog.Exit(err)
		}
	}()
	go func() {
		err = manager.RunControlLoop(ctx, conn, recoverer)
		if err != nil {
			klog.Exit(err)
		}
	}()
	go func() {
		err = updater.Run(ctx, conn)
		if err != nil {
			klog.Exit(err)
		}
	}()
	go func() {
		if err := c.WebugConfig.Start(ctx, conn); err != nil && !errors.Is(err, ctx.Err()) {
			klog.Exitf("Failed to start webug: %v", err)
		}
	}()

	<-ctx.Done()
}
