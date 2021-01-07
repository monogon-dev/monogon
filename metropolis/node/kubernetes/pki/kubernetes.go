// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pki

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"

	"go.etcd.io/etcd/clientv3"
	"k8s.io/client-go/tools/clientcmd"
	configapi "k8s.io/client-go/tools/clientcmd/api"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/pkg/logtree"
)

// KubeCertificateName is an enum-like unique name of a static Kubernetes certificate. The value of the name is used
// as the unique part of an etcd path where the certificate and key are stored.
type KubeCertificateName string

const (
	// The main Kubernetes CA, used to authenticate API consumers, and servers.
	IdCA KubeCertificateName = "id-ca"

	// Kubernetes apiserver server certificate.
	APIServer KubeCertificateName = "apiserver"

	// Kubelet client certificate, used to authenticate to the apiserver.
	KubeletClient KubeCertificateName = "kubelet-client"

	// Kubernetes Controller manager client certificate, used to authenticate to the apiserver.
	ControllerManagerClient KubeCertificateName = "controller-manager-client"
	// Kubernetes Controller manager server certificate, used to run its HTTP server.
	ControllerManager KubeCertificateName = "controller-manager"

	// Kubernetes Scheduler client certificate, used to authenticate to the apiserver.
	SchedulerClient KubeCertificateName = "scheduler-client"
	// Kubernetes scheduler server certificate, used to run its HTTP server.
	Scheduler KubeCertificateName = "scheduler"

	// Root-on-kube (system:masters) client certificate. Used to control the apiserver (and resources) by Metropolis
	// internally.
	Master KubeCertificateName = "master"

	// OpenAPI Kubernetes Aggregation CA.
	// See: https://kubernetes.io/docs/tasks/extend-kubernetes/configure-aggregation-layer/#ca-reusage-and-conflicts
	AggregationCA    KubeCertificateName = "aggregation-ca"
	FrontProxyClient KubeCertificateName = "front-proxy-client"
)

const (
	// serviceAccountKeyName is the etcd path part that is used to store the ServiceAccount authentication secret.
	// This is not a certificate, just an RSA key.
	serviceAccountKeyName = "service-account-privkey"
)

// KubernetesPKI manages all PKI resources required to run Kubernetes on Metropolis. It contains all static certificates,
// which can be retrieved, or be used to generate Kubeconfigs from.
type KubernetesPKI struct {
	logger       logtree.LeveledLogger
	KV           clientv3.KV
	Certificates map[KubeCertificateName]*Certificate
}

func NewKubernetes(l logtree.LeveledLogger, kv clientv3.KV) *KubernetesPKI {
	pki := KubernetesPKI{
		logger:       l,
		KV:           kv,
		Certificates: make(map[KubeCertificateName]*Certificate),
	}

	make := func(i, name KubeCertificateName, template x509.Certificate) {
		pki.Certificates[name] = New(pki.Certificates[i], string(name), template)
	}

	pki.Certificates[IdCA] = New(SelfSigned, string(IdCA), CA("Metropolis Kubernetes ID CA"))
	make(IdCA, APIServer, Server(
		[]string{
			"kubernetes",
			"kubernetes.default",
			"kubernetes.default.svc",
			"kubernetes.default.svc.cluster",
			"kubernetes.default.svc.cluster.local",
			"localhost",
		},
		[]net.IP{{10, 0, 255, 1}, {127, 0, 0, 1}}, // TODO(q3k): add service network internal apiserver address
	))
	make(IdCA, KubeletClient, Client("metropolis:apiserver-kubelet-client", nil))
	make(IdCA, ControllerManagerClient, Client("system:kube-controller-manager", nil))
	make(IdCA, ControllerManager, Server([]string{"kube-controller-manager.local"}, nil))
	make(IdCA, SchedulerClient, Client("system:kube-scheduler", nil))
	make(IdCA, Scheduler, Server([]string{"kube-scheduler.local"}, nil))
	make(IdCA, Master, Client("metropolis:master", []string{"system:masters"}))

	pki.Certificates[AggregationCA] = New(SelfSigned, string(AggregationCA), CA("Metropolis OpenAPI Aggregation CA"))
	make(AggregationCA, FrontProxyClient, Client("front-proxy-client", nil))

	return &pki
}

