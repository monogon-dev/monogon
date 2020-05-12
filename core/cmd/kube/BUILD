load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "git.monogon.dev/source/nexantic.git/core/cmd/kube",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_spf13_cobra//:go_default_library",
        "@com_github_spf13_pflag//:go_default_library",
        "@io_k8s_component_base//cli/flag:go_default_library",
        "@io_k8s_component_base//logs:go_default_library",
        "@io_k8s_component_base//metrics/prometheus/restclient:go_default_library",
        "@io_k8s_component_base//metrics/prometheus/version:go_default_library",
        "@io_k8s_kubernetes//cmd/kube-apiserver/app:go_default_library",
        "@io_k8s_kubernetes//cmd/kube-controller-manager/app:go_default_library",
        "@io_k8s_kubernetes//cmd/kube-scheduler/app:go_default_library",
        "@io_k8s_kubernetes//cmd/kubelet/app:go_default_library",
    ],
)

go_binary(
    name = "kube",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
)
