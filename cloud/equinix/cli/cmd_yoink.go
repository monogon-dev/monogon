package main

import (
	"bufio"
	"context"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/equinix/wrapngo"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
)

var yoinkCmd = &cobra.Command{
	Use: "yoink",
	Long: `This moves a specified amount of servers that match the given spec to a different metro.
While spec is a easy to find argument that matches the equinix system spec e.g. w3amd.75xx24c.512.8160.x86, 
metro does not represent the public facing name. Instead it is the acutal datacenter name e.g. fr2"`,
	Short: "Move a server base on the spec from one to another project",
	Args:  cobra.NoArgs,
	Run:   doYoink,
}

func init() {
	yoinkCmd.Flags().Int("count", 1, "how many machines should be moved")
	yoinkCmd.Flags().String("equinix_source_project", "", "from which project should the machine be yoinked")
	yoinkCmd.Flags().String("equinix_target_project", "", "to which project should the machine be moved")
	yoinkCmd.Flags().String("spec", "", "which device spec should be moved")
	yoinkCmd.Flags().String("metro", "", "to which metro should be moved")
	rootCmd.AddCommand(yoinkCmd)
}

func doYoink(cmd *cobra.Command, args []string) {
	srcProject, err := cmd.Flags().GetString("equinix_source_project")
	if err != nil {
		klog.Exitf("flag: %v", err)
	}

	dstProject, err := cmd.Flags().GetString("equinix_target_project")
	if err != nil {
		klog.Exitf("flag: %v", err)
	}

	if srcProject == "" || dstProject == "" {
		klog.Exitf("missing project flags")
	}

	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		klog.Exitf("flag: %v", err)
	}

	spec, err := cmd.Flags().GetString("spec")
	if err != nil {
		klog.Exitf("flag: %v", err)
	}

	if spec == "" {
		klog.Exitf("missing spec flag")
	}

	metro, err := cmd.Flags().GetString("metro")
	if err != nil {
		klog.Exitf("flag: %v", err)
	}

	if metro == "" {
		klog.Exitf("missing metro flag")
	}

	ctx := clicontext.WithInterrupt(context.Background())
	api := wrapngo.New(&c)

	klog.Infof("Listing reservations for %q", srcProject)
	reservations, err := api.ListReservations(ctx, srcProject)
	if err != nil {
		klog.Exitf("Failed to list reservations: %v", err)
	}

	type configDC struct {
		config string
		dc     string
	}
	mtypes := make(map[configDC]int)

	var matchingReservations []packngo.HardwareReservation
	reqType := configDC{config: strings.ToLower(spec), dc: strings.ToLower(metro)}

	klog.Infof("Got %d reservations", len(reservations))
	for _, r := range reservations {
		curType := configDC{config: strings.ToLower(r.Plan.Name), dc: strings.ToLower(r.Facility.Metro.Code)}

		mtypes[curType]++
		if curType == reqType {
			matchingReservations = append(matchingReservations, r)
		}
	}

	klog.Infof("Found the following configurations:")
	for dc, c := range mtypes {
		klog.Infof("%s | %s | %d", dc.dc, dc.config, c)
	}

	if len(matchingReservations) == 0 {
		klog.Exitf("Configuration not found: %s - %s", reqType.dc, reqType.config)
	}

	if len(matchingReservations)-count < 0 {
		klog.Exitf("Not enough machines with matching configuration found ")
	}

	// prefer hosts that are not deployed
	sort.Slice(matchingReservations, func(i, j int) bool {
		return matchingReservations[i].Device == nil && matchingReservations[j].Device != nil
	})

	toMove := matchingReservations[:count]
	var toDelete []string
	for _, r := range toMove {
		if r.Device != nil {
			toDelete = append(toDelete, r.Device.Hostname)
		}
	}

	stdInReader := bufio.NewReader(os.Stdin)
	klog.Infof("Will move %d machines with spec %s in %s from %s to %s.", count, spec, metro, srcProject, dstProject)
	if len(toDelete) > 0 {
		klog.Warningf("Not enough free machines found. This will delete %d provisioned hosts! Hosts scheduled for deletion: ", len(toDelete))
		klog.Warningf("%s", strings.Join(toDelete, ", "))
		klog.Warningf("Please confirm by inputting in the number of machines that will be moved.")

		read, err := stdInReader.ReadString('\n')
		if err != nil {
			klog.Exitf("failed reading input: %v", err)
		}

		atoi, err := strconv.Atoi(strings.TrimSpace(read))
		if err != nil {
			klog.Exitf("failed parsing number: %v", err)
		}

		if atoi != len(toDelete) {
			klog.Exitf("Confirmation failed! Wanted \"%q\" got \"%d\"", len(toDelete), atoi)
		} else {
			klog.Infof("Thanks for the confirmation! continuing...")
		}
	}

	klog.Infof("Note: It can be normal for a device move to fail for project validation issues. This is a known issue and can be ignored")
	for _, r := range matchingReservations[:count] {
		if r.Device != nil {
			klog.Warningf("Deleting server %s (%s) on %s", r.Device.ID, r.Device.Hostname, r.ID)

			if err := api.DeleteDevice(ctx, r.Device.ID); err != nil {
				klog.Errorf("failed deleting device %s (%s): %v", r.Device.ID, r.Device.Hostname, err)
				continue
			}
		}

		_, err := api.MoveReservation(ctx, r.ID, dstProject)
		if err != nil {
			klog.Errorf("failed moving device %s: %v", r.ID, err)
		}
	}
}
