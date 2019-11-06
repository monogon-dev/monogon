#!/bin/bash
# Copy generated Go protobuf libraries to a place where a non-Bazel-aware IDE can find them.
# Locally, a symlink will be sufficient.

mkdir -p core/generated
rsync -av --delete --exclude '*.a' bazel-bin/core/api/*/linux_amd64_stripped/*/git.monogon.dev/source/nexantic.git/core/generated/* core/generated/
