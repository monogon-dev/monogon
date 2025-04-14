// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/cenkalti/backoff/v4"
)

type tokenBody struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
}

// handleUnauthorized implements token authentication based on this
// specification: https://distribution.github.io/distribution/spec/auth/token/
//
// The registry will return Unauthorized if a token is required and it is
// missing or invalid (e.g. expired). We then need to ask the authorization
// service for a token, and retry the original request with the new token.
//
// Some registries (e.g. Docker Hub and ghcr.io) require a token even for public
// repositories. In this case, the authorization service returns tokens without
// requiring any credentials.
func (c *Client) handleUnauthorized(ctx context.Context, resp *http.Response) (retry bool, err error) {
	// Check if we have a Bearer challenge.
	challenges := parseAuthenticateHeader(resp.Header.Values("Www-Authenticate"))
	var params map[string]string
	for _, c := range challenges {
		if strings.EqualFold(c.scheme, "bearer") {
			params = c.params
			break
		}
	}
	realm := params["realm"]
	if realm == "" {
		// There is no challenge, return the original HTTP error.
		return false, nil
	}

	// Construct token URL.
	tokenURL, err := url.Parse(realm)
	if err != nil {
		return false, backoff.Permanent(fmt.Errorf("failed to parse realm: %w", err))
	}
	query := tokenURL.Query()
	service := params["service"]
	if service != "" {
		query.Set("service", service)
	}
	for scope := range strings.SplitSeq(params["scope"], " ") {
		if scope != "" {
			query.Add("scope", scope)
		}
	}
	tokenURL.RawQuery = query.Encode()

	// Do token request.
	req, err := http.NewRequestWithContext(ctx, "GET", tokenURL.String(), nil)
	if err != nil {
		return false, err
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	client := http.Client{Transport: c.Transport}
	tokenResp, err := client.Do(req)
	if err != nil {
		return false, redactURLError(err)
	}
	if tokenResp.StatusCode != http.StatusOK {
		return false, readClientError(tokenResp, req)
	}
	defer tokenResp.Body.Close()

	// Parse token response.
	bodyBytes, err := readFullBody(tokenResp, 1024*1024)
	if err != nil {
		return false, err
	}
	body := tokenBody{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return false, backoff.Permanent(fmt.Errorf("failed to parse token response: %w", err))
	}
	token := body.Token
	if token == "" {
		token = body.AccessToken
	}
	if token == "" {
		return false, backoff.Permanent(fmt.Errorf("missing token in token response"))
	}

	c.authMu.Lock()
	c.bearerToken = token
	c.authMu.Unlock()
	return true, nil
}

func (c *Client) addAuthorization(req *http.Request) {
	c.authMu.RLock()
	defer c.authMu.RUnlock()
	if c.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	}
}
