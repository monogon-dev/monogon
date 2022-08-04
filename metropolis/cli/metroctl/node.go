package main

import (
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Short:   "Updates and queries node information.",
	Use:     "node",
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
