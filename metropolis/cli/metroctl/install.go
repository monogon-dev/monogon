package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	_ "embed"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	"source.monogon.dev/metropolis/cli/pkg/datafile"
	"source.monogon.dev/metropolis/proto/api"
)

var installCmd = &cobra.Command{
	Short: "Contains subcommands to install Metropolis via different media.",
	Use:   "install",
}

var bundlePath = installCmd.PersistentFlags().StringP("bundle", "b", "", "Path to the Metropolis bundle to be installed")

var genusbCmd = &cobra.Command{
	Use:     "genusb target",
	Short:   "Generates a Metropolis installer disk or image.",
	Example: "metroctl install --bundle=metropolis-v0.1.zip genusb /dev/sdx",
	Args:    cobra.ExactArgs(1), // One positional argument: the target
	Run:     doGenUSB,
}

// bootstrap is a flag controlling node parameters included in the installer
// image. If set, the installed node will bootstrap a new cluster. Otherwise,
// it will try to connect to the cluster which endpoints were provided with
// the --endpoints flag.
var bootstrap bool

//go:embed metropolis/installer/kernel.efi
var installer []byte

func doGenUSB(cmd *cobra.Command, args []string) {
	var bundleReader io.Reader
	var bundleSize uint64
	if bundlePath == nil || *bundlePath == "" {
		// Attempt Bazel runfile bundle if not explicitly set
		bundle, err := datafile.Get("metropolis/node/bundle.zip")
		if err != nil {
			log.Fatalf("No bundle specified and fallback to runfiles failed: %v", err)
		}
		bundleReader = bytes.NewReader(bundle)
		bundleSize = uint64(len(bundle))
	} else {
		// Load bundle from specified path
		bundle, err := os.Open(*bundlePath)
		if err != nil {
			log.Fatalf("Failed to open specified bundle: %v", err)
		}
		bundleStat, err := bundle.Stat()
		if err != nil {
			log.Fatalf("Failed to stat specified bundle: %v", err)
		}
		bundleReader = bundle
		bundleSize = uint64(bundleStat.Size())
	}

	ctx := clicontext.WithInterrupt(context.Background())

	// TODO(lorenz): Have a key management story for this
	if err := os.MkdirAll(flags.configPath, 0700); err != nil && !os.IsExist(err) {
		log.Fatalf("Failed to create config directory: %v", err)
	}

	var params *api.NodeParameters
	if bootstrap {
		priv, err := core.GetOrMakeOwnerKey(flags.configPath)
		if err != nil {
			log.Fatalf("Failed to generate or get owner key: %v", err)
		}
		pub := priv.Public().(ed25519.PublicKey)
		params = &api.NodeParameters{
			Cluster: &api.NodeParameters_ClusterBootstrap_{
				ClusterBootstrap: &api.NodeParameters_ClusterBootstrap{
					OwnerPublicKey: pub,
				},
			},
		}
	} else {
		cc := dialAuthenticated(ctx)
		mgmt := api.NewManagementClient(cc)
		resT, err := mgmt.GetRegisterTicket(ctx, &api.GetRegisterTicketRequest{})
		if err != nil {
			log.Fatalf("While receiving register ticket: %v", err)
		}
		resI, err := mgmt.GetClusterInfo(ctx, &api.GetClusterInfoRequest{})
		if err != nil {
			log.Fatalf("While receiving cluster directory: %v", err)
		}

		params = &api.NodeParameters{
			Cluster: &api.NodeParameters_ClusterRegister_{
				ClusterRegister: &api.NodeParameters_ClusterRegister{
					RegisterTicket:   resT.Ticket,
					ClusterDirectory: resI.ClusterDirectory,
					CaCertificate:    resI.CaCertificate,
				},
			},
		}
	}

	installerImageArgs := core.MakeInstallerImageArgs{
		TargetPath:    args[0],
		Installer:     bytes.NewReader(installer),
		InstallerSize: uint64(len(installer)),
		NodeParams:    params,
		Bundle:        bundleReader,
		BundleSize:    bundleSize,
	}

	log.Printf("Generating installer image (this can take a while, see issues/92).")
	if err := core.MakeInstallerImage(installerImageArgs); err != nil {
		log.Fatalf("Failed to create installer: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)

	genusbCmd.Flags().BoolVar(&bootstrap, "bootstrap", false, "Create a bootstrap installer image.")
	installCmd.AddCommand(genusbCmd)
}
