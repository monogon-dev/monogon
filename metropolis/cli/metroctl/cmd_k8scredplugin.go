package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientauthentication "k8s.io/client-go/pkg/apis/clientauthentication/v1"

	"source.monogon.dev/metropolis/cli/metroctl/core"
)

var k8scredpluginCmd = &cobra.Command{
	Use:   "k8scredplugin",
	Short: "Kubernetes client-go credential plugin [internal use]",
	Long: `This implements a Kubernetes client-go credential plugin to
authenticate client-go based callers including kubectl against a Metropolis
cluster. This should never be directly called by end users.`,
	Args:   PrintUsageOnWrongArgs(cobra.ExactArgs(0)),
	Hidden: true,
	Run:    doK8sCredPlugin,
}

func doK8sCredPlugin(cmd *cobra.Command, args []string) {
	cert, key, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.ErrNoCredentials) {
		log.Fatal("No credentials found on your machine")
	}
	if err != nil {
		log.Fatalf("failed to get Metropolis credentials: %v", err)
	}

	pkcs8Key, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		// We explicitly pass an Ed25519 private key in, so this can't happen
		panic(err)
	}

	cred := clientauthentication.ExecCredential{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clientauthentication.SchemeGroupVersion.String(),
			Kind:       "ExecCredential",
		},
		Status: &clientauthentication.ExecCredentialStatus{
			ClientCertificateData: string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})),
			ClientKeyData:         string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: pkcs8Key})),
		},
	}
	if err := json.NewEncoder(os.Stdout).Encode(cred); err != nil {
		log.Fatalf("failed to encode ExecCredential: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(k8scredpluginCmd)
}
