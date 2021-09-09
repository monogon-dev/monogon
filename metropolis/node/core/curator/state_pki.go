package curator

import (
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/pkg/pki"
)

var (
	// pkiNamespace is the etcd/pki namespace in which the Metropolis cluster CA
	// data will live.
	pkiNamespace = pki.Namespaced("/cluster-pki/")
	// pkiCA is the main cluster CA, stored in etcd. It is used to emit cluster,
	// node and user certificates.
	pkiCA = &pki.Certificate{
		Namespace: &pkiNamespace,
		Issuer:    pki.SelfSigned,
		Template:  identity.CACertificate("Metropolis Cluster CA"),
		Name:      "cluster-ca",
	}
)
