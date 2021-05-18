#!/usr/bin/env bash

# This script can be run by Monogon Labs employees to push the Builder image
# built from //build/ci/Dockerfile into a public registry image. That image is
# then consumed by external, non-public infrastructure code as a basis to run
# Jenkins CI agents.
#
# For more information, see //build/ci/README.md.

set -euo pipefail

main() {
    if [[ "$HOME" == /user ]] && [[ -d /user ]] && [[ -d /home/ci ]]; then
        echo "WARNING: likely running within Builder image instead of the host environment." >&2
        echo "If this script was invoked using 'bazel run', please instead do:" >&2
        echo "    \$ scripts/bin/bazel build //build/ci:push_ci_image" >&2
        echo "    \$ bazel-bin/build/ci/push_ci_image" >&2
        echo "This will build the script within the container but run it on the host." >&2
    fi

    local podman="$(command -v podman || true)"
    if [[ -z "$podman" ]]; then
        echo "'podman' must be available in the system PATH to build the image." >&2
        exit 1
    fi

    local dockerfile="build/ci/Dockerfile"
    if [[ ! -f "${dockerfile}" ]]; then
        echo "Could not find ${dockerfile} - this script needs to be run from the root of the Monogon repository." >&2
        ecit 1
    fi

    local image="gcr.io/monogon-infra/monogon-builder:$(date +%s)"

    echo "Building image ${image} from ${dockerfile}..."
    "${podman}" build -t "${image}" - < "${dockerfile}"
    echo "Pushing image ${image}..."
    "${podman}" push "${image}"
    echo "Done, new image is ${image}"
}

main "$@"

