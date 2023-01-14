load("@rules_proto//proto:defs.bzl", "ProtoInfo")

def _proto_docs(ctx):
    protos = [proto[ProtoInfo] for proto in ctx.attr.protos]
    transitive_sources = depset(transitive = [proto.transitive_sources for proto in protos])
    transitive_proto_path = depset(transitive = [proto.transitive_proto_path for proto in protos])

    out = ctx.actions.declare_file(ctx.label.name + ".html")

    args = []
    args.append("--plugin")
    args.append(ctx.executable._protoc_gen_doc.path)
    args.append("--doc_out")
    args.append(out.dirname)
    args.append("--doc_opt=html," + out.basename)

    for include_path in transitive_proto_path.to_list():
        args.append("-I")
        args.append(include_path)

    for src in transitive_sources.to_list():
        # Due to the built-in import path for well-known types (see AddDefaultProtoPaths
        # in @com_google_protobuf//src/google/protobuf/compiler:command_line_interface.cc)
        # in protoc the Bazel-generated well-known protos are considered to contain
        #  "duplicate" types.
        # Since generating documentation for well-known types is not that useful just
        # skip them.
        if src.path.find("/bin/external/com_github_protocolbuffers_protobuf/_virtual_imports/") != -1:
            continue
        args.append(src.path)

    ctx.actions.run(
        tools = [ctx.executable._protoc_gen_doc],
        inputs = transitive_sources,
        outputs = [out],
        executable = ctx.executable._protoc,
        arguments = args,
    )
    return [DefaultInfo(files = depset([out]))]

proto_docs = rule(
    implementation = _proto_docs,
    doc = """
        Generate a single HTML documentation file documenting all types and services from the transitive set of
        Protobuf files referenced by all proto_libraries passed into `protos`.
    """,
    attrs = {
        "protos": attr.label_list(
            doc = "A list of protobuf libraries for which (and their dependencies) documentation should be generated for",
            providers = [ProtoInfo],
            default = [],
        ),
        "_protoc": attr.label(
            default = Label("@com_google_protobuf//:protoc"),
            cfg = "exec",
            executable = True,
            allow_files = True,
        ),
        "_protoc_gen_doc": attr.label(
            default = Label("@com_github_pseudomuto_protoc_gen_doc//cmd/protoc-gen-doc"),
            cfg = "exec",
            executable = True,
            allow_files = True,
        ),
    },
)
