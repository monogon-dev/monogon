load("@aspect_bazel_lib//lib:transitions.bzl", "platform_transition_filegroup")
load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "efistub_bin",
    srcs = [("src/boot/efi/%s" % v) for v in [
        "assert.c",
        "cpio.c",
        "disk.c",
        "graphics.c",
        "console.c",
        "devicetree.c",
        "linux.c",
        "linux_x86.c",
        "initrd.c",
        "measure.c",
        "ticks.c",
        "pe.c",
        "secure-boot.c",
        "splash.c",
        "stub.c",
        "efi-string.c",
        "util.c",
    ]] + glob(["src/boot/efi/*.h", "src/fundamental/*.c", "src/fundamental/*.h"]),
    includes = ["src/fundamental", "src/boot/efi"],
    copts = ["-std=gnu11", "-O2", "-DSD_BOOT", "-DGIT_VERSION=\\\"0.0.0-mngn\\\""],
    deps = ["@gnuefi//:gnuefi"],
    target_compatible_with = [
        "@platforms//os:uefi",
    ],
    visibility = ["//visibility:private"],
)

platform_transition_filegroup(
    name = "efistub",
    srcs = [":efistub_bin"],
    target_platform = select({
        "@platforms//cpu:x86_64": "@//build/platforms:uefi_x86_64",
        "@platforms//cpu:aarch64": "@//build/platforms:uefi_aarch64",
    }),
    visibility = ["//visibility:public"],
)
