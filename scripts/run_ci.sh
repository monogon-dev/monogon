#!/bin/bash
# This script is executed by our CI
set -euo pipefail

BUILD_ID=$1;
BUILD_PHID=$2;
shift; shift;

TAG=nexantic-build-${BUILD_ID}
POD=nexantic-build-${BUILD_ID}

# We keep one Bazel build cache per working copy to avoid concurrency issues
# (we cannot run multiple Bazel servers on a given _bazel_root)
function getWorkingCopyID {
  local pattern='/var/drydock/workingcopy-([0-9]+)/'
  [[ "$(pwd)" =~ $pattern ]]
  echo ${BASH_REMATCH[1]}
}

CACHE_VOLUME=bazel-cache-$(getWorkingCopyID)

# New image for each build - the Dockerfile might have changed.
# Rely on the build step cache to avoid costly rebuilds.
podman build -t ${TAG} build

# Keep this in sync with create_container.sh:

function cleanup {
  rc=$?
  ! podman pod rm $POD --force
  ! podman rmi $TAG --force
  exit $rc
}

trap cleanup EXIT

! podman volume create --opt o=nodev,exec ${CACHE_VOLUME}

podman pod create --name ${POD}

podman run -d \
    --pod ${POD} \
    --ulimit nofile=262144:262144 \
    --name=${POD}-cockroach \
    cockroachdb/cockroach:v19.1.5 start --insecure

podman run \
    -v $(pwd):/work \
    -v ${CACHE_VOLUME}:/user/.cache/bazel/_bazel_root \
    --device /dev/kvm \
    --privileged \
    --pod ${POD} \
    --name=${POD}-bazel \
    ${TAG} \
    $@

function conduit() {
  # Get Phabricator host from Git origin
  local pattern='ssh://(.+?):([0-9]+)'
  [[ "$(git remote get-url origin)" =~ $pattern ]];
  local host=${BASH_REMATCH[1]}
  local port=${BASH_REMATCH[2]}

  ssh "$host" -p "$port" conduit $@
}

# Report build results if we made it here successfully
conduit harbormaster.sendmessage <<EOF
{"params": "{\"buildTargetPHID\": \"${BUILD_PHID}\", \"type\": \"pass\"}"}
EOF
