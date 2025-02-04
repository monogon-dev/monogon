#!/bin/bash
#
# Copyright The Monogon Project Authors.
# SPDX-License-Identifier: Apache-2.0
#

exec "$1" -initrd-path "$2" -kernel-path "$3" -cmdline "$4"