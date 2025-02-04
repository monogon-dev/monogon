// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package manager, itself a part of BMaaS project, provides implementation
// governing Equinix bare metal server lifecycle according to conditions set by
// Bare Metal Database (BMDB).
//
// The implementation will attempt to provide as many machines as possible and
// register them with BMDB. This is limited by the count of Hardware
// Reservations available in the Equinix Metal project used. The BMaaS agent
// will then be started on these machines as soon as they become ready.
//
// The implementation is provided in the form of a library, to which interface is
// exported through Provisioner and Initializer types, each taking servers
// through a single stage of their lifecycle.
//
// See the included test code for usage examples.
//
// The terms "device" and "machine" are used interchangeably throughout this
// package due to differences in Equinix Metal and BMDB nomenclature.
package manager
