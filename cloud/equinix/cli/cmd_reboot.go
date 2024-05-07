package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/equinix/wrapngo"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot [project] [id]",
	Short: "Reboots all or one specific node",
	Args:  cobra.MaximumNArgs(1),
	Run:   doReboot,
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}

func doReboot(cmd *cobra.Command, args []string) {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	api := wrapngo.New(&c)

	klog.Infof("Requesting device list...")
	devices, err := api.ListDevices(ctx, args[0])
	if err != nil {
		klog.Fatal(err)
	}

	for _, d := range devices {
		if len(args) == 2 && args[1] != d.ID {
			continue
		}

		err := api.RebootDevice(ctx, d.ID)
		if err != nil {
			klog.Error(err)
			continue
		}
		klog.Infof("rebooted %s", d.ID)
	}
}
