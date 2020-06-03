load("@bazel_gazelle//:def.bzl", "gazelle")
load("//build/fietsje:def.bzl", "fietsje")
load("@io_bazel_rules_go//go:def.bzl", "go_path", "nogo")

# gazelle:prefix git.monogon.dev/source/nexantic.git
gazelle(name = "gazelle")

fietsje(name = "fietsje")

# Shortcut for the Go SDK
alias(
    name = "go",
    actual = "@go_sdk//:bin/go",
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
        # Disable cgocall because it fails processing com_github_mattn_go_sqlite3 before exclusions are applied
        #"@org_golang_x_tools//go/analysis/passes/cgocall:go_tool_library",
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

load("@rules_python//python:defs.bzl", "py_runtime_pair")

# Python toolchains - just use the host python for now.
# TODO(T649): move to external (nix?) interpreters.
py_runtime(
    name = "host_python3",
    interpreter_path = "/usr/bin/python3",
    python_version = "PY3",
)

py_runtime(
    name = "host_python2",
    interpreter_path = "/usr/bin/python2",
    python_version = "PY2",
)

py_runtime_pair(
    name = "host_python_pair",
    py2_runtime = ":host_python2",
    py3_runtime = ":host_python3",
)

toolchain(
    name = "host_python",
    toolchain = ":host_python_pair",
    toolchain_type = "@rules_python//python:toolchain_type",
)

# Shortcut for kubectl when running through bazel run
# (don't depend on this, it might turn into an env-based PATH shortcut, use
# @io_k8s_kubernetes//cmd/kubectl instead)
alias(
    name = "kubectl",
    actual = "@io_k8s_kubernetes//cmd/kubectl:kubectl",
)

# Shortcut for the Delve debugger for interactive debugging
alias(
    name = "dlv",
    actual = "@com_github_go_delve_delve//cmd/dlv:dlv",
)
