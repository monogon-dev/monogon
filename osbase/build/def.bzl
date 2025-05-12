load("@bazel_skylib//lib:paths.bzl", "paths")

def _build_static_transition_impl(_settings, _attr):
    """
    Transition that enables static build of Go and C binaries.
    """
    return {
        "@io_bazel_rules_go//go/config:static": True,
        "//build/platforms/linkmode:static": True,
    }

build_static_transition = transition(
    implementation = _build_static_transition_impl,
    inputs = [],
    outputs = [
        "@io_bazel_rules_go//go/config:static",
        "//build/platforms/linkmode:static",
    ],
)

def forward_impl(ctx):
    # We can't pass DefaultInfo through as-is, since Bazel forbids executable
    # if it's a file declared in a different target. To emulate that, symlink
    # to the original executable, if there is one.
    default_info = ctx.attr.dep[DefaultInfo]
    new_executable = None
    original_executable = default_info.files_to_run.executable
    runfiles = default_info.default_runfiles
    if original_executable:
        # In order for the symlink to have the same basename as the original
        # executable (important in the case of proto plugins), put it in a
        # subdirectory named after the label to prevent collisions.
        new_executable = ctx.actions.declare_file(paths.join(ctx.label.name, original_executable.basename))
        ctx.actions.symlink(
            output = new_executable,
            target_file = original_executable,
            is_executable = True,
        )
        runfiles = runfiles.merge(ctx.runfiles([new_executable]))

    return [DefaultInfo(
        files = default_info.files,
        runfiles = runfiles,
        executable = new_executable,
    )]

build_static_target = rule(
    cfg = build_static_transition,
    implementation = forward_impl,
    attrs = {
        "dep": attr.label(mandatory = True),
    },
    doc = """Applies build_static_transition to a target.""",
)

_new_settings = {
    # This list should be expanded with any configuration options that end
    # up reaching this rule with different values across different build
    # graph paths, but that do not actually influence the kernel build.
    # Force-setting them to a stable value forces the build configuration
    # to a stable hash.
    # See the transition's comment block for more information.
    "@io_bazel_rules_go//go/config:static": False,
    "//build/platforms/linkmode:static": False,
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

ignore_unused_configuration_target = rule(
    cfg = ignore_unused_configuration,
    implementation = forward_impl,
    attrs = {
        "dep": attr.label(mandatory = True),
    },
    doc = """Applies ignore_unused_configuration transition to a target.""",
)
