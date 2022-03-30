load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//metropolis/node/build:def.bzl", "node_initramfs")

go_library(
    name = "nanoswitch_lib",
    srcs = ["nanoswitch.go"],
    importpath = "source.monogon.dev/metropolis/test/nanoswitch",
    visibility = ["//visibility:private"],
    deps = [
        "//metropolis/node",
        "//metropolis/node/core/network/dhcp4c",
        "//metropolis/node/core/network/dhcp4c/callback",
        "//metropolis/pkg/logtree",
        "//metropolis/pkg/supervisor",
        "//metropolis/test/launch",
        "@com_github_google_nftables//:nftables",
        "@com_github_google_nftables//expr",
        "@com_github_insomniacslk_dhcp//dhcpv4",
        "@com_github_insomniacslk_dhcp//dhcpv4/server4",
        "@com_github_vishvananda_netlink//:netlink",
        "@org_golang_x_sys//unix",
    ],
)

go_binary(
    name = "nanoswitch",
    embed = [":nanoswitch_lib"],
    pure = "on",
    visibility = ["//visibility:private"],
)

node_initramfs(
    name = "initramfs",
    files = {
        ":nanoswitch": "/init",

        # CA Certificate bundle
        "@cacerts//file": "/etc/ssl/cert.pem",
    },
    visibility = ["//metropolis/test:__subpackages__"],
)
