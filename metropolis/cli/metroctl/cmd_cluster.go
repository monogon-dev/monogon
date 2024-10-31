package main

import "github.com/spf13/cobra"

var clusterCmd = &cobra.Command{
	Short: "Manages a running Metropolis cluster.",
	Use:   "cluster",
}

func init() {
	rootCmd.AddCommand(clusterCmd)
}
