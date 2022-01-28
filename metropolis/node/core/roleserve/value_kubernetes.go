package roleserve

import (
	"context"

	"source.monogon.dev/metropolis/node/kubernetes"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
)

// KubernetesStatus is an Event Value structure populated by a running
// Kubernetes instance. It allows external services to access the Kubernetes
// Service whenever available (ie. enabled and started by the Role Server).
type KubernetesStatus struct {
	Svc *kubernetes.Service
}

type KubernetesStatusValue struct {
	value memory.Value
}

func (k *KubernetesStatusValue) Watch() *KubernetesStatusWatcher {
	return &KubernetesStatusWatcher{
		Watcher: k.value.Watch(),
	}
}

func (k *KubernetesStatusValue) set(v *KubernetesStatus) {
	k.value.Set(v)
}

type KubernetesStatusWatcher struct {
	event.Watcher
}

// Get waits until the Kubernetes services is available. The returned
// KubernetesStatus is guaranteed to contain a kubernetes.Service that was
// running at the time of this call returning (but which might have since been
// stopped).
func (k *KubernetesStatusWatcher) Get(ctx context.Context) (*KubernetesStatus, error) {
	v, err := k.Watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	return v.(*KubernetesStatus), nil
}
