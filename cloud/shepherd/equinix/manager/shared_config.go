package manager

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/packethost/packngo"
	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"

	ecl "source.monogon.dev/cloud/shepherd/equinix/wrapngo"
)

var (
	NoSuchKey = errors.New("no such key")
)

// SharedConfig contains configuration options used by both the Initializer and
// Provisioner components of the Shepherd. In CLI scenarios, RegisterFlags should
// be called to configure this struct from CLI flags. Otherwise, this structure
// should be explicitly configured, as the default values are not valid.
type SharedConfig struct {
	// ProjectId is the Equinix project UUID used by the manager. See Equinix API
	// documentation for details. Must be set.
	ProjectId string

	// Label specifies the ID to use when handling the Equinix-registered SSH key
	// used to authenticate to newly created servers. Must be set.
	KeyLabel string

	// myKey guards Key.
	muKey sync.Mutex

	// SSH key to use when creating machines and then connecting to them. If not
	// provided, it will be automatically loaded from KeyPersistPath, and if that
	// doesn't exist either, it will be first generated and persisted there.
	Key ed25519.PrivateKey

	// Path at which the SSH key will be loaded from and persisted to, if Key is not
	// explicitly set. Either KeyPersistPath or Key must be set.
	KeyPersistPath string

	// Prefix applied to all devices (machines) created by the Provisioner, and used
	// by the Provisioner to identify machines which it managed. Must be set.
	DevicePrefix string

	// configPrefix will be set to the prefix of the latest RegisterFlags call and
	// will be then used by various methods to display the full name of a
	// misconfigured flag.
	configPrefix string
}

func (c *SharedConfig) check() error {
	if c.ProjectId == "" {
		return fmt.Errorf("-%sequinix_project_id must be set", c.configPrefix)
	}
	if c.KeyLabel == "" {
		return fmt.Errorf("-%sequinix_ssh_key_label must be set", c.configPrefix)
	}
	if c.DevicePrefix == "" {
		return fmt.Errorf("-%sequinix_device_prefix must be set", c.configPrefix)
	}
	return nil
}

func (k *SharedConfig) RegisterFlags(prefix string) {
	k.configPrefix = prefix

	flag.StringVar(&k.ProjectId, prefix+"equinix_project_id", "", "Equinix project ID where resources will be managed")
	flag.StringVar(&k.KeyLabel, prefix+"equinix_ssh_key_label", "shepherd-FIXME", "Label used to identify managed SSH key in Equinix project")
	flag.StringVar(&k.KeyPersistPath, prefix+"ssh_key_path", "shepherd-key.priv", "Local filesystem path to read SSH key from, and save generated key to")
	flag.StringVar(&k.DevicePrefix, prefix+"equinix_device_prefix", "shepherd-FIXME-", "Prefix applied to all devices (machines) in Equinix project, used to identify managed machines")
}

// sshKey returns the SSH key as defined by the Key and KeyPersistPath options,
// loading/generating/persisting it as necessary.
func (c *SharedConfig) sshKey() (ed25519.PrivateKey, error) {
	c.muKey.Lock()
	defer c.muKey.Unlock()

	if c.Key != nil {
		return c.Key, nil
	}
	if c.KeyPersistPath == "" {
		return nil, fmt.Errorf("-%sequinix_ssh_key_path must be set", c.configPrefix)
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

// sshPub returns the SSH public key marshaled for use, based on sshKey.
func (c *SharedConfig) sshPub() (string, error) {
	private, err := c.sshKey()
	if err != nil {
		return "", err
	}
	// Marshal the public key part in OpenSSH authorized_keys format that will be
	// registered with Equinix Metal.
	sshpub, err := ssh.NewPublicKey(private.Public())
	if err != nil {
		return "", fmt.Errorf("while building SSH public key: %w", err)
	}
	return string(ssh.MarshalAuthorizedKey(sshpub)), nil
}

// sshSigner builds an ssh.Signer (for use in SSH connections) based on sshKey.
func (c *SharedConfig) sshSigner() (ssh.Signer, error) {
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
func (c *SharedConfig) sshGenerateUnlocked() error {
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

// sshEquinixGet looks up the Equinix key matching SharedConfig.KeyLabel,
// returning its packngo.SSHKey instance.
func (c *SharedConfig) sshEquinix(ctx context.Context, cl ecl.Client) (*packngo.SSHKey, error) {
	ks, err := cl.ListSSHKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("while listing SSH keys: %w", err)
	}

	for _, k := range ks {
		if k.Label == c.KeyLabel {
			return &k, nil
		}
	}
	return nil, NoSuchKey
}

// sshEquinixId looks up the Equinix key identified by SharedConfig.KeyLabel,
// returning its Equinix-assigned UUID.
func (c *SharedConfig) sshEquinixId(ctx context.Context, cl ecl.Client) (string, error) {
	k, err := c.sshEquinix(ctx, cl)
	if err != nil {
		return "", err
	}
	return k.ID, nil
}

// sshEquinixUpdate makes sure the existing SSH key registered with Equinix
// matches the one from sshPub.
func (c *SharedConfig) sshEquinixUpdate(ctx context.Context, cl ecl.Client, kid string) error {
	pub, err := c.sshPub()
	if err != nil {
		return err
	}
	_, err = cl.UpdateSSHKey(ctx, kid, &packngo.SSHKeyUpdateRequest{
		Key: &pub,
	})
	if err != nil {
		return fmt.Errorf("while updating the SSH key: %w", err)
	}
	return nil
}

// sshEquinixUpload registers a new SSH key from sshPub.
func (c *SharedConfig) sshEquinixUpload(ctx context.Context, cl ecl.Client) error {
	pub, err := c.sshPub()
	if err != nil {
		return fmt.Errorf("while generating public key: %w", err)
	}
	_, err = cl.CreateSSHKey(ctx, &packngo.SSHKeyCreateRequest{
		Label:     c.KeyLabel,
		Key:       pub,
		ProjectID: c.ProjectId,
	})
	if err != nil {
		return fmt.Errorf("while creating an SSH key: %w", err)
	}
	return nil
}

// SSHEquinixEnsure initializes the locally managed SSH key (from a persistence
// path or explicitly set key) and updates or uploads it to Equinix. The key is
// generated as needed The key is generated as needed
func (c *SharedConfig) SSHEquinixEnsure(ctx context.Context, cl ecl.Client) error {
	k, err := c.sshEquinix(ctx, cl)
	switch err {
	case NoSuchKey:
		if err := c.sshEquinixUpload(ctx, cl); err != nil {
			return fmt.Errorf("while uploading key: %w", err)
		}
		return nil
	case nil:
		if err := c.sshEquinixUpdate(ctx, cl, k.ID); err != nil {
			return fmt.Errorf("while updating key: %w", err)
		}
		return nil
	default:
		return err
	}
}

// managedDevices provides a map of device provider IDs to matching
// packngo.Device instances. It calls Equinix API's ListDevices. The returned
// devices are filtered according to DevicePrefix provided through Opts. The
// returned error value, if not nil, will originate in wrapngo.
func (c *SharedConfig) managedDevices(ctx context.Context, cl ecl.Client) (map[string]packngo.Device, error) {
	ds, err := cl.ListDevices(ctx, c.ProjectId)
	if err != nil {
		return nil, err
	}
	dm := map[string]packngo.Device{}
	for _, d := range ds {
		if strings.HasPrefix(d.Hostname, c.DevicePrefix) {
			dm[d.ID] = d
		}
	}
	return dm, nil
}
