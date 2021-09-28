load("@rules_cc//cc:defs.bzl", "cc_library")

cc_library(
    name = "gnuefi",
    srcs = glob(["lib/*.c", "lib/runtime/*.c", "lib/x86_64/*.c"]),
    hdrs = glob(["inc/**/*.h"]),
    # Consumers expect this to be a system library, so add it to the transitive include paths
    includes = ["inc", "inc/x86_64"],
    visibility = ["//visibility:public"],
)
