package component

import (
	"database/sql"
	"flag"
	"net/url"
	"os"
	"sync"

	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/lib/pq"
	"k8s.io/klog/v2"

	"source.monogon.dev/metropolis/cli/pkg/datafile"
)

// CockroachConfig is the common configuration of a components' connection to
// CockroachDB. It's supposed to be instantiated within a Configuration struct
// of a component.
//
// It can be configured by flags (via RegisterFlags) or manually (eg. in tests).
type CockroachConfig struct {
	// Migrations is the go-migrate source of migrations for this database. Usually
	// this can be taken from a go-embedded set of migration files.
	Migrations source.Driver

	// EndpointHost is the host part of the endpoint address of the database server.
	EndpointHost string
	// TLSKeyPath is the filesystem path of the x509 key used to authenticate to the
	// database server.
	TLSKeyPath string
	// TLSKeyPath is the filesystem path of the x509 certificate used to
	// authenticate to the database server.
	TLSCertificatePath string
	// TLSCACertificatePath is the filesystem path of the x509 CA certificate used
	// to verify the database server's certificate.
	TLSCACertificatePath string
	// UserName is the username to be used on the database server.
	UserName string
	// UserName is the database name to be used on the database server.
	DatabaseName string

	// InMemory indicates that an in-memory CockroachDB instance should be used.
	// Data will be lost after the component shuts down.
	InMemory bool

	// mu guards inMemoryInstance.
	mu sync.Mutex
	// inMemoryInstance is populated with a CockroachDB test server handle when
	// InMemory is set and Connect()/MigrateUp() is called.
	inMemoryInstance testserver.TestServer
}

// RegisterFlags registers the connection configuration to be provided by flags.
// This must be called exactly once before then calling flags.Parse().
func (c *CockroachConfig) RegisterFlags(prefix string) {
	flag.StringVar(&c.EndpointHost, prefix+"_endpoint_host", "", "Host of CockroachDB endpoint for "+prefix)
	flag.StringVar(&c.TLSKeyPath, prefix+"_tls_key_path", "", "Path to CockroachDB TLS client key for "+prefix)
	flag.StringVar(&c.TLSCertificatePath, prefix+"_tls_certificate_path", "", "Path to CockroachDB TLS client certificate for "+prefix)
	flag.StringVar(&c.TLSCACertificatePath, prefix+"_tls_ca_certificate_path", "", "Path to CockroachDB CA certificate for "+prefix)
	flag.StringVar(&c.UserName, prefix+"_user_name", prefix, "CockroachDB user name for "+prefix)
	flag.StringVar(&c.DatabaseName, prefix+"_database_name", prefix, "CockroachDB database name for "+prefix)
	flag.BoolVar(&c.InMemory, prefix+"_eat_my_data", false, "Use in-memory CockroachDB for "+prefix+". Warning: Data will be lost at process shutdown!")
}

// startInMemory starts an in-memory cockroachdb server as a subprocess, and
// returns a DSN that connects to the newly created database.
func (c *CockroachConfig) startInMemory(scheme string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	klog.Warningf("STARTING IN-MEMORY COCKROACHDB FOR TESTS")
	klog.Warningf("ALL DATA WILL BE LOST AFTER SERVER SHUTDOWN!")

	if c.inMemoryInstance == nil {
		opts := []testserver.TestServerOpt{
			testserver.SecureOpt(),
		}
		if path, err := datafile.ResolveRunfile("external/cockroach/cockroach"); err == nil {
			opts = append(opts, testserver.CockroachBinaryPathOpt(path))
		} else {
			if os.Getenv("TEST_TMPDIR") != "" {
				klog.Exitf("In test which requires in-memory cockroachdb, but @cockroach//:cockroach missing as a dependency. Failing.")
			}
			klog.Warningf("CockroachDB in-memory database requested, but not available as a build dependency. Trying to download it...")
		}

		inst, err := testserver.NewTestServer(opts...)
		if err != nil {
			klog.Exitf("Failed to create crdb test server: %v", err)
		}
		c.inMemoryInstance = inst
	}

	u := *c.inMemoryInstance.PGURL()
	u.Scheme = scheme
	return u.String()
}

// buildDSN returns a DSN to the configured database connection with a given DSN
// scheme. The scheme will usually be 'postgres' or 'cockroach', depending on
// whether it's used for lib/pq or for golang-migrate.
func (c *CockroachConfig) buildDSN(scheme string) string {
	if c.InMemory {
		return c.startInMemory(scheme)
	}

	query := make(url.Values)
	query.Set("sslmode", "verify-full")
	query.Set("sslcert", c.TLSCertificatePath)
	query.Set("sslkey", c.TLSKeyPath)
	query.Set("sslrootcert", c.TLSCACertificatePath)
	u := url.URL{
		Scheme:   scheme,
		User:     url.User(c.UserName),
		Host:     c.EndpointHost,
		Path:     c.DatabaseName,
		RawQuery: query.Encode(),
	}
	return u.String()
}

// Connect returns a working *sql.DB handle to the database described by this
// CockroachConfig.
func (d *CockroachConfig) Connect() (*sql.DB, error) {
	dsn := d.buildDSN("postgres")
	klog.Infof("Connecting to %s...", dsn)
	return sql.Open("postgres", d.buildDSN("postgres"))
}

// MigrateUp performs all possible migrations upwards for the database described
// by this CockroachConfig.
func (d *CockroachConfig) MigrateUp() error {
	dsn := d.buildDSN("cockroachdb")
	klog.Infof("Running migrations on %s...", dsn)
	m, err := migrate.NewWithSourceInstance("iofs", d.Migrations, dsn)
	if err != nil {
		return err
	}
	return m.Up()
}
