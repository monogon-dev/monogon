#!/usr/bin/env bash
set -euo pipefail

mkdir -p .vendor/linux
curl -L https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.19.72.tar.xz | tar -xJf - -C .vendor/linux --strip-components 1
ln -rfs kernel/linux-smalltown.config .vendor/linux/.config
