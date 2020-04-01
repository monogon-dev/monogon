workspace(name = "nexantic")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")

# Load skylib

http_archive(
    name = "bazel_skylib",
    sha256 = "97e70364e9249702246c0e9444bccdc4b847bed1eb03c5a3ece4f83dfe6abc44",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.2/bazel-skylib-1.0.2.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.0.2/bazel-skylib-1.0.2.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

# Assert minimum Bazel version

load("@bazel_skylib//lib:versions.bzl", "versions")

versions.check(minimum_bazel_version = "2.2.0")

# Go and Gazelle

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "e6a6c016b0663e06fa5fccf1cd8152eab8aa8180c583ec20c872f4f9953a7ac5",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.22.1/rules_go-v0.22.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.22.1/rules_go-v0.22.1.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "86c6d481b3f7aedc1d60c1c211c6f76da282ae197c3b3160f54bd3a8f847896f",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# golang.org/x/sys is overridden by the go_rules protobuf dependency -> declare it first, since
# we need a newer version of it for the netlink package which would fail to compile otherwise.
load("@bazel_gazelle//:deps.bzl", "go_repository")

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:ZtoklVMHQy6BFRHkbG6JzK+S6rX82//Yeok1vMlizfQ=",
    version = "v0.0.0-20191018095205-727590c5006e",
)

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.13",
    nogo = "@//:nogo_vet",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

# Load Gazelle-generated local dependencies

# gazelle:repository_macro third_party/go/repositories.bzl%go_repositories
load("//third_party/go:repositories.bzl", "go_repositories")

go_repositories()

# Protobuf

http_archive(
    name = "com_google_protobuf",
    sha256 = "758249b537abba2f21ebc2d02555bf080917f0f2f88f4cbe2903e0e28c4187ed",
    strip_prefix = "protobuf-3.10.0",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.10.0.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# third_party external repositories
load("//third_party/linux:external.bzl", "linux_external")

linux_external(
    name = "linux",
    version = "5.6",
)

load("//third_party/edk2:external.bzl", "edk2_external")

edk2_external(name = "edk2")

load("//third_party/kubernetes:external.bzl", "kubernetes_external")

kubernetes_external(
    name = "kubernetes",
    version = "1.16.4",
)

load("//third_party/musl:external.bzl", "musl_external")

musl_external(
    name = "musl",
    version = "1.1.24",
)

load("//third_party/util-linux:external.bzl", "util_linux_external")

util_linux_external(
    name = "util_linux",
    version = "2.34",
)

load("//third_party/xfsprogs:external.bzl", "xfsprogs_external")

xfsprogs_external(
    name = "xfsprogs",
    version = "5.2.1",
)

load("@kubernetes//build:workspace_mirror.bzl", "mirror")

http_archive(
    name = "io_k8s_repo_infra",
    sha256 = "f6d65480241ec0fd7a0d01f432938b97d7395aeb8eefbe859bb877c9b4eafa56",
    strip_prefix = "repo-infra-9f4571ad7242bf3ec4b47365062498c2528f9a5f",
    urls = mirror("https://github.com/kubernetes/repo-infra/archive/9f4571ad7242bf3ec4b47365062498c2528f9a5f.tar.gz"),
)

register_toolchains("//:host_python")
