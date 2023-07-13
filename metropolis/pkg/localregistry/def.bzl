load("@io_bazel_rules_docker//container:providers.bzl", "ImageInfo")

def _localregistry_manifest_impl(ctx):
    manifest_out = ctx.actions.declare_file(ctx.label.name+".prototxt")
    
    images = []
    referenced = [manifest_out]
    for i in ctx.attr.images:
        image_info = i[ImageInfo].container_parts
        referenced.append(image_info['config'])
        referenced.append(image_info['config_digest'])
        image = struct(
            name = i.label.package + "/" + i.label.name,
            config = struct(
                file_path = image_info['config'].short_path,
                digest_path = image_info['config_digest'].short_path,
            ),
            layers = [],
        )
        for layer in zip(image_info['zipped_layer'], image_info['blobsum']):
            referenced.append(layer[0])
            referenced.append(layer[1])
            image.layers.append(struct(file_path = layer[0].short_path, digest_path=layer[1].short_path))
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
                List of images (with ImageInfo provider) to be served from the local registry.
            """,
           providers = [ImageInfo],
        ),
    },
)
