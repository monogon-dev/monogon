# sandboxroot

We use [bazeldnf](https://github.com/rmohr/bazeldnf) in order to reproducibly generate a sysroot
to use for Bazel's sandbox from Fedora packages.

bazeldnf is self-contained and requires only a Go toolchain, requiring minimal host dependencies.
This allows us to bootstrap without having to ship a prebuilt sysroot.

## How to update repository and build rules

Add any new packages to [regenerate.sh](./regenerate.sh) and regenerate definitions:

    third_party/sandboxroot/regenerate.sh

This will fetch the latest version of all required packages from Fedora's repos
and update repositories.bzl and BUILD.bazel.

The next time a bazel command is run, the wrapper will pick up the change
and rebuild the sandbox root.
