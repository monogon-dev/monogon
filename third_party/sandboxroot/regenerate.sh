#!/usr/bin/env bash
set -euo pipefail

# Tell wrapper to not touch sandbox
export MONOGON_SYSROOT_REBUILD=1

# Packages to install. Make sure to document the reason for including each package.
PKGS=(
  # Common base toolchain used across the tree.
  "binutils"
  "gcc"
  "python3"
  "python-unversioned-command"
  "glibc-static"

  # Required to build static CGO binaries
  # see monogon-dev/monogon#192
  "libstdc++-static"

  # Kernel build
  "flex"
  "bison"
  "elfutils-libelf-devel"
  "openssl-devel"
  "diffutils"
  "bc"
  "perl"
  "lz4"

  # EDK2
  "libuuid-devel"
  "util-linux"
  "nasm"
  "acpica-tools"

  # patch tool, as used by gazelle
  "patch"

  # Clang/LLVM (for EFI toolchain)
  "clang"
  "llvm"
  "lld"

  # image_gcp rule
  "tar"

  # ktest
  "qemu-system-x86-core"
  "qemu-img"

  # musl-host-gcc
  "rsync"
  "xz"

  # Packages included to stabilize SAT solution when there are equal scores.
  "fedora-release-identity-container"
  "coreutils-single"
  "curl"
  "libcurl"
  "glibc-langpack-en"
  "selinux-policy-minimum"
)

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
REPO=third_party/sandboxroot/repo.yaml
BAZEL_ARGS="--noworkspace_rc --bazelrc ${DIR}/../../.bazelrc.sandboxroot"

# Fetch latest repository metadata
bazel ${BAZEL_ARGS} run //:bazeldnf -- fetch --repofile $REPO

# Write BUILD.bazel template
cat <<EOF > ${DIR}/BUILD.bazel.in
load("@bazeldnf//:deps.bzl", "rpmtree")
load("@bazeldnf//:def.bzl", "bazeldnf")

bazeldnf(
    name = "sandboxroot",
    command = "sandbox",
    tar = ":sandbox",
)

EOF

# Create new sandbox root
bazel ${BAZEL_ARGS} \
  run //:bazeldnf -- rpmtree \
  --repofile third_party/sandboxroot/repo.yaml \
  --name sandbox \
  --nobest \
  --buildfile third_party/sandboxroot/BUILD.bazel.in \
  --to-macro third_party/sandboxroot/repositories.bzl%sandbox_dependencies \
  ${PKGS[@]}

# Verify package signatures
bazel ${BAZEL_ARGS} run //:bazeldnf -- verify \
  --repofile third_party/sandboxroot/repo.yaml \
  --from-macro third_party/sandboxroot/repositories.bzl%sandbox_dependencies

mv ${DIR}/BUILD.bazel.in ${DIR}/BUILD.bazel
