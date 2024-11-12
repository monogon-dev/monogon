def _build_logo_impl(ctx):
    arguments = ctx.actions.args()

    arguments.add_all(["--input"] + ctx.files.logo)
    output = ctx.actions.declare_file("logo.asm")
    arguments.add_all(["--output", output])

    ctx.actions.run(
        outputs = [output],
        inputs = ctx.files.logo,
        arguments = [arguments],
        executable = ctx.executable._genlogo,
    )

    return DefaultInfo(
        files = depset([output]),
    )

    pass

gen_logo = rule(
    implementation = _build_logo_impl,
    attrs = {
        "logo": attr.label(
            allow_single_file = True,
        ),
        "_genlogo": attr.label(
            default = Label(":genlogo"),
            allow_single_file = True,
            executable = True,
            cfg = "exec",
        ),
    },
)
