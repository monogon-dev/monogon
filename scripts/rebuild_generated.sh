#!/bin/bash
bazel build core/api/... && scripts/bazel_copy_generated_for_ide.sh
