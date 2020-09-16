#!/usr/bin/env bash
exec /usr/bin/gcc "$@" -specs build/toolchain/musl-host-gcc/musl.spec
