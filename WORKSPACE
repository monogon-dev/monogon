workspace(name = "dev_source_monogon")

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
    name = "io_bazel_rules_go",
    sha256 = "7904dbecbaffd068651916dce77ff3437679f9d20e1a7956bff43826e7645fcc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    patch_args = ["-p1"],
    patches = [
        "//third_party/gazelle:add-prepatching.patch",
    ],
    sha256 = "222e49f034ca7a1d1231422cdb67066b885819885c356673cb1f72f748a3c9d4",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "go_repository")

# golang.org/x/sys is overridden by the go_rules protobuf dependency -> declare it first, since
# we need a newer version of it for the netlink package which would fail to compile otherwise.
go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:q9u40nxWT5zRClI/uU9dHCiYGottAg6Nzz4YUQyHxdA=",
    version = "v0.0.0-20190927073244-c990c680b611",
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
    version = "5.10.4",
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
    version = "2.36.2",
)

load("//third_party/xfsprogs:external.bzl", "xfsprogs_external")

xfsprogs_external(
    name = "xfsprogs",
    version = "5.10.0",
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
    sha256 = "336f711f646974ce51cbd10f0cbaaeac17364d92a9ebbcb88b168088b0df2d58",
    strip_prefix = "rules_docker-622f9383d2a306840814c2aba8af0e5dae63c44e",
    # Pinned to a specific commit because of #1675 and #1676 (Bazel 4 compatibility)
    urls = ["https://github.com/bazelbuild/rules_docker/archive/622f9383d2a306840814c2aba8af0e5dae63c44e.tar.gz"],
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

# ini.h, a tiny ini parser library
http_archive(
    name = "inih",
    build_file = "@//third_party/inih:inih.bzl",
    sha256 = "01b0366fdfdf6363efc070c2f856f1afa33e7a6546548bada5456ad94a516241",
    strip_prefix = "inih-r53",
    urls = ["https://github.com/benhoyt/inih/archive/r53.tar.gz"],
)

# qboot bootloader for MicroVMs
http_archive(
    name = "com_github_bonzini_qboot",
    build_file = "//third_party/qboot:qboot.bzl",
    sha256 = "a643b2486fbee57b969659d408984094ca9afa1a048317dd3f5d3022e47213e8",
    strip_prefix = "qboot-a5300c4949b8d4de2d34bedfaed66793f48ec948",
    urls = ["https://github.com/bonzini/qboot/archive/a5300c4949b8d4de2d34bedfaed66793f48ec948.tar.gz"],
)

# Load musl toolchain Metropolis sysroot tarball into external repository.
load("//build/toolchain/musl-host-gcc:sysroot.bzl", "musl_sysroot_repositories")

musl_sysroot_repositories()
