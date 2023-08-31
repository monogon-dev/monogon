# Monogon Monorepo

This is the main repository containing the source code for the [Monogon Platform](https://monogon.tech).

*This is pre-release software - take a look, and check back later!*

## Environment

Our build environment is self-contained and requires only minimal host dependencies:

- A Linux machine or VM.
- [Bazelisk](https://github.com/bazelbuild/bazelisk) >= v1.15.0 (or a working Nix environment).
- A reasonably recent kernel with user namespaces enabled.
- Working KVM with access to `/dev/kvm` (if you want to run tests).

Our docs assume that Bazelisk is available as `bazel` on your PATH.

Refer to [SETUP.md](./SETUP.md) for detailed instructions.

## Monogon OS

### Run a single node demo cluster

Build CLI and node image:

    bazel build //metropolis/cli/dbg //:launch --config dbg

Launch an ephemeral test node:

    bazel test //:launch --config dbg --test_output=streamed
    
Run a kubectl command while the test is running:

    bazel-bin/metropolis/cli/dbg/dbg_/dbg kubectl describe node
 
### Test suite

Run full test suite:

    bazel test --config dbg //...
