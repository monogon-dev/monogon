workspace(name = "dev_source_monogon")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")

# Assert minimum Bazel version

load("@bazel_skylib//lib:versions.bzl", "versions")

versions.check(minimum_bazel_version = "7.0.0")

# Register our custom CC toolchains. Order matters - more specific toolchains must be registered first.
# (host_cc_toolchain won't care about //build/platforms/linkmode, but musl_host_toolchain won't
# match anything unless its linkmode is set).
register_toolchains("//build/toolchain/musl-host-gcc:musl_host_toolchain")

register_toolchains("//build/toolchain/llvm-efi:efi_k8_toolchain")

register_toolchains("//build/toolchain:host_cc_toolchain")

# Go and Gazelle
http_archive(
    name = "io_bazel_rules_go",
    patch_args = ["-p1"],
    patches = [
        "//third_party/go/patches:rules_go_absolute_embedsrc.patch",
    ],
    sha256 = "80a98277ad1311dacd837f9b16db62887702e9f1d1c4c9f796d0121a46c8e184",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.46.0/rules_go-v0.46.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.46.0/rules_go-v0.46.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
    patch_args = ["-p1"],
    patches = [
        "//third_party/gazelle:add-prepatching.patch",
    ],
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("//third_party/go:repositories.bzl", "go_repositories")

# gazelle:repository_macro third_party/go/repositories.bzl%go_repositories
# gazelle:repository go_repository name=dev_source_monogon importpath=source.monogon.dev
go_repositories()

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.22.1",
    nogo = "@dev_source_monogon//build/analysis:nogo",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

# Load Gazelle-generated local dependencies
gazelle_dependencies()

# Protobuf

http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# GRPC
http_archive(
    name = "com_github_grpc_grpc",
    integrity = "sha256-Q3BouLd307M52pTTSY8dwgZCrJv6dttDq91SIYaxVCs=",
    patch_args = ["-p1"],
    patches = ["//third_party/go/patches:grpc_extra_deps.patch"],
    strip_prefix = "grpc-1.60.0",
    urls = [
        "https://github.com/grpc/grpc/archive/v1.60.0.tar.gz",
    ],
)

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")

grpc_extra_deps()

# Rust rules
http_archive(
    name = "rules_rust",
    integrity = "sha256-ZQGWDD5NoySV0eEAfe0HaaU0yxlcMN6jaqVPnYo/A2E=",
    patch_args = ["-p1"],
    patches = [
        "//third_party:rust-uefi-platform.patch",
        "//third_party:rust-prost-nostd.patch",
        "//third_party:rust-reproducibility.patch",
    ],
    urls = ["https://github.com/bazelbuild/rules_rust/releases/download/0.38.0/rules_rust-v0.38.0.tar.gz"],
)

load("@rules_rust//rust:repositories.bzl", "rules_rust_dependencies", "rust_register_toolchains")

rules_rust_dependencies()

rust_register_toolchains(
    edition = "2021",
    extra_target_triples = [
        "x86_64-unknown-uefi",
    ],
    versions = ["1.71.0"],
)

load("@rules_rust//crate_universe:repositories.bzl", "crate_universe_dependencies")

crate_universe_dependencies()

load("@rules_rust//crate_universe:defs.bzl", "crates_repository")

crates_repository(
    name = "crate_index",
    cargo_lockfile = "//third_party/rust:Cargo.lock",
    generate_binaries = True,
    lockfile = "//third_party/rust:Cargo.Bazel.lock",
    manifests = ["//third_party/rust:Cargo.toml"],
)

load("@crate_index//:defs.bzl", "crate_repositories")

crate_repositories()

crates_repository(
    name = "crate_index_efi",
    cargo_lockfile = "//third_party/rust_efi:Cargo.lock",
    generate_binaries = True,
    lockfile = "//third_party/rust_efi:Cargo.Bazel.lock",
    manifests = ["//third_party/rust_efi:Cargo.toml"],
    supported_platform_triples = [
        "x86_64-unknown-linux-gnu",
        "x86_64-unknown-uefi",
    ],
)

load("@crate_index_efi//:defs.bzl", crate_efi_repositories = "crate_repositories")

crate_efi_repositories()

load("@rules_rust//proto/prost:repositories.bzl", "rust_prost_dependencies")

rust_prost_dependencies()

load("@rules_rust//proto/prost:transitive_repositories.bzl", "rust_prost_transitive_repositories")

rust_prost_transitive_repositories()

register_toolchains("//build/rust:prost_efi_toolchain")

# third_party external repositories
load("//third_party/linux:external.bzl", "linux_external")

linux_external(
    name = "linux",
    version = "6.6.30",
)

load("//third_party/linux-firmware:external.bzl", "linux_firmware_external")

linux_firmware_external(
    name = "linux-firmware",
    version = "20240513",
)

load("//third_party/intel_ucode:external.bzl", "intel_ucode_external")

intel_ucode_external(
    name = "intel_ucode",
    version = "20231114",
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

load("//third_party/pixman:external.bzl", "pixman_external")

pixman_external(
    name = "pixman",
    version = "0.40.0",
)

load("//third_party/uring:external.bzl", "uring_external")

uring_external(
    name = "uring",
    version = "2.0",
)

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "gperf",
    commit = "de9373c2d48a3edf29862eb8be44764a7f7d24c6",
    remote = "https://github.com/monogon-dev/gperf.git",
    shallow_since = "1615306886 +0100",
)

load("//third_party/seccomp:external.bzl", "seccomp_external")

seccomp_external(
    name = "seccomp",
    version = "2.5.1",
)

load("//third_party/glib:external.bzl", "glib_external")

glib_external(
    name = "glib",
    version = "2.67.5",
)

load("//third_party/qemu:external.bzl", "qemu_external")

qemu_external(
    name = "qemu",
    version = "5.2.0",
)

load("//third_party/chrony:external.bzl", "chrony_external")

chrony_external(
    name = "chrony",
)

load("//third_party/cap:external.bzl", "cap_external")

cap_external(
    name = "cap",
    version = "1.2.55",
)

load("//third_party/gnuefi:external.bzl", "gnuefi_external")

gnuefi_external(
    name = "gnuefi",
    version = "3.0.14",
)

load("//third_party/efistub:external.bzl", "efistub_external")

efistub_external(
    name = "efistub",
    # Developed in the systemd monorepo, pinned to master as there have been a bunch of critical fixes for the
    # EFI stub since 249.
    version = "3542da2442d8b29661b47c42ad7e5fa9bc8562ec",
)

load("//third_party/libpg_query:external.bzl", "libpg_query_external")

libpg_query_external(
    name = "libpg_query",
    version = "15-4.2.3",
)

register_toolchains("//:host_python")

# Derived from Mozilla NSS, currently needed for containerd to be able to pull images
http_file(
    name = "cacerts",
    sha256 = "1794c1d4f7055b7d02c2170337b61b48a2ef6c90d77e95444fd2596f4cac609f",
    urls = ["https://curl.se/ca/cacert-2024-03-11.pem"],
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

load("//third_party/dosfstools:external.bzl", "dosfstools_external")

dosfstools_external(
    name = "com_github_dosfstools_dosfstools",
    version = "c888797b1d84ffbb949f147e3116e8bfb2e145a7",
)

# Load musl toolchain Metropolis sysroot tarball into external repository.
load("//build/toolchain/musl-host-gcc:sysroot.bzl", "musl_sysroot_repositories")

musl_sysroot_repositories()

# CockroachDB binary used for tests.
#
# WARNING: Not distributed under an OSI certified license. Must only be used in
# tests, not be redistributed!
http_archive(
    name = "cockroach",
    build_file_content = """
exports_files([
    "cockroach"
])
""",
    sha256 = "0821cff5770400fb94c8b6c2ab338d96f4114fbf2b3206bc8a6dcf62f9c0f4ea",
    strip_prefix = "cockroach-v22.1.6.linux-amd64",
    urls = [
        # TODO: select() to pick other host architectures.
        "https://binaries.cockroachdb.com/cockroach-v22.1.6.linux-amd64.tgz",
    ],
)

# CockroachDB repository used for linter passes.
http_archive(
    name = "com_github_cockroachdb_cockroach",
    integrity = "sha256-3xYgvXmuPvrGgtSzfoK/K9p/FCH0eMZywAAL10A41k0=",
    strip_prefix = "cockroach-23.2.4",
    urls = [
        "https://github.com/cockroachdb/cockroach/archive/v23.2.4.tar.gz",
    ],
)

# bazeldnf is used to generate our sandbox root.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazeldnf",
    sha256 = "cd75fbbad6f191c26b036132d57ca731cce067e9330306a8a2beb3e51af991a8",
    urls = [
        "https://github.com/rmohr/bazeldnf/releases/download/v0.5.8/bazeldnf-v0.5.8.tar.gz",
    ],
)

load("@bazeldnf//:deps.bzl", "bazeldnf_dependencies")

bazeldnf_dependencies()

load("//third_party/sandboxroot:repositories.bzl", "sandbox_dependencies")

sandbox_dependencies()

# Used by tests in cloud/takeover
http_file(
    name = "debian_11_cloudimage",
    sha256 = "14caeec68ba3129a115a9b57396d08dc0973cc9f569ce049232d7d15d768ad41",
    urls = [
        "https://cloud.debian.org/images/cloud/bullseye/20230124-1270/debian-11-genericcloud-amd64-20230124-1270.qcow2",
    ],
)

# Used to include staticcheck as nogo analyzer
http_archive(
    name = "com_github_sluongng_nogo_analyzer",
    integrity = "sha256-p0peRHUdKS0XvYeeWqi0C6qUtdwvBD3x46y7PiPq0HM=",
    strip_prefix = "nogo-analyzer-0.0.2",
    urls = [
        "https://github.com/sluongng/nogo-analyzer/archive/refs/tags/v0.0.2.tar.gz",
    ],
)

git_repository(
    name = "boringssl",
    commit = "d7278cebad5b8eda0901246f2215344cffece4f4",
    remote = "https://boringssl.googlesource.com/boringssl",
)

load("//third_party/libtpms:external.bzl", "libtpms_external")

libtpms_external(
    name = "libtpms",
    version = "93a827aeccd3ab2178281571b1545dcfffa2991b",
)

load("//third_party/swtpm:external.bzl", "swtpm_external")

swtpm_external(
    name = "swtpm",
    version = "0c9a6c4a12a63b86ab472e69e95bd75853d4fa96",
)
