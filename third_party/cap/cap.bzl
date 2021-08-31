#  Copyright 2021 The Monogon Project Authors.
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
load("@io_bazel_rules_go//go:def.bzl", "go_binary")

cc_library(
    name = "cap",
    srcs = [
        "cap_alloc.c",
        "cap_extint.c",
        "cap_file.c",
        "cap_flag.c",
        "cap_proc.c",
        "cap_text.c",
    ] + glob(["include/sys/*.h"]),  # UAPI is intentionally excluded as we want it from our own kernel headers
    hdrs = ["libcap.h", ":cap_names.h"],
    visibility = ["//visibility:public"],
    includes = [".", "include"],
)

go_binary(
    name = "makenames",
    srcs = ["makenames.go"],
)

genrule(
    name = "cap_names_hdr",
    srcs = [
        "include/uapi/linux/capability.h",
    ],
    outs = [
        "cap_names.h",
    ],
    cmd = "$(location :makenames) \"$(location include/uapi/linux/capability.h)\" \"$@\"",
    tools = [
        ":makenames",
    ],
)
