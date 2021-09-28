load(":cc_toolchain_config.bzl", "efi_k8_cc_toolchain_config")

package(default_visibility = ["//visibility:public"])

filegroup(name = "empty")

cc_toolchain_suite(
    name = "efi_cc_suite",
    toolchains = {
        "k8": ":efi_k8_cc_toolchain",
    },
)

filegroup(
    name = "fltused",
    srcs = ["fltused.o"],
)

cc_toolchain(
    name = "efi_k8_cc_toolchain",
    all_files = ":empty",
    compiler_files = ":empty",
    dwp_files = ":empty",
    linker_files = ":fltused",
    objcopy_files = ":empty",
    strip_files = ":empty",
    supports_param_files = 0,
    toolchain_config = ":efi_k8_cc_toolchain_config",
    toolchain_identifier = "efi-k8-toolchain",
)

efi_k8_cc_toolchain_config(name = "efi_k8_cc_toolchain_config")
