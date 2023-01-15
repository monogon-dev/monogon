# Monogon Monorepo

This is the main repository containing the source code for the [Monogon Platform](https://monogon.tech).

*This is pre-release software - take a look, and check back later!*

## Environment

Our build environment is self-contained and requires only minimal host dependencies:

- A Linux machine or VM.
- [Bazelisk](https://github.com/bazelbuild/bazelisk) >= v1.15.0
- A reasonably recent kernel with user namespaces enabled.
- Working KVM with access to `/dev/kvm` (if you want to run tests).

Our docs assume that Bazelisk is available as `bazel` on your PATH.

### IntelliJ support

This repository is compatible with the IntelliJ Bazel plugin out of the box, which enables
full autocompletion for external dependencies and generated code.

The following steps are necessary:

- Install Google's [Bazel plugin](https://plugins.jetbrains.com/plugin/8609-bazel) in IntelliJ.
 
- Make sure that Bazel "*Bazel Binary Location*" in Other Settings → Bazel Settings points to Bazelisk.
  
- Use _File → Import Bazel project_... and select your monorepo checkout.

After running the first sync, everything should now resolve in the IDE, including generated code.

## Monogon OS

### Run a single node demo cluster

Build CLI and node image:

    bazel build //metropolis/cli/dbg //:launch -c dbg

Launch an ephemeral test node:

    bazel test //:launch -c dbg --test_output=streamed
    
Run a kubectl command while the test is running:

    bazel-bin/metropolis/cli/dbg/dbg_/dbg kubectl describe node
 
### Test suite

Run full test suite:

    bazel test -c dbg //...
