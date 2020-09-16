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

load("@bazel_tools//tools/cpp:cc_toolchain_config_lib.bzl", "tool", "tool_path")

# This defines a minimal, barely parametrized toolchain configuration rule that
# uses the host GCC with some possible overrides.

def _host_cc_toolchain_impl(ctx):
    tool_paths = [
        tool_path(
            name = "gcc",
            path = ctx.attr.gcc,
        ),
        tool_path(
            name = "ld",
            path = "/usr/bin/ld",
        ),
        tool_path(
            name = "ar",
            path = "/usr/bin/ar",
        ),
        tool_path(
            name = "cpp",
            path = "/bin/false",
        ),
        tool_path(
            name = "gcov",
            path = "/bin/false",
        ),
        tool_path(
            name = "nm",
            path = "/bin/false",
        ),
        tool_path(
            name = "objdump",
            path = "/bin/false",
        ),
        tool_path(
            name = "strip",
            path = "/bin/false",
        ),
    ]
    return cc_common.create_cc_toolchain_config_info(
        ctx = ctx,
        cxx_builtin_include_directories = ctx.attr.host_includes,
        toolchain_identifier = "k8-toolchain",
        host_system_name = "local",
        target_system_name = "local",
        target_cpu = "k8",
        target_libc = "unknown",
        compiler = "gcc",
        abi_version = "unknown",
        abi_libc_version = "unknown",
        tool_paths = tool_paths,
        builtin_sysroot = ctx.attr.sysroot,
    )

host_cc_toolchain_config = rule(
    implementation = _host_cc_toolchain_impl,
    attrs = {
        "gcc": attr.string(
            default = "/usr/bin/gcc",
        ),
        "host_includes": attr.string_list(
            default = [
                "/usr/lib/gcc/x86_64-redhat-linux/10/include/",
                "/usr/include",
            ],
        ),
        "sysroot": attr.string(
            default = "",
        ),
    },
    provides = [CcToolchainConfigInfo],
)
