# Checking out and building

## Stamping

The Metropolis OS image is stamped with info from the current git commit (commit hash, commit date and dirty flag).
This is useful when the image is deployed, as you know exactly which version is running in your cluster.
Each time you make a commit or change the dirty state during development, the stamping info changes, forcing a rebuild of the OS image.
This rebuild is quite cheap, since no binaries are rebuilt.
However, it does invalidate cached test results for all end-to-end tests which depend on the OS image.
If you prefer, you can disable stamping.

To disable stamping, pass the `--config=nostamp` flag.
Note that the builtin Bazel flag `--nostamp` does not work in this repo.
To set this flag for all builds, create the file `.bazelrc.user` in the repository root with content `build --config=nostamp`.
