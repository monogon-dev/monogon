workspace(name = "nexantic")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
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
    # Pin slightly above 0.23.0 (and 0.23.1 prerelease at time of writing) to pull in this commit.
    # This fixes https://github.com/bazelbuild/rules_go/issues/2499, which started manifesting at 0.22.4.
    name = "io_bazel_rules_go",
    sha256 = "89501e6e6ae6308e82239f0c8e53dceaa428ad8de471a7f8be8b99a1717bb7d8",
    strip_prefix = "rules_go-c07100d793fc0cdb20bc4f0361c1d53987ba259b",
    urls = [
        "https://github.com/bazelbuild/rules_go/archive/c07100d793fc0cdb20bc4f0361c1d53987ba259b.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    patch_args = ["-p1"],
    patches = [
        "//third_party/gazelle:add-prepatching.patch",
    ],
    sha256 = "bfd86b3cbe855d6c16c6fce60d76bd51f5c8dbc9cfcaef7a2bb5c1aafd0710e8",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.0/bazel-gazelle-v0.21.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.0/bazel-gazellev0.21.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# golang.org/x/sys is overridden by the go_rules protobuf dependency -> declare it first, since
# we need a newer version of it for the netlink package which would fail to compile otherwise.
load("@bazel_gazelle//:deps.bzl", "go_repository")

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:q9u40nxWT5zRClI/uU9dHCiYGottAg6Nzz4YUQyHxdA=",
    version = "v0.0.0-20190927073244-c990c680b611",
)

# we also pin github.com/golang/protobuf to 1.3.2, because we use gRPC 1.26 (can bump to 1.27
# once https://github.com/etcd-io/etcd/issues/11563 is resolved and merged.

go_repository(
    name = "com_github_golang_protobuf",
    build_file_proto_mode = "disable_global",
    commit = "6c65a5562fc06764971b7c5d05c76c75e84bdbf7",
    importpath = "github.com/golang/protobuf",
    patch_args = ["-p1"],
    patches = [
        "@io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch",
    ],
)

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.14",
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

# Build packages
http_archive(
    name = "rules_pkg",
    sha256 = "5bdc04987af79bd27bc5b00fe30f59a858f77ffa0bd2d8143d5b31ad8b1bd71c",
    url = "https://github.com/bazelbuild/rules_pkg/releases/download/0.2.0/rules_pkg-0.2.0.tar.gz",
)

# third_party external repositories
load("//third_party/linux:external.bzl", "linux_external")

linux_external(
    name = "linux",
    version = "5.6",
)

load("//third_party/edk2:external.bzl", "edk2_external")

edk2_external(name = "edk2")

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

register_toolchains("//:host_python")

# python dependencies. Currently we don't use Python, but some of our deps (ie. gvisor) do expect @pydeps// to exist, even
# if it's not being used.

load("@rules_python//python:pip.bzl", "pip_import")

pip_import(
    name = "pydeps",
    requirements = "//third_party/py:requirements.txt",
)

load("@pydeps//:requirements.bzl", "pip_install")

pip_install()

# same for gvisor/rules_docker.

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "14ac30773fdb393ddec90e158c9ec7ebb3f8a4fd533ec2abbfd8789ad81a284b",
    strip_prefix = "rules_docker-0.12.1",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.12.1/rules_docker-v0.12.1.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    go_image_repos = "repositories",
)

go_image_repos()

# Derived from Mozilla NSS, currently needed for containerd to be able to pull images
http_file(
    name = "cacerts",
    sha256 = "adf770dfd574a0d6026bfaa270cb6879b063957177a991d453ff1d302c02081f",
    urls = ["https://curl.haxx.se/ca/cacert-2020-01-01.pem"],
)

# lz4, the library and the tool.
http_archive(
    name = "com_github_lz4_lz4",
    patch_args = ["-p1"],
    patches = ["//third_party/lz4:build.patch"],
    sha256 = "658ba6191fa44c92280d4aa2c271b0f4fbc0e34d249578dd05e50e76d0e5efcc",
    strip_prefix = "lz4-1.9.2",
    urls = ["https://github.com/lz4/lz4/archive/v1.9.2.tar.gz"],
)

# qboot bootloader for MicroVMs
http_archive(
    name = "com_github_bonzini_qboot",
    build_file = "//third_party/qboot:qboot.bzl",
    sha256 = "a643b2486fbee57b969659d408984094ca9afa1a048317dd3f5d3022e47213e8",
    strip_prefix = "qboot-a5300c4949b8d4de2d34bedfaed66793f48ec948",
    urls = ["https://github.com/bonzini/qboot/archive/a5300c4949b8d4de2d34bedfaed66793f48ec948.tar.gz"],
)
