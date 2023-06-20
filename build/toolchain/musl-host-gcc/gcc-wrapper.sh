#!/usr/bin/env bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
exec /usr/bin/gcc "$@" -specs $SCRIPT_DIR/musl.spec
