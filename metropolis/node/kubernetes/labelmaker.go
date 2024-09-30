package kubernetes

import (
	"context"
	"strings"
	"time"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/osbase/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

// labelmaker is a service responsible for keeping Kubernetes node labels in sync
// with Metropolis data. Currently, it synchronized Metropolis node roles into
// corresponding Kubernetes roles (implemented as labels).
//
// The labelmaker service runs on all Kubernetes controller nodes. This is safe,
// but it can cause spurious updates or update errors.
//
// TODO(q3k): leader elect a node responsible for this (and other Kubernetes
// control loops, too).
type labelmaker struct {
	clientSet kubernetes.Interface
	curator   ipb.CuratorClient
}

// managedLabelPrefixes are string prefixes to the keys of labels that we manage.
// This is used to filter out labels when doing a local diff of state vs. intent,
// before Kubernetes server-side apply is used.
var managedLabelPrefixes = []string{
	"node-role.kubernetes.io/",
}

// filterManaged filters out all labels not falling under managedLabelPrefixes.
func filterManaged(l common.Labels) common.Labels {
	res := make(map[string]string)
	for k, v := range l {
		for _, prefix := range managedLabelPrefixes {
			if strings.HasPrefix(k, prefix) {
				res[k] = v
				break
			}
		}
	}
	return res
}

// getIntendedLabelsPerNode returns labels that a node should have according to
// Curator data.
func getIntendedLabelsPerNode(node *ipb.Node) common.Labels {
	labels := make(map[string]string)
	// Mark role labels with empty string as content, following convention set by
	// kubeadm et al.
	if node.Roles.ConsensusMember != nil {
		labels["node-role.kubernetes.io/ConsensusMember"] = ""
	}
	if node.Roles.KubernetesController != nil {
		labels["node-role.kubernetes.io/KubernetesController"] = ""
	}
	if node.Roles.KubernetesWorker != nil {
		labels["node-role.kubernetes.io/KubernetesWorker"] = ""
	}
	return labels
}

// getCurrentLabelsForNode returns the current labels in Kubernetes for a given
// node.
//
// If the given node does not exist in Kubernetes, an empty label map is returned.
func getCurrentLabelsForNode(ctx context.Context, clientset kubernetes.Interface, nid string) (common.Labels, error) {
	node, err := clientset.CoreV1().Nodes().Get(ctx, nid, v1.GetOptions{})
	if kerrors.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return node.Labels, nil
}

// run the labelmaker service.
func (l *labelmaker) run(ctx context.Context) error {
	for {
		ctxT, ctxC := context.WithTimeout(ctx, time.Minute*10)
		err := l.runInner(ctxT)
		// runInner will return an error either on context expiry (which should lead to a
		// silent restart) or any other error (either processing error or parent context
		// timeout) which should be bubbled up.
		if err == nil {
			// This shouldn't happen, but let's handle it just in case.
			ctxC()
			continue
		}
		if ctxT.Err() != nil && ctx.Err() == nil {
			// We should check with errors.Is whether the returned error is influenced by
			// context expiry, but gRPC apparently doesn't (always?) wrap the context error
			// cause into the returned error. Skipping this check might lead us to
			// misidentifying other errors as context expiry if these two race, but there's
			// not much we can do about this.
			ctxC()
			continue
		}
		// Probably some other kind of error. Return it.
		ctxC()
		return err
	}
}

// runInner runs the labelmaker service with one active curator Watch call. This
// will be interrupted by the given context timing out (as provided by run) in
// order to ensure nodes that couldn't get processed have another chance later
// on.
func (l *labelmaker) runInner(ctx context.Context) error {
	srv, err := l.curator.Watch(ctx, &ipb.WatchRequest{
		Kind: &ipb.WatchRequest_NodesInCluster_{
			NodesInCluster: &ipb.WatchRequest_NodesInCluster{},
		},
	})
	if err != nil {
		return err
	}
	defer srv.CloseSend()

	for {
		ev, err := srv.Recv()
		if err != nil {
			return err
		}

		supervisor.Logger(ctx).Infof("Processing %d nodes...", len(ev.Nodes))

		for _, node := range ev.Nodes {
			if err := l.processClusterNode(ctx, node); err != nil {
				supervisor.Logger(ctx).Warningf("Failed to process labels on node %s: %v", node.Id, err)
			}
		}

		if ev.Progress == ipb.WatchEvent_PROGRESS_LAST_BACKLOGGED {
			supervisor.Logger(ctx).Infof("Caught up with node backlog, now watching for updates.")
		}
	}
}

// processClusterNodes runs the label reconciliation algorithm on a single node.
// It requests current labels from Kubernetes and issues a server-side-apply
// operation to bring them into the requested state, if there is a diff.
func (l *labelmaker) processClusterNode(ctx context.Context, node *ipb.Node) error {
	intent := getIntendedLabelsPerNode(node)
	state, err := getCurrentLabelsForNode(ctx, l.clientSet, node.Id)
	if err != nil {
		return err
	}
	if state == nil {
		return nil
	}
	intent = filterManaged(intent)
	state = filterManaged(state)

	if intent.Equals(state) {
		return nil
	}
	supervisor.Logger(ctx).Infof("Updating labels on node %s... ", node.Id)
	for k, v := range state {
		if intent[k] != v {
			supervisor.Logger(ctx).Infof("  Removing %s=%q", k, v)
		}
	}
	for k, v := range intent {
		if state[k] != v {
			supervisor.Logger(ctx).Infof("  Adding %s=%q", k, v)
		}
	}

	cfg := applycorev1.Node(node.Id)
	cfg.Labels = intent
	_, err = l.clientSet.CoreV1().Nodes().Apply(ctx, cfg, v1.ApplyOptions{
		FieldManager: "metropolis-labelmaker",
		Force:        true,
	})
	return err
}