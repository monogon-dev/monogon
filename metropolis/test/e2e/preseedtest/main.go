// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello world from preseeded image")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world from preseeded image\n")
	})
	http.HandleFunc("/ready_userns", func(w http.ResponseWriter, r *http.Request) {
		uidMapRaw, err := os.ReadFile("/proc/self/uid_map")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		uidMapFields := strings.Fields(string(uidMapRaw))
		if len(uidMapFields) != 3 {
			http.Error(w, fmt.Sprintf("Bad uid_map contents, not 3 fields: %q", string(uidMapRaw)), http.StatusInternalServerError)
			return
		}
		startId, err := strconv.ParseUint(uidMapFields[1], 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("while parsing start ID: %v", err), http.StatusInternalServerError)
			return
		}
		if startId == 0 {
			http.Error(w, "Not in a non-initial user namespace, UID space starts at 0", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Hello world from a user namespace\n")
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Serve failed: %v\n", err)
	}
}
