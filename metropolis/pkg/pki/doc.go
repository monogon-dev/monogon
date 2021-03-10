// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package pki implements an x509 PKI (Public Key Infrastructure) system backed
// on etcd.
//
// The PKI is made of Certificates, all constrained within a Namespace. The
// Namespace allows for multiple users of this library to co-exist on a single
// etcd server.
//
// Any time a Certificate object is created, it describes the promise (or
// intent) of an x509 certificate to exist. For every created Certifiacte an
// Issuer must be specified - either another Certificate (which will act as a
// CA and sign that Certificate), or SelfSigned (which will cause the
// Certificate to be self-signed when generated).
//
// Once a Certificate object is created, a call to Ensure() must be placed to
// turn the intent of a certificate into physical bytes that can then be
// accessed by the appliaction.
//
// Two kinds of Certificates can be created:
//  - Named certificates are stored in etcd, and an Ensure call will either
//    create them, or return a Certificate already stored in etcd. Multiple
//    concurrent calls to Ensure for a Certificate with the same name are
//    permitted, even across machines, as long as the Certificate intent data
//    is the same. If not, it is still safe to perform this action
//    concurrently, but the first transaction will win, causing the losing
//    transaction to return the Ensure call with a certificate that was not
//    based on the same intent.
//    It is the responsibility of the caller to ensure these cases are handled
//    gracefully.
//  - Volatile certificates are stored in memory, and have an empty ("") name.
//    Any time Ensure is called, the certificate already present in memory is
//    returned, or one is created if it does not yet exist.
//    Currently, these certificates live fully in memory, but in the future we
//    will likely perform audit logging (and revocation) of these certificate
//    within etcd, too.
//
package pki
