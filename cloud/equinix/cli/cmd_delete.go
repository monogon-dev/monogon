package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/equinix/wrapngo"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [target]",
	Short: "Delete all devices from one project",
	Args:  cobra.ExactArgs(1),
	Run:   doDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func doDelete(cmd *cobra.Command, args []string) {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	api := wrapngo.New(&c)

	klog.Infof("Listing devices for %q", args[0])

	devices, err := api.ListDevices(ctx, args[0])
	if err != nil {
		klog.Exitf("failed listing devices: %v", err)
	}

	if len(devices) == 0 {
		klog.Infof("No devices found in %s", args[0])
		return
	}

	klog.Infof("Deleting %d Devices in %s. THIS WILL DELETE SERVERS! You have five seconds to cancel!", len(devices), args[0])
	time.Sleep(5 * time.Second)

	for _, d := range devices {
		h := "deleted-" + d.Hostname
		_, err := api.UpdateDevice(ctx, d.ID, &packngo.DeviceUpdateRequest{
			Hostname: &h,
		})
		if err != nil {
			klog.Infof("failed updating device %s (%s): %v", d.ID, d.Hostname, err)
			continue
		}

		klog.Infof("deleting %s (%s)...", d.ID, d.Hostname)
		if err := api.DeleteDevice(ctx, d.ID); err != nil {
			klog.Infof("failed deleting device %s (%s): %v", d.ID, d.Hostname, err)
			continue
		}
	}
}
