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
load("@bazel_skylib//lib:paths.bzl", "paths")

def _build_pure_transition_impl(settings, attr):
    """
    Transition that enables pure, static build of Go binaries.
    """
    race = settings['@io_bazel_rules_go//go/config:race']
    pure = not race

    return {
        "@io_bazel_rules_go//go/config:pure": pure,
        "@io_bazel_rules_go//go/config:static": True,
    }

build_pure_transition = transition(
    implementation = _build_pure_transition_impl,
    inputs = [
        "@io_bazel_rules_go//go/config:race",
    ],
    outputs = [
        "@io_bazel_rules_go//go/config:pure",
        "@io_bazel_rules_go//go/config:static",
    ],
)

def _build_static_transition_impl(settings, attr):
    """
    Transition that enables static builds with CGo and musl for Go binaries.
    """
    return {
        "@io_bazel_rules_go//go/config:static": True,
        "//command_line_option:platforms": "//build/platforms:linux_amd64_static",
    }

build_static_transition = transition(
    implementation = _build_static_transition_impl,
    inputs = [],
    outputs = [
        "@io_bazel_rules_go//go/config:static",
        "//command_line_option:platforms",
    ],
)

FSSpecInfo = provider(
    "Provides parts of an FSSpec used to assemble filesystem images",
    fields = {
        "spec": "File containing the partial FSSpec as prototext",
        "referenced": "Files (potentially) referenced by the spec",
    },
)

def _fsspec_core_impl(ctx, tool, output_file):
    """
    _fsspec_core_impl implements the core of an fsspec-based rule. It takes
    input from the `files`,`files_cc`, `symlinks` and `fsspecs` attributes
    and calls `tool` with the `-out` parameter pointing to `output_file`
    and paths to all fsspecs as positional arguments.
    """
    fs_spec_name = ctx.label.name + ".prototxt"
    fs_spec = ctx.actions.declare_file(fs_spec_name)

    fs_files = []
    inputs = []
    for label, p in ctx.attr.files.items() + ctx.attr.files_cc.items():
        if not p.startswith("/"):
            fail("file {} invalid: must begin with /".format(p))

        # Figure out if this is an executable.
        is_executable = True

        di = label[DefaultInfo]
        if di.files_to_run.executable == None:
            # Generated non-executable files will have DefaultInfo.files_to_run.executable == None
            is_executable = False
        elif di.files_to_run.executable.is_source:
            # Source files will have executable.is_source == True
            is_executable = False

        # Ensure only single output is declared.
        # If you hit this error, figure out a better logic to find what file you need, maybe looking at providers other
        # than DefaultInfo.
        files = di.files.to_list()
        if len(files) > 1:
            fail("file {} has more than one output: {}", p, files)
        src = files[0]
        inputs.append(src)

        mode = 0o555 if is_executable else 0o444
        fs_files.append(struct(path = p, source_path = src.path, mode = mode, uid = 0, gid = 0))

    fs_symlinks = []
    for target, p in ctx.attr.symlinks.items():
        fs_symlinks.append(struct(path = p, target_path = target))

    fs_spec_content = struct(file = fs_files, directory = [], symbolic_link = fs_symlinks)
    ctx.actions.write(fs_spec, proto.encode_text(fs_spec_content))

    extra_specs = []

    for fsspec in ctx.attr.fsspecs:
        if FSSpecInfo in fsspec:
            fsspecInfo = fsspec[FSSpecInfo]
            extra_specs.append(fsspecInfo.spec)
            for f in fsspecInfo.referenced:
                inputs.append(f)
        else:
            # Raw .fsspec prototext. No referenced data allowed.
            di = fsspec[DefaultInfo]
            extra_specs += di.files.to_list()

    ctx.actions.run(
        outputs = [output_file],
        inputs = [fs_spec] + inputs + extra_specs,
        tools = [tool],
        executable = tool,
        arguments = ["-out", output_file.path, fs_spec.path] + [s.path for s in extra_specs],
    )
    return

def _node_initramfs_impl(ctx):
    initramfs_name = ctx.label.name + ".cpio.zst"
    initramfs = ctx.actions.declare_file(initramfs_name)

    _fsspec_core_impl(ctx, ctx.executable._mkcpio, initramfs)

    # TODO(q3k): Document why this is needed
    return [DefaultInfo(runfiles = ctx.runfiles(files = [initramfs]), files = depset([initramfs]))]

