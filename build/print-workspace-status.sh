#!/usr/bin/env bash

# Workspace status used for build stamping.
# This is currently only consumed by //third_party/go/kubernetes_version_def.bzl.

set -o errexit
set -o nounset
set -o pipefail

KUBERNETES_gitTreeState="clean"
if [ ! -z "$(git status --porcelain)" ]; then
    KUBERNETES_gitTreeState="dirty"
fi

# TODO(q3k): unify with //third_party/go/repsitories.bzl.
KUBERNETES_gitMajor="1"
KUBERNETES_gitMinor="19"
KUBERNETES_gitVersion="v1.19.0-alpha.2+nxt"

cat <<EOF
KUBERNETES_gitCommit $(git rev-parse "HEAD^{commit}")
KUBERNETES_gitTreeState $KUBERNETES_gitTreeState
STABLE_KUBERNETES_gitVersion $KUBERNETES_gitVersion
STABLE_KUBERNETES_gitMajor $KUBERNETES_gitMajor
STABLE_KUBERNETES_gitMinor $KUBERNETES_gitMinor
KUBERNETES_buildDate $(date \
  ${SOURCE_DATE_EPOCH:+"--date=@${SOURCE_DATE_EPOCH}"} \
 -u +'%Y-%m-%dT%H:%M:%SZ')
EOF
