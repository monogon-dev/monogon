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

load("@bazel_gazelle//:deps.bzl", "go_repository")
load(
    "@io_bazel_rules_go//go:def.bzl",
    "GoLibrary",
    "go_context",
    "go_library",
)

def _bindata_impl(ctx):
    out = ctx.actions.declare_file("bindata.go")

    arguments = ctx.actions.args()
    arguments.add_all([
        "-pkg",
        ctx.attr.package,
        "-prefix",
        ctx.label.workspace_root,
        "-o",
        out,
    ])
    arguments.add_all(ctx.files.srcs)

    ctx.actions.run(
        inputs = ctx.files.srcs,
        outputs = [out],
        executable = ctx.file.bindata,
        arguments = [arguments],
    )

    go = go_context(ctx)

    source_files = [out]

    library = go.new_library(
        go,
        srcs = source_files,
    )
    source = go.library_to_source(go, None, library, False)
    providers = [library, source]
    output_groups = {
        "go_generated_srcs": source_files,
    }

    return providers + [OutputGroupInfo(**output_groups)]

bindata = rule(
    implementation = _bindata_impl,
    attrs = {
        "srcs": attr.label_list(
            mandatory = True,
            allow_files = True,
         ),
        "package": attr.string(
            mandatory = True,
         ),
        "bindata": attr.label(
            allow_single_file = True,
            default = Label("@com_github_kevinburke_go_bindata//go-bindata"),
        ),
        "_go_context_data": attr.label(
            default = "@io_bazel_rules_go//:go_context_data",
        ),
    },
    toolchains = ["@io_bazel_rules_go//go:toolchain"],
)
