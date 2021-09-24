load("@io_bazel_rules_go//go:def.bzl", "go_context", "GoSource")

# This implements the toolchain_library rule, which is used to generate a
# rules_go compatible go_library-style target which contains toolchain.go.in
# augmented with information about the Go SDK's toolchain used by Bazel.
#
# The library can then be used to build tools that call the `go` tool, for
# example to perform static analysis or dependency management.

def _toolchain_library_impl(ctx):
    go = go_context(ctx)

    importpath = ctx.attr.importpath

    out = go.declare_file(go, ext = ".go")
    ctx.actions.expand_template(
        template = ctx.file._template,
        output = out,
        substitutions = {
            'GOROOT': go.root,
            'GOTOOL': go.go.path,
        },
    )

    library = go.new_library(go)
    source = go.library_to_source(go, struct(
        srcs = [struct(files = [out])],
        deps = ctx.attr.deps,
    ), library, ctx.coverage_instrumented())

    # Hack: we want to inject runfiles into the generated GoSource, because
    # there's no other way to make rules_go pick up runfiles otherwise.
    runfiles = ctx.runfiles(files = [
        go.go,
        go.sdk_root,
    ] + go.sdk_files)
    source = {
        key: getattr(source, key)
        for key in dir(source)
        if key not in ['to_json', 'to_proto']
    }
    source['runfiles'] = runfiles
    source = GoSource(**source)
    archive = go.archive(go, source)


    return [
        library,
        source,
        archive,
        DefaultInfo(
            files = depset([archive.data.file]),
            runfiles = runfiles,
        ),
        OutputGroupInfo(
            cgo_exports = archive.cgo_exports,
            compilation_outputs = [archive.data.file],
        ),
    ]


toolchain_library = rule(
    implementation = _toolchain_library_impl,
    attrs = {
        "importpath": attr.string(
            mandatory = True,
        ),
        "deps": attr.label_list(),
        "_template": attr.label(
            allow_single_file = True,
            default = ":toolchain.go.in",
        ),
        "_go_context_data": attr.label(
            default = "@io_bazel_rules_go//:go_context_data",
        ),
    },
    toolchains = ["@io_bazel_rules_go//go:toolchain"],
)
