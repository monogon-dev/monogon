package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
)

func init() {
	certCmd.AddCommand(certExportCmd)

	rootCmd.AddCommand(certCmd)
}

var certCmd = &cobra.Command{
	Short: "Certificate utilities",
	Use:   "cert",
}

var certExportCmd = &cobra.Command{
	Short:   "Exports certificates for use in other programs",
	Use:     "export",
	Example: "metroctl cert export",
	Run: func(cmd *cobra.Command, args []string) {
		ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
		if errors.Is(err, core.ErrNoCredentials) {
			log.Fatalf("You have to take ownership of the cluster first: %v", err)
		}

		pkcs8Key, err := x509.MarshalPKCS8PrivateKey(opkey)
		if err != nil {
			// We explicitly pass an Ed25519 private key in, so this can't happen
			panic(err)
		}

		if err := os.WriteFile("owner.crt", pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ocert.Raw}), 0755); err != nil {
			log.Fatal(err)
		}

		if err := os.WriteFile("owner.key", pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: pkcs8Key}), 0755); err != nil {
			log.Fatal(err)
		}
		log.Println("Wrote files to current dir: cert.pem, key.pem")
	},
	Args: cobra.NoArgs,
}
