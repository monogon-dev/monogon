// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package wrapngo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/packethost/packngo"
	"k8s.io/klog/v2"
)

// wrap a given fn in some reliability-increasing duct tape: context support and
// exponential backoff retries for intermittent connectivity issues. This allows
// us to use packngo code instead of writing our own API stub for Equinix Metal.
//
// The given fn will be retried until it returns a 'permanent' Equinix error (see
// isPermanentEquinixError) or the given context expires. Additionally, fn will
// be called with a brand new packngo client tied to the context of the wrap
// call. Finally, the given client will also have some logging middleware
// attached to it which can be activated by setting verbosity 5 (or greater) on
// this file.
//
// The wrapped fn can be either just a plain packngo method or some complicated
// idempotent logic, as long as it cooperates with the above contract.
func wrap[U any](ctx context.Context, cl *client, fn func(*packngo.Client) (U, error)) (U, error) {
	var zero U
	if err := cl.serializer.up(ctx); err != nil {
		return zero, err
	}
	defer cl.serializer.down()

	bc := backoff.WithContext(cl.o.BackOff(), ctx)
	pngo, err := cl.clientForContext(ctx)
	if err != nil {
		// Generally this shouldn't happen other than with programming errors, so we
		// don't back this off.
		return zero, fmt.Errorf("could not crate equinix client: %w", err)
	}

	var res U
	err = backoff.Retry(func() error {
		res, err = fn(pngo)
		if isPermanentEquinixError(err) {
			return backoff.Permanent(err)
		}
		return err
	}, bc)
	if err != nil {
		return zero, err
	}
	return res, nil
}

type injectContextRoundTripper struct {
	ctx      context.Context
	original http.RoundTripper
	metrics  *metricsSet
}

func (r *injectContextRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	klog.V(5).Infof("Request -> %v", req.URL.String())
	start := time.Now()
	res, err := r.original.RoundTrip(req.WithContext(r.ctx))
	latency := time.Since(start)
	r.metrics.onAPIRequestDone(req, res, err, latency)

	if err != nil {
		klog.V(5).Infof("HTTP error <- %v", err)
	} else {
		klog.V(5).Infof("Response <- %v", res.Status)
	}
	return res, err
}

func (c *client) clientForContext(ctx context.Context) (*packngo.Client, error) {
	httpcl := &http.Client{
		Transport: &injectContextRoundTripper{
			ctx:      ctx,
			original: http.DefaultTransport,
			metrics:  c.metrics,
		},
	}
	return packngo.NewClient(packngo.WithAuth(c.username, c.token), packngo.WithHTTPClient(httpcl))
}

// httpStatusCode extracts the status code from error values returned by
// packngo methods.
func httpStatusCode(err error) int {
	var er *packngo.ErrorResponse
	if err != nil && errors.As(err, &er) {
		return er.Response.StatusCode
	}
	return -1
}

// IsNotFound returns true if the given error is an Equinix packngo/wrapngo 'not
// found' error.
func IsNotFound(err error) bool {
	return httpStatusCode(err) == http.StatusNotFound
}

func isPermanentEquinixError(err error) bool {
	// Invalid argument/state errors from wrapping.
	if errors.Is(err, ErrRaceLost) {
		return true
	}
	if errors.Is(err, ErrNoReservationProvided) {
		return true
	}
	// Real errors returned from equinix.
	st := httpStatusCode(err)
	switch st {
	case http.StatusUnauthorized:
		return true
	case http.StatusForbidden:
		return true
	case http.StatusNotFound:
		return true
	case http.StatusUnprocessableEntity:
		return true
	}
	return false
}
