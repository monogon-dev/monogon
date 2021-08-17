// package curator implements the Curator, a service responsible for management
// of the Metropolis cluster that it is running on.
//
// The Curator is implemented as a leader-elected service. Instances of the
// service are running colocated with all nodes that run a consensus (etcd)
// server.
// Each instance listens locally over gRPC for requests from code running on the
// same node, and publicly over gRPC for traffic from other nodes (eg. ones that
// do not run an instance of the Curator) and external users.
// The curator leader keeps its state fully in etcd. Followers forward all
// requests to the active leader.
package curator

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3/concurrency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// Config is the configuration of the curator.
type Config struct {
	// Etcd is an etcd client in which all curator storage and leader election
	// will be kept.
	Etcd client.Namespaced
	// NodeID is the ID of the node that this curator will run on. It's used to
	// populate the leader election lock.
	NodeID string
	// LeaderTTL is the timeout on the lease used to perform leader election.
	// Any active leader must continue updating its lease at least this often,
	// or the lease (and leadership) will be lost.
	// Lower values allow for faster failovers. Higher values allow for higher
	// resiliency against short network partitions.
	// A value less or equal to zero will default to 60 seconds.
	LeaderTTL time.Duration
	// Directory is the curator ephemeral directory in which the curator will
	// store its local domain socket for connections from the node.
	Directory         *localstorage.EphemeralCuratorDirectory
	ServerCredentials credentials.TransportCredentials
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
	status memory.Value
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

type electionStatusFollower struct {
	lock *ppb.LeaderElectionValue
}

func (s *Service) electionWatch() electionWatcher {
	return electionWatcher{
		Watcher: s.status.Watch(),
	}
}

// electionWatcher is a type-safe wrapper around event.Watcher which provides
// electionStatus values.
type electionWatcher struct {
	event.Watcher
}

// get retrieves an electionStatus from the electionWatcher.
func (w *electionWatcher) get(ctx context.Context) (*electionStatus, error) {
	val, err := w.Watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	status := val.(electionStatus)
	return &status, err
}

// buildLockValue returns a serialized etcd value that will be set by the
// instance when it becomes a leader. This value is a serialized
// LeaderElectionValue from private/storage.proto.
func (c *Config) buildLockValue(ttl int) ([]byte, error) {
	v := &ppb.LeaderElectionValue{
		NodeId: c.NodeID,
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

	// Establish a lease/session with etcd.
	session, err := concurrency.NewSession(s.config.Etcd.ThinClient(ctx),
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
		case o := <-observerC:
			var lock ppb.LeaderElectionValue
			if err := proto.Unmarshal(o.Kvs[0].Value, &lock); err != nil {
				return fmt.Errorf("parsing existing lock value failed: %w", err)
			}
			s.status.Set(electionStatus{
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
	s.status.Set(electionStatus{
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

func (s *Service) Run(ctx context.Context) error {
	// Start local election watcher. This logs what this curator knows about its own
	// leadership.
	go func() {
		w := s.electionWatch()
		for {
			s, err := w.get(ctx)
			if err != nil {
				supervisor.Logger(ctx).Warningf("Election watcher existing: get(): %w", err)
				return
			}
			if l := s.leader; l != nil {
				supervisor.Logger(ctx).Infof("Election watcher: this node's curator is leader (lock key %q, rev %d)", l.lockKey, l.lockRev)
			} else {
				supervisor.Logger(ctx).Infof("Election watcher: this node's curator is a follower")
			}
		}
	}()

	// Start listener. This is a gRPC service listening on a local socket,
	// providing the Curator API to consumers, dispatching to either a locally
	// running leader, or forwarding to a remotely running leader.
	lis := listener{
		directory:     s.config.Directory,
		publicCreds:   s.config.ServerCredentials,
		electionWatch: s.electionWatch,
		etcd:          s.config.Etcd,
		dispatchC:     make(chan dispatchRequest),
	}
	if err := supervisor.Run(ctx, "listener", lis.run); err != nil {
		return fmt.Errorf("when starting listener: %w", err)
	}

	// Calculate effective TTL. This replicates the behaviour of clientv3's WithTTL,
	// but allows us to explicitly log the used TTL.
	s.ttl = int(s.config.LeaderTTL.Seconds())
	if s.ttl <= 0 {
		s.ttl = 60
	}
	supervisor.Logger(ctx).Infof("Curator starting on prefix %q with lease TTL of %d seconds...", electionPrefix, s.ttl)

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	for {
		s.status.Set(electionStatus{})
		err := s.elect(ctx)
		s.status.Set(electionStatus{})

		if err != nil && errors.Is(err, ctx.Err()) {
			return fmt.Errorf("election round failed due to context cancelation, not attempting to re-elect: %w", err)
		}
		supervisor.Logger(ctx).Infof("Curator election round done: %v", err)
		supervisor.Logger(ctx).Info("Curator election restarting...")
	}
}

func (s *Service) DialCluster(ctx context.Context) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, fmt.Sprintf("unix://%s", s.config.Directory.ClientSocket.FullPath()), grpc.WithInsecure())
}
