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

/*
Package logtree implements a tree-shaped logger for debug events. It provides log publishers (ie. Go code) with a
glog-like API, with loggers placed in a hierarchical structure defined by a dot-delimited path (called a DN, short for
Distinguished Name).

    tree.MustLeveledFor("foo.bar.baz").Warningf("Houston, we have a problem: %v", err)

Logs in this context are unstructured, operational and developer-centric human readable text messages presented as lines
of text to consumers, with some attached metadata. Logtree does not deal with 'structured' logs as some parts of the
industry do, and instead defers any machine-readable logs to either be handled by metrics systems like Prometheus or
event sourcing systems like Kafka.

Tree Structure

As an example, consider an application that produces logs with the following DNs:

    listener.http
    listener.grpc
    svc
    svc.cache
    svc.cache.gc

This would correspond to a tree as follows:

                          .------.
                         |   ""   |
                         | (root) |
                          '------'
           .----------------'   '------.
    .--------------.           .---------------.
    |     svc      |           |    listener   |
    '--------------'           '---------------'
           |                   .----'      '----.
    .--------------.  .---------------.  .---------------.
    |  svc.cache   |  | listener.http |  | listener.grpc |
    '--------------'  '---------------'  '---------------'
           |
    .--------------.
    | svc.cache.gc |
    '--------------'

In this setup, every DN acts as a separate logging target, each with its own retention policy and quota. Logging to a DN
under foo.bar does NOT automatically log to foo - all tree mechanisms are applied on log access by consumers. Loggers
are automatically created on first use, and importantly, can be created at any time, and will automatically be created
if a sub-DN is created that requires a parent DN to exist first. Note, for instance, that a `listener` logging node was
created even though the example application only logged to `listener.http` and `listener.grpc`.

An implicit root node is always present in the tree, accessed by DN "" (an empty string). All other logger nodes are
children (or transitive children) of the root node.

Log consumers (application code that reads the log and passes them on to operators, or ships them off for aggregation in
other systems) to select subtrees of logs for readout. In the example tree, a consumer could select to either read all
logs of the entire tree, just a single DN (like svc), or a subtree (like everything under listener, ie. messages emitted
to listener.http and listener.grpc).

Log Producer API

As part of the glog-like logging API available to producers, the following metadata is attached to emitted logs in
addition to the DN of the logger to which the log entry was emitted:

 - timestamp at which the entry was emitted
 - a severity level (one of FATAL, ERROR, WARN or INFO)
 - a source of the message (file name and line number)

In addition, the logger mechanism supports a variable verbosity level (so-called 'V-logging') that can be set at every
node of the tree. For more information about the producer-facing logging API, see the documentation of the LeveledLogger
interface, which is the main interface exposed to log producers.

Log Access API

The Log Access API is mostly exposed via a single function on the LogTree struct: Read. It allows access to log entries
that have been already buffered inside LogTree and to subscribe to receive future entries over a channel. As outlined
earlier, any access can specify whether it is just interested in a single logger (addressed by DN), or a subtree of
loggers.

Due to the current implementation of the logtree, subtree accesses of backlogged data is significantly slower than
accessing data of just one DN, or the whole tree (as every subtree backlog access performs a scan on all logged data).
Thus, log consumers should be aware that it is much better to stream and buffer logs specific to some long-standing
logging request on their own, rather than repeatedly perform reads of a subtree backlog.

*/
package logtree
