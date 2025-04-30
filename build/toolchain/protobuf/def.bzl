load("//osbase/build:def.bzl", "forward_impl")

# This is a copy of ignore_unused_configuration_target from
# //osbase/build:def.bzl with specific settings for the protobuf toolchain.
_new_settings = {
    "@io_bazel_rules_go//go/config:race": False,
    "@io_bazel_rules_go//go/config:pure": False,
    "@io_bazel_rules_go//go/config:static": False,
    "@io_bazel_rules_go//go/config:tags": [],

    # These private configs show up because of a bug in rules_go, which is
    # missing a non_go_tool_transition on the proto toolchain when
    # --incompatible_enable_proto_toolchain_resolution is enabled.
    "@io_bazel_rules_go//go/private/rules:original_pure": "",
    "@io_bazel_rules_go//go/private/rules:original_tags": "",
}

def _ignore_unused_configuration_impl(_settings, _attr):
    return _new_settings

_ignore_unused_configuration = transition(
    implementation = _ignore_unused_configuration_impl,
    inputs = [],
    outputs = list(_new_settings.keys()),
)

ignore_unused_configuration_target = rule(
    cfg = _ignore_unused_configuration,
    implementation = forward_impl,
    attrs = {
        "dep": attr.label(mandatory = True),
    },
    doc = """Applies ignore_unused_configuration transition to a target.""",
)
