# Copyright The Monogon Project Authors.
# SPDX-License-Identifier: Apache-2.0

def _test_product_info_impl(ctx):
    raw_product_info = json.encode({
        "id": ctx.attr.os_id,
        "name": ctx.attr.os_name,
        "version": "0.0.0",
        "variant": ctx.attr.architecture,
        "architecture": ctx.attr.architecture,
    })
    product_info_file = ctx.actions.declare_file(ctx.label.name + ".json")
    ctx.actions.write(product_info_file, raw_product_info)
    return [DefaultInfo(files = depset([product_info_file]))]

_test_product_info = rule(
    implementation = _test_product_info_impl,
    attrs = {
        "os_name": attr.string(mandatory = True),
        "os_id": attr.string(mandatory = True),
        "architecture": attr.string(mandatory = True),
    },
)

def _test_product_info_macro_impl(**kwargs):
    _test_product_info(
        architecture = select({
            "@platforms//cpu:x86_64": "x86_64",
            "@platforms//cpu:aarch64": "aarch64",
        }),
        **kwargs
    )

test_product_info = macro(
    inherit_attrs = _test_product_info,
    attrs = {"architecture": None},
    implementation = _test_product_info_macro_impl,
    doc = "This is a simplified variant of product_info for use in tests.",
)
