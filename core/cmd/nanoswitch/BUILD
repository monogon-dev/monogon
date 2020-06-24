load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//core/build:def.bzl", "smalltown_initramfs")

go_library(
    name = "go_default_library",
    srcs = ["nanoswitch.go"],
    importpath = "git.monogon.dev/source/nexantic.git/core/cmd/nanoswitch",
    visibility = ["//visibility:private"],
    deps = [
        "//core/internal/common:go_default_library",
        "//core/internal/common/supervisor:go_default_library",
        "//core/internal/launch:go_default_library",
        "//core/internal/network/dhcp:go_default_library",
        "@com_github_google_nftables//:go_default_library",
        "@com_github_google_nftables//expr:go_default_library",
        "@com_github_insomniacslk_dhcp//dhcpv4:go_default_library",
        "@com_github_insomniacslk_dhcp//dhcpv4/server4:go_default_library",
        "@com_github_vishvananda_netlink//:go_default_library",
        "@org_golang_x_sys//unix:go_default_library",
        "@org_uber_go_zap//:go_default_library",
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
