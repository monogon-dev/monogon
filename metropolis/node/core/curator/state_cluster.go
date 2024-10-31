package curator

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/rpc"

	cpb "source.monogon.dev/metropolis/proto/common"
)

var (
	clusterConfigurationKey = "/cluster/configuration"
)

// Cluster is the cluster's configuration, as (un)marshaled to/from
// common.ClusterConfiguration.
type Cluster struct {
	ClusterDomain                       string
	TPMMode                             cpb.ClusterConfiguration_TPMMode
	StorageSecurityPolicy               cpb.ClusterConfiguration_StorageSecurityPolicy
	NodeLabelsToSynchronizeToKubernetes []*cpb.ClusterConfiguration_Kubernetes_NodeLabelsToSynchronize
}

// DefaultClusterConfiguration is the default cluster configuration for a newly
// bootstrapped cluster if no initial cluster configuration was specified by the
// user.
func DefaultClusterConfiguration() *Cluster {
	return &Cluster{
		ClusterDomain:                       "cluster.internal",
		TPMMode:                             cpb.ClusterConfiguration_TPM_MODE_REQUIRED,
		StorageSecurityPolicy:               cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION,
		NodeLabelsToSynchronizeToKubernetes: nil,
	}
}

// ClusterConfigurationFromInitial converts a user-provided initial cluster
// configuration proto into a Cluster, checking that the provided values are
// valid.
func ClusterConfigurationFromInitial(icc *cpb.ClusterConfiguration) (*Cluster, error) {
	// clusterFromProto performs type checks.
	return clusterFromProto(icc)
}

// NodeShouldUseTPM returns whether a node should use a TPM or not for this Cluster
// and a given node's TPM availability.
//
// A user-facing error is returned if the combination of local cluster policy and
// node TPM availability is invalid.
func (c *Cluster) NodeShouldUseTPM(available bool) (bool, error) {
	switch c.TPMMode {
	case cpb.ClusterConfiguration_TPM_MODE_DISABLED:
		return false, nil
	case cpb.ClusterConfiguration_TPM_MODE_REQUIRED:
		if !available {
			return false, fmt.Errorf("TPM required but not available")
		}
		return true, nil
	case cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT:
		return available, nil
	default:
		return false, fmt.Errorf("invalid TPM mode")
	}
}

// NodeTPMUsage returns the NodeTPMUsage (whether a node should use a TPM or not
// plus information whether it has a TPM in the first place) for this Cluster and
// a given node's TPM availability.
//
// A user-facing error is returned if the combination of local cluster policy and
// node TPM availability is invalid.
func (c *Cluster) NodeTPMUsage(have bool) (usage cpb.NodeTPMUsage, err error) {
	var use bool
	use, err = c.NodeShouldUseTPM(have)
	if err != nil {
		return
	}
	switch {
	case have && use:
		usage = cpb.NodeTPMUsage_NODE_TPM_PRESENT_AND_USED
	case have && !use:
		usage = cpb.NodeTPMUsage_NODE_TPM_PRESENT_BUT_UNUSED
	case !have:
		usage = cpb.NodeTPMUsage_NODE_TPM_NOT_PRESENT
	}
	return
}

// NodeStorageSecurity returns the recommended NodeStorageSecurity for nodes
// joining the cluster.
func (c *Cluster) NodeStorageSecurity() (security cpb.NodeStorageSecurity, err error) {
	switch c.StorageSecurityPolicy {
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE:
		// TODO(q3k): allow per-node configuration. Be conservative for now.
		return cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_AUTHENTICATED_ENCRYPTED, nil
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION:
		return cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_AUTHENTICATED_ENCRYPTED, nil
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION:
		// TODO(q3k): allow per-node configuration. Be conservative for now.
		return cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_ENCRYPTED, nil
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE:
		return cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_INSECURE, nil
	default:
		return cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_INVALID, fmt.Errorf("invalid cluster storage policy %d", c.StorageSecurityPolicy)
	}
}

// ValidateNodeStorage checks the given NodeStorageSecurity and returns a gRPC
// status if the security setting is not compliant with the cluster node storage
// policy.
func (c *Cluster) ValidateNodeStorage(ns cpb.NodeStorageSecurity) error {
	switch c.StorageSecurityPolicy {
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE:
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE:
		if ns != cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_INSECURE {
			return status.Error(codes.FailedPrecondition, "cluster policy requires insecure node storage")
		}
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION:
		switch ns {
		case cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_AUTHENTICATED_ENCRYPTED:
		case cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_ENCRYPTED:
		default:
			return status.Error(codes.FailedPrecondition, "cluster policy requires encrypted node storage")
		}
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION:
		if ns != cpb.NodeStorageSecurity_NODE_STORAGE_SECURITY_AUTHENTICATED_ENCRYPTED {
			return status.Error(codes.FailedPrecondition, "cluster policy requires encrypted and authenticated node storage")
		}
	default:
		return status.Error(codes.Internal, "cannot interpret cluster node storage policy")
	}
	return nil
}

