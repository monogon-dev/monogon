#!/bin/bash
set -euo pipefail

podman exec -it monogon-dev $@
