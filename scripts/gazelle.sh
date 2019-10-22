#!/bin/bash
# gazelle.sh regenerates BUILD.bazel files for Go source files.

bazelisk run //:gazelle -- update
bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=repositories.bzl%go_repositories
