load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//metropolis/node/build:def.bzl", "erofs_image")
load("//metropolis/node/build:efi.bzl", "efi_unified_kernel_image")
load("@rules_pkg//:pkg.bzl", "pkg_zip")

erofs_image(
    name = "rootfs",
    files = {
        ":testos": "/init",
    },
)

efi_unified_kernel_image(
    name = "kernel_efi",
    cmdline = "loglevel=0 console=ttyS0 root=PARTLABEL=METROPOLIS-SYSTEM rootfstype=erofs init=/init",
    kernel = "//third_party/linux",
)

# An intermediary "bundle" format until we finalize the actual bundle format. This is NOT stable until migrated
# to the actual bundle format.
# TODO(lorenz): Replace this
pkg_zip(
    name = "testos_bundle",
    srcs = [
        ":kernel_efi",
        ":rootfs",
    ],
    visibility = ["//metropolis/installer/test:__subpackages__"],
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "source.monogon.dev/metropolis/installer/test/testos",
    visibility = ["//visibility:private"],
    deps = ["@org_golang_x_sys//unix:go_default_library"],
)

go_binary(
    name = "testos",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
