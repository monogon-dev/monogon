// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package reconciler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"google.golang.org/protobuf/proto"
	"k8s.io/client-go/kubernetes"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/curator"
	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/node/core/productinfo"
	"source.monogon.dev/osbase/event/etcd"
	"source.monogon.dev/osbase/event/memory"
	"source.monogon.dev/osbase/supervisor"
	"source.monogon.dev/version"
	vpb "source.monogon.dev/version/spec"
)

// This file contains the reconciler Service, whose purpose is to run
// reconcileAll in a controlled way (leader elected and only when other nodes
// are compatible) and to set the reconciler status in etcd.
// The file also contains WaitReady, which watches the status and returns
// when the apiserver can start serving.
// These two form the public interface of the reconciler.

const (
	// statusKey is the key in the curator etcd namespace
	// under which the reconciler status is stored.
	// At some point, we do a transaction involving both this key and
	// the nodes prefix, so both must be in the same namespace.
	statusKey = "/kubernetes/reconciler/status"
	// electionPrefix is the etcd prefix where
	// a node is elected to run the reconciler.
	electionPrefix = "/kubernetes/reconciler/leader"
)

var (
	// minReconcilerRelease is the minimum Metropolis release which
	// the node last performing reconciliation must have
	// for the local node to be able to start serving.
	// Set this to the next release when making changes to reconciled state
	// which must be applied before starting to serve.
	minReconcilerRelease = &vpb.Version_Release{Major: 0, Minor: 1, Patch: 0}
	// minApiserverRelease is the minimum Metropolis release which all Kubernetes
	// controller nodes must have before the local node can reconcile.
	// This will be written to minimum_compatible_release in the reconciler status,
	// and thus block any reappearing apiservers with a lower release from serving,
	// until a reconciler of a lower release has run.
	// Increase this when making changes to reconciled state which are
	// incompatible with apiservers serving at the current minApiserverRelease.
	minApiserverRelease = &vpb.Version_Release{Major: 0, Minor: 1, Patch: 0}
)

// reconcileWait is the wait time between getting elected and
// starting to reconcile.
// It is a variable to allow changing it from tests.
var reconcileWait = 5 * time.Second

// WaitReady watches the reconciler status and returns once initial
// reconciliation is done and the reconciled state is compatible.
func (s *Service) WaitReady(ctx context.Context) error {
	value := etcd.NewValue(s.Etcd, statusKey, func(_, data []byte) (*ppb.KubernetesReconcilerStatus, error) {
		status := &ppb.KubernetesReconcilerStatus{}
		if err := proto.Unmarshal(data, status); err != nil {
			return nil, fmt.Errorf("could not unmarshal: %w", err)
		}
		return status, nil
	})

	w := value.Watch()
	defer w.Close()

	for {
		status, err := w.Get(ctx)
		if err != nil {
			return err
		}

		state := "unknown"
		switch status.State {
		case ppb.KubernetesReconcilerStatus_STATE_DONE:
			state = "done"
		case ppb.KubernetesReconcilerStatus_STATE_WORKING:
			state = "working"
		}
		supervisor.Logger(ctx).Infof("Reconciler status: %s, version: %s, minimum compatible release: %s. Local node version: %s, minimum reconciler release: %s.",
			state,
			version.Semver(status.Version),
			version.Release(status.MinimumCompatibleRelease),
			version.Semver(productinfo.Get().Version),
			version.Release(minReconcilerRelease),
		)

		if version.ReleaseLessThan(productinfo.Get().Version.Release, status.MinimumCompatibleRelease) {
			supervisor.Logger(ctx).Info("Not ready, because the local node release is below the reconciler minimum compatible release. Waiting for status change.")
			continue
		}

		if version.ReleaseLessThan(status.Version.Release, minReconcilerRelease) {
			supervisor.Logger(ctx).Info("Not ready, because the reconciler release is below the local required minimum. Waiting for status change.")
			continue
		}

		// Startup is intentionally not blocked by state=working.
		// As long as a node is compatible with both the before and after state,
		// it can continue running, and startup should also be allowed.
		// This way, disruption is minimized in case reconciliation fails
		// to complete and the status stays in working state.
		// For the initial reconcile, a status is only created after it is complete.

		return nil
	}
}

// Service is the reconciler service.
type Service struct {
	// Etcd is an etcd client for the curator namespace.
	Etcd client.Namespaced
	// ClientSet is what the reconciler uses to interact with the apiserver.
	ClientSet kubernetes.Interface
	// NodeID is the ID of the local node.
	NodeID string
	// releases is set by watchNodes and watched by other parts of the service.
	releases memory.Value[*nodeReleases]
}

