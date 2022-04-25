package roleserve

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"fmt"
	"time"

	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/pki"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
	ppb "source.monogon.dev/metropolis/proto/private"
)

// workerControlPlane is the Control Plane Worker, responsible for maintaining a
// locally running Control Plane (Consensus and Curator service pair) if needed.
//
// The Control Plane will run under the following conditions:
//  - This node has been started in BOOTSTRAP mode and bootstrapData was provided
//    by the cluster enrolment logic. In this case, the Control Plane Worker will
//    perform the required bootstrap steps, creating a local node with appropriate
//    roles, and will start Consensus and the Curator.
//  - This node has the ConsensusMember Node Role. This will be true for nodes
//    which are REGISTERing into the cluster, as well as already running nodes that
//    have been assigned the role.
//
// In either case, ClusterMembership will be updated to allow connecting to the
// newly locally running control plane. For nodes that are bootstrapping the
// cluster, this will be the fist time the rest of the node can reach the
// Curator. For other cases, this will be the new, preferred way to reach
// consensus, without having to rely on external Control Plane nodes.
type workerControlPlane struct {
	storageRoot *localstorage.Root

	// bootstrapData will be read.
	bootstrapData *bootstrapDataValue
	// clusterMembership will be read and written.
	clusterMembership *ClusterMembershipValue
	// localRoles will be read.
	localRoles *localRolesValue
}

// controlPlaneStartup is used internally to provide a reduced (as in MapReduce)
// datum for the main Control Plane launcher responsible for launching the
// Control Plane Services, if at all.
type controlPlaneStartup struct {
	// consensusConfig is set if the node should run the control plane, and will
	// contain the configuration of the Consensus service.
	consensusConfig *consensus.Config
	// bootstrap is set if this node should bootstrap consensus. It contains all
	// data required to perform this bootstrap step.
	bootstrap *bootstrapData

	// existingMembership is the ClusterMembership that the node already had
	// available before deciding to run the Control Plane. This will be used to
	// carry over existing data from the membership into the new membership as
	// affected by starting the control plane.
	existingMembership *ClusterMembership
}

// changed informs the Control Plane launcher whether two different
// controlPlaneStartups differ to the point where a restart of the control plane
// should happen.
//
// Currently this is only true when a node switches to/from having a Control
// Plane role.
func (c *controlPlaneStartup) changed(o *controlPlaneStartup) bool {
	hasConsensusA := c.consensusConfig != nil
	hasConsensusB := o.consensusConfig != nil
	if hasConsensusA != hasConsensusB {
		return true
	}

	return false
}

