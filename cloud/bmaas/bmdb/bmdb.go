// Package bmdb implements a connector to the Bare Metal Database, which is the
// main data store backing information about bare metal machines.
//
// All components of the BMaaS project connect directly to the underlying
// CockroachDB database storing this data via this library. In the future, this
// library might turn into a shim which instead connects to a coordinator
// service over gRPC.
package bmdb
