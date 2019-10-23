#!/bin/bash
# Copy generated Go protobuf libraries to a place where a non-Bazel-aware IDE can find them.
# Locally, a symlink will be sufficient.

mkdir -p smalltown/generated
rsync -av --delete --exclude '*.a' bazel-bin/smalltown/api/*/linux_amd64_stripped/*/git.monogon.dev/source/nexantic.git/smalltown/generated/* smalltown/generated/
