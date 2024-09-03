package roleserve

import (
	"context"

	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/curator/watcher"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/osbase/event/memory"
	"source.monogon.dev/osbase/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// workerRoleFetch is the Role Fetcher, an internal bookkeeping service
// responsible for populating localRoles based on a curatorConnection whenever
// the node is HOME and cluster credentials / curator access is available.
type workerRoleFetch struct {
	storageRoot       *localstorage.Root
	curatorConnection *memory.Value[*CuratorConnection]

	// localRoles will be written.
	localRoles *memory.Value[*cpb.NodeRoles]
}

func (s *workerRoleFetch) run(ctx context.Context) error {
	// Wait for a 'curator connection' first. This isn't a real connection but just a
	// set of credentials and a potentially working resolver. Here, we just use its
	// presence as a marking that the local disk has been mounted and thus we can
	// attempt to read a persisted cluster directory.
	w := s.curatorConnection.Watch()
	_, err := w.Get(ctx)
	if err != nil {
		w.Close()
		return err
	}
	w.Close()

	// Read persisted roles if available.
	exists, _ := s.storageRoot.Data.Node.PersistedRoles.Exists()
	if exists {
		supervisor.Logger(ctx).Infof("Attempting to read persisted node roles...")
		data, err := s.storageRoot.Data.Node.PersistedRoles.Read()
		if err != nil {
			supervisor.Logger(ctx).Errorf("Failed to read persisted roles: %w", err)
		} else {
			var nr cpb.NodeRoles
			if err := proto.Unmarshal(data, &nr); err != nil {
				supervisor.Logger(ctx).Errorf("Failed to unmarshal persisted roles: %w", err)
			} else {
				supervisor.Logger(ctx).Infof("Got persisted role data from disk:")
			}
			if nr.ConsensusMember != nil {
				supervisor.Logger(ctx).Infof(" - control plane member, existing peers: %+v", nr.ConsensusMember.Peers)
			}
			if nr.KubernetesController != nil {
				supervisor.Logger(ctx).Infof(" - kubernetes controller")
			}
			if nr.KubernetesWorker != nil {
				supervisor.Logger(ctx).Infof(" - kubernetes worker")
			}
			s.localRoles.Set(&nr)
		}
	} else {
		supervisor.Logger(ctx).Infof("No persisted node roles.")
	}

	// Run networked part in a sub-runnable so that network errors don't cause us to
	// retry the above and make us possibly trigger spurious restarts.
	supervisor.Run(ctx, "watcher", func(ctx context.Context) error {
		cw := s.curatorConnection.Watch()
		defer cw.Close()
		cc, err := cw.Get(ctx)
		if err != nil {
			return err
		}

		nodeID := cc.nodeID()
		cur := ipb.NewCuratorClient(cc.conn)
		w := watcher.WatchNode(ctx, cur, nodeID)
		defer w.Close()

		supervisor.Signal(ctx, supervisor.SignalHealthy)

		// Run watch for current node, update localRoles whenever we get something
		// new.
		for w.Next() {
			n := w.Node()
			supervisor.Logger(ctx).Infof("Got new node data. Roles:")
			if n.Roles.ConsensusMember != nil {
				supervisor.Logger(ctx).Infof(" - control plane member, existing peers: %+v", n.Roles.ConsensusMember.Peers)
			}
			if n.Roles.KubernetesController != nil {
				supervisor.Logger(ctx).Infof(" - kubernetes controller")
			}
			if n.Roles.KubernetesWorker != nil {
				supervisor.Logger(ctx).Infof(" - kubernetes worker")
			}
			s.localRoles.Set(n.Roles)

			// Persist role data to disk.
			bytes, err := proto.Marshal(n.Roles)
			if err != nil {
				supervisor.Logger(ctx).Errorf("Failed to marshal node roles: %w", err)
			} else {
				err = s.storageRoot.Data.Node.PersistedRoles.Write(bytes, 0400)
				if err != nil {
					supervisor.Logger(ctx).Errorf("Failed to write node roles: %w", err)
				}
			}
		}
		return w.Error()
	})

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return ctx.Err()

}
