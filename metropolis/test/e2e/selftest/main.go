package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// test1InClusterKubernetes exercises connectivity to the cluster-local
// Kubernetes API server. It expects to be able to connect to the APIserver using
// the ServiceAccount and cluster CA injected by the Kubelet.
//
// The entire functionality is reimplemented without relying on Kubernetes
// client code to make the expected behaviour clear.
func test1InClusterKubernetes(ctx context.Context) error {
	token, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return fmt.Errorf("failed to read serviceaccount token: %w", err)
	}

	cacert, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/ca.crt")
	if err != nil {
		return fmt.Errorf("failed to read cluster CA certificate: %w", err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(cacert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: pool,
			},
		},
	}

	req, err := http.NewRequestWithContext(ctx, "GET", "https://kubernetes.default.svc.cluster.local/api", nil)
	if err != nil {
		return fmt.Errorf("creating request failed: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+string(token))

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	j := struct {
		Kind    string `json:"kind"`
		Message string `json:"message"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
		return fmt.Errorf("json parse error: %w", err)
	}

	if j.Kind == "Status" {
		return fmt.Errorf("API server responded with error: %q", j.Message)
	}
	if j.Kind != "APIVersions" {
		return fmt.Errorf("unexpected response from server (kind: %q)", j.Kind)
	}

	return nil
}

func main() {
	log.Printf("Metropolis Kubernetes self-test starting...")
	ctx, ctxC := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxC()

	log.Printf("1. In-cluster Kubernetes client...")
	if err := test1InClusterKubernetes(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	log.Printf("All tests passed.")
}
