// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package model

import "context"

// MetricValue is a prometheus-style labeled numerical metric value. In other
// words, it's a number accompanied by string key/value pairs.
type MetricValue struct {
	Count  int64
	Labels map[string]string
}

// WrapSimpleMetric turns a SQL model function which returns a single number into
// a function which returns one-length MetricValue list with no labels.
func WrapSimpleMetric(fn func(*Queries, context.Context) (int64, error)) func(*Queries, context.Context) ([]MetricValue, error) {
	return func(q *Queries, ctx context.Context) ([]MetricValue, error) {
		v, err := fn(q, ctx)
		if err != nil {
			return nil, err
		}
		return []MetricValue{
			{
				Count:  v,
				Labels: nil,
			},
		}, nil
	}
}

// A SQLMetricRow is a row that is the result of some kind of SQL 'metric query'.
// For each such query we define in our *.sql files, a corresponding
// implementation exists here.
type SQLMetricRow interface {
	Value() MetricValue
}

// Value implements SQLMetricRow for a row of the result of the
// CountActiveBackoffs SQL metric query.
func (c CountActiveBackoffsRow) Value() MetricValue {
	return MetricValue{
		Count: c.Count,
		Labels: map[string]string{
			"process": string(c.Process),
		},
	}
}

// Value implements SQLMetricRow for a row of the result of the
// CountActiveWork SQL metric query.
func (c CountActiveWorkRow) Value() MetricValue {
	return MetricValue{
		Count: c.Count,
		Labels: map[string]string{
			"process": string(c.Process),
		},
	}
}

// WrapLabeledMetric turns a SQL model function which returns a list of rows
// implementing SQLMetricRow into a function which returns a list of MetricValues
// with labels corresponding to the data returned in the rows.
func WrapLabeledMetric[M SQLMetricRow](fn func(*Queries, context.Context) ([]M, error)) func(*Queries, context.Context) ([]MetricValue, error) {
	return func(q *Queries, ctx context.Context) ([]MetricValue, error) {
		v, err := fn(q, ctx)
		if err != nil {
			return nil, err
		}
		res := make([]MetricValue, len(v))
		for i, vv := range v {
			res[i] = vv.Value()
		}
		return res, nil
	}
}
