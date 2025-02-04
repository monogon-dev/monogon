// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package kubernetes

import (
	"testing"

	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/component-base/featuregate"
)

func TestFeatureGateDefaults(t *testing.T) {
	for f, en := range extraFeatureGates {
		if utilfeature.DefaultFeatureGate.Enabled(f) == en {
			t.Errorf("Feature gate %q is already %v by default, remove it from extraFeatureGates", string(f), en)
		}
	}
}

func TestAsFlags(t *testing.T) {
	for _, c := range []struct {
		name     string
		fg       featureGates
		expected string
	}{
		{"None", featureGates{}, "--feature-gates="},
		{"Single", featureGates{featuregate.Feature("Test"): true}, "--feature-gates=Test=true"},
		{"Multiple", featureGates{featuregate.Feature("Test"): true, featuregate.Feature("Test2"): false}, "--feature-gates=Test=true,Test2=false"},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := c.fg.AsFlag()
			if got != c.expected {
				t.Errorf("Expected %q, got %q", c.expected, got)
			}
		})
	}
}