func clusterUnmarshal(data []byte) (*Cluster, error) {
	var msg cpb.ClusterConfiguration
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("could not unmarshal proto: %w", err)
	}
	if msg.ClusterDomain == "" {
		// Backward compatibility for clusters which did not have this field
		// initially.
		msg.ClusterDomain = "cluster.internal"
	}
	return clusterFromProto(&msg)
}

func clusterFromProto(cc *cpb.ClusterConfiguration) (*Cluster, error) {
	if err := common.ValidateClusterDomain(cc.ClusterDomain); err != nil {
		return nil, fmt.Errorf("invalid ClusterDomain: %w", err)
	}

	switch cc.TpmMode {
	case cpb.ClusterConfiguration_TPM_MODE_REQUIRED:
	case cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT:
	case cpb.ClusterConfiguration_TPM_MODE_DISABLED:
	default:
		return nil, fmt.Errorf("invalid TpmMode: %v", cc.TpmMode)
	}

	switch cc.StorageSecurityPolicy {
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE:
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION:
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION:
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE:
	default:
		return nil, fmt.Errorf("invalid StorageSecurityPolicy: %v", cc.StorageSecurityPolicy)
	}

	c := &Cluster{
		ClusterDomain:         cc.ClusterDomain,
		TPMMode:               cc.TpmMode,
		StorageSecurityPolicy: cc.StorageSecurityPolicy,
	}
	if kc := cc.Kubernetes; kc != nil {
		c.NodeLabelsToSynchronizeToKubernetes = kc.NodeLabelsToSynchronize
	}

	return c, nil
}

func (c *Cluster) proto() (*cpb.ClusterConfiguration, error) {
	switch c.TPMMode {
	case cpb.ClusterConfiguration_TPM_MODE_REQUIRED:
	case cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT:
	case cpb.ClusterConfiguration_TPM_MODE_DISABLED:
	default:
		return nil, fmt.Errorf("invalid TPMMode %d", c.TPMMode)
	}

	switch c.StorageSecurityPolicy {
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE:
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION:
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION:
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE:
	default:
		return nil, fmt.Errorf("invalid StorageSecurityPolicy %d", c.StorageSecurityPolicy)
	}

	return &cpb.ClusterConfiguration{
		ClusterDomain:         c.ClusterDomain,
		TpmMode:               c.TPMMode,
		StorageSecurityPolicy: c.StorageSecurityPolicy,
		Kubernetes: &cpb.ClusterConfiguration_Kubernetes{
			NodeLabelsToSynchronize: c.NodeLabelsToSynchronizeToKubernetes,
		},
	}, nil
}

func clusterLoad(ctx context.Context, l *leadership) (*Cluster, error) {
	rpc.Trace(ctx).Printf("loadCluster...")
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(clusterConfigurationKey))
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return nil, rpcErr
		}
		rpc.Trace(ctx).Printf("could not retrieve cluster configuartion: %v", err)
		return nil, status.Errorf(codes.Unavailable, "could not retrieve cluster configuration: %v", err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	rpc.Trace(ctx).Printf("loadCluster: %d KVs", len(kvs))
	if len(kvs) != 1 {
		return nil, errNodeNotFound
	}
	node, err := clusterUnmarshal(kvs[0].Value)
	if err != nil {
		rpc.Trace(ctx).Printf("could not unmarshal cluster: %v", err)
		return nil, status.Errorf(codes.Unavailable, "could not unmarshal cluster")
	}
	rpc.Trace(ctx).Printf("loadCluster: unmarshal ok")
	return node, nil
}

func clusterSave(ctx context.Context, l *leadership, c *Cluster) error {
	rpc.Trace(ctx).Printf("clusterSave...")
	clusterProto, err := c.proto()
	if err != nil {
		rpc.Trace(ctx).Printf("could not convert updated cluster configuration: %v", err)
		return status.Errorf(codes.Unavailable, "could not convert updated cluster")
	}
	clusterBytes, err := proto.Marshal(clusterProto)
	if err != nil {
		rpc.Trace(ctx).Printf("could not marshal updated cluster configuration: %v", err)
		return status.Errorf(codes.Unavailable, "could not marshal updated cluster")
	}

	ocs := clientv3.OpPut(clusterConfigurationKey, string(clusterBytes))
	_, err = l.txnAsLeader(ctx, ocs)
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return rpcErr
		}
		rpc.Trace(ctx).Printf("could not save updated cluster configuration: %v", err)
		return status.Error(codes.Unavailable, "could not save updated cluster configuration")
	}
	rpc.Trace(ctx).Printf("clusterSave: write ok")
	return nil
}
