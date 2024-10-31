package curator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	cpb "source.monogon.dev/metropolis/proto/common"
)

func TestReconfigureCluster(t *testing.T) {
	mkCfg := func(regexes ...string) *cpb.ClusterConfiguration {
		res := &cpb.ClusterConfiguration{
			TpmMode:               cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT,
			StorageSecurityPolicy: cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE,
			Kubernetes:            &cpb.ClusterConfiguration_Kubernetes{},
		}
		for _, regex := range regexes {
			res.Kubernetes.NodeLabelsToSynchronize = append(res.Kubernetes.NodeLabelsToSynchronize, &cpb.ClusterConfiguration_Kubernetes_NodeLabelsToSynchronize{
				Regexp: regex,
			})
		}
		return res
	}

	for i, te := range []struct {
		base       *cpb.ClusterConfiguration
		new        *cpb.ClusterConfiguration
		existing   *cpb.ClusterConfiguration
		mask       *fieldmaskpb.FieldMask
		result     *cpb.ClusterConfiguration
		shouldFail bool
	}{
		// Case 0: no-op on an empty config.
		{
			base:     &cpb.ClusterConfiguration{},
			new:      &cpb.ClusterConfiguration{},
			existing: &cpb.ClusterConfiguration{},
			mask:     &fieldmaskpb.FieldMask{Paths: []string{}},
			result:   &cpb.ClusterConfiguration{},
		},
		// Case 1: no-op with an empty base.
		{
			new:      &cpb.ClusterConfiguration{},
			existing: &cpb.ClusterConfiguration{},
			mask:     &fieldmaskpb.FieldMask{Paths: []string{}},
			result:   &cpb.ClusterConfiguration{},
		},
		// Case 2: no-op on a populated config.
		{
			base:     &cpb.ClusterConfiguration{},
			new:      &cpb.ClusterConfiguration{},
			existing: mkCfg("^foo$"),
			mask:     &fieldmaskpb.FieldMask{Paths: []string{}},
			result:   mkCfg("^foo$"),
		},
		// Case 3: reconfigure kubernetes node labels.
		{
			base: &cpb.ClusterConfiguration{
				Kubernetes: &cpb.ClusterConfiguration_Kubernetes{
					NodeLabelsToSynchronize: []*cpb.ClusterConfiguration_Kubernetes_NodeLabelsToSynchronize{
						{Regexp: "^foo$"},
					},
				},
			},
			new: &cpb.ClusterConfiguration{
				Kubernetes: &cpb.ClusterConfiguration_Kubernetes{
					NodeLabelsToSynchronize: []*cpb.ClusterConfiguration_Kubernetes_NodeLabelsToSynchronize{
						{Regexp: "^bar$"},
					},
				},
			},
			existing: mkCfg("^foo$"),
			mask:     &fieldmaskpb.FieldMask{Paths: []string{"kubernetes.node_labels_to_synchronize"}},
			result:   mkCfg("^bar$"),
		},
		// Case 4: reconfigure kubernetes node labels without base.
		{
			new: &cpb.ClusterConfiguration{
				Kubernetes: &cpb.ClusterConfiguration_Kubernetes{
					NodeLabelsToSynchronize: []*cpb.ClusterConfiguration_Kubernetes_NodeLabelsToSynchronize{
						{Regexp: "^bar$"},
					},
				},
			},
			existing: mkCfg("^foo$"),
			mask:     &fieldmaskpb.FieldMask{Paths: []string{"kubernetes.node_labels_to_synchronize"}},
			result:   mkCfg("^bar$"),
		},
		// Case 5: no-op with an empty base.
		{
			new:      &cpb.ClusterConfiguration{},
			existing: &cpb.ClusterConfiguration{},
			mask:     &fieldmaskpb.FieldMask{Paths: []string{}},
			result:   &cpb.ClusterConfiguration{},
		},
		// Case 6: missing new.
		{
			new:        nil,
			existing:   &cpb.ClusterConfiguration{},
			mask:       &fieldmaskpb.FieldMask{Paths: []string{}},
			result:     &cpb.ClusterConfiguration{},
			shouldFail: true,
		},
		// Case 7: missing mask.
		{
			new:        &cpb.ClusterConfiguration{},
			existing:   &cpb.ClusterConfiguration{},
			mask:       nil,
			result:     &cpb.ClusterConfiguration{},
			shouldFail: true,
		},
		// Case 8: mask references unknown field.
		{
			new:        &cpb.ClusterConfiguration{},
			existing:   &cpb.ClusterConfiguration{},
			mask:       &fieldmaskpb.FieldMask{Paths: []string{"foo"}},
			result:     &cpb.ClusterConfiguration{},
			shouldFail: true,
		},
		// Case 9: mask references field unset in new.
		{
			new:        &cpb.ClusterConfiguration{},
			existing:   &cpb.ClusterConfiguration{},
			mask:       &fieldmaskpb.FieldMask{Paths: []string{"kubernetes.node_labels_to_synchronize"}},
			result:     &cpb.ClusterConfiguration{},
			shouldFail: true,
		},
		// Case 10: mask references field unset in base.
		{
			base: &cpb.ClusterConfiguration{},
			new: &cpb.ClusterConfiguration{
				Kubernetes: &cpb.ClusterConfiguration_Kubernetes{
					NodeLabelsToSynchronize: []*cpb.ClusterConfiguration_Kubernetes_NodeLabelsToSynchronize{},
				},
			},
			existing:   &cpb.ClusterConfiguration{},
			mask:       &fieldmaskpb.FieldMask{Paths: []string{"kubernetes.node_labels_to_synchronize"}},
			result:     &cpb.ClusterConfiguration{},
			shouldFail: true,
		},
	} {
		{
			got, err := reconfigureCluster(te.base, te.new, te.existing, te.mask)
			if !te.shouldFail {
				if err != nil {
					t.Errorf("Case %d: %v", i, err)
					continue
				}
				if diff := cmp.Diff(te.result, got, protocmp.Transform()); diff != "" {
					t.Errorf("Case %d: %s", i, diff)
				}
			} else {
				if err == nil {
					t.Errorf("Case %d: should've failed, got success", i)
				}
			}
		}
	}
}
