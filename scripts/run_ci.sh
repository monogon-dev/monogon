#!/bin/bash
# This helper scripts executes all Bazel tests in our CI environment.
# https://phab.monogon.dev/harbormaster/plan/2/
set -euo pipefail

DOCKERFILE_HASH=$(sha1sum build/Dockerfile | cut -c -8)

BUILD_ID=$1;
BUILD_PHID=$2;
shift; shift;

TAG=nexantic-version-${DOCKERFILE_HASH}
POD=nexantic-build-${BUILD_ID}

# We keep one Bazel build cache per working copy to avoid concurrency issues
# (we cannot run multiple Bazel servers on a given _bazel_root)
function getWorkingCopyID {
  local pattern='/var/drydock/workingcopy-([0-9]+)/'
  [[ "$(pwd)" =~ $pattern ]]
  echo ${BASH_REMATCH[1]}
}

CACHE_VOLUME=bazel-cache-$(getWorkingCopyID)

# We do our own image caching since the podman build step cache does
# not work across different repository checkouts and is also easily
# invalidated by multiple in-flight revisions with different Dockerfiles.
if ! podman image inspect "$TAG" >/dev/null; then
  echo "Could not find $TAG, building..."
  podman build -t ${TAG} build
fi

# Keep this in sync with create_container.sh:

function cleanup {
  rc=$?
  ! podman pod rm $POD --force
  exit $rc
}

trap cleanup EXIT

! podman volume create --opt o=nodev,exec ${CACHE_VOLUME}

podman pod create --name ${POD}

podman run \
    --rm \
    -v $(pwd):/work \
    -v ${CACHE_VOLUME}:/user/.cache/bazel/_bazel_root \
    --privileged \
    ${TAG} \
    scripts/gazelle.sh

if [[ ! -z "$(git status --porcelain)" ]]; then
  echo "Unclean working directory after running scripts/gazelle.sh:"
  git diff HEAD
  exit 1
fi

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
    bazel test //...

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
