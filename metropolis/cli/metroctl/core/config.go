package core

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	clientauthentication "k8s.io/client-go/pkg/apis/clientauthentication/v1"
	"k8s.io/client-go/tools/clientcmd"
	clientapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	// OwnerKeyFileName is the filename of the owner key in a metroctl config
	// directory.
	OwnerKeyFileName = "owner-key.pem"
	// OwnerCertificateFileName is the filename of the owner certificate in a
	// metroctl config directory.
	OwnerCertificateFileName = "owner.pem"
)

// NoCredentialsError indicates that the requested datum (eg. owner key or owner
// certificate) is not present in the requested directory.
var NoCredentialsError = errors.New("owner certificate or key does not exist")

// A PEM block type for a Metropolis initial owner private key
const ownerKeyType = "METROPOLIS INITIAL OWNER PRIVATE KEY"

// GetOrMakeOwnerKey returns the owner key for a given metroctl configuration
// directory path, generating and saving it first if it doesn't exist.
func GetOrMakeOwnerKey(path string) (ed25519.PrivateKey, error) {
	existing, err := GetOwnerKey(path)
	switch err {
	case nil:
		return existing, nil
	case NoCredentialsError:
	default:
		return nil, err
	}

	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("when generating key: %w", err)
	}
	if err := WriteOwnerKey(path, priv); err != nil {
		return nil, err
	}
	return priv, nil
}

// WriteOwnerKey saves a given raw ED25519 private key as the owner key at a
// given metroctl configuration directory path.
func WriteOwnerKey(path string, priv ed25519.PrivateKey) error {
	pemPriv := pem.EncodeToMemory(&pem.Block{Type: ownerKeyType, Bytes: priv})
	if err := os.WriteFile(filepath.Join(path, OwnerKeyFileName), pemPriv, 0600); err != nil {
		return fmt.Errorf("when saving key: %w", err)
	}
	return nil
}

// GetOwnerKey loads and returns a raw ED25519 private key from the saved owner
// key in a given metroctl configuration directory path. If the owner key doesn't
// exist, NoCredentialsError will be returned.
func GetOwnerKey(path string) (ed25519.PrivateKey, error) {
	ownerPrivateKeyPEM, err := os.ReadFile(filepath.Join(path, OwnerKeyFileName))
	if os.IsNotExist(err) {
		return nil, NoCredentialsError
	} else if err != nil {
		return nil, fmt.Errorf("failed to load owner private key: %w", err)
	}
	block, _ := pem.Decode(ownerPrivateKeyPEM)
	if block == nil {
		return nil, errors.New("owner-key.pem contains invalid PEM armoring")
	}
	if block.Type != ownerKeyType {
		return nil, fmt.Errorf("owner-key.pem contains a PEM block that's not a %v", ownerKeyType)
	}
	if len(block.Bytes) != ed25519.PrivateKeySize {
		return nil, errors.New("owner-key.pem contains a non-Ed25519 key")
	}
	return block.Bytes, nil
}

// WriteOwnerCertificate saves a given DER-encoded X509 certificate as the owner
// key for a given metroctl configuration directory path.
func WriteOwnerCertificate(path string, cert []byte) error {
	ownerCertPEM := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	}
	if err := os.WriteFile(filepath.Join(path, OwnerCertificateFileName), pem.EncodeToMemory(&ownerCertPEM), 0644); err != nil {
		return err
	}
	return nil
}

// GetOwnerCredentials loads and returns a raw ED25519 private key alongside a
// DER-encoded X509 certificate from the saved owner key and certificate in a
// given metroctl configuration directory path. If either the key or certificate
// doesn't exist, NoCredentialsError will be returned.
func GetOwnerCredentials(path string) (cert *x509.Certificate, key ed25519.PrivateKey, err error) {
	key, err = GetOwnerKey(path)
	if err != nil {
		return nil, nil, err
	}

	ownerCertPEM, err := os.ReadFile(filepath.Join(path, OwnerCertificateFileName))
	if os.IsNotExist(err) {
		return nil, nil, NoCredentialsError
	} else if err != nil {
		return nil, nil, fmt.Errorf("failed to load owner certificate: %w", err)
	}
	block, _ := pem.Decode(ownerCertPEM)
	if block == nil {
		return nil, nil, errors.New("owner.pem contains invalid PEM armoring")
	}
	if block.Type != "CERTIFICATE" {
		return nil, nil, fmt.Errorf("owner.pem contains a PEM block that's not a CERTIFICATE")
	}
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("owner.pem contains an invalid X.509 certificate: %w", err)
	}
	return
}

// InstallK8SWrapper configures the current user's kubectl to connect to a
// Kubernetes cluster as defined by server (Metropolis wrapped APIServer
// endpoint), proxyURL (optional proxy URL) and metroctlPath (binary managing
// credentials for this cluster, and used to implement the client-side part of
// the Metropolis-wrapped APIServer protocol). The configuration will be saved to
// the 'configName' context in kubectl.
func InstallK8SWrapper(metroctlPath, configName, server, proxyURL string) error {
	ca := clientcmd.NewDefaultPathOptions()
	config, err := ca.GetStartingConfig()
	if err != nil {
		return fmt.Errorf("getting initial config failed: %w", err)
	}

	config.AuthInfos[configName] = &clientapi.AuthInfo{
		Exec: &clientapi.ExecConfig{
			APIVersion: clientauthentication.SchemeGroupVersion.String(),
			Command:    metroctlPath,
			Args:       []string{"k8scredplugin"},
			InstallHint: `Authenticating to Metropolis clusters requires metroctl to be present.
Running metroctl takeownership creates this entry and either points to metroctl as a command in
PATH if metroctl is in PATH at that time or to the absolute path to metroctl at that time.
If you moved metroctl afterwards or want to switch to PATH resolution, edit $HOME/.kube/config and
change users.metropolis.exec.command to the required path (or just metroctl if using PATH resolution).`,
			InteractiveMode: clientapi.NeverExecInteractiveMode,
		},
	}

	config.Clusters[configName] = &clientapi.Cluster{
		// MVP: This is insecure, but making this work would be wasted effort
		// as all of it will be replaced by the identity system.
		// TODO(issues/144): adjust cluster endpoints once have functioning roles
		// implemented.
		InsecureSkipTLSVerify: true,
		Server:                server,
		ProxyURL:              proxyURL,
	}

	config.Contexts[configName] = &clientapi.Context{
		AuthInfo:  configName,
		Cluster:   configName,
		Namespace: "default",
	}

	// Only set us as the current context if no other exists. Changing that
	// unprompted would be kind of rude.
	if config.CurrentContext == "" {
		config.CurrentContext = configName
	}

	if err := clientcmd.ModifyConfig(ca, *config, true); err != nil {
		return fmt.Errorf("modifying config failed: %w", err)
	}
	return nil
}
