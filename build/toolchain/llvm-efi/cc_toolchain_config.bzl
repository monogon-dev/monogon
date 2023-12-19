load("@bazel_tools//tools/cpp:cc_toolchain_config_lib.bzl", "feature", "flag_group", "flag_set", "tool", "tool_path", "with_feature_set")
load("@bazel_tools//tools/build_defs/cc:action_names.bzl", "ACTION_NAMES")

all_compile_actions = [
    ACTION_NAMES.c_compile,
    ACTION_NAMES.cpp_compile,
    ACTION_NAMES.linkstamp_compile,
    ACTION_NAMES.assemble,
    ACTION_NAMES.preprocess_assemble,
    ACTION_NAMES.cpp_header_parsing,
    ACTION_NAMES.cpp_module_compile,
    ACTION_NAMES.cpp_module_codegen,
    ACTION_NAMES.clif_match,
    ACTION_NAMES.lto_backend,
]

all_cpp_compile_actions = [
    ACTION_NAMES.cpp_compile,
    ACTION_NAMES.linkstamp_compile,
    ACTION_NAMES.cpp_header_parsing,
    ACTION_NAMES.cpp_module_compile,
    ACTION_NAMES.cpp_module_codegen,
    ACTION_NAMES.clif_match,
]

preprocessor_compile_actions = [
    ACTION_NAMES.c_compile,
    ACTION_NAMES.cpp_compile,
    ACTION_NAMES.linkstamp_compile,
    ACTION_NAMES.preprocess_assemble,
    ACTION_NAMES.cpp_header_parsing,
    ACTION_NAMES.cpp_module_compile,
    ACTION_NAMES.clif_match,
]

codegen_compile_actions = [
    ACTION_NAMES.c_compile,
    ACTION_NAMES.cpp_compile,
    ACTION_NAMES.linkstamp_compile,
    ACTION_NAMES.assemble,
    ACTION_NAMES.preprocess_assemble,
    ACTION_NAMES.cpp_module_codegen,
    ACTION_NAMES.lto_backend,
]

all_link_actions = [
    ACTION_NAMES.cpp_link_executable,
    ACTION_NAMES.cpp_link_dynamic_library,
    ACTION_NAMES.cpp_link_nodeps_dynamic_library,
]

lto_index_actions = [
    ACTION_NAMES.lto_index_for_executable,
    ACTION_NAMES.lto_index_for_dynamic_library,
    ACTION_NAMES.lto_index_for_nodeps_dynamic_library,
]

# This defines a relatively minimal EFI toolchain based on host LLVM and no standard library or headers.
def _efi_k8_cc_toolchain_impl(ctx):
    default_compile_flags_feature = feature(
        name = "default_compile_flags",
        enabled = True,
        flag_sets = [
            flag_set(
                actions = all_compile_actions,
                flag_groups = ([
                    flag_group(
                        flags = ["-target", "x86_64-unknown-windows"],
                    ),
                ]),
            ),
            flag_set(
                actions = all_compile_actions,
                flag_groups = ([
                    flag_group(
                        flags = ["-g"],
                    ),
                ]),
                with_features = [with_feature_set(features = ["dbg"])],
            ),
            flag_set(
                actions = all_compile_actions,
                flag_groups = ([
                    flag_group(
                        # Don't bother with O3, this is an EFI toolchain. It's unlikely to gain much performance here
                        # and increases the risk of dangerous optimizations.
                        flags = ["-O2", "-DNDEBUG"],
                    ),
                ]),
                with_features = [with_feature_set(features = ["opt"])],
            ),
        ],
    )

    # "Hybrid" mode disables some MSVC C extensions (but keeps its ABI), but still identifies as MSVC.
    # This is useful if code requires GNU extensions to work which are silently ignored in full MSVC mode.
    # As EFI does not include Windows headers which depend on nonstandard C behavior this should be fine for most code.
    # If this feature is disabled, the toolchain runs with MSVC C extensions fully enabled.
    hybrid_gnu_msvc_feature = feature(
        name = "hybrid_gnu_msvc",
        enabled = True,
        flag_sets = [
            flag_set(
                actions = all_compile_actions,
                flag_groups = [
                    flag_group(
                        flags = ["-D_MSC_VER=1920", "-fno-ms-compatibility", "-fno-ms-extensions"],
                    ),
                ],
            ),
        ],
    )

    default_link_flags_feature = feature(
        name = "default_link_flags",
        enabled = True,
        flag_sets = [
            flag_set(
                actions = all_link_actions + lto_index_actions,
                flag_groups = ([
                    flag_group(
                        flags = [
                            "-target",
                            "x86_64-unknown-windows",
                            "-fuse-ld=lld",
                            "-Wl,-entry:efi_main",
                            "-Wl,-subsystem:efi_application",
                            "-Wl,/BASE:0x0",
                            "-Wl,/Brepro",
                            "-nostdlib",
                            "build/toolchain/llvm-efi/fltused.o",
                        ],
                    ),
                ]),
            ),
        ],
    )

    lto_feature = feature(
        name = "lto",
        enabled = False,
        flag_sets = [
            flag_set(
                actions = all_compile_actions + all_link_actions,
                flag_groups = ([
                    flag_group(
                        flags = [
                            "-flto",
                        ],
                    ),
                ]),
            ),
        ],
    )

    tool_paths = [
        tool_path(
            name = "gcc",
            path = "/usr/bin/clang",
        ),
        tool_path(
            name = "ld",
            path = "/usr/bin/lld-link",
        ),
        tool_path(
            name = "ar",
            path = "/usr/bin/llvm-ar",
        ),
        tool_path(
            name = "cpp",
            path = "/bin/false",
        ),
        tool_path(
            name = "gcov",
            path = "/bin/false",
        ),
        tool_path(
            name = "nm",
            path = "/usr/bin/llvm-nm",
        ),
        tool_path(
            name = "objcopy",
            # We can't use LLVM's objcopy until we pick up https://reviews.llvm.org/D106942
            path = "/usr/bin/objcopy",
        ),
        tool_path(
            name = "objdump",
            path = "/usr/bin/llvm-objdump",
        ),
        tool_path(
            name = "strip",
            path = "/usr/bin/llvm-strip",
        ),
    ]

    return cc_common.create_cc_toolchain_config_info(
        ctx = ctx,
        features = [default_link_flags_feature, default_compile_flags_feature, hybrid_gnu_msvc_feature, lto_feature],
        # Needed for various compiler built-in headers and auxiliary data. No system libraries are being used.
        cxx_builtin_include_directories = [
            "/usr/lib64/clang",
        ],
        toolchain_identifier = "k8-toolchain",
        host_system_name = "local",
        target_system_name = "x86_64-efi",
        target_cpu = "k8",
        target_libc = "none",
        compiler = "clang",
        abi_version = "none",
        abi_libc_version = "none",
        tool_paths = tool_paths,
    )

efi_k8_cc_toolchain_config = rule(
    implementation = _efi_k8_cc_toolchain_impl,
    attrs = {},
    provides = [CcToolchainConfigInfo],
)
