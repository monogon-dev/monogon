package core

import (
	"bufio"
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

// CertificateTOFU is an interface to different providers of a user interaction
// to confirm the validity of a CA certificate.
type CertificateTOFU interface {
	// Ask is called whenever the user needs to confirm some certificate as being the
	// CA certificate presented as the result of connection via given ConnectOptions.
	// If true is returned, the certificate is accepted and persisted as the
	// canonical CA certificate of the cluster pointed to by ConnectOptions.
	Ask(ctx context.Context, connection *ConnectOptions, cert *x509.Certificate) (bool, error)
}

// TerminalTOFU implements CertificateTOFU in an interactive way, similar to SSH.
type TerminalTOFU struct {
	// Out will be used to output prompts to the user. If not set, defaults to
	// os.Stdout.
	Out io.Writer
	// In will be used to read responses from the user. If not set, defaults to
	// os.Stdin.
	In io.Reader
}

func (i *TerminalTOFU) Ask(ctx context.Context, connection *ConnectOptions, cert *x509.Certificate) (bool, error) {
	out := i.Out
	if out == nil {
		out = os.Stdout
	}
	in := i.In
	if in == nil {
		in = os.Stdin
	}

	clusterIdentity := fmt.Sprintf("at endpoints %s", strings.Join(connection.Endpoints, ", "))
	if connection.ProxyServer != "" {
		clusterIdentity += fmt.Sprintf(" via proxy %s", connection.ProxyServer)
	}
	fmt.Fprintf(out, "The authenticity of the cluster %s can't be established.\n", clusterIdentity)

	sum := sha256.New()
	sum.Write(cert.Raw)
	fingerprint := "SHA256:" + hex.EncodeToString(sum.Sum(nil))
	fmt.Fprintf(out, "ED25519 key fingerprint of the cluster CA is %s.\n", fingerprint)

	fmt.Fprintf(out, "Are you sure you want to continue connecting (yes/no)? ")

	reader := bufio.NewReader(in)

	resC := make(chan string)
	errC := make(chan error)
	go func() {
		// This goroutine will run until we read a full newline-delimited string from the
		// input. It will leak if the context is canceled, but only until a string is
		// fully read. This is fine for now as context cancellation indicates a shutdown
		// of metroctl, but might complicate things whenever we cancel for other reasons
		// and then attempt subsequent reads from the same input.
		res, err := reader.ReadString('\n')
		if err != nil {
			errC <- err
		} else {
			resC <- res
		}
	}()

	var res string
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case err := <-errC:
		return false, err
	case res = <-resC:
	}
	res = strings.ToLower(strings.TrimSpace(res))
	return res == "yes", nil
}

// GetClusterCAWithTOFU returns the CA certificate of the cluster, performing
// trust-on-first-use (TOFU) checks per ConnectOptions first if necessary.
//
// If no locally persisted CA is found, this will connect to the cluster and
// retrieve it. Then, if now owner certificate is present, a TOFU prompt will be
// shown to the user. Otherwise, the retrieved CA will be verified against the
// local owner certificate.
//
// If the above logic accepts the CA it will be written to the configuration
// directory and used automatically on subsequent connections.
//
// An error will be returned if the user rejects the certificate as part of the
// TOFU process, if the returned CA does not matched persisted owner certificate
// (if available) or if retrieving the certificate from the cluster fails for
// some other reason.
func GetClusterCAWithTOFU(ctx context.Context, c *ConnectOptions) (*x509.Certificate, error) {
	ca, err := GetClusterCA(c.ConfigPath)
	if err == nil {
		return ca, nil
	}
	if !errors.Is(err, NoCACertificateError) {
		return nil, err
	}

	// Connect to cluster with credentials. If possible, use owner credentials.
	// Otherwise, use ephemeral credentials with owner key.
	var creds credentials.TransportCredentials

	tlsc, err := GetOwnerTLSCredentials(c.ConfigPath)

	// If we have an owner certificate, simplify TOFU by just checking the cluster CA
	// against it, and don't ask the user.
	var ocert *x509.Certificate
	if err != nil {
		if errors.Is(err, NoCredentialsError) {
			okey, err := GetOwnerKey(c.ConfigPath)
			if err != nil {
				return nil, err
			}
			creds, err = rpc.NewEphemeralCredentials(okey, rpc.WantInsecure())
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		creds = rpc.NewAuthenticatedCredentials(*tlsc, rpc.WantInsecure())
		ocert, err = x509.ParseCertificate(tlsc.Certificate[0])
		if err != nil {
			return nil, err
		}
	}

	opts, err := DialOpts(ctx, c)
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	cc, err := grpc.Dial(resolver.MetropolisControlAddress, opts...)
	if err != nil {
		return nil, fmt.Errorf("while dialing cluster to retrieve CA: %w", err)
	}
	cur := ipb.NewCuratorLocalClient(cc)
	res, err := cur.GetCACertificate(ctx, &ipb.GetCACertificateRequest{})
	if err != nil {
		return nil, fmt.Errorf("while retrieving cluster CA certificate: %w", err)
	}
	if len(res.IdentityCaCertificate) == 0 {
		return nil, fmt.Errorf("cluster returned empty CA certificate")
	}

	cert, err := x509.ParseCertificate(res.IdentityCaCertificate)
	if err != nil {
		return nil, fmt.Errorf("cluster returned invalid CA certificate: %w", err)
	}

	var okay bool
	if ocert != nil {
		// Simplified process.
		if err := ocert.CheckSignatureFrom(cert); err != nil {
			return nil, fmt.Errorf("server CA doesn't match owner certificate")
		}
		okay = true
	} else {
		// Full TOFU.
		tofu := c.TOFU
		if tofu == nil {
			tofu = &TerminalTOFU{}
		}
		okay, err = tofu.Ask(ctx, c, cert)
		if err != nil {
			return nil, err
		}
	}

	if !okay {
		return nil, fmt.Errorf("cluster CA rejected by user")
	}
	if err := WriteCACertificate(c.ConfigPath, res.IdentityCaCertificate); err != nil {
		return nil, err
	}
	return GetClusterCA(c.ConfigPath)
}
