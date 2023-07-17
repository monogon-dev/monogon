package metrics

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/metropolis/cli/pkg/datafile"
	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
	"source.monogon.dev/metropolis/test/util"
)

// TestMetricsForwarder exercises the metrics forwarding functionality of the
// metrics service. That is, it makes sure that the service starts some fake
// exporters and then forwards HTTP traffic to them.
func TestMetricsForwarder(t *testing.T) {
	path, _ := datafile.ResolveRunfile("metropolis/node/core/metrics/fake_exporter/fake_exporter_/fake_exporter")

	exporters := []Exporter{
		{
			Name:       "test1",
			Port:       node.Port(8081),
			Executable: path,
			Arguments: []string{
				"-listen", "127.0.0.1:8081",
				"-value", "100",
			},
		},
		{
			Name:       "test2",
			Port:       node.Port(8082),
			Executable: path,
			Arguments: []string{
				"-listen", "127.0.0.1:8082",
				"-value", "200",
			},
		},
	}

	eph := util.NewEphemeralClusterCredentials(t, 1)

	svc := Service{
		Credentials: eph.Nodes[0],
		Exporters:   exporters,

		enableDynamicAddr: true,
		dynamicAddr:       make(chan string),
	}

	supervisor.TestHarness(t, svc.Run)
	addr := <-svc.dynamicAddr

	pool := x509.NewCertPool()
	pool.AddCert(eph.CA)

	cl := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				ServerName: eph.Nodes[0].ID(),
				RootCAs:    pool,

				Certificates: []tls.Certificate{eph.Manager},
			},
		},
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	util.TestEventual(t, "retrieve-test1", ctx, 10*time.Second, func(ctx context.Context) error {
		url := (&url.URL{
			Scheme: "https",
			Host:   addr,
			Path:   "/metrics/test1",
		}).String()
		req, _ := http.NewRequest("GET", url, nil)
		res, err := cl.Do(req)
		if err != nil {
			return fmt.Errorf("Get(%q): %v", url, err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return fmt.Errorf("Get(%q): code %d", url, res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		want := "test 100"
		if !strings.Contains(string(body), want) {
			return util.Permanent(fmt.Errorf("did not find expected value %q in %q", want, string(body)))
		}
		return nil
	})
	util.TestEventual(t, "retrieve-test2", ctx, 10*time.Second, func(ctx context.Context) error {
		url := (&url.URL{
			Scheme: "https",
			Host:   addr,
			Path:   "/metrics/test2",
		}).String()
		req, _ := http.NewRequest("GET", url, nil)
		res, err := cl.Do(req)
		if err != nil {
			return fmt.Errorf("Get(%q): %v", url, err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return fmt.Errorf("Get(%q): code %d", url, res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		want := "test 200"
		if !strings.Contains(string(body), want) {
			return util.Permanent(fmt.Errorf("did not find expected value %q in %q", want, string(body)))
		}
		return nil
	})
}

func TestDiscovery(t *testing.T) {
	eph := util.NewEphemeralClusterCredentials(t, 1)

	curator, ccl := util.MakeTestCurator(t)
	defer ccl.Close()

	svc := Service{
		Credentials:       eph.Nodes[0],
		Curator:           apb.NewCuratorClient(ccl),
		LocalRoles:        &memory.Value[*cpb.NodeRoles]{},
		Exporters:         []Exporter{},
		enableDynamicAddr: true,
		dynamicAddr:       make(chan string),
	}

	supervisor.TestHarness(t, svc.Run)
	addr := <-svc.dynamicAddr

	pool := x509.NewCertPool()
	pool.AddCert(eph.CA)

	cl := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				ServerName: eph.Nodes[0].ID(),
				RootCAs:    pool,

				Certificates: []tls.Certificate{eph.Manager},
			},
		},
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	util.TestEventual(t, "inactive-discovery", ctx, 10*time.Second, func(ctx context.Context) error {
		url := (&url.URL{
			Scheme: "https",
			Host:   addr,
			Path:   "/discovery",
		}).String()
		req, _ := http.NewRequest("GET", url, nil)
		res, err := cl.Do(req)
		if err != nil {
			return fmt.Errorf("Get(%q): %v", url, err)
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusNotImplemented {
			return fmt.Errorf("Get(%q): code %d", url, res.StatusCode)
		}
		return nil
	})

	// First set the local roles to be a consensus member which starts a watcher,
	// create a fake node after that
	svc.LocalRoles.Set(&cpb.NodeRoles{ConsensusMember: &cpb.NodeRoles_ConsensusMember{}})
	curator.NodeWithPrefixes(wgtypes.Key{}, "metropolis-fake-1", "1.2.3.4")

	util.TestEventual(t, "active-discovery", ctx, 10*time.Second, func(ctx context.Context) error {
		url := (&url.URL{
			Scheme: "https",
			Host:   addr,
			Path:   "/discovery",
		}).String()
		req, _ := http.NewRequest("GET", url, nil)
		res, err := cl.Do(req)
		if err != nil {
			return fmt.Errorf("Get(%q): %v", url, err)
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Get(%q): code %d", url, res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		want := `[{"targets":["1.2.3.4"],"labels":{"consensus_member":"true","kubernetes_controller":"false","kubernetes_worker":"false"}}]`
		if !strings.Contains(string(body), want) {
			return util.Permanent(fmt.Errorf("did not find expected value %q in %q", want, string(body)))
		}
		return nil
	})
}
