load("@@//build/utils:template_file.bzl", "template_file")
load("@rules_cc//cc:defs.bzl", "cc_library")

template_file(
    name = "config.h",
    src = ":config.h.in",
    substitutions = {},
)

cc_library(
    name = "urcu",
    srcs = glob(
        [
            "src/*.c",
            "src/*.h",
        ],
    ),
    hdrs = glob(["include/**/*.h"]),
    includes = ["include"],
    local_defines = ["RCU_MEMBARRIER", "_GNU_SOURCE"],
    visibility = ["//visibility:public"],
)
