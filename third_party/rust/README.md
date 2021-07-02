Rust dependency management
==========================

You will need cargo-raze installed on your host operating system, as currently building cargo-raze with Bazel [seems to be broken](https://github.com/google/cargo-raze/issues/423) (and pulls in _a lot_ of transitive dependencies).

    $ cargo install cargo-raze

Dependencies are defined in Cargo.toml. Raze is used to lock these into concrete versions (in `//third_party/rust/cargo/Cargo.raze.lock`) and to generate BUILDfiles (in `//third_party/rust/cargo/remote/...` and `//third_party/rust/BUILD.bazel`).

In contrast to Gazelle/go dependencies, the BUILD files for external packages are actually commited into the repository instead of being generated on demand during analysis phase. This makes Raze a bit more noisy in Git history, but vastly speeds up analysis phase, and doesn't rely on an early, pre-analysis Go toolchain that Gazelle relies on.

To relock dependencies and regenerate BUILDfiles:

    $ cd third_party/rust/
    $ cargo raze

For more information on the process, consult the official [cargo-raze documentation](https://github.com/google/cargo-raze).
