#!/usr/bin/env bash
set -euo pipefail

# Our local user needs write access to /dev/kvm (best accomplished by
# adding your user to the kvm group).
if ! touch /dev/kvm; then
  echo "Cannot write to /dev/kvm - please verify permissions."
  exit 1
fi

# The KVM module needs to be loaded, since our container is unprivileged
# and won't be able to do it itself.
if ! [[ -d /sys/module/kvm ]]; then
  echo "kvm module not loaded - please modprobe kvm"
  exit 1
fi

# Rebuild base image
podman build -t monogon-builder build

# Keep this in sync with ci.sh:

podman pod create --name monogon

# Mount bazel root to identical paths inside and outside the container.
# This caches build state even if the container is destroyed, and
BAZEL_ROOT=${HOME}/.cache/bazel-nxt
mkdir -p ${BAZEL_ROOT}

# The Bazel plugin injects a Bazel repository into the sync command line,
# We need to copy the aspect repository and apply a custom patch.

# TODO(leo): the IntelliJ path changed to ~/.config on new setups, we should look for that as well

IJ_HOME=$(echo ${HOME}/.IntelliJIdea* | tr ' ' '\n' | sort | tail -n 1)
ASPECT_ORIG=${IJ_HOME}/config/plugins/ijwb/aspect
ASPECT_PATH=${BAZEL_ROOT}/ijwb_aspect

if [[ -d "$IJ_HOME" ]]; then
    echo "IntelliJ found, copying aspect file to Bazel root"
    rm -rf "$ASPECT_PATH"
    cp -r "$ASPECT_ORIG" "$ASPECT_PATH"
    patch -d "$ASPECT_PATH" -p1 < scripts/patches/bazel_intellij_aspect_filter.patch
fi

podman run -it -d \
    -v $(pwd):$(pwd):z \
    -w $(pwd) \
    --volume=${BAZEL_ROOT}:${BAZEL_ROOT} \
    --device /dev/kvm \
    --privileged \
    --pod monogon \
    --name=monogon-dev \
    --net=host \
    monogon-builder
