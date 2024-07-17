#  Copyright 2020 The Monogon Project Authors.
#
#  SPDX-License-Identifier: Apache-2.0
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

load("@rules_cc//cc:defs.bzl", "cc_library")
load("@@//build/utils:template_file.bzl", "template_file")

template_file(
    name = "pixman-version.h",
    src = "pixman-version.h.in",
    substitutions = {
        "@PIXMAN_VERSION_MAJOR@": "0",
        "@PIXMAN_VERSION_MINOR@": "40",
        "@PIXMAN_VERSION_MICRO@": "1",
    },
)

genrule(
    name = "config-h",
    outs = ["config.h"],
    cmd = "echo \"\" > \"$@\"",
)

cc_library(
    name = "pixman",
    srcs = [
        ":config-h",
        "pixman.c",
        "pixman-access.c",
        "pixman-access-accessors.c",
        "pixman-accessor.h",
        "pixman-arm.c",
        "pixman-bits-image.c",
        "pixman-combine-float.c",
        "pixman-combine32.c",
        "pixman-combine32.h",
        "pixman-compiler.h",
        "pixman-conical-gradient.c",
        "pixman-edge.c",
        "pixman-edge-accessors.c",
        "pixman-edge-imp.h",
        "pixman-fast-path.c",
        "pixman-filter.c",
        "pixman-general.c",
        "pixman-glyph.c",
        "pixman-gradient-walker.c",
        "pixman-image.c",
        "pixman-implementation.c",
        "pixman-inlines.h",
        "pixman-linear-gradient.c",
        "pixman-matrix.c",
        "pixman-mips.c",
        "pixman-noop.c",
        "pixman-ppc.c",
        "pixman-private.h",
        "pixman-radial-gradient.c",
        "pixman-region16.c",
        "pixman-region32.c",
        "pixman-solid-fill.c",
        "pixman-sse2.c",
        "pixman-ssse3.c",
        "pixman-timer.c",
        "pixman-trap.c",
        "pixman-utils.c",
        "pixman-x86.c",
        ":pixman-version.h",
        "dither/blue-noise-64x64.h",
    ],
    hdrs = [
        "pixman.h",
        # Please never include these, this is some next-level insanity
        "pixman-region.c",
        "pixman-edge.c",
        "pixman-access.c",
    ],
    copts = ["-mssse3"],
    includes = ["."],
    local_defines = [
        "PACKAGE=foo",
        "HAVE_FLOAT128=1",
        "HAVE_BUILTIN_CLZ=1",
        "USE_SSSE3=1",
        "USE_SSE2=1",
        "TLS=__thread",
    ],
    visibility = ["//visibility:public"],
)
