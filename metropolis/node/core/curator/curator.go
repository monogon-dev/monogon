// package curator implements the Curator, a service responsible for management
// of the Metropolis cluster that it is running on.
//
// The Curator is implemented as a leader-elected service. Instances of the
// service are running colocated with all nodes that run a consensus (etcd)
// server.
//
// Each instance listens on all network interfaces, for requests both from the
// code running on the same node, for traffic from other nodes (eg. ones that do
// not run an instance of the Curator) and from external users.
//
// The curator leader keeps its state fully in etcd. Followers forward all
// requests to the active leader.
package curator

import (
	"context"
	"errors"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// Config is the configuration of the curator.
type Config struct {
	// NodeCredentials are the identity credentials for the node that is running
	// this curator.
	NodeCredentials *identity.NodeCredentials
	Consensus       consensus.ServiceHandle
	// LeaderTTL is the timeout on the lease used to perform leader election.
	// Any active leader must continue updating its lease at least this often,
	// or the lease (and leadership) will be lost.
	// Lower values allow for faster failovers. Higher values allow for higher
	// resiliency against short network partitions.
	// A value less or equal to zero will default to 60 seconds.
	LeaderTTL time.Duration
}

// Service is the Curator service. See the package-level documentation for more
// information.
type Service struct {
	// config is the configuration with which the service was started.
	config *Config

	// ttl is the effective TTL value from Config.LeaderTTL (if given as <= 0,
	// this is the value that has been fixed up to some default).
	ttl int

	// status is a memory Event Value for keeping the electionStatus of this
	// instance. It is not exposed to users of the Curator.
	status memory.Value[*electionStatus]
}

// New creates a new curator Service.
func New(cfg Config) *Service {
	return &Service{
		config: &cfg,
	}
}

// electionStatus represents the status of this curator's leader election
// attempts within the cluster.
type electionStatus struct {
	// leader is set if this curator is a leader, nil otherwise. This cannot be set
	// if follower is also set, both leader and follower might be nil to signify
	// that this curator instance is not part of a quorum.
	leader *electionStatusLeader
	// follower is set if the curator is a follower for another leader, nil
	// otherwise. This cannot be set if leader is also set. However, both leader and
	// follower might be nil to signify that this curator instance is not part of a
	// quorum.
	follower *electionStatusFollower
}

type electionStatusLeader struct {
	// lockKey is the etcd key for which a lease value is set with revision
	// lockRev. This key/revision must be ensured to exist when any etcd access
	// is performed by the curator to ensure that it is still the active leader
	// according to the rest of the cluster.
	lockKey string
	lockRev int64
}

func (e *electionStatusLeader) equal(o *electionStatusLeader) bool {
	if e.lockKey != o.lockKey {
		return false
	}
	if e.lockRev != o.lockRev {
		return false
	}
	return true
}

type electionStatusFollower struct {
	lock *ppb.LeaderElectionValue
}

func (e *electionStatusFollower) equal(o *electionStatusFollower) bool {
	if e.lock.NodeId != o.lock.NodeId {
		return false
	}
	if e.lock.Ttl != o.lock.Ttl {
		return false
	}
	return true
}

// buildLockValue returns a serialized etcd value that will be set by the
// instance when it becomes a leader. This value is a serialized
// LeaderElectionValue from private/storage.proto.
func (c *Config) buildLockValue(ttl int) ([]byte, error) {
	v := &ppb.LeaderElectionValue{
		NodeId: c.NodeCredentials.ID(),
		Ttl:    uint64(ttl),
	}
	bytes, err := proto.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("when marshaling value: %w", err)
	}
	return bytes, nil
}

var (
	// electionPrefix is the prefix under which the curator instances will
	// attempt to perform leader election.
	//
	// The trailing slash is omitted, as the etcd concurrency library appends one.
	electionPrefix = "/leader"
)

