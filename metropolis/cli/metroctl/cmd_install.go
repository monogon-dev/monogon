package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	_ "embed"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	"source.monogon.dev/metropolis/pkg/blkio"
	"source.monogon.dev/metropolis/pkg/fat32"
)

var installCmd = &cobra.Command{
	Short: "Contains subcommands to install Metropolis via different media.",
	Use:   "install",
}

var bundlePath = installCmd.PersistentFlags().StringP("bundle", "b", "", "Path to the Metropolis bundle to be installed")
var installerPath = installCmd.PersistentFlags().StringP("installer", "i", "", "Path to the Metropolis installer to use when installing")

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

var bootstrapTPMMode string
var bootstrapStorageSecurityPolicy string

type externalFile struct {
	reader io.Reader
	size   uint64
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

	f, err := blkio.NewFileReader(*bundlePath)
	if err != nil {
		log.Fatalf("Failed to open specified %s: %v", name, err)
	}

	return f
}

func doGenUSB(cmd *cobra.Command, args []string) {
	var tpmMode cpb.ClusterConfiguration_TPMMode
	switch strings.ToLower(bootstrapTPMMode) {
	case "required", "require":
		tpmMode = cpb.ClusterConfiguration_TPM_MODE_REQUIRED
	case "best-effort", "besteffort":
		tpmMode = cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT
	case "disabled", "disable":
		tpmMode = cpb.ClusterConfiguration_TPM_MODE_DISABLED
	default:
		log.Fatalf("Invalid --bootstrap-tpm-mode (must be one of: required, best-effort, disabled)")
	}

	var bootstrapStorageSecurity cpb.ClusterConfiguration_StorageSecurityPolicy
	switch strings.ToLower(bootstrapStorageSecurityPolicy) {
	case "permissive":
		bootstrapStorageSecurity = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE
	case "needs-encryption":
		bootstrapStorageSecurity = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION
	case "needs-encryption-and-authentication":
		bootstrapStorageSecurity = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION
	case "needs-insecure":
		bootstrapStorageSecurity = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE
	default:

		log.Fatalf("Invalid --bootstrap-storage-security (must be one of: permissive, needs-encryption, needs-encryption-and-authentication, needs-insecure)")
	}

	bundle := external("bundle", "metropolis/node/bundle.zip", bundlePath)
	installer := external("installer", "metropolis/installer/kernel.efi", installerPath)
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
					InitialClusterConfiguration: &cpb.ClusterConfiguration{
						StorageSecurityPolicy: bootstrapStorageSecurity,
						TpmMode:               tpmMode,
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

	installerImageArgs := core.MakeInstallerImageArgs{
		TargetPath: args[0],
		Installer:  installer,
		NodeParams: params,
		Bundle:     bundle,
	}

	log.Printf("Generating installer image (this can take a while, see issues/92).")
	if err := core.MakeInstallerImage(installerImageArgs); err != nil {
		log.Fatalf("Failed to create installer: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)

	genusbCmd.Flags().BoolVar(&bootstrap, "bootstrap", false, "Create a bootstrap installer image.")
	genusbCmd.Flags().StringVar(&bootstrapTPMMode, "bootstrap-tpm-mode", "required", "TPM mode to set on cluster (required, best-effort, disabled)")
	genusbCmd.Flags().StringVar(&bootstrapStorageSecurityPolicy, "bootstrap-storage-security", "needs-encryption-and-authentication", "Storage security policy to set on cluster (permissive, needs-encryption, needs-encryption-and-authentication, needs-insecure)")
	installCmd.AddCommand(genusbCmd)
}
