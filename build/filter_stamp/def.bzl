def _filtered_stamp_impl(ctx):
    output = ctx.actions.declare_file(ctx.label.name + ".txt")
    ctx.actions.run(
        mnemonic = "FilterStamp",
        executable = ctx.executable._filter_stamp,
        arguments = ["-status", ctx.info_file.path, "-out", output.path, "--"] + ctx.attr.vars,
        inputs = [ctx.info_file],
        outputs = [output],
    )
    return [DefaultInfo(files = depset([output]))]

filtered_stamp = rule(
    implementation = _filtered_stamp_impl,
    doc = """
        Build a stamp file with a subset of
        variables from the stable status file.
    """,
    attrs = {
        "vars": attr.string_list(
            doc = """
                List of variables to include in the output.
            """,
            mandatory = True,
        ),
        # Tool
        "_filter_stamp": attr.label(
            default = Label("//build/filter_stamp"),
            executable = True,
            cfg = "exec",
        ),
    },
)
