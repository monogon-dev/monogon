load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# This is a hack to make go modules ignore all of kube-controlplane since Kubernetes is not importable
# and we still need normal go mod tooling to work. Instead we're just depending on our own Kubernetes
# which is already being built with Bazel and thus works fine as a dependency.

genrule(
    name = "hack_ignore",
    srcs = [
        "main.go",
    ],
    outs = [
        "main_patched.go",
    ],
    cmd = """
    sed '/+build ignore/d' $(location main.go) > "$@"
    """,
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [":main_patched.go"],
    importpath = "git.monogon.dev/source/nexantic.git/core/cmd/kubemaster",
    visibility = ["//visibility:private"],
    deps = [
        "@kubernetes//cmd/kube-apiserver/app:go_default_library",
        "@kubernetes//cmd/kube-controller-manager/app:go_default_library",
        "@kubernetes//cmd/kube-scheduler/app:go_default_library",
        "@kubernetes//staging/src/k8s.io/component-base/cli/flag:go_default_library",
        "@kubernetes//staging/src/k8s.io/component-base/logs:go_default_library",
        "@kubernetes//staging/src/k8s.io/component-base/metrics/prometheus/restclient:go_default_library",
        "@kubernetes//staging/src/k8s.io/component-base/metrics/prometheus/version:go_default_library",
        "@kubernetes//vendor/github.com/spf13/cobra:go_default_library",
        "@kubernetes//vendor/github.com/spf13/pflag:go_default_library",
    ],
)

go_binary(
    name = "kube-controlplane",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
)
