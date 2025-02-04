// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// httpserver serves a test HTTP endpoint for E2E testing.
package main

import (
	"net/http"
	"os"
)

func main() {
	nodeName := os.Getenv("NODE_NAME")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Node-Name", nodeName)
		w.Header().Set("X-Remote-IP", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		// Send a big chunk to root out MTU/MSS issues.
		testPayload := make([]byte, 2000)
		w.Write(testPayload)
	}))
}
