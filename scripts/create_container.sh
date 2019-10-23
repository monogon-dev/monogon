#!/bin/bash
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
podman build -t nexantic-builder build

# Set up SELinux contexts to prevent the container from writing to
# files that would allow for easy breakouts via tools ran on the host.
chcon -R system_u:object_r:container_file_t:s0 .
chcon -R unconfined_u:object_r:user_home_t:s0 \
  .arcconfig .idea .git

# TODO(leo): mount .cache/bazel on a volume (waiting for podman issue to be fixed)
# https://github.com/containers/libpod/issues/4318
podman run -it -d \
    -v $(pwd):/work \
    -v smalltown-gopath:/user/go/pkg \
    -v smalltown-gobuildcache:/user/.cache/go-build \
    -v smalltown-bazelcache:/user/.cache/bazel/_bazel_root/cache \
    --tmpfs=/user/.cache/bazel:exec \
    --device /dev/kvm \
    --privileged \
    --userns=keep-id \
    --name=nexantic-dev \
    nexantic-builder
