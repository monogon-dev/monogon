package main

import (
	"context"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/shepherd/equinix/wrapngo"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

var moveCmd = &cobra.Command{
	Use:   "move [source] [target]",
	Short: "Move all reserved hardware from one to another project",
	Args:  cobra.ExactArgs(2),
	Run:   doMove,
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func doMove(cmd *cobra.Command, args []string) {
	ctx := clicontext.WithInterrupt(context.Background())
	api := wrapngo.New(&c)

	klog.Infof("Listing reservations for %q", args[0])
	reservations, err := api.ListReservations(ctx, args[0])
	if err != nil {
		klog.Exitf("failed listing reservations: %v", err)
	}

	klog.Infof("Got %d reservations. Moving machines", len(reservations))
	for _, r := range reservations {
		_, err := api.MoveReservation(ctx, r.ID, args[1])
		if err != nil {
			klog.Errorf("failed moving reservation: %v", err)
			continue
		}
		klog.Infof("Moved Device %s", r.ID)
	}
}
