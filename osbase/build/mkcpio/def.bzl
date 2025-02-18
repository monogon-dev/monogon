load("//osbase/build:def.bzl", "build_pure_transition", "build_static_transition")
load("//osbase/build/fsspec:def.bzl", "FSSpecInfo", "fsspec_core_impl")

def _node_initramfs_impl(ctx):
    initramfs_name = ctx.label.name + ".cpio.zst"
    initramfs = ctx.actions.declare_file(initramfs_name)

    fsspec_core_impl(ctx, ctx.executable._mkcpio, initramfs)

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
        "files": attr.string_keyed_label_dict(
            mandatory = True,
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

        # Tool
        "_mkcpio": attr.label(
            default = Label("//osbase/build/mkcpio"),
            executable = True,
            cfg = "exec",
        ),
    },
)
