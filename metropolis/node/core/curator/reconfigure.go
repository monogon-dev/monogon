package curator

import (
	"strings"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	cpb "source.monogon.dev/metropolis/proto/common"
)

// reconfigureCluster does a three-way merge of a given cluster configuration
// (new, existing and optional base) into a merged configuration based on the
// fields set in the given fieldmask.
//
// An error will be returned if the mask contains paths that reference unknown
// fields or references fields which cannot be changed.
func reconfigureCluster(base, new, existing *cpb.ClusterConfiguration, mask *fieldmaskpb.FieldMask) (*cpb.ClusterConfiguration, error) {
	if new == nil {
		return nil, status.Error(codes.InvalidArgument, "new_config must be set")
	}
	if mask == nil {
		return nil, status.Error(codes.InvalidArgument, "update_mask must be set")
	}

	mask.Normalize()
	if !mask.IsValid(new) {
		return nil, status.Error(codes.InvalidArgument, "update_mask is invalid for new_config")
	}

	if base != nil {
		if !mask.IsValid(base) {
			return nil, status.Error(codes.InvalidArgument, "update_mask is invalid for base_config")
		}
	}

	// Merged proto, start with deep copy of existing configuration.
	merged := proto.Clone(existing).(*cpb.ClusterConfiguration)

	for _, path := range mask.Paths {
		handled, err := reconfigureKubernetes(base, new, existing, merged, path)
		if err != nil {
			return nil, err
		}
		if !handled {
			return nil, status.Errorf(codes.InvalidArgument, "cannot modify %s", path)
		}
	}

	return merged, nil
}

// reconfigureKubernetes does a three-way merge of Kubernetes configuration
// (new, existing and optional base) of a given protobuf field path into merged.
//
// An error is returned if there is an issue applying the given change.
// Otherwise, a boolean value is returned, indicating whether this given filed
// path was handled.
func reconfigureKubernetes(base, new, existing, merged *cpb.ClusterConfiguration, path string) (bool, error) {
	if path == "kubernetes" {
		return false, status.Error(codes.InvalidArgument, "cannot mutate kubernetes directly, only subfields")
	}
	if !strings.HasPrefix(path, "kubernetes.") {
		return false, nil
	}

	// Kubernetes should always exist in stored configs.
	if merged.Kubernetes == nil {
		merged.Kubernetes = &cpb.ClusterConfiguration_Kubernetes{}
	}
	if existing.Kubernetes == nil {
		existing.Kubernetes = &cpb.ClusterConfiguration_Kubernetes{}
	}

	// Check Kubernetes in user structs.
	if new.Kubernetes == nil {
		return false, status.Errorf(codes.InvalidArgument, "cannot reference field %s in new_config", path)
	}
	if base != nil && base.Kubernetes == nil {
		return false, status.Errorf(codes.InvalidArgument, "cannot reference field %s in old_config", path)
	}

	switch path {
	case "kubernetes.node_labels_to_synchronize":
		if base != nil && cmp.Diff(base.Kubernetes.NodeLabelsToSynchronize, existing.Kubernetes.NodeLabelsToSynchronize, protocmp.Transform()) != "" {
			return false, status.Error(codes.FailedPrecondition, "base_config.kubernetes.node_labels_to_synchronize different from current value")
		}
		merged.Kubernetes.NodeLabelsToSynchronize = new.Kubernetes.NodeLabelsToSynchronize
	default:
		return false, status.Errorf(codes.InvalidArgument, "cannot mutate %s", path)
	}

	return true, nil
}
