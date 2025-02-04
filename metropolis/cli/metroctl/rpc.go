// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/net/proxy"
	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
)

func dialAuthenticated(ctx context.Context) (*grpc.ClientConn, error) {
	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.ErrNoCredentials) {
		return nil, fmt.Errorf("you have to take ownership of the cluster first: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get owner credentials: %w", err)
	}
	if len(flags.clusterEndpoints) == 0 {
		return nil, fmt.Errorf("please provide at least one cluster endpoint using the --endpoint parameter")
	}

	ca, err := core.GetClusterCAWithTOFU(ctx, connectOptions())
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster CA: %w", err)
	}

	tlsc := tls.Certificate{
		Certificate: [][]byte{ocert.Raw},
		PrivateKey:  opkey,
	}
	creds := rpc.NewAuthenticatedCredentials(tlsc, rpc.WantRemoteCluster(ca))
	opts, err := core.DialOpts(ctx, connectOptions())
	if err != nil {
		return nil, fmt.Errorf("while configuring dial options: %w", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	cc, err := grpc.Dial(resolver.MetropolisControlAddress, opts...)
	if err != nil {
		return nil, fmt.Errorf("while dialing cluster: %w", err)
	}
	return cc, nil
}

func dialAuthenticatedNode(ctx context.Context, id, address string, cacert *x509.Certificate) (*grpc.ClientConn, error) {
	// Collect credentials, validate command parameters, and try dialing the
	// cluster.
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.ErrNoCredentials) {
		return nil, fmt.Errorf("you have to take ownership of the cluster first: %w", err)
	}
	cc, err := core.DialNode(ctx, opkey, ocert, cacert, flags.proxyAddr, id, address)
	if err != nil {
		return nil, fmt.Errorf("while dialing node: %w", err)
	}
	return cc, nil
}

func newAuthenticatedNodeHTTPTransport(ctx context.Context, id string) (*http.Transport, error) {
	cacert, err := core.GetClusterCAWithTOFU(ctx, connectOptions())
	if err != nil {
		return nil, fmt.Errorf("could not get CA certificate: %w", err)
	}
	ocert, opkey, err := core.GetOwnerCredentials(flags.configPath)
	if errors.Is(err, core.ErrNoCredentials) {
		return nil, fmt.Errorf("you have to take ownership of the cluster first: %w", err)
	}
	tlsc := tls.Certificate{
		Certificate: [][]byte{ocert.Raw},
		PrivateKey:  opkey,
	}
	tlsconf := rpc.NewAuthenticatedTLSConfig(tlsc, rpc.WantRemoteCluster(cacert), rpc.WantRemoteNode(id))
	transport := &http.Transport{
		TLSClientConfig: tlsconf,
	}
	if flags.proxyAddr != "" {
		dialer, err := proxy.SOCKS5("tcp", flags.proxyAddr, nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("failed to create proxy dialer: %w", err)
		}
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			// TODO(q3k): handle context
			return dialer.Dial(network, addr)
		}
	}
	return transport, nil
}
