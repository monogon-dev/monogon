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
    "@io_bazel_rules_go//go:def.bzl",
    _go_context = "go_context",
)
load(
    "@bazel_skylib//lib:shell.bzl",
    "shell",
)

def _fietsje_runner_impl(ctx):
    go = _go_context(ctx)
    out_file = ctx.actions.declare_file(ctx.label.name + ".bash")
    substitutions = {
        "@@GOTOOL@@": shell.quote(go.go.path),
        "@@FIETSJE_SHORT_PATH@@": shell.quote(ctx.executable.fietsje.short_path),
    }
    ctx.actions.expand_template(
        template = ctx.file._template,
        output = out_file,
        substitutions = substitutions,
        is_executable = True,
    )
    runfiles = ctx.runfiles(files = [
        ctx.executable.fietsje,
        go.go,
    ])
    return [DefaultInfo(
        files = depset([out_file]),
        runfiles = runfiles,
        executable = out_file,
    )]

_fietsje_runner = rule(
    implementation = _fietsje_runner_impl,
    attrs = {
        "fietsje": attr.label(
            default = "//build/fietsje",
            executable = True,
            cfg = 'host',
        ),
        "_template": attr.label(
            default = "//build/fietsje:fietsje.bash.in",
            allow_single_file = True,
        ),
        "_go_context_data": attr.label(
            default = "@io_bazel_rules_go//:go_context_data",
        ),
    },
    toolchains = ["@io_bazel_rules_go//go:toolchain"],
)

def fietsje(name):
    runner_name = name + "-runner"
    _fietsje_runner(
        name = runner_name,
    )
    native.sh_binary(
        name = name,
        srcs = [runner_name],
    )