// nodeReleases contains a summary of the releases of all
// Kubernetes controller nodes currently in the cluster.
type nodeReleases struct {
	minRelease *vpb.Version_Release
	maxRelease *vpb.Version_Release
	// revision is the etcd revision at which this info is valid.
	revision int64
}

// The reconciler service has a tree of runnables:
//
// - watch-nodes: Watches nodes in etcd and sets releases.
// - watch-releases: Watches releases and runs elect while the local node is
//   the latest release.
//   - elect: Performs etcd leader election and starts lead once elected.
//     - lead: Checks current status, watches releases until incompatible
//       nodes disappear, updates status, runs reconcileAll.

// Run is the root runnable of the reconciler service.
func (s *Service) Run(ctx context.Context) error {
	err := supervisor.Run(ctx, "watch-nodes", s.watchNodes)
	if err != nil {
		return fmt.Errorf("could not run watch-nodes: %w", err)
	}

	err = supervisor.Run(ctx, "watch-releases", s.watchReleases)
	if err != nil {
		return fmt.Errorf("could not run watch-releases: %w", err)
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}

// watchNodes watches nodes in etcd, and publishes a summary of
// releases of Kubernetes controller nodes in s.releases.
func (s *Service) watchNodes(ctx context.Context) error {
	nodesStart, nodesEnd := curator.NodeEtcdPrefix.KeyRange()

	var revision int64
	nodeToRelease := make(map[string]string)
	releaseCount := make(map[string]int)
	releaseStruct := make(map[string]*vpb.Version_Release)

	updateNode := func(kv *mvccpb.KeyValue) {
		nodeKey := string(kv.Key)
		// Subtract the previous release of this node if any.
		if prevRelease, ok := nodeToRelease[nodeKey]; ok {
			delete(nodeToRelease, nodeKey)
			releaseCount[prevRelease] -= 1
			if releaseCount[prevRelease] == 0 {
				delete(releaseCount, prevRelease)
				delete(releaseStruct, prevRelease)
			}
		}

		// Parse the node release. Skip if the node was deleted, is not a
		// Kubernetes controller, or does not have a release.
		if len(kv.Value) == 0 {
			return
		}
		var node ppb.Node
		if err := proto.Unmarshal(kv.Value, &node); err != nil {
			supervisor.Logger(ctx).Errorf("Failed to unmarshal node %q: %w", nodeKey, err)
			return
		}
		if node.Roles.KubernetesController == nil {
			return
		}
		if node.Status == nil || node.Status.Version == nil {
			return
		}
		release := version.Release(node.Status.Version.Release)
		// Add the new release.
		nodeToRelease[nodeKey] = release
		if releaseCount[release] == 0 {
			releaseStruct[release] = node.Status.Version.Release
		}
		releaseCount[release] += 1
	}

	publish := func() {
		minRelease := productinfo.Get().Version.Release
		maxRelease := productinfo.Get().Version.Release
		for _, release := range releaseStruct {
			if version.ReleaseLessThan(release, minRelease) {
				minRelease = release
			}
			if version.ReleaseLessThan(maxRelease, release) {
				maxRelease = release
			}
		}
		s.releases.Set(&nodeReleases{
			minRelease: minRelease,
			maxRelease: maxRelease,
			revision:   revision,
		})
	}

	// Get the initial nodes data.
	get, err := s.Etcd.Get(ctx, nodesStart, clientv3.WithRange(nodesEnd))
	if err != nil {
		return fmt.Errorf("when retrieving initial nodes: %w", err)
	}

	for _, kv := range get.Kvs {
		updateNode(kv)
	}
	revision = get.Header.Revision
	publish()

	supervisor.Signal(ctx, supervisor.SignalHealthy)

	// Watch for changes.
	wch := s.Etcd.Watch(ctx, nodesStart, clientv3.WithRange(nodesEnd), clientv3.WithRev(revision+1))
	for resp := range wch {
		if err := resp.Err(); err != nil {
			return fmt.Errorf("watch failed: %w", err)
		}
		for _, ev := range resp.Events {
			updateNode(ev.Kv)
		}
		revision = resp.Header.Revision
		publish()
	}
	return fmt.Errorf("channel closed: %w", ctx.Err())
}

// watchReleases watches s.releases, and runs elect for as long as
// the local node has the latest release.
func (s *Service) watchReleases(ctx context.Context) error {
	w := s.releases.Watch()
	defer w.Close()

	r, err := w.Get(ctx)
	if err != nil {
		return err
	}

	shouldRun := !version.ReleaseLessThan(productinfo.Get().Version.Release, r.maxRelease)
	if shouldRun {
		supervisor.Logger(ctx).Info("This Kubernetes controller node has the latest release, starting election.")
		err := supervisor.Run(ctx, "elect", s.elect)
		if err != nil {
			return fmt.Errorf("could not run elect: %w", err)
		}
	} else {
		supervisor.Logger(ctx).Infof("This Kubernetes controller node does not have the latest release, not starting election. Latest release: %s", version.Release(r.maxRelease))
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)

	for {
		r, err := w.Get(ctx)
		if err != nil {
			return err
		}
		shouldRunNow := !version.ReleaseLessThan(productinfo.Get().Version.Release, r.maxRelease)
		if shouldRunNow != shouldRun {
			return errors.New("latest release changed, restarting")
		}
	}
}

func (s *Service) elect(ctx context.Context) error {
	session, err := concurrency.NewSession(s.Etcd.ThinClient(ctx))
	if err != nil {
		return fmt.Errorf("creating session failed: %w", err)
	}

	defer func() {
		session.Orphan()
		// ctx may be canceled, but we still try to revoke with a short timeout.
		revokeCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err := s.Etcd.Revoke(revokeCtx, session.Lease())
		cancel()
		if err != nil {
			supervisor.Logger(ctx).Warningf("Failed to revoke lease: %v", err)
		}
	}()

	supervisor.Signal(ctx, supervisor.SignalHealthy)

	supervisor.Logger(ctx).Infof("Campaigning. Lease ID: %x", session.Lease())
	election := concurrency.NewElection(session, electionPrefix)

	// The election value is unused; we put the node ID there for manual inspection.
	err = election.Campaign(ctx, s.NodeID)
	if err != nil {
		return fmt.Errorf("campaigning failed: %w", err)
	}
	supervisor.Logger(ctx).Info("Elected.")

	leadCtx, leadCancel := context.WithCancel(ctx)
	go func() {
		<-session.Done()
		leadCancel()
	}()

	isLeaderCmp := clientv3.Compare(clientv3.CreateRevision(election.Key()), "=", election.Rev())
	return s.lead(leadCtx, isLeaderCmp)
}

func (s *Service) lead(ctx context.Context, isLeaderCmp clientv3.Cmp) error {
	log := supervisor.Logger(ctx)

	// Retrieve the initial status.
	status := &ppb.KubernetesReconcilerStatus{}
	statusGet, err := s.Etcd.Get(ctx, statusKey)
	if err != nil {
		return fmt.Errorf("when getting status: %w", err)
	}
	if len(statusGet.Kvs) == 1 {
		err := proto.Unmarshal(statusGet.Kvs[0].Value, status)
		if err != nil {
			log.Warningf("Could not unmarshal status: %v", err)
			status = nil
		}
	} else {
		status = nil
	}

	doneStatus := &ppb.KubernetesReconcilerStatus{
		State:                    ppb.KubernetesReconcilerStatus_STATE_DONE,
		Version:                  productinfo.Get().Version,
		MinimumCompatibleRelease: minApiserverRelease,
	}
	doneStatusBytes, err := proto.Marshal(doneStatus)
	if err != nil {
		return fmt.Errorf("could not marshal status: %w", err)
	}

	if status == nil {
		// The status does not exist yet. Reconcile, then create the status.
		log.Info("Status does not exist yet.")
	} else if proto.Equal(status, doneStatus) {
		// The status is already what we would set, so leave it as is.
		log.Info("Status is already up to date.")
	} else if !version.ReleaseLessThan(productinfo.Get().Version.Release, status.Version.Release) &&
		!version.ReleaseLessThan(status.MinimumCompatibleRelease, minApiserverRelease) {
		// The status does not allow apiservers to start serving which would be
		// incompatible after we reconcile. So just set the state to working.
		log.Info("Status is compatible, setting state to working.")
		if status.State != ppb.KubernetesReconcilerStatus_STATE_WORKING {
			status.State = ppb.KubernetesReconcilerStatus_STATE_WORKING

			workingStatusBytes, err := proto.Marshal(status)
			if err != nil {
				return fmt.Errorf("could not marshal status: %w", err)
			}
			resp, err := s.Etcd.Txn(ctx).If(isLeaderCmp).Then(
				clientv3.OpPut(statusKey, string(workingStatusBytes)),
			).Commit()
			if err != nil {
				return fmt.Errorf("failed to update status: %w", err)
			}
			if !resp.Succeeded {
				return errors.New("lost leadership, could not update status")
			}
		}
	} else {
		// The status allows apiservers to start which would be incompatible after
		// we reconcile. We need to wait for any such nodes to disappear, then set
		// the status to disallow these nodes from starting before reconciling.
		// While reconciliation is ongoing, we are in an intermediate state
		// between the previous and the new reconciled state, and we only want
		// to allow nodes that are compatible with both. So we use the minimum of
		// the two versions and the maximum of the two MinimumCompatibleReleases,
		// which results in allowing the intersection of the two statuses.
		log.Info("Status allows incompatible releases, need to restrict.")

		status.State = ppb.KubernetesReconcilerStatus_STATE_WORKING
		if !version.ReleaseLessThan(status.Version.Release, productinfo.Get().Version.Release) {
			status.Version = productinfo.Get().Version
		}
		if version.ReleaseLessThan(status.MinimumCompatibleRelease, minApiserverRelease) {
			status.MinimumCompatibleRelease = minApiserverRelease
		}
		restrictedStatusBytes, err := proto.Marshal(status)
		if err != nil {
			return fmt.Errorf("could not marshal status: %w", err)
		}

		releasesW := s.releases.Watch()
		defer releasesW.Close()

		lastLogRelease := ""
		for {
			releases, err := releasesW.Get(ctx)
			if err != nil {
				return err
			}
			if version.ReleaseLessThan(productinfo.Get().Version.Release, releases.maxRelease) {
				// We will likely get canceled soon by watchReleases restarting, unless
				// this is a very short transient that is not noticed by watchReleases.
				continue
			}
			if version.ReleaseLessThan(releases.minRelease, minApiserverRelease) {
				rel := version.Release(releases.minRelease)
				if rel != lastLogRelease {
					lastLogRelease = rel
					log.Infof("There are incompatible nodes, waiting for node changes. Minimum node release: %s Need at least: %s", rel, version.Release(minApiserverRelease))
				}
				continue
			}

			nodesStart, nodesEnd := curator.NodeEtcdPrefix.KeyRange()
			resp, err := s.Etcd.Txn(ctx).If(
				isLeaderCmp,
				clientv3.Compare(clientv3.ModRevision(nodesStart).WithRange(nodesEnd), "<", releases.revision+1),
			).Then(
				clientv3.OpPut(statusKey, string(restrictedStatusBytes)),
			).Commit()
			if err != nil {
				return fmt.Errorf("failed to update status: %w", err)
			}
			if !resp.Succeeded {
				// This could happen either if we lost leadership, or any node was
				// modified since we got the releases. If a node was modified, this
				// should be seen soon by the nodes watcher. If we lost leadership,
				// we will get canceled soon, and it's fine to go back to watching.
				log.Info("Transaction failed, retrying.")
				continue
			}
			break
		}
	}

	if status != nil {
		// A status exists, which means a reconciler has been running before.
		// Wait a bit for any still outstanding Kubernetes API requests by the
		// previous reconciler to be processed.
		// The Kubernetes API does not support making requests conditional on an
		// etcd lease, so requests can still be processed after leadership expired.
		// This is best effort, since requests could take arbitrarily long to be
		// processed. The periodic reconcile below ensures that we eventually
		// reach the desired state and stay there.
		select {
		case <-time.After(reconcileWait):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	log.Info("Performing initial resource reconciliation...")
	// If the apiserver was just started, reconciliation will fail until the
	// apiserver is ready. To keep the logs clean, retry with exponential
	// backoff and only start logging errors after some time has passed.
	startLogging := time.Now().Add(2 * time.Second)
	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = 100 * time.Millisecond
	bo.MaxElapsedTime = 0
	err = backoff.Retry(func() error {
		err := reconcileAll(ctx, s.ClientSet)
		if err != nil && time.Now().After(startLogging) {
			log.Errorf("Still couldn't do initial reconciliation: %v", err)
			startLogging = time.Now().Add(10 * time.Second)
		}
		return err
	}, backoff.WithContext(bo, ctx))
	if err != nil {
		return err
	}
	log.Infof("Initial resource reconciliation succeeded.")

	// Update status.
	if !proto.Equal(status, doneStatus) {
		resp, err := s.Etcd.Txn(ctx).If(isLeaderCmp).Then(
			clientv3.OpPut(statusKey, string(doneStatusBytes)),
		).Commit()
		if err != nil {
			return fmt.Errorf("failed to update status: %w", err)
		}
		if !resp.Succeeded {
			return errors.New("lost leadership, could not update status")
		}
	}

	// Reconcile at a regular interval.
	t := time.NewTicker(30 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			err := reconcileAll(ctx, s.ClientSet)
			if err != nil {
				log.Warning(err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
