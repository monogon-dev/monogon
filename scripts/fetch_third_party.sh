#!/usr/bin/env bash
set -euo pipefail

mkdir -p third_party/linux
curl -L https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.19.72.tar.xz | tar -xJf - -C third_party/linux --strip-components 1
ln -fs ../../kernel/linux-smalltown.config third_party/linux/.config