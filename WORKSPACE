workspace(name = "dev_source_monogon")

# Assert minimum Bazel version
load("@bazel_skylib//lib:versions.bzl", "versions")

versions.check(minimum_bazel_version = "7.2.1")

# third_party external repositories
load("//third_party/linux:external.bzl", "linux_external")

linux_external(
    name = "linux",
    version = "6.6.42",
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
load("//third_party/urcu:external.bzl", "urcu_external")

urcu_external(
    name = "urcu",
    version = "0.14.0",
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

load("//third_party/dosfstools:external.bzl", "dosfstools_external")

dosfstools_external(
    name = "com_github_dosfstools_dosfstools",
    version = "c888797b1d84ffbb949f147e3116e8bfb2e145a7",
)

# Load musl toolchain Metropolis sysroot tarball into external repository.
load("//build/toolchain/musl-host-gcc:sysroot.bzl", "musl_sysroot_repositories")

musl_sysroot_repositories()

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
