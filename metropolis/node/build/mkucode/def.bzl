def _cpio_ucode_impl(ctx):
    ucode_spec = ctx.actions.declare_file(ctx.label.name + "_spec.prototxt")

    vendors = []
    inputs = []
    for label, vendor in ctx.attr.ucode.items():
        files = label[DefaultInfo].files.to_list()
        inputs += files
        vendors.append(struct(id = vendor, file = [f.path for f in files]))

    ctx.actions.write(ucode_spec, proto.encode_text(struct(vendor = vendors)))

    output_file = ctx.actions.declare_file(ctx.label.name + ".cpio")
    ctx.actions.run(
        outputs = [output_file],
        inputs = [ucode_spec] + inputs,
        tools = [ctx.executable._mkucode],
        executable = ctx.executable._mkucode,
        arguments = ["-out", output_file.path, "-spec", ucode_spec.path],
    )
    return [DefaultInfo(files = depset([output_file]))]

cpio_ucode = rule(
    implementation = _cpio_ucode_impl,
    doc = """
        Builds a cpio archive with microcode for the Linux early microcode loader.
    """,
    attrs = {
        "ucode": attr.label_keyed_string_dict(
            mandatory = True,
            allow_files = True,
            doc = """
                Dictionary of Labels to String. Each label is a list of microcode files and the string label
                is the vendor ID corresponding to that microcode.
            """,
        ),

        # Tool
        "_mkucode": attr.label(
            default = Label("//metropolis/node/build/mkucode"),
            executable = True,
            cfg = "exec",
        ),
    },
)
