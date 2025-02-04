// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"

	xssh "golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/shepherd/manager"
	"source.monogon.dev/go/net/ssh"
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

func (sc *sshConfig) NewClient() (*ssh.DirectClient, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}

	c := ssh.DirectClient{
		Username: sc.User,
	}

	switch {
	case sc.Pass != "":
		c.AuthMethods = []xssh.AuthMethod{xssh.Password(sc.Pass)}
	case sc.SSHKey.KeyPersistPath != "":
		signer, err := sc.SSHKey.Signer()
		if err != nil {
			return nil, err
		}

		pubKey, err := sc.SSHKey.PublicKey()
		if err != nil {
			return nil, err
		}

		klog.Infof("Using ssh key auth with public key: %s", pubKey)

		c.AuthMethods = []xssh.AuthMethod{xssh.PublicKeys(signer)}
	}
	return &c, nil
}
