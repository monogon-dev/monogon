def _build_efi_transition_impl(settings, attr):
    """
    Transition that enables building for an EFI environment. Currently ony supports C code.
    """
    return {
        "//command_line_option:crosstool_top": "//build/toolchain/llvm-efi:efi_cc_suite",
    }

build_efi_transition = transition(
    implementation = _build_efi_transition_impl,
    inputs = [],
    outputs = [
        "//command_line_option:crosstool_top",
    ],
)