// elect runs a single leader election attempt. The status of the service will
// be updated with electionStatus values as the election makes progress.
func (s *Service) elect(ctx context.Context) error {
	lv, err := s.config.buildLockValue(s.ttl)
	if err != nil {
		return fmt.Errorf("building lock value failed: %w", err)
	}

	w := s.config.Consensus.Watch()
	defer w.Close()
	st, err := w.Get(ctx)
	if err != nil {
		return fmt.Errorf("getting consensus status failed: %w", err)
	}
	cl, err := st.CuratorClient()
	if err != nil {
		return fmt.Errorf("getting consensus client failed: %w", err)
	}

	if err := s.cleanupPreviousLifetime(ctx, cl); err != nil {
		supervisor.Logger(ctx).Warningf("Failed to cleanup previous lifetime: %v", err)
	}

	// Establish a lease/session with etcd.
	session, err := concurrency.NewSession(cl.ThinClient(ctx),
		concurrency.WithContext(ctx),
		concurrency.WithTTL(s.ttl))
	if err != nil {
		return fmt.Errorf("creating session failed: %w", err)
	}

	// Kill the session whenever we lose leadership or error out.
	defer func() {
		err := session.Close()
		if err != nil {
			supervisor.Logger(ctx).Warningf("Failed to close session: %v", err)
		}
	}()

	supervisor.Logger(ctx).Infof("Curator established lease, ID: %d", session.Lease())
	election := concurrency.NewElection(session, electionPrefix)

	// Observer context, we need to cancel it to not leak the observer
	// goroutine/channel.
	octx, octxC := context.WithCancel(ctx)
	defer octxC()

	// Channel that gets updates about the current leader in the cluster.
	observerC := election.Observe(octx)
	// Channel that gets updates about this instance becoming a leader.
	campaignerC := make(chan error)

	// Campaign to become leader. This blocks until leader election is successful
	// and this instance is now the leader.
	//
	// The lock value is converted to string from raw binary bytes, but that's fine
	// as that string is converted back to []byte within the etcd client library (in
	// OpPut).
	go func() {
		campaignerC <- election.Campaign(ctx, string(lv))
	}()

	// While campaigning, update the electionStatus with information about the
	// current leader.
	for {
		select {
		case o, ok := <-observerC:
			if !ok {
				return errors.New("election observation failed")
			}
			var lock ppb.LeaderElectionValue
			if err := proto.Unmarshal(o.Kvs[0].Value, &lock); err != nil {
				return fmt.Errorf("parsing existing lock value failed: %w", err)
			}
			s.status.Set(&electionStatus{
				follower: &electionStatusFollower{
					lock: &lock,
				},
			})
		case err = <-campaignerC:
			if err == nil {
				goto campaigned
			}
			return fmt.Errorf("campaigning failed: %w", err)
		}
	}

campaigned:
	supervisor.Logger(ctx).Info("Curator became leader.")

	// Update status, watchers will now know that this curator is the leader.
	s.status.Set(&electionStatus{
		leader: &electionStatusLeader{
			lockKey: election.Key(),
			lockRev: election.Rev(),
		},
	})

	// Wait until either we loose the lease/session or our context expires.
	select {
	case <-ctx.Done():
		supervisor.Logger(ctx).Warningf("Context canceled, quitting.")
		return fmt.Errorf("curator session canceled: %w", ctx.Err())
	case <-session.Done():
		supervisor.Logger(ctx).Warningf("Session done, quitting.")
		return fmt.Errorf("curator session done")
	}

}

