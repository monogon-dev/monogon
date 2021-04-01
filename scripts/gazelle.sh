#!/usr/bin/env bash
# gazelle.sh regenerates BUILD.bazel files for Go source files.
set -euo pipefail

bazel run //:gazelle -- update
