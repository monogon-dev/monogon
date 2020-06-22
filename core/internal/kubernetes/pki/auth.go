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
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"path"
	"time"

	"git.monogon.dev/source/nexantic.git/core/internal/common"

	"go.etcd.io/etcd/clientv3"
	"k8s.io/client-go/tools/clientcmd"
	configapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	EtcdPath = "/kube-pki/"
)

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

// Directly derived from Kubernetes PKI requirements documented at
// https://kubernetes.io/docs/setup/best-practices/certificates/#configure-certificates-manually
func ClientCertTemplate(identity string, groups []string) x509.Certificate {
	return x509.Certificate{
		Subject: pkix.Name{
			CommonName:   identity,
			Organization: groups,
		},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
}
func ServerCertTemplate(dnsNames []string, ips []net.IP) x509.Certificate {
	return x509.Certificate{
		Subject:     pkix.Name{},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    dnsNames,
		IPAddresses: ips,
	}
}

// Workaround for https://github.com/golang/go/issues/26676 in Go's crypto/x509. Specifically Go
// violates Section 4.2.1.2 of RFC 5280 without this.
// Fixed for 1.15 in https://go-review.googlesource.com/c/go/+/227098/.
//
// Taken from https://github.com/FiloSottile/mkcert/blob/master/cert.go#L295 written by one of Go's
// crypto engineers
func calculateSKID(pubKey crypto.PublicKey) ([]byte, error) {
	spkiASN1, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	var spki struct {
		Algorithm        pkix.AlgorithmIdentifier
		SubjectPublicKey asn1.BitString
	}
	_, err = asn1.Unmarshal(spkiASN1, &spki)
	if err != nil {
		return nil, err
	}
	skid := sha1.Sum(spki.SubjectPublicKey.Bytes)
	return skid[:], nil
}

func newCA(name string) ([]byte, ed25519.PrivateKey, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return []byte{}, privKey, fmt.Errorf("Failed to generate serial number: %w", err)
	}

	skid, err := calculateSKID(pubKey)
	if err != nil {
		return []byte{}, privKey, err
	}

	caCert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: name,
		},
		IsCA:                  true,
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              unknownNotAfter,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageOCSPSigning},
		AuthorityKeyId:        skid,
		SubjectKeyId:          skid,
	}

	caCertRaw, err := x509.CreateCertificate(rand.Reader, caCert, caCert, pubKey, privKey)
	return caCertRaw, privKey, err
}

func storeCert(consensusKV clientv3.KV, name string, cert []byte, key []byte) error {
	certPath := path.Join(EtcdPath, fmt.Sprintf("%v-cert.der", name))
	keyPath := path.Join(EtcdPath, fmt.Sprintf("%v-key.der", name))
	if _, err := consensusKV.Put(context.Background(), certPath, string(cert)); err != nil {
		return fmt.Errorf("failed to store certificate: %w", err)
	}
	if _, err := consensusKV.Put(context.Background(), keyPath, string(key)); err != nil {
		return fmt.Errorf("failed to store key: %w", err)
	}
	return nil
}

func GetCert(consensusKV clientv3.KV, name string) (cert []byte, key []byte, err error) {
	certPath := path.Join(EtcdPath, fmt.Sprintf("%v-cert.der", name))
	keyPath := path.Join(EtcdPath, fmt.Sprintf("%v-key.der", name))
	certRes, err := consensusKV.Get(context.Background(), certPath)
	if err != nil {
		err = fmt.Errorf("failed to get certificate: %w", err)
		return
	}
	keyRes, err := consensusKV.Get(context.Background(), keyPath)
	if err != nil {
		err = fmt.Errorf("failed to get certificate: %w", err)
		return
	}
	if len(certRes.Kvs) != 1 || len(keyRes.Kvs) != 1 {
		err = fmt.Errorf("failed to find certificate %v", name)
		return
	}
	cert = certRes.Kvs[0].Value
	key = keyRes.Kvs[0].Value
	return
}

func GetSingle(consensusKV clientv3.KV, name string) ([]byte, error) {
	res, err := consensusKV.Get(context.Background(), path.Join(EtcdPath, name))
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get PKI item: %w", err)
	}
	if len(res.Kvs) != 1 {
		return []byte{}, fmt.Errorf("failed to find PKI item %v", name)
	}
	return res.Kvs[0].Value, nil
}

