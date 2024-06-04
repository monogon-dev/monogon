package cluster

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bazelbuild/rules_go/go/runfiles"

	"source.monogon.dev/metropolis/test/launch"
)

// A TPMFactory manufactures virtual TPMs using swtpm.
//
// A factory has an assigned state directory into which it will write per-factory
// data (like CA certificates and keys). Each manufactured TPM also has a state
// directory, which is first generated on manufacturing, and then passed to an
// swtpm instance.
type TPMFactory struct {
	stateDir string
}

// NewTPMFactory creates a new TPM factory at a given state path. The state path
// is a directory used to persist TPM factory data. It will be created if needed,
// and can be reused across TPM factories (but not used in parallel).
func NewTPMFactory(stateDir string) (*TPMFactory, error) {
	if err := os.MkdirAll(stateDir, 0744); err != nil {
		return nil, fmt.Errorf("could not create state directory: %w", err)
	}

	f := &TPMFactory{
		stateDir: stateDir,
	}

	if err := os.MkdirAll(f.caDir(), 0700); err != nil {
		return nil, fmt.Errorf("could not create CA state directory: %w", err)
	}
	err := writeSWTPMConfig(f.localCAConfPath(), map[string]string{
		"statedir":   f.caDir(),
		"signingkey": filepath.Join(f.caDir(), "signkey.pem"),
		"issuercert": filepath.Join(f.caDir(), "issuercert.pem"),
		"certserial": filepath.Join(f.caDir(), "certserial"),
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *TPMFactory) caDir() string {
	return filepath.Join(f.stateDir, "ca")
}

func (f *TPMFactory) localCAConfPath() string {
	return filepath.Join(f.caDir(), "swtpm-localca.conf")
}

func (f *TPMFactory) localCAOptionsPath() string {
	return filepath.Join(f.caDir(), "swtpm-localca.options")
}

func (f *TPMFactory) swtpmConfPath() string {
	return filepath.Join(f.stateDir, "swtpm.conf")
}

// writeSWTPMConfig serializes a key/value config file for swtpm tools into a
// path.
func writeSWTPMConfig(path string, data map[string]string) error {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, k := range keys {
		if _, err := fmt.Fprintf(f, "%s = %s\n", k, data[k]); err != nil {
			return err
		}
	}
	return nil
}

// A TPMPlatform defines a platform that a TPM is part of. This will usually be
// some kind of device, in this case a virtual device.
type TPMPlatform struct {
	Manufacturer string
	Version      string
	Model        string
}

// Manufacture builds a new TPM for a given platform at a path. The path points
// to a directory that will be created if it doens't exist yet, and can be passed
// to swtpm to actually emulate the created TPM.
func (f *TPMFactory) Manufacture(ctx context.Context, path string, platform *TPMPlatform) error {
	launch.Log("Starting to manufacture TPM for %s... (%+v)", path, platform)

	// Path to state file. Used to make sure Manufacture runs only once.
	permall := filepath.Join(path, "tpm2-00.permall")

	if _, err := os.Stat(permall); err == nil {
		launch.Log("Skipping manufacturing TPM for %s, already exists", path)
		return nil
	}

	// Find all tools.
	swtpm, err := runfiles.Rlocation("swtpm/swtpm")
	if err != nil {
		return fmt.Errorf("could not find swtpm: %w", err)
	}
	swtpmSetup, err := runfiles.Rlocation("swtpm/swtpm_setup")
	if err != nil {
		return fmt.Errorf("could not find swtpm_setup: %w", err)
	}
	swtpmLocalca, err := runfiles.Rlocation("swtpm/swtpm_localca")
	if err != nil {
		return fmt.Errorf("could not find swtpm_localca: %w", err)
	}
	swtpmCert, err := runfiles.Rlocation("_main/metropolis/test/swtpm/swtpm_cert/swtpm_cert_/swtpm_cert")
	if err != nil {
		return fmt.Errorf("could not find swtpm_cert: %w", err)
	}
	certtool, err := runfiles.Rlocation("_main/metropolis/test/swtpm/certtool/certtool_/certtool")
	if err != nil {
		return fmt.Errorf("could not find certtool: %w", err)
	}

	// Prepare swtpm-localca.options.
	options := []string{
		"--platform-manufacturer " + platform.Manufacturer,
		"--platform-version " + platform.Version,
		"--platform-model " + platform.Model,
		"",
	}
	err = os.WriteFile(f.localCAOptionsPath(), []byte(strings.Join(options, "\n")), 0600)
	if err != nil {
		return fmt.Errorf("could not write local options: %w", err)
	}

	// Prepare swptm.conf.
	err = writeSWTPMConfig(f.swtpmConfPath(), map[string]string{
		"create_certs_tool":         swtpmLocalca,
		"create_certs_tool_config":  f.localCAConfPath(),
		"create_certs_tool_options": f.localCAOptionsPath(),
	})
	if err != nil {
		return fmt.Errorf("could not write swtpm.conf: %w", err)
	}

	if err := os.MkdirAll(path, 0700); err != nil {
		return fmt.Errorf("could not make output path: %w", err)
	}
	cmd := exec.CommandContext(ctx, swtpmSetup,
		"--tpm", fmt.Sprintf("%s socket", swtpm),
		"--tpmstate", path,
		"--create-ek-cert",
		"--create-platform-cert",
		"--allow-signing",
		"--tpm2",
		"--display",
		"--pcr-banks", "sha1,sha256,sha384,sha512",
		"--config", f.swtpmConfPath())
	cmd.Env = append(cmd.Env, fmt.Sprintf("PATH=%s:%s", filepath.Dir(swtpmCert), filepath.Dir(certtool)))
	cmd.Env = append(cmd.Env, "MONOGON_LIBTPMS_ACKNOWLEDGE_UNSAFE=yes")
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("Manufacturing TPM for %s failed: swtm_setup: %s", path, out)
		return fmt.Errorf("swtpm_setup failed: %w", err)
	}

	if _, err := os.Stat(permall); os.IsNotExist(err) {
		log.Printf("Manufacturing TPM for %s failed: state file did not get created", path)
		return fmt.Errorf("%s did not get created during TPM manufacture", permall)
	}

	launch.Log("Successfully manufactured TPM for %s", path)
	return nil
}