node_initramfs = rule(
    implementation = _node_initramfs_impl,
    doc = """
        Build a node initramfs. The initramfs will contain a basic /dev directory and all the files specified by the
        `files` attribute. Executable files will have their permissions set to 0755, non-executable files will have
        their permissions set to 0444. All parent directories will be created with 0755 permissions.
    """,
    attrs = {
        "files": attr.label_keyed_string_dict(
            mandatory = True,
            allow_files = True,
            doc = """
                Dictionary of Labels to String, placing a given Label's output file in the initramfs at the location
                specified by the String value. The specified labels must only have a single output.
            """,
            # Attach pure transition to ensure all binaries added to the initramfs are pure/static binaries.
            cfg = build_pure_transition,
        ),
        "files_cc": attr.label_keyed_string_dict(
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

        # Tool
        "_mkcpio": attr.label(
            default = Label("//osbase/build/mkcpio"),
            executable = True,
            cfg = "exec",
        ),
    },
)

def _erofs_image_impl(ctx):
    fs_name = ctx.label.name + ".img"
    fs_out = ctx.actions.declare_file(fs_name)

    _fsspec_core_impl(ctx, ctx.executable._mkerofs, fs_out)

    return [DefaultInfo(files = depset([fs_out]))]

erofs_image = rule(
    implementation = _erofs_image_impl,
    doc = """
        Build an EROFS. All files specified in files, files_cc and all specified symlinks will be contained.
        Executable files will have their permissions set to 0555, non-executable files will have
        their permissions set to 0444. All parent directories will be created with 0555 permissions.
    """,
    attrs = {
        "files": attr.label_keyed_string_dict(
            mandatory = True,
            allow_files = True,
            doc = """
                Dictionary of Labels to String, placing a given Label's output file in the EROFS at the location
                specified by the String value. The specified labels must only have a single output.
            """,
            # Attach pure transition to ensure all binaries added to the initramfs are pure/static binaries.
            cfg = build_pure_transition,
        ),
        "files_cc": attr.label_keyed_string_dict(
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

        # Tools, implicit dependencies.
        "_mkerofs": attr.label(
            default = Label("//osbase/build/mkerofs"),
            executable = True,
            cfg = "host",
        ),
    },
)

# VerityConfig is emitted by verity_image, and contains a file enclosing a
# singular dm-verity target table.
VerityConfig = provider(
    "Configuration necessary to mount a single dm-verity target.",
    fields = {
        "table": "A file containing the dm-verity target table. See: https://www.kernel.org/doc/html/latest/admin-guide/device-mapper/verity.html",
    },
)

def _verity_image_impl(ctx):
    """
    Create a new file containing the source image data together with the Verity
    metadata appended to it, and provide an associated DeviceMapper Verity target
    table in a separate file, through VerityConfig provider.
    """

    # Run mkverity.
    image = ctx.actions.declare_file(ctx.attr.name + ".img")
    table = ctx.actions.declare_file(ctx.attr.name + ".dmt")
    ctx.actions.run(
        mnemonic = "GenVerityImage",
        progress_message = "Generating a dm-verity image",
        inputs = [ctx.file.source],
        outputs = [
            image,
            table,
        ],
        executable = ctx.file._mkverity,
        arguments = [
            "-input=" + ctx.file.source.path,
            "-output=" + image.path,
            "-table=" + table.path,
            "-data_alias=" + ctx.attr.rootfs_partlabel,
            "-hash_alias=" + ctx.attr.rootfs_partlabel,
        ],
    )

    return [
        DefaultInfo(
            files = depset([image]),
            runfiles = ctx.runfiles(files = [image]),
        ),
        VerityConfig(
            table = table,
        ),
    ]

verity_image = rule(
    implementation = _verity_image_impl,
    doc = """
      Build a dm-verity target image by appending Verity metadata to the source
      image. A corresponding dm-verity target table will be made available
      through VerityConfig provider.
  """,
    attrs = {
        "source": attr.label(
            doc = "A source image.",
            allow_single_file = True,
        ),
        "rootfs_partlabel": attr.string(
            doc = "GPT partition label of the rootfs to be used with dm-mod.create.",
            default = "PARTLABEL=METROPOLIS-SYSTEM-X",
        ),
        "_mkverity": attr.label(
            doc = "The mkverity executable needed to generate the image.",
            default = "//osbase/build/mkverity",
            allow_single_file = True,
            executable = True,
            cfg = "host",
        ),
    },
)

# From Aspect's bazel-lib under Apache 2.0
def _transition_platform_impl(_, attr):
    return {"//command_line_option:platforms": str(attr.target_platform)}

# Transition from any input configuration to one that includes the
# --platforms command-line flag.
_transition_platform = transition(
    implementation = _transition_platform_impl,
    inputs = [],
    outputs = ["//command_line_option:platforms"],
)


def _platform_transition_binary_impl(ctx):
    # We need to forward the DefaultInfo provider from the underlying rule.
    # Unfortunately, we can't do this directly, because Bazel requires that the executable to run
    # is actually generated by this rule, so we need to symlink to it, and generate a synthetic
    # forwarding DefaultInfo.

    result = []
    binary = ctx.attr.binary[0]

    default_info = binary[DefaultInfo]
    files = default_info.files
    new_executable = None
    original_executable = default_info.files_to_run.executable
    runfiles = default_info.default_runfiles

    if not original_executable:
        fail("Cannot transition a 'binary' that is not executable")

    new_executable_name = ctx.attr.basename if ctx.attr.basename else original_executable.basename

    # In order for the symlink to have the same basename as the original
    # executable (important in the case of proto plugins), put it in a
    # subdirectory named after the label to prevent collisions.
    new_executable = ctx.actions.declare_file(paths.join(ctx.label.name, new_executable_name))
    ctx.actions.symlink(
        output = new_executable,
        target_file = original_executable,
        is_executable = True,
    )
    files = depset(direct = [new_executable], transitive = [files])
    runfiles = runfiles.merge(ctx.runfiles([new_executable]))

    result.append(
        DefaultInfo(
            files = files,
            runfiles = runfiles,
            executable = new_executable,
        ),
    )

    return result

platform_transition_binary = rule(
    implementation = _platform_transition_binary_impl,
    attrs = {
        "basename": attr.string(),
        "binary": attr.label(allow_files = True, cfg = _transition_platform),
        "target_platform": attr.label(
            doc = "The target platform to transition the binary.",
            mandatory = True,
        ),
        "_allowlist_function_transition": attr.label(
            default = "@bazel_tools//tools/allowlists/function_transition_allowlist",
        ),
    },
    executable = True,
    doc = "Transitions the binary to use the provided platform.",
)