package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	"source.monogon.dev/metropolis/proto/api"
)

var installCmd = &cobra.Command{
	Short: "Contains subcommands to install Metropolis over different mediums.",
	Use:   "install",
}

// install flags
var installer *string
var bundle *string

var genusbCmd = &cobra.Command{
	Use:     "genusb target --installer=inst.efi --bundle=bundle.bin",
	Short:   "Generates a Metropolis installer disk or image.",
	Example: "metroctl install genusb /dev/sdx --installer=installer_x86_64.efi --bundle=metropolis_dev_x86_64.tar.xz",
	Args:    cobra.ExactArgs(1), // One positional argument: the target
	Run:     doGenUSB,
}

// If useInTreeArtifacts is true metroctl should use a bundle and installer
// directly from the build tree. It is automatically set to true if metroctl is
// running under bazel run. Specifying either one manually still overrides
// the in-tree artifacts.
var useInTreeArtifacts = os.Getenv("BUILD_WORKSPACE_DIRECTORY") != ""

var inTreeInstaller = "metropolis/node/installer/kernel.efi"
var inTreeBundle = "metropolis/node/node.zip"

// A PEM block type for a Metropolis initial owner private key
const ownerKeyType = "METROPOLIS INITIAL OWNER PRIVATE KEY"

func doGenUSB(cmd *cobra.Command, args []string) {
	if useInTreeArtifacts && *installer == "" {
		installer = &inTreeInstaller
	}
	if useInTreeArtifacts && *bundle == "" {
		bundle = &inTreeBundle
	}
	installerFile, err := os.Open(*installer)
	if err != nil {
		log.Fatalf("Failed to open installer: %v", err)
	}
	installerFileStat, err := installerFile.Stat()
	if err != nil {
		log.Fatalf("Failed to stat installer: %v", err)
	}
	var bundleFile *os.File
	var bundleFileStat os.FileInfo
	if bundle != nil && *bundle != "" {
		bundleFile, err = os.Open(*bundle)
		if err != nil {
			log.Fatalf("Failed to open bundle: %v", err)
		}
		bundleFileStat, err = bundleFile.Stat()
		if err != nil {
			log.Fatalf("Failed to stat bundle: %v", err)
		}
	}

	// TODO(lorenz): Have a key management story for this
	if err := os.MkdirAll(filepath.Join(xdg.ConfigHome, "metroctl"), 0700); err != nil {
		log.Fatalf("Failed to create config directory: %v", err)
	}
	var ownerPublicKey ed25519.PublicKey
	ownerPrivateKeyPEM, err := os.ReadFile(filepath.Join(xdg.ConfigHome, "metroctl/owner-key.pem"))
	if os.IsNotExist(err) {
		pub, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			log.Fatalf("Failed to generate owner private key: %v", err)
		}
		pemPriv := pem.EncodeToMemory(&pem.Block{Type: ownerKeyType, Bytes: priv})
		if err := os.WriteFile(filepath.Join(xdg.ConfigHome, "metroctl/owner-key.pem"), pemPriv, 0600); err != nil {
			log.Fatalf("Failed to store owner private key: %v", err)
		}
		ownerPublicKey = pub
	} else if err != nil {
		log.Fatalf("Failed to load owner private key: %v", err)
	} else {
		block, _ := pem.Decode(ownerPrivateKeyPEM)
		if block == nil {
			log.Fatalf("owner-key.pem contains invalid PEM")
		}
		if block.Type != ownerKeyType {
			log.Fatalf("owner-key.pem contains a PEM block that's not a %v", ownerKeyType)
		}
		if len(block.Bytes) != ed25519.PrivateKeySize {
			log.Fatal("owner-key.pem contains non-Ed25519 key")
		}
		ownerPrivateKey := ed25519.PrivateKey(block.Bytes)
		ownerPublicKey = ownerPrivateKey.Public().(ed25519.PublicKey)
	}

	// TODO(lorenz): This can only bootstrap right now. As soon as @serge's role
	// management has stabilized we can replace this with a proper
	// implementation.
	params := &api.NodeParameters{
		Cluster: &api.NodeParameters_ClusterBootstrap_{
			ClusterBootstrap: &api.NodeParameters_ClusterBootstrap{
				OwnerPublicKey: ownerPublicKey,
			},
		},
	}

	installerImageArgs := core.MakeInstallerImageArgs{
		TargetPath:    args[0],
		Installer:     installerFile,
		InstallerSize: uint64(installerFileStat.Size()),
		NodeParams:    params,
	}

	if bundleFile != nil {
		installerImageArgs.Bundle = bundleFile
		installerImageArgs.BundleSize = uint64(bundleFileStat.Size())
	}

	if err := core.MakeInstallerImage(installerImageArgs); err != nil {
		log.Fatalf("Failed to create installer: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.AddCommand(genusbCmd)

	bundle = installCmd.PersistentFlags().StringP("bundle", "b", "", "Metropolis bundle file to use")
	installer = installCmd.PersistentFlags().StringP("installer", "i", "", "Metropolis installer file to use")
}