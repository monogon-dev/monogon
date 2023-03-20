package roleserve

import (
	"source.monogon.dev/metropolis/node/kubernetes"
)

// KubernetesStatus is an Event Value structure populated by a running
// Kubernetes instance. It allows external services to access the Kubernetes
// Service whenever available (ie. enabled and started by the Role Server).
type KubernetesStatus struct {
	Controller *kubernetes.Controller
}
