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
podman build -t smalltown-builder .

# Set up SELinux contexts to prevent the container from writing to
# files that would allow for easy breakouts via tools ran on the host.
chcon -R system_u:object_r:container_file_t:s0 .
chcon -R unconfined_u:object_r:user_home_t:s0 \
  .arcconfig .idea .git

# Create cache volume if it does not yet exist
! podman volume create repo-cache

podman run -it -d \
    -v $(pwd):/work \
    -v repo-cache:/root/repo-cache \
    --tmpfs=/root/.cache/bazel:exec \
    --device /dev/kvm \
    --net=host \
    --name=smalltown-dev \
    smalltown-builder