func (s *workerControlPlane) run(ctx context.Context) error {
	// Map/Reduce a *controlPlaneStartup from different data sources. This will then
	// populate an Event Value that the actual launcher will use to start the
	// Control Plane.
	//
	//     bootstrapData -M-> bootstrapDataC ------.
	//                                             |
	// ClusterMambership -M-> clusterMembershipC --R---> startupV
	//                                             |
	//         NodeRoles -M-> rolesC --------------'
	//
	var startupV memory.Value

	// Channels are used as intermediaries between map stages and the final reduce,
	// which is okay as long as the entire tree restarts simultaneously (which we
	// ensure via RunGroup).
	bootstrapDataC := make(chan *bootstrapData)
	clusterMembershipC := make(chan *ClusterMembership)
	rolesC := make(chan *cpb.NodeRoles)

	supervisor.RunGroup(ctx, map[string]supervisor.Runnable{
		// Plain conversion from Event Value to channel.
		"map-bootstrap-data": func(ctx context.Context) error {
			w := s.bootstrapData.Watch()
			defer w.Close()
			for {
				v, err := w.Get(ctx)
				if err != nil {
					return err
				}
				bootstrapDataC <- v
			}
		},
		// Plain conversion from Event Value to channel.
		"map-cluster-membership": func(ctx context.Context) error {
			supervisor.Signal(ctx, supervisor.SignalHealthy)
			w := s.clusterMembership.Watch()
			defer w.Close()
			for {
				v, err := w.GetHome(ctx)
				if err != nil {
					return err
				}
				clusterMembershipC <- v
			}
		},
		// Plain conversion from Event Value to channel.
		"map-roles": func(ctx context.Context) error {
			supervisor.Signal(ctx, supervisor.SignalHealthy)
			w := s.localRoles.Watch()
			defer w.Close()
			for {
				v, err := w.Get(ctx)
				if err != nil {
					return err
				}
				rolesC <- v
			}
		},
		// Provide config from clusterMembership and roles.
		"reduce-config": func(ctx context.Context) error {
			supervisor.Signal(ctx, supervisor.SignalHealthy)
			var lr *cpb.NodeRoles
			var cm *ClusterMembership
			var bd *bootstrapData
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case lr = <-rolesC:
				case cm = <-clusterMembershipC:
				case bd = <-bootstrapDataC:
				}

				// If we have any bootstrap config ever, always use that.
				//
				// If there is a conflict between two available configuration methods (bootstrap
				// and non-bootstrap) there effectively shouldn't be any difference between the
				// two and it shouldn't matter which one we pick. That is because the bootstrap
				// data is only effectively used to populate the JoinCluster parameter of etcd,
				// which in turns is only used when a node is starting without any data present.
				// And since we managed to get our own node roles and that won the race against
				// bootstrap data, it means the bootstrap was successful and we can now start
				// without the bootstrap data.
				//
				// The only problem is when we remove a ConsensusMember from a node which still
				// has BootstrapData lingering from first bootup. However, we currently do not
				// support removing consensus roles (or any roles for that matter).
				//
				// TODO(q3k): support the above edge case. This can be done, for example, by
				// rewriting the reduction to wait for all data to be available and by
				// pre-populating all values to be nil at startup, thereby allowing for priority
				// encoding and removing the above race condition.
				if bd != nil {
					supervisor.Logger(ctx).Infof("Using bootstrap data...")
					startupV.Set(&controlPlaneStartup{
						consensusConfig: &consensus.Config{
							Data:           &s.storageRoot.Data.Etcd,
							Ephemeral:      &s.storageRoot.Ephemeral.Consensus,
							NodePrivateKey: bd.nodePrivateKey,
						},
						bootstrap: bd,
					})
					continue
				}

				// Otherwise, try to interpret node roles if available.
				if lr != nil && cm != nil {
					supervisor.Logger(ctx).Infof("Using role assigned by cluter...")
					role := lr.ConsensusMember
					if role == nil {
						supervisor.Logger(ctx).Infof("Not a control plane node.")
						startupV.Set(&controlPlaneStartup{})
						continue
					}
					supervisor.Logger(ctx).Infof("Control plane node, building config...")

					// Parse X509 data from NodeRoles.
					caCert, err := x509.ParseCertificate(role.CaCertificate)
					if err != nil {
						supervisor.Logger(ctx).Errorf("Could not parse CA certificate: %v", err)
						continue
					}
					peerCert, err := x509.ParseCertificate(role.PeerCertificate)
					if err != nil {
						supervisor.Logger(ctx).Errorf("Could not parse peer certificate: %v", err)
						continue
					}
					crl, err := x509.ParseCRL(role.InitialCrl)
					if err != nil {
						supervisor.Logger(ctx).Errorf("Could not parse CRL: %v", err)
						continue
					}

					// Convert NodeRoles peers into consensus peers. Let the user know what peers
					// we're starting with.
					supervisor.Logger(ctx).Infof("Node role mandates cluster membership with initial peers:")
					for _, p := range role.Peers {
						supervisor.Logger(ctx).Infof("  - %s (%s)", p.Name, p.URL)
					}
					nodes := make([]consensus.ExistingNode, len(role.Peers))
					for i, p := range role.Peers {
						nodes[i].Name = p.Name
						nodes[i].URL = p.URL
					}

					// Build and submit config to startup V.
					startupV.Set(&controlPlaneStartup{
						consensusConfig: &consensus.Config{
							Data:           &s.storageRoot.Data.Etcd,
							Ephemeral:      &s.storageRoot.Ephemeral.Consensus,
							NodePrivateKey: cm.credentials.TLSCredentials().PrivateKey.(ed25519.PrivateKey),
							JoinCluster: &consensus.JoinCluster{
								CACertificate:   caCert,
								NodeCertificate: peerCert,
								InitialCRL: &pki.CRL{
									Raw:  role.InitialCrl,
									List: crl,
								},
								ExistingNodes: nodes,
							},
						},
						existingMembership: cm,
					})
				}
			}
		},
	})

	// Run main Control Plane launcher. This depends on a config being put to
	// startupV.
	supervisor.Run(ctx, "launcher", func(ctx context.Context) error {
		supervisor.Logger(ctx).Infof("Waiting for start data...")

		// Read config from startupV.
		w := startupV.Watch()
		defer w.Close()
		startupI, err := w.Get(ctx)
		if err != nil {
			return err
		}
		startup := startupI.(*controlPlaneStartup)

		// Start Control Plane if we have a config.
		if startup.consensusConfig == nil {
			supervisor.Logger(ctx).Infof("No consensus config, not starting up control plane.")
		} else {
			supervisor.Logger(ctx).Infof("Got config, starting consensus and curator...")

			// Start consensus with config from startupV. This bootstraps the consensus
			// service if needed.
			con := consensus.New(*startup.consensusConfig)
			if err := supervisor.Run(ctx, "consensus", con.Run); err != nil {
				return fmt.Errorf("failed to start consensus service: %w", err)
			}

			// Prepare curator config, notably performing a bootstrap step if necessary. The
			// preparation will result in a set of node credentials to run the curator with
			// and a previously used cluster directory to be passed over to the new
			// ClusterMembership, if any.
			var creds *identity.NodeCredentials
			var caCert []byte
			var directory *cpb.ClusterDirectory
			if b := startup.bootstrap; b != nil {
				supervisor.Logger(ctx).Infof("Bootstrapping control plane. Waiting for consensus...")

				// Connect to etcd as curator to perform the bootstrap step.
				w := con.Watch()
				st, err := w.Get(ctx)
				if err != nil {
					return fmt.Errorf("while waiting for consensus for bootstrap: %w", err)
				}
				ckv, err := st.CuratorClient()
				if err != nil {
					return fmt.Errorf("when retrieving curator client for bootstarp: %w", err)
				}

				supervisor.Logger(ctx).Infof("Bootstrapping control plane. Performing bootstrap...")

				// Perform curator bootstrap step in etcd.
				//
				// This is all idempotent, so there's no harm in re-running this on every
				// curator startup.
				//
				// TODO(q3k): collapse the curator bootstrap shenanigans into a single function.
				npub := b.nodePrivateKey.Public().(ed25519.PublicKey)
				jpub := b.nodePrivateJoinKey.Public().(ed25519.PublicKey)
				n := curator.NewNodeForBootstrap(b.clusterUnlockKey, npub, jpub)
				n.EnableKubernetesWorker()

				var nodeCert []byte
				caCert, nodeCert, err = curator.BootstrapNodeFinish(ctx, ckv, &n, b.initialOwnerKey)
				if err != nil {
					return fmt.Errorf("while bootstrapping node: %w", err)
				}
				// ... and build new credentials from bootstrap step.
				creds, err = identity.NewNodeCredentials(b.nodePrivateKey, nodeCert, caCert)
				if err != nil {
					return fmt.Errorf("when creating bootstrap node credentials: %w", err)
				}
				supervisor.Logger(ctx).Infof("Control plane bootstrap complete, starting curator...")
			} else {
				// Not bootstrapping, just starting consensus with credentials we already have.

				// First, run a few assertions. This should never happen with the Map/Reduce
				// logic above, ideally we would encode this in the type system.
				if startup.existingMembership == nil {
					panic("no existingMembership but not bootstrapping either")
				}
				if startup.existingMembership.credentials == nil {
					panic("no existingMembership.credentials but not bootstrapping either")
				}
				if startup.existingMembership.remoteCurators == nil {
					panic("no existingMembership.remoteCurators but not bootstrapping either")
				}

				// Use already existing credentials, and pass over already known curators (as
				// we're not the only node, and we'd like downstream consumers to be able to
				// keep connecting to existing curators in case the local one fails).
				creds = startup.existingMembership.credentials
				directory = startup.existingMembership.remoteCurators
			}

			// Ensure this node is present in the cluster directory.
			if directory == nil {
				directory = &cpb.ClusterDirectory{}
			}
			missing := true
			for _, n := range directory.Nodes {
				if bytes.Equal(n.PublicKey, creds.PublicKey()) {
					missing = false
					break
				}
			}
			if missing {
				directory.Nodes = append(directory.Nodes, &cpb.ClusterDirectory_Node{
					PublicKey: creds.PublicKey(),
					Addresses: []*cpb.ClusterDirectory_Node_Address{
						{
							Host: "127.0.0.1",
						},
					},
				})
			}

			// Save this node's credentials, cluster directory and configuration as
			// part of the control plane bootstrap process.
			if b := startup.bootstrap; b != nil && caCert != nil {
				if err = creds.Save(&s.storageRoot.Data.Node.Credentials); err != nil {
					return fmt.Errorf("while saving node credentials: %w", err)
				}

				cdirRaw, err := proto.Marshal(directory)
				if err != nil {
					return fmt.Errorf("couldn't marshal ClusterDirectory: %w", err)
				}
				if err = s.storageRoot.ESP.Metropolis.ClusterDirectory.Write(cdirRaw, 0644); err != nil {
					return err
				}

				sc := ppb.SealedConfiguration{
					NodeUnlockKey: b.nodeUnlockKey,
					JoinKey:       b.nodePrivateJoinKey,
					ClusterCa:     caCert,
				}
				if err = s.storageRoot.ESP.Metropolis.SealedConfiguration.SealSecureBoot(&sc); err != nil {
					return err
				}

				supervisor.Logger(ctx).Infof("Saved bootstrapped node's credentials.")
				unix.Sync()
			}

			// Start curator.
			cur := curator.New(curator.Config{
				NodeCredentials: creds,
				Consensus:       con,
				LeaderTTL:       10 * time.Second,
			})
			if err := supervisor.Run(ctx, "curator", cur.Run); err != nil {
				return fmt.Errorf("failed to start curator: %w", err)
			}

			supervisor.Signal(ctx, supervisor.SignalHealthy)
			supervisor.Logger(ctx).Infof("Control plane running, submitting clusterMembership.")

			// We now have a locally running ControlPlane. Reflect that in a new
			// ClusterMembership.
			s.clusterMembership.set(&ClusterMembership{
				localConsensus: con,
				credentials:    creds,
				remoteCurators: directory,
				pubkey:         creds.PublicKey(),
			})
		}

		// Restart everything if we get a significantly different config (ie., a config
		// whose change would/should either turn up or tear down the Control Plane).
		//
		// Not restarting on every single change prevents us from going in a
		// ClusterMembership -> ClusterDirectory -> ClusterMembership thrashing loop.
		for {
			ncI, err := w.Get(ctx)
			if err != nil {
				return err
			}
			nc := ncI.(*controlPlaneStartup)
			if nc.changed(startup) {
				supervisor.Logger(ctx).Infof("Configuration changed, restarting...")
				return fmt.Errorf("config changed, restarting")
			}
		}
	})

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return ctx.Err()
}
