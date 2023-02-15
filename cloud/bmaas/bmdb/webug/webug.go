// Package webug implements a web-based debug/troubleshooting/introspection
// system for the BMDB. It's optimized for use by developers and trained
// operators, prioritizing information density, fast navigation and heavy
// interlinking.
package webug

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/reflection"
)

var (
	//go:embed templates/*html
	templateFS embed.FS
	templates  = template.Must(template.New("base.html").Funcs(templateFuncs).ParseFS(templateFS, "templates/*"))
)

// server holds the state of an active webug interface.
type server struct {
	// connection pool to BMDB.
	conn *bmdb.Connection
	// schema retrieved from BMDB.
	schema *reflection.Schema
	// muSchema locks schema for updates.
	muSchema sync.RWMutex
}

// curSchema returns the current cached BMDB schema.
func (s *server) curSchema() *reflection.Schema {
	s.muSchema.Lock()
	defer s.muSchema.Unlock()
	return s.schema
}

// schemaWorker runs a background goroutine which attempts to update the server's
// cached BMDB schema every hour.
func (s *server) schemaWorker(ctx context.Context) {
	t := time.NewTicker(time.Hour)
	defer t.Stop()

	for {
		// Wait for the timer to tick, or for the context to expire.
		select {
		case <-ctx.Done():
			klog.Infof("Schema fetch worker exiting: %v", ctx.Err())
			return
		case <-t.C:
		}

		// Time to check the schema. Do that in an exponential backoff loop until
		// successful.
		bo := backoff.NewExponentialBackOff()
		bo.MaxElapsedTime = 0
		var schema *reflection.Schema
		err := backoff.Retry(func() error {
			var err error
			schema, err = s.conn.Reflect(ctx)
			if err != nil {
				klog.Warningf("Failed to fetch new schema: %v", err)
				return err
			}
			return nil
		}, backoff.WithContext(bo, ctx))
		// This will only happen due to context expiration.
		if err != nil {
			klog.Errorf("Giving up on schema fetch: %v", err)
			continue
		}

		// Swap the current schema if necessary.
		cur := s.curSchema().Version
		new := schema.Version
		if cur != new {
			klog.Infof("Got new schema: %s -> %s", cur, new)
			s.muSchema.Lock()
			s.schema = schema
			s.muSchema.Unlock()
		}
	}
}

// Register webug on an HTTP mux, using a BMDB connection pool.
//
// The given context will be used not only to time out the registration call, but
// also used to run a BMDB schema fetching goroutine that will attempt to fetch
// newer versions of the schema every hour.
//
// This is a low-level function useful when tying webug into an existing web
// application. If you just want to run webug on a separate port that's
// configured by flags, use Config and Config.RegisterFlags.
func Register(ctx context.Context, conn *bmdb.Connection, mux *http.ServeMux) error {
	schema, err := conn.Reflect(ctx)
	if err != nil {
		return fmt.Errorf("could not get BMDB schema for webug: %w", err)
	}
	s := server{
		conn:   conn,
		schema: schema,
	}
	go s.schemaWorker(ctx)

	type route struct {
		pattern *regexp.Regexp
		handler func(w http.ResponseWriter, r *http.Request, args ...string)
	}

	routes := []route{
		{regexp.MustCompile(`^/$`), s.viewMachines},
		{regexp.MustCompile(`^/machine/([a-fA-F0-9\-]+)$`), s.viewMachineDetail},
		{regexp.MustCompile(`^/provider/([^/]+)/([^/]+)$`), s.viewProviderRedirect},
		{regexp.MustCompile(`^/session/([^/]+)`), s.viewSession},
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, route := range routes {
			match := route.pattern.FindStringSubmatch(r.URL.Path)
			if match == nil {
				continue
			}
			route.handler(w, r, match[1:]...)
			return
		}
		http.NotFound(w, r)
	})
	return nil
}

// Config describes the webug interface configuration. This should be embedded
// inside your component's Config object.
//
// To configure, either set values or call RegisterFlags before flag.Parse.
//
// To run after configuration, call Start.
type Config struct {
	// If set, start a webug interface on an HTTP listener bound to the given address.
	WebugListenAddress string
}

// RegisterFlags for webug interface.
func (c *Config) RegisterFlags() {
	flag.StringVar(&c.WebugListenAddress, "webug_listen_address", "", "Address to start BMDB webug on. If not set, webug will not be started.")
}

// Start the webug interface in the foreground if enabled. The returned error
// will be either a configuration/connection error returned as soon as possible,
// or a context expiration error.
//
// The given context will be used for all connections from the webug interface to
// the given BMDB connection.
func (c *Config) Start(ctx context.Context, conn *bmdb.Connection) error {
	if c.WebugListenAddress == "" {
		return nil
	}
	mux := http.NewServeMux()
	if err := Register(ctx, conn, mux); err != nil {
		return err
	}

	klog.Infof("Webug listening at %s", c.WebugListenAddress)
	return http.ListenAndServe(c.WebugListenAddress, mux)
}
