package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"k8s.io/klog"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/webug"
	"source.monogon.dev/cloud/lib/component"
	"source.monogon.dev/cloud/shepherd/equinix/manager"
	"source.monogon.dev/cloud/shepherd/equinix/wrapngo"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

type Config struct {
	Component component.ComponentConfig
	BMDB      bmdb.BMDB

	SharedConfig      manager.SharedConfig
	ProvisionerConfig manager.ProvisionerConfig
	InitializerConfig manager.InitializerConfig
	WebugConfig       webug.Config
	API               wrapngo.Opts
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

	c.SharedConfig.RegisterFlags("")
	c.ProvisionerConfig.RegisterFlags()
	c.InitializerConfig.RegisterFlags()
	c.WebugConfig.RegisterFlags()
	c.API.RegisterFlags()
}

func main() {
	c := &Config{}
	c.RegisterFlags()
	flag.Parse()

	ctx := clicontext.WithInterrupt(context.Background())

	if c.API.APIKey == "" || c.API.User == "" {
		klog.Exitf("-equinix_api_username and -equinix_api_key must be set")
	}
	api := wrapngo.New(&c.API)

	// These variables are _very_ important to configure correctly, otherwise someone
	// running this locally with prod creds will actually destroy production
	// data.
	if strings.Contains(c.SharedConfig.KeyLabel, "FIXME") {
		klog.Exitf("refusing to run with -equinix_ssh_key_label %q, please set it to something unique", c.SharedConfig.KeyLabel)
	}
	if strings.Contains(c.SharedConfig.DevicePrefix, "FIXME") {
		klog.Exitf("refusing to run with -equinix_device_prefix %q, please set it to something unique", c.SharedConfig.DevicePrefix)
	}

	klog.Infof("Ensuring our SSH key is configured...")
	if err := c.SharedConfig.SSHEquinixEnsure(ctx, api); err != nil {
		klog.Exitf("Ensuring SSH key failed: %v", err)
	}

	provisioner, err := c.ProvisionerConfig.New(api, &c.SharedConfig)
	if err != nil {
		klog.Exitf("%v", err)
	}

	initializer, err := manager.NewInitializer(api, c.InitializerConfig, &c.SharedConfig)
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
		if err := c.WebugConfig.Start(ctx, conn); err != nil && err != ctx.Err() {
			klog.Exitf("Failed to start webug: %v", err)
		}
	}()

	<-ctx.Done()
}
