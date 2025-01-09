def _build_efi_transition_impl(_settings, _attr):
    """
    Transition that enables building for an EFI environment. Currently only supports C code.
    """
    return {
        "//command_line_option:platforms": "//build/platforms:efi_amd64",
    }

build_efi_transition = transition(
    implementation = _build_efi_transition_impl,
    inputs = [],
    outputs = [
        "//command_line_option:platforms",
    ],
)
