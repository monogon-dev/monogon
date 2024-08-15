def _node_image_impl(ctx):
    img_file = ctx.actions.declare_file(ctx.label.name + ".img")
    ctx.actions.run(
        mnemonic = "MkImage",
        executable = ctx.executable._mkimage,
        arguments = [
            "-efi",
            ctx.file.kernel.path,
            "-system",
            ctx.file.system.path,
            "-out",
            img_file.path,
        ],
        inputs = [
            ctx.file.kernel,
            ctx.file.system,
        ],
        outputs = [img_file],
    )

    return [DefaultInfo(files = depset([img_file]), runfiles = ctx.runfiles(files = [img_file]))]

node_image = rule(
    implementation = _node_image_impl,
    doc = """
        Build a disk image from an EFI kernel payload and system partition
        contents. See //osbase/build/mkimage for more information.
    """,
    attrs = {
        "kernel": attr.label(
            doc = "EFI binary containing a kernel.",
            mandatory = True,
            allow_single_file = True,
        ),
        "system": attr.label(
            doc = "Contents of the system partition.",
            mandatory = True,
            allow_single_file = True,
        ),
        "_mkimage": attr.label(
            doc = "The mkimage executable.",
            default = "//osbase/build/mkimage",
            allow_single_file = True,
            executable = True,
            cfg = "exec",
        ),
    },
)
