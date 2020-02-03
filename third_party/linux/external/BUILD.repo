filegroup(
    name = "all",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)

load("@rules_cc//cc:defs.bzl", "cc_binary")

# Build gen_init_cpio separately for the initramfs generation stage
cc_binary(
    name = "gen_init_cpio",
    srcs = ["usr/gen_init_cpio.c"],
    visibility = ["//visibility:public"],
)
