// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/shepherd/manager"
)

type sshConfig struct {
	User   string
	Pass   string
	SSHKey manager.SSHKey
}

func (sc *sshConfig) check() error {
	if sc.User == "" {
		return fmt.Errorf("-ssh_user must be set")
	}

	if sc.Pass == "" && sc.SSHKey.KeyPersistPath == "" {
		//TODO: The flag name -ssh_key_path could change, which would make this
		// error very confusing.
		return fmt.Errorf("-ssh_pass or -ssh_key_path must be set")
	}

	return nil
}

func (sc *sshConfig) RegisterFlags() {
	flag.StringVar(&sc.User, "ssh_user", "", "SSH username to log into the machines")
	flag.StringVar(&sc.Pass, "ssh_pass", "", "SSH password to log into the machines")
	sc.SSHKey.RegisterFlags()
}

func (sc *sshConfig) Configure(config *ssh.ClientConfig) error {
	if err := sc.check(); err != nil {
		return err
	}

	config.User = sc.User

	switch {
	case sc.Pass != "":
		config.Auth = []ssh.AuthMethod{ssh.Password(sc.Pass)}
	case sc.SSHKey.KeyPersistPath != "":
		signer, err := sc.SSHKey.Signer()
		if err != nil {
			return err
		}

		pubKey, err := sc.SSHKey.PublicKey()
		if err != nil {
			return err
		}

		klog.Infof("Using ssh key auth with public key: %s", pubKey)

		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}

	// Ignore the host key, since it's likely the first time anything logs into
	// this device, and also because there's no way of knowing its fingerprint.
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	return nil
}
