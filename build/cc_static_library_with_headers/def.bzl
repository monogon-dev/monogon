load("@rules_cc//cc/common:cc_info.bzl", "CcInfo")

# This is a workaround for the linux kernel build, as it requires static
# linked libraries but also their headers. This is fairly cursed and should
# be removed as fast as possible.
def cc_static_library_with_headers(name, dep):
    # The artifact name has to be the same as the name of the
    # static library, so that the linker can find it.
    artifact_name = Label(dep).name
    native.cc_static_library(
        name = artifact_name,
        deps = [dep],
    )

    _cc_static_library_wrapper(
        name = name,
        static_library = artifact_name,
        library = dep,
    )

def _cc_static_library_wrapper_impl(ctx):
    return [
        ctx.attr.static_library[DefaultInfo],
        ctx.attr.static_library[OutputGroupInfo],
        ctx.attr.library[CcInfo],
    ]

_cc_static_library_wrapper = rule(
    implementation = _cc_static_library_wrapper_impl,
    attrs = {
        "static_library": attr.label(),
        "library": attr.label(),
    },
)
