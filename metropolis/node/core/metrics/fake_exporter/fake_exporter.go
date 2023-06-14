// fake_exporter is a tiny Prometheus-compatible metrics exporter which exports a
// single metric with a value configured at startup. It is used to test the
// metrics service.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	flagListen string
	flagValue  int
)

func main() {
	flag.StringVar(&flagListen, "listen", ":8080", "address to listen on")
	flag.IntVar(&flagValue, "value", 1234, "value of 'test' metric to serve")
	flag.Parse()

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test %d\n", flagValue)
	})
	log.Printf("Listening on %s", flagListen)
	http.ListenAndServe(flagListen, nil)
}
