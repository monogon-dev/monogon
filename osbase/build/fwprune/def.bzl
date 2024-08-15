load("//osbase/build:def.bzl", "FSSpecInfo")

def _fsspec_linux_firmware(ctx):
    fsspec_out = ctx.actions.declare_file(ctx.label.name + ".prototxt")

    fwlist = ctx.actions.declare_file(ctx.label.name + "-fwlist.txt")
    ctx.actions.write(
        output = fwlist,
        content = "\n".join([f.path for f in ctx.files.firmware_files]),
    )

    modinfo = ctx.attr.kernel[OutputGroupInfo].modinfo.to_list()[0]
    modules = ctx.attr.kernel[OutputGroupInfo].modules.to_list()[0]

    meta_out = ctx.actions.declare_file(ctx.label.name + "-meta.pb")

    ctx.actions.run(
        outputs = [fsspec_out, meta_out],
        inputs = [fwlist, modinfo, modules, ctx.file.metadata] + ctx.files.firmware_files,
        tools = [ctx.executable._fwprune],
        executable = ctx.executable._fwprune,
        arguments = [
            "-modinfo",
            modinfo.path,
            "-modules",
            modules.path,
            "-firmware-file-list",
            fwlist.path,
            "-firmware-whence",
            ctx.file.metadata.path,
            "-out-meta",
            meta_out.path,
            "-out-fsspec",
            fsspec_out.path,
        ],
    )

    return [DefaultInfo(files = depset([fsspec_out])), FSSpecInfo(spec = fsspec_out, referenced = ctx.files.firmware_files + [modules, meta_out])]

fsspec_linux_firmware = rule(
    implementation = _fsspec_linux_firmware,
    doc = """
         Generates a partial filesystem spec containing all firmware files required by a given kernel at the
         default firmware load path (/lib/firmware).
    """,
    attrs = {
        "firmware_files": attr.label_list(
            mandatory = True,
            allow_files = True,
            doc = """
               List of firmware files. Generally at least a filegroup of the linux-firmware repository should
               be in here.
            """,
        ),
        "metadata": attr.label(
            mandatory = True,
            allow_single_file = True,
            doc = """
                The metadata file for the Linux firmware. Currently this is the WHENCE file at the root of the
                linux-firmware repository. Used for resolving additional links.
            """,
        ),
        "kernel": attr.label(
            doc = """
                Kernel for which firmware should be selected. Needs to have a modinfo OutputGroup.
            """,
        ),

        # Tool
        "_fwprune": attr.label(
            default = Label("//osbase/build/fwprune"),
            executable = True,
            cfg = "exec",
        ),
    },
)
