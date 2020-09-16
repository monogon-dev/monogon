load(":cc_toolchain_config.bzl", "host_cc_toolchain_config")

# Toolchain definitions.
#
# We currently define two toolchains:
#
#  - //build/toolchain:host_cc_suite , which is a fully unhermetic host toolchain,
#    that can be used to build tools for the host.
#  - //build/toolchain/musl-host-gcc:musl_host_cc_suite , which combines the host's
#    gcc compiler with a sysroot tarball that targets Smalltown. This can be used to
#    build C libraries/tools for Smalltown.
#

# This file defines //build/toolchain:host_cc_suite.
#
# This is a C++ toolchain that uses GCC from the host at hardcoded paths. We
# can get away with this, as currently the entire build is performed in a known
# container (see: //scripts:create_container.sh). We define this toolchain so
# that we have full control over all configuration of it, which we need as we
# are building some fairly odd C binaries (notably, a qboot bootloader for
# testing).
#
# The host_cc toolchain suite is enabled for all cc_* targets that aren't
# building host tools by setting --crosstool_top in .bazelrc. In the future,
# this should only be triggered by transitions where necessary.
#
# In the future, the host_cc toolchains should be replaced by a hermetic
# toolchain that's built locally, or downloaded from the Internet - as
# github.com/bazelbuild/bazel-toolchains does it. As that's being built, we
# should then also have another toolchain definition for C binaries that
# target Smalltown static binaries, so that mkfs.xfs can be built using native
# cc_* rules, too.
#
# This, and :cc_toolchain_config.bzl is based on the following tutorial:
# https://docs.bazel.build/versions/master/tutorial/cc-toolchain-config.html

package(default_visibility = ["//visibility:public"])

filegroup(name = "empty")

cc_toolchain_suite(
    name = "host_cc_suite",
    toolchains = {
        "k8": ":host_cc_k8_toolchain",
    },
)

cc_toolchain(
    name = "host_cc_k8_toolchain",
    toolchain_identifier = "host-k8-toolchain",
    toolchain_config = ":host_cc_k8_toolchain_config",
    all_files = ":empty",
    compiler_files = ":empty",
    dwp_files = ":empty",
    linker_files = ":empty",
    objcopy_files = ":empty",
    strip_files = ":empty",
    supports_param_files = 0,
)

host_cc_toolchain_config(name = "host_cc_k8_toolchain_config")
