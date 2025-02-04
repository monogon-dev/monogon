// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package scruffy

import (
	"github.com/prometheus/client_golang/prometheus"

	"source.monogon.dev/go/algorithm/cartesian"
)

// A labelDefinition describes a key/value pair that's a metric dimension. It
// consists of the label key/name (a string), and a list of possible values of
// this key. The list of values will be used to initialize the metrics at startup
// with zero values.
//
// The initialValues system is intended to be used with labels that are
// low-cardinality enums, e.g. the name of a subsystem.
//
// All labelDefinitions for a single metric will then create a cartesian product
// of all initialValues.
type labelDefinition struct {
	// name/key of the label.
	name string
	// initialValues defines the default values for this label key/name that will be
	// used to generate a list of initial zero-filled metrics.
	initialValues []string
}

// labelDefinitions is a list of labelDefinition which define the label
// dimensions of a metric. All the initialValues of the respective
// labelDefinitions will create a cartesian set of default zero-filled metric
// values when the metric susbsystem gets initialized. These zero values will
// then get overridden by real data as it is collected.
type labelDefinitions []labelDefinition

// initialLabels generates the list of initial labels key/values that should be
// used to generate zero-filled metrics on startup. This is a cartesian product
// of all initialValues of all labelDefinitions.
func (l labelDefinitions) initialLabels() []prometheus.Labels {
	// Nothing to do if this is an empty labelDefinitions.
	if len(l) == 0 {
		return nil
	}

	// Given:
	//
	// labelDefinitions := []labelDefinition{
	//    { name: "a", initialValues: []string{"foo", "bar"}},
	//    { name: "b", initialValues: []string{"baz", "barf"}},
	// }
	//
	// This creates:
	//
	// values := []string{
	//    { "foo", "bar" }, // label 'a'
	//    { "baz", "barf" }, // label 'b'
	// }
	var values [][]string
	for _, ld := range l {
		values = append(values, ld.initialValues)
	}

	// Given the above:
	//
	// valuesProduct := []string{
	//    //  a      b
	//    { "foo", "baz" },
	//    { "foo", "barf" },
	//    { "bar", "baz" },
	//    { "bar", "barf" },
	// }
	valuesProduct := cartesian.Product[string](values...)

	// This converts valuesProduct into an actual prometheus-compatible type,
	// re-attaching the label names back into the columns as seen above.
	var res []prometheus.Labels
	for _, vp := range valuesProduct {
		labels := make(prometheus.Labels)
		for i, lv := range vp {
			labelDef := l[i]
			labels[labelDef.name] = lv
		}
		res = append(res, labels)
	}
	return res
}

// names returns the keys/names of all the metric labels from these
// labelDefinitions.
func (l labelDefinitions) names() []string {
	var res []string
	for _, ld := range l {
		res = append(res, ld.name)
	}
	return res
}
