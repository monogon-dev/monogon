load(
    "@io_bazel_rules_go//go:def.bzl",
    "GoLibrary",
    "go_context",
    "go_library",
)

def _go_version_library_impl(ctx):
    output = ctx.actions.declare_file(ctx.attr.name + "_generated.go")

    ctx.actions.run(
        mnemonic = "GenVersion",
        progress_message = "Generating version file",
        inputs = [ctx.info_file],
        outputs = [output],
        executable = ctx.executable._genversion,
        arguments = [
            "-importpath",
            ctx.attr.importpath,
            "-product",
            ctx.attr.product,
            "-status_file",
            ctx.info_file.path,
            "-out_file",
            output.path,
        ],
    )

    go = go_context(ctx)
    source_files = [output]
    library = go.new_library(
        go,
        srcs = source_files,
    )
    source = go.library_to_source(go, ctx.attr, library, False)
    providers = [library, source]
    output_groups = {
        "go_generated_srcs": source_files,
    }
    return providers + [OutputGroupInfo(**output_groups)]

go_version_library = rule(
    doc = """
        Generate a Go library target which can be further embedded/depended upon
        by other Go code. This library contains a Version proto field which will
        be automatically populated with version based on build state data.
    """,
    implementation = _go_version_library_impl,
    attrs = {
        "importpath": attr.string(
            mandatory = True,
        ),
        "product": attr.string(
            mandatory = True,
            doc = """
                Name of Monogon product that for which this version library will
                be generated. This must correspond to the product name as used in
                Git tags, which in turn is used to extract a release version
                during a build.
            """,
        ),
        "_genversion": attr.label(
            default = Label("//version/stampgo"),
            cfg = "host",
            executable = True,
            allow_files = True,
        ),
        "_go_context_data": attr.label(
            default = "@io_bazel_rules_go//:go_context_data",
        ),
        "deps": attr.label_list(
            default = [
                "@org_golang_google_protobuf//proto",
                "//version/spec",
            ],
        ),
    },
    toolchains = ["@io_bazel_rules_go//go:toolchain"],
)
