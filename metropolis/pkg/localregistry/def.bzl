#load("@io_bazel_rules_docker//container:providers.bzl", "ImageInfo")

def _localregistry_manifest_impl(ctx):
    manifest_out = ctx.actions.declare_file(ctx.label.name+".prototxt")

    images = []
    referenced = [manifest_out]
    for i in ctx.attr.images:
        image_file = i.files.to_list()[0]
        image = struct(
            name = i.label.package + "/" + i.label.name,
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
        "images": attr.label_list(
            mandatory = True,
            doc = """
                List of images to be served from the local registry.
            """,
           providers = [],
        ),
    },
)