// EnsureAll ensures that all static certificates (and the serviceaccount key) are present on etcd.
func (k *KubernetesPKI) EnsureAll(ctx context.Context) error {
	for n, v := range k.Certificates {
		k.logger.Infof("Ensuring %s exists", string(n))
		_, _, err := v.Ensure(ctx, k.KV)
		if err != nil {
			return fmt.Errorf("could not ensure certificate %q exists: %w", n, err)
		}
	}
	_, err := k.ServiceAccountKey(ctx)
	if err != nil {
		return fmt.Errorf("could not ensure service account key exists: %w", err)
	}
	return nil
}

// Kubeconfig generates a kubeconfig blob for a given certificate name. The same lifetime semantics as in .Certificate
// apply.
func (k *KubernetesPKI) Kubeconfig(ctx context.Context, name KubeCertificateName) ([]byte, error) {
	c, ok := k.Certificates[name]
	if !ok {
		return nil, fmt.Errorf("no certificate %q", name)
	}
	return c.Kubeconfig(ctx, k.KV)
}

// Certificate retrieves an x509 DER-encoded (but not PEM-wrapped) key and certificate for a given certificate name.
// If the requested certificate is volatile, it will be created on demand. Otherwise it will be created on etcd (if not
// present), and retrieved from there.
func (k *KubernetesPKI) Certificate(ctx context.Context, name KubeCertificateName) (cert, key []byte, err error) {
	c, ok := k.Certificates[name]
	if !ok {
		return nil, nil, fmt.Errorf("no certificate %q", name)
	}
	return c.Ensure(ctx, k.KV)
}

// Kubeconfig generates a kubeconfig blob for this certificate. The same lifetime semantics as in .Ensure apply.
func (c *Certificate) Kubeconfig(ctx context.Context, kv clientv3.KV) ([]byte, error) {

	cert, key, err := c.Ensure(ctx, kv)
	if err != nil {
		return nil, fmt.Errorf("could not ensure certificate exists: %w", err)
	}

	kubeconfig := configapi.NewConfig()

	cluster := configapi.NewCluster()
	cluster.Server = fmt.Sprintf("https://127.0.0.1:%v", common.KubernetesAPIPort)

	ca, err := c.issuer.CACertificate(ctx, kv)
	if err != nil {
		return nil, fmt.Errorf("could not get CA certificate: %w", err)
	}
	if ca != nil {
		cluster.CertificateAuthorityData = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ca})
	}
	kubeconfig.Clusters["default"] = cluster

	authInfo := configapi.NewAuthInfo()
	authInfo.ClientCertificateData = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})
	authInfo.ClientKeyData = pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})
	kubeconfig.AuthInfos["default"] = authInfo

	ct := configapi.NewContext()
	ct.Cluster = "default"
	ct.AuthInfo = "default"
	kubeconfig.Contexts["default"] = ct

	kubeconfig.CurrentContext = "default"
	return clientcmd.Write(*kubeconfig)
}

// ServiceAccountKey retrieves (and possible generates and stores on etcd) the Kubernetes service account key. The
// returned data is ready to be used by Kubernetes components (in PKIX form).
func (k *KubernetesPKI) ServiceAccountKey(ctx context.Context) ([]byte, error) {
	// TODO(q3k): this should be abstracted away once we abstract away etcd access into a library with try-or-create
	// semantics.

	path := etcdPath("%s.der", serviceAccountKeyName)

	// Try loading  key from etcd.
	keyRes, err := k.KV.Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to get key from etcd: %w", err)
	}

	if len(keyRes.Kvs) == 1 {
		// Certificate and key exists in etcd, return that.
		return keyRes.Kvs[0].Value, nil
	}

	// No key found - generate one.
	keyRaw, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	key, err := x509.MarshalPKCS8PrivateKey(keyRaw)
	if err != nil {
		panic(err) // Always a programmer error
	}

	// Save to etcd.
	_, err = k.KV.Put(ctx, path, string(key))
	if err != nil {
		err = fmt.Errorf("failed to write newly generated key: %w", err)
	}
	return key, nil
}
