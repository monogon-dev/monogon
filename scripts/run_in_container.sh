#!/usr/bin/env bash
set -euo pipefail

if [[ "$(podman inspect monogon-dev --format "{{ (index .Mounts 0).Source }}")" != "$(pwd)" ]]; then
    echo "Please run this wrapper from the original checkout"
    exit 1
fi

podman exec -it monogon-dev $@
