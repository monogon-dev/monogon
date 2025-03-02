load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@rules_pkg//:mappings.bzl", "pkg_files")
load("@rules_pkg//:pkg.bzl", "pkg_zip")
load("//osbase/build/mkerofs:def.bzl", "erofs_image")
load("//osbase/build/mkpayload:def.bzl", "efi_unified_kernel_image")
load("//osbase/build/mkverity:def.bzl", "verity_image")

# Macro for generating multiple TestOS instances to check if the updater works.
# buildifier: disable=unnamed-macro
def testos(variant):
    erofs_image(
        name = "rootfs_" + variant,
        files = {
            "/init": ":testos_" + variant,
            "/etc/resolv.conf": "//osbase/net/dns:resolv.conf",
        },
        fsspecs = [
            "//osbase/build:earlydev.fsspec",
            ":rootfs.fsspec",
        ],
    )

    verity_image(
        name = "verity_rootfs_" + variant,
        source = ":rootfs_" + variant,
        visibility = ["//metropolis/node/core/update/e2e:__pkg__"],
    )

    efi_unified_kernel_image(
        name = "kernel_efi_" + variant,
        cmdline = "console=ttyS0 quiet rootfstype=erofs init=/init loadpin.exclude=kexec-image,kexec-initramfs",
        kernel = "//third_party/linux",
        verity = ":verity_rootfs_" + variant,
        visibility = ["//metropolis/node/core/update/e2e:__pkg__"],
    )

    # An intermediary "bundle" format until we finalize the actual bundle format. This is NOT stable until migrated
    # to the actual bundle format.
    # TODO(lorenz): Replace this
    pkg_files(
        name = "testos_bundle_files_" + variant,
        srcs = [
            ":kernel_efi_" + variant,
            ":verity_rootfs_" + variant,
        ],
        renames = {
            ":kernel_efi_" + variant: "kernel_efi.efi",
            ":verity_rootfs_" + variant: "verity_rootfs.img",
        },
    )
    pkg_zip(
        name = "testos_bundle_" + variant,
        srcs = [
            ":testos_bundle_files_" + variant,
        ],
        visibility = ["//metropolis/node/core/update/e2e:__pkg__"],
    )

    go_binary(
        name = "testos_" + variant,
        embed = [":testos_lib"],
        visibility = ["//visibility:public"],
        x_defs = {"source.monogon.dev/metropolis/node/core/update/e2e/testos.Variant": variant.upper()},
    )
