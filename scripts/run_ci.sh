#!/usr/bin/env bash
# This helper scripts executes all Bazel tests in our CI environment.
# https://phab.monogon.dev/harbormaster/plan/2/
set -euo pipefail

DOCKERFILE_HASH=$(sha1sum build/Dockerfile | cut -c -8)

BUILD_ID=$1;
BUILD_PHID=$2;
shift; shift;

TAG=monogon-version-${DOCKERFILE_HASH}
CONTAINER=monogon-build-${BUILD_ID}

# We keep one set of Bazel build caches per working copy to avoid concurrency
# issues (we cannot run multiple Bazel servers on a given _bazel_root).
function getWorkingCopyID {
  local pattern='/var/drydock/workingcopy-([0-9]+)/'
  [[ "$(pwd)" =~ $pattern ]]
  echo ${BASH_REMATCH[1]}
}

# Main Bazel cache, used as Bazel outputRoot/outputBase.
CACHE_VOLUME=bazel-cache-$(getWorkingCopyID)
# Secondary disk cache for Bazel, used to keep build data between configuration
# switches (saving from spurious rebuilds when switchint from debug to
# non-debug builds).
SECONDARY_CACHE_VOLUME=bazel-secondary-cache-$(getWorkingCopyID)
SECONDARY_CACHE_LOCATION="/user/.cache/bazel-secondary"
# TODO(q3k): Neither the main nor secondary caches are garbage collected and
# they will slowly fill up the disk of the CI builder.

# The Go pkg cache is safe to use concurrently.
GOPKG_VOLUME=gopkg-cache

cat > ci.bazelrc <<EOF
build --disk_cache=${SECONDARY_CACHE_LOCATION}
EOF

# We do our own image caching since the podman build step cache does
# not work across different repository checkouts and is also easily
# invalidated by multiple in-flight revisions with different Dockerfiles.
if ! podman image inspect "$TAG" >/dev/null; then
  echo "Could not find $TAG, building..."
  podman build -t ${TAG} build
fi

function cleanup {
  rc=$?
  ! podman kill $CONTAINER
  ! podman rm $CONTAINER --force
  exit $rc
}

trap cleanup EXIT

! podman kill $CONTAINER
! podman rm $CONTAINER --force

! podman volume create --opt o=nodev,exec ${CACHE_VOLUME}
! podman volume create --opt o=nodev ${SECONDARY_CACHE_VOLUME}
! podman volume create --opt o=nodev ${GOPKG_VOLUME}

function bazel() {
    podman run \
        --rm \
        --name $CONTAINER \
        -v $(pwd):/work \
        -v ${CACHE_VOLUME}:/user/.cache/bazel/_bazel_root \
        -v ${SECONDARY_CACHE_VOLUME}:${SECONDARY_CACHE_LOCATION} \
        -v ${GOPKG_VOLUME}:/user/go/pkg \
        --privileged \
        ${TAG} \
        bazel "$@"
}

bazel run //:fietsje
bazel run //:gazelle -- update

if [[ ! -z "$(git status --porcelain)" ]]; then
  echo "Unclean working directory after running gazelle and fietsje:"
  git diff HEAD
  cat <<EOF
Please run:

  $ bazel run //:fietsje
  $ bazel run //:gazelle -- update

in your local branch and add the resulting changes to this diff.
EOF
  exit 1
fi

bazel test //...
bazel test //... -c dbg

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
