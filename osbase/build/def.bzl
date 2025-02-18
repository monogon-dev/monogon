def _build_pure_transition_impl(settings, _attr):
    """
    Transition that enables pure, static build of Go binaries.
    """
    race = settings["@io_bazel_rules_go//go/config:race"]
    pure = not race

    return {
        "@io_bazel_rules_go//go/config:pure": pure,
        "@io_bazel_rules_go//go/config:static": True,
        "//command_line_option:platforms": "//build/platforms:linux_amd64_static",
    }

build_pure_transition = transition(
    implementation = _build_pure_transition_impl,
    inputs = [
        "@io_bazel_rules_go//go/config:race",
    ],
    outputs = [
        "@io_bazel_rules_go//go/config:pure",
        "@io_bazel_rules_go//go/config:static",
        "//command_line_option:platforms",
    ],
)

def _build_static_transition_impl(_settings, _attr):
    """
    Transition that enables static builds with CGo and musl for Go binaries.
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