// cleanupPreviousLifetime checks if we just started up after ungracefully losing
// our leadership, and attempts to clean up if so.
//
// Having the rest of the cluster assume this node is still the leader is not a
// problem from a correctness point of view (as the node will refuse to serve
// leader requests with Unimplemented), but it is quite an eyesore to operators,
// as all nodes just end up having a ton of client processes all complain with
// odd 'Unimplemented' errors.
func (s *Service) cleanupPreviousLifetime(ctx context.Context, cl client.Namespaced) error {
	// Get the active leader key and value.
	resp, err := cl.Get(ctx, electionPrefix, clientv3.WithFirstCreate()...)
	if err != nil {
		return err
	}
	if len(resp.Kvs) < 1 {
		return nil
	}

	// Check that this key belonged to use by comparing the embedded node ID with our
	// own node ID.
	key := string(resp.Kvs[0].Key)
	rev := resp.Kvs[0].ModRevision
	var lock ppb.LeaderElectionValue
	if err := proto.Unmarshal(resp.Kvs[0].Value, &lock); err != nil {
		return fmt.Errorf("parsing existing lock value failed: %w", err)
	}

	// Not our node? Nothing to do.
	if lock.NodeId != s.config.NodeCredentials.ID() {
		return nil
	}

	// Now here's the sketchy part: removing the leader election key if we think it
	// used to be ours.
	supervisor.Logger(ctx).Infof("Detecting our own stale lock, attempting to remove...")

	// Just removing the key should be correct, as the key encodes the original
	// session ID that created the leadership key, and the session ID is unique per
	// leader. So if the key exists, it is guaranteed to have been created and updated
	// by exactly one session. And having read it earlier, we know it is a session
	// that was owned by this node, as it proclaimed this node as the leader.
	//
	// The only scenario in which this can fail is if we have the same node ID
	// running more than one curator service and thus more than one leader election.
	// But that would be a serious programming/design bug and other things would
	// likely break at this point, anyway.
	//
	// For safety, we add a check that it is still the same ModRevision as when we
	// checked it earlier on, but that's probably unnecessary.
	txn := cl.Txn(ctx).If(clientv3.Compare(clientv3.ModRevision(key), "=", rev))
	txn = txn.Then(clientv3.OpDelete(key))
	resp2, err := txn.Commit()
	if err != nil {
		return err
	}
	if resp2.Succeeded {
		supervisor.Logger(ctx).Infof("Cleanup successful")
	} else {
		// This will happen if the key expired by itself already.
		supervisor.Logger(ctx).Warningf("Cleanup failed - maybe our old lease already expired...")
	}
	return nil
}

func (s *Service) Run(ctx context.Context) error {
	// Start local election watcher. This logs what this curator knows about its own
	// leadership.
	go func() {
		w := s.status.Watch()
		for {
			s, err := w.Get(ctx)
			if err != nil {
				supervisor.Logger(ctx).Warningf("Election watcher exiting: get(): %v", err)
				return
			}
			if l := s.leader; l != nil {
				supervisor.Logger(ctx).Infof("Election watcher: this node's curator is leader (lock key %q, rev %d)", l.lockKey, l.lockRev)
			} else {
				supervisor.Logger(ctx).Infof("Election watcher: this node's curator is a follower")
			}
		}
	}()

	supervisor.Logger(ctx).Infof("Waiting for consensus...")
	w := s.config.Consensus.Watch()
	defer w.Close()
	st, err := w.Get(ctx, consensus.FilterRunning)
	if err != nil {
		return fmt.Errorf("while waiting for consensus: %w", err)
	}
	supervisor.Logger(ctx).Infof("Got consensus, starting up...")
	etcd, err := st.CuratorClient()
	if err != nil {
		return fmt.Errorf("while retrieving consensus client: %w", err)
	}

	// Start listener. This is a gRPC service listening on all interfaces, providing
	// the Curator API to consumers, dispatching to either a locally running leader,
	// or forwarding to a remotely running leader.
	lis := listener{
		node:        s.config.NodeCredentials,
		etcd:        etcd,
		etcdCluster: st.ClusterClient(),
		consensus:   s.config.Consensus,
		status:      &s.status,
	}
	if err := supervisor.Run(ctx, "listener", lis.run); err != nil {
		return fmt.Errorf("when starting listener: %w", err)
	}

	// Calculate effective TTL. This replicates the behaviour of clientv3's WithTTL,
	// but allows us to explicitly log the used TTL.
	s.ttl = int(s.config.LeaderTTL.Seconds())
	if s.ttl <= 0 {
		s.ttl = 10
	}
	supervisor.Logger(ctx).Infof("Curator starting on prefix %q with lease TTL of %d seconds...", electionPrefix, s.ttl)

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	for {
		s.status.Set(&electionStatus{})
		err := s.elect(ctx)
		s.status.Set(&electionStatus{})

		if err != nil && errors.Is(err, ctx.Err()) {
			return fmt.Errorf("election round failed due to context cancelation, not attempting to re-elect: %w", err)
		}
		supervisor.Logger(ctx).Infof("Curator election round done: %v", err)
		supervisor.Logger(ctx).Info("Curator election restarting...")
	}
}
