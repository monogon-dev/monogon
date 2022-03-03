load("//metropolis/node/build:def.bzl", "FSSpecInfo")

def _fsspec_linux_firmware(ctx):
    fsspec_out = ctx.actions.declare_file(ctx.label.name + ".prototxt")

    fwlist = ctx.actions.declare_file(ctx.label.name + "-fwlist.txt")
    ctx.actions.write(
        output = fwlist,
        content = "\n".join([f.path for f in ctx.files.firmware_files]),
    )

    modinfo = ctx.attr.kernel[OutputGroupInfo].modinfo.to_list()[0]

    ctx.actions.run(
        outputs = [fsspec_out],
        inputs = [fwlist, modinfo, ctx.file.metadata] + ctx.files.firmware_files,
        tools = [ctx.executable._fwprune],
        executable = ctx.executable._fwprune,
        arguments = [modinfo.path, fwlist.path, ctx.file.metadata.path, fsspec_out.path],
    )

    return [DefaultInfo(files = depset([fsspec_out])), FSSpecInfo(spec = fsspec_out, referenced = ctx.files.firmware_files)]

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
            default = Label("//metropolis/node/build/fwprune"),
            executable = True,
            cfg = "exec",
        ),
    },
)
