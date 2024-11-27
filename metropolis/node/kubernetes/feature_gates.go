package kubernetes

import (
	"fmt"
	"strings"

	"k8s.io/component-base/featuregate"
	"k8s.io/kubernetes/pkg/features"
)

type featureGates map[featuregate.Feature]bool

// AsFlag returns the feature gates as a --feature-gate flag.
func (fgs featureGates) AsFlag() string {
	var strb strings.Builder
	strb.WriteString("--feature-gates=")
	i := 0
	for f, en := range fgs {
		fmt.Fprintf(&strb, "%s=%v", string(f), en)
		if i++; i != len(fgs) {
			strb.WriteByte(',')
		}
	}
	return strb.String()
}

// AsConfigObject returns the feature gates as a plain map for K8s configs.
func (fgs featureGates) AsMap() map[string]bool {
	out := make(map[string]bool)
	for f, en := range fgs {
		out[string(f)] = en
	}
	return out
}

var extraFeatureGates = featureGates{
	features.UserNamespacesSupport:              true,
	features.UserNamespacesPodSecurityStandards: true,
}
