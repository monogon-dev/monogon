package manager

import (
	"crypto/ed25519"
	"crypto/rand"
	"flag"
	"fmt"
	"os"
	"sync"

	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"
)

type SSHKey struct {
	// myKey guards Key.
	muKey sync.Mutex

	// SSH key to use when creating machines and then connecting to them. If not
	// provided, it will be automatically loaded from KeyPersistPath, and if that
	// doesn't exist either, it will be first generated and persisted there.
	Key ed25519.PrivateKey

	// Path at which the SSH key will be loaded from and persisted to, if Key is not
	// explicitly set. Either KeyPersistPath or Key must be set.
	KeyPersistPath string
}

func (c *SSHKey) RegisterFlags() {
	flag.StringVar(&c.KeyPersistPath, "ssh_key_path", "", "Local filesystem path to read SSH key from, and save generated key to")
}

// sshKey returns the SSH key as defined by the Key and KeyPersistPath options,
// loading/generating/persisting it as necessary.
func (c *SSHKey) sshKey() (ed25519.PrivateKey, error) {
	c.muKey.Lock()
	defer c.muKey.Unlock()

	if c.Key != nil {
		return c.Key, nil
	}
	if c.KeyPersistPath == "" {
		return nil, fmt.Errorf("-ssh_key_path must be set")
	}

	data, err := os.ReadFile(c.KeyPersistPath)
	switch {
	case err == nil:
		if len(data) != ed25519.PrivateKeySize {
			return nil, fmt.Errorf("%s is not a valid ed25519 private key", c.KeyPersistPath)
		}
		c.Key = data
		klog.Infof("Loaded SSH key from %s", c.KeyPersistPath)
		return c.Key, nil
	case os.IsNotExist(err):
		if err := c.sshGenerateUnlocked(); err != nil {
			return nil, err
		}
		if err := os.WriteFile(c.KeyPersistPath, c.Key, 0400); err != nil {
			return nil, fmt.Errorf("could not persist key: %w", err)
		}
		return c.Key, nil
	default:
		return nil, fmt.Errorf("could not load peristed key: %w", err)
	}
}

// PublicKey returns the SSH public key marshaled for use, based on sshKey.
func (c *SSHKey) PublicKey() (string, error) {
	private, err := c.sshKey()
	if err != nil {
		return "", err
	}
	// Marshal the public key part in OpenSSH authorized_keys.
	sshpub, err := ssh.NewPublicKey(private.Public())
	if err != nil {
		return "", fmt.Errorf("while building SSH public key: %w", err)
	}
	return string(ssh.MarshalAuthorizedKey(sshpub)), nil
}

// Signer builds an ssh.Signer (for use in SSH connections) based on sshKey.
func (c *SSHKey) Signer() (ssh.Signer, error) {
	private, err := c.sshKey()
	if err != nil {
		return nil, err
	}
	// Set up the internal ssh.Signer to be later used to initiate SSH
	// connections with newly provided hosts.
	signer, err := ssh.NewSignerFromKey(private)
	if err != nil {
		return nil, fmt.Errorf("while building SSH signer: %w", err)
	}
	return signer, nil
}

// sshGenerateUnlocked saves a new private key into SharedConfig.Key.
func (c *SSHKey) sshGenerateUnlocked() error {
	if c.Key != nil {
		return nil
	}
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("while generating SSH key: %w", err)
	}
	c.Key = priv
	return nil
}
