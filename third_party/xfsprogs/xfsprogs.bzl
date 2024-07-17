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

load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")
load("@@//build/utils:template_file.bzl", "template_file")

defs = [
    "NDEBUG",  # Doesn't compile without it because their assertions reference non-existent fields
    "_GNU_SOURCE",
    "ENABLE_BLKID",
    "HAVE_FSETXATTR",
    "HAVE_GETFSMAP",
    "HAVE_GETMNTENT",
    "HAVE_MNTENT",
    "VERSION=\\\"0.0.0\\\"",
]

template_file(
    name = "platform_defs.h",
    src = "include/platform_defs.h.in",
    substitutions = {
        "#undef SIZEOF_LONG": "#define SIZEOF_LONG sizeof(long)",  # Because C reasons
    },
)

cc_library(
    name = "util",
    srcs = [
        "libfrog/util.c",
        ":platform_defs.h",
    ],
    hdrs = ["libfrog/util.h"],
    local_defines = defs,
)

cc_library(
    name = "radix_tree",
    srcs = [
        "libfrog/radix-tree.c",
        ":platform_defs.h",
    ],
    hdrs = ["libfrog/radix-tree.h"],
    local_defines = defs,
)

cc_binary(
    name = "gen_crc32table",
    srcs = [
        "libfrog/crc32defs.h",
        "libfrog/gen_crc32table.c",
    ],
)

genrule(
    name = "crc32table",
    srcs = [],
    outs = ["crc32table.h"],
    cmd = "./$(location :gen_crc32table) > \"$@\"",
    tools = [":gen_crc32table"],
)

cc_library(
    name = "crc32c",
    srcs = [
        "include/xfs_arch.h",
        "libfrog/crc32.c",
        "libfrog/crc32defs.h",
        ":crc32table",
        ":platform_defs.h",
    ],
    hdrs = ["libfrog/crc32c.h"],
    local_defines = defs,
)

cc_library(
    name = "list_sort",
    srcs = ["libfrog/list_sort.c"],
    hdrs = ["include/list.h"],
    local_defines = defs,
)

cc_library(
    name = "fsgeom",
    srcs = ["libfrog/fsgeom.c"],
    hdrs = ["libfrog/fsgeom.h"],
    local_defines = defs,
    deps = [
        ":libxfs",
        ":util",
    ],
)

cc_library(
    name = "projects",
    srcs = [
        "include/input.h",
        "libfrog/projects.c",
        "libfrog/projects.h",
        ":platform_defs.h",
    ],
    hdrs = ["libfrog/projects.h"],
    local_defines = defs,
    deps = [
        ":libxfs",
    ],
)

cc_library(
    name = "convert",
    srcs = [
        "include/input.h",
        "libfrog/convert.c",
        ":platform_defs.h",
    ],
    hdrs = ["libfrog/convert.h"],
    local_defines = defs,
    deps = [
        ":projects",
    ],
)

cc_library(
    name = "topology",
    srcs = [
        "include/xfs_multidisk.h",
        "libfrog/topology.c",
    ],
    hdrs = [
        "include/libxcmd.h",
        "libfrog/topology.h",
    ],
    local_defines = defs,
    visibility = ["//visibility:public"],
    deps = [
        ":libxfs",
        "@util_linux//:blkid",
    ],
)

cc_library(
    name = "platform",
    srcs = ["libfrog/linux.c"],
    local_defines = defs,
    visibility = ["//visibility:public"],
    deps = [":libxfs"],
)

cc_library(
    name = "libxfs",
    srcs = glob([
        "libxfs/*.c",
        "libxfs/*.h",
    ]) + [
        ":platform_defs.h",
        "include/xfs.h",
        "libfrog/platform.h",
        "include/linux.h",
        "include/hlist.h",
        "include/cache.h",
        "include/bitops.h",
        "include/kmem.h",
        "include/atomic.h",
        "include/xfs_mount.h",
        "include/xfs_inode.h",
        "include/xfs_trans.h",
        "include/xfs_trace.h",
        "libfrog/linux.c",
        "include/xfs_fs_compat.h",
    ],
    hdrs = ["include/libxfs.h"],
    local_defines = defs,
    deps = [
        ":crc32c",
        ":list_sort",
        ":radix_tree",
        "@util_linux//:uuid",
    ],
)

cc_binary(
    name = "mkfs",
    srcs = [
        "include/xfs_multidisk.h",
        "mkfs/proto.c",
        "mkfs/xfs_mkfs.c",
    ],
    linkopts = ["-lpthread"],
    local_defines = defs,
    deps = [
        ":convert",
        ":fsgeom",
        ":libxfs",
        ":platform",
        ":topology",
        ":util",
        "@inih",
    ],
    visibility = ["//visibility:public"],
)
