load("//build/toolchain:cc_toolchain_config.bzl", "host_cc_toolchain_config")

# This file defines //build/toolchain/musl-host-gcc:musl_host_cc_suite.
#
# This is a C++ toolchain that uses GCC from the host at hardcoded paths, with
# a pre-built sysroot tarball that targets Smalltown with musl and Linux headers.
# It's a superset of //build/toolchain:host_cc_suite.
# For more information, see README.md.

cc_toolchain_suite(
    name = "musl_host_cc_suite",
    toolchains = {
        "k8": ":musl_host_cc_k8_toolchain",
    },
    visibility = ["//visibility:public"],
)

cc_toolchain(
    name = "musl_host_cc_k8_toolchain",
    all_files = ":musl_toolchain_files",
    compiler_files = ":musl_toolchain_files",
    dwp_files = ":musl_toolchain_files",
    linker_files = ":musl_toolchain_files",
    objcopy_files = ":musl_toolchain_files",
    strip_files = ":musl_toolchain_files",
    supports_param_files = 0,
    toolchain_config = ":musl_host_cc_k8_toolchain_config",
    toolchain_identifier = "host-musl-k8-toolchain",
)

host_cc_toolchain_config(
    name = "musl_host_cc_k8_toolchain_config",
    gcc = "gcc-wrapper.sh",
    host_includes = [],
    sysroot = "external/musl_sysroot",
)

filegroup(
    name = "musl_toolchain_files",
    srcs = [
        ":gcc-wrapper.sh",
        ":musl.spec",
        "@musl_sysroot//:all",
    ],
)
