package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/equinix/wrapngo"
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
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
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
