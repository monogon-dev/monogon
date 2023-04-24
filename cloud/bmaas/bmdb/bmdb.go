// Package bmdb implements a connector to the Bare Metal Database, which is the
// main data store backing information about bare metal machines.
//
// All components of the BMaaS project connect directly to the underlying
// CockroachDB database storing this data via this library. In the future, this
// library might turn into a shim which instead connects to a coordinator
// service over gRPC.
package bmdb

import (
	"github.com/prometheus/client_golang/prometheus"

	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
	"source.monogon.dev/cloud/lib/component"
)

// BMDB is the Bare Metal Database, a common schema to store information about
// bare metal machines in CockroachDB. This struct is supposed to be
// embedded/contained by different components that interact with the BMDB, and
// provides a common interface to BMDB operations to these components.
//
// The BMDB provides two mechanisms facilitating a 'reactive work system' being
// implemented on the bare metal machine data:
//
//   - Sessions, which are maintained by heartbeats by components and signal the
//     liveness of said components to other components operating on the BMDB. These
//     effectively extend CockroachDB's transactions to be visible as row data. Any
//     session that is not actively being updated by a component can be expired by a
//     component responsible for lease garbage collection.
//   - Work locking, which bases on Sessions and allows long-standing
//     multi-transaction work to be performed on given machines, preventing
//     conflicting work from being performed by other components. As both Work
//     locking and Sessions are plain row data, other components can use SQL queries
//     to exclude machines to act on by constraining SELECT queries to not return
//     machines with some active work being performed on them.
type BMDB struct {
	Config

	metrics *metrics.MetricsSet
}

// Config is the configuration of the BMDB connector.
type Config struct {
	Database component.CockroachConfig

	// ComponentName is a human-readable name of the component connecting to the
	// BMDB, and is stored in any Sessions managed by this component's connector.
	ComponentName string
	// RuntimeInfo is a human-readable 'runtime information' (eg. software version,
	// host machine/job information, IP address, etc.) stored alongside the
	// ComponentName in active Sessions.
	RuntimeInfo string
}

// EnableMetrics configures BMDB metrics collection and registers it on the given
// registry. This method should only be called once, and is not goroutine safe.
func (b *BMDB) EnableMetrics(registry *prometheus.Registry) {
	if b.metrics == nil {
		b.metrics = metrics.New(registry)
	}
}
