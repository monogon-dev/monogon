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
	// proxyAddr is a SOCKS5 proxy address the cluster will be accessed through.
	proxyAddr string
}

var flags metroctlFlags

func init() {
	rootCmd.PersistentFlags().StringArrayVar(&flags.clusterEndpoints, "endpoints", nil, "A list of the target cluster's endpoints.")
	rootCmd.PersistentFlags().StringVar(&flags.proxyAddr, "proxy", "", "SOCKS5 proxy address")
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
