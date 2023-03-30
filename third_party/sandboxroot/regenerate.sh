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

  # TPM emulator for testing
  "swtpm-tools"

  # Clang/LLVM (for EFI toolchain)
  "clang"
  "llvm"
  "lld"

  # image_gcp rule
  "tar"

  # ktest
  "qemu-system-x86-core"

  # musl-host-gcc
  "rsync"
  "xz"

  # Packages included to stabilize SAT solution when there are equal scores.
  "fedora-release-identity-container"
  "coreutils-single"
  "curl-minimal"
  "libcurl-minimal"
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

echo > ${DIR}/repositories.bzl.in

# Create new sandbox root
bazel ${BAZEL_ARGS} \
  run //:bazeldnf -- rpmtree \
  --repofile third_party/sandboxroot/repo.yaml \
  --name sandbox \
  --nobest \
  --buildfile third_party/sandboxroot/BUILD.bazel.in \
  --workspace third_party/sandboxroot/repositories.bzl.in \
  ${PKGS[@]}

# Verify package signatures
bazel ${BAZEL_ARGS} run //:bazeldnf -- verify \
  --repofile third_party/sandboxroot/repo.yaml \
  --workspace third_party/sandboxroot/repositories.bzl.in

# Write out repositories.bzl and clean up.
#
# Ideally, bazeldnf would support the format natively:
# https://github.com/rmohr/bazeldnf/issues/26
cat <<EOF > ${DIR}/repositories.bzl
load("@bazeldnf//:deps.bzl", "rpm")

def sandbox_dependencies():
$(cat ${DIR}/repositories.bzl.in | sed 's/^/    /')
EOF

mv ${DIR}/BUILD.bazel.in ${DIR}/BUILD.bazel
rm ${DIR}/repositories.bzl.in
