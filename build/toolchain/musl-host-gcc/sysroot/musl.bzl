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

def _musl_headers(ctx):
    hdrs_name = ctx.attr.name + "_headers"
    hdrs_dir = ctx.actions.declare_directory(hdrs_name)

    root = detect_root(ctx.attr.src)
    ctx.actions.run_shell(
        inputs = ctx.files.src,
        outputs = [hdrs_dir],
        progress_message = "Collecting musl headers",
        mnemonic = "MuslCollectHeaders",
        arguments = [root, ctx.attr.arch, hdrs_dir.path],
        use_default_shell_env = True,
        command = "make -C \"$1\" install-headers ARCH=\"$2\" includedir=\"$(pwd)/$3\" > /dev/null",
    )
    return [DefaultInfo(files = depset([hdrs_dir]))]

musl_headers = rule(
    implementation = _musl_headers,
    attrs = {
        "src": attr.label(mandatory = True),
        "arch": attr.string(mandatory = True),
    },
)
