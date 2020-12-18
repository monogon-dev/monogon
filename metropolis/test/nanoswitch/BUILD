load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//metropolis/node/build:def.bzl", "smalltown_initramfs")

go_library(
    name = "go_default_library",
    srcs = ["nanoswitch.go"],
    importpath = "git.monogon.dev/source/nexantic.git/metropolis/test/nanoswitch",
    visibility = ["//visibility:private"],
    deps = [
        "//metropolis/node:go_default_library",
        "//metropolis/node/common/supervisor:go_default_library",
        "//metropolis/node/core/network/dhcp4c:go_default_library",
        "//metropolis/node/core/network/dhcp4c/callback:go_default_library",
        "//metropolis/test/launch:go_default_library",
        "@com_github_google_nftables//:go_default_library",
        "@com_github_google_nftables//expr:go_default_library",
        "@com_github_insomniacslk_dhcp//dhcpv4:go_default_library",
        "@com_github_insomniacslk_dhcp//dhcpv4/server4:go_default_library",
        "@com_github_vishvananda_netlink//:go_default_library",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)

go_binary(
    name = "nanoswitch",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
)

smalltown_initramfs(
    name = "initramfs",
    files = {
        ":nanoswitch": "/init",

        # CA Certificate bundle
        "@cacerts//file": "/etc/ssl/cert.pem",
    },
    visibility = ["//visibility:public"],
)
