package main

import (
	"flag"

	"github.com/spf13/cobra"

	"k8s.io/klog"

	"source.monogon.dev/cloud/shepherd/equinix/wrapngo"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if c.APIKey == "" || c.User == "" {
			klog.Exitf("-equinix_api_username and -equinix_api_key must be set")
		}
		return nil
	},
}

var c wrapngo.Opts

func init() {
	c.RegisterFlags()
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
