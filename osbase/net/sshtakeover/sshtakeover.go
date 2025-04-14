// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package sshtakeover provides an [ssh.Client] wrapper which provides utilities
// for taking over a machine over ssh, by uploading an executable and other
// payloads, and then executing the executable.
package sshtakeover

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	"source.monogon.dev/osbase/structfs"
)

type Client struct {
	cl       *ssh.Client
	sc       *sftp.Client
	progress func(int64)
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

type wrappedReader struct {
	r        io.Reader
	ctx      context.Context
	progress func(int64)
}

func (r *wrappedReader) Read(p []byte) (n int, err error) {
	if r.ctx.Err() != nil {
		return 0, r.ctx.Err()
	}
	n, err = r.r.Read(p)
	if r.progress != nil {
		r.progress(int64(n))
	}
	return
}

// Upload a given blob to a targetPath on the system.
func (p *Client) Upload(ctx context.Context, targetPath string, src io.Reader) error {
	src = &wrappedReader{r: src, ctx: ctx, progress: p.progress}

	df, err := p.sc.Create(targetPath)
	if err != nil {
		return fmt.Errorf("while creating file on the host: %w", err)
	}
	_, err = df.ReadFromWithConcurrency(src, 0)
	closeErr := df.Close()
	if err != nil {
		return err
	}
	return closeErr
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

func (p *Client) UploadTree(ctx context.Context, targetPath string, tree structfs.Tree) error {
	if err := p.sc.RemoveAll(targetPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("RemoveAll: %w", err)
	}
	if err := p.sc.Mkdir(targetPath); err != nil {
		return err
	}
	for nodePath, node := range tree.Walk() {
		fullPath := targetPath + "/" + nodePath
		switch {
		case node.Mode.IsDir():
			if err := p.sc.Mkdir(fullPath); err != nil {
				return fmt.Errorf("sftp mkdir %q: %w", fullPath, err)
			}
		case node.Mode.IsRegular():
			reader, err := node.Content.Open()
			if err != nil {
				return fmt.Errorf("upload %q: %w", nodePath, err)
			}
			if err := p.Upload(ctx, fullPath, reader); err != nil {
				reader.Close()
				return fmt.Errorf("upload %q: %w", fullPath, err)
			}
			reader.Close()
		default:
			return fmt.Errorf("upload %q: unsupported file type %s", nodePath, node.Mode.Type().String())
		}
	}
	return nil
}

// SetProgress sets a callback which will be called repeatedly during uploads
// with a number of bytes that have been read.
func (p *Client) SetProgress(callback func(int64)) {
	p.progress = callback
}

func (p *Client) Close() error {
	scErr := p.sc.Close()
	clErr := p.cl.Close()
	if clErr != nil {
		return clErr
	}
	return scErr
}
