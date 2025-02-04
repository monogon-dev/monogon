// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"slices"
	"strings"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/equinix/wrapngo"
)

var listCmd = &cobra.Command{
	Use:  "list",
	Long: `This lists all hardware reservations inside a specified organization or project.`,
	Args: cobra.NoArgs,
	Run:  doList,
}

func init() {
	listCmd.Flags().String("equinix_organization", "", "from which organization to list from")
	listCmd.Flags().String("equinix_project", "", "from which project to list from")
	rootCmd.AddCommand(listCmd)
}

func doList(cmd *cobra.Command, args []string) {
	organization, err := cmd.Flags().GetString("equinix_organization")
	if err != nil {
		klog.Exitf("flag: %v", err)
	}

	project, err := cmd.Flags().GetString("equinix_project")
	if err != nil {
		klog.Exitf("flag: %v", err)
	}

	if organization == "" && project == "" {
		klog.Exitf("missing organization or project flag")
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	api := wrapngo.New(&c)

	var (
		reservations []packngo.HardwareReservation
	)
	switch {
	case project != "" && organization == "":
		klog.Infof("Listing reservations for project: %s", project)
		reservations, err = api.ListReservations(ctx, project)
	case organization != "" && project == "":
		klog.Infof("Listing reservations for organization: %s", organization)
		reservations, err = api.ListOrganizationReservations(ctx, organization)
	default:
		klog.Exitf("exactly one of organization or project flags has to be set")
	}

	if err != nil {
		klog.Fatalf("Failed to list reservations: %v", err)
	}

	type configDC struct {
		config string
		dc     string
	}
	type configDCP struct {
		configDC
		project string
	}
	mtypes := make(map[configDC]int)
	mptypes := make(map[configDCP]int)

	klog.Infof("Got %d reservations", len(reservations))
	for _, r := range reservations {
		curType := configDC{config: strings.ToLower(r.Plan.Name), dc: strings.ToLower(r.Facility.Metro.Code)}
		curPType := configDCP{curType, r.Project.Name}
		mtypes[curType]++
		mptypes[curPType]++
	}

	klog.Infof("Found the following configurations:")
	var mStrings []string
	for dc, c := range mtypes {
		mStrings = append(mStrings, fmt.Sprintf("%s | %s | %d", dc.dc, dc.config, c))
	}
	slices.Sort(mStrings)
	for _, s := range mStrings {
		klog.Info(s)
	}

	klog.Infof("Found the following configurations (per project):")
	var mpStrings []string
	for dc, c := range mptypes {
		mpStrings = append(mpStrings, fmt.Sprintf("%s | %s | %s | %d", dc.project, dc.dc, dc.config, c))
	}
	slices.Sort(mpStrings)
	for _, s := range mpStrings {
		klog.Info(s)
	}
}
