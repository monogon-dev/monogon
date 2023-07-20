package roleserve

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/metropolis/node/core/metrics"
	kpki "source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

// workerMetrics runs the Metrics Service, which runs local Prometheus collectors
// (themselves usually instances of existing Prometheus Exporters running as
// sub-processes), and a forwarding service that lets external users access them
// over HTTPS using the Cluster CA.
type workerMetrics struct {
	curatorConnection *memory.Value[*curatorConnection]
	localRoles        *memory.Value[*cpb.NodeRoles]
	localControlplane *memory.Value[*localControlPlane]
}

func (s *workerMetrics) run(ctx context.Context) error {
	w := s.curatorConnection.Watch()
	defer w.Close()

	supervisor.Logger(ctx).Infof("Waiting for curator connection")
	cc, err := w.Get(ctx)
	if err != nil {
		return err
	}
	supervisor.Logger(ctx).Infof("Got curator connection, starting...")

	lw := s.localControlplane.Watch()
	defer lw.Close()
	cp, err := lw.Get(ctx)
	if err != nil {
		return err
	}

	pki, err := kpki.FromLocalConsensus(ctx, cp.consensus)
	if err != nil {
		return err
	}

	// TODO(q3k): move this to IssueCertificates and replace with dedicated certificate
	cert, key, err := pki.Certificate(ctx, kpki.Master)
	if err != nil {
		return fmt.Errorf("could not load certificate %q from PKI: %w", kpki.Master, err)
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return fmt.Errorf("failed to parse key for cert %q: %w", kpki.Master, err)
	}

	caCert, _, err := pki.Certificate(ctx, kpki.IdCA)
	if err != nil {
		return fmt.Errorf("could not load certificate %q from PKI: %w", kpki.IdCA, err)
	}
	parsedCACert, err := x509.ParseCertificate(caCert)
	if err != nil {
		return fmt.Errorf("failed to parse cert %q: %w", kpki.IdCA, err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AddCert(parsedCACert)

	kubeTLSConfig := &tls.Config{
		RootCAs: rootCAs,
		Certificates: []tls.Certificate{{
			Certificate: [][]byte{cert},
			PrivateKey:  parsedKey,
		}},
	}

	svc := metrics.Service{
		Credentials:   cc.credentials,
		Curator:       ipb.NewCuratorClient(cc.conn),
		LocalRoles:    s.localRoles,
		KubeTLSConfig: kubeTLSConfig,
	}
	return svc.Run(ctx)
}
