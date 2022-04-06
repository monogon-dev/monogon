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

NOTE: currently the first part (`go mod tidy`) doesn't work without performing some in-place symlinking in the repository. TODO(lorenz): document this

Regenerating BUILDfiles
-----------------------

    $ bazel run //:gazelle
