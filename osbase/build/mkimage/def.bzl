def _node_image_impl(ctx):
    img_file = ctx.actions.declare_file(ctx.label.name + ".img")

    arguments = ctx.actions.args()
    arguments.add_all([
        "-efi",
        ctx.file.kernel.path,
        "-system",
        ctx.file.system.path,
        "-abloader",
        ctx.file.abloader.path,
        "-out",
        img_file.path,
    ])

    if len(ctx.files.bios_bootcode) != 0:
        arguments.add_all(["-bios_bootcode", ctx.file.bios_bootcode.path])

    ctx.actions.run(
        mnemonic = "MkImage",
        executable = ctx.executable._mkimage,
        arguments = [arguments],
        inputs = [
            ctx.file.kernel,
            ctx.file.system,
            ctx.file.abloader,
            ctx.file.bios_bootcode,
        ],
        outputs = [img_file],
    )

    return [DefaultInfo(files = depset([img_file]), runfiles = ctx.runfiles(files = [img_file]))]

node_image = rule(
    implementation = _node_image_impl,
    doc = """
        Build a disk image from an EFI kernel payload, ABLoader and system partition
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
        "abloader": attr.label(
            doc = "ABLoader binary",
            mandatory = True,
            allow_single_file = True,
        ),
        "bios_bootcode": attr.label(
            doc = """
            Optional label to the BIOS bootcode which gets placed at the start of the first block of the image.
            Limited to 440 bytes, padding is not required. It is only used by legacy BIOS boot.
        """,
            mandatory = False,
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
