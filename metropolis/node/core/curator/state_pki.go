package curator

import (
	"source.monogon.dev/metropolis/pkg/pki"
)

var (
	// pkiNamespace is the etcd/pki namespace in which the Metropolis cluster CA
	// data will live.
	pkiNamespace = pki.Namespaced("/cluster-pki/")
	// pkiCA is the main cluster CA, stored in etcd. It is used to emit cluster,
	// node and user certificates.
	pkiCA = pkiNamespace.New(pki.SelfSigned, "cluster-ca", pki.CA("Metropolis Cluster CA"))
)
