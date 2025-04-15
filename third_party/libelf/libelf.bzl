load("@rules_cc//cc:defs.bzl", "cc_library")

filegroup(
    name = "all",
    srcs = glob(["**/*"]),
    visibility = ["//visibility:public"],
)

cc_library(
    name = "elf",
    srcs = glob(["src/*.c", "src/*.h"]),
    hdrs = glob(["include/*.h"]),
    textual_hdrs = glob(["src/*.c"]),
    copts = [
        "-I{path}/include".format(path = package_relative_label(":all").workspace_root),
        "-I{path}/src".format(path = package_relative_label(":all").workspace_root),
    ],
    local_defines = [
        "HAVE_CONFIG_H",
    ],
    linkstatic = True,
    deps = [
        "@zlib//:zlib",
        "@zstd//:zstd",
    ],
    visibility = ["//visibility:public"],
)
