def _localregistry_manifest_impl(ctx):
    manifest_out = ctx.actions.declare_file(ctx.label.name + ".prototxt")

    images = []
    referenced = [manifest_out]
    for key, label in ctx.attr.images.items():
        image_file = label[DefaultInfo].files.to_list()[0]
        repository, _, tag = key.partition(":")
        image = struct(
            repository = repository,
            tag = tag,
            path = image_file.short_path,
        )
        referenced.append(image_file)
        images.append(image)

    ctx.actions.write(manifest_out, proto.encode_text(struct(images = images)))
    return [DefaultInfo(runfiles = ctx.runfiles(files = referenced), files = depset([manifest_out]))]

localregistry_manifest = rule(
    implementation = _localregistry_manifest_impl,
    doc = """
        Builds a manifest for serving images directly from the build files.
    """,
    attrs = {
        "images": attr.string_keyed_label_dict(
            mandatory = True,
            doc = """
                Images to be served from the local registry.
                The key defines the repository and tag, separated by ':'.
                The value is a label which must contain an OCI layout directory.
            """,
            providers = [],
        ),
    },
)
