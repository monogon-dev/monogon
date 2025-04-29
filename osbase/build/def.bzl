def _build_static_transition_impl(_settings, _attr):
    """
    Transition that enables static build of Go and C binaries.
    """
    return {
        "@io_bazel_rules_go//go/config:static": True,
        "//command_line_option:platforms": "//build/platforms:linux_amd64_static",
    }

build_static_transition = transition(
    implementation = _build_static_transition_impl,
    inputs = [],
    outputs = [
        "@io_bazel_rules_go//go/config:static",
        "//command_line_option:platforms",
    ],
)

_new_settings = {
    # This list should be expanded with any configuration options that end
    # up reaching this rule with different values across different build
    # graph paths, but that do not actually influence the kernel build.
    # Force-setting them to a stable value forces the build configuration
    # to a stable hash.
    # See the transition's comment block for more information.
    "@io_bazel_rules_go//go/config:pure": False,
    "@io_bazel_rules_go//go/config:static": False,

    # Note: this toolchain is not actually used to perform the build.
    "//command_line_option:platforms": "//build/platforms:linux_amd64_static",
}

def _ignore_unused_configuration_impl(_settings, _attr):
    return _new_settings

# Transition to flip all known-unimportant but varying configuration options to
# a known, stable value.
# This is to prevent Bazel from creating extra configurations for possible
# combinations of options in case the linux_image rule is pulled through build
# graph fragments that have different options set.
#
# Ideally, Bazel would let us mark in a list that we only care about some set
# of options (or at least let us mark those that we explicitly don't care
# about, instead of manually setting them to some value). However, this doesn't
# seem to be possible, thus this transition is a bit of a hack.
ignore_unused_configuration = transition(
    implementation = _ignore_unused_configuration_impl,
    inputs = [],
    outputs = list(_new_settings.keys()),
)
