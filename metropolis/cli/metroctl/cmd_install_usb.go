package main

import (
	_ "embed"
	"log"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
)

var genusbCmd = &cobra.Command{
	Use:     "genusb target",
	Short:   "Generates a Metropolis installer disk or image.",
	Example: "metroctl install --bundle=metropolis-v0.1.zip genusb /dev/sdx",
	Args:    PrintUsageOnWrongArgs(cobra.ExactArgs(1)), // One positional argument: the target
	Run:     doGenUSB,
}

func doGenUSB(cmd *cobra.Command, args []string) {
	params := makeNodeParams()

	installerPath, err := cmd.Flags().GetString("installer")
	if err != nil {
		log.Fatal(err)
	}

	installer := external("installer", "_main/metropolis/installer/kernel.efi", &installerPath)
	bundle := external("bundle", "_main/metropolis/node/bundle.zip", bundlePath)

	installerImageArgs := core.MakeInstallerImageArgs{
		TargetPath: args[0],
		Installer:  installer,
		NodeParams: params,
		Bundle:     bundle,
	}

	log.Printf("Generating installer image (this can take a while, see issues/92).")
	if err := core.MakeInstallerImage(installerImageArgs); err != nil {
		log.Fatalf("Failed to create installer: %v", err)
	}
}

func init() {
	genusbCmd.Flags().StringP("installer", "i", "", "Path to the Metropolis installer to use when installing")
	installCmd.AddCommand(genusbCmd)
}
