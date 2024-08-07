package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	_ "embed"
	"log"
	"os"
	"os/signal"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/metropolis/cli/flagdefs"
	"source.monogon.dev/metropolis/cli/metroctl/core"
	"source.monogon.dev/osbase/blkio"
	"source.monogon.dev/osbase/fat32"
)

var installCmd = &cobra.Command{
	Short: "Contains subcommands to install Metropolis via different media.",
	Use:   "install",
}

// bootstrap is a flag controlling node parameters included in the installer
// image. If set, the installed node will bootstrap a new cluster. Otherwise,
// it will try to connect to the cluster which endpoints were provided with
// the --endpoints flag.
var bootstrap = installCmd.PersistentFlags().Bool("bootstrap", false, "Create a bootstrap installer image.")
var bootstrapTPMMode = flagdefs.TPMModePflag(installCmd.PersistentFlags(), "bootstrap-tpm-mode", cpb.ClusterConfiguration_TPM_MODE_REQUIRED, "TPM mode to set on cluster")
var bootstrapStorageSecurityPolicy = flagdefs.StorageSecurityPolicyPflag(installCmd.PersistentFlags(), "bootstrap-storage-security", cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION, "Storage security policy to set on cluster")
var bundlePath = installCmd.PersistentFlags().StringP("bundle", "b", "", "Path to the Metropolis bundle to be installed")

func makeNodeParams() *api.NodeParameters {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	if err := os.MkdirAll(flags.configPath, 0700); err != nil && !os.IsExist(err) {
		log.Fatalf("Failed to create config directory: %v", err)
	}

	var params *api.NodeParameters
	if *bootstrap {
		// TODO(lorenz): Have a key management story for this
		priv, err := core.GetOrMakeOwnerKey(flags.configPath)
		if err != nil {
			log.Fatalf("Failed to generate or get owner key: %v", err)
		}
		pub := priv.Public().(ed25519.PublicKey)
		params = &api.NodeParameters{
			Cluster: &api.NodeParameters_ClusterBootstrap_{
				ClusterBootstrap: &api.NodeParameters_ClusterBootstrap{
					OwnerPublicKey: pub,
					InitialClusterConfiguration: &cpb.ClusterConfiguration{
						StorageSecurityPolicy: *bootstrapStorageSecurityPolicy,
						TpmMode:               *bootstrapTPMMode,
					},
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
	return params
}

func external(name, datafilePath string, flag *string) fat32.SizedReader {
	if flag == nil || *flag == "" {
		rPath, err := runfiles.Rlocation(datafilePath)
		if err != nil {
			log.Fatalf("No %s specified", name)
		}
		df, err := os.ReadFile(rPath)
		if err != nil {
			log.Fatalf("Cant read file: %v", err)
		}
		return bytes.NewReader(df)
	}

	f, err := blkio.NewFileReader(*flag)
	if err != nil {
		log.Fatalf("Failed to open specified %s: %v", name, err)
	}

	return f
}

func init() {
	rootCmd.AddCommand(installCmd)
}
