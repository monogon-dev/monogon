package main

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "metroctl",
	Short: "metroctl controls Metropolis nodes and clusters.",
}

type metroctlFlags struct {
	// clusterEndpoints is a list of the targeted cluster's endpoints, used by
	// commands that perform RPC on it.
	clusterEndpoints []string
}

var flags metroctlFlags

func init() {
	rootCmd.PersistentFlags().StringArrayVar(&flags.clusterEndpoints, "endpoints", nil, "A list of the target cluster's endpoints.")
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
