load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_path", "nogo")

# gazelle:prefix git.monogon.dev/source/nexantic.git
# gazelle:exclude core/generated
# gazelle:exclude tools.go
# gazelle:exclude core/cmd/kube-controlplane
# gazelle:resolve go k8s.io/client-go/tools/clientcmd @kubernetes//staging/src/k8s.io/client-go/tools/clientcmd:go_default_library
# gazelle:resolve go k8s.io/client-go/tools/clientcmd/api @kubernetes//staging/src/k8s.io/client-go/tools/clientcmd/api:go_default_library
gazelle(name = "gazelle")

# Shortcut for the Go SDK
alias(
    name = "go",
    actual = "@go_sdk//:bin/go",
    visibility = ["//visibility:public"],
)

# Shortcut for kubectl
alias(
    name = "kubectl",
    actual = "@kubernetes//cmd/kubectl:kubectl",
    visibility = ["//visibility:public"],
)

# nogo linters
nogo(
    name = "nogo_vet",
    config = "nogo_config.json",
    visibility = ["//visibility:public"],
    # These deps enable the analyses equivalent to running `go vet`.
    # Passing vet = True enables only a tiny subset of these (the ones
    # that are always correct).
    #
    # You can see the what `go vet` does by running `go doc cmd/vet`.
    deps = [
        "@org_golang_x_tools//go/analysis/passes/asmdecl:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/assign:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/atomic:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/bools:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/buildtag:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/cgocall:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/composite:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/copylock:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/httpresponse:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/loopclosure:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/lostcancel:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/nilfunc:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/printf:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/shift:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/stdmethods:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/structtag:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/tests:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unmarshal:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unreachable:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unsafeptr:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unusedresult:go_tool_library",
    ],
)

# Synthesize a fake GOPATH in bazel-bin/gopath for IDEs without Bazel support
go_path(
    name = "gopath",
    mode = "link",
    deps = [
        # All top-level Go targets that need IDE integration need to be listed here
        "//core/cmd/init",
    ],
)
