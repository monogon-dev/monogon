musl-host-gcc
=============

musl-host-gcc is a Bazel C++ toolchain that uses the machine's host gcc in combination with a pre-built musl, musl headers, and Linux headers.

It is currently used to build the few C binaries we need in Smalltown' runtime.

At some point, this toolchain should be replaced by a fully hermetic toolchain that doesn't depend on the host environment.

Usage
-----

To use this toolchain explicitely while building a `cc_binary`, do:

    bazel build --crosstool_top=//build/toolchain/musl-host-gcc:musl_host_cc_suite //foo/bar

During an actual build however, the right toolchain should be selected using aspects or other Bazel configurability features, instead of a hardcoded `--crosstool_top`.

Building Toolchain Sysroot Tarball
----------------------------------

The toolchain's musl/linux components are currently built ahead of time and committed to this repository as `//build/toolchain/musl-host-gcc/toolchain.tar.xz`. This is the 'sysroot' tarball, that contains all headers and libraries required to build against Smalltown.

To build this tarball, run the following commands:

    bazel build //build/toolchain/musl-host-gcc/sysroot
    cp -f bazel-bin/build/toolchain/musl-host-gcc/sysroot/sysroot.tar.xz build/toolchain/musl-host-gcc/sysroot.tar.xz

Internals
---------

The toolchain is implemented in the following way:

1. `//build/toolchain/musl-host-gcc/sysroot` is used to build `//build/toolchain/musl-host-gcc/sysroot.tar.xz` which is a tarball that contains all include and binary library files for building against musl for Smalltown (x86\_64 / k8) - thes are musl headers, musl libraries, and linux headers. This tarball is commited to source control.
1. When building a target that uses the toolchain, the `sysroot.tar.xz` tarball is extracted into an external repository `@musl_sysroot`, via `sysroot.bzl` and `sysroot_repository.bzl`.
1. A toolchain config is built using `//build/toolchain:cc_toolchain_config.bzl`, which points at `gcc-wrapper.sh` as its gcc entrypoint. `gcc-wrapper.sh` expects to be able to call the host gcc with `musl.spec`.
1. A toolchain is built in `//build/toolchain/musl-host-gcc:musl_host_cc_suite`, which uses the previously mentioned config, and builds it to contain `gcc-wrapper.sh`, `musl.spec`, and the sysroot tarball.

Quirks
------

As mentioned above, the musl sysroot is kept in a tarball in this repository. This is obviously suboptimal, but on the other hand gives us an effectively pre-built part of a toolchain. In the future, once we have a hermetic toolchain, a similar tarball might actually contain a fully hermetic toolchain pre-built for k8.
