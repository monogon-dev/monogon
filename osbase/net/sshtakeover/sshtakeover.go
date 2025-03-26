// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package sshtakeover provides an [ssh.Client] wrapper which provides utilities
// for taking over a machine over ssh, by uploading an executable and other
// payloads, and then executing the executable.
package sshtakeover

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	cl *ssh.Client
	sc *sftp.Client
}

// Dial starts an ssh client connection.
func Dial(ctx context.Context, address string, config *ssh.ClientConfig) (*Client, error) {
	d := net.Dialer{
		Timeout: config.Timeout,
	}
	conn, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	conn2, chanC, reqC, err := ssh.NewClientConn(conn, address, config)
	if err != nil {
		return nil, err
	}
	cl := ssh.NewClient(conn2, chanC, reqC)

	sc, err := sftp.NewClient(cl, sftp.UseConcurrentWrites(true), sftp.MaxConcurrentRequestsPerFile(1024))
	if err != nil {
		cl.Close()
		return nil, fmt.Errorf("while building sftp client: %w", err)
	}
	return &Client{
		cl: cl,
		sc: sc,
	}, nil
}

// Execute a given command on a remote host synchronously, passing in stdin as
// input, and returning a captured stdout/stderr. The returned data might be
// valid even when err != nil, which might happen if the remote side returned a
// non-zero exit code.
func (p *Client) Execute(ctx context.Context, command string, stdin []byte) (stdout []byte, stderr []byte, err error) {
	sess, err := p.cl.NewSession()
	if err != nil {
		return nil, nil, fmt.Errorf("while creating SSH session: %w", err)
	}
	stdoutBuf := bytes.NewBuffer(nil)
	stderrBuf := bytes.NewBuffer(nil)
	sess.Stdin = bytes.NewBuffer(stdin)
	sess.Stdout = stdoutBuf
	sess.Stderr = stderrBuf
	defer sess.Close()

	if err := sess.Start(command); err != nil {
		return nil, nil, err
	}
	doneC := make(chan error, 1)
	go func() {
		doneC <- sess.Wait()
	}()
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case err := <-doneC:
		return stdoutBuf.Bytes(), stderrBuf.Bytes(), err
	}
}

// Upload a given blob to a targetPath on the system.
func (p *Client) Upload(ctx context.Context, targetPath string, src io.Reader) error {
	df, err := p.sc.Create(targetPath)
	if err != nil {
		return fmt.Errorf("while creating file on the host: %w", err)
	}

	doneC := make(chan error, 1)

	go func() {
		_, err := df.ReadFromWithConcurrency(src, 0)
		df.Close()
		doneC <- err
	}()

	select {
	case err := <-doneC:
		if err != nil {
			return fmt.Errorf("while copying file: %w", err)
		}
	case <-ctx.Done():
		df.Close()
		return ctx.Err()
	}
	return nil
}

// UploadExecutable uploads a given blob to a targetPath on the system
// and makes it executable.
func (p *Client) UploadExecutable(ctx context.Context, targetPath string, src io.Reader) error {
	if err := p.Upload(ctx, targetPath, src); err != nil {
		return err
	}
	if err := p.sc.Chmod(targetPath, 0755); err != nil {
		return fmt.Errorf("while setting file permissions: %w", err)
	}
	return nil
}

func (p *Client) Close() error {
	scErr := p.sc.Close()
	clErr := p.cl.Close()
	if clErr != nil {
		return clErr
	}
	return scErr
}