// newCluster initializes the whole PKI for Kubernetes. It issues a single certificate per control
// plane service since it assumes that etcd is already a secure place to store data. This removes
// the need for revocation and makes the logic much simpler. Thus PKI data can NEVER be stored
// outside of etcd or other secure storage locations. All PKI data is stored in DER form and not
// PEM encoded since that would require more logic to deal with it.
func NewCluster(consensusKV clientv3.KV) error {
	// This whole issuance procedure is pretty repetitive, but abstracts badly because a lot of it
	// is subtly different.
	idCA, idKey, err := newCA("Smalltown Kubernetes ID CA")
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes ID CA: %w", err)
	}
	if err := storeCert(consensusKV, "id-ca", idCA, idKey); err != nil {
		return err
	}
	aggregationCA, aggregationKey, err := newCA("Smalltown OpenAPI Aggregation CA")
	if err != nil {
		return fmt.Errorf("failed to create OpenAPI Aggregation CA: %w", err)
	}
	if err := storeCert(consensusKV, "aggregation-ca", aggregationCA, aggregationKey); err != nil {
		return err
	}

	// ServiceAccounts don't support ed25519 yet, so use RSA (better side-channel resistance than ECDSA)
	serviceAccountPrivKeyRaw, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	serviceAccountPrivKey, err := x509.MarshalPKCS8PrivateKey(serviceAccountPrivKeyRaw)
	if err != nil {
		panic(err) // Always a programmer error
	}
	_, err = consensusKV.Put(context.Background(), path.Join(EtcdPath, "service-account-privkey.der"),
		string(serviceAccountPrivKey))
	if err != nil {
		return fmt.Errorf("failed to store service-account-privkey.der: %w", err)
	}

	apiserverCert, apiserverKey, err := IssueCertificate(
		ServerCertTemplate([]string{
			"kubernetes",
			"kubernetes.default",
			"kubernetes.default.svc",
			"kubernetes.default.svc.cluster",
			"kubernetes.default.svc.cluster.local",
			"localhost",
		}, []net.IP{{127, 0, 0, 1}}, // TODO: Add service internal IP
		),
		idCA, idKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for apiserver: %w", err)
	}
	if err := storeCert(consensusKV, "apiserver", apiserverCert, apiserverKey); err != nil {
		return err
	}

	kubeletClientCert, kubeletClientKey, err := IssueCertificate(
		ClientCertTemplate("smalltown:apiserver-kubelet-client", []string{}),
		idCA, idKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for kubelet client: %w", err)
	}
	if err := storeCert(consensusKV, "kubelet-client", kubeletClientCert, kubeletClientKey); err != nil {
		return err
	}

	frontProxyClientCert, frontProxyClientKey, err := IssueCertificate(
		ClientCertTemplate("front-proxy-client", []string{}),
		aggregationCA, aggregationKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for OpenAPI frontend: %w", err)
	}
	if err := storeCert(consensusKV, "front-proxy-client", frontProxyClientCert, frontProxyClientKey); err != nil {
		return err
	}

	controllerManagerClientCert, controllerManagerClientKey, err := IssueCertificate(
		ClientCertTemplate("system:kube-controller-manager", []string{}),
		idCA, idKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for controller-manager client: %w", err)
	}

	controllerManagerKubeconfig, err := MakeLocalKubeconfig(idCA, controllerManagerClientCert,
		controllerManagerClientKey)
	if err != nil {
		return fmt.Errorf("failed to create kubeconfig for controller-manager: %w", err)
	}

	_, err = consensusKV.Put(context.Background(), path.Join(EtcdPath, "controller-manager.kubeconfig"),
		string(controllerManagerKubeconfig))
	if err != nil {
		return fmt.Errorf("failed to store controller-manager kubeconfig: %w", err)
	}

	controllerManagerCert, controllerManagerKey, err := IssueCertificate(
		ServerCertTemplate([]string{"kube-controller-manager.local"}, []net.IP{}),
		idCA, idKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for controller-manager: %w", err)
	}
	if err := storeCert(consensusKV, "controller-manager", controllerManagerCert, controllerManagerKey); err != nil {
		return err
	}

	schedulerClientCert, schedulerClientKey, err := IssueCertificate(
		ClientCertTemplate("system:kube-scheduler", []string{}),
		idCA, idKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for scheduler client: %w", err)
	}

	schedulerKubeconfig, err := MakeLocalKubeconfig(idCA, schedulerClientCert, schedulerClientKey)
	if err != nil {
		return fmt.Errorf("failed to create kubeconfig for scheduler: %w", err)
	}

	_, err = consensusKV.Put(context.Background(), path.Join(EtcdPath, "scheduler.kubeconfig"),
		string(schedulerKubeconfig))
	if err != nil {
		return fmt.Errorf("failed to store controller-manager kubeconfig: %w", err)
	}

	schedulerCert, schedulerKey, err := IssueCertificate(
		ServerCertTemplate([]string{"kube-scheduler.local"}, []net.IP{}),
		idCA, idKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for scheduler: %w", err)
	}
	if err := storeCert(consensusKV, "scheduler", schedulerCert, schedulerKey); err != nil {
		return err
	}

	masterClientCert, masterClientKey, err := IssueCertificate(
		ClientCertTemplate("smalltown:master", []string{"system:masters"}),
		idCA, idKey,
	)
	if err != nil {
		return fmt.Errorf("failed to issue certificate for master client: %w", err)
	}

	masterClientKubeconfig, err := MakeLocalKubeconfig(idCA, masterClientCert,
		masterClientKey)
	if err != nil {
		return fmt.Errorf("failed to create kubeconfig for master client: %w", err)
	}

	_, err = consensusKV.Put(context.Background(), path.Join(EtcdPath, "master.kubeconfig"),
		string(masterClientKubeconfig))
	if err != nil {
		return fmt.Errorf("failed to store master kubeconfig: %w", err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	if err := bootstrapLocalKubelet(consensusKV, hostname); err != nil {
		return err
	}

	return nil
}

func IssueCertificate(template x509.Certificate, caCert []byte, privateKey interface{}) (cert []byte, privkey []byte, err error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("Failed to generate serial number: %w", err)
		return
	}

	caCertObj, err := x509.ParseCertificate(caCert)
	if err != nil {
		err = fmt.Errorf("failed to parse CA certificate: %w", err)
	}

	pubKey, privKeyRaw, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	privkey, err = x509.MarshalPKCS8PrivateKey(privKeyRaw)
	if err != nil {
		return
	}

	template.SerialNumber = serialNumber
	template.IsCA = false
	template.BasicConstraintsValid = true
	template.NotBefore = time.Now()
	template.NotAfter = unknownNotAfter

	cert, err = x509.CreateCertificate(rand.Reader, &template, caCertObj, pubKey, privateKey)
	return
}

