musl-host-gcc
=============

musl-host-gcc is a Bazel C++ toolchain that uses the sandbox sysroot gcc in combination with a pre-built musl, musl headers, and Linux headers.

It is currently used to build the few C binaries we need on Metropolis nodes.

At some point, this toolchain should be improved to directly consume a static compiler toolchain and sysroot, so we can eventually get rid of the sandbox (like Aspect's [gcc-toolchain](https://github.com/aspect-build/gcc-toolchain) is doing).

Usage
-----

To use this toolchain explicitly while building a `cc_binary`, do:

    bazel build --platforms=//build/platforms:linux_arm64_static //foo/bar

During an actual build however, the right toolchain should be selected using transitions
or other configuration mechanisms.

Building Toolchain Sysroot Tarball
----------------------------------

The toolchain's musl/linux components are currently built ahead of time and committed to this repository as `//build/toolchain/musl-host-gcc/toolchain.tar.xz`. This is the 'sysroot' tarball, that contains all headers and libraries required to build for Metropolis nodes.

To build this tarball, run the following commands:

    bazel build //build/toolchain/musl-host-gcc/sysroot
    cp -f bazel-bin/build/toolchain/musl-host-gcc/sysroot/sysroot.tar.xz build/toolchain/musl-host-gcc/sysroot.tar.xz

As a temporary hack the compiler-specific headers of our current development container have been manually merged in. This is expected to be replaced by a proper LLVM-based toolchain.

Internals
---------

The toolchain is implemented in the following way:

1. `//build/toolchain/musl-host-gcc/sysroot` is used to build `//build/toolchain/musl-host-gcc/sysroot.tar.xz` which is a tarball that contains all include and binary library files for building against musl for Metropolis nodes (x86\_64 / k8) - these are musl headers, musl libraries, and linux headers. This tarball is committed to source control.
1. When building a target that uses the toolchain, the `sysroot.tar.xz` tarball is extracted into an external repository `@musl_sysroot`, via `sysroot.bzl` and `sysroot_repository.bzl`.
1. A toolchain config is built using `//build/toolchain:cc_toolchain_config.bzl`, which points at `gcc-wrapper.sh` as its gcc entrypoint. `gcc-wrapper.sh` expects to be able to call the host gcc with `musl.spec`.
1. A toolchain is defined in `//build/toolchain/musl-host-gcc:musl_host_toolchain` with a `//build/platforms/linkmode:musl-static` constraint, which is selected by the `//build/platforms:linux_amd64_static` platform.

Quirks
------

As mentioned above, the musl sysroot is kept in a tarball in this repository. This is obviously suboptimal, but on the other hand gives us an effectively pre-built part of a toolchain. In the future, once we have a hermetic toolchain, a similar tarball might actually contain a fully hermetic toolchain pre-built for k8.
