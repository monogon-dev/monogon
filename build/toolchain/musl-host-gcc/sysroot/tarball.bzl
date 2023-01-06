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

load(
    "//build/utils:detect_root.bzl",
    "detect_root",
)

"""
Build a sysroot-style tarball containing musl/linux headers and libraries.

This can then be used to build a C toolchain that builds C/C++ binaries for Metropolis nodes.
"""

def _musl_gcc_tarball(ctx):
    tarball_name = ctx.attr.name + ".tar.xz"
    tarball = ctx.actions.declare_file(tarball_name)

    musl_headers = ctx.file.musl_headers
    musl_headers_path = musl_headers.path
    linux_headers = ctx.file.linux_headers
    linux_headers_path = linux_headers.path

    compiler_headers_path = "lib/gcc/x86_64-redhat-linux/12/include"

    musl_root = detect_root(ctx.attr.musl)
    musl_files = ctx.files.musl

    # This builds a tarball containing musl, musl headers and linux headers.
    # This is done by some carefully crafted tar command line arguments that rewrite
    # paths to ensure that everything lands in lib/ and include/ in the tarball.

    # TODO(q3k): write nice, small static Go utility for this.

    arguments = [tarball.path]
    command = "tar -chJf $1"

    # Order is important here as this is a terrible hack producing a tar file with duplicate files. The decompressor
    # will then overwrite the wrong one with the correct one for us.
    arguments += [compiler_headers_path]
    command += " --transform 's|^'$2'|include|' /$2"

    arguments += [musl_headers_path]
    command += " --transform 's|^'$3'|include|' $3"

    arguments += [linux_headers_path]
    command += " --transform 's|^'$4'|include|' $4"

    arguments += [musl_root]
    command += " --transform 's|^'$5'|lib|' $5"

    ctx.actions.run_shell(
        inputs = [musl_headers, linux_headers] + ctx.files.musl,
        outputs = [tarball],
        progress_message = "Building toolchain tarball",
        mnemonic = "BuildToolchainTarball",
        arguments = arguments,
        use_default_shell_env = True,
        command = command,
    )
    return [DefaultInfo(files = depset([tarball]))]

musl_gcc_tarball = rule(
    implementation = _musl_gcc_tarball,
    attrs = {
        "musl": attr.label(mandatory = True),
        "musl_headers": attr.label(mandatory = True, allow_single_file = True),
        "linux_headers": attr.label(mandatory = True, allow_single_file = True),
    },
)
