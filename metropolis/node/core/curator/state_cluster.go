package curator

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/rpc"
	cpb "source.monogon.dev/metropolis/proto/common"
)

var (
	clusterConfigurationKey = "/cluster/configuration"
)

// Cluster is the cluster's configuration, as (un)marshaled to/from
// common.ClusterConfiguration.
type Cluster struct {
	TPMMode cpb.ClusterConfiguration_TPMMode
}

// DefaultClusterConfiguration is the default cluster configuration for a newly
// bootstrapped cluster if no initial cluster configuration was specified by the
// user.
func DefaultClusterConfiguration() *Cluster {
	return &Cluster{
		TPMMode: cpb.ClusterConfiguration_TPM_MODE_REQUIRED,
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

func clusterUnmarshal(data []byte) (*Cluster, error) {
	msg := cpb.ClusterConfiguration{}
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("could not unmarshal proto: %w", err)
	}
	return clusterFromProto(&msg)
}

func clusterFromProto(cc *cpb.ClusterConfiguration) (*Cluster, error) {
	switch cc.TpmMode {
	case cpb.ClusterConfiguration_TPM_MODE_REQUIRED:
	case cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT:
	case cpb.ClusterConfiguration_TPM_MODE_DISABLED:
	default:
		return nil, fmt.Errorf("invalid TpmMode: %v", cc.TpmMode)
	}

	c := &Cluster{
		TPMMode: cc.TpmMode,
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
	return &cpb.ClusterConfiguration{
		TpmMode: c.TPMMode,
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
