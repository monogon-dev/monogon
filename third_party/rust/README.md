Rust dependency management
==========================

Dependencies are defined in Cargo.toml. Dependency syncing and updating is done in the repository rule which means it’s done during the analysis phase of builds.

To render a new lock file:

    $ CARGO_BAZEL_REPIN=1 bazel sync

For more information on the process, consult the official [rules_rust/crate_universe documentation](https://bazelbuild.github.io/rules_rust/crate_universe.html).
