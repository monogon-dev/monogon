package core

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"

	clientauthentication "k8s.io/client-go/pkg/apis/clientauthentication/v1"
	"k8s.io/client-go/tools/clientcmd"
	clientapi "k8s.io/client-go/tools/clientcmd/api"

	"source.monogon.dev/metropolis/node"
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

// InstallKubeletConfig modifies the default kubelet kubeconfig of the host
// system to be able to connect via a metroctl (and an associated ConnectOptions)
// to a Kubernetes apiserver at IP address/hostname 'server'.
//
// The kubelet's kubeconfig changes will be limited to contexts/configs/... named
// configName. The configName context will be made the default context only if
// there is no other default context in the current subconfig.
//
// Kubeconfigs can only take a single Kubernetes server address, so this function
// similarly only allows you to specify only a single server address.
func InstallKubeletConfig(metroctlPath string, opts *ConnectOptions, configName, server string) error {
	ca := clientcmd.NewDefaultPathOptions()
	config, err := ca.GetStartingConfig()
	if err != nil {
		return fmt.Errorf("getting initial config failed: %w", err)
	}

	args := []string{
		"k8scredplugin",
	}
	args = append(args, opts.ToFlags()...)

	config.AuthInfos[configName] = &clientapi.AuthInfo{
		Exec: &clientapi.ExecConfig{
			APIVersion: clientauthentication.SchemeGroupVersion.String(),
			Command:    metroctlPath,
			Args:       args,
			InstallHint: `Authenticating to Metropolis clusters requires metroctl to be present.
Running metroctl takeownership creates this entry and either points to metroctl as a command in
PATH if metroctl is in PATH at that time or to the absolute path to metroctl at that time.
If you moved metroctl afterwards or want to switch to PATH resolution, edit $HOME/.kube/config and
change users.metropolis.exec.command to the required path (or just metroctl if using PATH resolution).`,
			InteractiveMode: clientapi.NeverExecInteractiveMode,
		},
	}

	var u url.URL
	u.Scheme = "https"
	u.Host = net.JoinHostPort(server, node.KubernetesAPIWrappedPort.PortString())
	config.Clusters[configName] = &clientapi.Cluster{
		// MVP: This is insecure, but making this work would be wasted effort
		// as all of it will be replaced by the identity system.
		// TODO(issues/144): adjust cluster endpoints once have functioning roles
		// implemented.
		InsecureSkipTLSVerify: true,
		Server:                u.String(),
		ProxyURL:              opts.ProxyURL(),
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

// ConnectOptions define how to reach a Metropolis cluster from metroctl.
//
// This structure can be built directly. All unset fields mean 'default'. It can
// then be used to generate the equivalent flags to passs to metroctl.
//
// Nil pointers to ConnectOptions are equivalent to an empty ConneectOptions when
// methods on it are called.
type ConnectOptions struct {
	// ConfigPath is the path at which the metroctl configuration/credentials live.
	// If not set, the default will be used.
	ConfigPath string
	// ProxyServer is a host:port pair that indicates the metropolis cluster should
	// be reached via the given SOCKS5 proxy. If not set, the cluster can be reached
	// directly from the host networking stack.
	ProxyServer string
	// Endpoints are the IP addresses/hostnames (without port part) of the Metropolis
	// instances that metroctl should use to establish connectivity to a cluster.
	// These instances should have the ControlPlane role set.
	Endpoints []string
}

// ToFlags returns the metroctl flags corresponding to the options described by
// this ConnectionOptions struct.
func (c *ConnectOptions) ToFlags() []string {
	var res []string

	if c == nil {
		return res
	}

	if c.ConfigPath != "" {
		res = append(res, "--config", c.ConfigPath)
	}
	if c.ProxyServer != "" {
		res = append(res, "--proxy", c.ProxyServer)
	}
	for _, ep := range c.Endpoints {
		res = append(res, "--endpoints", ep)
	}

	return res
}

// ProxyURL returns a kubeconfig-compatible URL of the proxy server configured by
// ConnectOptions, or an empty string if not set.
func (c *ConnectOptions) ProxyURL() string {
	if c == nil {
		return ""
	}
	if c.ProxyServer == "" {
		return ""
	}
	var u url.URL
	u.Scheme = "socks5"
	u.Host = c.ProxyServer
	return u.String()
}