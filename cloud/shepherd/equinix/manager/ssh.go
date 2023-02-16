package manager

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SSHClient defines a simple interface to an abstract SSH client. Usually this
// would be PlainSSHClient, but tests can use this interface to dependency-inject
// fake SSH connections.
type SSHClient interface {
	// Dial returns an SSHConnection to a given address (host:port pair) using a
	// given username/sshkey for authentication, and with a timeout for connection.
	Dial(ctx context.Context, address string, username string, sshkey ssh.Signer, connectTimeout time.Duration) (SSHConnection, error)
}

type SSHConnection interface {
	// Execute a given command on a remote host synchronously, passing in stdin as
	// input, and returning a captured stdout/stderr. The returned data might be
	// valid even when err != nil, which might happen if the remote side returned a
	// non-zero exit code.
	Execute(ctx context.Context, command string, stdin []byte) (stdout []byte, stderr []byte, err error)
	// Upload a given blob to a targetPath on the system and make executable.
	Upload(ctx context.Context, targetPath string, data []byte) error
	// Close this connection.
	Close() error
}

// PlainSSHClient implements SSHClient (and SSHConnection) using
// golang.org/x/crypto/ssh.
type PlainSSHClient struct {
}

type plainSSHConn struct {
	cl *ssh.Client
}

func (p *PlainSSHClient) Dial(ctx context.Context, address, username string, sshkey ssh.Signer, connectTimeout time.Duration) (SSHConnection, error) {
	d := net.Dialer{
		Timeout: connectTimeout,
	}
	conn, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	conf := &ssh.ClientConfig{
		// Equinix OS installations always use root.
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(sshkey),
		},
		// Ignore the host key, since it's likely the first time anything logs into
		// this device, and also because there's no way of knowing its fingerprint.
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// Timeout sets a bound on the time it takes to set up the connection, but
		// not on total session time.
		Timeout: connectTimeout,
	}
	conn2, chanC, reqC, err := ssh.NewClientConn(conn, address, conf)
	if err != nil {
		return nil, err
	}
	cl := ssh.NewClient(conn2, chanC, reqC)
	return &plainSSHConn{
		cl: cl,
	}, nil
}

func (p *plainSSHConn) Execute(ctx context.Context, command string, stdin []byte) (stdout []byte, stderr []byte, err error) {
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

func (p *plainSSHConn) Upload(ctx context.Context, targetPath string, data []byte) error {
	sc, err := sftp.NewClient(p.cl)
	if err != nil {
		return fmt.Errorf("while building sftp client: %w", err)
	}
	defer sc.Close()

	acrdr := bytes.NewReader(data)
	df, err := sc.Create(targetPath)
	if err != nil {
		return fmt.Errorf("while creating file on the host: %w", err)
	}

	doneC := make(chan error, 1)

	go func() {
		_, err := io.Copy(df, acrdr)
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

	if err := sc.Chmod(targetPath, 0755); err != nil {
		return fmt.Errorf("while setting file permissions: %w", err)
	}
	return nil
}

func (p *plainSSHConn) Close() error {
	return p.cl.Close()
}
