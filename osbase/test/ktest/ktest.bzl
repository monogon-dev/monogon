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

"""
Ktest provides a macro to run tests under a normal Metropolis node kernel
"""

load("//osbase/build:def.bzl", "FSSpecInfo", "build_pure_transition", "build_static_transition", "fsspec_core_impl")

_KTEST_SCRIPT = """
#!/usr/bin/env bash

exec "{ktest}" -initrd-path "{initrd}" -kernel-path "{kernel}" -cmdline "{cmdline}"
"""

def _ktest_impl(ctx):
    initramfs_name = ctx.label.name + ".cpio.zst"
    initramfs = ctx.actions.declare_file(initramfs_name)

    fsspec_core_impl(ctx, ctx.executable._mkcpio, initramfs, [("/init", ctx.attr._ktest_init[0]), ("/tester", ctx.attr.tester[0])], [ctx.attr._earlydev])

    script_file = ctx.actions.declare_file(ctx.label.name + ".sh")

    ctx.actions.write(
        output = script_file,
        content = _KTEST_SCRIPT.format(
            ktest = ctx.executable._ktest.short_path,
            initrd = initramfs.short_path,
            kernel = ctx.file.kernel.short_path,
            cmdline = ctx.attr.cmdline,
        ),
        is_executable = True,
    )

    return [DefaultInfo(
        executable = script_file,
        runfiles = ctx.runfiles(
            files = [ctx.files._ktest[0], initramfs, ctx.file.kernel, ctx.file.tester],
        ),
    )]

k_test = rule(
    implementation = _ktest_impl,
    doc = """
        Run a given test program under the Monogon kernel. 
    """,
    attrs = {
        "tester": attr.label(
            mandatory = True,
            executable = True,
            allow_single_file = True,
            # Runs inside the given kernel, needs to be build for Linux/static
            cfg = build_static_transition,
        ),
        "files": attr.string_keyed_label_dict(
            allow_files = True,
            doc = """
                Dictionary of Labels to String, placing a given Label's output file in the initramfs at the location
                specified by the String value. The specified labels must only have a single output.
            """,
            # Attach pure transition to ensure all binaries added to the initramfs are pure/static binaries.
            cfg = build_pure_transition,
        ),
        "files_cc": attr.string_keyed_label_dict(
            allow_files = True,
            doc = """
                 Special case of 'files' for compilation targets that need to be built with the musl toolchain like
                 go_binary targets which need cgo or cc_binary targets.
            """,
            # Attach static transition to all files_cc inputs to ensure they are built with musl and static.
            cfg = build_static_transition,
        ),
        "symlinks": attr.string_dict(
            default = {},
            doc = """
                Symbolic links to create. Similar format as in files and files_cc, so the target of the symlink is the
                key and the value of it is the location of the symlink itself. Only raw strings are allowed as targets,
                labels are not permitted. Include the file using files or files_cc, then symlink to its location.
            """,
        ),
        "fsspecs": attr.label_list(
            default = [],
            doc = """
                List of file system specs (osbase.build.fsspec.FSSpec) to also include in the resulting image.
                These will be merged with all other given attributes.
            """,
            providers = [FSSpecInfo],
            allow_files = True,
        ),
        "kernel": attr.label(
            default = Label("//osbase/test/ktest:linux-testing"),
            cfg = "exec",
            allow_single_file = True,
        ),
        "cmdline": attr.string(
            default = "",
        ),
        # Tool
        "_ktest": attr.label(
            default = Label("//osbase/test/ktest"),
            cfg = "exec",
            executable = True,
            allow_files = True,
        ),
        "_ktest_init": attr.label(
            default = Label("//osbase/test/ktest/init"),
            cfg = build_pure_transition,
            executable = True,
            allow_single_file = True,
        ),
        "_mkcpio": attr.label(
            default = Label("//osbase/build/mkcpio"),
            executable = True,
            cfg = "exec",
        ),
        "_earlydev": attr.label(
            default = Label("//osbase/build:earlydev.fsspec"),
            allow_files = True,
        ),
    },
    test = True,
)
