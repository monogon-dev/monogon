package ssh

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

// Client defines a simple interface to an abstract SSH client. Usually this
// would be DirectClient, but tests can use this interface to dependency-inject
// fake SSH connections.
type Client interface {
	// Dial returns an Connection to a given address (host:port pair) with
	// a timeout for connection.
	Dial(ctx context.Context, address string, connectTimeout time.Duration) (Connection, error)
}

type Connection interface {
	// Execute a given command on a remote host synchronously, passing in stdin as
	// input, and returning a captured stdout/stderr. The returned data might be
	// valid even when err != nil, which might happen if the remote side returned a
	// non-zero exit code.
	Execute(ctx context.Context, command string, stdin []byte) (stdout []byte, stderr []byte, err error)
	// Upload a given blob to a targetPath on the system and make executable.
	Upload(ctx context.Context, targetPath string, src io.Reader) error
	// Close this connection.
	Close() error
}

// DirectClient implements Client (and Connection) using
// golang.org/x/crypto/ssh.
type DirectClient struct {
	AuthMethod ssh.AuthMethod
	Username   string
}

type directConn struct {
	cl *ssh.Client
}

func (p *DirectClient) Dial(ctx context.Context, address string, connectTimeout time.Duration) (Connection, error) {
	d := net.Dialer{
		Timeout: connectTimeout,
	}
	conn, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	conf := &ssh.ClientConfig{
		User: p.Username,
		Auth: []ssh.AuthMethod{
			p.AuthMethod,
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
	return &directConn{
		cl: cl,
	}, nil
}

func (p *directConn) Execute(ctx context.Context, command string, stdin []byte) (stdout []byte, stderr []byte, err error) {
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

func (p *directConn) Upload(ctx context.Context, targetPath string, src io.Reader) error {
	sc, err := sftp.NewClient(p.cl, sftp.UseConcurrentWrites(true), sftp.MaxConcurrentRequestsPerFile(1024))
	if err != nil {
		return fmt.Errorf("while building sftp client: %w", err)
	}
	defer sc.Close()

	df, err := sc.Create(targetPath)
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

	if err := sc.Chmod(targetPath, 0755); err != nil {
		return fmt.Errorf("while setting file permissions: %w", err)
	}
	return nil
}

func (p *directConn) Close() error {
	return p.cl.Close()
}