func MakeLocalKubeconfig(ca, cert, key []byte) ([]byte, error) {
	kubeconfig := configapi.NewConfig()
	cluster := configapi.NewCluster()
	cluster.Server = fmt.Sprintf("https://127.0.0.1:%v", common.KubernetesAPIPort)
	cluster.CertificateAuthorityData = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ca})
	kubeconfig.Clusters["default"] = cluster
	authInfo := configapi.NewAuthInfo()
	authInfo.ClientCertificateData = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})
	authInfo.ClientKeyData = pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})
	kubeconfig.AuthInfos["default"] = authInfo
	ctx := configapi.NewContext()
	ctx.Cluster = "default"
	ctx.AuthInfo = "default"
	kubeconfig.Contexts["default"] = ctx
	kubeconfig.CurrentContext = "default"
	return clientcmd.Write(*kubeconfig)
}

func bootstrapLocalKubelet(consensusKV clientv3.KV, nodeName string) error {
	idCA, idKeyRaw, err := GetCert(consensusKV, "id-ca")
	if err != nil {
		return err
	}
	idKey := ed25519.PrivateKey(idKeyRaw)
	cert, key, err := IssueCertificate(ClientCertTemplate("system:node:"+nodeName, []string{"system:nodes"}), idCA, idKey)
	if err != nil {
		return err
	}
	kubeconfig, err := MakeLocalKubeconfig(idCA, cert, key)
	if err != nil {
		return err
	}

	serverCert, serverKey, err := IssueCertificate(ServerCertTemplate([]string{nodeName}, []net.IP{}), idCA, idKey)
	if err != nil {
		return err
	}
	if err := os.MkdirAll("/data/kubernetes", 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/kubelet.kubeconfig", kubeconfig, 0400); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/ca.crt", pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: idCA}), 0400); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/kubelet.crt", pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: serverCert}), 0400); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/data/kubernetes/kubelet.key", pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: serverKey}), 0400); err != nil {
		return err
	}

	return nil
}
