package main

import (
	"context"
	"crypto/x509"
	"log"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
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
	// format refers to how the output data, except logs, is formatted.
	format string
	// filter specifies a CEL filter used to narrow down the set of output
	// objects.
	filter string
	// output is an optional output file path the resulting data will be saved
	// at. If unspecified, the data will be written to stdout.
	output string
	// acceptAnyCA will persist the first encountered (while connecting) CA
	// certificate of the cluster as the trusted CA certificate for this cluster.
	// This is unsafe and should only be used for testing.
	acceptAnyCA bool
	// columns is a comma-separated list of column names which selects which columns
	// will be output to the user. An empty string means all columns will be
	// displayed.
	columns string
}

var flags metroctlFlags

func init() {
	rootCmd.PersistentFlags().StringSliceVar(&flags.clusterEndpoints, "endpoints", nil, "A list of the target cluster's endpoints.")
	rootCmd.PersistentFlags().StringVar(&flags.proxyAddr, "proxy", "", "SOCKS5 proxy address")
	rootCmd.PersistentFlags().StringVar(&flags.configPath, "config", filepath.Join(xdg.ConfigHome, "metroctl"), "An alternative cluster config path")
	rootCmd.PersistentFlags().BoolVar(&flags.verbose, "verbose", false, "Log additional runtime information")
	rootCmd.PersistentFlags().StringVar(&flags.format, "format", "plaintext", "Data output format")
	rootCmd.PersistentFlags().StringVar(&flags.filter, "filter", "", "The object filter applied to the output data")
	rootCmd.PersistentFlags().StringVar(&flags.columns, "columns", "", "Comma-separated list of column names to show. If not set, all columns will be shown")
	rootCmd.PersistentFlags().StringVarP(&flags.output, "output", "o", "", "Redirects output to the specified file")
	rootCmd.PersistentFlags().BoolVar(&flags.acceptAnyCA, "insecure-accept-and-persist-first-encountered-ca", false, "Accept the first encountered CA while connecting as the trusted CA for future metroctl connections with this config path. This is very insecure and should only be used for testing.")
}

// rpcLogger passes through the cluster resolver logs, if "--verbose" flag was
// used.
func rpcLogger(f string, args ...interface{}) {
	if flags.verbose {
		log.Printf("resolver: "+f, args...)
	}
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}

type acceptall struct{}

func (a *acceptall) Ask(ctx context.Context, _ *core.ConnectOptions, _ *x509.Certificate) (bool, error) {
	return true, nil
}

// connectOptions returns core.ConnectOptions as defined by the metroctl flags
// currently set.
func connectOptions() *core.ConnectOptions {
	var tofu core.CertificateTOFU
	if flags.acceptAnyCA {
		tofu = &acceptall{}
	}
	return &core.ConnectOptions{
		ConfigPath:     flags.configPath,
		ProxyServer:    flags.proxyAddr,
		Endpoints:      flags.clusterEndpoints,
		ResolverLogger: rpcLogger,
		TOFU:           tofu,
	}
}
