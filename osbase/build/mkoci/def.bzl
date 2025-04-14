def _oci_os_image_impl(ctx):
    inputs = []
    arguments = []
    for name, label in ctx.attr.srcs.items():
        files = label[DefaultInfo].files.to_list()
        if len(files) != 1:
            fail("payload {} does not have exactly one file: {}", name, files)
        file = files[0]
        inputs.append(file)
        arguments += [
            "-payload_name",
            name,
            "-payload_file",
            file.path,
        ]

    arguments += ["-compression_level", str(ctx.attr.compression_level)]
    runfiles = None
    if ctx.attr.compression_level == 0:
        # When not compressing, the inputs are referenced by symlinks.
        runfiles = ctx.runfiles(files = inputs)

    output = ctx.actions.declare_directory(ctx.label.name)
    arguments += ["-out", output.path]

    ctx.actions.run(
        mnemonic = "MkOCI",
        executable = ctx.executable._mkoci,
        arguments = arguments,
        inputs = inputs,
        outputs = [output],
    )

    return [DefaultInfo(
        files = depset([output]),
        runfiles = runfiles,
    )]

oci_os_image = rule(
    implementation = _oci_os_image_impl,
    doc = """
        Build an OS image OCI artifact.
    """,
    attrs = {
        "srcs": attr.string_keyed_label_dict(
            doc = """
                Payloads to include in the OCI artifact.
                The key defines the name of the payload.
                The value is a label which must contain one file.
            """,
            mandatory = True,
            allow_files = True,
        ),
        "compression_level": attr.int(
            default = 2,
            doc = """
                The compression level to use for payloads,
                1 is the fastest, 4 gives the smallest results.
                0 disables compression, payloads are symlinked.
            """,
        ),

        # Tool
        "_mkoci": attr.label(
            default = Label("//osbase/build/mkoci"),
            executable = True,
            cfg = "exec",
        ),
    },
)
