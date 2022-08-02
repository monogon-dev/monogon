load("@rules_cc//cc:defs.bzl", "cc_binary")
load("@dev_source_monogon//build/utils:template_file.bzl", "template_file")

cc_binary(
    name = "fsck",
    srcs = [
        "src/boot.c",
        "src/boot.h",
        "src/charconv.c",
        "src/charconv.h",
        "src/check.c",
        "src/check.h",
        "src/common.c",
        "src/common.h",
        "src/endian_compat.h",
        "src/fat.c",
        "src/fat.h",
        "src/file.c",
        "src/file.h",
        "src/fsck.fat.c",
        "src/fsck.fat.h",
        "src/io.c",
        "src/io.h",
        "src/lfn.c",
        "src/lfn.h",
        "src/msdos_fs.h",
        ":version.h",
    ],
    copts = ["-DHAVE_ENDIAN_H"],
    visibility = ["//visibility:public"],
    includes = ["."],
)

template_file(
    name = "version.h",
    src = "src/version.h.in",
    substitutions = {
        # ONCHANGE(//third_party/dosfstools:external.bzl): version needs to be kept in sync
        "@PACKAGE_VERSION@": "unstable-2022-07-25",
        "@RELEASE_DATE@": "2022-07-25",
    },
)
