package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
)

var k8sCommand = &cobra.Command{
	Short: "Manages kubernetes-specific functionality in Metropolis.",
	Use:   "k8s",
}

var k8sConfigureCommand = &cobra.Command{
	Use:   "configure",
	Short: "Configures local `kubectl` for use with a Metropolis cluster.",
	Long: `Configures a local kubectl instance (or any other Kubernetes application)
to connect to a Metropolis cluster. A cluster endpoint must be provided with the
--endpoints parameter.`,
	Args: PrintUsageOnWrongArgs(cobra.ExactArgs(0)),
	RunE: func(cmd *cobra.Command, _ []string) error {
		ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
		if len(flags.clusterEndpoints) < 1 {
			return fmt.Errorf("k8s configure requires at least one cluster endpoint to be provided with the --endpoints parameter")
		}

		contextName, err := cmd.Flags().GetString("context")
		if err != nil || contextName == "" {
			return fmt.Errorf("k8s configure requires a valid context name to be provided with the --context parameter")
		}

		// If the user has metroctl in their path, use the metroctl from path as
		// a credential plugin. Otherwise use the path to the currently-running
		// metroctl.
		metroctlPath := "metroctl"
		if _, err := exec.LookPath("metroctl"); err != nil {
			metroctlPath, err = os.Executable()
			if err != nil {
				return fmt.Errorf("failed to create kubectl entry as metroctl is neither in PATH nor can its absolute path be determined: %w", err)
			}
		}
		// TODO(q3k, issues/144): this only works as long as all nodes are kubernetes controller
		// nodes. This won't be the case for too long. Figure this out.
		if err := core.InstallKubeletConfig(ctx, metroctlPath, connectOptions(), contextName, flags.clusterEndpoints[0]); err != nil {
			return fmt.Errorf("failed to install metroctl/k8s integration: %w", err)
		}
		log.Printf("Success! kubeconfig is set up. You can now run kubectl --context=%s ... to access the Kubernetes cluster.", contextName)
		return nil
	},
}

func init() {
	k8sConfigureCommand.Flags().String("context", "metroctl", "The name for the kubernetes context to configure")
	k8sCommand.AddCommand(k8sConfigureCommand)
	rootCmd.AddCommand(k8sCommand)
}
