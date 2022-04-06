Go dependency managment
=======================

Status: managed by [Gazelle](https://github.com/bazelbuild/bazel-gazelle).


    .--------.
    | go.mod |------------.
    '--------'            |
        | go mod tidy     | bazel //:gazelle-update-repos
        V                 |
    .--------.            |
    | go.sum |-----------.|
    '--------'            |
                          V
    .---------------------------------.
    | third_party/go/repositories.bzl |
    '---------------------------------'
                          | bazel run //:gazelle
                          V
                   .----------------.
                   | **/BUILD.bazel |.
                   '----------------'|
                    '----------------'
                          | bazel build //...
                          V
                   .-----------------.
                   | build artifacts |
                   '-----------------'

Updating and adding new dependencies
------------------------------------

Add a Go dependency to your code, then:

    $ go mod tidy
    $ bazel run //:gazelle-update-repos

All generated sources (eg. protobuf stubs) that are usually built by Bazel are invisible to go(mod)-based tooling. To get around this, we place `gomod-generated-placeholder.go` files in package directories that would otherwise contain generated files. These are ignored by Gazelle (and thus by Bazel builds) but not by go(mod)-based tooling.

Regenerating BUILDfiles
-----------------------

    $ bazel run //:gazelle
