package main

import (
	"log"
	"path/filepath"

	"github.com/adrg/xdg"
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
	// configPath overrides the default XDG config path
	configPath string
	// verbose, if set, will make this utility log additional runtime
	// information.
	verbose bool
}

var flags metroctlFlags

func init() {
	rootCmd.PersistentFlags().StringArrayVar(&flags.clusterEndpoints, "endpoints", nil, "A list of the target cluster's endpoints.")
	rootCmd.PersistentFlags().StringVar(&flags.proxyAddr, "proxy", "", "SOCKS5 proxy address")
	rootCmd.PersistentFlags().StringVar(&flags.configPath, "config", filepath.Join(xdg.ConfigHome, "metroctl"), "An alternative cluster config path")
	rootCmd.PersistentFlags().BoolVar(&flags.verbose, "verbose", false, "Log additional runtime information")
}

// rpcLogger passes through the cluster resolver logs, if "--verbose" flag was
// used.
func rpcLogger(f string, args ...interface{}) {
	if flags.verbose {
		log.Printf("resolver: " + f, args...)
	}
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
