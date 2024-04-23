package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/bmaas/bmdb/webug"
	"source.monogon.dev/cloud/lib/component"
	"source.monogon.dev/cloud/shepherd"
	"source.monogon.dev/cloud/shepherd/manager"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

type Config struct {
	Component   component.ComponentConfig
	BMDB        bmdb.BMDB
	WebugConfig webug.Config

	InitializerConfig manager.InitializerConfig
	ProvisionerConfig manager.ProvisionerConfig
	RecovererConfig   manager.RecovererConfig

	SSHConfig        sshConfig
	DeviceListSource string
	ProviderType     model.Provider
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
	c.BMDB.ComponentName = "shepherd-mini"
	c.BMDB.RuntimeInfo = runtimeInfo()
	c.BMDB.Database.RegisterFlags("bmdb")
	c.WebugConfig.RegisterFlags()

	c.InitializerConfig.RegisterFlags()
	c.ProvisionerConfig.RegisterFlags()
	c.RecovererConfig.RegisterFlags()

	c.SSHConfig.RegisterFlags()
	flag.StringVar(&c.DeviceListSource, "mini_device_list_url", "", "The url from where to fetch the device list. For local paths use file:// as scheme")
	flag.Func("mini_provider", "The provider this mini shepherd should emulate. Supported values are: lumen,equinix", func(s string) error {
		switch s {
		case strings.ToLower(string(model.ProviderEquinix)):
			c.ProviderType = model.ProviderEquinix
		case strings.ToLower(string(model.ProviderLumen)):
			c.ProviderType = model.ProviderLumen
		default:
			return fmt.Errorf("invalid provider name")
		}
		return nil
	})
}

type deviceList []machine

func (dl deviceList) asMap() map[shepherd.ProviderID]machine {
	mm := make(map[shepherd.ProviderID]machine)
	for _, m := range dl {
		mm[m.ProviderID] = m
	}
	return mm
}

func fetchDeviceList(s string) (deviceList, error) {
	var r io.Reader
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("failed parsing device list url: %v", err)
	}

	if u.Scheme != "file" {
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("invalid status code: %d != %v", http.StatusOK, resp.StatusCode)
		}
		r = resp.Body
	} else {
		f, err := os.Open(u.Path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r = f
	}

	var d deviceList
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&d); err != nil {
		return nil, err
	}

	klog.Infof("Fetched device list with %d entries", len(d))

	return d, nil
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

	ctx := clicontext.WithInterrupt(context.Background())
	c.Component.StartPrometheus(ctx)

	conn, err := c.BMDB.Open(true)
	if err != nil {
		klog.Exitf("Failed to open BMDB connection: %v", err)
	}

	sshClient, err := c.SSHConfig.NewClient()
	if err != nil {
		klog.Exitf("Failed to create SSH client: %v", err)
	}

	if c.DeviceListSource == "" {
		klog.Exitf("-mini_device_list_source must be set")
	}

	list, err := fetchDeviceList(c.DeviceListSource)
	if err != nil {
		klog.Exitf("Failed to fetch device list: %v", err)
	}

	mini := &provider{
		providerType: c.ProviderType,
		machines:     list.asMap(),
	}

	provisioner, err := manager.NewProvisioner(mini, c.ProvisionerConfig)
	if err != nil {
		klog.Exitf("%v", err)
	}

	initializer, err := manager.NewInitializer(mini, sshClient, c.InitializerConfig)
	if err != nil {
		klog.Exitf("%v", err)
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
		if err := c.WebugConfig.Start(ctx, conn); err != nil && !errors.Is(err, ctx.Err()) {
			klog.Exitf("Failed to start webug: %v", err)
		}
	}()

	<-ctx.Done()
}
