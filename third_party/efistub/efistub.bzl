load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "efistub",
    srcs = [("src/boot/efi/%s" % v) for v in [
        "assert.c",
        "cpio.c",
        "disk.c",
        "graphics.c",
        "linux.c",
        "measure.c",
        "pe.c",
        "secure-boot.c",
        "splash.c",
        "stub.c",
        "util.c",
    ]] + glob(["src/boot/efi/*.h", "src/fundamental/*.c", "src/fundamental/*.h"]),
    includes = ["src/fundamental"],
    copts = ["-std=gnu99", "-DSD_BOOT", "-DGIT_VERSION=\\\"0.0.0-mngn\\\""],
    deps = ["@gnuefi//:gnuefi"],
    visibility = ["//visibility:public"],
)
